SET NAMES 'utf8mb4';
SET CHARACTER SET 'utf8mb4';

CREATE TABLE videos (
    id INT AUTO_INCREMENT PRIMARY KEY,
    title VARCHAR(255) NOT NULL,
    description TEXT,
    release_year YEAR NOT NULL,
    director VARCHAR(255),
    stars VARCHAR(255),  -- Atores principais, separados por vírgula
    genre VARCHAR(100),  -- Gênero do vídeo
    thumbnail_url VARCHAR(255),  -- URL da imagem em miniatura
    url VARCHAR(255) NOT NULL,  -- URL do vídeo
    duration INT NOT NULL,  -- Duração em minutos
    rating DECIMAL(3, 2),  -- Avaliação (por exemplo, 8.75)
    created_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP,  -- Data de criação
    updated_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP  -- Data da última atualização
);

INSERT INTO videos (title, description, release_year, director, stars, genre, thumbnail_url, url, duration, rating) VALUES
('O Mistério da Casa Abandonada', 'Um grupo de amigos descobre segredos sombrios em uma casa abandonada.', 2022, 'Luiz Fernando Carvalho', 'João Silva, Maria Oliveira', 'Suspense, Mistério', 'http://url-da-thumb.com/thumb1.jpg', 'http://url-do-video.com/video1.mp4', 120, 7.2),
('Caminhos da Aventura', 'Uma jornada emocionante através de paisagens desconhecidas.', 2020, 'Ana Paula Mendes', 'Carlos Lima, Fernanda Costa', 'Aventura', 'http://url-da-thumb.com/thumb2.jpg', 'http://url-do-video.com/video2.mp4', 95, 8.1),
('Amores e Desamores', 'Histórias de amor em tempos difíceis.', 2021, 'Pedro Almodóvar', 'Roberto Santos, Juliana Ribeiro', 'Romance, Drama', 'http://url-da-thumb.com/thumb3.jpg', 'http://url-do-video.com/video3.mp4', 85, 6.5),
('A Última Fronteira', 'Exploradores enfrentam desafios em uma terra inexplorada.', 2019, 'James Cameron', 'Tom Hanks, Nicole Kidman', 'Aventura, Ficção Científica', 'http://url-da-thumb.com/thumb4.jpg', 'http://url-do-video.com/video4.mp4', 150, 8.5),
('Os Guardiões da Galáxia', 'Uma equipe de heróis improváveis se une para salvar o universo.', 2018, 'James Gunn', 'Chris Pratt, Zoe Saldana', 'Ação, Aventura', 'http://url-da-thumb.com/thumb5.jpg', 'http://url-do-video.com/video5.mp4', 122, 8.0),
('Mistério na Neve', 'Uma investigação em um vilarejo gelado revela verdades ocultas.', 2023, 'Alfred Hitchcock', 'Guilherme Almeida, Larissa Santos', 'Suspense, Drama', 'http://url-da-thumb.com/thumb6.jpg', 'http://url-do-video.com/video6.mp4', 110, 7.9),
('A Viagem dos Sonhos', 'Um casal decide fazer uma viagem inesperada pelo mundo.', 2021, 'Wes Anderson', 'Leonardo Oliveira, Mariana Lima', 'Comédia, Aventura', 'http://url-da-thumb.com/thumb7.jpg', 'http://url-do-video.com/video7.mp4', 92, 8.3),
('O Último Herói', 'Um herói improvável surge em tempos de crise.', 2022, 'Christopher Nolan', 'Daniel Ferreira, Beatriz Nunes', 'Ação, Drama', 'http://url-da-thumb.com/thumb8.jpg', 'http://url-do-video.com/video8.mp4', 130, 9.0),
('Segredos do Passado', 'Uma jovem descobre segredos familiares que mudam sua vida.', 2020, 'Sofia Coppola', 'Gabriel Ribeiro, Ana Clara', 'Drama', 'http://url-da-thumb.com/thumb9.jpg', 'http://url-do-video.com/video9.mp4', 100, 7.0),
('Entre Mundos', 'Um viajante do tempo tenta corrigir erros do passado.', 2023, 'Steven Spielberg', 'Felipe Martins, Clara Souza', 'Ficção Científica, Aventura', 'http://url-da-thumb.com/thumb10.jpg', 'http://url-do-video.com/video10.mp4', 145, 8.7);