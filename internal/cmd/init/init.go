package init

import (
	"database/sql"
	"fmt"
	_ "github.com/mattn/go-sqlite3"
	"github.com/spf13/cobra"
	"github.com/yukitaka/longlong/internal/cli"
	"github.com/yukitaka/longlong/internal/util"
	"golang.org/x/crypto/bcrypt"
	"os"
)

type Options struct {
	CmdParent string
	cli.IOStream
}

func NewInitOptions(parent string, streams cli.IOStream) *Options {
	return &Options{
		CmdParent: parent,
		IOStream:  streams,
	}
}

func NewCmdInit(parent string, streams cli.IOStream) *cobra.Command {
	o := NewInitOptions(parent, streams)

	cmd := &cobra.Command{
		Use:   "init",
		Short: "Display one or many resources",
		Run: func(cmd *cobra.Command, args []string) {
			util.CheckErr(o.Run(args))
		},
	}

	cmd.AddCommand(&cobra.Command{
		Use:   "sqlite",
		Short: "Display one or many organizations",
		Run: func(cmd *cobra.Command, args []string) {
			util.CheckErr(o.Sqlite())
		},
	})

	return cmd
}

func (o *Options) Run(args []string) error {
	fmt.Printf("Args is %v.", args)
	return nil
}

func (o *Options) Sqlite() error {
	err := os.Remove("./longlong.db")
	if err != nil {
		return err
	}
	db, err := sql.Open("sqlite3", "./longlong.db")
	if err != nil {
		return err
	}
	defer func(db *sql.DB) {
		err := db.Close()
		if err != nil {
			fmt.Printf("Error: %v", err)
		}
	}(db)

	hash, err := bcrypt.GenerateFromPassword([]byte("password"), bcrypt.DefaultCost)
	if err != nil {
		return err
	}

	query := `
	create table authenticates (id integer not null primary key, name text, token text);
	create table organizations (id integer not null primary key, name text);
	create table users (id integer not null primary key);
	create table profiles (id integer not null primary key, name text, full_name text);
	create table user_profiles(user_id integer not null, profile_id integer not null);
	insert into organizations (id, name) values (1, 'longlong');
	insert into users (id) values (1);
	insert into profiles (id, name, full_name) values (1, 'yukitaka', 'Yuki Sato');
	insert into user_profiles (user_id, profile_id) values (1, 1);
	insert into authenticates (id, name, token) values (1, 'yukitaka', '%s');
	`
	_, err = db.Exec(fmt.Sprintf(query, hash))
	if err != nil {
		return err
	}

	return nil
}
