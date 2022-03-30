package database

import (
	"BugFreeCompilation/go-project/entity"
)

type Database interface {
	Add(post *entity.Post) error
	Delete(post *entity.Post) error
	GetAll() ([]entity.Post, error)
}
