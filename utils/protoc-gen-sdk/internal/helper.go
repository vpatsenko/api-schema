package internal

import (
	"sort"
	"strings"

	"github.com/gogo/protobuf/protoc-gen-gogo/descriptor"
)

func NewFileName(src string) string {

	const suffix = ".proto"
	if !strings.HasSuffix(src, suffix) {
		return src
	}

	newSuffix := ".pb.go"
	return strings.TrimSuffix(src, suffix) + newSuffix
}

func NewPackageName(file *descriptor.FileDescriptorProto) string {

	retval := file.GetPackage()

	if file.Options != nil && file.Options.GoPackage != nil {
		retval = *file.Options.GoPackage
		if i := strings.LastIndex(retval, "/"); i != -1 {
			retval = retval[i+1:]
		}
	}

	return retval
}

func NewPath(file *descriptor.FileDescriptorProto) string {

	var retval string

	if file.Options != nil && file.Options.GoPackage != nil {
		if strings.Contains(file.Options.GetGoPackage(), "/") {
			retval = *file.Options.GoPackage
		}
	}

	return retval
}

func StringPtr(src string) *string {
	return &src
}

func FindString(src []string, val string) int {

	idx := sort.Search(len(src), func(i int) bool {
		return strings.Compare(src[i], val) >= 0
	})
	if idx < len(src) && src[idx] == val {
		return idx
	}

	return -1
}

func SortStrings(vals []string) []string {

	sort.Slice(vals, func(i, j int) bool {
		return strings.Compare(vals[i], vals[j]) < 0
	})

	return vals
}

func AddCustomImport(dest *[]string, name string) {

	if name != "" {
		SortStrings(*dest)
		if cIdx := FindString(*dest, name); cIdx == -1 {
			*dest = append(*dest, name)
			SortStrings(*dest)
		}
	}
}
