package poststore

import (
	"fmt"

	"github.com/google/uuid"
)

const (
	posts       = "configs/%s/%s"
	postsLabels = "configs/%s/%s/%s"
	all         = "configs"
)

func generateKey(version string, labels string) (string, string) {
	id := uuid.New().String()
	if labels != "" {
		return fmt.Sprintf(postsLabels, id, version, labels), id
	} else {
		return fmt.Sprintf(posts, id, version), id
	}

}

func constructKey(id string, version string, labels string) string {
	if labels != "" {
		return fmt.Sprintf(postsLabels, id, version, labels)
	} else {
		return fmt.Sprintf(posts, id, version)
	}

}
