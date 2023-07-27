package middleware

import (
	"bytes"
	"github.com/jpdel518/go-graphql-gateway/gateway/controllers"
	"io"
	"log"
	"net/http"
)

func Auth(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		log.Println(r.URL)
		log.Println(r.RequestURI)
		log.Println(r.RemoteAddr)
		log.Println(r.Host)
		log.Println(r.Cookies())
		// bodyをeofまで読み込まずに中身を参照する
		buf := new(bytes.Buffer)
		io.Copy(buf, r.Body)
		body := buf.Bytes()
		log.Println(string(body))
		r.Body = io.NopCloser(buf)

		token := r.Header.Get("Authorization")
		if token == "" {
			log.Println("No Authorization Header")
			next.ServeHTTP(w, r)
			return
		} else {
			// tokenを設定していたら別サービスへ連携する（IDが1のユーザー情報をuserサービスへ取りにいく）
			id := "1"
			get, err := controllers.NewUserController().Get(r.Context(), &id)
			if err != nil {
				log.Printf("Auth get user error: %v", err)
				return
			} else {
				log.Printf("Auth get user: %v", get)
			}

			next.ServeHTTP(w, r)
			return
		}
	})
}
