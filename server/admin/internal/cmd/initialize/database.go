package initialize

import (
	"github.com/jmoiron/sqlx"
)

type Database struct {
	*sqlx.DB
}

func NewDatabase(db *sqlx.DB) *Database {
	return &Database{db}
}

func (d *Database) Init() error {
	query := `
	drop table if exists user_profiles;
	drop table if exists individuals;
	drop table if exists profiles;
	drop table if exists users;
	drop table if exists organization_members;
	drop table if exists organizations;
	drop table if exists authentications;
	drop table if exists oauth_authentications;
	drop table if exists habits_timers;
	drop table if exists timers;
	drop table if exists habits;
	create table authentications (id integer not null primary key, identify text not null, token text not null, individual_id integer);
	create table oauth_authentications (identify text not null primary key, access_token text not null, refresh_token text, expiry timestamp, individual_id integer);
	create table organizations (id integer not null primary key, parent_id integer not null default 0, name text);
	create table organization_members (organization_id integer not null, individual_id integer not null, role integer);
	create table users (id integer not null primary key);
	create table profiles (id integer not null primary key, nick_name text, full_name text, biography text);
	create table individuals (id integer not null primary key, name text, user_id integer, profile_id integer);
	create table user_profiles (user_id integer not null, profile_id integer not null);
	create table timers (id integer not null primary key, duration_type varchar(16) not null, number integer, interval integer, reference_at timestamp);
	create table habits (id integer not null primary key, name text, exp integer, start_at timestamp, end_at timestamp);
	create table habits_timers (habit_id integer not null, timer_id integer not null, foreign key (habit_id) references habits (id), foreign key (timer_id) references timers (id));
	insert into organizations (id, name) values (1, 'longlong');
	insert into users (id) values (1);
	insert into profiles (id, nick_name, full_name, biography) values (1, 'admin', 'Admin', 'I am a administrator.');
	insert into user_profiles (user_id, profile_id) values (1, 1);
	insert into individuals (id, name, user_id, profile_id) values (1, 'admin', 1, 1);
	insert into organization_members (organization_id, individual_id, role) values (1, 1, 0);
	`
	if _, err := d.Exec(query); err != nil {
		return err
	}

	return nil
}
