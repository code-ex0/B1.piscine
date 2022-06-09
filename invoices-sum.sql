SELECT SUM(Total) as AllInvoicesTotalPrice
FROM invoices
WHERE CustomerId IN (
    SELECT CustomerId
    FROM customers
    WHERE FirstName = 'Tim'
      AND  LastName = 'Goyer'
);