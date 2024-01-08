CREATE TABLE pessoas (
    id UUID NOT NULL UNIQUE,
    nome VARCHAR(100) NOT NULL,
    apelido VARCHAR(32) NOT NULL UNIQUE,
    stack TEXT,
    nascimento DATE NOT NULL
);