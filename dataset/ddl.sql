create table dataset
(
    id bigint auto_increment
        primary key,
    user_id bigint not null,
    dataset_no bigint not null,
    url varchar(1024) not null,
    name varchar(100) null,
    description varchar(2000) null,
    public tinyint(1) null,
    status varchar(10) not null,
    create_time datetime default current_timestamp() not null,
    update_time datetime default current_timestamp() not null on update current_timestamp()
);

