INSERT INTO collection (name, description) VALUES 
('collection 1', 'description 1'),
('collection 2', 'description 2'),
('collection 3', 'description 3'),
('faculdade', 'desc faculdade'),
('trabalho', 'desc trabalho'),
('pessoal', 'desc pessoal');

INSERT INTO request (name, url, method_id, collection_id) VALUES
('request 1', 'localhost:8080/ola', 1, 1),
('request 2', 'localhost:8080/cliente', 1, 1),
('usuarios', 'localhost:8080/usuario', 1, 1),
('usuario teste', 'localhost:8080/usuario/teste', 1, 1),
('json placeholder', 'https://jsonplaceholder.typicode.com/posts/1', 1, 2),
('POST json placeholder', 'https://jsonplaceholder.typicode.com/posts', 2, 2),
('PUT json placeholder', 'https://jsonplaceholder.typicode.com/posts', 3, 2),
('example', 'https://www.example.com', 1, 6);