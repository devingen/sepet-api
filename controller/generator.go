package controller

import (
	service_controller "github.com/devingen/sepet-api/controller/service-controller"
	data_service "github.com/devingen/sepet-api/data-service"
	file_service "github.com/devingen/sepet-api/file-service"
	is "github.com/devingen/sepet-api/interceptor-service"
)

// New generates new ServiceController
func New(dataService data_service.ISepetDataService, fileService file_service.ISepetService, interceptorService is.ISepetInterceptorService) *service_controller.ServiceController {
	return &service_controller.ServiceController{
		DataService:        dataService,
		FileService:        fileService,
		InterceptorService: interceptorService,
	}
}
