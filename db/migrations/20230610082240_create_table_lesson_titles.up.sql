create table lesson_titles(
                              id int not null auto_increment primary key ,
                              course_id int not null,
                              title varchar(255) not null default 'untitled',
                              created_at datetime not null default current_timestamp,
                              updated_at datetime not null default current_timestamp
)engine = innodb;