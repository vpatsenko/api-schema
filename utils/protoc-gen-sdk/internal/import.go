package internal

import "fmt"

type Import string

func (i Import) String() string {

	tmp := string(i)

	_, alias := splitByLast(tmp, "/")

	if alias == "" {
		return fmt.Sprintf("%q", tmp)
	}

	return fmt.Sprintf("%s %q", alias, tmp)
}
