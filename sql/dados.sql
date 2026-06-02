insert into usuarios (nome, nick, email, senha) 
values 
("Usuário 1", "usuario1", "usuario1@example.com", "$2a$10$TyvOcxChQIrIyYuff9mn9ekT6Wj8WLfaufMPunRcbofH2nkh2Ff5."),
("Usuário 2", "usuario2", "usuario2@example.com", "$2a$10$TyvOcxChQIrIyYuff9mn9ekT6Wj8WLfaufMPunRcbofH2nkh2Ff5."),
("Usuário 3", "usuario3", "usuario3@example.com", "$2a$10$TyvOcxChQIrIyYuff9mn9ekT6Wj8WLfaufMPunRcbofH2nkh2Ff5.");

insert into seguidores (usuario_id, seguidor_id) 
values 
(1, 2),
(1, 3),
(2, 1),
(2, 3),
(3, 1),
(3, 2);