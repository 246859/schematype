package schematype

import "entgo.io/ent/dialect"

type GremlinSchemaType map[string]string

func (g *GremlinSchemaType) Set(typ string) {
	(*g)[dialect.Gremlin] = typ
}

func (g *GremlinSchemaType) Postgres() *PostgresSchemaType {
	return (*PostgresSchemaType)(g)
}

func (g *GremlinSchemaType) SQLite() *SQLiteSchemaType {
	return (*SQLiteSchemaType)(g)
}

func (g *GremlinSchemaType) MYSQL() *MysqlSchemaType {
	return (*MysqlSchemaType)(g)
}
