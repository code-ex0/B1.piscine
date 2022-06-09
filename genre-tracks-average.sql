SELECT genres.Name, AVG(Milliseconds) as AverageDuration
FROM genres
JOIN tracks on genres.GenreId = tracks.GenreId
GROUP BY genres.Name
ORDER BY AverageDuration DESC;