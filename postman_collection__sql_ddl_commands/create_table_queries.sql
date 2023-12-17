
-- zopsmart_task_db.trains definition

CREATE TABLE `trains` (
  `train_number` int NOT NULL,
  `name` varchar(100) DEFAULT NULL,
  `status` varchar(100) DEFAULT NULL,
  PRIMARY KEY (`train_number`)
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_0900_ai_ci;


-- zopsmart_task_db.platforms definition

CREATE TABLE `platforms` (
  `platform_number` int NOT NULL,
  `train_number` int DEFAULT '-1',
  PRIMARY KEY (`platform_number`)
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_0900_ai_ci;

