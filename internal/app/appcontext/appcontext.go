package appcontext

import (
	"github.com/pembajak/personal-finance/internal/pkg/driver/driversql"
	"github.com/spf13/viper"
)

type AppContext struct{}

// NewAppContext ...
func NewAppContext() *AppContext {
	appCtx := &AppContext{}

	return appCtx
}

// GetDB ...
func (act *AppContext) GetDB() (db *driversql.Database, err error) {
	dbConfig := driversql.DBMysqlOption{
		Host:     viper.GetString("database.host"),
		Port:     viper.GetString("database.port"),
		DB:       viper.GetString("database.db"),
		Username: viper.GetString("database.username"),
		Password: viper.GetString("database.password"),
	}

	db, err = driversql.NewMysqlDatabase(dbConfig)

	return
}
