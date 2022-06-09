SELECT employees.FirstName || ' ' || UPPER(employees.LastName) AS FullName, COUNT(customers.SupportRepId) AS NumberOfCustomers
FROM employees
JOIN customers on employees.EmployeeId = customers.SupportRepId
GROUP BY employees.FirstName
ORDER BY NumberOfCustomers;