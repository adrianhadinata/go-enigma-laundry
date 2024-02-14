-- Basic query
SELECT * FROM mst_services;
SELECT * FROM mst_customers;
SELECT * FROM mst_employees;
SELECT * FROM tx_enigma_laundry;

-- Relation Query
SELECT 
	t1.id,
	t2.employee_name,
	t3.customer_name,
	t4.service_name,
	t4.price,
	t1.amount,
	t4.unit,
	(t1.amount * t4.price) AS total
FROM tx_enigma_laundry AS t1
JOIN mst_employees AS t2 ON t1.id_employee = t2.id
JOIN mst_customers AS t3 ON t1.id_customer = t3.id
JOIN mst_services AS t4 ON t1.id_service = t4.id;