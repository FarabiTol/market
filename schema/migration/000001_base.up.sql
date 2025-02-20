CREATE TABLE IF NOT EXISTS collection
(
    "id"                UUID PRIMARY KEY            NOT NULL,
    "user_id"           UUID UNIQUE                 NOT NULL,
    "short_name"        varchar(100)                NOT NULL,
    "full_name"         varchar(150)                NOT NULL,
    "status"            varchar(30)                 NOT NULL,
    "amount"            numeric(18, 2)              NOT NULL,
    "currency"          varchar(250)                NOT NULL,
    "register_date"     timestamp without time zone NOT NULL,
    "created_at"        timestamp without time zone NOT NULL,
    "updated_at"        timestamp without time zone NOT NULL
    );

CREATE TABLE IF NOT EXISTS users
(
    "id"                    UUID PRIMARY KEY            NOT NULL,
    "account_id"            UUID UNIQUE                 NOT NULL,
    "first_name"            varchar(100)                NOT NULL,
    "last_name"             varchar(150)                NOT NULL,
    "phone_number"          varchar(255)                NOT NULL,
    "email"                 varchar(50)                 NOT NULL,
    );

CREATE INDEX IF NOT EXISTS idx_collection_id ON collection (id);