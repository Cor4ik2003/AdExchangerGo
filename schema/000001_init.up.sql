CREATE TABLE Users
(
    UserID serial not null unique,
    name varchar(255) not null,
    email varchar(255) not null,
    password_hash varchar(255) not null
);

CREATE TABLE Images
(
    ImageID INT PRIMARY KEY,
    AdID INT,
    ImageURL VARCHAR(255) NOT NULL,
);


CREATE TABLE Advertisements
(
    AdID serial not null unique,
    title varchar(255) not null,
    description varchar(255),
    category varchar(255) not null,
    price DECIMAL(10,2)
);

CREATE TABLE Categories
(
    CategoryID INT PRIMARY KEY,
    CategoryName VARCHAR(255) NOT NULL
);