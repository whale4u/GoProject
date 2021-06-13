## 测试记录

1

```shell
curl -X GET -I http://127.0.0.1:9000/api/v1/hello
```

=== admin /api/v1/hello GET 很遗憾,权限验证没有通过

2

```shell
curl -X POST -I http://127.0.0.1:9000/api/v1/add
```

增加Policy 增加成功

3

```shell
curl -X GET -I http://127.0.0.1:9000/api/v1/get
```

查看policy value: admin value: /api/v1/hello value: GET

4

```shell
curl -X GET -I http://127.0.0.1:9000/api/v1/hello
```

恭喜您,权限验证通过 Hello 接收到GET请求..

5

```shell
curl -X DELETE -I http://127.0.0.1:9000/api/v1/delete
```

删除Policy 删除成功