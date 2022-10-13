package id

import "github.com/gofrs/uuid"

var (
	Channel chan string
)

func init() {
	Channel = make(chan string)
	go func() {
		for {
			str, _ := uuid.NewV4()
			Channel <- str.String()
		}
	}()
}
