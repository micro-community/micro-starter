package models

import "time"

type ModelExtension struct {
	CreatedAt time.Time `json:"createdAt"`
	UpdatedAt time.Time `json:"updatedAt"`
	DeletedAt time.Time `json:"deletedAt"`
	IsSoftDel bool      `json:"isSoftDelete"` //软删除
}
