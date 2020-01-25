package storage

const (
	Schema = `
        CREATE TABLE user(
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

		CREATE TABLE series_type( -- TV, Movie, OVA, etc.
            series_type_id INTEGER PRIMARY KEY NOT NULL,
            name VARCHAR(255) UNIQUE
        );


        CREATE TABLE review( -- User series join table
            review_id INTEGER PRIMARY KEY NOT NULL,
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
            comments TEXT,
            tags TEXT,
            times_watched INTEGER,
            rewatching INTEGER,
            rewatch_value INTEGER,
            rewatching_ep INTEGER,	
            
            UNIQUE(user_id, series_id)
        );

        CREATE TABLE review_status( -- Plan to watch, Watched, etc.
            review_status_id INTEGER PRIMARY KEY NOT NULL,
            name VARCHAR(255) UNIQUE
        );

        CREATE TABLE review_storage_type( -- HDD, DVD, etc.
            review_storage_type_id INTEGER PRIMARY KEY NOT NULL,
            name VARCHAR(255) UNIQUE
        );

		CREATE TABLE review_rewatch_value( -- Very Low, Low, etc.
            review_rewatch_value_id INTEGER PRIMARY KEY NOT NULL,
            name VARCHAR(255) UNIQUE
        );
    `
)
