SELECT customers.FirstName || ' ' || UPPER(customers.LastName) AS FullName, SUM(Total) as AllInvoices
FROM customers
INNER JOIN invoices on customers.CustomerId = invoices.CustomerId
GROUP BY FullName
HAVING AllInvoices >= 38
ORDER BY FullName;