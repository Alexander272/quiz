-- +goose Up
-- +goose StatementBegin
CREATE TABLE IF NOT EXISTS public.answer
(
    id uuid NOT NULL,
    question_id uuid NOT NULL,
    number integer NOT NULL,
    text text COLLATE pg_catalog."default" NOT NULL,
    image text COLLATE pg_catalog."default" DEFAULT ''::text,
    is_correct boolean DEFAULT false,
    created_at timestamp with time zone DEFAULT now(),
    CONSTRAINT answer_pkey PRIMARY KEY (id),
    CONSTRAINT answer_question_id_fkey FOREIGN KEY (question_id)
        REFERENCES public.question (id) MATCH SIMPLE
        ON UPDATE CASCADE
        ON DELETE CASCADE
)
TABLESPACE pg_default;

ALTER TABLE IF EXISTS public.answer
    OWNER to postgres;
-- +goose StatementEnd

-- +goose Down
-- +goose StatementBegin
DROP TABLE IF EXISTS public.answer;
-- +goose StatementEnd
