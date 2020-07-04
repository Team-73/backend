package entity

// InitialConfig entity
type InitialConfig struct {
	Username         string `json:"username"`
	Password         string `json:"password"`
	Host             string `json:"host"`
	Port             string `json:"port"`
	DBName           string `json:"db_name"`
	MaxLifeInMinutes int    `json:"max_life_minutes"`
	MaxIdleConns     int    `json:"max_idle_conns"`
	MaxOpenConns     int    `json:"max_open_conns"`
}
