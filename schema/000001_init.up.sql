CREATE TABLE IF NOT EXISTS deletion_requests
(
    bannerID INT UNIQUE,
    tagID INT,
    featureID INT UNIQUE,
    expires_at DATE NOT NULL
);