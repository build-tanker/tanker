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
  app_group UUID,
  expiry TIMESTAMP,
  deleted BOOLEAN DEFAULT false,
  created_at TIMESTAMP DEFAULT NOW(),
  updated_at TIMESTAMP DEFAULT NOW()
);

CREATE TABLE build (
  id UUID NOT NULL PRIMARY KEY,
  file_name VARCHAR(256),
  shipper UUID REFERENCES shipper(id),
  bundle_id VARCHAR(128),
  platform VARCHAR(128),
  extension VARCHAR(128),
  upload_complete BOOLEAN DEFAULT false,
  deleted BOOLEAN DEFAULT false,
  created_at TIMESTAMP DEFAULT NOW(),
  updated_at TIMESTAMP DEFAULT NOW()
);