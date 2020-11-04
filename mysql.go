package cycapi

import (
	"database/sql"
	"fmt"
	"github.com/getsentry/sentry-go"
	_ "github.com/go-sql-driver/mysql"
)

type MysqlConn struct {
	DBConnection *sql.DB
}

func CreateMysqlConnection(host string, username string, password string, dbname string) (Conn *MysqlConn, err error) {
	dbConn, err := sql.Open(
		"mysql",
		fmt.Sprintf("%s:%s@tcp(%s)/%s", username, password, host, dbname ),
	)

	Conn.DBConnection = dbConn

	return Conn, err
}

func (conn *mysqlConn) Query(query string, args interface{}) (rows *sql.Rows, err error) {
	rows, err = conn.DBConnection.Query(query, args)
	if err != nil {
		sentry.CaptureException(err)
	}
	return
}

func (conn *mysqlConn) QueryRow(query string, args interface{}) (row *sql.Row) {
	row = conn.DBConnection.QueryRow(query, args)

	return
}

// CloseConnection Closes a DB Connection
func (conn *mysqlConn) CloseConnection() error {
	return conn.DBConnection.Close()
}
