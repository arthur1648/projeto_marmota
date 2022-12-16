create database db_marmota;

use db_marmota;

CREATE TABLE IF NOT EXISTS produtos (
    id         INTEGER        PRIMARY KEY AUTO_INCREMENT UNIQUE  NOT NULL,
    name       VARCHAR(120)   NOT NULL,
    code       VARCHAR(20)    NOT NULL,
    price      DECIMAL(10, 2) NOT NULL,
	created_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP,
    updated_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP
);

DELIMITER //
CREATE TRIGGER tb_produto_updated_at_trig after update ON produtos
      FOR EACH ROW
BEGIN
    UPDATE produtos
       SET updated_at = CURRENT_TIMESTAMP
     WHERE id = NEW.id;
END //
DELIMITER ;

INSERT INTO tb_produto (name, code, price) VALUES
    ('Notebook 13 XPS', 'DELL', 14500),
    ('Notebook 15', 'DELL', 2500),
    ('Notebook 14', 'DELL', 2500),
    ('Notebook 15', 'Lenovo', 3500),
    ('Notebook 13', 'Lenovo', 4568),
    ('Tablet', 'SAM', 8450),
    ('Macbook 13 Pro M1', 'Apple', 18500.00),
    ('TV 55', 'SONY', 4500.00),
    ('TV 45', 'SONY', 3500.00),
    ('TV 32', 'SAM', 2500.00),
    ('TV 60', 'LG', 6500.00),
    ('TV 50', 'LG', 4800.00);