package service

import (
	"lampung_trip/model/domain"
	"mime/multipart"
)

type FileServ interface {
	UploadFile(files []*multipart.FileHeader) ([]domain.File, error)
	DeleteFile(fileURL string) error
}
