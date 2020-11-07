create table message_datas
(
    id bigint not null primary key,
    user_id int not null,
    message text not null,
    created_at datetime default current_timestamp() not null,
    updated_at datetime default current_timestamp() not null,
    deleted_at datetime null
);