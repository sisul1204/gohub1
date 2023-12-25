package config

import "gohub/pkg/config"

func init() {
	config.Add("database", func() map[string]interface{} {
		return map[string]interface{}{
			"connection": config.Env("DB_CONNECTION", "mysql"),
			"mysql": map[string]interface{}{
				"host":     config.Env("DB_HOST", "124.71.189.186"),
				"port":     config.Env("DB_PORT", "3306"),
				"database": config.Env("DB_DATABASE", "gohub"),
				"username": config.Env("DB_USERNAME", ""),
				"password": config.Env("DB_PASSWORD", ""),
				"charset":  "utf8mb4",

				"max_idle_connections": config.Env("DB_MAX_IDLE_CONNECTIONS", 25),
				"max_open_connections": config.Env("DB_MAX_OPEN_CONNECTIONS", 100),
				"max_life__seconds":    config.Env("DB_MAX_LIFE__SECONDS", 5*60),
			},
			"sqlite": map[string]interface{}{
				"database": config.Env("DB_SQL_FILE", "database/database.db"),
			},
		}
	})
}
