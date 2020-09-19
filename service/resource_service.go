package service

import (
	"github.com/micro-community/auth/repository"
)

//ResourceService for sdb
type ResourceService struct {
	repo repository.IResource
}

// NewResource return ResourceService
func NewResource(repo repository.IResource) *ResourceService {
	return &ResourceService{
		repo: repo,
	}
}
