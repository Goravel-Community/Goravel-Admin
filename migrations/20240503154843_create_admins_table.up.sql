CREATE TABLE admins (
  id INTEGER PRIMARY KEY AUTOINCREMENT NOT NULL,
  created_at DATETIME NOT NULL,
  updated_at DATETIME NOT NULL,
  email VARCHAR(255) NOT NULL,
  password VARCHAR(255) NOT NULL
);
CREATE INDEX admins_email_index ON admins (email);
