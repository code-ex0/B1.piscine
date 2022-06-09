SELECT DISTINCT tracks.Name as TrackName, playlists.Name as PlaylistName
FROM playlist_track
INNER JOIN playlists on playlist_track.PlaylistId = playlists.PlaylistId
INNER JOIN tracks on playlist_track.TrackId = tracks.TrackId
WHERE playlists.Name = 'TV Shows'
LIMIT 100;