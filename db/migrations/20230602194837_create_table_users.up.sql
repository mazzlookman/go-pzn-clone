create table users(
                      id int not null auto_increment,
                      name varchar(255) not null ,
                      email varchar(255) not null ,
                      password varchar(255) not null ,
                      avatar varchar(255),
                      primary key (id),
                      created_at datetime not null default current_timestamp,
                      updated_at datetime not null default current_timestamp
)engine = innodb;