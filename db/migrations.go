package db

import (
	"database/sql"
	"fmt"
	"log"

	migrate "github.com/rubenv/sql-migrate"
)

// MigrateInfo struct
type MigrateInfo struct {
	Db *sql.DB
}

// RunMigrationUp up
func (s *MigrateInfo) RunMigrationUp() {
	up := "SELECT 0"
	down := "SELECT 1"
	migrations := &migrate.MemoryMigrationSource{
		Migrations: []*migrate.Migration{
			&migrate.Migration{
				Id:   "1",
				Up:   []string{up},
				Down: []string{down},
			},
		},
	}

	n, err := migrate.Exec(s.Db, "mysql", migrations, migrate.Up)
	if err != nil {
		log.Println(err)
	}
	fmt.Println(n)
}

func main() {
	s := MigrateInfo{}
	s.RunMigrationUp()
}
