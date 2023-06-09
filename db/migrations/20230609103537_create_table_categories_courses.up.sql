create table categories_courses(
                                        category_id int not null ,
                                        course_id int not null ,
                                        created_at datetime not null default current_timestamp,
                                        updated_at datetime not null default current_timestamp,
                                        primary key (category_id, course_id)
)engine = innodb;