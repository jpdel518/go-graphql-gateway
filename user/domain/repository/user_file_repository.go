package repository

import "mime/multipart"

type UserFileRepository interface {
	UploadFile(file multipart.File, fileHeader *multipart.FileHeader) (string, error)
}
