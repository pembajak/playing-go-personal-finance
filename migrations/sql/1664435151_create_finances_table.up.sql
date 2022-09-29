CREATE TABLE IF NOT EXISTS `finances` (
  `id` int NOT NULL AUTO_INCREMENT,
  `title` varchar(255) DEFAULT NULL,
  `account_id` int NOT NULL,
  `amount` decimal(12,2) DEFAULT NULL,
  `description` text,
  `user_id` int NOT NULL,
  `type` varchar(255) DEFAULT NULL,
  `transaction_date` datetime DEFAULT NULL,
  `created_at` timestamp NULL DEFAULT CURRENT_TIMESTAMP,
  `updated_at` timestamp NULL DEFAULT CURRENT_TIMESTAMP,
  `deleted_at` timestamp NULL DEFAULT NULL,
  PRIMARY KEY (`id`),
  KEY `finance_account` (`account_id`),
  KEY `finance_user` (`user_id`),
  CONSTRAINT `finance_account` FOREIGN KEY (`account_id`) REFERENCES `accounts` (`id`),
  CONSTRAINT `finance_user` FOREIGN KEY (`user_id`) REFERENCES `users` (`id`)
)