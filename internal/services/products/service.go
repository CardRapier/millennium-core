package products


type ProductsService struct {
	productsStore *ProductsStore
}

func NewProductsService(productsStore *ProductsStore) *ProductsService {
	return &ProductsService{productsStore}
}