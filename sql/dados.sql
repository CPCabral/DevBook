insert into usuarios (nome, nick, email, senha) values
("Usuario_1", "usuario_1", "usuario_1@email.com", "$2a$10$TRIribXX5oa8s5YFUhphZuMtX2jwTG1B31Aw9uWCNNa.nNzJ90tIW"),
("Usuario_2", "usuario_2", "usuario_2@email.com", "$2a$10$TRIribXX5oa8s5YFUhphZuMtX2jwTG1B31Aw9uWCNNa.nNzJ90tIW"),
("Usuario_3", "usuario_3", "usuario_3@email.com", "$2a$10$TRIribXX5oa8s5YFUhphZuMtX2jwTG1B31Aw9uWCNNa.nNzJ90tIW");

insert into seguidores (usuario_id, seguidor_id) values
(1, 2),
(1, 3),
(2, 1),
(2, 3),
(3, 1),
(3, 2);