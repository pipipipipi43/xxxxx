// Code generated by protoc-gen-go-register. DO NOT EDIT.
// Sources: resource.proto

package pb

import (
	reflect "reflect"

	transport "github.com/erda-project/erda-infra/pkg/transport"
)

// RegisterResourceServiceImp resource.proto
func RegisterResourceServiceImp(regester transport.Register, srv ResourceServiceServer, opts ...transport.ServiceOption) {
	_ops := transport.DefaultServiceOptions()
	for _, op := range opts {
		op(_ops)
	}
	RegisterResourceServiceHandler(regester, ResourceServiceHandler(srv), _ops.HTTP...)
	RegisterResourceServiceServer(regester, srv, _ops.GRPC...)
}

// ServiceNames return all service names
func ServiceNames(svr ...string) []string {
	return append(svr,
		"erda.msp.resource.ResourceService",
	)
}

var (
	resourceServiceClientType  = reflect.TypeOf((*ResourceServiceClient)(nil)).Elem()
	resourceServiceServerType  = reflect.TypeOf((*ResourceServiceServer)(nil)).Elem()
	resourceServiceHandlerType = reflect.TypeOf((*ResourceServiceHandler)(nil)).Elem()
)

// ResourceServiceClientType .
func ResourceServiceClientType() reflect.Type { return resourceServiceClientType }

// ResourceServiceServerType .
func ResourceServiceServerType() reflect.Type { return resourceServiceServerType }

// ResourceServiceHandlerType .
func ResourceServiceHandlerType() reflect.Type { return resourceServiceHandlerType }

func Types() []reflect.Type {
	return []reflect.Type{
		// client types
		resourceServiceClientType,
		// server types
		resourceServiceServerType,
		// handler types
		resourceServiceHandlerType,
	}
}
