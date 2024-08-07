package database

import "fmt"

func RunMigration() {
	conn, err := Connection()
	if err != nil {
		panic(err)
	}

	defer conn.Close()

	querys := []string{
		`CREATE TABLE IF NOT EXISTS polls (
			id VARCHAR PRIMARY KEY,
			title VARCHAR (255) NOT NULL,
			created_at TIMESTAMP NOT NULL DEFAULT CURRENT_TIMESTAMP
		)`,

		`CREATE TABLE IF NOT EXISTS questions (
  		id VARCHAR PRIMARY KEY,
  		poll_id VARCHAR REFERENCES polls(id) ON DELETE CASCADE,
  		option TEXT NOT NULL
		)`,

		`CREATE TABLE IF NOT EXISTS votes (
			id VARCHAR PRIMARY KEY,
			ip_address VARCHAR NOT NULL,
			voted_at TIMESTAMP NOT NULL DEFAULT CURRENT_TIMESTAMP,
			poll_id VARCHAR REFERENCES polls(id) ON DELETE CASCADE,
			voted_in VARCHAR REFERENCES questions(id) ON DELETE CASCADE
		)
		`,
	}

	for _, query := range querys {
		_, err = conn.Query(query)

		if err != nil {
			fmt.Println("Error has happen: ", err)
			panic(err)
		}
	}

	fmt.Println("Migration execution success!")
}
