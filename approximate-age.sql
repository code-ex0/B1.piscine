SELECT FirstName, LastName, STRFTIME('%Y', HireDate) - STRFTIME('%Y', BirthDate) AS ApproximateAge
FROM employees
ORDER BY ApproximateAge;