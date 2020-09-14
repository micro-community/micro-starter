package models

type Operation int

const (
	Query = iota
	Add
	Delete
	Update
	Get
	List
	Watch
)
