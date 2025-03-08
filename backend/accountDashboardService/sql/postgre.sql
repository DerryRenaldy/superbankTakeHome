-- Create schema if needed (optional)
CREATE SCHEMA IF NOT EXISTS account_dashboard;

-- Customers Table
CREATE TABLE account_dashboard.customers (
    customer_id BIGSERIAL PRIMARY KEY,
    name VARCHAR(255) NOT NULL,
    email VARCHAR(100) NOT NULL UNIQUE,
    phone VARCHAR(20) NOT NULL UNIQUE,
    created_at TIMESTAMP DEFAULT NOW(),
    updated_at TIMESTAMP DEFAULT NOW()
);

-- Bank Accounts Table
CREATE TABLE account_dashboard.bank_accounts (
    account_id BIGSERIAL PRIMARY KEY,
    customer_id BIGINT NOT NULL,
    account_number VARCHAR(20) NOT NULL UNIQUE,
    balance NUMERIC(15,2) NOT NULL DEFAULT 0.00,
    currency VARCHAR(10) NOT NULL DEFAULT 'IDR',
    created_at TIMESTAMP DEFAULT NOW(),
    updated_at TIMESTAMP DEFAULT NOW(),
    
    CONSTRAINT bank_accounts_customer_FK FOREIGN KEY (customer_id)
        REFERENCES account_dashboard.customers(customer_id) ON DELETE CASCADE
);

-- Pockets Table
CREATE TABLE account_dashboard.pockets (
    pocket_id BIGSERIAL PRIMARY KEY,
    customer_id BIGINT NOT NULL,
    name VARCHAR(100) NOT NULL,
    balance NUMERIC(15,2) NOT NULL DEFAULT 0.00,
    currency VARCHAR(10) NOT NULL DEFAULT 'IDR',
    created_at TIMESTAMP DEFAULT NOW(),
    updated_at TIMESTAMP DEFAULT NOW(),
    
    CONSTRAINT pockets_customer_FK FOREIGN KEY (customer_id)
        REFERENCES account_dashboard.customers(customer_id) ON DELETE CASCADE
);

-- Term Deposits Table
CREATE TABLE account_dashboard.term_deposits (
    term_deposit_id BIGSERIAL PRIMARY KEY,
    customer_id BIGINT NOT NULL,
    amount NUMERIC(15,2) NOT NULL,
    interest_rate NUMERIC(5,2) NOT NULL,
    maturity_date DATE NOT NULL,
    currency VARCHAR(10) NOT NULL DEFAULT 'IDR',
    created_at TIMESTAMP DEFAULT NOW(),

    CONSTRAINT term_deposits_customer_FK FOREIGN KEY (customer_id)
        REFERENCES account_dashboard.customers(customer_id) ON DELETE CASCADE
);

-- Indexes
CREATE INDEX idx_customers_created_at ON account_dashboard.customers (created_at);
CREATE INDEX idx_bank_accounts_customer_id ON account_dashboard.bank_accounts (customer_id);
CREATE INDEX idx_pockets_customer_id ON account_dashboard.pockets (customer_id);
CREATE INDEX idx_term_deposits_customer_id ON account_dashboard.term_deposits (customer_id);

-- Ensure you're in the correct database
\c account_dashboard_db;

-- Insert customers
INSERT INTO account_dashboard.customers (name, email, phone) VALUES
('John Doe', 'john.doe@email.com', '081234567890'),
('Jane Smith', 'jane.smith@email.com', '082345678901'),
('Alice Johnson', 'alice.johnson@email.com', '083456789012'),
('Bob Williams', 'bob.williams@email.com', '084567890123'),
('Charlie Brown', 'charlie.brown@email.com', '085678901234'),
('Emma Watson', 'emma.watson@email.com', '086789012345'),
('Liam Johnson', 'liam.johnson@email.com', '087890123456'),
('Sophia Miller', 'sophia.miller@email.com', '088901234567'),
('William Davis', 'william.davis@email.com', '089012345678')
ON CONFLICT DO NOTHING;

-- Insert bank accounts
INSERT INTO account_dashboard.bank_accounts (customer_id, account_number, balance, currency) VALUES
(1, '1234567890', 5000000.00, 'IDR'),
(1, '9876543210', 7500000.00, 'IDR'),
(2, '1122334455', 12000000.00, 'USD'),
(3, '5566778899', 2000000.00, 'IDR'),
(4, '9988776655', 30000000.00, 'IDR'),
(5, '3344556677', 6000000.00, 'IDR'),
(6, '2233445566', 9500000.00, 'USD'),
(7, '7788990011', 11000000.00, 'IDR'),
(8, '9900112233', 5000000.00, 'USD'),
(9, '1122003344', 4500000.00, 'IDR')
ON CONFLICT DO NOTHING;

-- Insert pockets
INSERT INTO account_dashboard.pockets (customer_id, name, balance, currency) VALUES
(1, 'Vacation Savings', 2000000.00, 'IDR'),
(2, 'Emergency Fund', 5000000.00, 'IDR'),
(2, 'Shopping Wallet', 1500000.00, 'USD'),
(3, 'Groceries Budget', 700000.00, 'IDR'),
(4, 'Investment Savings', 10000000.00, 'IDR'),
(5, 'Car Savings', 3000000.00, 'IDR'),
(6, 'Health Fund', 8000000.00, 'USD'),
(7, 'Education Savings', 2500000.00, 'IDR'),
(8, 'House Fund', 12000000.00, 'USD'),
(9, 'Business Investment', 5000000.00, 'IDR')
ON CONFLICT DO NOTHING;

-- Insert term deposits
INSERT INTO account_dashboard.term_deposits (customer_id, amount, interest_rate, maturity_date, currency) VALUES
(1, 10000000.00, 5.5, '2026-03-07', 'IDR'),
(2, 25000000.00, 4.8, '2027-06-15', 'USD'),
(3, 15000000.00, 6.2, '2025-12-20', 'IDR'),
(4, 50000000.00, 5.0, '2028-01-10', 'IDR'),
(5, 20000000.00, 4.5, '2026-08-15', 'IDR'),
(6, 18000000.00, 5.0, '2027-09-10', 'USD'),
(7, 30000000.00, 6.1, '2025-11-25', 'IDR'),
(8, 25000000.00, 5.2, '2028-02-15', 'USD'),
(9, 40000000.00, 4.9, '2026-07-30', 'IDR')
ON CONFLICT DO NOTHING;
