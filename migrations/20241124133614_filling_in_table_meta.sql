-- +goose Up
-- +goose StatementBegin
INSERT INTO meta (current_page, total_pages) VALUES (1, 13);
-- +goose StatementEnd

-- +goose Down
-- +goose StatementBegin
DELETE FROM meta;
-- +goose StatementEnd
