-- Dumped from database version 10.11 (Ubuntu 10.11-1.pgdg16.04+1)

BEGIN;

-- kv table, for key-value store

CREATE TABLE IF NOT EXISTS public.kv (
    key character varying NOT NULL,
    value character varying,
    "timestamp" timestamp without time zone NOT NULL
);

-- meme_names

CREATE TABLE IF NOT EXISTS public.meme_names
(
    id integer NOT NULL,
    name character varying,
    "timestamp" timestamp without time zone NOT NULL,
    author character varying NOT NULL,
    meme_id integer
);

CREATE SEQUENCE IF NOT EXISTS public.meme_names_id_seq
    -- AS integer
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


CREATE SEQUENCE IF NOT EXISTS public.meme_urls_id_seq
    -- AS integer
    START WITH 1
    INCREMENT BY 1
    NO MINVALUE
    NO MAXVALUE
    CACHE 1;

ALTER TABLE ONLY public.meme_urls ALTER COLUMN id SET DEFAULT nextval('public.meme_urls_id_seq'::regclass);

-- memes

CREATE TABLE IF NOT EXISTS public.memes (
    id integer NOT NULL
);

CREATE SEQUENCE IF NOT EXISTS public.memes_id_seq
    -- AS integer
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

CREATE SEQUENCE IF NOT EXISTS public.seq_role_acl_row_id
    START WITH 1
    INCREMENT BY 1
    NO MINVALUE
    NO MAXVALUE
    CACHE 1;

-- user_acl table, for ACL features

CREATE TABLE IF NOT EXISTS public.user_acl (
    row_id integer NOT NULL,
    acl_id character varying NOT NULL,
    user_id character varying NOT NULL,
    details character varying
);

CREATE SEQUENCE IF NOT EXISTS public.seq_user_acl_row_id
    START WITH 1
    INCREMENT BY 1
    NO MINVALUE
    NO MAXVALUE
    CACHE 1;

-- constraints

ALTER TABLE ONLY public.kv DROP CONSTRAINT IF EXISTS kv_pkey CASCADE;
ALTER TABLE ONLY public.kv
    ADD CONSTRAINT kv_pkey PRIMARY KEY (key);

ALTER TABLE ONLY public.meme_names DROP CONSTRAINT IF EXISTS meme_names_name_key CASCADE;
ALTER TABLE ONLY public.meme_names
    ADD CONSTRAINT meme_names_name_key UNIQUE (name);

ALTER TABLE ONLY public.meme_names DROP CONSTRAINT IF EXISTS meme_names_pkey CASCADE;
ALTER TABLE ONLY public.meme_names
    ADD CONSTRAINT meme_names_pkey PRIMARY KEY (id);

ALTER TABLE ONLY public.meme_urls DROP CONSTRAINT IF EXISTS meme_urls_pkey CASCADE;
ALTER TABLE ONLY public.meme_urls
    ADD CONSTRAINT meme_urls_pkey PRIMARY KEY (id);

ALTER TABLE ONLY public.memes DROP CONSTRAINT IF EXISTS memes_pkey CASCADE;
ALTER TABLE ONLY public.memes
    ADD CONSTRAINT memes_pkey PRIMARY KEY (id);

ALTER TABLE ONLY public.role_acl DROP CONSTRAINT IF EXISTS role_acl_pkey CASCADE;
ALTER TABLE ONLY public.role_acl
    ADD CONSTRAINT role_acl_pkey PRIMARY KEY (row_id);

ALTER TABLE ONLY public.user_acl DROP CONSTRAINT IF EXISTS user_acl_pkey CASCADE;
ALTER TABLE ONLY public.user_acl
    ADD CONSTRAINT user_acl_pkey PRIMARY KEY (row_id);

-- foreign constraints

ALTER TABLE ONLY public.meme_names DROP CONSTRAINT IF EXISTS meme_names_meme_id_fkey CASCADE;
ALTER TABLE ONLY public.meme_names
    ADD CONSTRAINT meme_names_meme_id_fkey FOREIGN KEY (meme_id) REFERENCES public.memes(id);

ALTER TABLE ONLY public.meme_urls DROP CONSTRAINT IF EXISTS meme_urls_meme_id_fkey CASCADE;
ALTER TABLE ONLY public.meme_urls
    ADD CONSTRAINT meme_urls_meme_id_fkey FOREIGN KEY (meme_id) REFERENCES public.memes(id);

-- indices

CREATE INDEX IF NOT EXISTS idx_role_acl__acl_group ON public.role_acl USING btree (acl_id, role_id);

CREATE INDEX IF NOT EXISTS idx_user_acl__acl_user ON public.user_acl USING btree (acl_id, user_id);

COMMIT;