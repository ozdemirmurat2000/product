package defotanim

type IDefoTanimService interface {
	GetDefoTanimList(request DefoTanimRequest) ([]DefoTanimResponse, error)
	GetDefoByName(defoIsmi string) (DefoTanimResponse, error)
}

type DefoTanimServiceImpl struct {
	repo IDefoTanimRepository
}

func NewDefoTanimServiceImpl(repo IDefoTanimRepository) IDefoTanimService {
	return &DefoTanimServiceImpl{repo: repo}
}

func (s *DefoTanimServiceImpl) GetDefoTanimList(request DefoTanimRequest) ([]DefoTanimResponse, error) {
	return s.repo.GetDefoTanimList(request)
}

func (s *DefoTanimServiceImpl) GetDefoByName(defoIsmi string) (DefoTanimResponse, error) {
	return s.repo.GetDefoByName(defoIsmi)
}
