INSERT INTO products (name, description, price, category, image_path)
VALUES
('Laptop TUF Gaming', 'Laptop dengan performa tinggi untuk gaming', 17000000.00, 'Elektronik', 'image 1'),
('Iphone XI', 'Smartphone ', 15000000.00, 'Elektronik', 'image 2'),
('Headphone Sony', 'Headphone kualitas suara premium', 850000.00, 'Aksesoris', 'image 3'),
('Kursi Gaming Secret Lab', 'Kursi Game ergonomis untuk kenyamanan maksimal', 2200000.00, 'Furniture', 'image 4'),
('Meja Gaming', 'Meja Gaming', 3000000.00, 'Furniture', 'image 5');

INSERT INTO inventory (product_id, quantity, location)
VALUES
(1, 20, 'Gudang A'),
(2, 75, 'Gudang B'),
(3, 50, 'Gudang C'),
(4, 20, 'Gudang D'),
(5, 15, 'Gudang E');

INSERT INTO orders (product_id, quantity, order_date) VALUES
(1, 2, '2024-11-24 22:37:00')

