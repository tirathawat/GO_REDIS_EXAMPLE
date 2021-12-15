package demo_service

import (
	"example/models/entities"
	"example/repositories"
)

type demoService struct {
	Repo repositories.IDemoRepository
}

func NewDemoService(repo repositories.IDemoRepository) demoService {
	return demoService{Repo: repo}
}

func (s demoService) GetDemo() (demo []entities.Demo, err error) {

	demoDB, err := s.Repo.GetDemo()
	if err != nil {
		return nil, err
	}

	for _, d := range demoDB {
		demo = append(demo, entities.Demo{
			ID:       d.ID,
			Name:     d.Name,
			Priority: d.Priority,
		})
	}

	return demo, nil
}
