package gpa

import (
	"database/sql"
	"utils"
	"strings"
	"github.com/cihub/seelog"
)

func Init(Driver, Dsn string, models ...interface{}) *Gpa {
	impl := &Gpa{driver: Driver, dsn: Dsn}
	var err error
	impl.Conn, err = sql.Open(impl.driver, impl.dsn)
	if err != nil {
		panic("数据库连接错误:driver=" + impl.driver + ";" + impl.dsn)
	} else {
		impl.Conn.SetMaxOpenConns(5)
		//	dao.Conn.SetMaxIdleConns(0)
		impl.Conn.Ping()
	}
	for _, d := range models {
		impl.setMethodImpl(d)
	}
	return impl
}

func InitGpa(db string, models ...interface{}) *Gpa {
	if utils.EnvIsDev {
		ip := utils.GetIp()
		utils.EnvParamSet("ImgHost", "http://"+ip)
	}
	if strings.Index(db, ":") < 0 {
		utils.EnvParamSet("DbDsn", "root:root@tcp(127.0.0.1:3306)/" + db+
			"?timeout=30s&charset=utf8mb4&parseTime=true")
	} else {
		utils.EnvParamSet("DbDsn", db)
	}
	dsn := utils.EnvParam("DbDsn")
	dv := utils.EnvParam("DbDriver")
	if dv == "" {
		dv = "mysql"
	}
	dbx := Init(dv, dsn, models...)
	seelog.Info(dsn)
	return dbx
}
