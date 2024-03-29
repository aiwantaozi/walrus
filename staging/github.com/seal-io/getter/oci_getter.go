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
// TODO(michelia): borrow the code from helm.sh/helm/v3/pkg/getter/http.go
package getter

import (
	"bytes"

	"helm.sh/helm/v3/pkg/getter"
)

// HTTPGetter is the default HTTP(/S) backend handler
type OCIGetter struct {
	getter.OCIGetter
	opts options
}

// Get performs a Get from repo.Getter and returns the body.
func (g *OCIGetter) Get(href string, options ...Option) (*bytes.Buffer, error) {
	for _, opt := range options {
		opt(&g.opts)
	}

	return g.OCIGetter.Get(
		href,
		getter.WithTransport(g.opts.transport),
		getter.WithTimeout(g.opts.timeout),
		getter.WithInsecureSkipVerifyTLS(g.opts.insecureSkipVerifyTLS),
		getter.WithBasicAuth(g.opts.username, g.opts.password),
		getter.WithPassCredentialsAll(g.opts.passCredentialsAll),
		getter.WithTLSClientConfig(g.opts.certFile, g.opts.keyFile, g.opts.caFile),
		getter.WithUserAgent(g.opts.userAgent),
		getter.WithPlainHTTP(g.opts.plainHTTP),
		getter.WithRegistryClient(g.opts.registryClient),
		getter.WithURL(g.opts.url),
		getter.WithUntar())
}
