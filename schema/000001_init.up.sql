CREATE TABLE clients (
    id SERIAL PRIMARY KEY,
    full_name VARCHAR(100) NOT NULL,
    email VARCHAR(100) UNIQUE,
    phone VARCHAR(20)
);

CREATE TABLE policies (
    id SERIAL PRIMARY KEY,
    client_id INT REFERENCES clients(id) ON DELETE CASCADE,
    type VARCHAR(50) NOT NULL,  -- Например: "авто", "здоровье", "имущество"
    start_date DATE NOT NULL,
    end_date DATE NOT NULL,
    price DECIMAL(10, 2) NOT NULL, 
    is_active BOOLEAN DEFAULT TRUE
);

CREATE TABLE claims (
    id SERIAL PRIMARY KEY,
    policy_id INT REFERENCES policies(id),
    incident_date DATE NOT NULL,
    description TEXT NOT NULL,
    status VARCHAR(20) DEFAULT 'pending'  -- pending/approved/rejected
);