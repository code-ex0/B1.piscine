SELECT Title AS AlbumName,
       Name AS ArtistName
FROM albums JOIN artists
    on albums.ArtistId = artists.ArtistId
    LIMIT 100;