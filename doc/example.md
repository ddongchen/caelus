# 编译caelus
（1）编译二进制
```shell
[xx]# make build
[xx]# ll _output/bin/
total 129236
-rwxr-xr-x 1 root root 51634520 Dec 12 16:28 caelus
-rwxr-xr-x 1 root root  6819965 Dec 12 16:28 caelus_metric_adapter
-rwxr-xr-x 1 root root 16416404 Dec 12 16:28 lighthouse
-rwxr-xr-x 1 root root 17624320 Dec 12 16:28 nm-operator
-rwxr-xr-x 1 root root 39832989 Dec 12 16:28 plugin-server

其中
caelus：daemonset部署。实时计算节点空闲资源，并执行离线资源隔离和资源上报。同时通过干扰检测保证在线的服务质量
caelus_metric_adapter:
nm-operator:
lighthouse:
plugin-server:
```



（2）编译caelus镜像
make image

# 编译lighthouse

cd contrib/lighthouse
make

_output/bin/lighthouse

make rpm

_output/RPMS/x86_64/lighthouse-0.2.1-47.el7.x86_64.rpm

cd contrib/lighthouse-plugin

make

_output/bin/plugin-server

make rpm

_output/RPMS/x86_64/plugin-server-0.3.0-47.el7.x86_64.rpm

# 安装
（1）安装lighthouse

（2）安装plugin-server

（3）提交caelus workload

# 提交离线作业
1、离线作业通过kubernetes提交

2、离线作业通过YARN提交
