CREATE TABLE `dy_xingzuo` (
  `id` int(11) NOT NULL AUTO_INCREMENT,
  `title` varchar(255) DEFAULT NULL COMMENT '标题',
  `author` varchar(255) DEFAULT NULL COMMENT '作者',
  `cover` varchar(255) DEFAULT NULL COMMENT '封面',
  `desc` varchar(255) DEFAULT NULL COMMENT '简介',
  `content` text NOT NULL COMMENT '内容',
  `query` varchar(255) DEFAULT NULL COMMENT '来源',
  `content_time` varchar(20) DEFAULT NULL COMMENT '内容时间',
  `create_time` datetime DEFAULT NULL COMMENT '创建时间',
  PRIMARY KEY (`id`) USING BTREE
) ENGINE=InnoDB DEFAULT CHARSET=utf8 ROW_FORMAT=DYNAMIC;

CREATE TABLE `shi_xingzuo` (
  `id` int(11) NOT NULL AUTO_INCREMENT,
  `title` varchar(255) DEFAULT NULL COMMENT '标题',
  `author` varchar(255) DEFAULT NULL COMMENT '作者',
  `cover` varchar(255) DEFAULT NULL COMMENT '封面',
  `desc` varchar(255) DEFAULT NULL COMMENT '简介',
  `content` text NOT NULL COMMENT '内容',
  `query` varchar(255) DEFAULT NULL COMMENT '来源',
  `content_time` varchar(20) DEFAULT NULL COMMENT '内容时间',
  `create_time` datetime DEFAULT NULL COMMENT '创建时间',
  PRIMARY KEY (`id`) USING BTREE
) ENGINE=InnoDB DEFAULT CHARSET=utf8 ROW_FORMAT=DYNAMIC;

SELECT COUNT(*) FROM dy_xingzuo
