/*
Copyright The Kubernetes Authors.

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

// Code generated by informer-gen. DO NOT EDIT.

package v1

import (
	internalinterfaces "github.com/crunchydata/postgres-operator/pkg/generated/informers/externalversions/internalinterfaces"
)

// Interface provides access to all the informers in this group version.
type Interface interface {
	// Pgclusters returns a PgclusterInformer.
	Pgclusters() PgclusterInformer
	// Pgpolicies returns a PgpolicyInformer.
	Pgpolicies() PgpolicyInformer
	// Pgreplicas returns a PgreplicaInformer.
	Pgreplicas() PgreplicaInformer
	// Pgtasks returns a PgtaskInformer.
	Pgtasks() PgtaskInformer
}

type version struct {
	factory          internalinterfaces.SharedInformerFactory
	namespace        string
	tweakListOptions internalinterfaces.TweakListOptionsFunc
}

// New returns a new Interface.
func New(f internalinterfaces.SharedInformerFactory, namespace string, tweakListOptions internalinterfaces.TweakListOptionsFunc) Interface {
	return &version{factory: f, namespace: namespace, tweakListOptions: tweakListOptions}
}

// Pgclusters returns a PgclusterInformer.
func (v *version) Pgclusters() PgclusterInformer {
	return &pgclusterInformer{factory: v.factory, namespace: v.namespace, tweakListOptions: v.tweakListOptions}
}

// Pgpolicies returns a PgpolicyInformer.
func (v *version) Pgpolicies() PgpolicyInformer {
	return &pgpolicyInformer{factory: v.factory, namespace: v.namespace, tweakListOptions: v.tweakListOptions}
}

// Pgreplicas returns a PgreplicaInformer.
func (v *version) Pgreplicas() PgreplicaInformer {
	return &pgreplicaInformer{factory: v.factory, namespace: v.namespace, tweakListOptions: v.tweakListOptions}
}

// Pgtasks returns a PgtaskInformer.
func (v *version) Pgtasks() PgtaskInformer {
	return &pgtaskInformer{factory: v.factory, namespace: v.namespace, tweakListOptions: v.tweakListOptions}
}