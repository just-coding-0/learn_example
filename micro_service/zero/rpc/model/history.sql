CREATE TABLE `history`  (
  `msg` varchar(255)  NOT NULL,
  `last_echo` datetime(0) NOT NULL ,
  `times` bigint UNSIGNED NOT NULL,
  PRIMARY KEY (`msg`) USING BTREE
) ENGINE = InnoDB CHARACTER SET = utf8mb4 COLLATE = utf8mb4_general_ci ROW_FORMAT = Dynamic;