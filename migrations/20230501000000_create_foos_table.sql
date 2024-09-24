-- +goose Up
CREATE TABLE foos
(
    id   VARCHAR(36) PRIMARY KEY,
    name VARCHAR(255) NOT NULL
);

INSERT INTO foos (id, name)
VALUES ('1', 'foo'),
       ('2', 'bar'),
       ('3', 'baz'),
       ('4', 'qux'),
       ('5', 'quux');

-- +goose Down
DROP TABLE foos;