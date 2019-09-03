package handler

import (
	"crypto/md5"
	"encoding/hex"
	"image"
	"io"
	"net/http"
	"strings"

	"github.com/makiuchi-d/gozxing"
	"github.com/makiuchi-d/gozxing/qrcode"
)

// Upload 上传文件
func Upload(w http.ResponseWriter, r *http.Request) {
	f, _, err := r.FormFile("file")
	if err != nil {
		jFailed(w, http.StatusBadRequest, err.Error())
		return
	}
	defer f.Close()

	t := strings.TrimSpace(r.PostForm.Get("type"))
	switch t {
	case "1":
		u := md5.New()
		io.Copy(u, f)
		s := hex.EncodeToString(u.Sum(nil))
		jFailed(w, http.StatusCreated, s)

	case "2":
		img, _, err := image.Decode(f)
		if err != nil {
			jFailed(w, http.StatusBadRequest, err.Error())
			return
		}
		bmp, err := gozxing.NewBinaryBitmapFromImage(img)
		if err != nil {
			jFailed(w, http.StatusBadRequest, err.Error())
			return
		}
		s, err := qrcode.NewQRCodeReader().Decode(bmp, nil)
		if err != nil {
			jFailed(w, http.StatusBadRequest, err.Error())
			return
		}
		jFailed(w, http.StatusCreated, s.GetText())

	default:
		jFailed(w, http.StatusBadRequest, "unsupport %s", t)
	}
}
