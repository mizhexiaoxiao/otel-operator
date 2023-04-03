# otel-operator
Opentelemetry Operator

## Description
Customized operator, for reference only

## Getting Started

### Running on the cluster
1. Build and push your image to the location specified by `IMG`:
	
```sh
make docker-build docker-push IMG=<some-registry>/otel-operator:tag
docker push <some-registry>/otel-operator:tag
```
	
2. Deploy the controller to the cluster with the image specified by `IMG`:

```sh
make deploy IMG=<some-registry>/otel-operator:tag
```

3. Install Instances of Custom Resources:

```sh
kubectl apply -f config/samples/
```

### Undeploy controller
UnDeploy the controller to the cluster:

```sh
make undeploy
```


## api docs

### otel.mzx.org/v1

Resource Types:
- OpentelemetryCollector
- OpentelemetryAgent

#### OpentelemetryCollector.Spec

OpenTelemetryCollectorSpec defines the desired state of OpenTelemetryCollector(deploy by statefulset)

| Name             | Type              | Description                                                  | Required |
| :--------------- | ----------------- | ------------------------------------------------------------ | -------- |
| image            | string            | Image indicates the container image to use for the OpenTelemetry Collector. | true     |
| imagePullPolicy  | string            | ImagePullPolicy indicates the pull policy to be used for retrieving the container image (Always, Never, IfNotPresent) | false    |
| imagePullSecrets | []object          | Tencent cloud image download key                             | false    |
| replicas         | integer           | Replicas is the number of pod instances for the underlying OpenTelemetry Collector. Format: int32 | true     |
| ports            | []object          | ports defines the ports exposed by the pod, and the operator will infer the ports required by the service and create | false    |
| resources        | []object          | Resources to set on the OpenTelemetry Collector pods.        | false    |
| config           | String            | Config is the raw JSON to be used as the collector's configuration. | true     |
| podAnnotations   | map[string]string | PodAnnotations is the set of annotations that will be attached to Collector pods | false    |
| dnsPolicy        | String            | DNSPolicy defines how a pod's DNS will be configured(ClusterFirstWithHostNet, ClusterFirst, Default, None) | false    |
| hostNetwork      | boolean           | HostNetwork indicates if the pod should run in the host networking namespace. | false    |

#### OpentelemetryAgent.Spec

OpenTelemetryAgentSpec defines the desired state of OpenTelemetryAgent(deploy by daemonset)

| Name             | Type              | Description                                                  | Required |
| :--------------- | ----------------- | ------------------------------------------------------------ | -------- |
| image            | string            | Image indicates the container image to use for the OpenTelemetry Collector. | true     |
| imagePullPolicy  | string            | ImagePullPolicy indicates the pull policy to be used for retrieving the container image (Always, Never, IfNotPresent) | false    |
| imagePullSecrets | []object          | Tencent cloud image download key                             | false    |
| ports            | []object          | ports defines the ports exposed by the pod, and the operator will infer the ports required by the service and create | false    |
| resources        | []object          | Resources to set on the OpenTelemetry Collector pods.        | false    |
| config           | String            | Config is the raw JSON to be used as the collector's configuration. | true     |
| podAnnotations   | map[string]string | PodAnnotations is the set of annotations that will be attached to Collector pods | false    |
| dnsPolicy        | String            | DNSPolicy defines how a pod's DNS will be configured(ClusterFirstWithHostNet, ClusterFirst, Default, None) | false    |
| hostNetwork      | boolean           | HostNetwork indicates if the pod should run in the host networking namespace. | false    |