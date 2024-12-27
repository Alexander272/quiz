-- +goose Up
-- +goose StatementBegin
CREATE TABLE IF NOT EXISTS public.attempt
(
    id uuid NOT NULL,
    schedule_id uuid NOT NULL,
    user_id uuid NOT NULL,
    username text COLLATE pg_catalog."default" DEFAULT ''::text,
    start_time integer DEFAULT 0,
    end_time integer DEFAULT 0,
    correct integer DEFAULT 0,
    total integer DEFAULT 0,
    points integer DEFAULT 0,
    total_points integer DEFAULT 0,
    created_at timestamp with time zone DEFAULT now(),
    CONSTRAINT attempt_pkey PRIMARY KEY (id),
    CONSTRAINT attempt_schedule_id_fkey FOREIGN KEY (schedule_id)
        REFERENCES public.schedule (id) MATCH SIMPLE
        ON UPDATE CASCADE
        ON DELETE CASCADE
)
TABLESPACE pg_default;

ALTER TABLE IF EXISTS public.attempt
    OWNER to postgres;
-- +goose StatementEnd

-- +goose Down
-- +goose StatementBegin
DROP TABLE IF EXISTS public.attempt;
-- +goose StatementEnd
