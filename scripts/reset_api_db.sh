# truncate all merchant db tables
mysql -u root -h 127.0.0.1 -e "drop database if exists bookstore_db;"
mysql -u root -h 127.0.0.1 -e "create database bookstore_db;"
mysql -u root -h 127.0.0.1 bookstore_db < data/db.sql