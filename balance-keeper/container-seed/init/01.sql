use balances;

-- CLIENTS
CREATE TABLE IF NOT EXISTS balances(
    id varchar(255),
    account_id varchar(255),
    amount float,
    created_at date,
    updated_at date
);

REPLACE INTO balances 
    (id, account_id, amount, created_at, updated_at)
VALUES 
    ('7cb27100-a783-4cd3-bd6a-12c0be9aa8a8', '8b253de7-99a5-491d-8f5c-9d1242155cdc', 970.0, '2023-05-10', '2023-05-10'),
    ('a6abfd81-ef27-4941-91e7-c68d87bb4976', '99b23063-6b7e-42b0-ba5b-9af20250881f',  30.0, '2023-05-10', '2023-05-10');