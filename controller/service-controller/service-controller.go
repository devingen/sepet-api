package srvcont

import (
	data_service "github.com/devingen/sepet-api/data-service"
	file_service "github.com/devingen/sepet-api/file-service"
)

// ServiceController implements IServiceController interface by using IDamgaService
type ServiceController struct {
	DataService data_service.ISepetDataService
	FileService file_service.ISepetService
}
