-- MySQL dump 10.13  Distrib 5.7.16, for Linux (x86_64)
--
-- Host: localhost    Database: im
-- ------------------------------------------------------
-- Server version	5.7.16-0ubuntu0.16.10.1

/*!40101 SET @OLD_CHARACTER_SET_CLIENT=@@CHARACTER_SET_CLIENT */;
/*!40101 SET @OLD_CHARACTER_SET_RESULTS=@@CHARACTER_SET_RESULTS */;
/*!40101 SET @OLD_COLLATION_CONNECTION=@@COLLATION_CONNECTION */;
/*!40101 SET NAMES utf8 */;
/*!40103 SET @OLD_TIME_ZONE=@@TIME_ZONE */;
/*!40103 SET TIME_ZONE='+00:00' */;
/*!40014 SET @OLD_UNIQUE_CHECKS=@@UNIQUE_CHECKS, UNIQUE_CHECKS=0 */;
/*!40014 SET @OLD_FOREIGN_KEY_CHECKS=@@FOREIGN_KEY_CHECKS, FOREIGN_KEY_CHECKS=0 */;
/*!40101 SET @OLD_SQL_MODE=@@SQL_MODE, SQL_MODE='NO_AUTO_VALUE_ON_ZERO' */;
/*!40111 SET @OLD_SQL_NOTES=@@SQL_NOTES, SQL_NOTES=0 */;

--
-- Table structure for table `message`
--

DROP TABLE IF EXISTS `message`;
/*!40101 SET @saved_cs_client     = @@character_set_client */;
/*!40101 SET character_set_client = utf8 */;
CREATE TABLE `message` (
  `id` int(11) NOT NULL AUTO_INCREMENT,
  `cmd` varchar(45) DEFAULT NULL,
  `toid` varchar(45) DEFAULT NULL,
  `fromid` varchar(45) DEFAULT NULL,
  `created` datetime DEFAULT NULL,
  `msg` varchar(256) DEFAULT NULL,
  `read` varchar(45) DEFAULT NULL,
  PRIMARY KEY (`id`)
) ENGINE=InnoDB AUTO_INCREMENT=49 DEFAULT CHARSET=utf8;
/*!40101 SET character_set_client = @saved_cs_client */;

--
-- Dumping data for table `message`
--

LOCK TABLES `message` WRITE;
/*!40000 ALTER TABLE `message` DISABLE KEYS */;
INSERT INTO `message` VALUES (26,'chat','002','001','2016-11-29 12:29:43','message','1'),(27,'chat','002','001','2016-11-29 12:29:46','message','1'),(28,'chat','002','001','2016-11-29 12:29:46','message','1'),(29,'chat','002','001','2016-11-29 12:29:47','message','1'),(30,'chat','002','001','2016-11-29 12:29:47','message','1'),(31,'chat','002','001','2016-11-29 12:29:47','message','1'),(32,'chat','002','001','2016-11-29 12:29:47','message','1'),(33,'chat','002','001','2016-11-29 12:29:47','message','1'),(34,'chat','002','001','2016-11-29 12:29:47','message','1'),(35,'chat','002','001','2016-11-29 12:29:48','message','1'),(36,'chat','002','001','2016-11-29 12:29:48','message','1'),(37,'chat','002','001','2016-11-29 12:29:48','message','1'),(38,'chat','002','001','2016-11-29 12:29:48','message','1'),(39,'chat','002','001','2016-11-29 12:29:48','message','1'),(40,'chat','002','001','2016-11-29 12:29:48','message','1'),(41,'chat','002','001','2016-11-29 12:29:48','message','1'),(42,'chat','002','001','2016-11-29 12:30:04','message','1'),(43,'chat','002','001','2016-11-29 12:30:04','message','1'),(44,'chat','002','001','2016-11-29 12:30:05','message','1'),(45,'chat','002','001','2016-11-29 12:30:05','message','1'),(46,'chat','002','001','2016-11-29 12:30:05','message','1'),(47,'chat','002','001','2016-11-29 12:30:05','message','1'),(48,'chat','002','001','2016-11-29 12:30:05','message','1');
/*!40000 ALTER TABLE `message` ENABLE KEYS */;
UNLOCK TABLES;
/*!40103 SET TIME_ZONE=@OLD_TIME_ZONE */;

/*!40101 SET SQL_MODE=@OLD_SQL_MODE */;
/*!40014 SET FOREIGN_KEY_CHECKS=@OLD_FOREIGN_KEY_CHECKS */;
/*!40014 SET UNIQUE_CHECKS=@OLD_UNIQUE_CHECKS */;
/*!40101 SET CHARACTER_SET_CLIENT=@OLD_CHARACTER_SET_CLIENT */;
/*!40101 SET CHARACTER_SET_RESULTS=@OLD_CHARACTER_SET_RESULTS */;
/*!40101 SET COLLATION_CONNECTION=@OLD_COLLATION_CONNECTION */;
/*!40111 SET SQL_NOTES=@OLD_SQL_NOTES */;

-- Dump completed on 2016-11-29 20:45:09
