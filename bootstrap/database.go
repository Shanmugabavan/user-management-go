package bootstrap

import (
	"context"
	"fmt"
	"log"
	"time"

	"github.com/jackc/pgx/v5/pgxpool"
)

func GetConnectionPool(env *Env) *pgxpool.Pool {
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()

	connectionString := fmt.Sprintf(
		"postgres://%s:%s@%s:%s/%s?sslmode=%s",
		env.DBUser,
		env.DBPass,
		env.DBHost,
		env.DBPort,
		env.DBName,
		env.DBSSLMode,
	)

	var err error

	db, err := pgxpool.New(ctx, connectionString)
	if err != nil {
		log.Fatal("Unable to create connection pool:", err)
	}

	if err = db.Ping(ctx); err != nil {
		log.Fatal("Database connection failed:", err)
	}

	fmt.Println("Successfully connected to database")

	return db
}

func CloseConnectionPool(db *pgxpool.Pool) {
	db.Close()
}
