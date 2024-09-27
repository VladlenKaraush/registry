package main

import (
	"database/sql"
	"log"
)

type DbRepo struct {
	dbConn *sql.DB
}

func CreateDbRepo(config Config) DbRepo {
	db_conn, err := sql.Open(config.Db.Driver, config.Db.Filepath)
	if err != nil {
		log.Fatal(err)
	}
	repo := DbRepo{dbConn: db_conn}
	return repo
}

func (db *DbRepo) executeStatement(sql_statement string) {
	statement, err := db.dbConn.Prepare(sql_statement) // Prepare SQL Statement
	if err != nil {
		log.Fatal(err.Error())
	}
	statement.Exec()
}

func (db *DbRepo) CreatePackageTable() {
	createStudentTableSQL := `CREATE TABLE IF NOT EXISTS packages (
		"idPackage" integer NOT NULL PRIMARY KEY AUTOINCREMENT,		
		"repository_id" integer NOT NULL,
		"release_id" integer NOT NULL,		
		"name" TEXT NOT NULL ,
		"version" TEXT NOT NULL,
		"architecture" TEXT NOT NULL,
		"file_path" TEXT NOT NULL,
		"description" TEXT
	  );` // SQL Statement for Create Table

	log.Println("Create packages table...")
	db.executeStatement(createStudentTableSQL)
	log.Println("packages table created")
}
