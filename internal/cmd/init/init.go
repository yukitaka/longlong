package init

import (
	"database/sql"
	"fmt"
	_ "github.com/mattn/go-sqlite3"
	"github.com/spf13/cobra"
	"github.com/yukitaka/longlong/internal/cli"
	"github.com/yukitaka/longlong/internal/util"
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
			util.CheckErr(o.Run(cmd, args))
		},
	}

	cmd.AddCommand(&cobra.Command{
		Use:   "sqlite",
		Short: "Display one or many organizations",
		Run: func(cmd *cobra.Command, args []string) {
			util.CheckErr(o.Sqlite(cmd, args))
		},
	})

	return cmd
}

func (o *Options) Run(cmd *cobra.Command, args []string) error {
	fmt.Printf("Args is %v.", args)
	return nil
}

func (o *Options) Sqlite(cmd *cobra.Command, args []string) error {
	os.Remove("./longlong.db")
	db, err := sql.Open("sqlite3", "./longlong.db")
	if err != nil {
		return err
	}
	defer db.Close()

	query := `
	create table organizations (id integer not null primary key, name text);
	create table users (id integer not null primary key);
	create table profiles (id integer not null primary key, name text, full_name text);
	create table user_profiles(user_id integer not null, profile_id integer not null);
	`
	_, err = db.Exec(query)
	if err != nil {
		return err
	}

	return nil
}
