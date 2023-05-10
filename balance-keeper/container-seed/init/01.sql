use balances;

-- CLIENTS
CREATE TABLE IF NOT EXISTS balances(
    id varchar(255),
    account_id varchar(255),
    amount float,
    created_at date,
    updated_at date
);

-- REPLACE INTO balances 
--     (id, account_id, amount, created_at, updated_at)
-- VALUES 
--     (),
--     ();