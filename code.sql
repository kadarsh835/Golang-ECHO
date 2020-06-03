CREATE TABLE IF NOT EXISTS events (
    id INT PRIMARY KEY,
    title varchar(50) NOT NULL,
    prize INT NOT NULL,
    head varchar(50) NOT NULL,
    phone varchar(12) NOT NULL
);

INSERT INTO events VALUES (
    0, 
    'Hover Rumble', 
    15000,
    'Adarsh Kumar',
    7355789654
);

INSERT INTO events VALUES (
    1, 
    'P=NP', 
    25000,
    'Adarsh',
    7356789654
);