INSERT INTO `system_config`
(`name`, `config`, `created_by`, `updated_by`, `created_at`, `updated_at`)
VALUES('event_updater', '{\"countries\":[\"ID\", \"MY\", \"PH\", \"SG\", \"TW\", \"TH\", \"VN\", \"BR\"]}', 'system', 'system', UNIX_TIMESTAMP(NOW())*1000, UNIX_TIMESTAMP(NOW())*1000);
