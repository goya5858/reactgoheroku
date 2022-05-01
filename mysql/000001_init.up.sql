USE heroku_4c1b06639a7b676;
DROP TABLE IF EXISTS test_table;
 
CREATE TABLE test_table (
id INT NOT NULL AUTO_INCREMENT PRIMARY KEY,
ItemName TEXT NOT NULL,
Price INT NOT NULL,
Stock INT NOT NULL
);

 
INSERT INTO test_table (ItemName, Price, Stock) VALUES ("item_1", 250, 100);
INSERT INTO test_table (ItemName, Price, Stock) VALUES ("item_2", 120, 200);