CREATE EXTENSION IF NOT EXISTS pgcrypto;
CREATE TABLE products (
          id UUID PRIMARY KEY DEFAULT gen_random_uuid(),
          title varchar(255),
          description varchar(255),
          rating int,
          image varchar(255),
          created_at TIMESTAMPTZ NOT NULL DEFAULT NOW(),
          updated_at TIMESTAMPTZ,
          deleted_at TIMESTAMPTZ
);