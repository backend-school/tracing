CREATE TABLE IF NOT EXISTS token
(
  id          UUID PRIMARY KEY NOT NULL,
  name        VARCHAR(100)       NOT NULL,
  network_id  INT                NOT NULL,
  currency_id INTEGER            NOT NULL,
  is_active   BOOLEAN            NOT NULL
);