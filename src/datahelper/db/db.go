package db

import (
	"database/sql"
	"errors"
	"utils/config"

	"github.com/go-redis/redis"
	_ "github.com/go-sql-driver/mysql"
	"github.com/hongxufeng/fileLogger"
)

var MysqlMain *sql.DB
var RedisCache *redis.Client
var DBLog *fileLogger.FileLogger

func Init(config *config.Config) (err error) {
	DBLog = fileLogger.NewDefaultLogger(config.LogDir, "DB_LOG.log")
	DBLog.SetPrefix("[DB] ")
	MysqlMain, err = sql.Open("mysql", config.Mysql)
	if err != nil {
		return // Just for example purpose. You should use proper error handling instead of panic
	}
	err = MysqlMain.Ping()
	if err != nil {
		return // proper error handling instead of panic in your app
	}

	opt, err := redis.ParseURL(config.Redis)
	if err != nil {
		return
	}
	// Create client as usually.
	RedisCache = redis.NewClient(opt)

	_, err = RedisCache.Ping().Result()
	if err != nil {
		return
	}
	return
}

func Query(query string) (*sql.Rows, error) {
	//还是不要设置redis
	return MysqlMain.Query(query)
}
func QueryRow(query string) *sql.Row {
	//还是不要设置redis
	return MysqlMain.QueryRow(query)
}
func FetchRows(sqlstr string, args ...interface{}) ([]map[string]string, error) {
	stmtOut, err := MysqlMain.Prepare(sqlstr)
	if err != nil {
		return nil, err
	}
	defer stmtOut.Close()

	rows, err := stmtOut.Query(args...)
	if err != nil {
		return nil, err
	}

	columns, err := rows.Columns()
	if err != nil {
		return nil, err
	}

	values := make([]sql.RawBytes, len(columns))
	scanArgs := make([]interface{}, len(values))

	ret := make([]map[string]string, 0)
	for i := range values {
		scanArgs[i] = &values[i]
	}

	for rows.Next() {
		err = rows.Scan(scanArgs...)
		if err != nil {
			return nil, err
		}
		var value string
		vmap := make(map[string]string, len(scanArgs))
		for i, col := range values {
			if col == nil {
				value = "NULL"
			} else {
				value = string(col)
			}
			vmap[columns[i]] = value
		}
		ret = append(ret, vmap)
	}
	return ret, nil
}

//修改和删除
func Exec(sqlstr string, args ...interface{}) (err error) {
	stmtIns, err := MysqlMain.Prepare(sqlstr)
	if err != nil {
		return err
	}
	defer stmtIns.Close()

	result, err := stmtIns.Exec(args...)
	if err != nil {
		return err
	}
	_, err = result.RowsAffected()
	if err != nil {
		return err
	}
	defer func() {

		if r := recover(); r != nil {
			//check exactly what the panic was and create error.
			switch x := r.(type) {
			case string:
				err = errors.New(x)
			case error:
				err = x
			default:
				err = errors.New("数据库错误！")
			}
		}
	}()
	return nil
}

//增加
func Insert(sqlstr string, args ...interface{}) (lastInsertId int64, err error) {
	stmtIns, err := MysqlMain.Prepare(sqlstr)
	if err != nil {
		return 0, err
	}
	defer stmtIns.Close()

	result, err := stmtIns.Exec(args...)
	if err != nil {
		return 0, err
	}
	lastInsertId, err = result.LastInsertId()
	if err != nil {
		return 0, err
	}
	defer func() {

		if r := recover(); r != nil {
			//check exactly what the panic was and create error.
			switch x := r.(type) {
			case string:
				err = errors.New(x)
			case error:
				err = x
			default:
				err = errors.New("数据库错误！")
			}
		}
	}()
	return lastInsertId, nil
}
