-- Inserir valores na tabela 'usuarios'
INSERT INTO usuarios (nome, nick, email, senha)
VALUES
('Usuário 1', 'usuario_1', 'usuario1@gmail.com', '$2a$10$0iGYlKCAYTyJV/vC6nLGgeWFwD6AhSkWLsVRO/.M4lNK8OtIkfggy'), -- usuario1
('Usuário 2', 'usuario_2', 'usuario2@gmail.com', '$2a$10$0iGYlKCAYTyJV/vC6nLGgeWFwD6AhSkWLsVRO/.M4lNK8OtIkfggy'), -- usuario2
('Usuário 3', 'usuario_3', 'usuario3@gmail.com', '$2a$10$0iGYlKCAYTyJV/vC6nLGgeWFwD6AhSkWLsVRO/.M4lNK8OtIkfggy'); -- usuario3

-- Inserir valores na tabela 'seguidores'
INSERT INTO seguidores (usuario_id, seguidor_id)
VALUES
(1, 2),
(3, 1),
(1, 3);

-- Inserir valores na tabela 'publicacoes'
INSERT INTO publicacoes (titulo, conteudo, autor_id)
VALUES
('Publicação do Usuário 1', 'Essa é a publicação do usuário 1! Oba!', 1),
('Publicação do Usuário 2', 'Essa é a publicação do usuário 2! Oba!', 2),
('Publicação do Usuário 3', 'Essa é a publicação do usuário 3! Oba!', 3);
