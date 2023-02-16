// Copyright (c) 2022 Alibaba Group Holding Ltd.
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

// Code generated by informer-gen. DO NOT EDIT.

package v1alpha1

import (
	"context"
	time "time"

	extensionsv1alpha1 "github.com/alibaba/higress/client/pkg/apis/extensions/v1alpha1"
	versioned "github.com/alibaba/higress/client/pkg/clientset/versioned"
	internalinterfaces "github.com/alibaba/higress/client/pkg/informers/externalversions/internalinterfaces"
	v1alpha1 "github.com/alibaba/higress/client/pkg/listers/extensions/v1alpha1"
	v1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	runtime "k8s.io/apimachinery/pkg/runtime"
	watch "k8s.io/apimachinery/pkg/watch"
	cache "k8s.io/client-go/tools/cache"
)

// WasmPluginInformer provides access to a shared informer and lister for
// WasmPlugins.
type WasmPluginInformer interface {
	Informer() cache.SharedIndexInformer
	Lister() v1alpha1.WasmPluginLister
}

type wasmPluginInformer struct {
	factory          internalinterfaces.SharedInformerFactory
	tweakListOptions internalinterfaces.TweakListOptionsFunc
	namespace        string
}

// NewWasmPluginInformer constructs a new informer for WasmPlugin type.
// Always prefer using an informer factory to get a shared informer instead of getting an independent
// one. This reduces memory footprint and number of connections to the server.
func NewWasmPluginInformer(client versioned.Interface, namespace string, resyncPeriod time.Duration, indexers cache.Indexers) cache.SharedIndexInformer {
	return NewFilteredWasmPluginInformer(client, namespace, resyncPeriod, indexers, nil)
}

// NewFilteredWasmPluginInformer constructs a new informer for WasmPlugin type.
// Always prefer using an informer factory to get a shared informer instead of getting an independent
// one. This reduces memory footprint and number of connections to the server.
func NewFilteredWasmPluginInformer(client versioned.Interface, namespace string, resyncPeriod time.Duration, indexers cache.Indexers, tweakListOptions internalinterfaces.TweakListOptionsFunc) cache.SharedIndexInformer {
	return cache.NewSharedIndexInformer(
		&cache.ListWatch{
			ListFunc: func(options v1.ListOptions) (runtime.Object, error) {
				if tweakListOptions != nil {
					tweakListOptions(&options)
				}
				return client.ExtensionsV1alpha1().WasmPlugins(namespace).List(context.TODO(), options)
			},
			WatchFunc: func(options v1.ListOptions) (watch.Interface, error) {
				if tweakListOptions != nil {
					tweakListOptions(&options)
				}
				return client.ExtensionsV1alpha1().WasmPlugins(namespace).Watch(context.TODO(), options)
			},
		},
		&extensionsv1alpha1.WasmPlugin{},
		resyncPeriod,
		indexers,
	)
}

func (f *wasmPluginInformer) defaultInformer(client versioned.Interface, resyncPeriod time.Duration) cache.SharedIndexInformer {
	return NewFilteredWasmPluginInformer(client, f.namespace, resyncPeriod, cache.Indexers{cache.NamespaceIndex: cache.MetaNamespaceIndexFunc}, f.tweakListOptions)
}

func (f *wasmPluginInformer) Informer() cache.SharedIndexInformer {
	return f.factory.InformerFor(&extensionsv1alpha1.WasmPlugin{}, f.defaultInformer)
}

func (f *wasmPluginInformer) Lister() v1alpha1.WasmPluginLister {
	return v1alpha1.NewWasmPluginLister(f.Informer().GetIndexer())
}
