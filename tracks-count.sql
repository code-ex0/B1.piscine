SELECT genres.Name, count(tracks.GenreId) as NumberOfTracks
FROM genres
JOIN tracks on genres.GenreId = tracks.GenreId
GROUP BY genres.Name;