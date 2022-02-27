CREATE OR REPLACE FUNCTION update_time_column()
    RETURNS TRIGGER AS $$
BEGIN
    NEW.update_time = now();
    RETURN NEW;
END;
$$ language 'plpgsql';

CREATE TABLE IF NOT EXISTS addresses (
        id serial NOT NULL PRIMARY KEY,
        city text NOT NULL,
        street text NOT NULL,
        house text NOT NULL
);

CREATE TABLE IF NOT EXISTS pharmacy (
        id serial NOT NULL PRIMARY KEY,
        address_id int NOT NULL REFERENCES addresses(id),
        name text NOT NULL,
        creation_time timestamp NOT NULL DEFAULT NOW(),
        update_time timestamp NOT NULL DEFAULT NOW()
);

CREATE TRIGGER update_pharmacy_time BEFORE UPDATE
    ON pharmacy FOR EACH ROW EXECUTE PROCEDURE
    update_time_column();

CREATE TABLE IF NOT EXISTS  rights(
        id serial NOT NULL PRIMARY KEY,
        comment text NOT NULL DEFAULT '',
        creation_time timestamp NOT NULL DEFAULT NOW(),
        update_time timestamp NOT NULL DEFAULT NOW()
);
COMMENT ON table rights is 'Права пользователей с их описанием';

CREATE TRIGGER update_rights_update_time BEFORE UPDATE
    ON rights FOR EACH ROW EXECUTE PROCEDURE
    update_time_column();

CREATE TABLE IF NOT EXISTS users (
        id serial NOT NULL PRIMARY KEY,
        name text NOT NULL DEFAULT '',
        surname text NOT NULL DEFAULT '',
        email text NOT NULL DEFAULT '',
        password_hash text NOT NULL,
        default_pharmacy_id int REFERENCES pharmacy(id),
        creation_time timestamp NOT NULL DEFAULT NOW(),
        update_time timestamp NOT NULL DEFAULT NOW()
);

COMMENT ON table users is 'Пользователи админки аптеки с определенными правами';

CREATE TRIGGER update_users_update_time BEFORE UPDATE
    ON users FOR EACH ROW EXECUTE PROCEDURE
    update_time_column();

CREATE TABLE IF NOT EXISTS users_rights (
        id serial PRIMARY KEY NOT NULL,
        user_id int NOT NULL REFERENCES users(id),
        right_id int NOT NULL REFERENCES rights(id),
        authorized_user_id int REFERENCES users(id),
        is_active boolean NOT NULL DEFAULT true,
        comment text NOT NULL DEFAULT '',
        creation_time timestamp NOT NULL DEFAULT NOW(),
        update_time timestamp NOT NULL DEFAULT NOW(),
        UNIQUE (user_id, right_id)
);

CREATE TRIGGER update_users_rights_update_time BEFORE UPDATE
    ON users_rights FOR EACH ROW EXECUTE PROCEDURE
    update_time_column();

CREATE TABLE IF NOT EXISTS rights_group (
        id serial PRIMARY KEY NOT NULL,
        rights int[] NOT NULL DEFAULT '{}'::int[],
        last_edit_user int NOT NULL,
        creation_time timestamp NOT NULL DEFAULT NOW(),
        update_time timestamp NOT NULL DEFAULT NOW()
);

CREATE TRIGGER update_rights_group_update_time BEFORE UPDATE
    ON rights_group FOR EACH ROW EXECUTE PROCEDURE
    update_time_column();

CREATE TABLE IF NOT EXISTS product (
        id serial NOT NULL PRIMARY KEY,
        price int NOT NULL,
        expiration_date interval NOT NULL,
        instruction_url text NOT NULL DEFAULT '',
        img_url text NOT NULL DEFAULT '',
        comment text NOT NULL DEFAULT '',
        recipe_only boolean NOT NULL DEFAULT false,
        creation_time timestamp NOT NULL DEFAULT NOW(),
        update_time timestamp NOT NULL DEFAULT NOW()
);
CREATE TRIGGER update_product_update_time BEFORE UPDATE
    ON product FOR EACH ROW EXECUTE PROCEDURE
    update_time_column();

CREATE TABLE IF NOT EXISTS receipt (
        id bigserial NOT NULL PRIMARY KEY,
        user_id int NOT NULL REFERENCES users(id),
        pharmacy_id int NOT NULL REFERENCES pharmacy(id),
        sum double precision NOT NULL,
        discount int NOT NULL DEFAULT 0,
        creation_time timestamp NOT NULL DEFAULT NOW(),
        update_time timestamp NOT NULL DEFAULT NOW()
);

CREATE TRIGGER update_receipt_update_time BEFORE UPDATE
    ON receipt FOR EACH ROW EXECUTE PROCEDURE
    update_time_column();

CREATE TABLE IF NOT EXISTS product_item (
        id bigserial NOT NULL PRIMARY KEY,
        item_description_id int NOT NULL REFERENCES product(id),
        receipt_id bigint REFERENCES receipt(id),
        pharmacy_id int NOT NULL REFERENCES pharmacy(id),
        position text NOT NULL DEFAULT '',
        manufactured_time timestamp NOT NULL,
        reservation uuid,
        is_sold boolean NOT NULL DEFAULT false,
        is_expired boolean NOT NULL DEFAULT false,
        priority int NOT NULL DEFAULT 0,
        creation_time timestamp NOT NULL DEFAULT NOW(),
        update_time timestamp NOT NULL DEFAULT NOW()
);

COMMENT ON column product_item.reservation IS '
    uuid резервации. Резервация сохраняется в nosql хранилище для быстроты.
    Раз в 30 минут проверяется резервация. Если она была отменена, то в sql это поле станет null';

CREATE TRIGGER update_product_item_update_time BEFORE UPDATE
    ON product_item FOR EACH ROW EXECUTE PROCEDURE
    update_time_column();

CREATE TABLE IF NOT EXISTS complaints (
        id serial NOT NULL PRIMARY KEY,
        name text NOT NULL DEFAULT '',
        email text NOT NULL DEFAULT '',
        complaint text NOT NULL DEFAULT '',
        worker_name text NOT NULL DEFAULT '',
        pharmacy_id int REFERENCES pharmacy(id)
)

