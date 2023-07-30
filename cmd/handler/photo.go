package handler

import (
	"errors"
	"net/http"
	"strings"
)

var _ IPhoto = (*Photo)(nil)

const (
	MIME_PHTOT_WITH_URL = iota
	MIME_PHOTO_WITH_BASE64
	MIME_PHOTO_WITH_MULTIPART
)

type IPhoto interface {
	Upload(w http.ResponseWriter, r *http.Request)
}

type Photo struct {
	Handler
}

func newPhoto() *Photo {
	return &Photo{}
}

type Builder struct {
	photoType int
}

func (pb *Builder) build(w http.ResponseWriter, r *http.Request) (IPhoto, error) {

	contentType := r.Header.Get("Content-Type")

	switch {
	case strings.HasPrefix(contentType, "application/x-www-form-urlencoded"):

		pb.photoType = MIME_PHTOT_WITH_URL
		return newPhotowithURL(), nil

	case strings.HasPrefix(contentType, "multipart/form-data"):

		pb.photoType = MIME_PHOTO_WITH_MULTIPART
		return newPhotoWithMultiPart(), nil

	case strings.HasPrefix(contentType, "application/json"):

		pb.photoType = MIME_PHOTO_WITH_BASE64
		return newPhotowithBase64(), nil

	default:
		return nil, errors.New("invalid content-type")
	}
}

func (p *Photo) newBuilder() *Builder {
	return &Builder{}
}

func (p *Photo) Upload(w http.ResponseWriter, r *http.Request) {

	pb := p.newBuilder()

	b, err := pb.build(w, r)
	if err != nil {
		p.handleResponse(w, http.StatusBadRequest, err.Error(), err)
		return
	}

	b.Upload(w, r)
}
