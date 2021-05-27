package handler

import (
	"crypto/md5"
	"encoding/hex"
	"encoding/xml"
	"image/gif"
	"io"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/o1egl/govatar"
)

// Upload 上传文件
func Upload(c *gin.Context) {
	fh, err := c.FormFile("file")
	if err != nil {
		c.Error(err)
		return
	}
	f, err := fh.Open()
	if err != nil {
		c.Error(err)
		return
	}
	defer f.Close()

	u := md5.New()
	io.Copy(u, f)
	s := hex.EncodeToString(u.Sum(nil))
	c.Set("Data", s)
}

// Avatar 头像生成
func Avatar(c *gin.Context) {
	name := c.DefaultQuery("user", "guest")
	img, err := govatar.GenerateForUsername(govatar.MALE, name)
	if err != nil {
		return
	}
	c.Header("Content-Type", "image/gif")

	gif.Encode(c.Writer, img, nil)
}

// BingPic 每日背景
func BingPic(c *gin.Context) {
	resp, err := http.Get("https://cn.bing.com/HPImageArchive.aspx?idx=0&n=1")
	if err != nil {
		c.AbortWithError(http.StatusBadGateway, err)
		return
	}
	defer resp.Body.Close()

	var s struct {
		XMLName xml.Name `xml:"images"`
		Image   []struct {
			URL string `xml:"url"`
		} `xml:"image"`
	}
	if err = xml.NewDecoder(resp.Body).Decode(&s); err != nil {
		c.AbortWithError(http.StatusBadGateway, err)
		return
	}
	for _, v := range s.Image {
		location := "https://cn.bing.com" + v.URL
		c.Redirect(http.StatusFound, location)
		return
	}
}
