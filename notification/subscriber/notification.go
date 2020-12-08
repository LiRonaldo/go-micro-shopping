package subscriber

import (
	"context"
	"fmt"
	_ "github.com/LiRonaldo/l-log"
	log "github.com/LiRonaldo/l-log"
	notification "go-micro-shopping/notification/proto/notification"
)

type Notification struct{}

func (e *Notification) Handle(ctx context.Context, req *notification.SubmitRequest) error {
	log.Info(fmt.Sprintf("Handler Received message: ID为%v 的用户购买了商品ID为：%v 的物品", req.Uid, req.ProductId))
	return nil
}
