# schematype
schema type helper for ent fields

## install
```bash
go get -u github.com/246859/schematype
```

## example
```go
func (User) Fields() []ent.Field {
    return []ent.Field{
		// mysql varchar
        field.String("name").SchemaType(schematype.MySQL().VarChar(255)),
        field.Int("age").SchemaType(schematype.MySQL().TinyInt(0)),
    }
}
```
