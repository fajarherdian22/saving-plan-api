CREATE TABLE `users` (
  `id` VARCHAR(255) PRIMARY KEY NOT NULL,
  `email` VARCHAR(255) UNIQUE NOT NULL,
  `created_at` DATETIME NOT NULL DEFAULT NOW()
);

CREATE TABLE `saving_goals` (
  `id` VARCHAR(255) PRIMARY KEY NOT NULL,
  `user_id` VARCHAR(255) NOT NULL,
  `name` VARCHAR(255) NOT NULL,
  `target_amount` DOUBLE NOT NULL,
  `current_amount` DOUBLE NOT NULL DEFAULT 0,
  `target_time` DATE NOT NULL,
  `created_at` DATETIME NOT NULL DEFAULT NOW(),
  FOREIGN KEY (`user_id`) REFERENCES `users`(`id`) ON DELETE CASCADE
);

CREATE TABLE `transactions_record` (
  `id` VARCHAR(255) PRIMARY KEY NOT NULL,
  `user_id` VARCHAR(255) NOT NULL,
  `goal_id` VARCHAR(255) NULL,
  `amount` DOUBLE NOT NULL,
  `type` ENUM('income', 'expense') NOT NULL,
  `category` VARCHAR(255) NOT NULL,
  `name` VARCHAR(255) NOT NULL,
  `created_at` DATETIME NOT NULL DEFAULT NOW(),
  FOREIGN KEY (`user_id`) REFERENCES `users`(`id`) ON DELETE CASCADE,
  FOREIGN KEY (`goal_id`) REFERENCES `saving_goals`(`id`) ON DELETE SET NULL
);
