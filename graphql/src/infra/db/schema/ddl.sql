CREATE TABLE `account` (
  `id` bigint(20) NOT NULL AUTO_INCREMENT COMMENT 'アカウントID',
  `name` text COLLATE utf8mb4_unicode_ci NOT NULL COMMENT '名前',
  `email` text COLLATE utf8mb4_unicode_ci NOT NULL COMMENT 'メールアドレス',
  `created_at` datetime NOT NULL DEFAULT CURRENT_TIMESTAMP COMMENT '作成日時',
  `updated_at` datetime NOT NULL DEFAULT CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP COMMENT '更新日時',
  PRIMARY KEY (`id`)
) ENGINE=InnoDB AUTO_INCREMENT=25 DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_unicode_ci COMMENT='アカウント';

CREATE TABLE `account_password` (
  `account_id` bigint(20) NOT NULL COMMENT 'アカウントID',
  `hashed_password` varchar(256) COLLATE utf8mb4_unicode_ci NOT NULL COMMENT 'ハッシュ済みパスワード',
  `set_at` datetime NOT NULL COMMENT 'パスワード設定日時',
  `logged_in_at` datetime NOT NULL COMMENT 'ログイン日時',
  `created_at` datetime NOT NULL DEFAULT CURRENT_TIMESTAMP COMMENT '作成日時',
  `updated_at` datetime NOT NULL DEFAULT CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP COMMENT '更新日時',
  PRIMARY KEY (`account_id`)
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_unicode_ci COMMENT='アカウントパスワード';
