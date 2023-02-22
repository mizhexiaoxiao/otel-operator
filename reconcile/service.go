package reconcile

import (
	"context"
	"fmt"

	"github.com/mizhexiaoxiao/otel-operator/pkg/collector"
	corev1 "k8s.io/api/core/v1"
	k8serrors "k8s.io/apimachinery/pkg/api/errors"
	"k8s.io/apimachinery/pkg/types"
	"sigs.k8s.io/controller-runtime/pkg/client"
	"sigs.k8s.io/controller-runtime/pkg/controller/controllerutil"
)

func Service(ctx context.Context, params Params) error {

	if len(params.Instance.Spec.Ports) == 0 {
		params.Log.V(1).Info("the instance's configuration didn't yield any ports to open, skipping service", "instance.name", params.Instance.Name, "instance.namespace", params.Instance.Namespace)
	}
	desired := collector.Service(params.Instance)

	if err := expectedService(ctx, desired, params); err != nil {
		return fmt.Errorf("failed to reconcile the expected service sets: %w", err)
	}

	return nil
}

func expectedService(ctx context.Context, desired corev1.Service, params Params) error {

	if err := controllerutil.SetControllerReference(&params.Instance, &desired, params.Scheme); err != nil {
		return fmt.Errorf("failed to set controller reference: %w", err)
	}

	existing := &corev1.Service{}
	nns := types.NamespacedName{Namespace: desired.Namespace, Name: desired.Name}
	err := params.Client.Get(ctx, nns, existing)
	if err != nil && k8serrors.IsNotFound(err) {
		if clientErr := params.Client.Create(ctx, &desired); clientErr != nil {
			return fmt.Errorf("failed to create: %w", clientErr)
		}
		params.Log.V(1).Info("service created", "service.name", desired.Name, "service.namespace", desired.Namespace)
		return nil
	} else if err != nil {
		return fmt.Errorf("failed to get: %w", err)
	}

	updated := existing.DeepCopy()
	if updated.Annotations == nil {
		updated.Annotations = map[string]string{}
	}
	if updated.Labels == nil {
		updated.Labels = map[string]string{}
	}
	updated.ObjectMeta.OwnerReferences = desired.ObjectMeta.OwnerReferences

	for k, v := range desired.ObjectMeta.Annotations {
		updated.ObjectMeta.Annotations[k] = v
	}
	for k, v := range desired.ObjectMeta.Labels {
		updated.ObjectMeta.Labels[k] = v
	}
	updated.Spec.Ports = desired.Spec.Ports
	updated.Spec.Selector = desired.Spec.Selector

	patch := client.MergeFrom(existing)
	if err := params.Client.Patch(ctx, updated, patch); err != nil {
		return fmt.Errorf("failed to apply changes: %w", err)
	}

	params.Log.V(1).Info("applied", "service.name", desired.Name, "service.namespace", desired.Namespace)
	return nil
}
