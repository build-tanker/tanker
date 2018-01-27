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
  bundle_id VARCHAR(128) NOT NULL UNIQUE,
  size int,
  checksum VARCHAR(256),
  upload_complete BOOLEAN,
  migrated BOOLEAN,
  created_at TIMESTAMP default NOW(),
  updated_at TIMESTAMP DEFAULT now()
);

CREATE TABLE build_chunks (
  id BIGSERIAL NOT NULL PRIMARY KEY,
  build_id BIGINT,
  upload_url VARCHAR(256),
  disk_path VARCHAR(256),
  checksum VARCHAR(256),
  upload_complete BOOLEAN,
  created_at TIMESTAMP default NOW(),
  updated_at TIMESTAMP DEFAULT now()
);