package handle

import (
	"fmt"
	"strings"
)

type SqlList map[string]string

func (f SqlList) String() string {
	return fmt.Sprintf("%v", map[string]string(f))
}

func (f SqlList) Set(value string) error {
	split := strings.SplitN(value, "=", 2)
	f[split[0]] = split[1]
	return nil
}
