-- Create a `sessions` table.
CREATE TABLE sessions (
    token CHAR(43) PRIMARY KEY,
    data BLOB NOT NULL,
    expiry TIMESTAMP(6) NOT NULL
);

-- Add an index on the expiry column.
CREATE INDEX sessions_expiry_idx ON sessions(expiry);