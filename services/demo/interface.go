package demo_service

import "example/models/entities"

type IDemoService interface {
	GetDemo() (demo []entities.Demo, err error)
}
