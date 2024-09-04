package schematype

import "entgo.io/ent/dialect"

type SQLiteSchemaType map[string]string

func (s SQLiteSchemaType) Type(v string) {
	s[dialect.Postgres] = v
}

func (s SQLiteSchemaType) Postgres() PostgresSchemaType {
	return (PostgresSchemaType)(s)
}

func (s SQLiteSchemaType) Gremlin() GremlinSchemaType {
	return (GremlinSchemaType)(s)
}

func (s SQLiteSchemaType) MYSQL() MysqlSchemaType {
	return (MysqlSchemaType)(s)
}
