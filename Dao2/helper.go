package Dao2

import (
	"fmt"
	"github.com/google/uuid"
)

const (
	posts  = "key"
	postid = "key/%s"
)

func generateKey() (string, string) {
	id := uuid.New().String()

	return fmt.Sprintf(postid, id), id

}

func constructKey(id string) string {

	return fmt.Sprintf(postid, id)
}
