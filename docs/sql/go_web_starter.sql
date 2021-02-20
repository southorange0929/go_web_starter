create table if not exists user
(
    id          int auto_increment
    primary key,
    username    varchar(255)                         not null,
    password    varchar(255)                         not null,
    create_time datetime default current_timestamp() not null,
    update_time datetime default current_timestamp() not null on update current_timestamp()
    );

insert into go_web_starter.user (id, username, password, create_time, update_time) values  (1, 'admin', 'admin', '2021-02-20 02:25:56', '2021-02-20 02:25:56');