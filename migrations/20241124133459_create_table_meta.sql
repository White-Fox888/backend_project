-- +goose Up
-- +goose StatementBegin
CREATE TABLE meta (
    current_page integer,
    total_pages integer
);
-- +goose StatementEnd

-- +goose Down
-- +goose StatementBegin
DROP TABLE meta;
-- +goose StatementEnd
