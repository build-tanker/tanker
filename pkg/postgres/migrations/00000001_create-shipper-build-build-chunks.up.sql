CREATE TABLE person (
  id UUID NOT NULL PRIMARY KEY,
  source VARCHAR(128),
  name VARCHAR(128),
  email VARCHAR(128),
  picture_url VARCHAR(512),
  deleted BOOLEAN DEFAULT false,
  created_at TIMESTAMP DEFAULT NOW(),
  updated_at TIMESTAMP DEFAULT NOW()
);

CREATE TABLE token (
  id UUID NOT NULL PRIMARY KEY,
  person UUID NOT NULL REFERENCES person(id),
  source VARCHAR(128),
  access_token VARCHAR(256),
  refresh_token VARCHAR(256),
  expires_in INT,
  token_type VARCHAR(128),
  created_at TIMESTAMP DEFAULT NOW(),
  updated_at TIMESTAMP DEFAULT NOW()
);

CREATE TABLE app_group (
  id UUID NOT NULL PRIMARY KEY,
  name VARCHAR(128),
  image_url VARCHAR(512),
  deleted BOOLEAN DEFAULT false,
  created_at TIMESTAMP DEFAULT NOW(),
  updated_at TIMESTAMP DEFAULT NOW()
);

CREATE TABLE app (
  id UUID NOT NULL PRIMARY KEY,
  app_group UUID NOT NULL REFERENCES app_group(id),
  name VARCHAR(128),
  bundle_id VARCHAR(128),
  platform VARCHAR(128),
  deleted BOOLEAN DEFAULT false,
  created_at TIMESTAMP DEFAULT NOW(),
  updated_at TIMESTAMP DEFAULT NOW()
);

CREATE TABLE access (
  id UUID NOT NULL PRIMARY KEY,
  person UUID NOT NULL REFERENCES person(id),
  app_group UUID REFERENCES app_group(id),
  app UUID REFERENCES app(id),
  access_level VARCHAR(16) DEFAULT 'normal',
  access_given_by UUID REFERENCES person(id),
  deleted BOOLEAN DEFAULT false,
  created_at TIMESTAMP DEFAULT NOW(),
  updated_at TIMESTAMP DEFAULT NOW()
);

CREATE TABLE shipper (
  id UUID NOT NULL PRIMARY KEY,
  app_group UUID REFERENCES app_group(id),
  expiry TIMESTAMP,
  deleted BOOLEAN DEFAULT false,
  created_at TIMESTAMP DEFAULT NOW(),
  updated_at TIMESTAMP DEFAULT NOW()
);

CREATE TABLE build (
  id UUID NOT NULL PRIMARY KEY,
  file_name VARCHAR(256),
  shipper_access_key UUID REFERENCES shipper(id),
  bundle_id VARCHAR(128),
  upload_complete BOOLEAN,
  deleted BOOLEAN DEFAULT false,
  created_at TIMESTAMP DEFAULT NOW(),
  updated_at TIMESTAMP DEFAULT NOW()
);