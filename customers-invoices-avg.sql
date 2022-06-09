SELECT FirstName, LastName, AVG(Total) as InvoicesAverage
FROM customers
INNER JOIN invoices on customers.CustomerId = invoices.CustomerId
GROUP BY FirstName, LastName
ORDER BY FirstName;