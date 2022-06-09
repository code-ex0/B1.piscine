SELECT Name
FROM artists
INNER JOIN albums on artists.ArtistId = albums.ArtistId
GROUP BY Name
HAVING COUNT(Title) >= 4
ORDER BY Name DESC;