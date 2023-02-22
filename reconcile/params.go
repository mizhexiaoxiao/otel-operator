package reconcile

import (
	"github.com/go-logr/logr"
	otelv1 "github.com/mizhexiaoxiao/otel-operator/api/v1"
	"k8s.io/apimachinery/pkg/runtime"
	"sigs.k8s.io/controller-runtime/pkg/client"
)

type Params struct {
	Client       client.Client
	Scheme       *runtime.Scheme
	Log          logr.Logger
	Instance     otelv1.OpenTelemetryCollector
	AgentInstace otelv1.OpenTelemetryAgent
}
