package main

import (
	"context"
	"embed"
	"fmt"
	"log"
	"myapp/internal/initializers"
	"sort"

	"github.com/jackc/pgx/v5"
)

//go:embed db/migrations/*.sql
var migrationsFS embed.FS

func init() {
	initializers.LoadEnv()
	initializers.ConnectDB()
}

func main() {
	log.Println("Starting database migrations...")

	// Get connection from pool
	conn, err := initializers.DB.Acquire(context.Background())
	if err != nil {
		log.Fatalf("Failed to acquire connection: %v", err)
	}
	defer conn.Release()

	// Use the underlying pgx connection
	err = runMigrations(conn.Conn())
	if err != nil {
		log.Fatalf("Migrations failed: %v", err)
	}

	log.Println("Database migrations completed successfully!")
}

func runMigrations(conn *pgx.Conn) error {
	// Create schema_versions table
	_, err := conn.Exec(context.Background(), `
		CREATE TABLE IF NOT EXISTS schema_versions (
			version INT PRIMARY KEY,
			applied_at TIMESTAMPTZ DEFAULT NOW()
		)
	`)
	if err != nil {
		return fmt.Errorf("failed to create versions table: %w", err)
	}

	// Get current version
	var currentVersion int
	err = conn.QueryRow(context.Background(),
		"SELECT COALESCE(MAX(version), 0) FROM schema_versions").Scan(&currentVersion)
	if err != nil {
		return fmt.Errorf("failed to get current version: %w", err)
	}

	// Read and process migration files
	migrations, err := readMigrations(currentVersion)
	if err != nil {
		return err
	}

	for _, migration := range migrations {
		err = executeMigration(conn, migration)
		if err != nil {
			return err
		}
	}
	return nil
}

func executeMigration(conn *pgx.Conn, m migration) error {
	tx, err := conn.Begin(context.Background())
	if err != nil {
		return fmt.Errorf("failed to begin transaction: %w", err)
	}
	defer tx.Rollback(context.Background())

	// Execute migration SQL
	_, err = tx.Exec(context.Background(), m.sql)
	if err != nil {
		return fmt.Errorf("migration %d failed: %w", m.version, err)
	}

	// Record version
	_, err = tx.Exec(context.Background(),
		"INSERT INTO schema_versions (version) VALUES ($1)", m.version)
	if err != nil {
		return fmt.Errorf("failed to record version: %w", err)
	}

	return tx.Commit(context.Background())
}

type migration struct {
	version int
	sql     string
}

func readMigrations(currentVersion int) ([]migration, error) {
	var migrations []migration

	// Correct path for ReadDir
	entries, err := migrationsFS.ReadDir("db/migrations")
	if err != nil {
		return nil, fmt.Errorf("failed to read migrations: %w", err)
	}

	for _, entry := range entries {
		version := parseVersion(entry.Name())
		if version > currentVersion {
			// Correct path for ReadFile
			content, err := migrationsFS.ReadFile("db/migrations/" + entry.Name())
			if err != nil {
				return nil, fmt.Errorf("failed to read migration file: %w", err)
			}

			migrations = append(migrations, migration{
				version: version,
				sql:     string(content),
			})
		}
	}

	// Sort migrations in ascending order
	sort.Slice(migrations, func(i, j int) bool {
		return migrations[i].version < migrations[j].version
	})

	return migrations, nil
}

func parseVersion(filename string) int {
	var version int
	fmt.Sscanf(filename, "%d_", &version)
	return version
}
