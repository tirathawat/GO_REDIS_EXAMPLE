package repositories

import (
	"example/models/storages"
)

type IDemoRepository interface {
	GetDemo() (demo []storages.DemoDB, err error)
}
