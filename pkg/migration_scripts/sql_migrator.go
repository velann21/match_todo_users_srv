package migration_scripts

import (
	"database/sql"
	"fmt"
	"github.com/golang-migrate/migrate"
	"github.com/golang-migrate/migrate/database/mysql"
	"github.com/golang-migrate/migrate/source/file"
)


func MigrateDb(stepCount uint) error {
	fmt.Println("In main")
	fsrc, err := (&file.File{}).Open("file://")
	if err != nil {
		fmt.Println(err)
		return err
	}
	db, err := sql.Open("mysql",  "root:Siar@123@tcp(localhost:3306)/todousersrv?"+"multiStatements=true")
	if err != nil {
		return err
	}
	driver, _ := mysql.WithInstance(db, &mysql.Config{})
	m, err := migrate.NewWithInstance("", fsrc, "todousersrv1", driver)
	defer func(m *migrate.Migrate){
		m.Close()
	}(m)
	if err != nil {
		fmt.Println("Eror")
		return err
	}
	err = m.Migrate(stepCount)
	if err != nil {
		fmt.Println(err)
		return err
	}
	fmt.Println("Done")

	return nil
}

