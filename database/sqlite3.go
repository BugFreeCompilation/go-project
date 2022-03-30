package database

import (
	"BugFreeCompilation/go-project/entity"
	err "BugFreeCompilation/go-project/error"
	"database/sql"
	"log"
	"os"

	_ "github.com/mattn/go-sqlite3"
)

type databaseStruct struct{}

func InitSQLite3() Database {
	os.Remove("database.db")

	_, err_create := os.Create("database.db")
	if err_create != nil {
		log.Fatal(err_create)
	}

	sqliteDatabase, err_open := sql.Open("sqlite3", "database.db")
	if err_open != nil {
		log.Fatal(err_open)
	}
	defer sqliteDatabase.Close()

	// creating database that won't add the same entry twice and auto-increments id
	// https://stackoverflow.com/questions/7905859/is-there-an-auto-increment-in-sqlite
	// https://stackoverflow.com/questions/29312882/sqlite-preventing-duplicate-rows
	sqlStatement := `
		CREATE TABLE posts (id INTEGER NOT NULL PRIMARY KEY UNIQUE, title TEXT UNIQUE);
	`
	_, err_exec := sqliteDatabase.Exec(sqlStatement)
	if err_exec != nil {
		log.Fatal(err_exec)
	}

	// adding one entry as a test
	sqlStatement = `INSERT INTO posts(title) VALUES (?)`
	statement, err_addon := sqliteDatabase.Prepare(sqlStatement)
	if err_addon != nil {
		log.Fatal(err_addon)
	}
	var test_name string = "Titanic"
	statement.Exec(test_name)

	return &databaseStruct{}
}

func (*databaseStruct) Add(post *entity.Post) error {
	sqliteDatabase, err_open := sql.Open("sqlite3", "database.db")
	if err_open != nil {
		return err.ErrDatabaseOpen
	}
	defer sqliteDatabase.Close()

	sqlStatement := `INSERT INTO posts(title) VALUES (?)`
	statement, err_prepare := sqliteDatabase.Prepare(sqlStatement)
	if err_prepare != nil {
		return err.ErrDatabasePrepare
	}
	statement.Exec(post.TITLE)

	// set id to current db id
	// LastInsertId() will only work if the element is not there and gets added
	// and won't work if it gets added a second time
	// workaround by doing a SELECT statement, adding error handling might be useful
	// https://go.dev/doc/database/querying
	sqliteDatabase.QueryRow(`SELECT id FROM posts WHERE title=(?)`,
		post.TITLE).Scan(&post.ID)

	return nil
}

func (*databaseStruct) Delete(post *entity.Post) error {
	sqliteDatabase, err_open := sql.Open("sqlite3", "database.db")
	if err_open != nil {
		return err.ErrDatabaseOpen
	}
	defer sqliteDatabase.Close()

	// get id before deleting
	// the database starts counting from 1, and that means that the
	// id will stay at 0 if delete is invalid and does not match in database
	sqliteDatabase.QueryRow(`SELECT id FROM posts WHERE title=(?)`,
		post.TITLE).Scan(&post.ID)

	if post.ID == 0 {
		return err.ErrDatabaseEntry
	}

	sqlStatement := `DELETE FROM posts WHERE title=(?)`
	statement, err_prepare := sqliteDatabase.Prepare(sqlStatement)
	if err_prepare != nil {
		return err.ErrDatabasePrepare
	}
	statement.Exec(post.TITLE)

	return nil
}

func (*databaseStruct) GetAll() ([]entity.Post, error) {
	sqliteDatabase, err_open := sql.Open("sqlite3", "database.db")
	if err_open != nil {
		return nil, err.ErrDatabaseOpen
	}
	defer sqliteDatabase.Close()

	var posts []entity.Post
	statement, err_query := sqliteDatabase.Query("SELECT * FROM posts")
	if err_query != nil {
		return nil, err.ErrDatabaseQuery
	}
	defer statement.Close()

	for statement.Next() {
		var id int
		var title string
		statement.Scan(&id, &title)
		post := entity.Post{
			ID:    id,
			TITLE: title,
		}
		posts = append(posts, post)
	}

	return posts, nil
}
