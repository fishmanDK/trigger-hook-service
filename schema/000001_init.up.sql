CREATE TABLE IF NOT EXISTS deletion_requests
(
    id SERIAL,
    bannerID INT UNIQUE,
    tagID INT,
    featureID INT UNIQUE,
    expires_at TIMESTAMP NOT NULL
);