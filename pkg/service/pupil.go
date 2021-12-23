package service

import (
  "github.com/Sedokovka/simple-to-do-app"
  "github.com/Sedokovka/simple-to-do-app/pkg/repository"
)

type PupilService struct {
  repo repository.Pupil
}

func NewPupilService(repo repository.Pupil) *PupilService {
  return &PupilService{repo: repo}
}

func (s *PupilService) CreatePupil(userId int, pupil crud.Pupil) (int, error) {
  return s.repo.CreatePupil(userId, pupil)
}

func (s *PupilService) GetPupilById(userId, pupilId int) (crud.Pupil, error) {
  return s.repo.GetPupilById(userId, pupilId)
}


func (s *PupilService) GetAllPupils(userId int) ([]crud.Pupil, error) {
  return s.repo.GetAllPupils(userId)
}
