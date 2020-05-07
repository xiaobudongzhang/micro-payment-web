package handler

import (
	"context"
	"net/http"
	"strconv"

	"github.com/micro/go-micro/v2/client"
	"github.com/micro/go-micro/v2/util/log"
	auth "github.com/xiaobudongzhang/micro-auth/proto/auth"
	payS "github.com/xiaobudongzhang/micro-payment-srv/proto/payment"
)

var (
	serviceClient payS.PaymentService
	authClient    auth.Service
)


//Error 错误结构体
type Error struct {
	Code   string `json:"code"`
	Detail string `json:"detail"`
}

func Init() {
	serviceClient = payS.NewPaymentService("mu.micro.book.service.payment", client.DefaultClient)
	authClient = auth.NewService("mu.micro.book.service.auth", client.DefaultClient)
}

func PayOrder(w http.ResponseWriter, r *http.Request) {
	if r.Method != "POST" {
		log.Logf("非法请求")
		http.Error(w, "非法请求", 400)
		return
	}

	r.ParseForm()
	orderId, _ := strconv.ParseInt(r.Form.Get("orderId"), 10, 10)
	_, err := serviceClient.PayOrder(context.TODO(), &payS.Request{
		orderId: orderId,
	})

	
}
