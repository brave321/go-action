# 使用官方Nginx基础镜像
FROM nginx:1.21

# 替换默认配置文件

# 复制静态文件到容器的指定目录（假设你有静态文件）
COPY hi.html /usr/share/nginx/html

# 设置环境变量（如果需要）

# 容器启动时执行的命令
CMD ["nginx", "-g", "daemon off;"]
