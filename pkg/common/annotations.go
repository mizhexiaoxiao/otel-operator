package common

import (
	"crypto/sha256"
	"fmt"
)

type Generic interface {
	GetAnnotations() map[string]string
	GetPodAnnotations() map[string]string
	GetConfig() string
	GetHostNetwork() bool
}

func Annotations[T Generic](instance T) map[string]string {
	annotations := map[string]string{}
	instanceAnnotations := instance.GetAnnotations()
	instanceConfig := instance.GetConfig()
	if instanceAnnotations != nil {
		for k, v := range instanceAnnotations {
			annotations[k] = v
		}
	}

	annotations["otel-operator-config/sha256"] = getConfigMapSHA(instanceConfig)

	return annotations
}

func PodAnnotations[T Generic](instance T) map[string]string {
	podAnnotations := map[string]string{}
	instancePodAnnotations := instance.GetPodAnnotations()
	instanceConfig := instance.GetConfig()
	for k, v := range instancePodAnnotations {
		podAnnotations[k] = v
	}

	podAnnotations["otel-operator-config/sha256"] = getConfigMapSHA(instanceConfig)

	return podAnnotations
}

func getConfigMapSHA(config string) string {
	h := sha256.Sum256([]byte(config))
	return fmt.Sprintf("%x", h)
}
