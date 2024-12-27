-- +goose Up
-- +goose StatementBegin
CREATE TABLE IF NOT EXISTS public.quiz
(
    id uuid NOT NULL,
    title text COLLATE pg_catalog."default" NOT NULL,
    description text COLLATE pg_catalog."default" DEFAULT ''::text,
    image text COLLATE pg_catalog."default" DEFAULT ''::text,
    is_drawing boolean DEFAULT true,
    number_of_attempts integer DEFAULT 0,
    time integer NOT NULL,
    author_id uuid NOT NULL,
    category_id uuid DEFAULT '00000000-0000-0000-0000-000000000000'::uuid,
    start_time integer DEFAULT 0,
    end_time integer DEFAULT 0,
    has_shuffle boolean DEFAULT false,
    has_skippable boolean DEFAULT false,
    show_list boolean DEFAULT false,
    show_answers boolean DEFAULT false,
    show_results boolean DEFAULT false,
    created_at timestamp with time zone DEFAULT now(),
    CONSTRAINT quiz_pkey PRIMARY KEY (id)
)
TABLESPACE pg_default;

ALTER TABLE IF EXISTS public.quiz
    OWNER to postgres;
-- +goose StatementEnd

-- +goose Down
-- +goose StatementBegin
DROP TABLE IF EXISTS public.quiz;
-- +goose StatementEnd
