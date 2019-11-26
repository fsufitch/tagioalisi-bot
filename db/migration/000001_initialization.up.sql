-- Dumped from database version 10.11 (Ubuntu 10.11-1.pgdg16.04+1)

BEGIN;

-- kv table, for key-value store

CREATE TABLE IF NOT EXISTS public.kv (
    key character varying NOT NULL,
    value character varying,
    "timestamp" timestamp without time zone NOT NULL
);

ALTER TABLE ONLY public.kv DROP CONSTRAINT IF EXISTS kv_pbkey;
ALTER TABLE ONLY public.kv
    ADD CONSTRAINT kv_pkey PRIMARY KEY (key);

-- meme_names

CREATE TABLE IF NOT EXISTS public.meme_names
(
    id integer NOT NULL,
    name character varying,
    "timestamp" timestamp without time zone NOT NULL,
    author character varying NOT NULL,
    meme_id integer
);

ALTER TABLE ONLY public.meme_names DROP CONSTRAINT IF EXISTS meme_names_name_key;
ALTER TABLE ONLY public.meme_names
    ADD CONSTRAINT meme_names_name_key UNIQUE (name);

ALTER TABLE ONLY public.meme_names DROP CONSTRAINT IF EXISTS meme_names_pkey;
ALTER TABLE ONLY public.meme_names
    ADD CONSTRAINT meme_names_pkey PRIMARY KEY (id);

ALTER TABLE ONLY public.meme_names DROP CONSTRAINT IF EXISTS meme_names_meme_id_fkey;
ALTER TABLE ONLY public.meme_names
    ADD CONSTRAINT meme_names_meme_id_fkey FOREIGN KEY (meme_id) REFERENCES public.memes(id);

CREATE SEQUENCE IF NOT EXISTS public.meme_names_id_seq
    AS integer
    START WITH 1
    INCREMENT BY 1
    NO MINVALUE
    NO MAXVALUE
    CACHE 1;

ALTER TABLE ONLY public.meme_names ALTER COLUMN id SET DEFAULT nextval('public.meme_names_id_seq'::regclass);

-- meme_urls

CREATE TABLE IF NOT EXISTS public.meme_urls
(
    id integer NOT NULL,
    url character varying NOT NULL,
    "timestamp" timestamp
    without time zone NOT NULL,
    author character varying NOT NULL,
    meme_id integer
);

ALTER TABLE ONLY public.meme_urls DROP CONSTRAINT IF EXISTS meme_urls_pkey;
ALTER TABLE ONLY public.meme_urls
    ADD CONSTRAINT meme_urls_pkey PRIMARY KEY (id);

CREATE SEQUENCE IF NOT EXISTS public.meme_urls_id_seq
    AS integer
    START WITH 1
    INCREMENT BY 1
    NO MINVALUE
    NO MAXVALUE
    CACHE 1;

ALTER TABLE ONLY public.meme_urls ALTER COLUMN id SET DEFAULT nextval('public.meme_urls_id_seq'::regclass);
ALTER TABLE ONLY public.meme_urls DROP CONSTRAINT IF EXISTS meme_urls_meme_id_fkey;
ALTER TABLE ONLY public.meme_urls
    ADD CONSTRAINT meme_urls_meme_id_fkey FOREIGN KEY (meme_id) REFERENCES public.memes(id);

-- memes

CREATE TABLE IF NOT EXISTS public.memes (
    id integer NOT NULL
);

ALTER TABLE ONLY public.memes DROP CONSTRAINT IF EXISTS memes_pkey;
ALTER TABLE ONLY public.memes
    ADD CONSTRAINT memes_pkey PRIMARY KEY (id);

CREATE SEQUENCE IF NOT EXISTS public.memes_id_seq
    AS integer
    START WITH 1
    INCREMENT BY 1
    NO MINVALUE
    NO MAXVALUE
    CACHE 1;

ALTER TABLE ONLY public.memes ALTER COLUMN id SET DEFAULT nextval('public.memes_id_seq'::regclass);

-- role_acl table, for ACL features

CREATE TABLE IF NOT EXISTS public.role_acl (
    row_id integer NOT NULL,
    acl_id character varying NOT NULL,
    role_id character varying NOT NULL,
    details character varying
);

ALTER TABLE ONLY public.role_acl DROP CONSTRAINT IF EXISTS role_acl_pkey;
ALTER TABLE ONLY public.role_acl
    ADD CONSTRAINT role_acl_pkey PRIMARY KEY (row_id);

CREATE SEQUENCE IF NOT EXISTS public.seq_role_acl_row_id
    START WITH 1
    INCREMENT BY 1
    NO MINVALUE
    NO MAXVALUE
    CACHE 1;

CREATE INDEX IF NOT EXISTS idx_role_acl__acl_group ON public.role_acl USING btree (acl_id, role_id);

-- user_acl table, for ACL features

CREATE TABLE IF NOT EXISTS public.user_acl (
    row_id integer NOT NULL,
    acl_id character varying NOT NULL,
    user_id character varying NOT NULL,
    details character varying
);

ALTER TABLE ONLY public.user_acl DROP CONSTRAINT IF EXISTS user_acl_pkey;
ALTER TABLE ONLY public.user_acl
    ADD CONSTRAINT user_acl_pkey PRIMARY KEY (row_id);

CREATE SEQUENCE IF NOT EXISTS public.seq_user_acl_row_id
    START WITH 1
    INCREMENT BY 1
    NO MINVALUE
    NO MAXVALUE
    CACHE 1;

CREATE INDEX IF NOT EXISTS idx_user_acl__acl_user ON public.user_acl USING btree (acl_id, user_id);

COMMIT;