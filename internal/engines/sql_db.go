package engines

import (
	"github.com/LycsHub/crawlab-sdk/internal/driver"
	"github.com/LycsHub/crawlab-sdk/internal/interfaces"
	"github.com/crawlab-team/go-trace"
	"github.com/zsmj-xu/gorm"
)

type SQLDb struct {
	interfaces.SQLDb
	_DB *gorm.DB
}

func NewSQLDb(name string) interfaces.SQLDb {
	db, err := driver.SQL.New(name)
	if err != nil {
		panic(trace.TraceError(err))
	}

	return &SQLDb{_DB: db}
}

func (my *SQLDb) RawSQL(sql string) error {
	var gormDb *gorm.DB
	gormDb = my._DB.Exec(sql)
	return gormDb.Error
}

func (my *SQLDb) CreateTB(models ...interface{}) error {
	for _, model := range models {
		err := my._DB.AutoMigrate(model)
		if err != nil {
			return err
		}
	}

	return nil
}

func (my *SQLDb) TB(name string) interfaces.SQLTb {
	return &SQLTb{_InstanceFn: func() *gorm.DB {
		return my._DB.Table(name)
	}}
}
