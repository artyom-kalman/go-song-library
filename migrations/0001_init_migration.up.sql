CREATE TABLE IF NOT EXISTS public.groups (
    id serial NOT NULL PRIMARY KEY,
    name text NOT NULL
);

CREATE INDEX idx_groups_name ON public.groups (name);

CREATE TABLE IF NOT EXISTS public.songs (
    id serial NOT NULL PRIMARY KEy,
    name text NOT NULL,
    group_id integer references groups (id),
    release_date date NOT NULL,
    link text NOT NULL
);

CREATE INDEX idx_songs_name ON public.songs (name);
CREATE INDEX idx_songs_group_id ON public.songs (group_id);

CREATE TABLE IF NOT EXISTS public.lyrics (
    song_id integer NOT NULL REFERENCES songs (id),
    order_n integer NOT NULL,
    lyrics text NOT NULL,
    PRIMARY KEY (song_id, order_n)
);

CREATE INDEX idx_lyrics_song_id ON public.lyrics (song_id);
