package agent

import (
	corev1 "k8s.io/api/core/v1"
)

func Volumes() []corev1.Volume {
	volumes := []corev1.Volume{{
		Name: "otel-agent-config-vol",
		VolumeSource: corev1.VolumeSource{
			ConfigMap: &corev1.ConfigMapVolumeSource{
				LocalObjectReference: corev1.LocalObjectReference{Name: "otel-agent-conf-sls"},
				Items: []corev1.KeyToPath{{
					Key:  "otel-agent-config",
					Path: "otel-agent-config-sls.yaml",
				}},
			},
		},
	}}
	return volumes
}
