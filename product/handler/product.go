package handler

import (
	"context"
	"fmt"
	"github.com/go-log/log"
	_ "github.com/go-sql-driver/mysql"
	"go-micro-shopping/product/model"
	product "go-micro-shopping/product/proto/product"
	"go-micro-shopping/product/repository"
)

type Product struct {
	Pro *repository.Product
}

func (h *Product) SerchR(ctx context.Context, in *product.SerchRequest, out *product.SerchResponse) error {
	var products []*product.Product
	if err := h.Pro.Repo.Where("name like ?", "%"+in.Name+"%").Limit(10).Find(&products).Error; err != nil {
		return err
	}
	out.Product = products
	out.Code = "200"
	out.Msg = fmt.Sprintf("总共%v条数据", len(products))
	return nil
}

func (h *Product) Detail(ctx context.Context, in *product.DetailRequest, out *product.DetailResponse) error {
	product := &product.Product{}
	if err := h.Pro.Repo.Where("id=?", in.Id).First(product).Error; err != nil {
		return err
	}

	out.Code = "200"
	out.Msg = "商品详细信息如下："
	out.Product = product
	return nil
}

func (h *Product) ReduceNumber(ctx context.Context, in *product.ReduceNumberRequest, out *product.ReduceNumberResponse) error {
	log.Log("Received Product.Detail request")

	var product = &product.Product{}
	if err := h.Pro.Repo.Where("id = ?", in.Id).First(product).Error; err != nil {
		return err
	}

	product.Number -= 1
	log.Log("库存数量为:", product.Number)
	if err := h.Pro.Repo.Model(&model.Product{}).Where("id = ?", product.Id).Update("number", product.Number).Error; err != nil {
		return err
	}

	out.Code = "200"
	out.Msg = fmt.Sprintf("库存更新成功,更新后的数量为%v", product.Number)
	return nil
}
