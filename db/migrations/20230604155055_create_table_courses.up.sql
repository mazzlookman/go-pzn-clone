create table courses(
                        id int not null auto_increment primary key,
                        author_id int not null ,
                        title varchar(255) not null ,
                        slug varchar(255) not null ,
                        description text not null ,
                        perks text not null ,
                        price int not null default 0,
                        progress int not null default 0,
                        banner varchar(255),
                        created_at datetime not null default current_timestamp,
                        updated_at datetime not null default current_timestamp
)engine = innodb;