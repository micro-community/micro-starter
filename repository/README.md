# repository 说明

- repository 映射到某个数据库的数据表。

- repository 中是纯粹的数据库操作的聚合。

- repository 通过内的 service 对象对 handler 提供数据操作能力。

- 一般情况情况下，只需要使用到数据源中的一种.

- 目前的考虑的数据源:

  - dgraph 图数据库 -- 是一个完整的 RBAC 的实现
    - 用户、角色、资源、操作
  - memory 内存 -- 内存数据库的实现
  - mongodb 事件 和 日志
  - mysql 用户、角色、资源
  - sqlite 用户、角色、资源
