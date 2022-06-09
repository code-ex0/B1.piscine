SELECT artists.Name,
   COUNT(DISTINCT albums.AlbumId) as NbAlbums,
       case
           WHEN COUNT(DISTINCT albums.AlbumId) = 1 then 'Unproductive'
           WHEN COUNT(DISTINCT albums.AlbumId) > 1 AND COUNT(DISTINCT albums.AlbumId) < 10 then 'Productive'
           ELSE 'Very Productive'
       end as IsProductive
FROM artists
     INNER JOIN albums on artists.ArtistId = albums.ArtistId
GROUP BY artists.Name
ORDER BY artists.ArtistId
LIMIT 100;