SELECT FirstName, LastName, Country,
   CASE Country
       WHEN 'Brazil' THEN 'America'
       WHEN 'Canada' THEN 'America'
       WHEN 'USA' THEN 'America'
       WHEN 'Argentina' THEN 'America'
       WHEN 'Chile' THEN 'America'

       WHEN 'Germany' THEN 'Europe'
       WHEN 'Norway' THEN 'Europe'
       WHEN 'Czech Republic' THEN 'Europe'
       WHEN 'Austria'  THEN 'Europe'
       WHEN 'Belgium' THEN 'Europe'
       WHEN 'Denmark' THEN 'Europe'
       WHEN 'Portugal' THEN 'Europe'
       WHEN 'France' THEN 'Europe'
       WHEN 'Finland' THEN 'Europe'
       WHEN 'Hungary' THEN 'Europe'
       WHEN 'Ireland' THEN 'Europe'
       WHEN 'Italy' THEN 'Europe'
       WHEN 'Netherlands' THEN 'Europe'
       WHEN 'Poland' THEN 'Europe'
       WHEN 'Spain' THEN 'Europe'
       WHEN 'Sweden' THEN 'Europe'
       WHEN 'Sweden' THEN 'Europe'
       WHEN 'United Kingdom' THEN 'Europe'

       WHEN 'India' THEN 'Asia'

       WHEN 'Australia' THEN 'Oceania'
       else 'NO CONTINENT'
       end as Continent
FROM customers