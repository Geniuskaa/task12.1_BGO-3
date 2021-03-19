CREATE TABLE clients
(
    id BIGSERIAL PRIMARY KEY,
    login TEXT   NOT NULL UNIQUE,
    password TEXT NOT NULL,
    full_name TEXT NOT NULL,
    passport TEXT NOT NULL,
    birthday DATE NOT NULL,
    status TEXT NOT NULL DEFAULT 'INACTIVE' CHECK (status IN ('INACTIVE', 'ACTIVE')),
    created TIMESTAMP NOT NULL DEFAULT CURRENT_TIMESTAMP
);

CREATE TABLE cards
(
    id BIGSERIAL PRIMARY KEY,
    number TEXT NOT NULL,
    balance BIGINT NOT NULL DEFAULT 0,
    issuer TEXT NOT NULL CHECK ( issuer IN ('VISA', 'MasterCard', 'MIR') ),
    holder TEXT NOT NULL,
    owner_id BIGINT NOT NULL REFERENCES clients,
    status TEXT NOT NULL DEFAULT 'INACTIVE' CHECK (status IN ('INACTIVE', 'ACTIVE')),
    created TIMESTAMP NOT NULL DEFAULT CURRENT_TIMESTAMP
);

CREATE TABLE icons_for_transactions
(
    mcc BIGINT NOT NULL UNIQUE PRIMARY KEY,
    icon TEXT NOT NULL UNIQUE,
    created TIMESTAMP NOT NULL DEFAULT current_timestamp
);

CREATE TABLE transactions
(
    id BIGSERIAL NOT NULL PRIMARY KEY,
    card_id BIGINT NOT NULL REFERENCES cards,
    sum BIGINT NOT NULL,
    mcc BIGINT NOT NULL REFERENCES icons_for_transactions,
    receiver TEXT DEFAULT 'CardHolder',
    created TIMESTAMP NOT NULL DEFAULT CURRENT_TIMESTAMP
);

