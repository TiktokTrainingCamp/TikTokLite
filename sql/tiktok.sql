CREATE DATABASE  IF NOT EXISTS `tiktok` /*!40100 DEFAULT CHARACTER SET utf8 */ /*!80016 DEFAULT ENCRYPTION='N' */;
USE `tiktok`;
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
-- Dumping data for table `comments`
--

LOCK TABLES `comments` WRITE;
/*!40000 ALTER TABLE `comments` DISABLE KEYS */;
INSERT INTO `comments` VALUES (1,3,1,'熊的力量！','2022-05-27 10:23:56'),(2,2,1,'毛子狂喜','2022-05-27 10:35:27'),(3,1,19,'七都！','2022-05-28 22:45:12'),(4,1,19,'新结局！','2022-05-28 23:07:39'),(6,2,19,'你删不掉我！','2022-05-28 23:39:30'),(7,1,19,'好耶！','2022-05-28 23:40:15');
/*!40000 ALTER TABLE `comments` ENABLE KEYS */;
UNLOCK TABLES;

--
-- Dumping data for table `favorites`
--

LOCK TABLES `favorites` WRITE;
/*!40000 ALTER TABLE `favorites` DISABLE KEYS */;
INSERT INTO `favorites` VALUES (1,2,6),(2,2,7),(3,3,7),(5,4,2),(9,4,23),(10,1,19),(13,4,7),(14,4,6),(15,4,8),(16,1,18),(17,1,12);
/*!40000 ALTER TABLE `favorites` ENABLE KEYS */;
UNLOCK TABLES;

--
-- Dumping data for table `relations`
--

LOCK TABLES `relations` WRITE;
/*!40000 ALTER TABLE `relations` DISABLE KEYS */;
INSERT INTO `relations` VALUES (1,1,2),(2,3,2),(3,1,3),(5,4,1),(7,1,4);
/*!40000 ALTER TABLE `relations` ENABLE KEYS */;
UNLOCK TABLES;

--
-- Dumping data for table `tokens`
--

LOCK TABLES `tokens` WRITE;
/*!40000 ALTER TABLE `tokens` DISABLE KEYS */;
/*!40000 ALTER TABLE `tokens` ENABLE KEYS */;
UNLOCK TABLES;

--
-- Dumping data for table `users`
--

LOCK TABLES `users` WRITE;
/*!40000 ALTER TABLE `users` DISABLE KEYS */;
INSERT INTO `users` VALUES (1,'admin','admin','123456'),(2,'zhanglei','zhanglei','douyin'),(3,'user','user','123456'),(4,'zjj','colorofnight','654321'),(7,'test','test','123456'),(8,'test2','test2','123456'),(9,'test3','test3','123456'),(10,'test4','test4','123456'),(11,'test5','test5','123456'),(16,'testt','testt','123456');
/*!40000 ALTER TABLE `users` ENABLE KEYS */;
UNLOCK TABLES;

--
-- Dumping data for table `videos`
--

LOCK TABLES `videos` WRITE;
/*!40000 ALTER TABLE `videos` DISABLE KEYS */;
INSERT INTO `videos` VALUES (1,1,'野生的熊','https://www.w3schools.com/html/movie.mp4','https://cdn.pixabay.com/photo/2016/03/27/18/10/bear-1283347_1280.jpg','2022-05-21 13:46:54'),(6,1,'最终幻想6.0pv','https://static.web.sdo.com/jijiamobile/pic/ff14/ffweb860/vidindex.mp4','https://i1.hdslb.com/bfs/archive/749e86d601231d04f7bd81aec0dac85b720e4f2a.jpg','2022-05-21 14:20:29'),(7,1,'万象物语璃sp','http://static.dragonest.com/sdoricaweb/qa/Public/Home/default/pc/images/bg_new.mp4','https://i1.hdslb.com/bfs/archive/a68a9b8fb40ec5c0962ab949b0c703a2fee79781.png','2022-05-25 14:25:18'),(8,1,'万象物语mv','http://static.dragonest.com/sdoricaweb/qa/Public/Home/default/pc/images/pop_video.mp4','https://i2.hdslb.com/bfs/archive/b135c0bcb357a51949da7c805c3b9819b3fbca30.png','2022-05-27 14:26:22'),(12,4,'永远的七日之都结局1','https://nie.v.netease.com/nie/2020/1229/c44c910dbf07ba82c1162448d96e6ec8.mp4','https://fc.res.netease.com/pc/zt/20190513121653/img/icons/pic123010_2bc1e12.png','2022-05-28 12:14:02'),(13,4,'永远的七日之都结局2','https://nie.v.netease.com/nie/2020/1229/38cb701148f26741d80a45bc361aa17e.mp4','https://fc.res.netease.com/pc/zt/20190513121653/img/icons/pic123014_d3fd5e7.png','2022-05-28 12:14:03'),(14,4,'永远的七日之都结局3','https://nie.v.netease.com/nie/2020/1229/44c3739d4351e65ad3c4c556c3f92fd6.mp4','https://fc.res.netease.com/pc/zt/20190513121653/img/icons/pic123011_0b33584.png','2022-05-28 12:14:04'),(15,4,'永远的七日之都结局4','https://nie.v.netease.com/nie/2020/1229/ffaf0297df5a2f92fcaef16acade323b.mp4','https://fc.res.netease.com/pc/zt/20190513121653/img/icons/pic12301_5b44e6f.png','2022-05-28 12:14:05'),(16,4,'永远的七日之都结局5','https://nie.v.netease.com/nie/2020/1229/6fff2fc8f764e0826101963168cd85e5.mp4','https://fc.res.netease.com/pc/zt/20190513121653/img/icons/pic123016_727c7dc.png','2022-05-28 12:14:06'),(17,4,'永远的七日之都结局6','https://nie.v.netease.com/nie/2020/1229/4013089c523ccd0c4ca452fee87fbf52.mp4','https://fc.res.netease.com/pc/zt/20190513121653/img/icons/pic12309_a7c4e7d.png','2022-05-28 12:14:07'),(18,4,'永远的七日之都结局7','https://nie.v.netease.com/nie/2020/1229/0847ae998103e2c722abae4b87a1b537.mp4','https://fc.res.netease.com/pc/zt/20190513121653/img/icons/pic12308_c3e264c.png','2022-05-28 12:14:08'),(19,4,'永远的七日之都结局8','https://nie.v.netease.com/nie/2020/1229/9b1fc12eab23553df6401a6ac4c9d22e.mp4','https://fc.res.netease.com/pc/zt/20190513121653/img/icons/pic12307_39db61f.png','2022-05-28 12:14:09'),(23,1,'是可爱的猫猫~','http://10.21.193.72:8080/static/1_2022-05-28-17-11-51_neko.mp4','http://10.21.193.72:8080/static/1_2022-05-28-17-11-51_neko.jpg','2022-05-28 17:11:52');
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

-- Dump completed on 2022-05-31 10:58:15
