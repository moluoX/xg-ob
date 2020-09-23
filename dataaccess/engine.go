package dataaccess

import (
	//mssql driver
	_ "github.com/denisenkom/go-mssqldb"
	"xorm.io/xorm"
	"xorm.io/xorm/names"

	"github.com/moluoX/xg-ob/xlog"
)

var engine *xorm.Engine

func init() {
	var err error
	engine, err = xorm.NewEngine("mssql", "server=127.0.0.1\\X;user id=sa;password=198633aa;database=OB")
	if err != nil {
		xlog.SugarLogger.Errorf("[create xorm engine] %v\n", err)
	}

	if err = engine.Ping(); err != nil {
		xlog.SugarLogger.Errorf("[ping xorm engine] %v\n", err)
	}

	engine.SetTableMapper(names.SameMapper{})
	engine.SetColumnMapper(names.SameMapper{})

	//engine.ShowSQL(true)
}
