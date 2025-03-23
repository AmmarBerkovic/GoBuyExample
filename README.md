## Proposed SQL Queries

For local testing, you can execute the following SQL commands:

```sql
CREATE DATABASE go_user_db;
CREATE USER 'go_user'@'localhost' IDENTIFIED BY 'go_password';
GRANT ALL PRIVILEGES ON go_user_db.* TO 'go_user'@'localhost';
FLUSH PRIVILEGES;

CREATE TABLE user_dbs (
    ID INT AUTO_INCREMENT PRIMARY KEY,
    Name VARCHAR(255) NOT NULL,
    Email VARCHAR(255) NOT NULL,
    Age INT NOT NULL
);

INSERT INTO user_dbs (Name, Email, Age)
VALUES
  ('Alice Johnson', 'alice.johnson@example.com', 30),
  ('Bob Smith', 'bob.smith@example.com', 25),
  ('Charlie Brown', 'charlie.brown@example.com', 35),
  ('Diana Prince', 'diana.prince@example.com', 28),
  ('Frank Castle', 'frank.castle@example.com', 38),
  ('Grace Harper', 'grace.harper@example.com', 27),
  ('Henry Adams', 'henry.adams@example.com', 45),
  ('Isabella Clarke', 'isabella.clarke@example.com', 32),
  ('Jack Dawson', 'jack.dawson@example.com', 29),
  ('Katherine Pierce', 'katherine.pierce@example.com', 34),
  ('Liam OBrien', 'liam.obrien@example.com', 31),
  ('Monica Bell', 'monica.bell@example.com', 26),
  ('Nathan Drake', 'nathan.drake@example.com', 37),
  ('Olivia Turner', 'olivia.turner@example.com', 33);
```
