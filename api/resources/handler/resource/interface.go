package resource

import (
	"github.com/zopdev/zopdev/api/resources/service/resource"
	"gofr.dev/pkg/gofr"

	"github.com/zopdev/zopdev/api/resources/models"
)

type Service interface {
	GetAll(ctx *gofr.Context, id int64, resourceType []string) ([]models.Instance, error)
	SyncResources(ctx *gofr.Context, id int64) ([]models.Instance, error)
	ChangeState(ctx *gofr.Context, resDetails resource.ResourceDetails) error
}
