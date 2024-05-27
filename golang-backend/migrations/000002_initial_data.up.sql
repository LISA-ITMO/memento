INSERT INTO users (username, created_at) VALUES
    ('Big baby tape', '1999-01-08 04:05:06'),
    ('Kizaru', '1999-02-08 04:05:06');

INSERT INTO folders (name, owner, parent, created_at) VALUES
    ('main', 1, null, '1999-02-08 04:05:06'),
    ('main', 2, null, '1999-02-08 04:05:06');

INSERT INTO users_to_main_folders ("user", folder) VALUES
    (1,1),
    (2,2);

INSERT INTO notes (owner, folder, created_at, name, is_public) VALUES
    (1, 1, '1999-02-08 04:05:06', 'tape note 1', false),
    (1, 1, '1999-02-08 04:05:06', 'tape note 2', false),
    (1, 1, '1999-02-08 04:05:06', 'tape note 3', false);

INSERT INTO questions (owner, note, body, answer, is_public) VALUES
    (1, 1, 'question body 1', 'answer 1', false),
    (1, 1, 'question body 2', 'answer 2', false),
    (1, 1, 'question body 3', 'answer 3', false);


