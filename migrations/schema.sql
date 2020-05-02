-- +migrate Up

CREATE EXTENSION IF NOT EXISTS citext;

CREATE TABLE user_interface (
                      nickname citext PRIMARY KEY,
                      email citext UNIQUE NOT NULL,
                      fullname text,
                      about text,
                      country text,
                      password text,
                      role int DEFAULT 1 -- 0 - banned; 1 - user; 2 - moderator; 3 - admin
);

CREATE INDEX users_cover_index
    ON user_interface (nickname, email, about, fullname, country);

CREATE TABLE thread (
                        thread_created timestamp with time zone DEFAULT now(),
                        thread_title text NOT NULL,
                        d_thread SERIAL PRIMARY KEY
);

CREATE TABLE post (
                      post_author citext REFERENCES user_interface ON DELETE CASCADE NOT NULL,
                      post_thread integer REFERENCES thread ON DELETE CASCADE NOT NULL,
                      post_created timestamp with time zone DEFAULT now(),
                      post_title text NOT NULL,
                      post_topic text NOT NULL,
                      post_id serial PRIMARY KEY

);

CREATE TABLE comment (
                         comment_author citext REFERENCES user_interface ON DELETE CASCADE NOT NULL,
                         comment_post integer REFERENCES post ON DELETE CASCADE NOT NULL ,
                         comment_content text NOT NULL,
                         comment_created timestamp with time zone DEFAULT now(),
                         likes integer DEFAULT 0,
                         id_comment serial PRIMARY KEY
);

CREATE TABLE likes (
                       nickname citext REFERENCES user_interface ON DELETE CASCADE NOT NULL,
                       id_comment integer REFERENCES comment ON DELETE CASCADE NOT NULL,
                       CONSTRAINT uniq_like UNIQUE (nickname, id_comment)
);


-- +migrate StatementBegin
CREATE OR REPLACE FUNCTION add_like() RETURNS TRIGGER AS $add_like$
BEGIN
    IF (TG_OP = 'INSERT') THEN
        UPDATE comment SET likes = likes + 1 WHERE id_comment = NEW.id_comment;
        RETURN NEW;
    END IF;
    RETURN NULL;
END;
$add_like$ LANGUAGE plpgsql;
-- +migrate StatementEnd

CREATE TRIGGER add_like AFTER INSERT ON likes
    FOR EACH ROW EXECUTE PROCEDURE add_like();

-- +migrate StatementBegin
CREATE OR REPLACE FUNCTION minus_like() RETURNS TRIGGER AS $minus_like$
BEGIN
    IF (TG_OP = 'DELETE') THEN
        UPDATE comment SET likes = likes - 1 WHERE id_comment = OLD.id_comment;
        RETURN NEW;
    END IF;
    RETURN NULL;
END;
$minus_like$ LANGUAGE plpgsql;
-- +migrate StatementEnd

CREATE TRIGGER minus_like BEFORE DELETE ON likes
    FOR EACH ROW EXECUTE PROCEDURE minus_like();