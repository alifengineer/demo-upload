package uploader

import (
	"image/jpeg"
	"io"
	"os"
	"path/filepath"

	"github.com/labstack/gommon/log"
	"github.com/nfnt/resize"
	"github.com/oliamb/cutter"
)

const (
	IMGs_PATH    = "static/imgs"
	PREVIEW_PATH = "static/preview"
	PREVIEW_SIZE = 200
)

var _ Uploader = (*uploader)(nil)

type Uploader interface {
	Upload(name string, file io.Reader) error
	PreviewImageUpload(name string) error
}

type uploader struct {
	originalImage string
	previewImage  string
}

func New() *uploader {
	return &uploader{
		originalImage: IMGs_PATH,
		previewImage:  PREVIEW_PATH,
	}
}

func (u *uploader) Upload(name string, file io.Reader) error {

	fullPath := filepath.Join(u.originalImage, name)
	f, err := os.OpenFile(fullPath, os.O_WRONLY|os.O_CREATE, 0666)
	if err != nil {
		return err
	}
	defer f.Close()

	_, err = io.Copy(f, file)

	return err
}

func (u *uploader) PreviewImageUpload(name string) error {

	file, err := os.Open(filepath.Join(u.originalImage, name))
	if err != nil {
		log.Errorf("failed to open image: %s", err.Error())
		return err
	}

	orgImg, err := jpeg.Decode(file)
	if err != nil {
		log.Errorf("failed to decode image: %s", err.Error())
		return err
	}

	width := orgImg.Bounds().Dx()
	height := orgImg.Bounds().Dy()

	if width > height {
		width = height
	} else {
		height = width
	}

	croppedImg, err := cutter.Crop(orgImg, cutter.Config{
		Width:  width,
		Height: height,
		Mode:   cutter.Centered,
	})
	if err != nil {
		log.Errorf("failed to crop image: %s", err.Error())
		return err
	}

	resizedImg := resize.Resize(PREVIEW_SIZE, PREVIEW_SIZE, croppedImg, resize.Lanczos3)

	fullPath := filepath.Join(u.previewImage, name)

	out, err := os.Create(fullPath)
	if err != nil {
		log.Errorf("failed to create preview image: %s", err.Error())
		return err
	}
	defer out.Close()

	err = jpeg.Encode(out, resizedImg, nil)
	if err != nil {
		log.Errorf("failed to encode preview image: %s", err.Error())
		return err
	}

	return nil
}
