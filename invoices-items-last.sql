SELECT Name FROM invoices
    INNER JOIN invoice_items on invoices.InvoiceId = invoice_items.InvoiceId
    JOIN tracks on tracks.TrackId = invoice_items.TrackId
WHERE invoice_items.InvoiceId  = (
SELECT invoice_items.InvoiceId FROM invoices
    JOIN invoice_items on invoices.InvoiceId = invoice_items.InvoiceId
    ORDER BY invoices.InvoiceDate DESC
    LIMIT 1
)