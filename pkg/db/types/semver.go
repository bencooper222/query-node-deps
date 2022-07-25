package types

import (
	"database/sql"
	"database/sql/driver"
)

// uses https://github.com/theory/pg-semver/ on the SQL side

type Semver string

func (semver *Semver) Scan(value interface{}) (err error) {
	nullStr := &sql.NullString{}
	err = nullStr.Scan(value)
	*semver = Semver(nullStr.String)

	return
}

func (semver Semver) Value() (driver.Value, error) {
	return semver, nil
}

func (Semver) GormDataType() string {
	return "semver"
}
