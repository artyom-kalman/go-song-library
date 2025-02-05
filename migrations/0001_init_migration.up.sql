--
-- PostgreSQL database dump
--
CREATE TABLE IF NOT EXISTS public.groups (
    id serial NOT NULL PRIMARY KEY,
    name text NOT NULL
);

CREATE TABLE IF NOT EXISTS public.lyrics (
    song_id integer NOT NULL,
    order_n integer NOT NULL,
    lyrics text NOT NULL,
    PRIMARY KEY (song_id, order_n)
);

CREATE TABLE IF NOT EXISTS public.songs (
    id serial NOT NULL PRIMARY KEy,
    name text NOT NULL,
    group_id integer references groups (id),
    release_date date NOT NULL,
    link text NOT NULL
);

--
-- PostgreSQL database dump complete
--
