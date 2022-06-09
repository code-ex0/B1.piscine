SELECT EmployeeId, LastName, FirstName, STRFTIME('%d/%m/%Y', HireDate) as HireFrenchDate
FROM employees
ORDER BY HireDate;