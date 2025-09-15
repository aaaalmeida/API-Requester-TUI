INSERT INTO collection (name, description) VALUES
('collection 1', 'description 1'),
('collection 2', 'description 2'),
('collection 3', 'description 3'),
('faculdade', 'desc faculdade'),
('trabalho', 'desc trabalho'),
('pessoal', 'desc pessoal');

INSERT INTO request (name,                  url,                                            method_id,  collection_id,  headers,    body) VALUES
                    ('ola mundo',           'http://localhost:8080/',                       1,          1,              null,       null),
                    ('time',                'http://localhost:8080/time',                   1,          1,              null,       null),
                    ('example',             'https://www.example.com',                      1,          6,              null,       null),
                    ('json placeholder',    'https://jsonplaceholder.typicode.com/posts/1', 1,          2,              null,       null),
                    ('json placeholder',    'https://jsonplaceholder.typicode.com/posts',   2,          2,              null,       null),
                    ('json placeholder',    'https://jsonplaceholder.typicode.com/posts',   3,          2,              null,       null);


INSERT INTO request (name,                          url,                        method_id,  collection_id,  body,                               bodyType, headers) VALUES
                    ('retorna body',                'http://localhost:8080/',   1,          1,              '{"hello": "world", "num": 123}',   2,        null),
                    ('retorna body com header',     'http://localhost:8080/',   1,          1,              '{"hello": "world", "num": 123}',   2,        '{"Content-type":"application/json","Teste":"ABCDE"}'),
