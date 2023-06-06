/*
Copyright 2020 The cert-manager Authors.

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

package feature

import (
	"k8s.io/apimachinery/pkg/util/runtime"
	"k8s.io/component-base/featuregate"

	utilfeature "github.com/cert-manager/istio-csr/pkg/util/feature"
)

const (
	// alpha: v0.7.2
	//
	// AdditionalAnnotations allows specification of an additional CR annotation
	// TODO: improve this wording
	AdditionalAnnotations featuregate.Feature = "AdditionalAnnotations"
)

func init() {
	runtime.Must(utilfeature.DefaultMutableFeatureGate.Add(defaultIstioCSRFeatureGates))
}

// defaultCertManagerFeatureGates consists of all known cert-manager feature keys.
// To add a new feature, define a key for it above and add it here. The features will be
// available on the cert-manager controller binary.
var defaultIstioCSRFeatureGates = map[featuregate.Feature]featuregate.FeatureSpec{
	AdditionalAnnotations: {Default: false, PreRelease: featuregate.Alpha},
}
