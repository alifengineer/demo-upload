package handler

import (
	"net/http"
	"path/filepath"
	"testoviy-zadaniya/photo-uploader/pkg/uploader"
)

var _ IPhoto = (*PhotoWithMultiPart)(nil)

type PhotoWithMultiPart struct {
	Handler
}

func newPhotoWithMultiPart() *PhotoWithMultiPart {
	return &PhotoWithMultiPart{}
}

func (p *PhotoWithMultiPart) Upload(w http.ResponseWriter, r *http.Request) {

	file, fileHeader, err := r.FormFile("data")
	if err != nil {
		p.handleResponse(w, http.StatusBadRequest, "invalid file", err)
		return
	}
	defer file.Close()

	fileName := r.FormValue("name")
	if fileName == "" {
		p.handleResponse(w, http.StatusBadRequest, "invalid file name", nil)
		return
	}

	fileSize := fileHeader.Size
	if fileSize > 10<<20 {
		p.handleResponse(w, http.StatusBadRequest, "file size too large", nil)
		return
	}

	fileExt := filepath.Ext(fileName)
	if fileExt != ".jpg" {
		p.handleResponse(w, http.StatusBadRequest, "invalid file extension", nil)
		return
	}

	err = uploader.New().Upload(fileName, file)
	if err != nil {
		p.handleResponse(w, http.StatusInternalServerError, "failed to upload file", err)
		return
	}

	err = uploader.New().PreviewImageUpload(fileName)
	if err != nil {
		p.handleResponse(w, http.StatusInternalServerError, "failed to upload preview image", err)
		return
	}

	p.handleResponse(w, http.StatusOK, "file uploaded", nil)
}
