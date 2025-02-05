--
-- PostgreSQL database dump
--

-- Dumped from database version 16.3
-- Dumped by pg_dump version 16.3

SET statement_timeout = 0;
SET lock_timeout = 0;
SET idle_in_transaction_session_timeout = 0;
SET client_encoding = 'UTF8';
SET standard_conforming_strings = on;
SELECT pg_catalog.set_config('search_path', '', false);
SET check_function_bodies = false;
SET xmloption = content;
SET client_min_messages = warning;
SET row_security = off;

SET default_tablespace = '';

SET default_table_access_method = heap;

--
-- Name: groups; Type: TABLE; Schema: public; Owner: admin
--

CREATE TABLE public.groups (
    id integer NOT NULL,
    name text NOT NULL
);


--
-- Name: groups_id_seq; Type: SEQUENCE; Schema: public; Owner: admin
--

CREATE SEQUENCE public.groups_id_seq
    AS integer
    START WITH 1
    INCREMENT BY 1
    NO MINVALUE
    NO MAXVALUE
    CACHE 1;


ALTER SEQUENCE public.groups_id_seq OWNER TO admin;

--
-- Name: groups_id_seq; Type: SEQUENCE OWNED BY; Schema: public; Owner: admin
--

ALTER SEQUENCE public.groups_id_seq OWNED BY public.groups.id;


--
-- Name: lyrics; Type: TABLE; Schema: public; Owner: admin
--

CREATE TABLE public.lyrics (
    song_id integer NOT NULL,
    order_n integer NOT NULL,
    lyrics text NOT NULL
);


ALTER TABLE public.lyrics OWNER TO admin;

--
-- Name: songs; Type: TABLE; Schema: public; Owner: admin
--

CREATE TABLE public.songs (
    id integer NOT NULL,
    name text NOT NULL,
    group_id integer,
    release_date date NOT NULL,
    link text NOT NULL
);


ALTER TABLE public.songs OWNER TO admin;

--
-- Name: songs_id_seq; Type: SEQUENCE; Schema: public; Owner: admin
--

CREATE SEQUENCE public.songs_id_seq
    AS integer
    START WITH 1
    INCREMENT BY 1
    NO MINVALUE
    NO MAXVALUE
    CACHE 1;


ALTER SEQUENCE public.songs_id_seq OWNER TO admin;

--
-- Name: songs_id_seq; Type: SEQUENCE OWNED BY; Schema: public; Owner: admin
--

ALTER SEQUENCE public.songs_id_seq OWNED BY public.songs.id;


--
-- Name: groups id; Type: DEFAULT; Schema: public; Owner: admin
--

ALTER TABLE ONLY public.groups ALTER COLUMN id SET DEFAULT nextval('public.groups_id_seq'::regclass);


--
-- Name: songs id; Type: DEFAULT; Schema: public; Owner: admin
--

ALTER TABLE ONLY public.songs ALTER COLUMN id SET DEFAULT nextval('public.songs_id_seq'::regclass);


--
-- Name: groups groups_pkey; Type: CONSTRAINT; Schema: public; Owner: admin
--

ALTER TABLE ONLY public.groups
    ADD CONSTRAINT groups_pkey PRIMARY KEY (id);


--
-- Name: lyrics lyrics_pkey; Type: CONSTRAINT; Schema: public; Owner: admin
--

ALTER TABLE ONLY public.lyrics
    ADD CONSTRAINT lyrics_pkey PRIMARY KEY (song_id, order_n);


--
-- Name: songs songs_pkey; Type: CONSTRAINT; Schema: public; Owner: admin
--

ALTER TABLE ONLY public.songs
    ADD CONSTRAINT songs_pkey PRIMARY KEY (id);


--
-- Name: lyrics lyrics_song_id_fkey; Type: FK CONSTRAINT; Schema: public; Owner: admin
--

ALTER TABLE ONLY public.lyrics
    ADD CONSTRAINT lyrics_song_id_fkey FOREIGN KEY (song_id) REFERENCES public.songs(id);


--
-- Name: songs songs_group_id_fkey; Type: FK CONSTRAINT; Schema: public; Owner: admin
--

ALTER TABLE ONLY public.songs
    ADD CONSTRAINT songs_group_id_fkey FOREIGN KEY (group_id) REFERENCES public.groups(id);


--
-- PostgreSQL database dump complete
--
