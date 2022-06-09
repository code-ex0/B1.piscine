SELECT invoices.InvoiceId, tracks.Name as InvoiceItem, invoice_items.UnitPrice
FROM invoices
INNER JOIN invoice_items on invoices.InvoiceId = invoice_items.InvoiceId
INNER JOIN tracks on invoice_items.TrackId = tracks.TrackId
WHERE invoices.InvoiceId = 10
ORDER BY Name;