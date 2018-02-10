CREATE TABLE shippers (
  id BIGSERIAL NOT NULL PRIMARY KEY,
  access_key VARCHAR(128) NOT NULL UNIQUE,
  name VARCHAR(256),
  machine_name VARCHAR(256),
  created_at TIMESTAMP default NOW(),
  updated_at TIMESTAMP DEFAULT now()
);

CREATE TABLE builds (
  id BIGSERIAL NOT NULL PRIMARY KEY,
  shipper VARCHAR(128),
  bundle_id VARCHAR(128),
  upload_complete BOOLEAN,
  migrated BOOLEAN,
  created_at TIMESTAMP default NOW(),
  updated_at TIMESTAMP DEFAULT now()
);