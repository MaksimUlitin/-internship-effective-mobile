CREATE TABLE artists (
                         id SERIAL PRIMARY KEY,
                         name VARCHAR(255) NOT NULL UNIQUE,
                         created_at TIMESTAMP WITH TIME ZONE DEFAULT CURRENT_TIMESTAMP,
                         updated_at TIMESTAMP WITH TIME ZONE DEFAULT CURRENT_TIMESTAMP,
                         deleted_at TIMESTAMP WITH TIME ZONE
);

CREATE INDEX idx_artists_name ON artists (name);
CREATE INDEX idx_artists_deleted_at ON artists (deleted_at);

CREATE TABLE songs (
                       id SERIAL PRIMARY KEY,
                       artist_id INTEGER NOT NULL REFERENCES artists(id),
                       title VARCHAR(255) NOT NULL,
                       release_date DATE NOT NULL,
                       text TEXT,
                       link VARCHAR(2083),
                       created_at TIMESTAMP WITH TIME ZONE DEFAULT CURRENT_TIMESTAMP,
                       updated_at TIMESTAMP WITH TIME ZONE DEFAULT CURRENT_TIMESTAMP,
                       deleted_at TIMESTAMP WITH TIME ZONE
);

CREATE INDEX idx_songs_artist_id ON songs (artist_id);
CREATE INDEX idx_songs_title ON songs (title);
CREATE INDEX idx_songs_release_date ON songs (release_date);
CREATE INDEX idx_songs_deleted_at ON songs (deleted_at);