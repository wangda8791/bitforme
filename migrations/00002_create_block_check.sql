-- +goose Up
-- SQL in this section is executed when the migration is applied.
CREATE TABLE `block_checks` (
  `id` int(11) NOT NULL AUTO_INCREMENT,
  `currency` int(11) DEFAULT NULL,
  `last` int(11) DEFAULT NULL,
  `created_at` datetime NOT NULL DEFAULT CURRENT_TIMESTAMP,
  `updated_at` datetime NOT NULL DEFAULT CURRENT_TIMESTAMP,
  `deleted_at` datetime DEFAULT NULL,
  
  PRIMARY KEY (`id`)
) ENGINE=InnoDB AUTO_INCREMENT=1 DEFAULT CHARSET=utf8;
-- +goose Down
-- SQL in this section is executed when the migration is rolled back.
DROP TABLE block_checks;