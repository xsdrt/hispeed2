package hispeed2

import (
	"log"

	"github.com/golang-migrate/migrate/v4"
)

func (h *HiSpeed2) MigrateUp(dsn string) error {
	m, err := migrate.New("file://"+h.RootPath+"/migrations", dsn) // create /open up the migration file...
	if err != nil {
		return err
	}
	defer m.Close()

	if err = m.Up(); err != nil {
		log.Println("Error running migration: ", err)
		return err
	}
	return nil
}

func (h *HiSpeed2) MigrateDownAll(dsn string) error {
	m, err := migrate.New("file://"+h.RootPath+"/migrations", dsn)
	if err != nil {
		return err
	}
	defer m.Close()

	if err := m.Down(); err != nil {
		return err
	}

	return nil
}

func (h *HiSpeed2) Steps(n int, dsn string) error {
	m, err := migrate.New("file://"+h.RootPath+"/migrations", dsn)
	if err != nil {
		return err
	}
	defer m.Close()

	if err := m.Steps(n); err != nil {
		return err
	}

	return nil
}

func (h *HiSpeed2) MigrateForce(dsn string) error { //if you have an error in the migration file, might be marked dirty in the DB , so force the migration...
	m, err := migrate.New("file://"+h.RootPath+"/migrations", dsn)
	if err != nil {
		return err
	}
	defer m.Close()

	if err := m.Force(-1); err != nil { // So we will force the migration down 1... allows oportunity to fix  and retry the migration after we fix the problem in our migration file
		return err
	}

	return nil
}
