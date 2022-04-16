DROP TABLE IF EXISTS products;
DROP TABLE IF EXISTS categories;

CREATE TABLE categories (
    id serial NOT NULL PRIMARY KEY,
    name varchar(15) NOT NULL,
    description text
);

CREATE TABLE products (
    id serial NOT NULL PRIMARY KEY,
    name varchar(40) NOT NULL,
    category_id integer,
    unit_price real,
    units_in_stock integer
);


INSERT INTO categories (name, description) VALUES ('飲料', 'ソフトドリンク、お茶、コーヒーなど');
INSERT INTO categories (name, description) VALUES ('生鮮食品', '生野菜、肉、魚など');
INSERT INTO categories (name, description) VALUES ('調味料', '塩、コショウ、醤油、みそなど');

INSERT INTO products (name, category_id, unit_price, units_in_stock) VALUES ('黒豆茶', 1, 300, 20);
INSERT INTO products (name, category_id, unit_price, units_in_stock) VALUES ('音速コーヒー', 1, 130, 40);
INSERT INTO products (name, category_id, unit_price, units_in_stock) VALUES ('メカコーラ', 1, 120, 10);
INSERT INTO products (name, category_id, unit_price, units_in_stock) VALUES ('レモン豆乳', 1, 30, 100);
INSERT INTO products (name, category_id, unit_price, units_in_stock) VALUES ('レタス', 2, 50, 200);
INSERT INTO products (name, category_id, unit_price, units_in_stock) VALUES ('鶏むね肉', 2, 70, 88);
INSERT INTO products (name, category_id, unit_price, units_in_stock) VALUES ('鯛あら', 2, 90, 23);
INSERT INTO products (name, category_id, unit_price, units_in_stock) VALUES ('岩塩', 3, 50, 11);
INSERT INTO products (name, category_id, unit_price, units_in_stock) VALUES ('みりん', 3, 100, 5);
