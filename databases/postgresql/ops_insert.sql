INSERT INTO public.artists (name)
VALUES ('Bring Me The Horizon'),
       ('Star Lord'),
       ('We Butter The Bread With Butter'),
       ('Tenacious D'),
       ('The Devil Wears Prada'),
       ('Every Time I Die');

INSERT INTO public.albums (name, artist_id)
VALUES ('Space Rider', 2),
       ('POST HUMAN: SURVIVAL', 1),
       ('amo', 1),
       ('Sempiternal', 1),
       ('ZII', 5),
       ('With Roots Above And Branches Below', 5),
       ('Tribute', 4),
       ('Das Album', 3),
       ('Das Monster aus dem Schrank', 3);

INSERT INTO public.songs (name, length, album_id)
VALUES ('Parasite Eve', 292, 2),
       ('Dear Diary', 165, 2),
       ('Ludens', 281, 2),
       ('Kingslayer', 221, 2),
       ('Forlorn', 241, 5),
       ('The Darkness Inside', 471, 1),
       ('Watch Me Shine', 302, 1),
       ('Zero To Hero', 200, 1);
