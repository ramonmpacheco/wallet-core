use wallet;

-- CLIENTS
CREATE TABLE IF NOT EXISTS clients(
    id varchar(255), 
    name varchar(255), 
    email varchar(255), 
    created_at date
);

REPLACE INTO clients 
    (id, name, email, created_at)
VALUES 
    ('4a2eb0f6-43ac-4620-9085-e66412e97836', 'John Doe', 'john@j.com', '2023-04-30'),
    ('a8750f17-e8f9-4bbd-a6df-ea7694897ee4', 'Jane Doe', 'jane@j.com', '2023-04-30');
  
-- ACCOUNTS
CREATE TABLE IF NOT EXISTS accounts(
    id varchar(255), 
    client_id varchar(255), 
    balance int, 
    created_at date
);

REPLACE INTO 
    accounts (id, client_id, balance, created_at)
VALUES
    ('8b253de7-99a5-491d-8f5c-9d1242155cdc', '4a2eb0f6-43ac-4620-9085-e66412e97836', 1000, '2023-04-30'),
    ('99b23063-6b7e-42b0-ba5b-9af20250881f', 'a8750f17-e8f9-4bbd-a6df-ea7694897ee4', 0, '2023-04-30');


-- TRANSACTIONS
CREATE TABLE IF NOT EXISTS transactions (
    id varchar(255), 
    account_id_from varchar(255), 
    account_id_to varchar(255), 
    amount int, 
    created_at date
);
