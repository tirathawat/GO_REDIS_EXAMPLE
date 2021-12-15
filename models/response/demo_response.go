package response

import "example/models/entities"

type DemoResponse struct {
	Status string          `json:"status"`
	Demo   []entities.Demo `json:"demo"`
}
