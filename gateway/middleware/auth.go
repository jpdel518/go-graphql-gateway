package middleware

import (
	"bytes"
	"context"
	"encoding/json"
	"github.com/google/uuid"
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
		var m map[string]interface{}
		json.Unmarshal(body, &m)
		log.Println(m)
		log.Println(len(m))
		r.Body = io.NopCloser(buf)

		// cookieの設定
		cookie, err := r.Cookie("session_id")
		sessionId := ""
		if cookie == nil || err != nil {
			sessionId = uuid.New().String()
			cookie := &http.Cookie{
				Name:  "session_id",
				Value: sessionId,
				// Secure:   true,
				HttpOnly: true,
			}
			http.SetCookie(w, cookie)
		} else {
			sessionId = cookie.Value
		}
		log.Printf("auth.go session_id: %v", sessionId)
		// session_idをcontextに設定
		ctx := context.WithValue(r.Context(), "session_id", sessionId)
		ctx = context.WithValue(ctx, "is_subscribe", len(m) <= 0)
		r = r.WithContext(ctx)
		log.Printf("auth.go context session_id: %v", r.Context().Value("session_id"))
		log.Printf("auth.go context is_subscribe: %v", r.Context().Value("is_subscribe"))

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
