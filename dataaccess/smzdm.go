package dataaccess

import (
	//mssql driver
	_ "github.com/denisenkom/go-mssqldb"
	"github.com/moluoX/xg-ob/model"
	"github.com/moluoX/xg-ob/xlog"
)

//SaveArticle save Article
func SaveArticle(m model.SmzdmArticle) error {
	has, err := engine.Exist(&model.SmzdmArticle{Id: m.Id})
	if err != nil {
		return err
	}

	if has {
		_, err := engine.ID(m.Id).Update(m)
		handleErr(err)
	} else {
		_, err = engine.Insert(m)
		handleErr(err)
	}
	return err
}

//ListArticle list article
func ListArticle(limit int, start int, title string) ([]model.SmzdmArticle, int64, error) {
	total, err := engine.Where("Title like ?", "%"+title+"%").Count(new(model.SmzdmArticle))
	handleErr(err)
	list := make([]model.SmzdmArticle, 0)
	err = engine.Where("Title like ?", "%"+title+"%").Desc("Time").Limit(limit, start).Find(&list)
	handleErr(err)
	return list, total, err
}

func handleErr(err error) {
	if err != nil {
		xlog.SugarLogger.Errorf("[dataaccess smzdm error] %v\n", err)
	}
}
