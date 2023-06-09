create table categories(
                           id int not null auto_increment primary key ,
                           name varchar(255) not null,
                           created_at datetime not null default current_timestamp,
                           updated_at datetime not null default current_timestamp
)engine = innodb;