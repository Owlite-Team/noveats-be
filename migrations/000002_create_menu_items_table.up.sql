CREATE TABLE IF NOT EXISTS MenuItems
(
    id          BIGSERIAL PRIMARY KEY,
    category_id int            NOT NULL,
    name        varchar(100)   NOT NULL,
    description varchar(255),
    price       decimal(10, 2) NOT NULL,
    image_url   varchar(255),
    created_at  timestamp DEFAULT (now()),
    updated_at  timestamp DEFAULT (now())
);

CREATE INDEX IF NOT EXISTS idx_menu_items_created_at ON MenuItems (created_at);
