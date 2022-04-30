package handle

import (
	"fmt"
	"strings"
)

type TableList []string

func (h *TableList) String() string {
	return fmt.Sprintf("%v", *h)
}

func (h *TableList) Set(s string) error {
	for _, v := range strings.Split(s, ",") {
		*h = append(*h, v)
	}
	return nil
}
