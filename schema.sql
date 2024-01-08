CREATE TABLE pessoas (
    id UUID NOT NULL UNIQUE,
    nome VARCHAR(100) NOT NULL,
    apelido VARCHAR(32) NOT NULL UNIQUE,
    stack TEXT,
    nascimento DATE NOT NULL,
    busca TSVECTOR GENERATED ALWAYS AS (to_tsvector('english', coalesce(stack, '') || ' ' || apelido || ' ' || nome)) STORED
);

CREATE INDEX busca_idx ON pessoas USING GIN (busca);