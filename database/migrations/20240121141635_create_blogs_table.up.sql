CREATE TABLE blogs (
  id bigint(20) unsigned NOT NULL AUTO_INCREMENT,
  created_at datetime(3) NOT NULL,
  updated_at datetime(3) NOT NULL,
  PRIMARY KEY (id),
  KEY idx_blogs_created_at (created_at),
  KEY idx_blogs_updated_at (updated_at)
) ENGINE = InnoDB DEFAULT CHARSET = utf8mb4;
