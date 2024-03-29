package collector

import (
	otelv1 "github.com/mizhexiaoxiao/otel-operator/api/v1"
	corev1 "k8s.io/api/core/v1"
)

func Container(instance otelv1.OpenTelemetryCollector) corev1.Container {
	image := instance.Spec.Image
	if len(image) == 0 {
		image = "otel/opentelemetry-collector-contrib:0.73.0"
	}
	command := []string{"/otelcol-contrib", "--config=/conf/otel-collector-config-sls.yaml"}

	volumeMounts := []corev1.VolumeMount{}
	volumeMounts = append(volumeMounts, corev1.VolumeMount{
		Name:      "otel-collector-config-vol",
		MountPath: "/conf",
	})

	return corev1.Container{
		Name:            instance.Name,
		Image:           image,
		Ports:           instance.Spec.Ports,
		Command:         command,
		VolumeMounts:    volumeMounts,
		ImagePullPolicy: instance.Spec.ImagePullPolicy,
		Resources:       instance.Spec.Resources,
	}
}
