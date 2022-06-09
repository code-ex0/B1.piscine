SELECT artists.Name as ArtistName, albums.Title as AlbumName, tracks.Name as TrackName, round(tracks.Bytes / 1000000.0, 2) || ' MB' as MegaBytes FROM albums
JOIN artists on artists.ArtistId = albums.ArtistId
JOIN tracks on albums.AlbumId = tracks.AlbumId
WHERE albums.Title = 'American Idiot';