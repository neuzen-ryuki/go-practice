CREATE TABLE users (
  id int unsigned PRIMARY KEY AUTO_INCREMENT,
  uuid varchar(64) NOT NULL UNIQUE,
  name varchar(255),
  email varchar(255) NOT NULL UNIQUE,
  password varchar(255) NOT NULL,
  created_at timestamp NOT NULL DEFAULT CURRENT_TIMESTAMP,
  updated_at timestamp NOT NULL DEFAULT CURRENT_TIMESTAMP
) ENGINE=InnoDB DEFAULT CHARSET=utf8;


CREATE TABLE sessions (
  id int unsigned PRIMARY KEY AUTO_INCREMENT,
  user_id int unsigned,
  uuid varchar(64) NOT NULL UNIQUE,
  email varchar(255),
  created_at timestamp NOT NULL DEFAULT CURRENT_TIMESTAMP,
  updated_at timestamp NOT NULL DEFAULT CURRENT_TIMESTAMP,
  CONSTRAINT fk_user_sessions
    FOREIGN KEY (user_id)
    REFERENCES users(id)
    ON DELETE SET NULL ON UPDATE CASCADE
) ENGINE=InnoDB DEFAULT CHARSET=utf8;


CREATE TABLE threads (
  id int unsigned PRIMARY KEY AUTO_INCREMENT,
  user_id int unsigned,
  uuid varchar(64) NOT NULL UNIQUE,
  topic varchar(255),
  created_at timestamp NOT NULL DEFAULT CURRENT_TIMESTAMP,
  updated_at timestamp NOT NULL DEFAULT CURRENT_TIMESTAMP,
  CONSTRAINT fk_user_threads
    FOREIGN KEY (user_id)
    REFERENCES users(id)
    ON DELETE SET NULL ON UPDATE CASCADE
) ENGINE=InnoDB DEFAULT CHARSET=utf8;


CREATE TABLE posts (
  id int unsigned PRIMARY KEY AUTO_INCREMENT,
  user_id int unsigned,
  thread_id int unsigned,
  uuid varchar(64) NOT NULL UNIQUE,
  body varchar(255),
  created_at timestamp NOT NULL DEFAULT CURRENT_TIMESTAMP,
  updated_at timestamp NOT NULL DEFAULT CURRENT_TIMESTAMP,
  CONSTRAINT fk_user_posts
    FOREIGN KEY (user_id)
    REFERENCES users(id)
    ON DELETE SET NULL ON UPDATE CASCADE,
  CONSTRAINT fk_thread_posts
    FOREIGN KEY (thread_id)
    REFERENCES threads(id)
    ON DELETE SET NULL ON UPDATE CASCADE
) ENGINE=InnoDB DEFAULT CHARSET=utf8;

