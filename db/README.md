# db

处理数据库连接和数据库关系上下文,不包含任何业务处理逻辑。

一般情况，一个微服务最多连接 1-2 个数据源（往往是 DB 和 Cache)

当前 case 里面持久化的的连接数据库是：

- sqlite
- mysql
- mongodb
- dgraph
- memory

缓存是：

- redis
