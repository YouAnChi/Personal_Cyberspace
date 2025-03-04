# Linux文件权限表

下面是关于 Linux 文件权限的总结表格，包含二进制、十进制和符号表示法及其对应的权限解释。

| 权限类别         | 二进制 | 十进制 | 符号表示法 | 权限描述           |
| :--------------- | :----- | :----- | :--------- | :----------------- |
| 无权限           | 000    | 0      | ---        | 无读、写、执行权限 |
| 执行权限         | 001    | 1      | --x        | 仅有执行权限       |
| 写权限           | 010    | 2      | -w-        | 仅有写权限         |
| 写和执行权限     | 011    | 3      | -wx        | 写和执行权限       |
| 读权限           | 100    | 4      | r--        | 仅有读权限         |
| 读和执行权限     | 101    | 5      | r-x        | 读和执行权限       |
| 读和写权限       | 110    | 6      | rw-        | 读和写权限         |
| 读、写和执行权限 | 111    | 7      | rwx        | 读、写和执行权限   |

### 示例解释

- `r` 代表读（read），值为 4。
- `w` 代表写（write），值为 2。
- `x` 代表执行（execute），值为 1。

### 权限设置实例

- 权限设置为 755
  - 二进制表示法：111 101 101
  - 符号表示法：rwxr-xr-x
  - 权限描述：所有者有读、写和执行权限，组用户和其他用户有读和执行权限。

```bash
chmod 755 filename
```

- 权限设置为 644
  - 二进制表示法：110 100 100
  - 符号表示法：rw-r--r--
  - 权限描述：所有者有读和写权限，组用户和其他用户只有读权限。

```bash
chmod 644 filename
```

这些权限设置是通过 `chmod` 命令来实现的。`chmod` 命令允许你通过指定权限的八进制表示法来设置文件或目录的权限。