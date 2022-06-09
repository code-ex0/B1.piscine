SELECT Title as AlbumName, Name as TrackName, Milliseconds
FROM albums
INNER JOIN tracks on albums.AlbumId = tracks.AlbumId
ORDER BY Milliseconds
LIMIT 50;