package collector

import (
	"fmt"

	otelv1 "github.com/mizhexiaoxiao/otel-operator/api/v1"
	"github.com/mizhexiaoxiao/otel-operator/pkg/common"
	corev1 "k8s.io/api/core/v1"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
)

func Service(instance otelv1.OpenTelemetryCollector) corev1.Service {
	labels := map[string]string{}
	labels["app"] = "opentelemetry"
	labels["component"] = common.CollectorName()

	selectorLabels := map[string]string{}
	selectorLabels["component"] = common.CollectorName()
	return corev1.Service{
		ObjectMeta: metav1.ObjectMeta{
			Name:        common.CollectorName(),
			Namespace:   instance.Namespace,
			Labels:      labels,
			Annotations: instance.Annotations,
		},
		Spec: corev1.ServiceSpec{
			Selector:  selectorLabels,
			ClusterIP: "None", // headless service
			Ports:     ServicePorts(instance.Spec.Ports),
		},
	}
}

func ServicePorts(ports []corev1.ContainerPort) []corev1.ServicePort {
	servicePorts := []corev1.ServicePort{}

	for _, k := range ports {
		servicePorts = append(servicePorts, corev1.ServicePort{
			Name:     fmt.Sprintf("port-%d", k.ContainerPort),
			Protocol: k.Protocol,
			Port:     k.ContainerPort,
		})
	}

	return servicePorts
}
