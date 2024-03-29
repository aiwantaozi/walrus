/*
Copyright The Helm Authors.

Licensed under the Apache License, Version 2.0 (the "License");
you may not use this file except in compliance with the License.
You may obtain a copy of the License at

    http://www.apache.org/licenses/LICENSE-2.0

Unless required by applicable law or agreed to in writing, software
distributed under the License is distributed on an "AS IS" BASIS,
WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
See the License for the specific language governing permissions and
limitations under the License.
*/

package tlsutil

import (
	"crypto/tls"
	"crypto/x509"

	"github.com/pkg/errors"
)

// NewClientTLS returns tls.Config appropriate for client auth.
func NewClientTLS(certPEMBlock, keyPEMBlock, caPEMBlock []byte, insecureSkipTLSVerify bool) (*tls.Config, error) {
	config := tls.Config{
		InsecureSkipVerify: insecureSkipTLSVerify,
	}

	if len(certPEMBlock) != 0 && len(keyPEMBlock) != 0 {
		cert, err := tls.X509KeyPair(certPEMBlock, keyPEMBlock)
		if err != nil {
			return nil, err
		}
		config.Certificates = []tls.Certificate{cert}
	}

	if len(caPEMBlock) != 0 {
		cp := x509.NewCertPool()
		if !cp.AppendCertsFromPEM(caPEMBlock) {
			return nil, errors.New("failed to append certificates")
		}

		config.RootCAs = cp
	}

	return &config, nil
}
