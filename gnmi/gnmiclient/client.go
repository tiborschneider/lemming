// Copyright 2022 Google LLC
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

// Package gnmiclient contains a funcs to create gNMI for the local cache.
package gnmiclient

import (
	"context"
	"crypto/tls"
	"fmt"

	gpb "github.com/openconfig/gnmi/proto/gnmi"
	"github.com/openconfig/lemming/gnmi/gnmit"
	"github.com/openconfig/ygnmi/ygnmi"
	"github.com/spf13/viper"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials"
	"google.golang.org/grpc/credentials/local"
)

// localClient is a client that connects to local gNMI cache.
type localClient struct {
	gpb.GNMIClient
	setClient gpb.GNMIClient
}

// Set uses the datastore client for Set, instead of the public cache endpoint.
func (m *localClient) Set(ctx context.Context, in *gpb.SetRequest, opts ...grpc.CallOption) (*gpb.SetResponse, error) {
	return m.setClient.Set(ctx, in, opts...)
}

// NewLocal returns a gNMI client connected to the local gNMI cache and datastore for Set.
func NewLocal() (gpb.GNMIClient, error) {
	setConn, err := grpc.Dial(fmt.Sprintf("unix:///%s", gnmit.DatastoreAddress), grpc.WithTransportCredentials(local.NewCredentials()))
	if err != nil {
		return nil, fmt.Errorf("failed to dial unix socket: %v", err)
	}
	opt := grpc.WithTransportCredentials(local.NewCredentials())
	if viper.GetBool("enable_tls") {
		//nolint:gosec // TODO: figure out long cert handling
		opt = grpc.WithTransportCredentials(credentials.NewTLS(&tls.Config{InsecureSkipVerify: true}))
	}
	cacheConn, err := grpc.Dial(fmt.Sprintf("localhost:%d", viper.GetInt("port")), opt)
	if err != nil {
		return nil, fmt.Errorf("failed to dial cache socket: %v", err)
	}
	return &localClient{
		GNMIClient: gpb.NewGNMIClient(cacheConn),
		setClient:  gpb.NewGNMIClient(setConn),
	}, nil
}

// NewYGNMIClient returns ygnmi client connected to the local gNMI cache.
func NewYGNMIClient() (*ygnmi.Client, error) {
	gClient, err := NewLocal()
	if err != nil {
		return nil, err
	}
	return ygnmi.NewClient(gClient, ygnmi.WithTarget(viper.GetString("target")))
}

// Update updates the configuration at the given query path with the val.
func Update[T any](ctx context.Context, c *ygnmi.Client, q ygnmi.SingletonQuery[T], val T) (*ygnmi.Result, error) {
	return ygnmi.Update[T](ctx, c, &singletonAsConfig[T]{SingletonQuery: q}, val)
}

// Replace replaces the configuration at the given query path with the val.
func Replace[T any](ctx context.Context, c *ygnmi.Client, q ygnmi.SingletonQuery[T], val T) (*ygnmi.Result, error) {
	return ygnmi.Replace[T](ctx, c, &singletonAsConfig[T]{SingletonQuery: q}, val)
}

// Delete deletes the configuration at the given query path.
func Delete[T any](ctx context.Context, c *ygnmi.Client, q ygnmi.SingletonQuery[T], val T) (*ygnmi.Result, error) {
	return ygnmi.Delete[T](ctx, c, &singletonAsConfig[T]{SingletonQuery: q})
}

// BatchUpdate stores an update operation in the SetBatch.
func BatchUpdate[T any](sb *ygnmi.SetBatch, q ygnmi.SingletonQuery[T], val T) {
	ygnmi.BatchUpdate[T](sb, &singletonAsConfig[T]{SingletonQuery: q}, val)
}

// BatchReplace stores an replace operation in the SetBatch.
func BatchReplace[T any](sb *ygnmi.SetBatch, q ygnmi.SingletonQuery[T], val T) {
	ygnmi.BatchReplace[T](sb, &singletonAsConfig[T]{SingletonQuery: q}, val)
}

// BatchDelete stores an delete operation in the SetBatch.
func BatchDelete[T any](sb *ygnmi.SetBatch, q ygnmi.SingletonQuery[T]) {
	ygnmi.BatchDelete[T](sb, &singletonAsConfig[T]{SingletonQuery: q})
}

// singletonAsConfig turns a SingletonQuery into ConfigQuery.
type singletonAsConfig[T any] struct {
	ygnmi.SingletonQuery[T]
}

func (*singletonAsConfig[T]) IsConfig() {}