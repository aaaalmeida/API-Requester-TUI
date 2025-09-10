INSERT INTO collection (name, description) VALUES
('collection 1', 'description 1'),
('collection 2', 'description 2'),
('collection 3', 'description 3'),
('faculdade', 'desc faculdade'),
('trabalho', 'desc trabalho'),
('pessoal', 'desc pessoal');

INSERT INTO request (name, url, method_id, collection_id, body) VALUES
('ola mundo', 'localhost:8080/', 1, 1, null),
('time', 'localhost:8080/time', 1, 1, null),
('retorna json', 'localhost:8080/', 2, 1, '"{\n\"string\":\"abcd\",\n\"num\":123,\n\"arr\": [\n{\n\"val1\":\"val1\",\n\"num1\":1\n},\n{\n\"val2\":\"val2\",\n\"bool\":true\n}\n]\n}"'),
('retorna parte do json', 'localhost:8080/obj', 2, 1, '{"string": "abcd","valor": [{"num1": 1,"num2": 2},{"bool1": true,"bool2": false},{"nome": "almeida","info": {  "idade": 24,  "sexo": "M"}}]}')
('json placeholder', 'https://jsonplaceholder.typicode.com/posts/1', 1, 2),
('json placeholder', 'https://jsonplaceholder.typicode.com/posts', 2, 2),
('json placeholder', 'https://jsonplaceholder.typicode.com/posts', 3, 2),
('teste', 'localhost:8080/cliente', 1, 3),
('teste', 'localhost:8080/cliente', 2, 3),
('teste', 'localhost:8080/cliente', 3, 3),
('teste', 'localhost:8080/cliente', 4, 3),
('teste', 'localhost:8080/cliente', 5, 3),
('teste', 'localhost:8080/cliente', 6, 3),
('teste', 'localhost:8080/cliente', 7, 3),
('teste', 'localhost:8080/cliente', 8, 3),
('example', 'https://www.example.com', 1, 6),
('nome muito grande para testar o tamanho da caixa de requisições', 'localhost:8080/cliente', 1, 6);
