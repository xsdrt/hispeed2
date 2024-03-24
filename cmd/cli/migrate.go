package main

func doMigrate(arg2, arg3 string) error {
	dsn := getDSN()

	// run the migration command
	switch arg2 {
	case "up":
		err := his.MigrateUp(dsn)
		if err != nil {
			return err
		}

	case "down":
		if arg3 == "all" { // migrate down all migrations....
			err := his.MigrateDownAll(dsn)
			if err != nil {
				return err
			}
		} else {
			err := his.Steps(-1, dsn) // otherwise, if no arg3;  migrate down/back only the recent migration
			if err != nil {
				return err
			}
		}
	case "reset": // if want to reset the database: run MigrateDownALl migrations and then run MigrateUp...
		err := his.MigrateDownAll(dsn)
		if err != nil {
			return err
		}
		err = his.MigrateUp(dsn)
		if err != nil {
			return err
		}
	default: // if things went wrong show help...
		showHelp()
	}

	return nil
}
