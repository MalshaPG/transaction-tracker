DROP TABLE IF EXISTS transactions;
CREATE TABLE transactions (
    id INT AUTO_INCREMENT NOT NULL,
    type VARCHAR(100) NOT NULL,
    description VARCHAR(200) NOT NULL,
    date DATE NOT NULL,
    amount DECIMAL(20, 2) NOT NULL,
    CONSTRAINT pk_transactions PRIMARY KEY (id)   
);

INSERT INTO transactions (type, description, date, amount) 
VALUES
('income', 'Salary', '2025-01-01', 5000.00),
('expense', 'Groceries', '2025-01-02', -150.00),
('income', 'Rent from 627/D', '2025-01-03', 1200.00),
('expense', 'Utilities', '2025-01-04', -200.00),
('income', 'Investment Return', '2025-01-05', 300.00),
('expense', 'Dining Out', '2025-01-06', -100.00);