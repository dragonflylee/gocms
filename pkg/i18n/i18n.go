package i18n

import (
	"io/fs"
	"net/http"
	"path/filepath"
	"strings"

	"gocms/pkg/errors"

	"github.com/gin-gonic/gin"
	"github.com/unknwon/i18n"
	"golang.org/x/text/language"
)

// Load 加载语言文件
func Load(fsys fs.FS, pattern string) error {
	list, err := fs.Glob(fsys, pattern)
	if err != nil {
		return err
	}
	for _, name := range list {
		buf, err := fs.ReadFile(fsys, name)
		if err != nil {
			return err
		}
		base, ext := filepath.Base(name), filepath.Ext(name)
		i18n.SetMessage(strings.TrimSuffix(base, ext), buf)
	}

	i18n.SetDefaultLang("en-US")
	return nil
}

func Handle(c *gin.Context) {

	local := &i18n.Locale{
		Lang: c.Query("lang"),
	}
	c.Set("i18n", local)

	tags, _, _ := language.ParseAcceptLanguage(c.GetHeader("Accept-Language"))
	for _, tag := range tags {
		if i18n.IsExist(local.Lang) {
			break
		}
		local.Lang = tag.String()
	}

	c.Next()

	if err := c.Errors.Last(); err != nil {
		switch err.Type {
		case errors.ErrorTypeMessage:
		case gin.ErrorTypeBind:
			err = errors.ErrBinding
		default:
			err = errors.ErrInternal
		}
		c.AbortWithStatusJSON(http.StatusOK, gin.H{
			"error":   err.Error(),
			"message": i18n.Tr(local.Lang, err.Error(), err.Meta),
		})
	}
}
