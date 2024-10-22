CREATE TABLE Status (
  id CHAR(36) PRIMARY KEY,
  name VARCHAR(255) NOT NULL
);

CREATE TABLE Category (
  id CHAR(36) PRIMARY KEY,
  name VARCHAR(255) NOT NULL
);

CREATE TABLE User (
  id CHAR(36) PRIMARY KEY,
  name VARCHAR(255) NOT NULL,
  email VARCHAR(255) NOT NULL
);

CREATE TABLE Tasks (
  id CHAR(36) PRIMARY KEY,
  title VARCHAR(255) NOT NULL,
  description VARCHAR(255),
  status_id CHAR(36) NULL,
  category_id CHAR(36) NULL,
  user_id CHAR(36) NULL,
  created_at DATETIME,
  FOREIGN KEY (status_id) REFERENCES Status(id) ON DELETE SET NULL,
  FOREIGN KEY (category_id) REFERENCES Category(id) ON DELETE SET NULL,
  FOREIGN KEY (user_id) REFERENCES User(id) ON DELETE SET NULL
);

-- Insert mock data

INSERT INTO Status (id, name) VALUES (UUID(), 'Pending'), (UUID(), 'In Progress'), (UUID(), 'Completed');
INSERT INTO Category (id, name) VALUES (UUID(), 'Work'), (UUID(), 'Personal'), (UUID(), 'Shopping');
INSERT INTO User (id, name, email) VALUES (UUID(), 'John Doe', 'john@example.com'), (UUID(), 'Jane Smith', 'jane@example.com');

-- Insert tasks with generated UUIDs for status_id, category_id, and user_id
INSERT INTO Tasks (id, title, description, status_id, category_id, user_id, created_at)
VALUES
(UUID(), 'Task 1', 'First task description', (SELECT id FROM Status WHERE name='Pending'), (SELECT id FROM Category WHERE name='Work'), (SELECT id FROM User WHERE name='John Doe'), '2024-01-12 12:00:00'),
(UUID(), 'Task 2', 'Second task description', (SELECT id FROM Status WHERE name='In Progress'), (SELECT id FROM Category WHERE name='Personal'), (SELECT id FROM User WHERE name='Jane Smith'), '2024-01-15 15:30:00'),
(UUID(), 'Task 3', 'Third task description', (SELECT id FROM Status WHERE name='Completed'), (SELECT id FROM Category WHERE name='Work'), (SELECT id FROM User WHERE name='John Doe'), '2024-02-01 09:00:00');
