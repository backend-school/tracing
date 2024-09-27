CREATE TABLE IF NOT EXISTS token
(
  id          UUID PRIMARY KEY  ,
  name        VARCHAR(100)      ,
  network_id  INT               ,
  currency_id INTEGER           ,
  is_active   BOOLEAN           
);