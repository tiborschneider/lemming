// Copyright 2021 Google LLC
//
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at
//
//      http://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and
// limitations under the License.

// Package gnmit is a single-target gNMI collector implementation that can be
// used as an on-device/fake device implementation. It supports the Subscribe RPC
// using the libraries from openconfig/gnmi.
package gnmit

import (
	"context"
	"fmt"
	"net"
	"sync"
	"time"

	log "github.com/golang/glog"
	"github.com/openconfig/gnmi/cache"
	"github.com/openconfig/gnmi/subscribe"
	"github.com/openconfig/lemming/gnmi/gnmistore"
	"github.com/openconfig/ygot/ygot"
	"github.com/openconfig/ygot/ytypes"
	"golang.org/x/exp/slices"
	"google.golang.org/grpc"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/metadata"
	"google.golang.org/grpc/status"

	gpb "github.com/openconfig/gnmi/proto/gnmi"
)

const (
	OpenconfigOrigin = "openconfig"
)

// GNMIServer implements the gNMI server interface.
type GNMIServer struct {
	// The subscribe Server implements only Subscribe for gNMI.
	*subscribe.Server
	c *Collector

	configMu     sync.Mutex
	configSchema *ytypes.Schema

	stateMu     sync.Mutex
	stateSchema *ytypes.Schema
}

// New returns a new collector server implementation that can be registered on
// an existing gRPC server.
//
// - configSchema is the specification of the schema if gnmi.Set on config paths is used.
// - stateSchema is the specification of the schema if gnmi.Set on state paths is used.
// - targetName is the name of the target.
// - sendMeta indicates whether metadata should be sent
func New(ctx context.Context, configSchema *ytypes.Schema, stateSchema *ytypes.Schema, targetName string, sendMeta bool) (*Collector, *GNMIServer, error) {
	c := NewCollector(targetName)
	subscribeSrv, err := c.Start(ctx, sendMeta)
	if err != nil {
		return nil, nil, err
	}

	// Initialize the cache with the input schema root.
	if configSchema != nil {
		if err := SetupSchema(configSchema); err != nil {
			return nil, nil, err
		}
		if err := ygot.PruneConfigFalse(configSchema.RootSchema(), configSchema.Root); err != nil {
			return nil, nil, fmt.Errorf("gnmit: %v", err)
		}
		updateCache(c.cache, configSchema.Root, nil, targetName, OpenconfigOrigin, true)
	}

	if stateSchema != nil {
		if err := SetupSchema(stateSchema); err != nil {
			return nil, nil, err
		}
		updateCache(c.cache, stateSchema.Root, nil, targetName, OpenconfigOrigin, true)
	}

	gnmiserver := &GNMIServer{
		Server:       subscribeSrv, // use the 'subscribe' implementation.
		c:            c,
		configSchema: configSchema,
		stateSchema:  stateSchema,
	}

	return c, gnmiserver, nil
}

// Start starts the collector-backed gNMI server that listens on the specified
// addr (in the form host:port).
//
// It returns the address it is listening on in the form hostname:port or any
// errors encounted whilst setting it up.
func (s *GNMIServer) Start(ctx context.Context, addr string, opts ...grpc.ServerOption) (string, error) {
	// Start gNMI server.
	srv := grpc.NewServer(opts...)
	gpb.RegisterGNMIServer(srv, s)
	// Forward streaming updates to clients.
	// Register listening port and start serving.
	lis, err := net.Listen("tcp", addr)
	if err != nil {
		return "", fmt.Errorf("failed to listen: %v", err)
	}

	go func() {
		if err := srv.Serve(lis); err != nil {
			log.Fatalf("Error while serving gnmi target: %v", err)
		}
	}()
	s.c.stopFn = func() {
		srv.GracefulStop()
	}
	return lis.Addr().String(), nil
}

type populateDefaultser interface {
	PopulateDefaults()
}

// SetupSchema takes in a ygot schema object which it assumes to be
// uninitialized. It initializes and validates it, returning any errors
// encountered.
func SetupSchema(schema *ytypes.Schema) error {
	if !schema.IsValid() {
		return fmt.Errorf("cannot obtain valid schema for GoStructs: %v", schema)
	}

	// Initialize the root with default values.
	schema.Root.(populateDefaultser).PopulateDefaults()
	if err := schema.Validate(); err != nil {
		return fmt.Errorf("default root of input schema fails validation: %v", err)
	}

	return nil
}

// updateCache updates the cache with the difference between the root ->
// dirtyRoot such that if the root represents the cache, then the dirtyRoot
// will represent the cache afterwards.
//
// If root is nil, then it is assumed the cache is empty, and the entirety of
// the dirtyRoot is put into the cache.
func updateCache(cache *cache.Cache, dirtyRoot, root ygot.GoStruct, target, origin string, preferShadowPath bool) error {
	var nos []*gpb.Notification
	if root == nil {
		var err error
		if nos, err = ygot.TogNMINotifications(dirtyRoot, time.Now().UnixNano(), ygot.GNMINotificationsConfig{
			UsePathElem: true,
		}); err != nil {
			return fmt.Errorf("gnmit: %v", err)
		}
	} else {
		n, err := ygot.Diff(root, dirtyRoot, &ygot.DiffPathOpt{PreferShadowPath: preferShadowPath})
		if err != nil {
			return fmt.Errorf("gnmit: error while creating update notification for Set: %v", err)
		}
		n.Timestamp = time.Now().UnixNano()
		nos = append(nos, n)
	}

	return updateCacheNotifs(cache, nos, target, origin)
}

// updateCacheNotifs updates the target cache with the given notifications.
func updateCacheNotifs(cache *cache.Cache, nos []*gpb.Notification, target, origin string) error {
	cacheTarget := cache.GetTarget(target)
	for _, n := range nos {
		n.Prefix = &gpb.Path{Origin: origin, Target: target}
		if n.Prefix.Origin == "" {
			n.Prefix.Origin = OpenconfigOrigin
		}

		var pathsForDelete []string
		for _, path := range n.Delete {
			p, err := ygot.PathToString(path)
			if err != nil {
				return fmt.Errorf("cannot convert deleted path to string: %v", err)
			}
			pathsForDelete = append(pathsForDelete, p)
		}
		log.V(1).Infof("datastore: updating the following values: %+v", n.Update)
		log.V(1).Infof("datastore: deleting the following paths: %+v", pathsForDelete)
		if err := cacheTarget.GnmiUpdate(n); err != nil {
			return err
		}
	}
	return nil
}

// set updates the datastore and intended configuration with the SetRequest,
// allowing read-only values to be updated.
//
// update indicates whether to update the cache with the values from the set
// request.
func set(schema *ytypes.Schema, cache *cache.Cache, target string, req *gpb.SetRequest, preferShadowPath bool) error {
	prevRoot, err := ygot.DeepCopy(schema.Root)
	if err != nil {
		return fmt.Errorf("gnmit: failed to ygot.DeepCopy the cached root object: %v", err)
	}

	success := false

	// Rollback function
	defer func() {
		if !success {
			log.V(1).Infof("Rolling back set request: %v", req)
			schema.Root = prevRoot
		}
	}()

	var unmarshalOpts []ytypes.UnmarshalOpt
	if preferShadowPath {
		unmarshalOpts = append(unmarshalOpts, &ytypes.PreferShadowPath{})
	}
	if err := ytypes.UnmarshalSetRequest(schema, req, unmarshalOpts...); err != nil {
		return fmt.Errorf("gnmit: %v", err)
	}

	if err := schema.Validate(); err != nil {
		return fmt.Errorf("gnmit: invalid SetRequest: %v", err)
	}

	success = true

	if err := updateCache(cache, schema.Root, prevRoot, target, req.Prefix.Origin, preferShadowPath); err != nil {
		return err
	}
	return nil
}

// Set is a prototype for a gNMI Set operation.
// TODO(wenbli): Add unit test.
func (s *GNMIServer) Set(ctx context.Context, req *gpb.SetRequest) (*gpb.SetResponse, error) {
	// Use ConfigMode by default so that external users don't need to set metadata.
	gnmiMode := gnmistore.ConfigMode
	md, ok := metadata.FromIncomingContext(ctx)
	if ok {
		switch {
		case slices.Contains(md.Get(gnmistore.GNMIModeMetadataKey), string(gnmistore.ConfigMode)):
			gnmiMode = gnmistore.ConfigMode
		case slices.Contains(md.Get(gnmistore.GNMIModeMetadataKey), string(gnmistore.StateMode)):
			gnmiMode = gnmistore.StateMode
		}
	}

	switch gnmiMode {
	case gnmistore.ConfigMode:
		s.configMu.Lock()
		defer s.configMu.Unlock()

		log.V(1).Infof("config datastore service received SetRequest: %v", req)
		if s.configSchema == nil {
			return s.UnimplementedGNMIServer.Set(ctx, req)
		}

		// TODO(wenbli): Reject paths that try to modify read-only values.
		// TODO(wenbli): Question: what to do if there are operational-state values in a container that is specified to be replaced or deleted?
		err := set(s.configSchema, s.c.cache, s.c.name, req, true)

		// TODO(wenbli): Currently the SetResponse is not filled.
		return &gpb.SetResponse{
			Timestamp: time.Now().UnixNano(),
		}, err
	case gnmistore.StateMode:
		s.stateMu.Lock()
		defer s.stateMu.Unlock()

		log.V(1).Infof("operational state datastore service received SetRequest: %v", req)
		if s.stateSchema == nil {
			return s.UnimplementedGNMIServer.Set(ctx, req)
		}
		// TODO(wenbli): Reject values that modify config values. We only allow modifying state in this mode.
		if err := set(s.stateSchema, s.c.cache, s.c.name, req, false); err != nil {
			return &gpb.SetResponse{}, status.Errorf(codes.Aborted, "%v", err)
		}

		// This mode is intended to be used internally, and the SetResponse doesn't matter.
		return &gpb.SetResponse{}, nil
	default:
		return nil, status.Errorf(codes.Internal, "incoming gNMI SetRequest must specify a valid GNMIMode via context metadata: %v", md)
	}
}
