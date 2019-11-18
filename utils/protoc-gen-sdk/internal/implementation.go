package internal

import (
	"fmt"

	"github.com/gogo/protobuf/protoc-gen-gogo/generator"
)

type Implementation struct {
	Package  string
	Header   FileHeader
	Services []Service
}

func NewImplementation() *Implementation {
	return &Implementation{}
}

func (i *Implementation) AddService(info *Service) error {

	if i.Package == "" {
		i.Package = info.Package
	} else if i.Package != "" && i.Package != info.Package {
		return fmt.Errorf("not unique service name '%s' '%s'", i.Package, info.Package)
	}

	i.Services = append(i.Services, *info)
	return nil
}

func (i Implementation) Empty() bool {
	return len(i.Services) == 0
}

func (i Implementation) GetName() string {
	return generator.CamelCase(i.Package)
}
