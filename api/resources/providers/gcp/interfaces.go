package gcp

import (
	"time"

	"gofr.dev/pkg/gofr"

	"github.com/zopdev/zopdev/api/resources/models"
)

type SQLClient interface {
	InstanceLister
	Idler
}

type MetricsClient interface {
	TimeSeriesLister
}

type InstanceLister interface {
	GetAllInstances(ctx *gofr.Context, projectID string) ([]models.Resource, error)
}

type TimeSeriesLister interface {
	GetTimeSeries(ctx *gofr.Context, start, end time.Time, projectID, filter string) ([]models.Metric, error)
}

type Idler interface {
	StartInstance(ctx *gofr.Context, projectID, instanceName string) error
	StopInstance(ctx *gofr.Context, projectID, instanceName string) error
}
