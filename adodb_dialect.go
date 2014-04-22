package gorp

import (
	"fmt"
	"reflect"
	"strings"
)

///////////////////////////////////////////////////////
// MSSQL -- go-adodb //
///////////

// Implementation of Dialect for MySQL databases.
type AdodbDialect struct {
	// Encoding is the character encoding to use for created tables
	Encoding string
}

func (m AdodbDialect) ToSqlType(val reflect.Type, maxsize int, isAutoIncr bool) string {
	switch val.Kind() {
	case reflect.Ptr:
		return m.ToSqlType(val.Elem(), maxsize, isAutoIncr)
	case reflect.Bool:
		return "bit"
	case reflect.Int8:
		return "tinyint"
	case reflect.Uint8: //	As MSSQL does not support unsigned int/smallint/tinyint
		return "smallint"
	case reflect.Int16:
		return "smallint"
	case reflect.Uint16: //	As MSSQL does not support unsigned int/smallint/tinyint
		return "int"
	case reflect.Int, reflect.Int32:
		return "int"
	case reflect.Uint, reflect.Uint32: //	As MSSQL does not support unsigned int/smallint/tinyint
		return "bigint"
	case reflect.Int64:
		return "bigint"
	case reflect.Uint64:
		return "bigint"
	case reflect.Float64, reflect.Float32:
		return "float"
	case reflect.Slice:
		if val.Elem().Kind() == reflect.Uint8 {
			return "varbinary"
		}
	}

	switch val.Name() {
	case "NullInt64":
		return "bigint"
	case "NullFloat64":
		return "float"
	case "NullBool":
		return "bit"
	case "Time":
		return "datetime"
	}

	if maxsize < 1 {
		maxsize = 255
	}
	return fmt.Sprintf("varchar(%d)", maxsize)
}

// Returns auto_increment
func (m AdodbDialect) AutoIncrStr() string {
	return "identity(1,1)"
}

func (m AdodbDialect) AutoIncrBindValue() string {
	return "null"
}

func (m AdodbDialect) AutoIncrInsertSuffix(col *ColumnMap) string {
	return ""
}

// Returns engine=%s charset=%s  based on values stored on struct
func (m AdodbDialect) CreateTableSuffix() string {
	return ""
}

func (m AdodbDialect) TruncateClause() string {
	return "truncate"
}

// Returns "?"
func (m AdodbDialect) BindVar(i int) string {
	return "?"
}

func (m AdodbDialect) InsertAutoIncr(exec SqlExecutor, insertSql string, params ...interface{}) (int64, error) {
	res, err := exec.Exec(insertSql, params...)
	if err != nil {
		return 0, err
	}
	return res.LastInsertId()
}

func (d AdodbDialect) QuoteField(f string) string {
	return "[" + f + "]"
}

// MySQL does not have schemas like PostgreSQL does, so just escape it like normal
func (d AdodbDialect) QuotedTableForQuery(schema string, table string) string {
	if strings.TrimSpace(schema) == "" {
		return d.QuoteField(table)
	}

	return schema + "." + d.QuoteField(table)
}
