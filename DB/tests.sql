CREATE TABLE todo (
    id SERIAL,
    description TEXT,
    titel TEXT
);

INSERT INTO todo (id, description, titel) VALUES (1,'test', 'Test1'), (2, 'test2', 'Test2');