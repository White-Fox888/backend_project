-- +goose Up
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
-- +goose StatementBegin
SELECT 'up SQL query';
-- +goose StatementEnd

-- +goose Down
DROP TABLE grants;
-- +goose StatementBegin
SELECT 'down SQL query';
-- +goose StatementEnd
