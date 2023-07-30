package handler

import (
	"bytes"
	"encoding/base64"
	"encoding/json"
	"net/http"
	"testoviy-zadaniya/photo-uploader/pkg/uploader"
)

var _ IPhoto = (*PhotowithBase64)(nil)

type PhotowithBase64 struct {
	Handler
}

func newPhotowithBase64() *PhotowithBase64 {
	return &PhotowithBase64{}
}

func (p *PhotowithBase64) Upload(w http.ResponseWriter, r *http.Request) {

	var body map[string]string

	err := json.NewDecoder(r.Body).Decode(&body)
	if err != nil {
		p.handleResponse(w, http.StatusBadRequest, "invalid request body", err)
		return
	}

	name, ok := body["name"]
	if !ok {
		p.handleResponse(w, http.StatusBadRequest, "invalid file name", nil)
		return
	}

	data, ok := body["data"]
	if !ok {
		p.handleResponse(w, http.StatusBadRequest, "invalid file data", nil)
		return
	}

	file, err := base64.StdEncoding.DecodeString(data)
	if err != nil {
		p.handleResponse(w, http.StatusBadRequest, "invalid file data", err)
		return
	}

	var buf bytes.Buffer
	buf.Write(file)
	defer buf.Reset()

	err = uploader.New().Upload(name, &buf)
	if err != nil {
		p.handleResponse(w, http.StatusInternalServerError, "failed to upload file", err)
		return
	}

	err = uploader.New().PreviewImageUpload(name)
	if err != nil {
		p.handleResponse(w, http.StatusInternalServerError, "failed to upload preview image", err)
		return
	}

	p.handleResponse(w, http.StatusOK, "success", nil)
}
