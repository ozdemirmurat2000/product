package uretim

type IUretimService interface {
	GetUretimList(request UretimRequest) ([]UretimResponse, error)
	DeleteUretim(id int) error
	AddUretim(request UretimAddRequest) error
}

type UretimServiceImpl struct {
	repo IUretimRepository
}

func NewUretimServiceImpl(repo IUretimRepository) IUretimService {
	return &UretimServiceImpl{repo: repo}
}

func (s *UretimServiceImpl) GetUretimList(request UretimRequest) ([]UretimResponse, error) {
	return s.repo.GetUretimList(request)
}

func (s *UretimServiceImpl) DeleteUretim(id int) error {
	err := s.repo.DeleteUploads(id)
	if err != nil {
		return err
	}
	return s.repo.DeleteUretim(id)
}

func (s *UretimServiceImpl) AddUretim(request UretimAddRequest) error {
	id, err := s.repo.AddUretim(request)
	if err != nil {
		return err
	}
	if err := s.addUretimUploads(request, id); err != nil {
		return err
	}
	return nil
}

func (s *UretimServiceImpl) addUretimUploads(request UretimAddRequest, id int) error {

	return nil
}
