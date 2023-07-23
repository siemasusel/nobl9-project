package randomorg

import "fmt"

type ClientError struct {
	Code    int
	Message string
}

func (e ClientError) Error() string {
	return fmt.Sprintf("randomorg api responde with code %d and message '%s'", e.Code, e.Message)
}
