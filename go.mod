module github.com/xiaobudongzhang/micro-payment-web

go 1.14

replace github.com/xiaobudongzhang/micro-basic => /wwwroot/microdemo/micro-basic

replace github.com/xiaobudongzhang/micro-inventory-srv => /wwwroot/microdemo/micro-inventory-srv

replace github.com/xiaobudongzhang/micro-payment-srv => /wwwroot/microdemo/micro-payment-srv

replace github.com/xiaobudongzhang/micro-order-srv => /wwwroot/microdemo/micro-order-srv

replace github.com/xiaobudongzhang/micro-plugins => /wwwroot/microdemo/micro-plugins


replace github.com/xiaobudongzhang/micro-auth => /wwwroot/microdemo/micro-auth


require (
	github.com/gorilla/sessions v1.2.0 // indirect
	github.com/micro/go-micro/v2 v2.6.0
	github.com/xiaobudongzhang/micro-auth v1.1.1
	github.com/xiaobudongzhang/micro-basic v1.1.5
	github.com/xiaobudongzhang/micro-payment-srv v0.0.0-00010101000000-000000000000
	github.com/xiaobudongzhang/micro-plugins v0.0.0-20200423150326-f4d282de91ed
)
