/*
Copyright 2019 Google LLC

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

package fake

import (
	"context"

	fake "github.com/GoogleCloudPlatform/cloud-run-events/pkg/client/injection/informers/events/factory/fake"
	pubsubsource "github.com/GoogleCloudPlatform/cloud-run-events/pkg/client/injection/informers/events/v1alpha1/pubsubsource"
	controller "github.com/knative/pkg/controller"
	injection "github.com/knative/pkg/injection"
)

var Get = pubsubsource.Get

func init() {
	injection.Fake.RegisterInformer(withInformer)
}

func withInformer(ctx context.Context) (context.Context, controller.Informer) {
	f := fake.Get(ctx)
	inf := f.Events().V1alpha1().PubSubSources()
	return context.WithValue(ctx, pubsubsource.Key{}, inf), inf.Informer()
}
