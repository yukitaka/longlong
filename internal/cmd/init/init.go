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
	create table authentications (id integer not null primary key, identify text not null, token text not null, individual_id integer);
	create table organizations (id integer not null primary key, parent_id integer not null default 0, name text);
	create table organization_belongings (organization_id integer not null, individual_id integer not null, role integer);
	create table users (id integer not null primary key);
	create table profiles (id integer not null primary key, name text, full_name text);
	create table individuals (id integer not null primary key, name text, user_id integer, profile_id integer);
	create table user_profiles(user_id integer not null, profile_id integer not null);
	insert into organizations (id, name) values (1, 'longlong');
	insert into users (id) values (1);
	insert into profiles (id, name, full_name) values (1, 'yukitaka', 'Yuki Sato');
	insert into user_profiles (user_id, profile_id) values (1, 1);
	insert into individuals (id, name, user_id, profile_id) values (1, 'yukitaka', 1, 1);
	insert into organization_belongings (organization_id, individual_id, role) values (1, 1, 0);
	insert into authentications (id, identify, token, individual_id) values (1, 'yukitaka', '%s', 1);
	`
	_, err = db.Exec(fmt.Sprintf(query, hash))
	if err != nil {
		return err
	}

	return nil
}
