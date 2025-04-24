-- users table
CREATE TABLE if NOT EXISTS users (
    id BIGSERIAL PRIMARY KEY,
    name VARCHAR(100) NOT NULL CHECK (char_length(name) >= 2),
    password TEXT NOT NULL,
    role VARCHAR(20) NOT NULL CHECK (role IN ('receptionist', 'doctor')),
    created_at TIMESTAMPTZ NOT NULL DEFAULT NOW()
);

-- patients table 
CREATE TABLE if NOT EXISTS patients (
    id BIGSERIAL PRIMARY KEY,
    name VARCHAR(100) NOT NULL CHECK (char_length(name) >= 2),
    age INT NOT NULL CHECK (age >= 0 AND age <= 110),
    gender VARCHAR(10) NOT NULL CHECK (gender IN ('male', 'female', 'other')),
    contact TEXT NOT NULL,
    address TEXT NOT NULL,
    disease TEXT NOT NULL,
    handled_by_doctor BIGINT NOT NULL,
    updated_by BIGINT,
    updated_at TIMESTAMPTZ NOT NULL DEFAULT NOW(),
    created_at TIMESTAMPTZ NOT NULL DEFAULT NOW()
);