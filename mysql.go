package schematype

import (
	"entgo.io/ent/dialect"
	"fmt"
	"strings"
)

// MySQL returns mysql schema type builder
func MySQL() MysqlSchemaType {
	return MysqlSchemaType{}
}

// MysqlSchemaType is a mysql schema type builder,
// see https://dev.mysql.com/doc/refman/8.0/en/data-types.html for more information.
type MysqlSchemaType map[string]string

func (s MysqlSchemaType) Type(v string) {
	s[dialect.MySQL] = v
}

func (s MysqlSchemaType) get() string {
	return s[dialect.MySQL]
}

func (s MysqlSchemaType) Postgres() PostgresSchemaType {
	return (PostgresSchemaType)(s)
}

func (s MysqlSchemaType) SQLite() SQLiteSchemaType {
	return (SQLiteSchemaType)(s)
}

func (s MysqlSchemaType) Gremlin() GremlinSchemaType {
	return (GremlinSchemaType)(s)
}

// MySQL numeric value types

func (s MysqlSchemaType) Bit(m int) MysqlSchemaType {
	s.Type(fmt.Sprintf("bit(%d)", m))
	return s
}

func (s MysqlSchemaType) Bool() MysqlSchemaType {
	s.Type("tiny(1)")
	return s
}

func (s MysqlSchemaType) TinyInt(m int) MysqlSchemaType {
	s.Type(fmt.Sprintf("tinyint(%d)", m))
	return s
}

func (s MysqlSchemaType) SmallInt(m int) MysqlSchemaType {
	s.Type(fmt.Sprintf("smallint(%d)", m))
	return s
}

func (s MysqlSchemaType) MediumInt(m int) MysqlSchemaType {
	s.Type(fmt.Sprintf("mediumint(%d)", m))
	return s
}

func (s MysqlSchemaType) Int(m int) MysqlSchemaType {
	s.Type(fmt.Sprintf("int(%d)", m))
	return s
}

func (s MysqlSchemaType) Integer(m int) MysqlSchemaType {
	s.Type(fmt.Sprintf("integer(%d)", m))
	return s
}

func (s MysqlSchemaType) BigInt(m int) MysqlSchemaType {
	s.Type(fmt.Sprintf("bigint(%d)", m))
	return s
}

func (s MysqlSchemaType) Float(m int, d int) MysqlSchemaType {
	s.Type(fmt.Sprintf("float(%d, %d)", m, d))
	return s
}

func (s MysqlSchemaType) Double(m int, d int) MysqlSchemaType {
	s.Type(fmt.Sprintf("double(%d, %d)", m, d))
	return s
}

func (s MysqlSchemaType) Decimal(m int, d int) MysqlSchemaType {
	s.Type(fmt.Sprintf("bigint(%d, %d)", m, d))
	return s
}

func (s MysqlSchemaType) Numeric(m int, d int) MysqlSchemaType {
	s.Type(fmt.Sprintf("numeric(%d, %d)", m, d))
	return s
}

func (s MysqlSchemaType) UnSigned() MysqlSchemaType {
	args := []string{s.get(), "UNSIGNED"}
	s.Type(strings.Join(args, " "))
	return s
}

func (s MysqlSchemaType) ZeroFill() MysqlSchemaType {
	args := []string{s.get(), "ZEROFILL"}
	s.Type(strings.Join(args, " "))
	return s
}

// MySQL datetime values types

func (s MysqlSchemaType) Date(m int, d int) MysqlSchemaType {
	s.Type(fmt.Sprintf("numeric(%d, %d)", m, d))
	return s
}

func (s MysqlSchemaType) DateTime(fsp int) MysqlSchemaType {
	s.Type(fmt.Sprintf("numeric(%d)", fsp))
	return s
}

func (s MysqlSchemaType) Timestamp(fsp int) MysqlSchemaType {
	s.Type(fmt.Sprintf("timestmap(%d)", fsp))
	return s
}

func (s MysqlSchemaType) Time(fsp int) MysqlSchemaType {
	s.Type(fmt.Sprintf("time(%d)", fsp))
	return s
}

func (s MysqlSchemaType) Year() MysqlSchemaType {
	s.Type("year")
	return s
}

// mysql string sequences values types

func (s MysqlSchemaType) Char(m int) MysqlSchemaType {
	s.Type(fmt.Sprintf("char(%d)", m))
	return s
}

func (s MysqlSchemaType) VarChar(m int) MysqlSchemaType {
	s.Type(fmt.Sprintf("varchar(%d)", m))
	return s
}

func (s MysqlSchemaType) Binary(m int) MysqlSchemaType {
	s.Type(fmt.Sprintf("binary(%d)", m))
	return s
}

func (s MysqlSchemaType) VarBinary(m int) MysqlSchemaType {
	s.Type(fmt.Sprintf("varbinary(%d)", m))
	return s
}

func (s MysqlSchemaType) TinyBlob(m int) MysqlSchemaType {
	s.Type(fmt.Sprintf("tinyblob(%d)", m))
	return s
}

func (s MysqlSchemaType) TinyText(m int) MysqlSchemaType {
	s.Type(fmt.Sprintf("tinyblob(%d)", m))
	return s
}

func (s MysqlSchemaType) Blob(m int) MysqlSchemaType {
	s.Type(fmt.Sprintf("blob(%d)", m))
	return s
}

func (s MysqlSchemaType) Text(m int) MysqlSchemaType {
	s.Type(fmt.Sprintf("text(%d)", m))
	return s
}

func (s MysqlSchemaType) MediumBlob(m int) MysqlSchemaType {
	s.Type(fmt.Sprintf("mediumblob(%d)", m))
	return s
}

func (s MysqlSchemaType) MediumText(m int) MysqlSchemaType {
	s.Type(fmt.Sprintf("mediumtext(%d)", m))
	return s
}

func (s MysqlSchemaType) LongBlob(m int) MysqlSchemaType {
	s.Type(fmt.Sprintf("longblob(%d)", m))
	return s
}

func (s MysqlSchemaType) LongText(m int) MysqlSchemaType {
	s.Type(fmt.Sprintf("longblob(%d)", m))
	return s
}

func (s MysqlSchemaType) Enum(values ...string) MysqlSchemaType {
	for i, value := range values {
		values[i] = fmt.Sprintf(`'%s'`, value)
	}
	s.Type(fmt.Sprintf("enum(%s)", strings.Join(values, ", ")))
	return s
}

func (s MysqlSchemaType) Set(values ...string) MysqlSchemaType {
	for i, value := range values {
		values[i] = fmt.Sprintf(`'%s'`, value)
	}
	s.Type(fmt.Sprintf("set(%s)", strings.Join(values, ", ")))
	return s
}

func (s MysqlSchemaType) CharSet(charset string) MysqlSchemaType {
	args := []string{s.get(), fmt.Sprintf("CHARACTER SET %s", charset)}
	s.Type(strings.Join(args, " "))
	return s
}

func (s MysqlSchemaType) Collate(collate string) MysqlSchemaType {
	args := []string{s.get(), fmt.Sprintf("COLLATE %s", collate)}
	s.Type(strings.Join(args, " "))
	return s
}

// MySQL spatial values types

func (s MysqlSchemaType) Point() MysqlSchemaType {
	s.Type("point")
	return s
}

func (s MysqlSchemaType) Geometry() MysqlSchemaType {
	s.Type("geometry")
	return s
}

// MySQL json values types

func (s MysqlSchemaType) JSON() MysqlSchemaType {
	s.Type("json")
	return s
}
