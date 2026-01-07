package bootstrap

import "github.com/jackc/pgx/v5/pgxpool"

type Application struct {
	Env            *Env
	ConnectionPool *pgxpool.Pool
}

func App() Application {
	app := &Application{}
	app.Env = NewEnv()
	app.ConnectionPool = GetConnectionPool(app.Env)
	return *app
}

func (app *Application) CloseDBConnectionPool() {
	CloseConnectionPool(app.ConnectionPool)
}
