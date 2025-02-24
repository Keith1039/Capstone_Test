CREATE TABLE IF NOT EXISTS SHOP (
    ITEM_ID INTEGER PRIMARY KEY,
    PRICE INTEGER,
    AMOUNT INTEGER
);

CREATE TABLE IF NOT EXISTS PURCHASE (
    PURCHASE_ID INTEGER PRIMARY KEY,
    ITEM_ID INTEGER,
    PAYMENT INTEGER,
    FOREIGN KEY (ITEM_ID) REFERENCES SHOP
)