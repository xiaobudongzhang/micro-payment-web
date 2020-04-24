package handler

import (
	"context"
	"encoding/json"
	"net/http"
	"strconv"
	"time"

	"github.com/micro/go-micro/v2/client"
	"github.com/micro/go-micro/v2/util/log"
	auth "github.com/xiaobudongzhang/micro-auth/proto/auth"
	payS "github.com/xiaobudongzhang/micro-payment-srv/proto/payment"
)

var (
	serviceClient payS.PaymentService
	authClient    auth.Service
)

func PaymentCall(w http.ResponseWriter, r *http.Request) {
	// decode the incoming request as json
	var request map[string]interface{}
	if err := json.NewDecoder(r.Body).Decode(&request); err != nil {
		http.Error(w, err.Error(), 500)
		return
	}

	// call the backend service
	paymentClient := payment.NewPaymentService("mu.micro.book.service.payment", client.DefaultClient)
	rsp, err := paymentClient.Call(context.TODO(), &payment.Request{
		Name: request["name"].(string),
	})
	if err != nil {
		http.Error(w, err.Error(), 500)
		return
	}

	// we want to augment the response
	response := map[string]interface{}{
		"msg": rsp.Msg,
		"ref": time.Now().UnixNano(),
	}

	// encode and write the response as json
	if err := json.NewEncoder(w).Encode(response); err != nil {
		http.Error(w, err.Error(), 500)
		return
	}
}

//Error 错误结构体
type Error struct {
	Code   string `json:"code"`
	Detail string `json:"detail"`
}

func Init() {
	serviceClient = payS.NewPaymentService("mu.micro.book.service.payment", client.DefaultClient)
	authClient = auth.NewPaymentService("mu.micro.book.service.auth", client.DefaultClient)
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
		orderId: orderId
	})

	
}
