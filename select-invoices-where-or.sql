SELECT InvoiceId, Total FROM invoices
WHERE BillingCity = 'Paris' OR BillingCity = 'London'
ORDER BY InvoiceId;