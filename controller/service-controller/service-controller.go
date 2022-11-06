package srvcont

import (
	data_service "github.com/devingen/sepet-api/data-service"
	file_service "github.com/devingen/sepet-api/file-service"
	is "github.com/devingen/sepet-api/interceptor-service"
)

// ServiceController implements IServiceController interface by using IDamgaService
type ServiceController struct {
	DataService        data_service.ISepetDataService
	FileService        file_service.ISepetService
	InterceptorService is.ISepetInterceptorService
}
