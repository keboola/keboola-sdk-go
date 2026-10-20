package keboola

import (
	"time"

	otelMetric "go.opentelemetry.io/otel/metric"
	otelTrace "go.opentelemetry.io/otel/trace"

	"github.com/keboola/keboola-sdk-go/v2/pkg/client"
)

type apiConfig struct {
	client           *client.Client
	onSuccessTimeout time.Duration
	tracerProvider   otelTrace.TracerProvider
	meterProvider    otelMetric.MeterProvider
	serviceURLs      map[ServiceType]string
}

type APIOption func(c *apiConfig)

func newAPIConfig(opts []APIOption) apiConfig {
	cfg := apiConfig{
		onSuccessTimeout: time.Minute,
	}
	for _, opt := range opts {
		opt(&cfg)
	}
	return cfg
}

func WithClient(cl *client.Client) APIOption {
	return func(c *apiConfig) {
		c.client = cl
	}
}

func WithOnSuccessTimeout(timeout time.Duration) APIOption {
	return func(c *apiConfig) {
		c.onSuccessTimeout = timeout
	}
}

func WithTracerProvider(v otelTrace.TracerProvider) APIOption {
	return func(c *apiConfig) {
		c.tracerProvider = v
	}
}

func WithMeterProvider(v otelMetric.MeterProvider) APIOption {
	return func(c *apiConfig) {
		c.meterProvider = v
	}
}

// WithServiceURL makes requests of the service go to baseURL instead of the URL the index advertises for it. It is how
// a service is reached by an address the index does not know, for example the in-cluster address of the Metastore.
// The index must still be given (NewPublicAPIFromIndex), the other services keep using it.
func WithServiceURL(service ServiceType, baseURL string) APIOption {
	return func(c *apiConfig) {
		if c.serviceURLs == nil {
			c.serviceURLs = make(map[ServiceType]string)
		}
		c.serviceURLs[service] = baseURL
	}
}
