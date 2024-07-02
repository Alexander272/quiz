-- +goose Up
-- +goose StatementBegin
CREATE TABLE IF NOT EXISTS public.question
(
    id uuid NOT NULL,
    quiz_id uuid NOT NULL,
    number integer NOT NULL,
    text text COLLATE pg_catalog."default" NOT NULL,
    description text COLLATE pg_catalog."default" DEFAULT ''::text,
    image text COLLATE pg_catalog."default" DEFAULT ''::text,
    has_shuffle boolean DEFAULT false,
    level integer DEFAULT 0,
    points integer DEFAULT 1,
    time integer NOT NULL,
    group_id uuid DEFAULT '00000000-0000-0000-0000-000000000000'::uuid,
    created_at timestamp with time zone DEFAULT now(),
    CONSTRAINT question_pkey PRIMARY KEY (id),
    CONSTRAINT question_quiz_id_fkey FOREIGN KEY (quiz_id)
        REFERENCES public.quiz (id) MATCH SIMPLE
        ON UPDATE CASCADE
        ON DELETE CASCADE
)
TABLESPACE pg_default;

ALTER TABLE IF EXISTS public.question
    OWNER to postgres;
-- +goose StatementEnd

-- +goose Down
-- +goose StatementBegin
DROP TABLE IF EXISTS public.question;
-- +goose StatementEnd
