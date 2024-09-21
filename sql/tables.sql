-- ranking.activity definition

CREATE TABLE `activity` (
  `id` int NOT NULL AUTO_INCREMENT,
  `name` varchar(100) DEFAULT NULL COMMENT '活动名称',
  `add_time` int DEFAULT NULL COMMENT '添加时间',
  PRIMARY KEY (`id`)
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb3;


-- ranking.player definition

CREATE TABLE `player` (
  `id` int NOT NULL AUTO_INCREMENT,
  `aid` int DEFAULT NULL,
  `ref` varchar(50) DEFAULT NULL,
  `nickname` varchar(100) DEFAULT NULL,
  `declaration` varchar(255) DEFAULT NULL,
  `avatar` varchar(255) DEFAULT NULL,
  `score` int DEFAULT NULL,
  `add_time` int DEFAULT NULL,
  `update_time` int DEFAULT NULL,
  PRIMARY KEY (`id`)
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb3;


-- ranking.`user` definition

CREATE TABLE `user` (
  `id` int NOT NULL AUTO_INCREMENT,
  `username` varchar(100) DEFAULT NULL,
  `password` varchar(50) DEFAULT NULL,
  `add_time` int DEFAULT NULL,
  `update_time` int DEFAULT NULL,
  PRIMARY KEY (`id`)
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb3;



-- ranking.vote definition

CREATE TABLE `vote` (
  `id` int NOT NULL AUTO_INCREMENT,
  `user_id` int DEFAULT NULL,
  `player_id` int DEFAULT NULL,
  `add_time` int DEFAULT NULL,
  PRIMARY KEY (`id`)
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb3;