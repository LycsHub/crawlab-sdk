package interfaces

type SQLDb interface {
	CreateTB(models ...interface{}) error
	RawSQL(sql string) error
	TB(name string) SQLTb
}
