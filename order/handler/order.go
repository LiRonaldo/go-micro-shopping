package handler

import (
	"context"
	"github.com/bwmarrin/snowflake"
	"github.com/micro/go-micro"
	"github.com/micro/go-micro/errors"
	product "go-mico-shopping/product/proto/product"
	"go-micro-shopping/order/model"
	order "go-micro-shopping/order/proto/order"
	"go-micro-shopping/order/repository"
)

/**
在order中引入product的客户端
*/
type Order struct {
	O          *repository.Order
	ProductCli product.ProductService
	Publisher  micro.Publisher
}

func (h *Order) Submit(ctx context.Context, in *order.SubmitRequest, out *order.Response) error {
	productDetail, err := h.ProductCli.Detail(ctx, &product.DetailRequest{Id: in.ProductId})
	if productDetail.Product.Number < 1 {
		return errors.BadRequest("go.micro.srv.order", "库存不足")
	}
	node, err := snowflake.NewNode(1)
	if err != nil {
		return err
	}

	// Generate a snowflake ID.
	orderId := node.Generate().String()
	order := &model.Order{
		Status:    1,
		OrderId:   orderId,
		ProductId: in.ProductId,
		Uid:       in.Uid,
	}
	if err = h.O.Create(order); err != nil {
		return err
	}

	//减库存
	reduce, err := h.ProductCli.ReduceNumber(ctx, &product.ReduceNumberRequest{Id: in.ProductId})
	if reduce == nil || reduce.Code != "200" {
		return errors.BadRequest("go.micro.srv.order", err.Error())
	}
	//异步发送通知给用户订单信息
	if err := h.Publisher.Publish(ctx, in); err != nil {
		return errors.BadRequest("notification", err.Error())
	}
	out.Ode = "200"
	out.Msg = "订单提交成功"
	return nil
}

func (h *Order) OrderDetail(ctx context.Context, in *order.OrderDetailRequest, out *order.Response) error {
	orderDetail, err := h.O.Find(in.OrdeId)

	if err != nil {
		return err
	}

	productDetail, err := h.ProductCli.Detail(ctx, &product.DetailRequest{Id: orderDetail.ProductId})

	out.Ode = "200"
	out.Msg = "订单详情如下：订单号为：" + orderDetail.OrderId + "。购买的产品名字为：" + productDetail.Product.Name
	return nil
}
