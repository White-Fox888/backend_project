-- +goose Up
INSERT INTO grants (id, title, source_url, project_directions, amount, legal_forms, age, cutting_off_criterea) VALUES 
(1, 'Стипендиальный конкурс' 'https://fondpotanin.ru/competitions/fellowships/', '[]', 0, '[]', 0, '[]');
INSERT INTO grants (id, title, source_url, project_directions, amount, legal_forms, age, cutting_off_criterea) VALUES 
(2, 'Спорт для всех', 'https://fondpotanin.ru/competitions/sport-dlya-vsekh/', '[]', 0, '[]', 0, '[]');
INSERT INTO grants (id, title, source_url, project_directions, amount, legal_forms, age, cutting_off_criterea) VALUES 
(3, 'Музей 4.0', 'https://fondpotanin.ru/competitions/muzey-4-0/', '[]', 0, '[]', 0, '[]');
INSERT INTO grants (id, title, source_url, project_directions, amount, legal_forms, age, cutting_off_criterea) VALUES 
(4, 'Грантовый конкурс для преподавателей магистратуры', 'https://fondpotanin.ru/competitions/professors-grants/', '[]', 0, '[]', 0, '[]');
INSERT INTO grants (id, title, source_url, project_directions, amount, legal_forms, age, cutting_off_criterea) VALUES 
(5, 'Индустриальный эксперимент', 'https://fondpotanin.ru/competitions/industrial/', '[]', 0, '[]', 0, '[]');
INSERT INTO grants (id, title, source_url, project_directions, amount, legal_forms, age, cutting_off_criterea) VALUES 
(6, 'Креативный музей', 'https://fondpotanin.ru/competitions/kreativnyy-muzey/', '[]', 0, '[]', 0, '[]');
INSERT INTO grants (id, title, source_url, project_directions, amount, legal_forms, age, cutting_off_criterea) VALUES 
(7, 'Профессиональное развитие', 'https://fondpotanin.ru/competitions/konkurs-professionalnogo-razvitiya/', '[]', 0, '[]', 0, '[]');
INSERT INTO grants (id, title, source_url, project_directions, amount, legal_forms, age, cutting_off_criterea) VALUES 
(8, 'Практики личной филантропии и альтруизма', 'https://fondpotanin.ru/competitions/individual-giving-and-volunteering/', '[]', 0, '[]', 0, '[]');
INSERT INTO grants (id, title, source_url, project_directions, amount, legal_forms, age, cutting_off_criterea) VALUES 
(9, '#фондпотанина25', 'https://fondpotanin.ru/competitions/fondpotanina25/', '[]', 0, '[]', 0, '[]');
INSERT INTO grants (id, title, source_url, project_directions, amount, legal_forms, age, cutting_off_criterea) VALUES 
(10, 'Школа Фонда', 'https://fondpotanin.ru/competitions/shkola-fonda/', '[]', 0, '[]', 0, '[]');
INSERT INTO grants (id, title, source_url, project_directions, amount, legal_forms, age, cutting_off_criterea) VALUES 
(11, 'Программа стажировки студентов МГИМО (совместно с МИД РФ)', 'https://fondpotanin.ru/competitions/programma-stazhirovki-studentov-mgimo-sovmestno-s-mid-rf/', '[]', 0, '[]', 0, '[]');
INSERT INTO grants (id, title, source_url, project_directions, amount, legal_forms, age, cutting_off_criterea) VALUES 
(12, 'Центры знаний по социальной поддержке', 'https://fondpotanin.ru/competitions/socialsupport/', '[]', 0, '[]', 0, '[]');
INSERT INTO grants (id, title, source_url, project_directions, amount, legal_forms, age, cutting_off_criterea) VALUES 
(13, 'Российско-европейская программа обмена Philanthropic Leadership Platform: Russia-Europe', 'https://fondpotanin.ru/competitions/rossiysko-evropeyskaya-programma-obmena-philanthropic-leadership-platform-russia-europe/', '[]', 0, '[]', 0, '[]');
INSERT INTO grants (id, title, source_url, project_directions, amount, legal_forms, age, cutting_off_criterea) VALUES 
(14, 'Социальные финансы', 'https://fondpotanin.ru/competitions/oxfordsocialfinance/', '[]', 0, '[]', 0, '[]');
INSERT INTO grants (id, title, source_url, project_directions, amount, legal_forms, age, cutting_off_criterea) VALUES 
(15, 'Олимпийские стипендии', 'https://fondpotanin.ru/competitions/olimpiyskie-stipendii/', '[]', 0, '[]', 0, '[]');
INSERT INTO grants (id, title, source_url, project_directions, amount, legal_forms, age, cutting_off_criterea) VALUES 
(16, 'Школа филантропии', 'https://fondpotanin.ru/competitions/shkola-filantropii/', '[]', 0, '[]', 0, '[]');
INSERT INTO grants (id, title, source_url, project_directions, amount, legal_forms, age, cutting_off_criterea) VALUES 
(17, 'Инициатива «Музей. Сила места»', 'https://fondpotanin.ru/competitions/muzey-sila-mesta/', '[]', 0, '[]', 0, '[]');
INSERT INTO grants (id, title, source_url, project_directions, amount, legal_forms, age, cutting_off_criterea) VALUES 
(18, 'Точка опоры', 'https://fondpotanin.ru/competitions/tochka-opory/', '[]', 0, '[]', 0, '[]');
INSERT INTO grants (id, title, source_url, project_directions, amount, legal_forms, age, cutting_off_criterea) VALUES 
(19, 'Конкурс на поддержку центров социальных инноваций в сфере культуры', 'https://fondpotanin.ru/competitions/konkurs-na-podderzhku-tsentrov-sotsialnykh-innovatsiy-v-sfere-kultury/', '[]', 0, '[]', 0, '[]');
INSERT INTO grants (id, title, source_url, project_directions, amount, legal_forms, age, cutting_off_criterea) VALUES 
(20, 'Конкурс по приглашению для грантополучателей антикризисных конкурсов', 'https://fondpotanin.ru/competitions/konkurs-po-priglasheniyu-dlya-grantopoluchateley-antikrizisnykh-konkursov/', '[]', 0, '[]', 0, '[]');
-- +goose StatementBegin
SELECT 'up SQL query';
-- +goose StatementEnd

-- +goose Down
DELETE FROM grants;
-- +goose StatementBegin
SELECT 'down SQL query';
-- +goose StatementEnd
