// +build clickhouse

package cli

import (
	_ "github.com/dunsal/migrate/v4/database/clickhouse"
	_ "github.com/kshvakov/clickhouse"
)
