package common

import (
	corev1 "k8s.io/api/core/v1"
)

func DNSPolicy[T Generic](instance T) corev1.DNSPolicy {
	dnsPolicy := corev1.DNSClusterFirst
	hostNetwork := instance.GetHostNetwork()
	if hostNetwork {
		dnsPolicy = corev1.DNSClusterFirstWithHostNet
	}
	return dnsPolicy
}
