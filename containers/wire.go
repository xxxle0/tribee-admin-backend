//go:generate wire

package containers

import (
	"github.com/google/wire"
	"github.com/xxxle0/tribee-admin-backend/controllers"
	"github.com/xxxle0/tribee-admin-backend/repositories"
	"github.com/xxxle0/tribee-admin-backend/services"
)

func InitializeAdminController() controllers.AdminController {
	wire.Build(controllers.AdminControllerInit, services.AdminServiceInit, repositories.UserRepositoryInit)
	return controllers.AdminController{}
}
