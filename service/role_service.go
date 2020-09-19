package service

import (
	"github.com/micro-community/auth/repository"
)

//RoleService for sdb
type RoleService struct {
	repo repository.IRole
}

func NewRole(repo repository.IRole) *RoleService {
	return &RoleService{
		repo: repo,
	}
}
