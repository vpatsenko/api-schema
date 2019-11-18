package internal

import "fmt"

type Import struct {
	Standard bool
	Alias    string
	Path     string
}

func NewImport(standard bool, alias, path string) *Import {
	return &Import{
		Standard: standard,
		Alias:    alias,
		Path:     path,
	}
}

func (i Import) String() string {
	if i.Alias == "" {
		return fmt.Sprintf("%q", i.Path)
	}
	return fmt.Sprintf("%s %q", i.Alias, i.Path)
}
