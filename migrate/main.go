package main

import (
	"errors"
	"fmt"
	"github.com/golang-migrate/migrate/v4"
	_ "github.com/golang-migrate/migrate/v4/database/postgres"
	_ "github.com/golang-migrate/migrate/v4/source/file"
)

func main() {
	m, err := migrate.New("file://./migrations", "postgres://postgres:root@localhost/song_library?sslmode=disable") // вынести куда нибудь путь
	if err != nil {
		panic(err)
	}

	if err := m.Up(); err != nil {
		if errors.Is(err, migrate.ErrNoChange) {
			fmt.Println("no migrations for apply")
		}
		panic(err)
	}
	fmt.Println("migrations successfully apply")
}