DELETE FROM tracks
where TrackId IN (
    SELECT TrackId FROM albums
        JOIN tracks on albums.AlbumId = tracks.AlbumId
    WHERE albums.Title = 'Facelift'
);