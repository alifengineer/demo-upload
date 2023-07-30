package handler

import (
	"net/http"
	"testoviy-zadaniya/photo-uploader/pkg/uploader"
	"time"
)

var _ IPhoto = (*PhotowithURL)(nil)

type PhotowithURL struct {
	Handler
}

func newPhotowithURL() *PhotowithURL {
	return &PhotowithURL{}
}

func (p *PhotowithURL) Upload(w http.ResponseWriter, r *http.Request) {

	name := r.FormValue("name")
	if name == "" {
		p.handleResponse(w, http.StatusBadRequest, "invalid file name", nil)
		return
	}

	url := r.FormValue("url")
	if url == "" {
		p.handleResponse(w, http.StatusBadRequest, "invalid file url", nil)
		return
	}

	client := http.Client{
		Timeout: 10 * time.Second,
	}

	creq, err := http.NewRequest("GET", url, nil)
	if err != nil {
		p.handleResponse(w, http.StatusInternalServerError, "failed to create request", err)
		return
	}

	cresp, err := client.Do(creq)
	if err != nil {
		p.handleResponse(w, http.StatusInternalServerError, "failed to do request", err)
		return
	}
	defer cresp.Body.Close()

	if cresp.StatusCode != http.StatusOK {
		p.handleResponse(w, http.StatusInternalServerError, "failed to get file", err)
		return
	}

	err = uploader.New().Upload(name, cresp.Body)
	if err != nil {
		p.handleResponse(w, http.StatusInternalServerError, "failed to upload file", err)
		return
	}

	err = uploader.New().PreviewImageUpload(name)
	if err != nil {
		p.handleResponse(w, http.StatusInternalServerError, "failed to upload preview image", err)
		return
	}

	p.handleResponse(w, http.StatusOK, "file uploaded", nil)

}
