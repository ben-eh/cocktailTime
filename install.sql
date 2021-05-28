DROP DATABASE IF EXISTS cocktailTime;
CREATE DATABASE cocktailTime;
USE cocktailTime;

CREATE TABLE ingredients (
  ingredient_id INT UNSIGNED NOT NULL AUTO_INCREMENT PRIMARY KEY,
  name VARCHAR(255) DEFAULT ""
);

CREATE TABLE measurements (
  measurement_id INT UNSIGNED NOT NULL AUTO_INCREMENT PRIMARY KEY,
  name VARCHAR(255) DEFAULT ""
);

CREATE TABLE step_ingredient_entry (
  quantity VARCHAR(255) DEFAULT "",
  ingredient_id INT UNSIGNED NOT NULL,
  measurement_id INT UNSIGNED NOT NULL,
  PRIMARY KEY (`quantity`, `ingredient_id`, `measurement_id`)
);



INSERT INTO measurements (name)
values
  ("Oz."),
  ("ml"),
  ("dash"),
  ("dashes"),
  ("tsp"),
  ("tbsp");
