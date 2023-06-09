create table user_courses(
                                    user_id int not null ,
                                    course_id int not null ,
                                    created_at datetime not null default current_timestamp,
                                    updated_at datetime not null default current_timestamp,
                                    primary key (user_id, course_id)
)engine = innodb;