DROP TABLE IF EXISTS `xingzuo`;
CREATE TABLE `xingzuo` (
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

SELECT COUNT(*) FROM xingzuo







-- INSERT INTO `xingzuo` (`title`,`cover`,`author`,`query`,`desc`,`content`,`content_time`,`create_time`) VALUES ('2021年下半年金牛座婚姻运如何 旺不旺','https://img.d1xz.net/d/2021/05/60b39c469370e.jpg_art','','第一星座网','婚姻对于每个人来说，都有着非同寻常的寓意，这是爱情最好的见证，也是一段感情有始有终的结局。想要了解每个人的婚姻运势如何，可以先从了解这个人的星座开始。在2021年的下半年中，十二星座当中的金牛座，会有...','<p>　　婚姻对于每个人来说，都有着非同寻常的寓意，这是爱情最好的见证，也是一段感情有始有终的结局。想要了解每个人的婚姻运势如何，可以先从了解这个人的星座开始。在2021年的下半年中，十二星座当中的金牛座，会有怎样的婚姻运势呢？<br/></p><p style="text-align: center;"><img src="https://img.d1xz.net/d/2021/05/60b39c469370e.jpg_art"/></p><p>　<strong>　金牛座2021年婚姻运如何</strong></p><p>　　2021年的金牛，在婚姻面前可能会面临一些考验，所以金牛需要提前做好准备。尤其是水逆结束之后，对于已婚的金牛座来说是件好事儿，夫妻之间的感情更加牢固;生活中彼此相互体谅，相互扶持。2021年金牛座进入下一阶段的几率非常高，可以期待一下新生命的到来，最适合备孕， 成功的可能性是很大的。</p><p>　　<strong>桃花运势发展</strong></p><p>　　2021年整体上对于单身的金牛座这一年身边桃花不断，但总是分分合合，感情非常不稳定;一点小事情都会掀起巨浪;目前最好保持着恋人的关系，不适合进入婚姻的状态;不要因为年龄而着急忙慌地结婚，不然未来的婚姻生活会成为生活当中的难题。</p><p style="text-align: center;"><img src="https://img.d1xz.net/d/2021/05/60b39c4e71d4a.jpg_art"/></p><p>　<strong>　如何提升感情运势</strong></p><p>　　金牛座的人在日常参与社交的时候，切勿表现的太过于沉默或者高冷，就算面对自己不感兴趣的话题也要礼貌微笑，不要将情绪直白的表露在脸上。若是相亲约会各种细节更是要注意，外在形象的管理要跟上，可以在穿衣打扮上下点功夫。结账的时候女生不要强调AA，男生主动帮相亲对象买单，总之不要表现的过分在乎金钱，否则最后的结局不太好。单身的金牛座今年可佩戴一串红豆增情饰品作为爱情吉祥物，寓意提升自信，增添魅力，期盼早日遇到良缘。这一年单身的金牛座还是要做到学习如何与异性相处，积极培养好感对象，不要着急脱单，否则反而不会有好的结果，会变得很失望。</p><p><br/></p>','2021-05-30 22:08:00','2021-05-31 16:34:15')
