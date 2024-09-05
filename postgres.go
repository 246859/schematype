package schematype

import (
	"entgo.io/ent/dialect"
	"fmt"
)

// Postgres returns Postgresql schema type builder
func Postgres() PostgresSchemaType {
	return PostgresSchemaType{}
}

// PostgresSchemaType is schema type builder for Postgresql
// see https://www.postgresql.org/docs/16/datatype.html for more information
type PostgresSchemaType map[string]string

func (pl PostgresSchemaType) Type(v string) {
	pl[dialect.Postgres] = v
}

func (pl PostgresSchemaType) SQLite() SQLiteSchemaType {
	return (SQLiteSchemaType)(pl)
}

func (pl PostgresSchemaType) Gremlin() GremlinSchemaType {
	return (GremlinSchemaType)(pl)
}

func (pl PostgresSchemaType) MYSQL() MysqlSchemaType {
	return (MysqlSchemaType)(pl)
}

// Numeric Types consist of two-, four-, and eight-byte integers,
// four- and eight-byte floating-point numbers, and selectable-precision decimals

func (pl PostgresSchemaType) SmallInt() PostgresSchemaType {
	pl.Type("smallint")
	return pl
}

func (pl PostgresSchemaType) Integer() PostgresSchemaType {
	pl.Type("integer")
	return pl
}

func (pl PostgresSchemaType) BigInt() PostgresSchemaType {
	pl.Type("bigint")
	return pl
}

func (pl PostgresSchemaType) Decimal(p, s int) PostgresSchemaType {
	pl.Type(fmt.Sprintf("decimal(%d, %d)", p, s))
	return pl
}

func (pl PostgresSchemaType) Numeric() PostgresSchemaType {
	pl.Type("numeric")
	return pl
}

func (pl PostgresSchemaType) Real() PostgresSchemaType {
	pl.Type("real")
	return pl
}

func (pl PostgresSchemaType) SmallSerial() PostgresSchemaType {
	pl.Type("smallserial")
	return pl
}

func (pl PostgresSchemaType) BigSerial() PostgresSchemaType {
	pl.Type("bigserial")
	return pl
}

func (pl PostgresSchemaType) Boolean() PostgresSchemaType {
	pl.Type("boolean")
	return pl
}

func (pl PostgresSchemaType) Cidr() PostgresSchemaType {
	pl.Type("cidr")
	return pl
}

func (pl PostgresSchemaType) DoublePrecision() PostgresSchemaType {
	pl.Type("double precision")
	return pl
}

// Geometric Types

func (pl PostgresSchemaType) Circle() PostgresSchemaType {
	pl.Type("cidr")
	return pl
}

func (pl PostgresSchemaType) Lseg() PostgresSchemaType {
	pl.Type("lseg")
	return pl
}

func (pl PostgresSchemaType) Box() PostgresSchemaType {
	pl.Type("box")
	return pl
}

func (pl PostgresSchemaType) Line() PostgresSchemaType {
	pl.Type("line")
	return pl
}

func (pl PostgresSchemaType) Path() PostgresSchemaType {
	pl.Type("path")
	return pl
}

func (pl PostgresSchemaType) Point() PostgresSchemaType {
	pl.Type("point")
	return pl
}

func (pl PostgresSchemaType) Polygon() PostgresSchemaType {
	pl.Type("polygon")
	return pl
}

// string sequence data type

func (pl PostgresSchemaType) Bytea() PostgresSchemaType {
	pl.Type("bytea")
	return pl
}

func (pl PostgresSchemaType) Bit(n int) PostgresSchemaType {
	pl.Type(fmt.Sprintf("bit(%d)", n))
	return pl
}

func (pl PostgresSchemaType) VarBit(n int) PostgresSchemaType {
	pl.Type(fmt.Sprintf("varbit(%d)", n))
	return pl
}

func (pl PostgresSchemaType) BpChar() PostgresSchemaType {
	pl.Type("bpchar")
	return pl
}

func (pl PostgresSchemaType) Character(n int) PostgresSchemaType {
	pl.Type(fmt.Sprintf("character(%d)", n))
	return pl
}

func (pl PostgresSchemaType) CharacterVarying(n int) PostgresSchemaType {
	pl.Type(fmt.Sprintf("character varying(%d)", n))
	return pl
}

func (pl PostgresSchemaType) Text() PostgresSchemaType {
	pl.Type("text")
	return pl
}

// network address type

func (pl PostgresSchemaType) Inet() PostgresSchemaType {
	pl.Type("inet")
	return pl
}

func (pl PostgresSchemaType) MacAddr() PostgresSchemaType {
	pl.Type("macaddr")
	return pl
}

func (pl PostgresSchemaType) MacAddr8() PostgresSchemaType {
	pl.Type("macaddr")
	return pl
}

// time data type

func (pl PostgresSchemaType) Date() PostgresSchemaType {
	pl.Type("date")
	return pl
}

func (pl PostgresSchemaType) Time(p int) PostgresSchemaType {
	pl.Type(fmt.Sprintf("time(%d)", p))
	return pl
}

func (pl PostgresSchemaType) TimeZ(p int) PostgresSchemaType {
	pl.Type(fmt.Sprintf("time(%d) with the zone", p))
	return pl
}

func (pl PostgresSchemaType) Timestamp(p int) PostgresSchemaType {
	pl.Type(fmt.Sprintf("timestamp(%d)", p))
	return pl
}

func (pl PostgresSchemaType) TimestampZ(p int) PostgresSchemaType {
	pl.Type(fmt.Sprintf("timestamp(%d)", p))
	return pl
}

// other

func (pl PostgresSchemaType) PgLsn() PostgresSchemaType {
	pl.Type("pg_lsn")
	return pl
}

func (pl PostgresSchemaType) PgSnapShot() PostgresSchemaType {
	pl.Type("pg_snapshot")
	return pl
}

func (pl PostgresSchemaType) TsQuery() PostgresSchemaType {
	pl.Type("tsquery")
	return pl
}

func (pl PostgresSchemaType) TsVector() PostgresSchemaType {
	pl.Type("tsvector")
	return pl
}

func (pl PostgresSchemaType) TxidSnapshot() PostgresSchemaType {
	pl.Type("txid_snapshot")
	return pl
}

func (pl PostgresSchemaType) Uuid() PostgresSchemaType {
	pl.Type("uuid")
	return pl
}

func (pl PostgresSchemaType) Xml() PostgresSchemaType {
	pl.Type("xml")
	return pl
}

func (pl PostgresSchemaType) Interval(p int) PostgresSchemaType {
	pl.Type(fmt.Sprintf("interval(%d)", p))
	return pl
}

func (pl PostgresSchemaType) Json() PostgresSchemaType {
	pl.Type("json")
	return pl
}

func (pl PostgresSchemaType) JsonB() PostgresSchemaType {
	pl.Type("jsonb")
	return pl
}

func (pl PostgresSchemaType) Money() PostgresSchemaType {
	pl.Type("money")
	return pl
}
