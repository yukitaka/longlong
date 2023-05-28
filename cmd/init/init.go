package main

import (
	"database/sql"
	"fmt"
	"github.com/joho/godotenv"
	_ "github.com/lib/pq"
	_ "github.com/mattn/go-sqlite3"
	"golang.org/x/crypto/bcrypt"
	"os"
)

func main() {
	err := sqlite()
	if err != nil {
		fmt.Printf("Error: %v", err)
	}
	err = postgres()
	if err != nil {
		fmt.Printf("Error: %v", err)
	}
}

func postgres() error {
	err := godotenv.Load(".env")
	if err != nil {
		return err
	}

	fmt.Println("Open db.")
	host := os.Getenv("POSTGRES_HOST")
	dbName := os.Getenv("POSTGRES_DB")
	user := os.Getenv("POSTGRES_USER")
	password := os.Getenv("POSTGRES_PASSWORD")
	port := os.Getenv("POSTGRES_PORT")
	dsl := fmt.Sprintf("host=%s port=%s user=%s dbname=%s password=%s sslmode=disable", host, port, user, dbName, password)
	db, err := sql.Open("postgres", dsl)
	if err != nil {
		return err
	}
	defer func(db *sql.DB) {
		err := db.Close()
		if err != nil {
			fmt.Printf("Error: %v", err)
		} else {
			fmt.Println("Close db.")
		}
	}(db)

	return initSql(db)
}

func sqlite() error {
	filename := "./longlong.db"
	if _, err := os.Stat(filename); err == nil {
		err := os.Remove(filename)
		if err != nil {
			return err
		}
	}
	fmt.Println("Open db.")
	db, err := sql.Open("sqlite3", filename)
	if err != nil {
		return err
	}
	defer func(db *sql.DB) {
		err := db.Close()
		if err != nil {
			fmt.Printf("Error: %v", err)
		} else {
			fmt.Println("Close db.")
		}
	}(db)

	return initSql(db)
}

func initSql(db *sql.DB) error {
	hash, err := bcrypt.GenerateFromPassword([]byte("password"), bcrypt.DefaultCost)
	if err != nil {
		return err
	}

	fmt.Println("Init data.")
	query := `
	drop table if exists user_profiles;
	drop table if exists individuals;
	drop table if exists profiles;
	drop table if exists users;
	drop table if exists organization_members;
	drop table if exists organizations;
	drop table if exists authentications;
	create table authentications (id integer not null primary key, identify text not null, token text not null, individual_id integer);
	create table organizations (id integer not null primary key, parent_id integer not null default 0, name text);
	create table organization_members (organization_id integer not null, individual_id integer not null, role integer);
	create table users (id integer not null primary key);
	create table profiles (id integer not null primary key, nick_name text, full_name text, biography text);
	create table individuals (id integer not null primary key, name text, user_id integer, profile_id integer);
	create table user_profiles(user_id integer not null, profile_id integer not null);
	insert into organizations (id, name) values (1, 'longlong');
	insert into users (id) values (1);
	insert into profiles (id, nick_name, full_name, biography) values (1, 'yukitaka', 'Takayuki Sato', 'I am a software engineer.');
	insert into user_profiles (user_id, profile_id) values (1, 1);
	insert into individuals (id, name, user_id, profile_id) values (1, 'yukitaka', 1, 1);
	insert into organization_members (organization_id, individual_id, role) values (1, 1, 0);
	insert into authentications (id, identify, token, individual_id) values (1, 'yukitaka', '%s', 1);
	`
	_, err = db.Exec(fmt.Sprintf(query, hash))
	if err != nil {
		return err
	}

	return nil

}
