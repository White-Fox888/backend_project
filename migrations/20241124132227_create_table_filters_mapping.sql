-- +goose Up
-- +goose StatementBegin
CREATE TABLE filters_mapping (
    age jsonb,
    project_direction jsonb,
    legal_form jsonb,
    cutting_off_criteria jsonb,
    amount jsonb
);
-- +goose StatementEnd

-- +goose Down
-- +goose StatementBegin
DROP TABLE filters_mapping;
-- +goose StatementEnd
