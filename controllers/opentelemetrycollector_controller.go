/*
Copyright 2023.

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

package controllers

import (
	"context"

	"k8s.io/apimachinery/pkg/runtime"
	ctrl "sigs.k8s.io/controller-runtime"
	"sigs.k8s.io/controller-runtime/pkg/client"
	"sigs.k8s.io/controller-runtime/pkg/log"

	otelv1 "github.com/mizhexiaoxiao/otel-operator/api/v1"
	"github.com/mizhexiaoxiao/otel-operator/reconcile"
	apierrors "k8s.io/apimachinery/pkg/api/errors"
)

// OpenTelemetryCollectorReconciler reconciles a OpenTelemetryCollector object
type OpenTelemetryCollectorReconciler struct {
	client.Client
	Scheme *runtime.Scheme
}

//+kubebuilder:rbac:groups=otel.mzx.org,resources=opentelemetrycollectors,verbs=get;list;watch;create;update;patch;delete
//+kubebuilder:rbac:groups=otel.mzx.org,resources=opentelemetrycollectors/status,verbs=get;update;patch
//+kubebuilder:rbac:groups=otel.mzx.org,resources=opentelemetrycollectors/finalizers,verbs=update
//+kubebuilder:rbac:groups="",resources=configmaps,verbs=get;list;watch;create;update;patch;delete
//+kubebuilder:rbac:groups="",resources=services,verbs=get;list;watch;create;update;patch;delete
//+kubebuilder:rbac:groups="apps",resources=statefulsets,verbs=get;list;watch;create;update;patch;delete
//+kubebuilder:rbac:groups="apps",resources=deployments,verbs=get;list;watch;create;update;patch;delete

// Reconcile is part of the main kubernetes reconciliation loop which aims to
// move the current state of the cluster closer to the desired state.
// TODO(user): Modify the Reconcile function to compare the state specified by
// the OpenTelemetryCollector object against the actual cluster state, and then
// perform operations to make the cluster state reflect the state specified by
// the user.
//
// For more details, check Reconcile and its Result here:
// - https://pkg.go.dev/sigs.k8s.io/controller-runtime@v0.12.2/pkg/reconcile
func (r *OpenTelemetryCollectorReconciler) Reconcile(ctx context.Context, req ctrl.Request) (ctrl.Result, error) {
	log := log.FromContext(ctx)

	instance := otelv1.OpenTelemetryCollector{}
	if err := r.Client.Get(ctx, req.NamespacedName, &instance); err != nil {
		if !apierrors.IsNotFound(err) {
			log.Error(err, "unable to fetch OpenTelemetryCollector")
		}
		return ctrl.Result{}, client.IgnoreNotFound(err)
	}

	if instance.DeletionTimestamp != nil {
		return ctrl.Result{}, nil
	}

	params := reconcile.Params{
		Client:   r.Client,
		Scheme:   r.Scheme,
		Log:      log,
		Instance: instance,
	}

	//handle configmap
	if err := reconcile.CollectorConfigMap(ctx, params); err != nil {
		return ctrl.Result{}, err
	}

	//handle statefulset
	if err := reconcile.StatefulSet(ctx, params); err != nil {
		return ctrl.Result{}, err
	}

	//handle headless service
	if err := reconcile.Service(ctx, params); err != nil {
		return ctrl.Result{}, err
	}

	return ctrl.Result{}, nil
}

// SetupWithManager sets up the controller with the Manager.
func (r *OpenTelemetryCollectorReconciler) SetupWithManager(mgr ctrl.Manager) error {
	return ctrl.NewControllerManagedBy(mgr).
		For(&otelv1.OpenTelemetryCollector{}).
		Complete(r)
}
