-- root用户下
create user 'spider'@'localhost' identified by '123456';

create database zhihuspider;

grant all on zhihuspider.* to 'spider'@'localhost';


DROP TABLE IF EXISTS `spider_user`;
CREATE TABLE `spider_user` (
  `username` varchar(255) NOT NULL,
  `name` varchar(255) NOT NULL,
  `vacation` varchar(255) DEFAULT NULL,
  `headLine` varchar(255) DEFAULT NULL,
  `sex` tinyint(1) DEFAULT NULL,
  PRIMARY KEY (`username`)
) ENGINE=InnoDB DEFAULT CHARSET=utf8 COMMENT='知乎用户表';