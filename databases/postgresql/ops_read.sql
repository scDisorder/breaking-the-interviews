select * from public.artists;

select * from public.albums;

select * from public.albums where artist_id = 2;

select * from public.songs where album_id = 1;

select a.artist_id, sum(s.length) from public.artists
    left join albums a on artists.artist_id = a.artist_id
    left join songs s on a.album_id = s.album_id
group by a.artist_id
having sum(s.length) > 0;

select * from public.artists where artist_id in (1, 2, 5);
