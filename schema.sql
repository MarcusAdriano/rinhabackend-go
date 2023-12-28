CREATE TABLE pessoas (
    id UUID PRIMARY KEY,
    nome VARCHAR(100) NOT NULL,
    apelido VARCHAR(32) NOT NULL UNIQUE,
    stack TEXT,
    nascimento DATE NOT NULL
);