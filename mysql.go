package schematype

import (
	"entgo.io/ent/dialect"
)

// MySQL returns mysql schema type builder
func MySQL() *MysqlSchemaType {
	return &MysqlSchemaType{}
}

type MysqlSchemaType map[string]string

func (s *MysqlSchemaType) Raw(v string) {
	(*s)[dialect.MySQL] = v
}

func (s *MysqlSchemaType) get() string {
	return (*s)[dialect.MySQL]
}

func (s *MysqlSchemaType) Postgres() *PostgresSchemaType {
	return (*PostgresSchemaType)(s)
}

func (s *MysqlSchemaType) SQLite() *SQLiteSchemaType {
	return (*SQLiteSchemaType)(s)
}

func (s *MysqlSchemaType) Gremlin() *GremlinSchemaType {
	return (*GremlinSchemaType)(s)
}
