package v1

import "github.com/musadhiekt/grpc-clean/models"

// this is an interface for repo models

type RepoInterface interface {
	// create a user with data supplied
	Create(models.User) error

	Get(id string) (models.User, error)
	Update(models.User) error
	Delete(id string) error
}

type UsecaseInterface interface {
	Create(models.User) error
	Get(id string) (models.User, error)
	Update(models.User) error
	Delete(id string) error
}
