package internal

import "fmt"

func NewLabel(namespace string, name string) []byte {
	label := fmt.Sprintf("%s/%s", namespace, name)
	return []byte(label)
}
