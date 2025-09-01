package simple

type FooBarService struct {
	*FooService
	*BarService
}

func NewFooBarService(FooService *FooService, BarService *BarService) *FooBarService {
	return &FooBarService{FooService: FooService, BarService: BarService}
}