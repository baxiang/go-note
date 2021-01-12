package services

import (
	"context"
)

type ProductSvc struct {
}

func (*ProductSvc) GetProdStock(ctx context.Context, req *ProductReq) (*ProductResp, error) {
	return &ProductResp{
		ProdStock:   1,
	},nil
}