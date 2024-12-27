-- +goose Up
-- +goose StatementBegin
CREATE TABLE IF NOT EXISTS public.attempt_details
(
    id uuid NOT NULL,
    attempt_id uuid NOT NULL,
    question_id uuid NOT NULL,
    answers uuid[],
    created_at timestamp with time zone DEFAULT now(),
    CONSTRAINT attempt_details_pkey PRIMARY KEY (id),
    CONSTRAINT attempt_details_attempt_id_fkey FOREIGN KEY (attempt_id)
        REFERENCES public.attempt (id) MATCH SIMPLE
        ON UPDATE CASCADE
        ON DELETE CASCADE,
    CONSTRAINT attempt_details_question_id_fkey FOREIGN KEY (question_id)
        REFERENCES public.question (id) MATCH SIMPLE
        ON UPDATE NO ACTION
        ON DELETE NO ACTION
)
TABLESPACE pg_default;

ALTER TABLE IF EXISTS public.attempt_details
    OWNER to postgres;
-- +goose StatementEnd

-- +goose Down
-- +goose StatementBegin
DROP TABLE IF EXISTS public.attempt_details;
-- +goose StatementEnd
