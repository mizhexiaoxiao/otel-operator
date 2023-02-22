package common

func CollectorConfigMapName() string {
	return "otel-collector-conf-sls"
}

func AgentConfigMapName() string {
	return "otel-agent-conf-sls"
}

func CollectorName() string {
	return "otel-collector"
}

func AgentName() string {
	return "otel-agent"
}
