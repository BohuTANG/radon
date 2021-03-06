/*
 * Radon
 *
 * Copyright 2020 The Radon Authors.
 * Code is licensed under the GPLv3.
 *
 */

package proxy

import (
	"github.com/xelabs/go-mysqlstack/driver"
	"github.com/xelabs/go-mysqlstack/sqlparser"
	"github.com/xelabs/go-mysqlstack/sqlparser/depends/sqltypes"
)

// handleHelp used to handle the HELP command.
func (spanner *Spanner) handleHelp(session *driver.Session, query string, node sqlparser.Statement) (*sqltypes.Result, error) {
	return spanner.ExecuteSingle(query)
}
