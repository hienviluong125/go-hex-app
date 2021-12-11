package productmodule

type ProductService interface {
	FindAll() ([]Product, error)
	Save(product Product) error
	FindById(id int) (*Product, error)
	UpdateById(id int, updateProductParams UpdateProduct) error
	DestroyById(id int) error
}

type DefaultProductService struct {
	repo ProductRepository
}

func NewDefaultProductService(repo ProductRepository) *DefaultProductService {
	return &DefaultProductService{repo: repo}
}

func (service *DefaultProductService) FindAll() ([]Product, error) {
	return service.repo.FindAll()
}

func (service *DefaultProductService) Save(product Product) error {
	return service.repo.Save(product)
}

func (service *DefaultProductService) FindById(id int) (*Product, error) {
	return service.repo.FindById(id)
}

func (service *DefaultProductService) UpdateById(id int, updateProductParams UpdateProduct) error {
	return service.repo.UpdateById(id, updateProductParams)
}

func (service *DefaultProductService) DestroyById(id int) error {
	return service.repo.DestroyById(id)
}
