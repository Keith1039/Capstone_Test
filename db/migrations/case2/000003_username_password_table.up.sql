CREATE TABLE IF NOT EXISTS USER_AUTH(
    USERNAME VARCHAR(255) NOT NULL,
    PASSWORD VARCHAR(255) NOT NULL,
    USER_ID INTEGER PRIMARY KEY,
    FOREIGN KEY (USER_ID) REFERENCES USERS
);