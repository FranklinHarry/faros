/*
Copyright 2018 Pusher Ltd.

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

package gittrackobject

import (
	"fmt"
	"log"

	"k8s.io/apimachinery/pkg/apis/meta/v1/unstructured"
)

// watch sets up an informer for the object kind and sends events to the
// ReconcileGitTrackObject's eventStream.
func (r *ReconcileGitTrackObject) watch(obj unstructured.Unstructured) error {
	if _, ok := r.informers[informerKey(obj)]; ok {
		// Informer already set up
		return nil
	}

	// Create new informer
	log.Printf("Creating new informer for kind %s", obj.GetObjectKind().GroupVersionKind().Kind)
	informer, err := r.cache.GetInformer(&obj)
	if err != nil {
		msg := fmt.Sprintf("error creating informer: %v", err)
		log.Printf(msg)
		return fmt.Errorf(msg)
	}

	// Add event handlers
	informer.AddEventHandler(&eventToChannelHandler{
		eventsChan: r.eventStream,
	})

	// Store and run informer
	r.informers[informerKey(obj)] = informer
	return nil
}

// informerKey creates a unique identifier containing the object's namespace,
// group, version and kind.
//
// This can be used to uniquely identify informers.
func informerKey(obj unstructured.Unstructured) string {
	return fmt.Sprintf("%s:%s", obj.GetNamespace(), obj.GroupVersionKind().String())
}
