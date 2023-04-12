CREATE SCHEMA unnispick_database;
use unnispick_database;

CREATE TABLE members (
    ID_MEMBER INT PRIMARY KEY NOT NULL,
    USERNAME VARCHAR(50) NOT NULL,
    GENDER ENUM('Male', 'Female') NOT NULL,
    SKIN_TYPE ENUM('Oily', 'Dry', 'Combination', 'Sensitive', 'Normal') NOT NULL,
    SKIN_COLOR ENUM('Fair', 'Medium', 'Olive', 'Brown', 'Dark') NOT NULL
);

ALTER TABLE members ADD COLUMN created_at TIMESTAMP NOT NULL DEFAULT CURRENT_TIMESTAMP;
ALTER TABLE members ADD COLUMN updated_at TIMESTAMP NOT NULL DEFAULT CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP;
ALTER TABLE members ADD COLUMN deleted_at TIMESTAMP NULL;

ALTER TABLE `unnispick_database`.`members` 
CHANGE COLUMN `SKINTYPE` `SKIN_TYPE` ENUM('Oily', 'Dry', 'Combination', 'Sensitive', 'Normal') NOT NULL ,
CHANGE COLUMN `SKINCOLOR` `SKIN_COLOR` ENUM('Fair', 'Medium', 'Olive', 'Brown', 'Dark') NOT NULL ;

ALTER TABLE `unnispick_database`.`review_product` DROP FOREIGN KEY `review_product_ibfk_2`;
ALTER TABLE `unnispick_database`.`like_review` DROP FOREIGN KEY `like_review_ibfk_2`;
ALTER TABLE `unnispick_database`.`members` 
CHANGE COLUMN `ID_MEMBER` `ID_MEMBER` INT NOT NULL AUTO_INCREMENT ;

ALTER TABLE `unnispick_database`.`review_product` ADD CONSTRAINT `review_product_ibfk_2` FOREIGN KEY (`ID_MEMBER`) REFERENCES `members`(`ID_MEMBER`) ON DELETE CASCADE ON UPDATE CASCADE;
ALTER TABLE `unnispick_database`.`like_review` ADD CONSTRAINT `like_review_ibfk_2` FOREIGN KEY (`ID_MEMBER`) REFERENCES `members`(`ID_MEMBER`) ON DELETE CASCADE ON UPDATE CASCADE;

SELECT * FROM members;

CREATE TABLE Product (
    ID_PRODUCT INT PRIMARY KEY NOT NULL,
    NAME_PRODUCT VARCHAR(50) NOT NULL,
    PRICE INT NOT NULL
);

SELECT * FROM product;

CREATE TABLE review_product (
    ID_REVIEW INT PRIMARY KEY NOT NULL,
    ID_PRODUCT INT NOT NULL,
    ID_MEMBER INT NOT NULL,
    DESC_REVIEW VARCHAR(500) NOT NULL,
    FOREIGN KEY (ID_PRODUCT) REFERENCES product(ID_PRODUCT),
    FOREIGN KEY (ID_MEMBER) REFERENCES member(ID_MEMBER)
);

SELECT * FROM review_product;

-- to make sure each review only can be liked once by each member so we should have to primary key in this table
CREATE TABLE Like_review (
    ID_REVIEW INT NOT NULL,
    ID_MEMBER INT NOT NULL,
    PRIMARY KEY (ID_REVIEW, ID_MEMBER),
    FOREIGN KEY (ID_REVIEW) REFERENCES Review_product(ID_REVIEW),
    FOREIGN KEY (ID_MEMBER) REFERENCES member(ID_MEMBER)
);

SELECT * FROM like_review;

SELECT * FROM `members` WHERE ID_MEMBER = '1' AND `members`.`deleted_at` IS NULL ORDER BY `members`.`id` LIMIT 1;