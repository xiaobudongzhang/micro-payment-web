package handler

import (
	"context"
	"net/http"

	auth "github.com/xiaobudongzhang/micro-auth/proto/auth"

	"github.com/micro/go-micro/v2/util/log"
	"github.com/xiaobudongzhang/micro-plugins/session"

	"github.com/xiaobudongzhang/micro-basic/common"
)

func AuthWrapper(h http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		ck, _ := r.Cookie(common.RememberMeCookieName)

		if ck == nil {
			log.Logf("token 不存在")
			http.Error(w, "非法请求", 400)
			return
		}

		sess := session.GetSession(w, r)

		if sess.ID != "" {
			if sess.Values["valid"] != nil {
				h.ServeHTTP(w, r)
				return
			} else {
				userId := sess.Values["userId"].(int64)
				if userId != 0 {
					rsp, err := authClient.GetCachedAccessToken(context.TODO(), &auth.Request{
						UserId: userId,
					})

					if err != nil {
						log.Logf("[authwrapper] err:%s", err)
						http.Error(w, "非法请求", 400)
						return
					}

					if rsp.Token != ck.Value {
						log.Logf("[authwrapper] token 不一致")
						http.Error(w, "非法请求", 400)
						return
					}
				} else {
					log.Logf("[authwrapper] session不合法,无用户id")
					http.Error(w, "非法请求", 400)
					return
				}
			}
		} else {
			http.Error(w, "非法请求", 400)
			return
		}

		h.ServeHTTP(w, r)
	})
}
