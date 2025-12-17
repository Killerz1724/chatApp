CREATE TABLE users(
  id BIGSERIAL PRIMARY KEY,
  email VARCHAR(255) UNIQUE NOT NULL,
  password VARCHAR(255) NOT NULL,
  username VARCHAR(255) NOT NULL,
  tag VARCHAR(255) NOT NULL,
  img VARCHAR DEFAULT '',
  created_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP,
  updated_at TIMESTAMP,
  deleted_at TIMESTAMP
);

CREATE TABLE friends(
  id BIGSERIAL PRIMARY KEY,
  user_id BIGINT NOT NULL REFERENCES users(id),
  friend_id BIGINT NOT NULL REFERENCES users(id),
  created_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP,
  updated_at TIMESTAMP,
  deleted_at TIMESTAMP,
  UNIQUE(user_id, friend_id)
);


CREATE TABLE conversations(
  id BIGSERIAL PRIMARY KEY,
  name VARCHAR(255),
  is_group BOOLEAN DEFAULT FALSE,
  created_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP,
  updated_at TIMESTAMP,
  deleted_at TIMESTAMP
);

CREATE TYPE role AS ENUM('admin', 'member');

CREATE TABLE conversationsMembers(
  id BIGSERIAL PRIMARY KEY,
  conversation_id BIGINT NOT NULL REFERENCES conversations(id),
  user_id BIGINT NOT NULL REFERENCES users(id),
  role ROLE DEFAULT 'member',
  created_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP,
  updated_at TIMESTAMP,
  deleted_at TIMESTAMP,
  UNIQUE(conversation_id, user_id)
);

CREATE TYPE attachment AS ENUM('text', 'image', 'video', 'audio', 'file');

CREATE TABLE messages(
  id BIGSERIAL PRIMARY KEY,
  sender_id BIGINT NOT NULL REFERENCES users(id),
  conversation_id BIGINT NOT NULL REFERENCES conversations(id),
  content TEXT NOT NULL,
  attachment ATTACHMENT DEFAULT 'text',
  created_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP,
  updated_at TIMESTAMP,
  deleted_at TIMESTAMP
);

INSERT INTO users(email, password, username, tag)
VALUES 
  ('user1@mail.com', '$2a$12$EAHF4JtPe/HKq6azH6Q9gedPfH50B2/qwDUEqSAT3nDl0OHeifPCC', 'user1', '#1111'),
  ('user2@mail.com', '$2a$12$EAHF4JtPe/HKq6azH6Q9gedPfH50B2/qwDUEqSAT3nDl0OHeifPCC', 'user2', '#2222'),
  ('user3@mail.com', '$2a$12$EAHF4JtPe/HKq6azH6Q9gedPfH50B2/qwDUEqSAT3nDl0OHeifPCC', 'user3', '#3333'),
  ('user4@mail.com', '$2a$12$EAHF4JtPe/HKq6azH6Q9gedPfH50B2/qwDUEqSAT3nDl0OHeifPCC', 'user4', '#4444'),
  ('user5@mail.com', '$2a$12$EAHF4JtPe/HKq6azH6Q9gedPfH50B2/qwDUEqSAT3nDl0OHeifPCC', 'user5', '#5555'),
  ('user6@mail.com', '$2a$12$EAHF4JtPe/HKq6azH6Q9gedPfH50B2/qwDUEqSAT3nDl0OHeifPCC', 'user6', '#6666');

INSERT INTO friends(user_id, friend_id)
VALUES
  (1, 2),
  (1, 3),
  (1, 4),
  (2, 3),
  (3, 4),
  (4, 5),
  (5, 6);

INSERT INTO conversations(name, is_group)
VALUES
  ('', true),
  ('group2', true);

INSERT INTO conversationsMembers(conversation_id, user_id)
VALUES
  (1, 1),
  (1, 2),
  (2, 1),
  (2, 2),
  (2, 3);

INSERT INTO messages(sender_id, conversation_id, content)
VALUES
  (1, 1, 'hello'),
  (2, 1, 'world'),
  (1, 2, 'hello'),
  (2, 2, 'world');