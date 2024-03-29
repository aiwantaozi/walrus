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

// TODO(michelia): borrow the code from helm.sh/helm/v3/pkg/getter/getter.go

package getter

import (
	"bytes"
	"net/http"
	"time"

	"helm.sh/helm/v3/pkg/registry"
)

// options are generic parameters to be provided to the getter during instantiation.
//
// Getters may or may not ignore these parameters as they are passed in.
type options struct {
	url                   string
	cert                  []byte
	certKey               []byte
	ca                    []byte
	unTar                 bool
	insecureSkipVerifyTLS bool
	plainHTTP             bool
	username              string
	password              string
	token                 string
	passCredentialsAll    bool
	userAgent             string
	version               string
	registryClient        *registry.Client
	timeout               time.Duration
	transport             *http.Transport
}

// Option allows specifying various settings configurable by the user for overriding the defaults
// used when performing Get operations with the Getter.
type Option func(*options)

// WithURL informs the getter the server name that will be used when fetching objects. Used in conjunction with
// WithTLSClientConfig to set the TLSClientConfig's server name.
func WithURL(url string) Option {
	return func(opts *options) {
		opts.url = url
	}
}

// WithBasicAuth sets the request's Authorization header to use the provided credentials
func WithBasicAuth(username, password string) Option {
	return func(opts *options) {
		opts.username = username
		opts.password = password
	}
}

// WithBeaterToken sets the request's Authorization header with bearer token
func WithBeaterToken(token string) Option {
	return func(opts *options) {
		opts.token = token
	}
}

func WithPassCredentialsAll(pass bool) Option {
	return func(opts *options) {
		opts.passCredentialsAll = pass
	}
}

// WithUserAgent sets the request's User-Agent header to use the provided agent name.
func WithUserAgent(userAgent string) Option {
	return func(opts *options) {
		opts.userAgent = userAgent
	}
}

// WithInsecureSkipVerifyTLS determines if a TLS Certificate will be checked
func WithInsecureSkipVerifyTLS(insecureSkipVerifyTLS bool) Option {
	return func(opts *options) {
		opts.insecureSkipVerifyTLS = insecureSkipVerifyTLS
	}
}

// WithTLSClientConfig sets the client auth with the provided credentials.
func WithTLSClientConfig(cert, key, ca []byte) Option {
	return func(opts *options) {
		opts.cert = cert
		opts.certKey = key
		opts.ca = ca
	}
}

// WithTLSClientConfig sets the client auth with the provided credentials.
func WithTLSClientFileConfig(cert, key, ca []byte) Option {
	return func(opts *options) {
		opts.cert = cert
		opts.certKey = key
		opts.ca = ca
	}
}

func WithPlainHTTP(plainHTTP bool) Option {
	return func(opts *options) {
		opts.plainHTTP = plainHTTP
	}
}

// WithTimeout sets the timeout for requests
func WithTimeout(timeout time.Duration) Option {
	return func(opts *options) {
		opts.timeout = timeout
	}
}

func WithTagName(tagname string) Option {
	return func(opts *options) {
		opts.version = tagname
	}
}

func WithRegistryClient(client *registry.Client) Option {
	return func(opts *options) {
		opts.registryClient = client
	}
}

func WithUntar() Option {
	return func(opts *options) {
		opts.unTar = true
	}
}

// WithTransport sets the http.Transport to allow overwriting the HTTPGetter default.
func WithTransport(transport *http.Transport) Option {
	return func(opts *options) {
		opts.transport = transport
	}
}

// Getter is an interface to support GET to the specified URL.
type Getter interface {
	// Get file content by url string
	Get(url string, options ...Option) (*bytes.Buffer, error)
}

// Constructor is the function for every getter which creates a specific instance
// according to the configuration
type Constructor func(options ...Option) (Getter, error)
