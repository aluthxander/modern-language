package simple


type FooRepository struct {

}

func NewFooRepository() *FooRepository {
	return &FooRepository{}
}

type FooService struct {
	*FooRepository
}

func NewFooService(FooRepository *FooRepository) *FooService {
	return &FooService{FooRepository: FooRepository}
}