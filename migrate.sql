-- MySQL dump 10.13  Distrib 8.0.19, for Win64 (x86_64)
--
-- Host: localhost    Database: database_1
-- ------------------------------------------------------
-- Server version	9.0.1

/*!40101 SET @OLD_CHARACTER_SET_CLIENT=@@CHARACTER_SET_CLIENT */;
/*!40101 SET @OLD_CHARACTER_SET_RESULTS=@@CHARACTER_SET_RESULTS */;
/*!40101 SET @OLD_COLLATION_CONNECTION=@@COLLATION_CONNECTION */;
/*!50503 SET NAMES utf8mb4 */;
/*!40103 SET @OLD_TIME_ZONE=@@TIME_ZONE */;
/*!40103 SET TIME_ZONE='+00:00' */;
/*!40014 SET @OLD_UNIQUE_CHECKS=@@UNIQUE_CHECKS, UNIQUE_CHECKS=0 */;
/*!40014 SET @OLD_FOREIGN_KEY_CHECKS=@@FOREIGN_KEY_CHECKS, FOREIGN_KEY_CHECKS=0 */;
/*!40101 SET @OLD_SQL_MODE=@@SQL_MODE, SQL_MODE='NO_AUTO_VALUE_ON_ZERO' */;
/*!40111 SET @OLD_SQL_NOTES=@@SQL_NOTES, SQL_NOTES=0 */;

--
-- Table structure for table `categories`
--

DROP TABLE IF EXISTS `categories`;
/*!40101 SET @saved_cs_client     = @@character_set_client */;
/*!50503 SET character_set_client = utf8mb4 */;
CREATE TABLE `categories` (
  `id` int NOT NULL AUTO_INCREMENT,
  `name` varchar(255) NOT NULL,
  `description` text,
  PRIMARY KEY (`id`)
) ENGINE=InnoDB AUTO_INCREMENT=116 DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_0900_ai_ci;
/*!40101 SET character_set_client = @saved_cs_client */;

--
-- Dumping data for table `categories`
--

LOCK TABLES `categories` WRITE;
/*!40000 ALTER TABLE `categories` DISABLE KEYS */;
INSERT INTO `categories` VALUES (1,'Electronics123','Electronics category'),(2,'Electronics','Electronics category'),(3,'Electronics','Electronics category'),(4,'Electronics','Electronics category'),(5,'Electronics','Electronics category'),(6,'Test','Test 123'),(7,'Test','Test 123'),(8,'Test','Test 123'),(14,'category 1','category 1'),(15,'category 2','category 2'),(16,'Category 1','This is category 1 description'),(17,'Category 2','This is category 2 description'),(18,'Category 3','This is category 3 description'),(19,'Category 4','This is category 4 description'),(20,'Category 5','This is category 5 description'),(21,'Category 6','This is category 6 description'),(22,'Category 7','This is category 7 description'),(23,'Category 8','This is category 8 description'),(24,'Category 9','This is category 9 description'),(25,'Category 10','This is category 10 description'),(26,'Category 11','This is category 11 description'),(27,'Category 12','This is category 12 description'),(28,'Category 13','This is category 13 description'),(29,'Category 14','This is category 14 description'),(30,'Category 15','This is category 15 description'),(31,'Category 16','This is category 16 description'),(32,'Category 17','This is category 17 description'),(33,'Category 18','This is category 18 description'),(34,'Category 19','This is category 19 description'),(35,'Category 20','This is category 20 description'),(36,'Category 21','This is category 21 description'),(37,'Category 22','This is category 22 description'),(38,'Category 23','This is category 23 description'),(39,'Category 24','This is category 24 description'),(40,'Category 25','This is category 25 description'),(41,'Category 26','This is category 26 description'),(42,'Category 27','This is category 27 description'),(43,'Category 28','This is category 28 description'),(44,'Category 29','This is category 29 description'),(45,'Category 30','This is category 30 description'),(46,'Category 31','This is category 31 description'),(47,'Category 32','This is category 32 description'),(48,'Category 33','This is category 33 description'),(49,'Category 34','This is category 34 description'),(50,'Category 35','This is category 35 description'),(51,'Category 36','This is category 36 description'),(52,'Category 37','This is category 37 description'),(53,'Category 38','This is category 38 description'),(54,'Category 39','This is category 39 description'),(55,'Category 40','This is category 40 description'),(56,'Category 41','This is category 41 description'),(57,'Category 42','This is category 42 description'),(58,'Category 43','This is category 43 description'),(59,'Category 44','This is category 44 description'),(60,'Category 45','This is category 45 description'),(61,'Category 46','This is category 46 description'),(62,'Category 47','This is category 47 description'),(63,'Category 48','This is category 48 description'),(64,'Category 49','This is category 49 description'),(65,'Category 50','This is category 50 description'),(66,'Category 51','This is category 51 description'),(67,'Category 52','This is category 52 description'),(68,'Category 53','This is category 53 description'),(69,'Category 54','This is category 54 description'),(70,'Category 55','This is category 55 description'),(71,'Category 56','This is category 56 description'),(72,'Category 57','This is category 57 description'),(73,'Category 58','This is category 58 description'),(74,'Category 59','This is category 59 description'),(75,'Category 60','This is category 60 description'),(76,'Category 61','This is category 61 description'),(77,'Category 62','This is category 62 description'),(78,'Category 63','This is category 63 description'),(79,'Category 64','This is category 64 description'),(80,'Category 65','This is category 65 description'),(81,'Category 66','This is category 66 description'),(82,'Category 67','This is category 67 description'),(83,'Category 68','This is category 68 description'),(84,'Category 69','This is category 69 description'),(85,'Category 70','This is category 70 description'),(86,'Category 71','This is category 71 description'),(87,'Category 72','This is category 72 description'),(88,'Category 73','This is category 73 description'),(89,'Category 74','This is category 74 description'),(90,'Category 75','This is category 75 description'),(91,'Category 76','This is category 76 description'),(92,'Category 77','This is category 77 description'),(93,'Category 78','This is category 78 description'),(94,'Category 79','This is category 79 description'),(95,'Category 80','This is category 80 description'),(96,'Category 81','This is category 81 description'),(97,'Category 82','This is category 82 description'),(98,'Category 83','This is category 83 description'),(99,'Category 84','This is category 84 description'),(100,'Category 85','This is category 85 description'),(101,'Category 86','This is category 86 description'),(102,'Category 87','This is category 87 description'),(103,'Category 88','This is category 88 description'),(104,'Category 89','This is category 89 description'),(105,'Category 90','This is category 90 description'),(106,'Category 91','This is category 91 description'),(107,'Category 92','This is category 92 description'),(108,'Category 93','This is category 93 description'),(109,'Category 94','This is category 94 description'),(110,'Category 95','This is category 95 description'),(111,'Category 96','This is category 96 description'),(112,'Category 97','This is category 97 description'),(113,'Category 98','This is category 98 description'),(114,'Category 99','This is category 99 description'),(115,'Category 100','This is category 100 description');
/*!40000 ALTER TABLE `categories` ENABLE KEYS */;
UNLOCK TABLES;

--
-- Table structure for table `customers`
--

DROP TABLE IF EXISTS `customers`;
/*!40101 SET @saved_cs_client     = @@character_set_client */;
/*!50503 SET character_set_client = utf8mb4 */;
CREATE TABLE `customers` (
  `id` int NOT NULL,
  `name` varchar(255) DEFAULT NULL,
  `email` varchar(255) DEFAULT NULL,
  `created_at` timestamp NULL DEFAULT CURRENT_TIMESTAMP,
  PRIMARY KEY (`id`)
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_0900_ai_ci;
/*!40101 SET character_set_client = @saved_cs_client */;

--
-- Dumping data for table `customers`
--

LOCK TABLES `customers` WRITE;
/*!40000 ALTER TABLE `customers` DISABLE KEYS */;
INSERT INTO `customers` VALUES (1,'John Doe','john@example.com','2022-01-01 00:00:00'),(2,'Jane Smith','jane@example.com','2022-01-02 00:00:00'),(3,'Bob Johnson','bob@example.com','2022-01-03 00:00:00');
/*!40000 ALTER TABLE `customers` ENABLE KEYS */;
UNLOCK TABLES;

--
-- Table structure for table `order_items`
--

DROP TABLE IF EXISTS `order_items`;
/*!40101 SET @saved_cs_client     = @@character_set_client */;
/*!50503 SET character_set_client = utf8mb4 */;
CREATE TABLE `order_items` (
  `id` int NOT NULL,
  `order_id` int DEFAULT NULL,
  `product_id` int DEFAULT NULL,
  `quantity` int DEFAULT NULL,
  `unit_price` decimal(10,2) DEFAULT NULL,
  PRIMARY KEY (`id`),
  KEY `order_id` (`order_id`),
  KEY `product_id` (`product_id`),
  CONSTRAINT `order_items_ibfk_1` FOREIGN KEY (`order_id`) REFERENCES `orders` (`id`),
  CONSTRAINT `order_items_ibfk_2` FOREIGN KEY (`product_id`) REFERENCES `products` (`id`)
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_0900_ai_ci;
/*!40101 SET character_set_client = @saved_cs_client */;

--
-- Dumping data for table `order_items`
--

LOCK TABLES `order_items` WRITE;
/*!40000 ALTER TABLE `order_items` DISABLE KEYS */;
INSERT INTO `order_items` VALUES (1,1,6,2,10.00),(2,1,6,1,10.00),(3,2,7,3,10.00),(4,2,7,2,10.00),(5,3,6,1,10.00),(6,4,9,4,10.00);
/*!40000 ALTER TABLE `order_items` ENABLE KEYS */;
UNLOCK TABLES;

--
-- Table structure for table `orders`
--

DROP TABLE IF EXISTS `orders`;
/*!40101 SET @saved_cs_client     = @@character_set_client */;
/*!50503 SET character_set_client = utf8mb4 */;
CREATE TABLE `orders` (
  `id` int NOT NULL,
  `created_at` timestamp NULL DEFAULT CURRENT_TIMESTAMP,
  `customer_id` int DEFAULT NULL,
  `total` decimal(10,0) DEFAULT NULL,
  PRIMARY KEY (`id`),
  KEY `orders_FK` (`customer_id`),
  CONSTRAINT `orders_FK` FOREIGN KEY (`customer_id`) REFERENCES `customers` (`id`)
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_0900_ai_ci;
/*!40101 SET character_set_client = @saved_cs_client */;

--
-- Dumping data for table `orders`
--

LOCK TABLES `orders` WRITE;
/*!40000 ALTER TABLE `orders` DISABLE KEYS */;
INSERT INTO `orders` VALUES (1,'2022-01-01 00:00:00',1,1000),(2,'2022-01-02 00:00:00',2,2000),(3,'2022-01-03 00:00:00',1,1500),(4,'2022-01-04 00:00:00',3,3000);
/*!40000 ALTER TABLE `orders` ENABLE KEYS */;
UNLOCK TABLES;

--
-- Table structure for table `products`
--

DROP TABLE IF EXISTS `products`;
/*!40101 SET @saved_cs_client     = @@character_set_client */;
/*!50503 SET character_set_client = utf8mb4 */;
CREATE TABLE `products` (
  `id` int NOT NULL AUTO_INCREMENT,
  `name` varchar(255) NOT NULL,
  `description` text,
  `price` decimal(10,2) NOT NULL,
  `category_id` int DEFAULT NULL,
  `stock_quantity` int DEFAULT NULL,
  `is_active` tinyint(1) DEFAULT '1',
  `created_at` datetime DEFAULT CURRENT_TIMESTAMP,
  `updated_at` datetime DEFAULT CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP,
  PRIMARY KEY (`id`),
  KEY `category_id` (`category_id`),
  CONSTRAINT `products_ibfk_1` FOREIGN KEY (`category_id`) REFERENCES `categories` (`id`)
) ENGINE=InnoDB AUTO_INCREMENT=90 DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_0900_ai_ci;
/*!40101 SET character_set_client = @saved_cs_client */;

--
-- Dumping data for table `products`
--

LOCK TABLES `products` WRITE;
/*!40000 ALTER TABLE `products` DISABLE KEYS */;
INSERT INTO `products` VALUES (6,'New Product','This is a new product',19.99,1,0,0,'2024-09-30 11:29:44','2024-09-30 11:29:44'),(7,'New Product','This is a new product',19.99,1,0,0,'2024-09-30 11:29:45','2024-09-30 11:29:45'),(9,'New Product','This is a new product',19.99,1,0,0,'2024-09-30 11:29:46','2024-09-30 11:29:46'),(11,'New Product','This is a new product',19.99,3,0,0,'2024-09-30 11:29:54','2024-09-30 11:29:54'),(14,'Product 2','This is a product description',19.99,1,10,1,'2024-10-03 04:37:29','2024-10-03 04:37:29'),(15,'Product 15','This is a product description',19.99,2,10,1,'2024-10-03 04:39:30','2024-10-03 04:40:19'),(18,'Product 1','This is product 1 description',10.99,1,100,1,'2022-01-01 00:00:00','2022-01-01 00:00:00'),(19,'Product 2','This is product 2 description',9.99,2,50,1,'2022-01-02 00:00:00','2022-01-02 00:00:00'),(20,'Product 3','This is product 3 description',12.99,3,200,1,'2022-01-03 00:00:00','2022-01-03 00:00:00'),(21,'Product 4','This is product 4 description',8.99,1,150,1,'2022-01-04 00:00:00','2022-01-04 00:00:00'),(22,'Product 5','This is product 5 description',11.99,2,75,1,'2022-01-05 00:00:00','2022-01-05 00:00:00'),(23,'Product 6','This is product 6 description',13.99,3,250,1,'2022-01-06 00:00:00','2022-01-06 00:00:00'),(24,'Product 7','This is product 7 description',7.99,1,100,1,'2022-01-07 00:00:00','2022-01-07 00:00:00'),(25,'Product 8','This is product 8 description',10.99,2,50,1,'2022-01-08 00:00:00','2022-01-08 00:00:00'),(26,'Product 9','This is product 9 description',12.99,3,200,1,'2022-01-09 00:00:00','2022-01-09 00:00:00'),(27,'Product 10','This is product 10 description',8.99,1,150,1,'2022-01-10 00:00:00','2022-01-10 00:00:00'),(28,'Product 11','This is product 11 description',11.99,2,75,1,'2022-01-11 00:00:00','2022-01-11 00:00:00'),(29,'Product 12','This is product 12 description',13.99,3,250,1,'2022-01-12 00:00:00','2022-01-12 00:00:00'),(30,'Product 13','This is product 13 description',7.99,1,100,1,'2022-01-13 00:00:00','2022-01-13 00:00:00'),(31,'Product 14','This is product 14 description',10.99,2,50,1,'2022-01-14 00:00:00','2022-01-14 00:00:00'),(32,'Product 15','This is product 15 description',12.99,3,200,1,'2022-01-15 00:00:00','2022-01-15 00:00:00'),(33,'Product 16','This is product 16 description',8.99,1,150,1,'2022-01-16 00:00:00','2022-01-16 00:00:00'),(34,'Product 17','This is product 17 description',11.99,2,75,1,'2022-01-17 00:00:00','2022-01-17 00:00:00'),(35,'Product 18','This is product 18 description',13.99,3,250,1,'2022-01-18 00:00:00','2022-01-18 00:00:00'),(36,'Product 19','This is product 19 description',7.99,1,100,1,'2022-01-19 00:00:00','2022-01-19 00:00:00'),(37,'Product 20','This is product 20 description',10.99,2,50,1,'2022-01-20 00:00:00','2022-01-20 00:00:00'),(38,'Product 21','This is product 21 description',12.99,3,200,1,'2022-01-21 00:00:00','2022-01-21 00:00:00'),(39,'Product 22','This is product 22 description',8.99,1,150,1,'2022-01-22 00:00:00','2022-01-22 00:00:00'),(40,'Product 23','This is product 23 description',11.99,2,75,1,'2022-01-23 00:00:00','2022-01-23 00:00:00'),(41,'Product 24','This is product 24 description',13.99,3,250,1,'2022-01-24 00:00:00','2022-01-24 00:00:00'),(42,'Product 25','This is product 25 description',7.99,1,100,1,'2022-01-25 00:00:00','2022-01-25 00:00:00'),(43,'Product 26','This is product 26 description',10.99,2,50,1,'2022-01-26 00:00:00','2022-01-26 00:00:00'),(44,'Product 27','This is product 27 description',12.99,3,200,1,'2022-01-27 00:00:00','2022-01-27 00:00:00'),(45,'Product 28','This is product 28 description',8.99,1,150,1,'2022-01-28 00:00:00','2022-01-28 00:00:00'),(46,'Product 29','This is product 29 description',11.99,2,75,1,'2022-01-29 00:00:00','2022-01-29 00:00:00'),(47,'Product 30','This is product 30 description',13.99,3,250,1,'2022-01-30 00:00:00','2022-01-30 00:00:00'),(48,'Product 31','This is product 31 description',7.99,1,100,1,'2022-01-31 00:00:00','2022-01-31 00:00:00'),(49,'Product 32','This is product 32 description',10.99,2,50,1,'2022-02-01 00:00:00','2022-02-01 00:00:00'),(50,'Product 33','This is product 33 description',12.99,3,200,1,'2022-02-02 00:00:00','2022-02-02 00:00:00'),(51,'Product 34','This is product 34 description',8.99,1,150,1,'2022-02-03 00:00:00','2022-02-03 00:00:00'),(52,'Product 35','This is product 35 description',11.99,2,75,1,'2022-02-04 00:00:00','2022-02-04 00:00:00'),(53,'Product 36','This is product 36 description',13.99,3,250,1,'2022-02-04 00:00:00','2022-02-04 00:00:00'),(54,'Product 37','This is product 37 description',7.99,1,100,1,'2022-02-05 00:00:00','2022-02-05 00:00:00'),(55,'Product 38','This is product 38 description',10.99,2,50,1,'2022-02-06 00:00:00','2022-02-06 00:00:00'),(56,'Product 39','This is product 39 description',12.99,3,200,1,'2022-02-07 00:00:00','2022-02-07 00:00:00'),(57,'Product 40','This is product 40 description',8.99,1,150,1,'2022-02-08 00:00:00','2022-02-08 00:00:00'),(58,'Product 41','This is product 41 description',11.99,2,75,1,'2022-02-09 00:00:00','2022-02-09 00:00:00'),(59,'Product 42','This is product 42 description',13.99,3,250,1,'2022-02-10 00:00:00','2022-02-10 00:00:00'),(60,'Product 43','This is product 43 description',7.99,1,100,1,'2022-02-11 00:00:00','2022-02-11 00:00:00'),(61,'Product 44','This is product 44 description',10.99,2,50,1,'2022-02-12 00:00:00','2022-02-12 00:00:00'),(62,'Product 45','This is product 45 description',12.99,3,200,1,'2022-02-13 00:00:00','2022-02-13 00:00:00'),(63,'Product 46','This is product 46 description',8.99,1,150,1,'2022-02-14 00:00:00','2022-02-14 00:00:00'),(64,'Product 47','This is product 47 description',11.99,2,75,1,'2022-02-15 00:00:00','2022-02-15 00:00:00'),(65,'Product 48','This is product 48 description',13.99,3,250,1,'2022-02-16 00:00:00','2022-02-16 00:00:00'),(66,'Product 49','This is product 49 description',7.99,1,100,1,'2022-02-17 00:00:00','2022-02-17 00:00:00'),(67,'Product 50','This is product 50 description',10.99,2,50,1,'2022-02-18 00:00:00','2022-02-18 00:00:00'),(68,'Product 51','This is product 51 description',12.99,3,200,1,'2022-02-19 00:00:00','2022-02-19 00:00:00'),(69,'Product 52','This is product 52 description',8.99,1,150,1,'2022-02-20 00:00:00','2022-02-20 00:00:00'),(70,'Product 53','This is product 53 description',11.99,2,75,1,'2022-02-21 00:00:00','2022-02-21 00:00:00'),(71,'Product 54','This is product 54 description',13.99,3,250,1,'2022-02-22 00:00:00','2022-02-22 00:00:00'),(72,'Product 55','This is product 55 description',7.99,1,100,1,'2022-02-23 00:00:00','2022-02-23 00:00:00'),(73,'Product 56','This is product 56 description',10.99,2,50,1,'2022-02-24 00:00:00','2022-02-24 00:00:00'),(74,'Product 57','This is product 57 description',12.99,3,200,1,'2022-02-25 00:00:00','2022-02-25 00:00:00'),(75,'Product 58','This is product 58 description',8.99,1,150,1,'2022-02-26 00:00:00','2022-02-26 00:00:00'),(76,'Product 59','This is product 59 description',11.99,2,75,1,'2022-02-27 00:00:00','2022-02-27 00:00:00'),(77,'Product 60','This is product 60 description',13.99,3,250,1,'2022-02-28 00:00:00','2022-02-28 00:00:00'),(78,'Product 61','This is product 61 description',7.99,1,100,1,'2022-03-01 00:00:00','2022-03-01 00:00:00'),(79,'Product 62','This is product 62 description',10.99,2,50,1,'2022-03-02 00:00:00','2022-03-02 00:00:00'),(80,'Product 63','This is product 63 description',12.99,3,200,1,'2022-03-03 00:00:00','2022-03-03 00:00:00'),(81,'Product 64','This is product 64 description',8.99,1,150,1,'2022-03-04 00:00:00','2022-03-04 00:00:00'),(82,'Product 65','This is product 65 description',11.99,2,75,1,'2022-03-05 00:00:00','2022-03-05 00:00:00'),(83,'Product 66','This is product 66 description',13.99,3,250,1,'2022-03-06 00:00:00','2022-03-06 00:00:00'),(84,'Product 67','This is product 67 description',7.99,1,100,1,'2022-03-07 00:00:00','2022-03-07 00:00:00'),(85,'Product 68','This is product 68 description',10.99,2,50,1,'2022-03-08 00:00:00','2022-03-08 00:00:00'),(86,'Product 69','This is product 69 description',12.99,3,200,1,'2022-03-09 00:00:00','2022-03-09 00:00:00'),(87,'Product 70','This is product 70 description',8.99,1,150,1,'2022-03-10 00:00:00','2022-03-10 00:00:00'),(88,'Product 71','This is product 71 description',11.99,2,75,1,'2022-03-11 00:00:00','2022-03-11 00:00:00'),(89,'Product 72','This is product 72 description',13.99,3,250,1,'2022-03-12 00:00:00','2022-03-12 00:00:00');
/*!40000 ALTER TABLE `products` ENABLE KEYS */;
UNLOCK TABLES;

--
-- Dumping routines for database 'database_1'
--
/*!40103 SET TIME_ZONE=@OLD_TIME_ZONE */;

/*!40101 SET SQL_MODE=@OLD_SQL_MODE */;
/*!40014 SET FOREIGN_KEY_CHECKS=@OLD_FOREIGN_KEY_CHECKS */;
/*!40014 SET UNIQUE_CHECKS=@OLD_UNIQUE_CHECKS */;
/*!40101 SET CHARACTER_SET_CLIENT=@OLD_CHARACTER_SET_CLIENT */;
/*!40101 SET CHARACTER_SET_RESULTS=@OLD_CHARACTER_SET_RESULTS */;
/*!40101 SET COLLATION_CONNECTION=@OLD_COLLATION_CONNECTION */;
/*!40111 SET SQL_NOTES=@OLD_SQL_NOTES */;

-- Dump completed on 2024-10-03 17:57:09
