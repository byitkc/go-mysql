CREATE DATABASE testing;
CREATE USER 'testing'@'%' IDENTIFIED BY 'password';
GRANT ALL PRIVILEGES ON testing.* to 'testing'@'%';