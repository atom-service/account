# Proto

## 文件名

由于 grpc 会将文件名作为 namespace 注入全局，所以必须保证文件名的全局唯一性，这个唯一性的范围是包括使用本服务的应用，所以我们制定了文件的命名规范，来避免冲突。

- 组织名
- 服务名
- 模块名
- 版本号(可选)

例如：

```shell
atom_service.account.common.proto
```
