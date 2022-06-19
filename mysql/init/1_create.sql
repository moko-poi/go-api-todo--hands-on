create table IF not exists `tasks`
(
    `id`               INT(20) AUTO_INCREMENT,
    `title`             VARCHAR(20) NOT NULL,
    `content`             VARCHAR(100) NOT NULL,
    PRIMARY KEY (`id`)
    ) DEFAULT CHARSET=utf8 COLLATE=utf8_bin;

insert into tasks(title, content) values("title1", "content1");
insert into tasks(title, content) values("title2", "content2");
insert into tasks(title, content) values("title3", "content3");
insert into tasks(title, content) values("title4", "content4");
