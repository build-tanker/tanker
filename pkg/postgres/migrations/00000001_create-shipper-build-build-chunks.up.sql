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
  shipper BIGINT,
  bundle_id VARCHAR(128) NOT NULL UNIQUE,
  upload_complete BOOLEAN,
  migrated BOOLEAN,
  created_at TIMESTAMP default NOW(),
  updated_at TIMESTAMP DEFAULT now()
);