package storage

const (
	Schema = `
        CREATE TABLE user( -- Users didn't need to be included but I want the flexibility
            user_id INTEGER PRIMARY KEY NOT NULL,
            name VARCHAR(255) UNIQUE
        );

        CREATE TABLE series( -- Individual shows
            series_id INTEGER PRIMARY KEY NOT NULL,
            animedb_id INTEGER,
            name VARCHAR(255),
            series_type_id INTEGER,
            episodes INTEGER
        );

        CREATE TABLE user_series( -- User series join table
            user_series_id INTEGER PRIMARY KEY NOT NULL,
            user_id INTEGER NOT NULL,
            series_id INTEGER NOT NULL,

            storage_type_id INTEGER,	
            user_status_id INTEGER,

            watched_episodes INTEGER,
            start_date INTEGER,
            finish_date INTEGER,
            rated INTEGER,
            score INTEGER,
            dvd INTEGER,
            tags TEXT,
            comments TEXT,
            times_watched INTEGER,
            rewatching INTEGER,
            rewatch_value INTEGER,
            rewatching_ep INTEGER,	
            
            UNIQUE(user_id, series_id)
        );


        CREATE TABLE series_type( -- TV, Movie, etc.
            series_type_id INTEGER PRIMARY KEY NOT NULL,
            name VARCHAR(255) UNIQUE
        );

        CREATE TABLE user_status( -- Plan to watch, Watched, etc.
            user_status_id INTEGER PRIMARY KEY NOT NULL,
            name VARCHAR(255) UNIQUE
        );

        CREATE TABLE storage_type( -- HDD, DVD, etc.
            storage_type_id INTEGER PRIMARY KEY NOT NULL,
            name VARCHAR(255) UNIQUE
        );

        CREATE TABLE user_tag( -- User defined tags
            user_tag_id INTEGER PRIMARY KEY NOT NULL,
            name VARCHAR(255) UNIQUE
        );

        CREATE TABLE user_series_tag ( -- User series tags join table
            user_series_tag_id INTEGER PRIMARY KEY NOT NULL,
            user_series_id INTEGER NOT NULL,
            user_tag_id INTEGER NOT NULL,

            UNIQUE(user_series_id, user_tag_id)
        );
    `
)
