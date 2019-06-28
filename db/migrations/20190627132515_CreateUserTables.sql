
-- +goose Up
-- SQL in section 'Up' is executed when this migration is applied

CREATE TABLE IF NOT EXISTS `user` (
  `id` INT UNSIGNED NOT NULL AUTO_INCREMENT,
  `name` VARCHAR(255) NOT NULL comment '名前',
  PRIMARY KEY (`id`)
)ENGINE = InnoDB
DEFAULT CHARSET=utf8mb4
comment='ユーザーのマスター情報';

CREATE TABLE IF NOT EXISTS `user_counter` (
  `user_id` INT UNSIGNED NOT NULL AUTO_INCREMENT,
  `reservation_count` INT UNSIGNED NOT NULL comment '予約回数',
  PRIMARY KEY (`user_id`),
  FOREIGN KEY (`user_id`) REFERENCES `user` (`id`) ON DELETE CASCADE
  )ENGINE = InnoDB
DEFAULT CHARSET=utf8mb4
comment='ユーザーの各種カウンター';

-- +goose Down
-- SQL section 'Down' is executed when this migration is rolled back

DROP TABLE `user_counter`;
DROP TABLE `user`;
