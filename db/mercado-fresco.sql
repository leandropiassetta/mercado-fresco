SET @OLD_UNIQUE_CHECKS=@@UNIQUE_CHECKS, UNIQUE_CHECKS=0;
SET @OLD_FOREIGN_KEY_CHECKS=@@FOREIGN_KEY_CHECKS, FOREIGN_KEY_CHECKS=0;
SET @OLD_SQL_MODE=@@SQL_MODE, SQL_MODE='ONLY_FULL_GROUP_BY,STRICT_TRANS_TABLES,NO_ZERO_IN_DATE,NO_ZERO_DATE,ERROR_FOR_DIVISION_BY_ZERO,NO_ENGINE_SUBSTITUTION';

-- -----------------------------------------------------
-- Schema mercado-fresco
-- -----------------------------------------------------

-- -----------------------------------------------------
-- Schema mercado-fresco
-- -----------------------------------------------------
-- DROP DATABASE `mercado-fresco`;
CREATE SCHEMA IF NOT EXISTS `mercado-fresco` DEFAULT CHARACTER SET utf8 ;
USE `mercado-fresco` ;

-- -----------------------------------------------------
-- Table `mercado-fresco`.`buyer`
-- -----------------------------------------------------
CREATE TABLE IF NOT EXISTS `mercado-fresco`.`buyer` (
    `id` SERIAL,
    `card_number_id` VARCHAR(45) NOT NULL,
    `first_name` VARCHAR(45) NOT NULL,
    `last_name` VARCHAR(45) NOT NULL,
    PRIMARY KEY (`id`))
    ENGINE = InnoDB;

-- -----------------------------------------------------
-- Table `mercado-fresco`.`warehouse`
-- -----------------------------------------------------
CREATE TABLE IF NOT EXISTS `mercado-fresco`.`warehouse` (
    `id` SERIAL,
    `warehouse_code` VARCHAR(20) NOT NULL,
    `address` VARCHAR(80) NOT NULL,
    `telephone` VARCHAR(15) NOT NULL,
    `minimun_capacity` INT NOT NULL,
    `minimun_temperature` INT NOT NULL,
    PRIMARY KEY (`id`))
    ENGINE = InnoDB;

-- -----------------------------------------------------
-- Table `mercado-fresco`.`employee`
-- -----------------------------------------------------
CREATE TABLE IF NOT EXISTS `mercado-fresco`.`employee` (
    `id` SERIAL,
    `card_number_id` VARCHAR(45) NOT NULL,
    `first_name` VARCHAR(45) NOT NULL,
    `last_name` VARCHAR(45) NOT NULL,
    `warehouse_id` BIGINT UNSIGNED,
    PRIMARY KEY (`id`))
    ENGINE = InnoDB;

-- -----------------------------------------------------
-- Table `mercado-fresco`.`seller`
-- -----------------------------------------------------
CREATE TABLE IF NOT EXISTS `mercado-fresco`.`seller` (
    `id` SERIAL,
    `cid` INT(11) UNSIGNED NOT NULL,
    `company_name` VARCHAR(80) NOT NULL,
    `address` VARCHAR(80) NOT NULL,
    `telephone` VARCHAR(15) NOT NULL,
    PRIMARY KEY (`id`))
    ENGINE = InnoDB;

-- -----------------------------------------------------
-- Table `mercado-fresco`.`product`
-- -----------------------------------------------------
CREATE TABLE IF NOT EXISTS `mercado-fresco`.`product` (
    `id` SERIAL,
    `product_code` INT(11) NULL,
    `description` VARCHAR(45) NULL,
    `width` VARCHAR(2) NULL,
    `height` VARCHAR(2) NULL,
    `length` VARCHAR(2) NULL,
    `net_weight` VARCHAR(2) NULL,
    `expiration_rate` VARCHAR(15) NULL,
    `recommended_freezing_temperature` DECIMAL(2) NULL,
    `freezing_rate` DECIMAL(2) NULL,
    `product_type_id` INT(11) NULL,
    `seller_id` BIGINT UNSIGNED,
    PRIMARY KEY (`id`))
    ENGINE = InnoDB;

-- -----------------------------------------------------
-- Table `mercado-fresco`.`section`
-- -----------------------------------------------------
CREATE TABLE IF NOT EXISTS `mercado-fresco`.`section` (
    `id` SERIAL,
    `section_number` INT(11) NULL,
    `current_temperature` INT(11) NULL,
    `minimum_temperature` INT(11) NULL,
    `current_capacity` INT(11) NULL,
    `minimum_capacity` INT(11) NULL,
    `maximum_capacity` INT(11) NULL,
    `warehouse_id` BIGINT UNSIGNED,
    `product_type_id` INT(11) NULL,
    PRIMARY KEY (`id`))
    ENGINE = InnoDB;

ALTER TABLE `mercado-fresco`.`product` ADD CONSTRAINT `FK_PRODUCT_SELLER` FOREIGN KEY (`seller_id`) REFERENCES `mercado-fresco`.`seller`(`id`);
ALTER TABLE `mercado-fresco`.`product` ADD UNIQUE KEY `PRODUCT_TYPE_ID` (`product_type_id`);

ALTER TABLE `mercado-fresco`.`employee` ADD CONSTRAINT `FK_EMPLOYEE_WAREHOUSE` FOREIGN KEY (`warehouse_id`) REFERENCES `mercado-fresco`.`warehouse`(`id`);

ALTER TABLE `mercado-fresco`.`section` ADD CONSTRAINT `FK_SECTION_WAREHOUSE` FOREIGN KEY (`warehouse_id`) REFERENCES `mercado-fresco`.`warehouse`(`id`);
ALTER TABLE `mercado-fresco`.`section` ADD CONSTRAINT `FK_SECTION_PRODUCT` FOREIGN KEY (`product_type_id`) REFERENCES `mercado-fresco`.`product`(`product_type_id`);

SET SQL_MODE=@OLD_SQL_MODE;
SET FOREIGN_KEY_CHECKS=@OLD_FOREIGN_KEY_CHECKS;
SET UNIQUE_CHECKS=@OLD_UNIQUE_CHECKS;