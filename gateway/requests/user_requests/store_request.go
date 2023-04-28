package user_requests

import (
	"fmt"
	"github.com/99designs/gqlgen/graphql"
	"io"
	"log"
	"mime/multipart"
	"net/textproto"
	"strconv"
)

type StoreRequest struct {
	firstName string
	lastName  string
	age       *int
	address   string
	email     string
	groupID   int
	avatar    *graphql.Upload
}

func NewStoreRequest(firstName string, lastName string, age *int, address string, email string, groupID int, avatar *graphql.Upload) *StoreRequest {
	return &StoreRequest{
		firstName: firstName,
		lastName:  lastName,
		age:       age,
		address:   address,
		email:     email,
		groupID:   groupID,
		avatar:    avatar,
	}
}

func (r *StoreRequest) MakeEndpoint() string {
	return "http://app1:8080/api/users"
}

func (r *StoreRequest) MakeRequest() (body io.Reader, contentType string, err error) {
	// multipartで送りたい場合
	// multipartデフォルト実装（メモリ消費が高い）
	// bodyBuf := &bytes.Buffer{}
	// Pipeを使ったメモリ消費に考慮した実装（Pipeは内部でバッファを持っているので、バッファを使いまわすことができる）
	body, pw := io.Pipe()
	bodyWriter := multipart.NewWriter(pw)

	// multipartデフォルト実装（メモリ消費が高い）
	// fileWriter, err := bodyWriter.CreateFormFile("avatar", avatar.Filename)
	// if err != nil {
	// 	log.Printf("Error creating form file: %v", err)
	// 	return nil, err
	// }
	// _, err = io.Copy(fileWriter, avatar.File)
	// if err != nil {
	// 	log.Printf("Error copying file: %v", err)
	// 	return nil, err
	// }

	// Pipeへの書き込みは別スレッドで行う
	go func() {
		if r.avatar != nil {
			// createformfileの代わりに、CreatePartを使うことで、ヘッダーを設定できる
			h := make(textproto.MIMEHeader)
			h.Set("Content-Disposition", fmt.Sprintf(`form-data; name="%s"; filename="%s"`, "avatar", r.avatar.Filename))
			h.Set("Content-Type", r.avatar.ContentType)
			part, err := bodyWriter.CreatePart(h)
			if err != nil {
				log.Printf("Error creating part: %v", err)
			}
			cnt, err := io.Copy(part, r.avatar.File)
			log.Printf("copy %d bytes from file %s in total\n", cnt, r.avatar.Filename)
			if err != nil {
				log.Printf("Error copying file: %v", err)
			}
		}

		firstNameField, _ := bodyWriter.CreateFormField("first_name")
		firstNameField.Write([]byte(r.firstName))
		lastNameField, _ := bodyWriter.CreateFormField("last_name")
		lastNameField.Write([]byte(r.lastName))
		ageField, _ := bodyWriter.CreateFormField("age")
		ageField.Write([]byte(strconv.Itoa(*r.age)))
		addressField, _ := bodyWriter.CreateFormField("address")
		addressField.Write([]byte(r.address))
		emailField, _ := bodyWriter.CreateFormField("email")
		emailField.Write([]byte(r.email))
		groupIDField, _ := bodyWriter.CreateFormField("group_id")
		groupIDField.Write([]byte(strconv.Itoa(r.groupID)))

		bodyWriter.Close()
		pw.Close()
	}()

	// multipartデフォルト実装
	// firstNameField, _ := bodyWriter.CreateFormField("first_name")
	// firstNameField.Write([]byte(firstName))
	// lastNameField, _ := bodyWriter.CreateFormField("last_name")
	// lastNameField.Write([]byte(lastName))
	// ageField, _ := bodyWriter.CreateFormField("age")
	// ageField.Write([]byte(strconv.Itoa(*age)))
	// addressField, _ := bodyWriter.CreateFormField("address")
	// addressField.Write([]byte(address))
	// emailField, _ := bodyWriter.CreateFormField("email")
	// emailField.Write([]byte(email))
	// groupIDField, _ := bodyWriter.CreateFormField("group_id")
	// groupIDField.Write([]byte(strconv.Itoa(groupID)))

	contentType = bodyWriter.FormDataContentType()
	// multipartデフォルト実装
	// bodyWriter.Close()

	return body, contentType, nil
}
