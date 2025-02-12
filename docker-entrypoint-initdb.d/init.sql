CREATE TABLE IF NOT EXISTS vettes (
    id SERIAL PRIMARY KEY,
    date DATE NOT NULL,
    user_id VARCHAR(255) NOT NULL,
    year SMALLINT CHECK (year BETWEEN 1953 AND EXTRACT(YEAR FROM CURRENT_DATE)) NOT NULL,
    miles INTEGER NOT NULL,
    cost DECIMAL(10,2) NOT NULL,
    transmission_type VARCHAR(255) NOT NULL,
    exterior_color VARCHAR(255) NOT NULL,
    interior_color VARCHAR(255) NOT NULL,
    submodel VARCHAR(255) NOT NULL,
    trim VARCHAR(255) NOT NULL,
    packages TEXT[] NOT NULL,
    link TEXT NOT NULL
);

-- Insert sample data
INSERT INTO vettes (
    date, 
    user_id, 
    year, 
    miles, 
    cost, 
    transmission_type, 
    exterior_color, 
    interior_color, 
    submodel, 
    trim, 
    packages, 
    link
) VALUES
(
    '2023-01-15', 
    'user123', 
    2020, 
    15000, 
    65000.00, 
    'Automatic', 
    'Torch Red', 
    'Jet Black', 
    'Stingray', 
    'Z51', 
    ARRAY['MRC', 'NPP', 'PDR'], 
    'https://example.com/vette1'
),
(
    '2023-02-20', 
    'user456', 
    2021, 
    8500, 
    72000.00, 
    'Manual', 
    'Arctic White', 
    'Natural', 
    'Z06', 
    '3LZ', 
    ARRAY['MRC', 'NPP',], 
    'https://example.com/vette2'
),
(
    '2023-03-10', 
    'user789', 
    2019, 
    25000, 
    58000.00, 
    'Manual', 
    'Black', 
    'Red', 
    'Grand Sport', 
    '2LT', 
    ARRAY[], 
    'https://example.com/vette3'
);
