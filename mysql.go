package cycapi

import (
	"database/sql"
	"fmt"
	"github.com/getsentry/sentry-go"
	_ "github.com/go-sql-driver/mysql"
)

type MysqlConn struct {
	Conn *sql.DB
}

func CreateMysqlConnection(host string, username string, password string, dbname string) (DB *MysqlConn, err error) {
	dbConn, err := sql.Open(
		"mysql",
		fmt.Sprintf("%s:%s@tcp(%s)/%s", username, password, host, dbname ),
	)

	if err != nil {
		return DB, err
	}

	DB.Conn = dbConn

	return DB, err
}

func (db *MysqlConn) Query(query string, args interface{}) (rows *sql.Rows, err error) {
	rows, err = db.Conn.Query(query, args)
	if err != nil {
		sentry.CaptureException(err)
	}
	return
}

func (db *MysqlConn) QueryRow(query string, args interface{}) (row *sql.Row) {
	row = db.Conn.QueryRow(query, args)

	return
}

// CloseConnection Closes a DB Connection
func (db *MysqlConn) CloseConnection() error {
	return db.Conn.Close()
}
