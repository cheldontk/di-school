package database

func CreateProductTable() {
	DB.Query(`CREATE TABLE IF NOT EXISTS products(
	id SERIAL PRIMARY KEY,
	amount INTEGER,
	name TEXT UNIQUE,
	description TEXT,
	category TEXT NOT NULL
)
`)
}
