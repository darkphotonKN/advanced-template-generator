-- Create items table
CREATE TABLE items (
    id UUID PRIMARY KEY DEFAULT uuid_generate_v4(),
    name VARCHAR(255) NOT NULL,
    description TEXT,
    created_at TIMESTAMP NOT NULL DEFAULT NOW(),
    updated_at TIMESTAMP NOT NULL DEFAULT NOW()
);

-- Create indexes
CREATE INDEX idx_items_created_at ON items(created_at);
CREATE INDEX idx_items_name ON items(name);

-- Add updated_at trigger
CREATE TRIGGER set_timestamp
BEFORE UPDATE ON items
FOR EACH ROW
EXECUTE FUNCTION trigger_set_timestamp();