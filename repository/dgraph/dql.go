package dgraph

import (
	"fmt"

	"github.com/micro-community/auth/models"
)

type Count struct {
	Count int `json:"count"`
}

type UID struct {
	UID string `json:"uid"`
}

type Root struct {
	Count []Count  `json:"counts"`
	UID   []string `json:"uids"`
}

type RoleResult struct {
	Roles []models.Role `json:"roles"`
}

// GetUserExistQuery return
func GetUserExistQuery(id string) string {
	return fmt.Sprintf(`
	{
		find(func: type(User)) @filter(eq(person.id, %s)) {
			uid
		}
	}
	`, id)

}

// GetUserResourceQuery return
func GetUserResourceQuery(id string) string {
	return fmt.Sprintf(`
	{
		find(func: type(User)) @filter(eq(person.id, %s)) @normalize {
			role {
				resource {
					resource.id
					resource.name
				}
			}
		}
	}
	`, id)

}
