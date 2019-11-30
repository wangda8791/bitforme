package models

import (
	"log"
	_ "net/url"

	"github.com/jinzhu/gorm"
	"github.com/qor/transition"
	// "database/sql"
	// "log"
	"github.com/gchaincl/dotsql"
	// "strings"
	_ "github.com/go-sql-driver/mysql"
)

var (
	url string
	// db  *sql.DB
	db  *gorm.DB
	dot *dotsql.DotSql
)

func Init(DB_URL string) {

	log.Printf("DB_URL: %v\n", DB_URL)
	d, err := gorm.Open("mysql", DB_URL)
	if err != nil {
		log.Fatal(err)
	}
	db = d

	err = db.DB().Ping()
	if err != nil {
		panic(err)
	}

	db.AutoMigrate(&transition.StateChangeLog{})
}

const (
	userDBSchema = "./database/schema.sql"

	updateQuery           = "update"
	updateResetTokenQuery = "update-reset-token"
	updatePasswordQuery   = "update-password"
	updateOrderQuery      = "update-order"
	updateAccountQuery    = "update-account"

	insertQuery               = "insert"
	insertAuthenticationQuery = "insert-authentication"
	insertMemberQuery         = "insert-member"
	selectMemberIdQuery       = "select-member-id"
	insertId_DocumentQuery    = "insert-Id_Document"
	insertAccountQuery        = "insert-Account"
	insertOrderQuery          = "insert-order"

	selectLoginQuery      = "select-login"
	selectEmailQuery      = "select-email"
	selectResetTokenQuery = "select-reset-token"
	selectMemberQuery     = "select-member"
	selectAccountQuery    = "select-account"

	LOCKING_BUFFER_FACTOR = "1.1"
)
