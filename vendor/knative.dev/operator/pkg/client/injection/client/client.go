/*
Copyright 2022 The Knative Authors

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

// Code generated by injection-gen. DO NOT EDIT.

package client

import (
	context "context"
	json "encoding/json"
	errors "errors"
	fmt "fmt"

	v1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	unstructured "k8s.io/apimachinery/pkg/apis/meta/v1/unstructured"
	runtime "k8s.io/apimachinery/pkg/runtime"
	schema "k8s.io/apimachinery/pkg/runtime/schema"
	types "k8s.io/apimachinery/pkg/types"
	watch "k8s.io/apimachinery/pkg/watch"
	discovery "k8s.io/client-go/discovery"
	dynamic "k8s.io/client-go/dynamic"
	rest "k8s.io/client-go/rest"
	v1alpha1 "knative.dev/operator/pkg/apis/operator/v1alpha1"
	v1beta1 "knative.dev/operator/pkg/apis/operator/v1beta1"
	versioned "knative.dev/operator/pkg/client/clientset/versioned"
	typedoperatorv1alpha1 "knative.dev/operator/pkg/client/clientset/versioned/typed/operator/v1alpha1"
	typedoperatorv1beta1 "knative.dev/operator/pkg/client/clientset/versioned/typed/operator/v1beta1"
	injection "knative.dev/pkg/injection"
	dynamicclient "knative.dev/pkg/injection/clients/dynamicclient"
	logging "knative.dev/pkg/logging"
)

func init() {
	injection.Default.RegisterClient(withClientFromConfig)
	injection.Default.RegisterClientFetcher(func(ctx context.Context) interface{} {
		return Get(ctx)
	})
	injection.Dynamic.RegisterDynamicClient(withClientFromDynamic)
}

// Key is used as the key for associating information with a context.Context.
type Key struct{}

func withClientFromConfig(ctx context.Context, cfg *rest.Config) context.Context {
	return context.WithValue(ctx, Key{}, versioned.NewForConfigOrDie(cfg))
}

func withClientFromDynamic(ctx context.Context) context.Context {
	return context.WithValue(ctx, Key{}, &wrapClient{dyn: dynamicclient.Get(ctx)})
}

// Get extracts the versioned.Interface client from the context.
func Get(ctx context.Context) versioned.Interface {
	untyped := ctx.Value(Key{})
	if untyped == nil {
		if injection.GetConfig(ctx) == nil {
			logging.FromContext(ctx).Panic(
				"Unable to fetch knative.dev/operator/pkg/client/clientset/versioned.Interface from context. This context is not the application context (which is typically given to constructors via sharedmain).")
		} else {
			logging.FromContext(ctx).Panic(
				"Unable to fetch knative.dev/operator/pkg/client/clientset/versioned.Interface from context.")
		}
	}
	return untyped.(versioned.Interface)
}

type wrapClient struct {
	dyn dynamic.Interface
}

var _ versioned.Interface = (*wrapClient)(nil)

func (w *wrapClient) Discovery() discovery.DiscoveryInterface {
	panic("Discovery called on dynamic client!")
}

func convert(from interface{}, to runtime.Object) error {
	bs, err := json.Marshal(from)
	if err != nil {
		return fmt.Errorf("Marshal() = %w", err)
	}
	if err := json.Unmarshal(bs, to); err != nil {
		return fmt.Errorf("Unmarshal() = %w", err)
	}
	return nil
}

// OperatorV1alpha1 retrieves the OperatorV1alpha1Client
func (w *wrapClient) OperatorV1alpha1() typedoperatorv1alpha1.OperatorV1alpha1Interface {
	return &wrapOperatorV1alpha1{
		dyn: w.dyn,
	}
}

type wrapOperatorV1alpha1 struct {
	dyn dynamic.Interface
}

func (w *wrapOperatorV1alpha1) RESTClient() rest.Interface {
	panic("RESTClient called on dynamic client!")
}

func (w *wrapOperatorV1alpha1) KnativeEventings(namespace string) typedoperatorv1alpha1.KnativeEventingInterface {
	return &wrapOperatorV1alpha1KnativeEventingImpl{
		dyn: w.dyn.Resource(schema.GroupVersionResource{
			Group:    "operator.knative.dev",
			Version:  "v1alpha1",
			Resource: "knativeeventings",
		}),

		namespace: namespace,
	}
}

type wrapOperatorV1alpha1KnativeEventingImpl struct {
	dyn dynamic.NamespaceableResourceInterface

	namespace string
}

var _ typedoperatorv1alpha1.KnativeEventingInterface = (*wrapOperatorV1alpha1KnativeEventingImpl)(nil)

func (w *wrapOperatorV1alpha1KnativeEventingImpl) Create(ctx context.Context, in *v1alpha1.KnativeEventing, opts v1.CreateOptions) (*v1alpha1.KnativeEventing, error) {
	in.SetGroupVersionKind(schema.GroupVersionKind{
		Group:   "operator.knative.dev",
		Version: "v1alpha1",
		Kind:    "KnativeEventing",
	})
	uo := &unstructured.Unstructured{}
	if err := convert(in, uo); err != nil {
		return nil, err
	}
	uo, err := w.dyn.Namespace(w.namespace).Create(ctx, uo, opts)
	if err != nil {
		return nil, err
	}
	out := &v1alpha1.KnativeEventing{}
	if err := convert(uo, out); err != nil {
		return nil, err
	}
	return out, nil
}

func (w *wrapOperatorV1alpha1KnativeEventingImpl) Delete(ctx context.Context, name string, opts v1.DeleteOptions) error {
	return w.dyn.Namespace(w.namespace).Delete(ctx, name, opts)
}

func (w *wrapOperatorV1alpha1KnativeEventingImpl) DeleteCollection(ctx context.Context, opts v1.DeleteOptions, listOpts v1.ListOptions) error {
	return w.dyn.Namespace(w.namespace).DeleteCollection(ctx, opts, listOpts)
}

func (w *wrapOperatorV1alpha1KnativeEventingImpl) Get(ctx context.Context, name string, opts v1.GetOptions) (*v1alpha1.KnativeEventing, error) {
	uo, err := w.dyn.Namespace(w.namespace).Get(ctx, name, opts)
	if err != nil {
		return nil, err
	}
	out := &v1alpha1.KnativeEventing{}
	if err := convert(uo, out); err != nil {
		return nil, err
	}
	return out, nil
}

func (w *wrapOperatorV1alpha1KnativeEventingImpl) List(ctx context.Context, opts v1.ListOptions) (*v1alpha1.KnativeEventingList, error) {
	uo, err := w.dyn.Namespace(w.namespace).List(ctx, opts)
	if err != nil {
		return nil, err
	}
	out := &v1alpha1.KnativeEventingList{}
	if err := convert(uo, out); err != nil {
		return nil, err
	}
	return out, nil
}

func (w *wrapOperatorV1alpha1KnativeEventingImpl) Patch(ctx context.Context, name string, pt types.PatchType, data []byte, opts v1.PatchOptions, subresources ...string) (result *v1alpha1.KnativeEventing, err error) {
	uo, err := w.dyn.Namespace(w.namespace).Patch(ctx, name, pt, data, opts)
	if err != nil {
		return nil, err
	}
	out := &v1alpha1.KnativeEventing{}
	if err := convert(uo, out); err != nil {
		return nil, err
	}
	return out, nil
}

func (w *wrapOperatorV1alpha1KnativeEventingImpl) Update(ctx context.Context, in *v1alpha1.KnativeEventing, opts v1.UpdateOptions) (*v1alpha1.KnativeEventing, error) {
	in.SetGroupVersionKind(schema.GroupVersionKind{
		Group:   "operator.knative.dev",
		Version: "v1alpha1",
		Kind:    "KnativeEventing",
	})
	uo := &unstructured.Unstructured{}
	if err := convert(in, uo); err != nil {
		return nil, err
	}
	uo, err := w.dyn.Namespace(w.namespace).Update(ctx, uo, opts)
	if err != nil {
		return nil, err
	}
	out := &v1alpha1.KnativeEventing{}
	if err := convert(uo, out); err != nil {
		return nil, err
	}
	return out, nil
}

func (w *wrapOperatorV1alpha1KnativeEventingImpl) UpdateStatus(ctx context.Context, in *v1alpha1.KnativeEventing, opts v1.UpdateOptions) (*v1alpha1.KnativeEventing, error) {
	in.SetGroupVersionKind(schema.GroupVersionKind{
		Group:   "operator.knative.dev",
		Version: "v1alpha1",
		Kind:    "KnativeEventing",
	})
	uo := &unstructured.Unstructured{}
	if err := convert(in, uo); err != nil {
		return nil, err
	}
	uo, err := w.dyn.Namespace(w.namespace).UpdateStatus(ctx, uo, opts)
	if err != nil {
		return nil, err
	}
	out := &v1alpha1.KnativeEventing{}
	if err := convert(uo, out); err != nil {
		return nil, err
	}
	return out, nil
}

func (w *wrapOperatorV1alpha1KnativeEventingImpl) Watch(ctx context.Context, opts v1.ListOptions) (watch.Interface, error) {
	return nil, errors.New("NYI: Watch")
}

func (w *wrapOperatorV1alpha1) KnativeServings(namespace string) typedoperatorv1alpha1.KnativeServingInterface {
	return &wrapOperatorV1alpha1KnativeServingImpl{
		dyn: w.dyn.Resource(schema.GroupVersionResource{
			Group:    "operator.knative.dev",
			Version:  "v1alpha1",
			Resource: "knativeservings",
		}),

		namespace: namespace,
	}
}

type wrapOperatorV1alpha1KnativeServingImpl struct {
	dyn dynamic.NamespaceableResourceInterface

	namespace string
}

var _ typedoperatorv1alpha1.KnativeServingInterface = (*wrapOperatorV1alpha1KnativeServingImpl)(nil)

func (w *wrapOperatorV1alpha1KnativeServingImpl) Create(ctx context.Context, in *v1alpha1.KnativeServing, opts v1.CreateOptions) (*v1alpha1.KnativeServing, error) {
	in.SetGroupVersionKind(schema.GroupVersionKind{
		Group:   "operator.knative.dev",
		Version: "v1alpha1",
		Kind:    "KnativeServing",
	})
	uo := &unstructured.Unstructured{}
	if err := convert(in, uo); err != nil {
		return nil, err
	}
	uo, err := w.dyn.Namespace(w.namespace).Create(ctx, uo, opts)
	if err != nil {
		return nil, err
	}
	out := &v1alpha1.KnativeServing{}
	if err := convert(uo, out); err != nil {
		return nil, err
	}
	return out, nil
}

func (w *wrapOperatorV1alpha1KnativeServingImpl) Delete(ctx context.Context, name string, opts v1.DeleteOptions) error {
	return w.dyn.Namespace(w.namespace).Delete(ctx, name, opts)
}

func (w *wrapOperatorV1alpha1KnativeServingImpl) DeleteCollection(ctx context.Context, opts v1.DeleteOptions, listOpts v1.ListOptions) error {
	return w.dyn.Namespace(w.namespace).DeleteCollection(ctx, opts, listOpts)
}

func (w *wrapOperatorV1alpha1KnativeServingImpl) Get(ctx context.Context, name string, opts v1.GetOptions) (*v1alpha1.KnativeServing, error) {
	uo, err := w.dyn.Namespace(w.namespace).Get(ctx, name, opts)
	if err != nil {
		return nil, err
	}
	out := &v1alpha1.KnativeServing{}
	if err := convert(uo, out); err != nil {
		return nil, err
	}
	return out, nil
}

func (w *wrapOperatorV1alpha1KnativeServingImpl) List(ctx context.Context, opts v1.ListOptions) (*v1alpha1.KnativeServingList, error) {
	uo, err := w.dyn.Namespace(w.namespace).List(ctx, opts)
	if err != nil {
		return nil, err
	}
	out := &v1alpha1.KnativeServingList{}
	if err := convert(uo, out); err != nil {
		return nil, err
	}
	return out, nil
}

func (w *wrapOperatorV1alpha1KnativeServingImpl) Patch(ctx context.Context, name string, pt types.PatchType, data []byte, opts v1.PatchOptions, subresources ...string) (result *v1alpha1.KnativeServing, err error) {
	uo, err := w.dyn.Namespace(w.namespace).Patch(ctx, name, pt, data, opts)
	if err != nil {
		return nil, err
	}
	out := &v1alpha1.KnativeServing{}
	if err := convert(uo, out); err != nil {
		return nil, err
	}
	return out, nil
}

func (w *wrapOperatorV1alpha1KnativeServingImpl) Update(ctx context.Context, in *v1alpha1.KnativeServing, opts v1.UpdateOptions) (*v1alpha1.KnativeServing, error) {
	in.SetGroupVersionKind(schema.GroupVersionKind{
		Group:   "operator.knative.dev",
		Version: "v1alpha1",
		Kind:    "KnativeServing",
	})
	uo := &unstructured.Unstructured{}
	if err := convert(in, uo); err != nil {
		return nil, err
	}
	uo, err := w.dyn.Namespace(w.namespace).Update(ctx, uo, opts)
	if err != nil {
		return nil, err
	}
	out := &v1alpha1.KnativeServing{}
	if err := convert(uo, out); err != nil {
		return nil, err
	}
	return out, nil
}

func (w *wrapOperatorV1alpha1KnativeServingImpl) UpdateStatus(ctx context.Context, in *v1alpha1.KnativeServing, opts v1.UpdateOptions) (*v1alpha1.KnativeServing, error) {
	in.SetGroupVersionKind(schema.GroupVersionKind{
		Group:   "operator.knative.dev",
		Version: "v1alpha1",
		Kind:    "KnativeServing",
	})
	uo := &unstructured.Unstructured{}
	if err := convert(in, uo); err != nil {
		return nil, err
	}
	uo, err := w.dyn.Namespace(w.namespace).UpdateStatus(ctx, uo, opts)
	if err != nil {
		return nil, err
	}
	out := &v1alpha1.KnativeServing{}
	if err := convert(uo, out); err != nil {
		return nil, err
	}
	return out, nil
}

func (w *wrapOperatorV1alpha1KnativeServingImpl) Watch(ctx context.Context, opts v1.ListOptions) (watch.Interface, error) {
	return nil, errors.New("NYI: Watch")
}

// OperatorV1beta1 retrieves the OperatorV1beta1Client
func (w *wrapClient) OperatorV1beta1() typedoperatorv1beta1.OperatorV1beta1Interface {
	return &wrapOperatorV1beta1{
		dyn: w.dyn,
	}
}

type wrapOperatorV1beta1 struct {
	dyn dynamic.Interface
}

func (w *wrapOperatorV1beta1) RESTClient() rest.Interface {
	panic("RESTClient called on dynamic client!")
}

func (w *wrapOperatorV1beta1) KnativeEventings(namespace string) typedoperatorv1beta1.KnativeEventingInterface {
	return &wrapOperatorV1beta1KnativeEventingImpl{
		dyn: w.dyn.Resource(schema.GroupVersionResource{
			Group:    "operator.knative.dev",
			Version:  "v1beta1",
			Resource: "knativeeventings",
		}),

		namespace: namespace,
	}
}

type wrapOperatorV1beta1KnativeEventingImpl struct {
	dyn dynamic.NamespaceableResourceInterface

	namespace string
}

var _ typedoperatorv1beta1.KnativeEventingInterface = (*wrapOperatorV1beta1KnativeEventingImpl)(nil)

func (w *wrapOperatorV1beta1KnativeEventingImpl) Create(ctx context.Context, in *v1beta1.KnativeEventing, opts v1.CreateOptions) (*v1beta1.KnativeEventing, error) {
	in.SetGroupVersionKind(schema.GroupVersionKind{
		Group:   "operator.knative.dev",
		Version: "v1beta1",
		Kind:    "KnativeEventing",
	})
	uo := &unstructured.Unstructured{}
	if err := convert(in, uo); err != nil {
		return nil, err
	}
	uo, err := w.dyn.Namespace(w.namespace).Create(ctx, uo, opts)
	if err != nil {
		return nil, err
	}
	out := &v1beta1.KnativeEventing{}
	if err := convert(uo, out); err != nil {
		return nil, err
	}
	return out, nil
}

func (w *wrapOperatorV1beta1KnativeEventingImpl) Delete(ctx context.Context, name string, opts v1.DeleteOptions) error {
	return w.dyn.Namespace(w.namespace).Delete(ctx, name, opts)
}

func (w *wrapOperatorV1beta1KnativeEventingImpl) DeleteCollection(ctx context.Context, opts v1.DeleteOptions, listOpts v1.ListOptions) error {
	return w.dyn.Namespace(w.namespace).DeleteCollection(ctx, opts, listOpts)
}

func (w *wrapOperatorV1beta1KnativeEventingImpl) Get(ctx context.Context, name string, opts v1.GetOptions) (*v1beta1.KnativeEventing, error) {
	uo, err := w.dyn.Namespace(w.namespace).Get(ctx, name, opts)
	if err != nil {
		return nil, err
	}
	out := &v1beta1.KnativeEventing{}
	if err := convert(uo, out); err != nil {
		return nil, err
	}
	return out, nil
}

func (w *wrapOperatorV1beta1KnativeEventingImpl) List(ctx context.Context, opts v1.ListOptions) (*v1beta1.KnativeEventingList, error) {
	uo, err := w.dyn.Namespace(w.namespace).List(ctx, opts)
	if err != nil {
		return nil, err
	}
	out := &v1beta1.KnativeEventingList{}
	if err := convert(uo, out); err != nil {
		return nil, err
	}
	return out, nil
}

func (w *wrapOperatorV1beta1KnativeEventingImpl) Patch(ctx context.Context, name string, pt types.PatchType, data []byte, opts v1.PatchOptions, subresources ...string) (result *v1beta1.KnativeEventing, err error) {
	uo, err := w.dyn.Namespace(w.namespace).Patch(ctx, name, pt, data, opts)
	if err != nil {
		return nil, err
	}
	out := &v1beta1.KnativeEventing{}
	if err := convert(uo, out); err != nil {
		return nil, err
	}
	return out, nil
}

func (w *wrapOperatorV1beta1KnativeEventingImpl) Update(ctx context.Context, in *v1beta1.KnativeEventing, opts v1.UpdateOptions) (*v1beta1.KnativeEventing, error) {
	in.SetGroupVersionKind(schema.GroupVersionKind{
		Group:   "operator.knative.dev",
		Version: "v1beta1",
		Kind:    "KnativeEventing",
	})
	uo := &unstructured.Unstructured{}
	if err := convert(in, uo); err != nil {
		return nil, err
	}
	uo, err := w.dyn.Namespace(w.namespace).Update(ctx, uo, opts)
	if err != nil {
		return nil, err
	}
	out := &v1beta1.KnativeEventing{}
	if err := convert(uo, out); err != nil {
		return nil, err
	}
	return out, nil
}

func (w *wrapOperatorV1beta1KnativeEventingImpl) UpdateStatus(ctx context.Context, in *v1beta1.KnativeEventing, opts v1.UpdateOptions) (*v1beta1.KnativeEventing, error) {
	in.SetGroupVersionKind(schema.GroupVersionKind{
		Group:   "operator.knative.dev",
		Version: "v1beta1",
		Kind:    "KnativeEventing",
	})
	uo := &unstructured.Unstructured{}
	if err := convert(in, uo); err != nil {
		return nil, err
	}
	uo, err := w.dyn.Namespace(w.namespace).UpdateStatus(ctx, uo, opts)
	if err != nil {
		return nil, err
	}
	out := &v1beta1.KnativeEventing{}
	if err := convert(uo, out); err != nil {
		return nil, err
	}
	return out, nil
}

func (w *wrapOperatorV1beta1KnativeEventingImpl) Watch(ctx context.Context, opts v1.ListOptions) (watch.Interface, error) {
	return nil, errors.New("NYI: Watch")
}

func (w *wrapOperatorV1beta1) KnativeServings(namespace string) typedoperatorv1beta1.KnativeServingInterface {
	return &wrapOperatorV1beta1KnativeServingImpl{
		dyn: w.dyn.Resource(schema.GroupVersionResource{
			Group:    "operator.knative.dev",
			Version:  "v1beta1",
			Resource: "knativeservings",
		}),

		namespace: namespace,
	}
}

type wrapOperatorV1beta1KnativeServingImpl struct {
	dyn dynamic.NamespaceableResourceInterface

	namespace string
}

var _ typedoperatorv1beta1.KnativeServingInterface = (*wrapOperatorV1beta1KnativeServingImpl)(nil)

func (w *wrapOperatorV1beta1KnativeServingImpl) Create(ctx context.Context, in *v1beta1.KnativeServing, opts v1.CreateOptions) (*v1beta1.KnativeServing, error) {
	in.SetGroupVersionKind(schema.GroupVersionKind{
		Group:   "operator.knative.dev",
		Version: "v1beta1",
		Kind:    "KnativeServing",
	})
	uo := &unstructured.Unstructured{}
	if err := convert(in, uo); err != nil {
		return nil, err
	}
	uo, err := w.dyn.Namespace(w.namespace).Create(ctx, uo, opts)
	if err != nil {
		return nil, err
	}
	out := &v1beta1.KnativeServing{}
	if err := convert(uo, out); err != nil {
		return nil, err
	}
	return out, nil
}

func (w *wrapOperatorV1beta1KnativeServingImpl) Delete(ctx context.Context, name string, opts v1.DeleteOptions) error {
	return w.dyn.Namespace(w.namespace).Delete(ctx, name, opts)
}

func (w *wrapOperatorV1beta1KnativeServingImpl) DeleteCollection(ctx context.Context, opts v1.DeleteOptions, listOpts v1.ListOptions) error {
	return w.dyn.Namespace(w.namespace).DeleteCollection(ctx, opts, listOpts)
}

func (w *wrapOperatorV1beta1KnativeServingImpl) Get(ctx context.Context, name string, opts v1.GetOptions) (*v1beta1.KnativeServing, error) {
	uo, err := w.dyn.Namespace(w.namespace).Get(ctx, name, opts)
	if err != nil {
		return nil, err
	}
	out := &v1beta1.KnativeServing{}
	if err := convert(uo, out); err != nil {
		return nil, err
	}
	return out, nil
}

func (w *wrapOperatorV1beta1KnativeServingImpl) List(ctx context.Context, opts v1.ListOptions) (*v1beta1.KnativeServingList, error) {
	uo, err := w.dyn.Namespace(w.namespace).List(ctx, opts)
	if err != nil {
		return nil, err
	}
	out := &v1beta1.KnativeServingList{}
	if err := convert(uo, out); err != nil {
		return nil, err
	}
	return out, nil
}

func (w *wrapOperatorV1beta1KnativeServingImpl) Patch(ctx context.Context, name string, pt types.PatchType, data []byte, opts v1.PatchOptions, subresources ...string) (result *v1beta1.KnativeServing, err error) {
	uo, err := w.dyn.Namespace(w.namespace).Patch(ctx, name, pt, data, opts)
	if err != nil {
		return nil, err
	}
	out := &v1beta1.KnativeServing{}
	if err := convert(uo, out); err != nil {
		return nil, err
	}
	return out, nil
}

func (w *wrapOperatorV1beta1KnativeServingImpl) Update(ctx context.Context, in *v1beta1.KnativeServing, opts v1.UpdateOptions) (*v1beta1.KnativeServing, error) {
	in.SetGroupVersionKind(schema.GroupVersionKind{
		Group:   "operator.knative.dev",
		Version: "v1beta1",
		Kind:    "KnativeServing",
	})
	uo := &unstructured.Unstructured{}
	if err := convert(in, uo); err != nil {
		return nil, err
	}
	uo, err := w.dyn.Namespace(w.namespace).Update(ctx, uo, opts)
	if err != nil {
		return nil, err
	}
	out := &v1beta1.KnativeServing{}
	if err := convert(uo, out); err != nil {
		return nil, err
	}
	return out, nil
}

func (w *wrapOperatorV1beta1KnativeServingImpl) UpdateStatus(ctx context.Context, in *v1beta1.KnativeServing, opts v1.UpdateOptions) (*v1beta1.KnativeServing, error) {
	in.SetGroupVersionKind(schema.GroupVersionKind{
		Group:   "operator.knative.dev",
		Version: "v1beta1",
		Kind:    "KnativeServing",
	})
	uo := &unstructured.Unstructured{}
	if err := convert(in, uo); err != nil {
		return nil, err
	}
	uo, err := w.dyn.Namespace(w.namespace).UpdateStatus(ctx, uo, opts)
	if err != nil {
		return nil, err
	}
	out := &v1beta1.KnativeServing{}
	if err := convert(uo, out); err != nil {
		return nil, err
	}
	return out, nil
}

func (w *wrapOperatorV1beta1KnativeServingImpl) Watch(ctx context.Context, opts v1.ListOptions) (watch.Interface, error) {
	return nil, errors.New("NYI: Watch")
}
