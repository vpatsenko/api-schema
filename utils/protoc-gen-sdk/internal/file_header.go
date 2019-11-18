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
	Imports    map[Import]struct{}
}

func NewFileHeader(file *descriptor.FileDescriptorProto) *FileHeader {

	return &FileHeader{
		SourceFile: file.GetName(),
		FileName:   NewFileName(file.GetName()),
		FilePath:   NewPath(file),
		PkgName:    NewPackageName(file),
		Imports:    make(map[Import]struct{}),
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

	for k := range val.Imports {
		h.AddCustomImport(k)
	}

	return nil
}

func (h *FileHeader) FullPath() string {

	if h.FilePath != "" {
		return filepath.Join(h.FilePath, h.FileName)
	}

	return h.FileName
}

func (h *FileHeader) AddCustomImport(val Import) {

	if h.Imports == nil {
		h.Imports = make(map[Import]struct{})
	}

	if len(val) > 0 {
		h.Imports[val] = struct{}{}
	}
}

func (h *FileHeader) GetImports() (list []string) {

	for k := range h.Imports {
		list = append(list, k.String())
	}

	return
}
