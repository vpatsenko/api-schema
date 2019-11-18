package internal

import "strings"

type Method struct {
	Name            string
	InType          string
	OutType         string
	ClientStreaming bool
	ServerStreaming bool
}

func SplitTypeName(src string) (string, string) {

	pkg, name := splitByLast(src, ".")
	if pkg != "" {
		return strings.TrimPrefix(pkg, "."), name
	}

	return "", src
}

func splitByLast(src string, delimeter string) (string, string) {

	i := strings.LastIndex(src, delimeter)
	if i != -1 {
		return src[:i], src[i+1:]
	}

	return "", src
}
