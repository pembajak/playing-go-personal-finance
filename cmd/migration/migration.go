package migration

import (
	"database/sql"
	"fmt"
	"os"
	"time"

	"github.com/golang-migrate/migrate/v4"
	"github.com/golang-migrate/migrate/v4/database"
	"github.com/golang-migrate/migrate/v4/database/mysql"
	_ "github.com/golang-migrate/migrate/v4/source/file"
	"github.com/pembajak/personal-finance/internal/pkg/driver/driversql"
)

// IMigration ...
type IMigration interface {
	Up() error
	Down() error
	Create(string, string) error
}

// Migration ...
type Migration struct {
	ormRaw *sql.DB
	driver *driversql.Database
}

// New ...
func New(driver *driversql.Database) IMigration {
	return &Migration{
		driver: driver,
	}
}

// Up ...
func (migration *Migration) Up() (err error) {
	mgr, err := migration.migrateInstance()
	if err != nil {
		return
	}
	defer mgr.Close()
	return mgr.Up()
}

// Down ...
func (migration *Migration) Down() (err error) {
	mgr, err := migration.migrateInstance()
	if err != nil {
		return
	}
	defer mgr.Close()
	return mgr.Down()
}

// Create ...
func (migration *Migration) Create(name string, ext string) (err error) {
	base := fmt.Sprintf("%v/%v_%v.", "migrations/sql", time.Now().Unix(), name)
	err = migration.createFile(base + "up." + ext)
	if err != nil {
		return
	}

	err = migration.createFile(base + "down." + ext)
	return
}

func (migration *Migration) createFile(fname string) (err error) {
	_, err = os.Create(fname)
	return
}

func (migration *Migration) migrateInstance() (mgr *migrate.Migrate, err error) {
	driver, err := migration.getDriver()
	if err != nil {
		return
	}

	mgr, err = migrate.NewWithDatabaseInstance(
		fmt.Sprintf("file://%s", "migrations/sql"),
		"mysql",
		driver,
	)
	if err != nil {
		return
	}

	return
}

func (migration *Migration) getDriver() (driver database.Driver, err error) {
	migration.ormRaw, _ = migration.driver.Instance.DB()

	driver, err = mysql.WithInstance(migration.ormRaw, &mysql.Config{})
	if err != nil {
		return
	}
	return
}
