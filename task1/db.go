package task1

import (
	"database/sql"
	"fmt"
	_ "github.com/lib/pq"
)

var (
	db    *sql.DB
	errDB error
)

func OpenSqlConnections() {
	db, errDB = sql.Open("postgres", fmt.Sprintf("postgres://%v:%v@%v/%v", UserEnv, PasswordEnv, HostEnv, DBEnv))
	if errDB == nil {
		errDB = db.Ping()
	}
	if errDB != nil {
		panic(errDB)
	}
	if MaxPoolSizeEnv > 0 {
		db.SetMaxOpenConns(int(MaxPoolSizeEnv))
	}
}
