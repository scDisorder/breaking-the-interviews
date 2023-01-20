create table public.artists
(
    artist_id serial
        constraint artists_pk
            primary key,
    name      text
);

create unique index artists_artist_id_uindex
    on public.artists (artist_id);

create table public.albums
(
    album_id  serial
        constraint albums_pk
            primary key,
    name      text    not null,
    artist_id integer not null
        constraint albums_artists_artist_id_fk
            references public.artists
            on update cascade on delete cascade
);

create unique index albums_album_id_uindex
    on public.albums (album_id);

create table public.songs
(
    song_id  serial
        constraint songs_pk
            primary key,
    name     text    not null,
    length   integer,
    album_id integer not null
        constraint songs_albums_album_id_fk
            references public.albums
            on update cascade on delete cascade
);

create unique index songs_song_id_uindex
    on public.songs (song_id);

create table public.users
(
    id           bigserial,
    shortlink    text,
    name         text not null,
    address      text,
    biography    text,
    email        text,
    image        text,
    background   text,
    discord_id   text,
    twitter_id   text,
    instagram_id text,
    github_id    text,
    flags        bigint,
    roles        bigint,
    created_at   bigint,
    updated_at   bigint
);

create unique index users_id_uindex
    on public.users (id);

create unique index users_shortlink_uindex
    on public.users (shortlink);

