package idx

import "github.com/google/uuid"

func Uuid() string {
	gen, err := uuid.NewUUID()
	if err != nil {
		panic(err)
	}
	return gen.String()
}
