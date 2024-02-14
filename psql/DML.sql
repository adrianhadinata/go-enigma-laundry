-- Data - data
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