create table lesson_contents(
                                id int not null auto_increment primary key ,
                                lesson_title_id int not null ,
                                content varchar(255) not null default 'unknown',
                                created_at datetime not null default current_timestamp,
                                updated_at datetime not null default current_timestamp
)engine = innodb;