-- MySQL dump 10.13  Distrib 8.0.22, for Win64 (x86_64)
--
-- Host: localhost    Database: tiktok
-- ------------------------------------------------------
-- Server version	8.0.22

/*!40101 SET @OLD_CHARACTER_SET_CLIENT=@@CHARACTER_SET_CLIENT */;
/*!40101 SET @OLD_CHARACTER_SET_RESULTS=@@CHARACTER_SET_RESULTS */;
/*!40101 SET @OLD_COLLATION_CONNECTION=@@COLLATION_CONNECTION */;
/*!50503 SET NAMES utf8 */;
/*!40103 SET @OLD_TIME_ZONE=@@TIME_ZONE */;
/*!40103 SET TIME_ZONE='+00:00' */;
/*!40014 SET @OLD_UNIQUE_CHECKS=@@UNIQUE_CHECKS, UNIQUE_CHECKS=0 */;
/*!40014 SET @OLD_FOREIGN_KEY_CHECKS=@@FOREIGN_KEY_CHECKS, FOREIGN_KEY_CHECKS=0 */;
/*!40101 SET @OLD_SQL_MODE=@@SQL_MODE, SQL_MODE='NO_AUTO_VALUE_ON_ZERO' */;
/*!40111 SET @OLD_SQL_NOTES=@@SQL_NOTES, SQL_NOTES=0 */;

--
-- Table structure for table `comments`
--

DROP TABLE IF EXISTS `comments`;
/*!40101 SET @saved_cs_client     = @@character_set_client */;
/*!50503 SET character_set_client = utf8mb4 */;
CREATE TABLE `comments` (
  `comment_id` int NOT NULL AUTO_INCREMENT,
  `user_id` int NOT NULL,
  `video_id` int NOT NULL,
  `content` varchar(200) NOT NULL,
  `create_time` datetime DEFAULT CURRENT_TIMESTAMP,
  PRIMARY KEY (`comment_id`),
  KEY `videoId` (`video_id`,`create_time` DESC)
) ENGINE=InnoDB AUTO_INCREMENT=28 DEFAULT CHARSET=utf8;
/*!40101 SET character_set_client = @saved_cs_client */;

--
-- Dumping data for table `comments`
--

LOCK TABLES `comments` WRITE;
/*!40000 ALTER TABLE `comments` DISABLE KEYS */;
INSERT INTO `comments` VALUES (1,3,1,'熊的力量！','2022-05-27 10:23:56'),(2,4,1,'毛子狂喜','2022-05-27 10:35:27'),(3,1,19,'七都！','2022-05-28 22:45:12'),(4,1,19,'新结局！','2022-05-28 23:07:39'),(6,19,19,'你删不掉我！','2022-05-28 23:39:30'),(7,1,19,'好耶！','2022-05-28 23:40:15'),(8,1,18,'又是BE呢QAQ','2022-06-03 18:28:43'),(10,19,27,'太可爱了吧','2022-06-18 11:37:58'),(15,4,33,'这只金渐层也太乖了吧','2022-06-18 22:11:56'),(16,1,33,'真的好可爱啊','2022-06-18 22:20:29'),(17,4,33,'小猫咪你喜欢什么颜色的麻袋呀','2022-06-18 22:21:24'),(27,19,33,'什么时候我才能有猫呢','2022-06-19 02:36:06');
/*!40000 ALTER TABLE `comments` ENABLE KEYS */;
UNLOCK TABLES;

--
-- Table structure for table `favorites`
--

DROP TABLE IF EXISTS `favorites`;
/*!40101 SET @saved_cs_client     = @@character_set_client */;
/*!50503 SET character_set_client = utf8mb4 */;
CREATE TABLE `favorites` (
  `id` int NOT NULL AUTO_INCREMENT,
  `user_id` int NOT NULL,
  `video_id` int NOT NULL,
  PRIMARY KEY (`id`),
  KEY `videoId` (`video_id`),
  KEY `userId&videoId` (`user_id`,`video_id`)
) ENGINE=InnoDB AUTO_INCREMENT=42 DEFAULT CHARSET=utf8;
/*!40101 SET character_set_client = @saved_cs_client */;

--
-- Dumping data for table `favorites`
--

LOCK TABLES `favorites` WRITE;
/*!40000 ALTER TABLE `favorites` DISABLE KEYS */;
INSERT INTO `favorites` VALUES (34,0,0),(17,1,12),(16,1,18),(10,1,19),(3,3,7),(22,3,12),(39,3,29),(37,3,30),(38,3,31),(36,3,33),(5,4,2),(14,4,6),(13,4,7),(15,4,8),(9,4,27),(35,4,33),(18,19,27),(41,19,33),(19,21,27);
/*!40000 ALTER TABLE `favorites` ENABLE KEYS */;
UNLOCK TABLES;

--
-- Table structure for table `relations`
--

DROP TABLE IF EXISTS `relations`;
/*!40101 SET @saved_cs_client     = @@character_set_client */;
/*!50503 SET character_set_client = utf8mb4 */;
CREATE TABLE `relations` (
  `id` int NOT NULL AUTO_INCREMENT,
  `follow_id` int NOT NULL,
  `follower_id` int NOT NULL,
  PRIMARY KEY (`id`),
  KEY `follow&follower` (`follow_id`,`follower_id`) /*!80000 INVISIBLE */,
  KEY `followerId` (`follower_id`)
) ENGINE=InnoDB AUTO_INCREMENT=21 DEFAULT CHARSET=utf8;
/*!40101 SET character_set_client = @saved_cs_client */;

--
-- Dumping data for table `relations`
--

LOCK TABLES `relations` WRITE;
/*!40000 ALTER TABLE `relations` DISABLE KEYS */;
INSERT INTO `relations` VALUES (1,1,4),(2,3,4),(3,1,3),(5,4,1),(8,1,19),(9,19,4),(20,3,19);
/*!40000 ALTER TABLE `relations` ENABLE KEYS */;
UNLOCK TABLES;

--
-- Table structure for table `tokens`
--

DROP TABLE IF EXISTS `tokens`;
/*!40101 SET @saved_cs_client     = @@character_set_client */;
/*!50503 SET character_set_client = utf8mb4 */;
CREATE TABLE `tokens` (
  `token_id` int NOT NULL AUTO_INCREMENT,
  `user_token` varchar(45) NOT NULL,
  `user_id` int NOT NULL,
  `expire_time` timestamp NOT NULL,
  PRIMARY KEY (`token_id`),
  KEY `userToken&expireTime` (`user_token`,`expire_time`)
) ENGINE=InnoDB AUTO_INCREMENT=78 DEFAULT CHARSET=utf8;
/*!40101 SET character_set_client = @saved_cs_client */;

--
-- Dumping data for table `tokens`
--

LOCK TABLES `tokens` WRITE;
/*!40000 ALTER TABLE `tokens` DISABLE KEYS */;
/*!40000 ALTER TABLE `tokens` ENABLE KEYS */;
UNLOCK TABLES;

--
-- Table structure for table `users`
--

DROP TABLE IF EXISTS `users`;
/*!40101 SET @saved_cs_client     = @@character_set_client */;
/*!50503 SET character_set_client = utf8mb4 */;
CREATE TABLE `users` (
  `user_id` int NOT NULL AUTO_INCREMENT,
  `name` varchar(45) CHARACTER SET utf8mb4 COLLATE utf8mb4_0900_ai_ci DEFAULT NULL,
  `username` varchar(45) CHARACTER SET utf8mb4 COLLATE utf8mb4_0900_ai_ci NOT NULL,
  `password` varchar(75) CHARACTER SET utf8mb4 COLLATE utf8mb4_0900_ai_ci NOT NULL,
  PRIMARY KEY (`user_id`),
  UNIQUE KEY `username_UNIQUE` (`username`)
) ENGINE=InnoDB AUTO_INCREMENT=82 DEFAULT CHARSET=utf8;
/*!40101 SET character_set_client = @saved_cs_client */;

--
-- Dumping data for table `users`
--

LOCK TABLES `users` WRITE;
/*!40000 ALTER TABLE `users` DISABLE KEYS */;
INSERT INTO `users` VALUES (1,'游戏资讯','admin','$2a$04$IX0rl8I7udvnBlZJK2lK7.zirSqBneTrvnX1srO68OhL.4Zx.WBk.'),(3,'萌宠之家','user','$2a$04$H2cI4uXKrRUcvE2evKjMPumPAAO6NJi1Du8nmRmT..LkDJaqqCZjS'),(4,'七日之都mv','zjj','$2a$04$HJQoumJ25/9O1SIjFo824OYqLEFDa5dlAa3m1lvEbBkJVathhxIRC'),(17,'tteesstt','tteesstt','$2a$04$d3dYKed5HiXjMH.ctF6dd.GzAJY38Q9JypW/z7.6FzxueKjo7jOFC'),(19,'colorofnight','colorofnight','$2a$04$OpkjAHmq3cyzeIQugUmgtegRYGnPR0bBlglTph3qaNFvkFKYqOEUq'),(22,'test','test','$2a$04$/WBg/GheUg0vxbuIEPCZfO2mM578Re5BlBAuhXW7PmtbKrK/iWR/a'),(80,'testtest','testtest','123456');
/*!40000 ALTER TABLE `users` ENABLE KEYS */;
UNLOCK TABLES;

--
-- Table structure for table `videos`
--

DROP TABLE IF EXISTS `videos`;
/*!40101 SET @saved_cs_client     = @@character_set_client */;
/*!50503 SET character_set_client = utf8mb4 */;
CREATE TABLE `videos` (
  `video_id` int NOT NULL AUTO_INCREMENT,
  `user_id` int NOT NULL,
  `title` varchar(45) NOT NULL,
  `play_url` varchar(2048) NOT NULL,
  `cover_url` varchar(2048) NOT NULL,
  `update_time` timestamp NOT NULL DEFAULT CURRENT_TIMESTAMP,
  PRIMARY KEY (`video_id`),
  KEY `updateTime&userId` (`update_time` DESC,`user_id`) /*!80000 INVISIBLE */,
  KEY `userId` (`user_id`)
) ENGINE=InnoDB AUTO_INCREMENT=50 DEFAULT CHARSET=utf8;
/*!40101 SET character_set_client = @saved_cs_client */;

--
-- Dumping data for table `videos`
--

LOCK TABLES `videos` WRITE;
/*!40000 ALTER TABLE `videos` DISABLE KEYS */;
INSERT INTO `videos` VALUES (1,3,'野生的熊','https://www.w3schools.com/html/movie.mp4','https://cdn.pixabay.com/photo/2016/03/27/18/10/bear-1283347_1280.jpg','2022-05-21 05:46:54'),(6,1,'最终幻想6.0pv','https://static.web.sdo.com/jijiamobile/pic/ff14/ffweb860/vidindex.mp4','https://i1.hdslb.com/bfs/archive/749e86d601231d04f7bd81aec0dac85b720e4f2a.jpg','2022-05-21 06:20:29'),(7,1,'万象物语璃sp','http://static.dragonest.com/sdoricaweb/qa/Public/Home/default/pc/images/bg_new.mp4','https://i1.hdslb.com/bfs/archive/a68a9b8fb40ec5c0962ab949b0c703a2fee79781.png','2022-05-25 06:25:18'),(8,1,'万象物语mv','http://static.dragonest.com/sdoricaweb/qa/Public/Home/default/pc/images/pop_video.mp4','https://i2.hdslb.com/bfs/archive/b135c0bcb357a51949da7c805c3b9819b3fbca30.png','2022-05-27 06:26:22'),(12,4,'永远的七日之都结局1','https://tiktok-video-1305174939.cos.ap-guangzhou.myqcloud.com/1深海火种.mp4','https://tiktok-img-1305174939.cos.ap-guangzhou.myqcloud.com/1深海火种.png','2022-05-28 04:14:02'),(13,4,'永远的七日之都结局2','https://tiktok-video-1305174939.cos.ap-guangzhou.myqcloud.com/2心之原野.mp4','https://tiktok-img-1305174939.cos.ap-guangzhou.myqcloud.com/2心之原野.png','2022-05-28 04:14:03'),(14,4,'永远的七日之都结局3','https://tiktok-video-1305174939.cos.ap-guangzhou.myqcloud.com/3神所留下的.mp4','https://tiktok-img-1305174939.cos.ap-guangzhou.myqcloud.com/3神所留下的.png','2022-05-28 04:14:04'),(15,4,'永远的七日之都结局4','https://tiktok-video-1305174939.cos.ap-guangzhou.myqcloud.com/4导师与学生.mp4','https://tiktok-img-1305174939.cos.ap-guangzhou.myqcloud.com/4导师与学生.png','2022-05-28 04:14:05'),(16,4,'永远的七日之都结局5','https://tiktok-video-1305174939.cos.ap-guangzhou.myqcloud.com/5终末的花火.mp4','https://tiktok-img-1305174939.cos.ap-guangzhou.myqcloud.com/5终末的花火.png','2022-05-28 04:14:06'),(17,4,'永远的七日之都结局6','https://tiktok-video-1305174939.cos.ap-guangzhou.myqcloud.com/6奇异恩典.mp4','https://tiktok-img-1305174939.cos.ap-guangzhou.myqcloud.com/6奇异恩典.png','2022-05-28 04:14:07'),(18,4,'永远的七日之都结局7','https://tiktok-video-1305174939.cos.ap-guangzhou.myqcloud.com/7希冀者的飘零.mp4','https://tiktok-img-1305174939.cos.ap-guangzhou.myqcloud.com/7希冀者的飘零.png','2022-05-28 04:14:08'),(19,4,'永远的七日之都结局8','https://tiktok-video-1305174939.cos.ap-guangzhou.myqcloud.com/8飘零者的希冀.mp4','https://tiktok-img-1305174939.cos.ap-guangzhou.myqcloud.com/8飘零者的希冀.png','2022-05-28 04:14:09'),(27,3,'野生猫猫','https://tiktok-video-1305174939.cos.ap-guangzhou.myqcloud.com/1_2022-06-18-11-35-07_VIDEO_20220613_152821.mp4','https://tiktok-img-1305174939.cos.ap-guangzhou.myqcloud.com/1_2022-06-18-11-35-07_VIDEO_20220613_152821.jpg','2022-06-18 03:35:09'),(28,4,'永远的七日之都结局9','https://tiktok-video-1305174939.cos.ap-guangzhou.myqcloud.com/9门扉之外.mp4','https://tiktok-img-1305174939.cos.ap-guangzhou.myqcloud.com/9门扉之外.png','2022-06-18 08:07:16'),(29,4,'永远的七日之都结局10','https://tiktok-video-1305174939.cos.ap-guangzhou.myqcloud.com/10微光.mp4','https://tiktok-img-1305174939.cos.ap-guangzhou.myqcloud.com/10微光.png','2022-06-18 08:07:17'),(30,4,'永远的七日之都结局11','https://tiktok-video-1305174939.cos.ap-guangzhou.myqcloud.com/11废土机神.mp4','https://tiktok-img-1305174939.cos.ap-guangzhou.myqcloud.com/11废土机神.png','2022-06-18 08:07:18'),(31,4,'永远的七日之都结局12','https://tiktok-video-1305174939.cos.ap-guangzhou.myqcloud.com/12相拥之时.mp4','https://tiktok-img-1305174939.cos.ap-guangzhou.myqcloud.com/12相拥之时.png','2022-06-18 08:07:19'),(32,3,'是可爱的猫猫哦','https://tiktok-video-1305174939.cos.ap-guangzhou.myqcloud.com/3_2022-06-18-16-29-40_neko.mp4','https://tiktok-img-1305174939.cos.ap-guangzhou.myqcloud.com/3_2022-06-18-16-29-40_neko.jpg','2022-06-18 08:29:42'),(33,3,'是可爱的猫猫~','https://tiktok-video-1305174939.cos.ap-guangzhou.myqcloud.com/3_2022-06-18-21-23-07_neko.mp4','https://tiktok-img-1305174939.cos.ap-guangzhou.myqcloud.com/3_2022-06-18-21-23-07_neko.jpg','2022-06-18 13:23:09');
/*!40000 ALTER TABLE `videos` ENABLE KEYS */;
UNLOCK TABLES;
/*!40103 SET TIME_ZONE=@OLD_TIME_ZONE */;

/*!40101 SET SQL_MODE=@OLD_SQL_MODE */;
/*!40014 SET FOREIGN_KEY_CHECKS=@OLD_FOREIGN_KEY_CHECKS */;
/*!40014 SET UNIQUE_CHECKS=@OLD_UNIQUE_CHECKS */;
/*!40101 SET CHARACTER_SET_CLIENT=@OLD_CHARACTER_SET_CLIENT */;
/*!40101 SET CHARACTER_SET_RESULTS=@OLD_CHARACTER_SET_RESULTS */;
/*!40101 SET COLLATION_CONNECTION=@OLD_COLLATION_CONNECTION */;
/*!40111 SET SQL_NOTES=@OLD_SQL_NOTES */;

-- Dump completed on 2022-06-19  2:38:18
