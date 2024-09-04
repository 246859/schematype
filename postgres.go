package schematype

import "entgo.io/ent/dialect"

type PostgresSchemaType map[string]string

func (p PostgresSchemaType) Type(v string) {
	p[dialect.Postgres] = v
}

func (p PostgresSchemaType) SQLite() SQLiteSchemaType {
	return (SQLiteSchemaType)(p)
}

func (p PostgresSchemaType) Gremlin() GremlinSchemaType {
	return (GremlinSchemaType)(p)
}

func (p PostgresSchemaType) MYSQL() MysqlSchemaType {
	return (MysqlSchemaType)(p)
}
