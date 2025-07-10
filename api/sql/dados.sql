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

insert into publicacoes (titulo, conteudo, autor_id)
values
("Publicação do usuario 1", "Essa é a publicação do usuario 1", 1),
("Publicação do usuario 2", "Essa é a publicação do usuario 2", 2),
("Publicação do usuario 3", "Essa é a publicação do usuario 3", 3);