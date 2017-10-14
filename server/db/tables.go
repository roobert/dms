package db

func CreateTables() {
	clientsSchema := `
		id    INTEGER PRIMARY KEY,
		name  TEXT    NOT NULL,
		date  TEXT    NOT NULL,

		UNIQUE (date, name),
		UNIQUE (id, name)
	`

	CreateTable("clients", clientsSchema)

	resultsSchema := `
		id     INT  NOT NULL,
		slot   INT  NOT NULL,

		UNIQUE (id, slot),

		FOREIGN KEY (id) REFERENCES client(id)
	`

	CreateTable("results", resultsSchema)
}
