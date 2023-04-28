package file

import (
	"github.com/jpdel518/go-graphql-gateway/user/domain/repository"
	"io"
	"log"
	"mime/multipart"
	"os"
)

type userFileRepository struct {
}

func NewUserFileRepository() repository.UserFileRepository {
	return &userFileRepository{}
}

func (u userFileRepository) UploadFile(file multipart.File, fileHeader *multipart.FileHeader) (string, error) {
	err := os.MkdirAll("./uploads", os.ModePerm)
	if err != nil {
		log.Printf("Error creating uploads directory: %v", err)
		return "", err
	}

	// 一度uploadsディレクトリに手動で画像を保存するもしくは、アップロードするものと同名のファイルを用意しておかないと失敗する
	// おそらく、Dockerのマウントの関係で、uploadsディレクトリにファイルが存在しないと、ディレクトリを認識できないのではないかと思われる
	dst, err := os.Create("./uploads/" + fileHeader.Filename)
	if err != nil {
		log.Printf("Error creating file: %v", err)
		return "", err
	}
	defer dst.Close()

	_, err = io.Copy(dst, file)
	if err != nil {
		log.Printf("Error copying file: %v", err)
		return "", err
	}

	log.Printf("File uploaded successfully: %s", fileHeader.Filename)

	return dst.Name(), nil
}
