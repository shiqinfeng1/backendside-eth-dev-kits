package common

import (
	"github.com/shiqinfeng1/chunyuyisheng/locale"

	"github.com/labstack/echo"
	"github.com/nicksnyder/go-i18n/i18n"
)

// TLang 返回绑定 accept-language 的i18n方法
func TLang(c echo.Context) i18n.TranslateFunc {
	return locale.Locate(GetAcceptLanguage(c))
}
