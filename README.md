# Go Fiber Clean Architecture

This project was created to learn golang with go fiber framework

## How To Run

1. Run `go mod vendor`
2. Change  value of file `.env` form smtp
3. Run docker compose with command `docker compose up`
4. manually seed to your db user and user_role
5. manually seed to your db product
```
INSERT INTO product (product_id, name, price, description, image) 
VALUES 
    ('f0b88290-7948-4d6f-91be-4b8896c4f6d2', 'Product A', 10.99, 'This is the description of Product A', 'https://example.com/product-a.jpg'),
    ('9b84a173-7e2b-4ce8-8f87-6d2461a2f04d', 'Product B', 19.99, 'This is the description of Product B', 'https://example.com/product-b.jpg'),
    ('c5dd6638-9ec1-41a3-ae5e-51a82df61a2d', 'Product C', 7.99, 'This is the description of Product C', 'https://example.com/product-c.jpg'),
    ('79a3f3d8-ba8f-43e9-b18e-9a365ed90e8e', 'Product D', 14.99, 'This is the description of Product D', 'https://example.com/product-d.jpg'),
    ('c1b407d7-434d-47da-97ee-8a56a554c3a7', 'Product E', 8.99, 'This is the description of Product E', 'https://example.com/product-e.jpg');

```
6. use cognotiv.postman.collection.json
