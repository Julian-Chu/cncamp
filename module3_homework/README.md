# Homework 3
1. 构建本地镜像。
2. 编写 Dockerfile 将练习 2.2 编写的 httpserver 容器化（请思考有哪些最佳实践可以引入到 Dockerfile 中来）。
3. 将镜像推送至 Docker 官方镜像仓库。
4. 通过 Docker 命令本地启动 httpserver。
5. 通过 nsenter 进入容器查看 IP 配置


```shell
// context: ../module2_homework
make docker-build
make docker-push
make docker-run
