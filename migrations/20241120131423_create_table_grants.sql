-- +goose Up
-- +goose StatementBegin
CREATE TABLE grants (
    id SERIAL PRIMARY KEY,
    title text,
    source_url text,
    project_directions jsonb,
    amount integer,
    legal_forms jsonb,
    age integer,
    cutting_off_criterea jsonb
);
-- +goose StatementEnd

-- +goose Down
-- +goose StatementBegin
DROP TABLE grants;
-- +goose StatementEnd
