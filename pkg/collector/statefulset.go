package collector

import (
	appsv1 "k8s.io/api/apps/v1"
	corev1 "k8s.io/api/core/v1"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"

	otelv1 "github.com/mizhexiaoxiao/otel-operator/api/v1"
	"github.com/mizhexiaoxiao/otel-operator/pkg/common"
)

// StatefulSet builds the statefulset for the given instance.
func StatefulSet(instance otelv1.OpenTelemetryCollector) appsv1.StatefulSet {
	labels := map[string]string{}
	labels["app"] = "opentelemetry"
	labels["component"] = common.CollectorName()

	annotations := common.Annotations(instance)
	podAnnotations := common.PodAnnotations(instance)

	return appsv1.StatefulSet{
		ObjectMeta: metav1.ObjectMeta{
			Name:        common.CollectorName(),
			Namespace:   instance.Namespace,
			Labels:      labels,
			Annotations: annotations,
		},
		Spec: appsv1.StatefulSetSpec{
			ServiceName: common.CollectorName(),
			Selector: &metav1.LabelSelector{
				MatchLabels: labels,
			},
			Template: corev1.PodTemplateSpec{
				ObjectMeta: metav1.ObjectMeta{
					Name:        common.CollectorName(),
					Namespace:   instance.Namespace,
					Labels:      labels,
					Annotations: podAnnotations,
				},
				Spec: corev1.PodSpec{
					Containers:       []corev1.Container{Container(instance)},
					Volumes:          Volumes(),
					DNSPolicy:        common.DNSPolicy(instance),
					HostNetwork:      instance.Spec.HostNetwork,
					ImagePullSecrets: instance.Spec.ImagePullSecrets,
				},
			},
			Replicas: instance.Spec.Replicas,
		},
	}
}
