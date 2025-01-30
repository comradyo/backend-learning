# use official Golang image
FROM golang:1.24rc2-alpine3.21
# Что это за инструкция?
# A base image is an image you designate (указываете) in a FROM directive in a Dockerfile. 
# It defines the starting point for your build. Dockerfile instructions create additional layers on top of the base image. 
# A Dockerfile with the FROM scratch directive uses an empty base image.
# То есть, мы наследуемся(?)/копируем(?) исходный образ, и уже на основе его добавляем свои инструкции

# set working directory
WORKDIR /app
# Что это за инструкция?
# Это папка внутри контейнера или внутри папки, где находится Dockerfile?
# В коде примера, который я смотрю, нет папки /app, то есть эта папка создается внутри контейнера, наверное
# WORKDIR instruction is used to set the working directory for all the subsequent Dockerfile instructions.
# # # # # # # # # # # # # # # # # # # # # # # # # # #
# All paths in a Dockerfile are always inside the image, except for the source paths for COPY and ADD instructions, 
# which are inside the build context directory on the host. 
# Absolute paths like /code will be directly inside the root directory in the image, following normal Unix conventions.
# # # # # # # # # # # # # # # # # # # # # # # # # # #
# То есть, да, внутри образа, в корне файловой системы создается новая папка, и всё, что происходит после этой инструкции, 
# происходит внутри этой папки. То есть, внутри папки будет docker-compose.yml, Dockerfile, go.mod, go.sum, main.go, 
# выполнится go get и т.д.

# Copy the source code
COPY . . 
# COPY <src-path> <destination-path>
# Копируем всё, что находится в той же папке, где и Dockerfile в контейнер, видимо (?)
# The COPY instruction copies new files or directories from <src> and adds them 
# to the filesystem of the image at the path <dest>
# .dockerignore - файл, который позволяет игнорировать выполнение этой инструкции для указанных в нём файлов.

# Отсюда, видимо, работаем уже изнутри контейнера
# А, или и так работаем уже изнутри контейнера, а COPY берет файлы извне контейнера.

# Download and install the dependencies
RUN go get -d -v ./...

# Build the Go app
RUN go build -o api .

#EXPOSE the port
EXPOSE 8000
# Это порт контейнера, который выставляем снаружи (?) 
# то есть, контейнер - это компьютер, и мы открываем порт 8000 для внешних сообщений (?)

# Run the executable
CMD ["./api"]

###################################
# Термины и команды

# И образ и контейнер - это просто файлы, внутри которых происходит 

# Образ Docker (Docker Image) - это шаблон, на основе которого создаются контейнеры. Это начальное, исходное состояние приложения. 
# Он содержит исходный код, библиотеки, зависимости, инструменты и другие файлы, в том числе и, к примеру, операционную систему 
# (её программу, видимо), всё, что необходимо для запуска приложения. 

# На основе образа создается контейнер - уже изменяемая сущность. В них запускается операционная система, имеются своя файловая система, 
# и т.д.
# Контейнеры можно запускать, останавливать, переносить, удалять и т.д.

# Создание образа 
# docker build .
# The final . in the command provides the path or URL to the build context. 
# At this location, the builder will find the Dockerfile and other referenced files.
# После выполнения инструкции, образ будет создан (но никакого контейнера не будет), 
# и его можно использовать для создания контейнера.

# Образу можно приписать тег (как при создании, так и потом).
# Точнее, так: образ однозначно описывается своим идентификатором (Image ID). Текстовый идентификатор конкретного образа 
# описывается как repository_name:tag_name, где repository_name - это имя репозитория - группы образов, а tag - название образа
# в данной группе. Эти параметры - необязательные.
# Создание образа с тегом (с репозиторием, видимо, нельзя):
# docker build -t my-image .
