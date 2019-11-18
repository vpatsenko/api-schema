package internal

import (
	"github.com/gogo/protobuf/protoc-gen-gogo/descriptor"
	"github.com/gogo/protobuf/protoc-gen-gogo/generator"
)

type Service struct {
	Package     string
	ServiceName string
	Methods     []Method
}

func NewService(file *descriptor.FileDescriptorProto, svc *descriptor.ServiceDescriptorProto) *Service {

	return &Service{
		Package:     NewPackageName(file),
		ServiceName: generator.CamelCase(svc.GetName()),
	}
}

func (s *Service) GetObjectName() string {
	return s.ServiceName + "Client"
}
