package agent

import (
	otelv1 "github.com/mizhexiaoxiao/otel-operator/api/v1"
	"github.com/mizhexiaoxiao/otel-operator/pkg/common"
	appsv1 "k8s.io/api/apps/v1"
	corev1 "k8s.io/api/core/v1"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	"k8s.io/apimachinery/pkg/util/intstr"
)

func Daemonset(instance otelv1.OpenTelemetryAgent) appsv1.DaemonSet {
	labels := map[string]string{}
	labels["app"] = "opentelemetry"
	labels["component"] = common.AgentName()

	annotations := common.Annotations(instance)
	podAnnotations := common.PodAnnotations(instance)

	return appsv1.DaemonSet{
		ObjectMeta: metav1.ObjectMeta{
			Name:        instance.Name,
			Namespace:   instance.Namespace,
			Labels:      labels,
			Annotations: annotations,
		},
		Spec: appsv1.DaemonSetSpec{
			Selector: &metav1.LabelSelector{
				MatchLabels: labels,
			},
			UpdateStrategy: appsv1.DaemonSetUpdateStrategy{
				Type: "RollingUpdate",
				RollingUpdate: &appsv1.RollingUpdateDaemonSet{
					MaxUnavailable: &intstr.IntOrString{
						Type:   intstr.String,
						StrVal: "100%",
					},
				},
			},
			Template: corev1.PodTemplateSpec{
				ObjectMeta: metav1.ObjectMeta{
					Name:        common.AgentName(),
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
		},
	}
}
