/*
Copyright 2017 The Kubernetes Authors.

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

// This file was automatically generated by informer-gen

package internalversion

import (
	v1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	runtime "k8s.io/apimachinery/pkg/runtime"
	watch "k8s.io/apimachinery/pkg/watch"
	cache "k8s.io/client-go/tools/cache"
	extensions "k8s.io/kubernetes/pkg/apis/extensions"
	internalclientset "k8s.io/kubernetes/pkg/client/clientset_generated/internalclientset"
	internalinterfaces "k8s.io/kubernetes/pkg/client/informers/informers_generated/internalversion/internalinterfaces"
	internalversion "k8s.io/kubernetes/pkg/client/listers/extensions/internalversion"
	time "time"
)

// DeploymentInformer provides access to a shared informer and lister for
// Deployments.
type DeploymentInformer interface {
	Informer() cache.SharedIndexInformer
	Lister() internalversion.DeploymentLister
}

type deploymentInformer struct {
	factory internalinterfaces.SharedInformerFactory
}

// NewDeploymentInformer constructs a new informer for Deployment type.
// Always prefer using an informer factory to get a shared informer instead of getting an independent
// one. This reduces memory footprint and number of connections to the server.
func NewDeploymentInformer(client internalclientset.Interface, namespace string, resyncPeriod time.Duration, indexers cache.Indexers) cache.SharedIndexInformer {
	return cache.NewSharedIndexInformer(
		&cache.ListWatch{
			ListFunc: func(options v1.ListOptions) (runtime.Object, error) {
				return client.Extensions().Deployments(namespace).List(options)
			},
			WatchFunc: func(options v1.ListOptions) (watch.Interface, error) {
				return client.Extensions().Deployments(namespace).Watch(options)
			},
		},
		&extensions.Deployment{},
		resyncPeriod,
		indexers,
	)
}

func defaultDeploymentInformer(client internalclientset.Interface, resyncPeriod time.Duration) cache.SharedIndexInformer {
	return NewDeploymentInformer(client, v1.NamespaceAll, resyncPeriod, cache.Indexers{cache.NamespaceIndex: cache.MetaNamespaceIndexFunc})
}

func (f *deploymentInformer) Informer() cache.SharedIndexInformer {
	return f.factory.InformerFor(&extensions.Deployment{}, defaultDeploymentInformer)
}

func (f *deploymentInformer) Lister() internalversion.DeploymentLister {
	return internalversion.NewDeploymentLister(f.Informer().GetIndexer())
}
