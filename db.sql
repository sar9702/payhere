CREATE DATABASE payhere;

CREATE TABLE payhere.Item (
	Id varchar(100) PRIMARY KEY,
	Category varchar(100),
	Name varchar(100),
	Price varchar(10),
	Cost varchar(10),
	Description varchar(100),
	Barcode varchar(100),
	ExpirationDate varchar(10),
	`Size` varchar(10),
 	`Chosung` varchar(100)
);

CREATE TABLE payhere.`User` (
	ID varchar(20) PRIMARY KEY,
	Password varchar(100),
	Token varchar(200),
	SignKey varchar(100)
);