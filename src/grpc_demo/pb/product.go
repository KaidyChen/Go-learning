package product

import "context"

var ProductService = &productService{}

type productService struct {
	UnimplementedProdServiceServer
}

func (p *productService) GetProductStock(ctx context.Context, request *ProductRequest) (*ProductResponse, error)  {
	//实现具体的业务逻辑
	stock := p.GetStockById(request.ProdId)
	return &ProductResponse{ProdStock: stock}, nil
}

func (p *productService) GetStockById(id int32) int32 {
	return 100
}
