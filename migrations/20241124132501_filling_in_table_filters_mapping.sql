-- +goose Up
-- +goose StatementBegin
INSERT INTO filters_mapping (age, project_direction, legal_form, cutting_off_criteria, amount) VALUES 
('{"title": "Возраст участников", "mapping": {}}',
 '{"title": "Направление проекта", "mapping": {"0": {"title": "Не указано"}, "1": {"title": "Выявление и поддержка молодых талантов"}, "2": {"title": "Защита прав и свобод"}, "3": {"title": "Охрана здоровья"}}}', 
 '{"title": "Отсекающие критерии", "mapping": {"0": {"title": "Не указано"}, "1": {"title": "Юр. лицо"}, "2": {"title": "Физ. лицо"}}}', 
 '{"title": "Отсекающие критерии", "mapping": {"0": {"title": "Не указано"}, "1": {"title": "Для школьников"}, "2": {"title": "Для студентов"}, "3": {"title": "Для асприантов"}}}', 
 '{"title": "Сумма", "mapping": {}}')
-- +goose StatementEnd

-- +goose Down
-- +goose StatementBegin
DELETE FROM filters_mapping;
-- +goose StatementEnd
