CREATE TABLE mst_services(
	id INT NOT NULL,
	service_name VARCHAR(100) NOT NULL,
	unit VARCHAR(100) NOT NULL,
	price INT,
	PRIMARY KEY(id)
);

CREATE TABLE mst_customers(
	id INT NOT NULL,
	customer_name VARCHAR(100) NOT NULL,
	phone_number VARCHAR(16) NOT NULL,
	date_created DATE,
	PRIMARY KEY(id)
);

CREATE TABLE mst_employees(
	id INT NOT NULL,
	employee_name VARCHAR(100) NOT NULL,
	date_created DATE,
	PRIMARY KEY(id)
);

CREATE TABLE tx_enigma_laundry(
	id INT NOT NULL,
	id_employee INT NOT NULL,
	id_customer INT NOT NULL,
	id_service INT NOT NULL,
	transaction_in DATE NOT NULL,
	transaction_out DATE NOT NULL,
	amount INT NOT NULL,
	PRIMARY KEY(id),
	FOREIGN KEY(id_employee) REFERENCES mst_employees(id),
	FOREIGN KEY(id_customer) REFERENCES mst_customers(id),
	FOREIGN KEY(id_service) REFERENCES mst_services(id)
);

INSERT INTO mst_services(id, service_name, unit, price) VALUES(1, 'Cuci + Setrika', 'KG', 7000);
INSERT INTO mst_services(id, service_name, unit, price) VALUES(2, 'Laundry Bedcover', 'Buah', 50000);
INSERT INTO mst_services(id, service_name, unit, price) VALUES(3, 'Laundry Boneka', 'Buah', 25000);
INSERT INTO mst_customers(id, customer_name, phone_number, date_created) VALUES(1, 'Trixie', '08888280098', '2024-01-10');
INSERT INTO mst_customers(id, customer_name, phone_number, date_created) VALUES(2, 'Adrian', '085802520642', '2023-02-10');
INSERT INTO mst_customers(id, customer_name, phone_number, date_created) VALUES(3, 'Jessica', '0812654987', '2023-09-11');
INSERT INTO mst_employees(id, employee_name, date_created) VALUES(1, 'Yeni', '2023-01-10');
INSERT INTO mst_employees(id, employee_name, date_created) VALUES(2, 'Mirna', '2024-01-05');
INSERT INTO tx_enigma_laundry(id, id_employee, id_customer, id_service, transaction_in, transaction_out, amount) VALUES(1, 2, 3, 1, '2022-08-18', '2022-08-20', 5);
INSERT INTO tx_enigma_laundry(id, id_employee, id_customer, id_service, transaction_in, transaction_out, amount) VALUES(2, 2, 3, 2, '2022-08-18', '2022-08-20', 1);
INSERT INTO tx_enigma_laundry(id, id_employee, id_customer, id_service, transaction_in, transaction_out, amount) VALUES(3, 2, 3, 3, '2022-08-18', '2022-08-20', 2);

SELECT * FROM mst_services;
SELECT * FROM mst_customers;
SELECT * FROM mst_employees;
SELECT * FROM tx_enigma_laundry;