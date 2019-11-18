package internal

import (
	"errors"
	"path/filepath"

	"github.com/gogo/protobuf/protoc-gen-gogo/descriptor"
)

type FileHeader struct {
	SourceFile string
	FileName   string
	FilePath   string
	PkgName    string
	Imports    []Import
}

func NewFileHeader(file *descriptor.FileDescriptorProto) *FileHeader {

	return &FileHeader{
		SourceFile: file.GetName(),
		FileName:   NewFileName(file.GetName()),
		FilePath:   NewPath(file),
		PkgName:    NewPackageName(file),
	}
}

func (h *FileHeader) Merge(val *FileHeader) error {

	if h.PkgName == "" {
		h.SourceFile = val.SourceFile
		h.FileName = val.FileName
		h.FilePath = val.FilePath
		h.PkgName = val.PkgName
		h.Imports = val.Imports
		return nil
	}

	if h.PkgName != val.PkgName {
		return errors.New("invalid package name")
	}

	if len(val.Imports) > 0 {
		h.Imports = append(h.Imports, val.Imports...)
	}

	return nil
}

func (h *FileHeader) FullPath() string {

	if h.FilePath != "" {
		return filepath.Join(h.FilePath, h.FileName)
	}

	return h.FileName
}

func (h *FileHeader) AddCustomImports(imports ...string) {

	for _, item := range imports {
		_, alias := splitByLast(item, "/")
		h.Imports = append(h.Imports, *NewImport(false, alias, item))
	}
}
