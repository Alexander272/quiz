package models

import "mime/multipart"

type MediaDTO struct {
	Path  string                `form:"path"`
	Image *multipart.FileHeader `form:"image"`
}
