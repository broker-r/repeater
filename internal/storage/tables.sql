create table word (
    name string primary_key,
    repeat_counter integer DEFAULT 0,
    last_repeat timestamp DEFAULT NULL
)