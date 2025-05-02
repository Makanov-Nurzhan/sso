package main

import (
	"errors"
	"flag"
	"fmt"
	// Библиотека для миграций
	"github.com/golang-migrate/migrate/v4"
	// Драйвер для выполнения миграций SQLite 3
	_ "github.com/golang-migrate/migrate/v4/database/sqlite"
	// Драйвер для получения миграций из файлов
	_ "github.com/golang-migrate/migrate/v4/source/file"
)

func main() {
	var storagePath, migrationsPath, migrationsTable string

	flag.StringVar(&storagePath, "storage-path", "", "path to store the migrations")
	flag.StringVar(&migrationsPath, "migrations-path", "", "path to store the migrations")
	flag.StringVar(&migrationsTable, "migrations-table", "migrations", "name of the migrations table")
	flag.Parse()

	if storagePath == "" {
		panic("--storage-path flag is required")
	}
	if migrationsPath == "" {
		panic("--migrations-path flag is required")
	}

	m, err := migrate.New(
		"file://"+migrationsPath,
		fmt.Sprintf("sqlite://%s?x-migrations-table=%s", storagePath, migrationsTable),
	)
	if err != nil {
		panic(err)
	}
	if err := m.Up(); err != nil {
		if errors.Is(err, migrate.ErrNoChange) {
			fmt.Println("No migrations to apply")
			return
		}
		panic(err)
	}
	fmt.Println("Applied migrations")
}
