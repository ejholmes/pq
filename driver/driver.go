// Package driver registers the ejholmes/pq driver for use with database/sql.
package driver

import (
	"database/sql"
	"database/sql/driver"

	"github.com/ejholmes/pq"
)

type drv struct{}

func (d *drv) Open(name string) (driver.Conn, error) {
	return pq.Open(name)
}

func init() {
	sql.Register("postgres", &drv{})
}
