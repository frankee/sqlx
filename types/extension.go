package types

import (
	"database/sql"
	"database/sql/driver"
)

//
type Extension interface {
	sql.Scanner
	driver.Valuer
}
