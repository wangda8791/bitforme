-- +goose Up
-- SQL in this section is executed when the migration is applied.
CREATE TABLE `payment_transactions` (
  `id` int(11) NOT NULL AUTO_INCREMENT,
  `txid` varchar(255) DEFAULT NULL,
  `amount` decimal(32,16) DEFAULT NULL,
  `confirmations` int(11) DEFAULT NULL,
  `address` varchar(255) DEFAULT NULL,
  `receive_at` datetime NOT NULL DEFAULT CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP,
  `currency` int(11) DEFAULT NULL,
  `txout` int(11) DEFAULT NULL,
  `created_at` datetime NOT NULL DEFAULT CURRENT_TIMESTAMP,
  `updated_at` datetime NOT NULL DEFAULT CURRENT_TIMESTAMP,
  `deleted_at` datetime DEFAULT NULL,
  `state` varchar(255) DEFAULT NULL,
  PRIMARY KEY (`id`)
) ENGINE=InnoDB AUTO_INCREMENT=1 DEFAULT CHARSET=utf8;
-- +goose Down
-- SQL in this section is executed when the migration is rolled back.
DROP TABLE payment_transactions;