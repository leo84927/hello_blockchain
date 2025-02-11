create table if not exists eth_address
(
    addr          char(42)                  not null comment '地址'
    primary key,
    id_rsa        varchar(200)              not null comment '私鑰',
    created_at    bigint unsigned default 0 not null comment '建立時間',
    updated_at    bigint unsigned default 0 not null comment '變更時間'
) comment 'eth地址表';