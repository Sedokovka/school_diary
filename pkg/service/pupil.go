package service

type PupilService struct {
  repo repository.Pupil
}

func NewPupilService(repo repository.Pupil) *PupilService {
  return &PupilService{repo: repo}
}

func (s *PupilService) CreatePupil(userId int, pupil crud.Pupil) (int, error) {
  return s.repo.CreatePupil(userId, pupil)
}
