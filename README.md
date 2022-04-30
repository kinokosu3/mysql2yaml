# mysql2yaml

Transfer mysql data to yaml file.

Example:

`-table` and `-cond` flag:

Notice: `-table` and `-cond` flag can be used at the same time.

```shell
./mysql2shell \
-host=127.0.0.1 \
-port=3306 \
-password=123456 \
-database=shop \
-table=user \
-cond="where id=1"
```

equal sql:

```SQL
select * from user where id=1
```

result:
```yaml
- created_at: "2021-09-06 16:01:28"
  id: "1"
  phone: "13800138000"
  username: admin
```

`-sql` flag:

```shell
./mysql2shell \
-host=127.0.0.1 \
-port=3306 \
-password=123456 \
-database=shop \
-sql inout_orders,inout_orders_detail="select * from inout_orders left join inout_orders_detail on inout_orders_detail.inout_id=inout_orders.id where inout_orders.id=1802"
```

equal sql:

```SQL
select inout_orders.* from inout_orders left join inout_orders_detail on inout_orders_detail.inout_id=inout_orders.id where inout_orders.id=1802
select inout_orders_detail.* from inout_orders left join inout_orders_detail on inout_orders_detail.inout_id=inout_orders.id where inout_orders.id=1802
```

result:

inout_orders.yaml
```yaml
- created_at: "2022-04-28 14:22:39"
  id: "1802"
  updated_at: "2022-04-28 14:22:39"
```

inout_orders_detail.yaml
```yaml
- created_at: "2022-04-28 14:22:39"
  id: "1766"
  inout_id: "1802"
  updated_at: "2022-04-28 14:22:39"
```