-- Mendapatkan total pesanan untuk setiap produk
SELECT products.name, SUM(orders.quantity) AS total_orders
FROM orders
JOIN products ON orders.product_id = products.id
GROUP BY products.name;

-- Melihat tingkat stok di lokasi tertentu
SELECT products.name, inventory.quantity, inventory.location
FROM inventory
JOIN products ON inventory.product_id = products.id
WHERE inventory.location = 'Gudang A';
