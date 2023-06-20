package Handler

import "fmt"

type Type string

const (
	errJson Type = "Error in json"
)

func (t Type) Error() string {
	return string(t)
}

func JsonWrong(e string) error {
	if e == "" {
		return fmt.Errorf("ss %s", errJson)
	}
	return nil
}
