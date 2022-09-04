CREATE TABLE `system_config` (
  `name` varchar(20) NOT NULL,
  `config` TEXT NOT NULL,
  `created_by` varchar(50) NOT NULL DEFAULT '',
  `updated_by` varchar(50) NOT NULL DEFAULT '',
  `created_at` bigint NOT NULL DEFAULT '0',
  `updated_at` bigint NOT NULL DEFAULT '0',
  PRIMARY KEY (`name`)
);