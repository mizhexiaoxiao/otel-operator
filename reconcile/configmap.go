package reconcile

import (
	"context"
	"errors"
	"fmt"
	"reflect"

	"github.com/mizhexiaoxiao/otel-operator/pkg/common"
	"gopkg.in/yaml.v2"
	corev1 "k8s.io/api/core/v1"
	k8serrors "k8s.io/apimachinery/pkg/api/errors"
	v1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	"k8s.io/apimachinery/pkg/types"
	"sigs.k8s.io/controller-runtime/pkg/client"
	"sigs.k8s.io/controller-runtime/pkg/controller/controllerutil"
)

func CollectorConfigMap(ctx context.Context, params Params) error {
	desired, err := desiredCollectorConfigMap(params)
	if err != nil {
		return err
	}
	if err := expectedConfigMap(ctx, desired, params); err != nil {
		return fmt.Errorf("failed to reconcile the expected configmap sets: %w", err)
	}

	return nil
}

func AgentConfigMap(ctx context.Context, params Params) error {
	desired, err := desiredAgentConfigMap(params)
	if err != nil {
		return err
	}
	if err := expectedConfigMap(ctx, desired, params); err != nil {
		return fmt.Errorf("failed to reconcile the expected configmap sets: %w", err)
	}

	return nil
}

func desiredCollectorConfigMap(params Params) (corev1.ConfigMap, error) {
	configStr, err := ConfigFromString(params.Instance.Spec.Config)
	if err != nil {
		return corev1.ConfigMap{}, err
	}
	name := common.CollectorConfigMapName()
	cm := corev1.ConfigMap{
		ObjectMeta: v1.ObjectMeta{
			Name:      name,
			Namespace: params.Instance.Namespace,
		},
		Data: map[string]string{
			"otel-collector-config": configStr,
		},
	}
	return cm, nil
}

func desiredAgentConfigMap(params Params) (corev1.ConfigMap, error) {
	configStr, err := ConfigFromString(params.AgentInstace.Spec.Config)
	if err != nil {
		return corev1.ConfigMap{}, err
	}
	name := common.AgentConfigMapName()
	cm := corev1.ConfigMap{
		ObjectMeta: v1.ObjectMeta{
			Name:      name,
			Namespace: params.AgentInstace.Namespace,
		},
		Data: map[string]string{
			"otel-agent-config": configStr,
		},
	}
	return cm, nil
}

func expectedConfigMap(ctx context.Context, desired corev1.ConfigMap, params Params) error {
	switch desired.Name {
	case common.AgentConfigMapName():
		if err := controllerutil.SetControllerReference(&params.AgentInstace, &desired, params.Scheme); err != nil {
			return fmt.Errorf("failed to set controller reference: %w", err)
		}
	case common.CollectorConfigMapName():
		if err := controllerutil.SetControllerReference(&params.Instance, &desired, params.Scheme); err != nil {
			return fmt.Errorf("failed to set controller reference: %w", err)
		}
	}

	existing := &corev1.ConfigMap{}
	nns := types.NamespacedName{Namespace: desired.Namespace, Name: desired.Name}
	err := params.Client.Get(ctx, nns, existing)
	if err != nil && k8serrors.IsNotFound(err) {
		if clientErr := params.Client.Create(ctx, &desired); clientErr != nil {
			return fmt.Errorf("failed to create %w", clientErr)
		}
		params.Log.V(1).Info("configmap created", "configmap.name", desired.Name, "configmap.namespace", desired.Namespace)
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

	updated.Data = desired.Data
	updated.BinaryData = desired.BinaryData
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

	if !reflect.DeepEqual(existing.Data, desired.Data) {
		params.Log.V(1).Info("Opentelemetry ConfigMap Changed", "Kind", desired.Kind)
	}

	params.Log.V(1).Info("applied", "configmap.name", desired.Name, "configmap.namespace", desired.Namespace)
	return nil
}

func ConfigFromString(configStr string) (string, error) {
	config := make(map[interface{}]interface{})
	if err := yaml.Unmarshal([]byte(configStr), &config); err != nil {
		return "", errors.New("couldn't parse the opentelemetry-collector configuration")
	}
	out, err := yaml.Marshal(config)
	if err != nil {
		return "", err
	}
	return string(out), nil
}
