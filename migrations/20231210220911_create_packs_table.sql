-- +goose Up
-- +goose StatementBegin
CREATE TABLE public.pack (
    id UUID PRIMARY KEY DEFAULT gen_random_uuid(),
    amount integer NOT NULL UNIQUE,
    created_at timestamp DEFAULT now(),
    updated_at timestamp
);

INSERT INTO public.pack(amount) VALUES 
    (250), 
    (500), 
    (1000),
    (2000), 
    (5000)
;

-- +goose StatementEnd

-- +goose Down
-- +goose StatementBegin
DROP TABLE IF EXISTS public.pack;

-- +goose StatementEnd
