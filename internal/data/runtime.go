package data

import (
	"fmt"
)

type Runtime int

func (r Runtime) MarshalJSON() ([]byte, error) {
	jsonValue := fmt.Sprintf(`"%v mins"`, r)

	return []byte(jsonValue), nil
}
