CREATE DATABASE Calculator;
CREATE TABLE logs (
  id INT AUTO_INCREMENT PRIMARY KEY,
  right_operand INT NOT NULL,
  operator VARCHAR(10) NOT NULL,
  left_operand INT NOT NULL,
  result FLOAT NOT NULL,
  date_created DATETIME NOT NULL DEFAULT CURRENT_TIMESTAMP
);