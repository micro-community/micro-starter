package models

type Operation int

const (
	OpGet = iota
	OpAdd
	OpDelete
	OpUpdate
)
