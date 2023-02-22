package reconcile

import (
	"context"
	"fmt"

	"github.com/mizhexiaoxiao/otel-operator/pkg/agent"
	appsv1 "k8s.io/api/apps/v1"
	k8serrors "k8s.io/apimachinery/pkg/api/errors"
	"k8s.io/apimachinery/pkg/types"
	"sigs.k8s.io/controller-runtime/pkg/client"
	"sigs.k8s.io/controller-runtime/pkg/controller/controllerutil"
)

func Daemonset(ctx context.Context, params Params) error {

	desired := agent.Daemonset(params.AgentInstace)

	if err := expectedDaemonset(ctx, desired, params); err != nil {
		return fmt.Errorf("failed to reconcile the expected daemonset: %w", err)
	}

	return nil
}

func expectedDaemonset(ctx context.Context, desired appsv1.DaemonSet, params Params) error {

	if err := controllerutil.SetControllerReference(&params.AgentInstace, &desired, params.Scheme); err != nil {
		return fmt.Errorf("failed to set controller reference: %w", err)
	}

	existing := &appsv1.DaemonSet{}
	nns := types.NamespacedName{Namespace: desired.Namespace, Name: desired.Name}
	err := params.Client.Get(ctx, nns, existing)
	if err != nil && k8serrors.IsNotFound(err) {
		if clientErr := params.Client.Create(ctx, &desired); clientErr != nil {
			return fmt.Errorf("failed to create: %w", clientErr)
		}
		params.Log.V(1).Info("daemonset created", "daemonset.name", desired.Name, "daemonset.namespace", desired.Namespace)
		return nil
	} else if err != nil {
		return fmt.Errorf("failed to get: %w", err)
	}

	updated := existing.DeepCopy()
	updated.Name = desired.Name
	updated.Namespace = desired.Namespace
	if updated.Annotations == nil {
		updated.Annotations = map[string]string{}
	}
	if updated.Labels == nil {
		updated.Labels = map[string]string{}
	}
	updated.Spec = desired.Spec
	updated.ObjectMeta.OwnerReferences = desired.ObjectMeta.OwnerReferences
	for k, v := range desired.ObjectMeta.Annotations {
		updated.ObjectMeta.Annotations[k] = v
	}
	for k, v := range desired.ObjectMeta.Labels {
		updated.ObjectMeta.Labels[k] = v
	}

	patch := client.MergeFrom(existing)
	if err := params.Client.Patch(ctx, updated, patch); err != nil {
		return fmt.Errorf("failed to apply changes: %w", err)
	}
	params.Log.V(1).Info("applied", "daemonset.name", desired.Name, "daemonset.namespace", desired.Namespace)
	return nil
}
