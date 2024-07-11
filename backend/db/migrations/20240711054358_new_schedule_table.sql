-- +goose Up
-- +goose StatementBegin
CREATE TABLE IF NOT EXISTS public.schedule
(
    id uuid NOT NULL,
    quiz_id uuid NOT NULL,
    start_time integer DEFAULT 0,
    end_time integer DEFAULT 0,
    number_of_attempts integer DEFAULT 0,
    created_at timestamp with time zone DEFAULT now(),
    CONSTRAINT schedule_pkey PRIMARY KEY (id),
    CONSTRAINT schedule_quiz_id_fkey FOREIGN KEY (quiz_id)
        REFERENCES public.quiz (id) MATCH SIMPLE
        ON UPDATE CASCADE
        ON DELETE CASCADE
)
TABLESPACE pg_default;

ALTER TABLE IF EXISTS public.schedule
    OWNER to postgres;
-- +goose StatementEnd

-- +goose Down
-- +goose StatementBegin
DROP TABLE IF EXISTS public.schedule;
-- +goose StatementEnd
