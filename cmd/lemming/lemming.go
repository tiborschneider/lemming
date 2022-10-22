package main

import (
	"crypto/rand"
	"crypto/rsa"
	"crypto/tls"
	"crypto/x509"
	"encoding/pem"
	"flag"
	"fmt"
	"math/big"
	"net"

	log "github.com/golang/glog"
	"github.com/openconfig/lemming"
	"github.com/openconfig/lemming/sysrib"
	"github.com/spf13/pflag"
	"github.com/spf13/viper"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials"
	"google.golang.org/grpc/credentials/insecure"
)

var (
	port   = pflag.Int("port", 6030, "localhost port to listen to.")
	target = pflag.String("target", "fakedut", "name of the fake target")
	// nolint:unused,varcheck
	enableDataplane = pflag.Bool("enable_dataplane", false, "Controls whether to enable dataplane")
	enableTLS       = pflag.Bool("enable_tls", false, "Controls whether to enable TLS for gNXI services. If enabled and TLS key/cert path unspecified, a generated cert is used.")
)

func main() {
	pflag.CommandLine.AddGoFlagSet(flag.CommandLine)
	pflag.Parse()
	viper.BindPFlags(pflag.CommandLine)

	lis, err := net.Listen("tcp", fmt.Sprintf(":%d", *port))
	if err != nil {
		log.Fatalf("Failed to start listener: %v", err)
	}
	creds, err := newCreds()
	if err != nil {
		log.Fatalf("failed to create credentials: %v", err)
	}

	f, err := lemming.New(lis, *target, sysrib.ZAPI_ADDR, grpc.Creds(creds))
	if err != nil {
		log.Fatalf("Failed to start lemming: %v", err)
	}
	defer f.Stop()

	log.Info("lemming initialization complete")
	select {}
}

// newCreds returns either insecure or tls credentials, depending the enable_tls flag.
// TODO: figure out long term plan for certs, this implementation is here to unblock using Ondatra KNEBind.
func newCreds() (credentials.TransportCredentials, error) {
	if !*enableTLS {
		return insecure.NewCredentials(), nil
	}
	key, err := rsa.GenerateKey(rand.Reader, 2048)
	if err != nil {
		return nil, err
	}
	tmpl := &x509.Certificate{
		SerialNumber: big.NewInt(1234),
	}

	certDer, err := x509.CreateCertificate(rand.Reader, tmpl, tmpl, &key.PublicKey, key)
	if err != nil {
		return nil, err
	}
	certPem := pem.EncodeToMemory(&pem.Block{Type: "CERTIFICATE", Bytes: certDer})

	keyDer, err := x509.MarshalPKCS8PrivateKey(key)
	if err != nil {
		return nil, err
	}
	keyPem := pem.EncodeToMemory(&pem.Block{Type: "PRIVATE KEY", Bytes: keyDer})

	serverCert, err := tls.X509KeyPair(certPem, keyPem)
	if err != nil {
		return nil, err
	}
	return credentials.NewTLS(&tls.Config{
		MinVersion:   tls.VersionTLS13,
		Certificates: []tls.Certificate{serverCert},
	}), nil
}
