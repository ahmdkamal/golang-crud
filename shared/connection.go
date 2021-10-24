package shared

import (
	"fmt"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

var (
	DBConnection *gorm.DB
	err error
)

func Connect() {
	var dns = fmt.Sprintf("%s:%s@tcp(%s)/%s?charset=utf8&parseTime=True&loc=Local",
		GetEnvVariable("DB_USER"), GetEnvVariable("DB_PASSWORD"),
		GetEnvVariable("DB_HOST"), GetEnvVariable("DB_DATABASE"))
	DBConnection, err = gorm.Open(mysql.New(mysql.Config{
		DSN: dns, // data source name
		DefaultStringSize: 256, // default size for string fields
		DisableDatetimePrecision: true, // disable datetime precision, which not supported before MySQL 5.6
		DontSupportRenameIndex: true, // drop & create when rename index, rename index not supported before MySQL 5.7, MariaDB
		DontSupportRenameColumn: false, // `change` when rename column, rename column not supported before MySQL 8, MariaDB
		SkipInitializeWithVersion: false, // auto configure based on currently MySQL version
	}), &gorm.Config{})

	if err != nil {
		panic(err)
	}
}
