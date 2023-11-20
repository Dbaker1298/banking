CREATE DATABASE banking;
USE banking;

DROP TABLE IF EXISTS `customers`;
CREATE TABLE `customers` (
  `customer_id` int(11) NOT NULL AUTO_INCREMENT,
  `name` varchar(100) NOT NULL,
  `date_of_birth` date NOT NULL,
  `city` varchar(100) NOT NULL,
  `zipcode` varchar(10) NOT NULL,
  `status` tinyint(1) NOT NULL DEFAULT '1',
  PRIMARY KEY (`customer_id`)
  ) ENGINE=InnoDB AUTO_INCREMENT=2006 DEFAULT CHARSET=utf8;

  INSERT INTO `customers` VALUES
  (2001,'John Doe','1980-01-01','New York, NY','10001',1),
  (2002,'Jane Doe','1981-02-02','New York, NY','10665',0),
  (2003,'John Smith','1982-03-03','New York NY','12311',1),
  (2004,'Jane Smith','1983-04-04','New York, NY','29301',0),
  (2005,'John Doe','1984-05-05','New York, NY','10394',1);

  DROP TABLE IF EXISTS `accounts`;
  CREATE TABLE `accounts` (
    `account_id` int(11) NOT NULL AUTO_INCREMENT,
    `customer_id` int(11) NOT NULL,
    `opening_date` datetime NOT NULL DEFAULT CURRENT_TIMESTAMP,
    `account_type` varchar(10) NOT NULL,
    `pin` varchar(10) NOT NULL,
    `status` tinyint(4) NOT NULL DEFAULT '1',
    PRIMARY KEY (`account_id`),
    KEY `accounts_FK` (`customer_id`),
    CONSTRAINT `accounts_FK` FOREIGN KEY (`customer_id`) REFERENCES `customers` (`customer_id`) 
    ) ENGINE=InnoDB AUTO_INCREMENT=95476 DEFAULT CHARSET=utf8;

    INSERT INTO `accounts` VALUES
    (95471,2001,'2018-01-01 00:00:00','savings','1234',1),
    (95472,2002,'2018-01-01 00:00:00','savings','1342',1),
    (95473,2003,'2018-01-01 00:00:00','checking','2234',1),
    (95474,2004,'2018-01-01 00:00:00','savings','1299',1),
    (95475,2005,'2018-01-01 00:00:00','checking','1264',0);

    DROP TABLE IF EXISTS `transactions`;

    CREATE TABLE `transactions` (
      `transaction_id` int(11) NOT NULL AUTO_INCREMENT,
      `account_id` int(11) NOT NULL,
      `amount` int(11) NOT NULL,
      `transaction_type` varchar(10) NOT NULL,
      `transaction_date` datetime NOT NULL DEFAULT CURRENT_TIMESTAMP,
      PRIMARY KEY (`transaction_id`),
      KEY `transactions_FK` (`account_id`),
      CONSTRAINT `transactions_FK` FOREIGN KEY (`account_id`) REFERENCES `accounts` (`account_id`)
      ) ENGINE=InnoDB AUTO_INCREMENT=1000001 DEFAULT CHARSET=utf8;
    
