package web

import (
	"log"
	"net/http"
	"strconv"

	"github.com/labstack/echo/v4"

	"github.com/moluoX/xg-ob/dataaccess"
	"github.com/moluoX/xg-ob/model"
)

func smzdm(c echo.Context) error {
	return c.Render(http.StatusOK, "smzdm", "World")
}

func listSmzdm(c echo.Context) error {
	limit, err := strconv.Atoi(c.FormValue("limit"))
	if err != nil {
		limit = 20
	}
	offset, err := strconv.Atoi(c.FormValue("offset"))
	if err != nil {
		offset = 0
	}
	search := c.FormValue("search")
	list, total, err := dataaccess.ListArticle(limit, offset, search)
	if err != nil {
		return err
	}
	return c.JSON(http.StatusOK, &model.SmzdmArticlePaged{Rows: list, Total: total})
}

func handleErr(err error) {
	if err != nil {
		log.Printf("[web smzdm error] %v\n", err)
	}
}
