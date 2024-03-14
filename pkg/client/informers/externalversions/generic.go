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

package externalversions

import (
	"fmt"

	schema "k8s.io/apimachinery/pkg/runtime/schema"
	cache "k8s.io/client-go/tools/cache"
	v1 "sigs.k8s.io/gateway-api/apis/v1"
	v1alpha2 "sigs.k8s.io/gateway-api/apis/v1alpha2"
	v1beta1 "sigs.k8s.io/gateway-api/apis/v1beta1"
)

// GenericInformer is type of SharedIndexInformer which will locate and delegate to other
// sharedInformers based on type
type GenericInformer interface {
	Informer() cache.SharedIndexInformer
	Lister() cache.GenericLister
}

type genericInformer struct {
	informer cache.SharedIndexInformer
	resource schema.GroupResource
}

// Informer returns the SharedIndexInformer.
func (f *genericInformer) Informer() cache.SharedIndexInformer {
	return f.informer
}

// Lister returns the GenericLister.
func (f *genericInformer) Lister() cache.GenericLister {
	return cache.NewGenericLister(f.Informer().GetIndexer(), f.resource)
}

// ForResource gives generic access to a shared informer of the matching type
// TODO extend this to unknown resources with a client pool
func (f *sharedInformerFactory) ForResource(resource schema.GroupVersionResource) (GenericInformer, error) {
	switch resource {
	// Group=gateway.networking.k8s.io, Version=v1
	case v1.SchemeGroupVersion.WithResource("gateways"):
		return &genericInformer{resource: resource.GroupResource(), informer: f.Gateway().V1().Gateways().Informer()}, nil
	case v1.SchemeGroupVersion.WithResource("gatewayclasses"):
		return &genericInformer{resource: resource.GroupResource(), informer: f.Gateway().V1().GatewayClasses().Informer()}, nil
	case v1.SchemeGroupVersion.WithResource("httproutes"):
		return &genericInformer{resource: resource.GroupResource(), informer: f.Gateway().V1().HTTPRoutes().Informer()}, nil

		// Group=gateway.networking.k8s.io, Version=v1alpha2
	case v1alpha2.SchemeGroupVersion.WithResource("backendtlspolicies"):
		return &genericInformer{resource: resource.GroupResource(), informer: f.Gateway().V1alpha2().BackendTLSPolicies().Informer()}, nil
	case v1alpha2.SchemeGroupVersion.WithResource("egressroutes"):
		return &genericInformer{resource: resource.GroupResource(), informer: f.Gateway().V1alpha2().EgressRoutes().Informer()}, nil
	case v1alpha2.SchemeGroupVersion.WithResource("grpcroutes"):
		return &genericInformer{resource: resource.GroupResource(), informer: f.Gateway().V1alpha2().GRPCRoutes().Informer()}, nil
	case v1alpha2.SchemeGroupVersion.WithResource("gateways"):
		return &genericInformer{resource: resource.GroupResource(), informer: f.Gateway().V1alpha2().Gateways().Informer()}, nil
	case v1alpha2.SchemeGroupVersion.WithResource("gatewayclasses"):
		return &genericInformer{resource: resource.GroupResource(), informer: f.Gateway().V1alpha2().GatewayClasses().Informer()}, nil
	case v1alpha2.SchemeGroupVersion.WithResource("httproutes"):
		return &genericInformer{resource: resource.GroupResource(), informer: f.Gateway().V1alpha2().HTTPRoutes().Informer()}, nil
	case v1alpha2.SchemeGroupVersion.WithResource("referencegrants"):
		return &genericInformer{resource: resource.GroupResource(), informer: f.Gateway().V1alpha2().ReferenceGrants().Informer()}, nil
	case v1alpha2.SchemeGroupVersion.WithResource("tcproutes"):
		return &genericInformer{resource: resource.GroupResource(), informer: f.Gateway().V1alpha2().TCPRoutes().Informer()}, nil
	case v1alpha2.SchemeGroupVersion.WithResource("tlsroutes"):
		return &genericInformer{resource: resource.GroupResource(), informer: f.Gateway().V1alpha2().TLSRoutes().Informer()}, nil
	case v1alpha2.SchemeGroupVersion.WithResource("udproutes"):
		return &genericInformer{resource: resource.GroupResource(), informer: f.Gateway().V1alpha2().UDPRoutes().Informer()}, nil

		// Group=gateway.networking.k8s.io, Version=v1beta1
	case v1beta1.SchemeGroupVersion.WithResource("gateways"):
		return &genericInformer{resource: resource.GroupResource(), informer: f.Gateway().V1beta1().Gateways().Informer()}, nil
	case v1beta1.SchemeGroupVersion.WithResource("gatewayclasses"):
		return &genericInformer{resource: resource.GroupResource(), informer: f.Gateway().V1beta1().GatewayClasses().Informer()}, nil
	case v1beta1.SchemeGroupVersion.WithResource("httproutes"):
		return &genericInformer{resource: resource.GroupResource(), informer: f.Gateway().V1beta1().HTTPRoutes().Informer()}, nil
	case v1beta1.SchemeGroupVersion.WithResource("referencegrants"):
		return &genericInformer{resource: resource.GroupResource(), informer: f.Gateway().V1beta1().ReferenceGrants().Informer()}, nil

	}

	return nil, fmt.Errorf("no informer found for %v", resource)
}
