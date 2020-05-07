module github.com/xiaobudongzhang/micro-payment-web

go 1.14

replace github.com/xiaobudongzhang/micro-payment-srv => /data/ndemo/micro-payment-srv

require (
	github.com/micro/go-micro/v2 v2.5.0
	github.com/xiaobudongzhang/micro-auth v1.1.1
	github.com/xiaobudongzhang/micro-payment-srv v0.0.0-00010101000000-000000000000
)
