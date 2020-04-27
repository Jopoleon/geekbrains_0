CREATE TABLE IF NOT EXISTS books  (
    name varchar(36),
    pages integer ,
    created_on TIMESTAMP WITH TIME ZONE DEFAULT CURRENT_TIMESTAMP
    );


CREATE TABLE IF NOT EXISTS authors  (
    first_name varchar(36),
    second_name varchar(36) ,
    created_on TIMESTAMP WITH TIME ZONE DEFAULT CURRENT_TIMESTAMP
    );

INSERT INTO books (name, pages) VALUES ('War And Peace',1225);
INSERT INTO books (name, pages) VALUES ('The Catcher in the Rye', 277);
INSERT INTO authors (first_name, second_name) VALUES ('Aleksey', 'Tolstoi');
INSERT INTO authors (first_name, second_name) VALUES ('Jerome','Salinger');
