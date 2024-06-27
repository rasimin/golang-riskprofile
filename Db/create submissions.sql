-- Table: public.submissions

-- DROP TABLE IF EXISTS public.submissions;

CREATE TABLE IF NOT EXISTS public.submissions
(
    id integer NOT NULL DEFAULT nextval('submissions_id_seq'::regclass),
    user_id integer NOT NULL,
    answers json NOT NULL,
    risk_score integer,
    risk_category character varying(50) COLLATE pg_catalog."default",
    created_at timestamp without time zone NOT NULL DEFAULT CURRENT_TIMESTAMP,
    updated_at timestamp without time zone NOT NULL DEFAULT CURRENT_TIMESTAMP,
    CONSTRAINT submissions_pkey PRIMARY KEY (id),
    CONSTRAINT submissions_user_id_fkey FOREIGN KEY (user_id)
        REFERENCES public.users (id) MATCH SIMPLE
        ON UPDATE NO ACTION
        ON DELETE NO ACTION
)

TABLESPACE pg_default;

ALTER TABLE IF EXISTS public.submissions
    OWNER to postgres;