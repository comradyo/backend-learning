# Знакомство с вебсокетами

### basic-websockets
Делал по https://www.golinuxcloud.com/golang-websocket/

Как тестировать:
1. Запустить Postman
2. Создать новый запрос:
    1. `New` -> `WebSocket`
    2. Ввести в качестве url `ws://localhost:8080/`
3. В message вводить данные, они будут приходить в ответе

### go-socket-io
https://github.com/googollee/go-socket.io

Делал по https://github.com/googollee/go-socket.io/blob/master/_examples/go-echo/main.go

Методы:
1. `OnConnect` - действие, выполняемое при установке соединения
2. `Emit` = send (отослать ответ)
3. `Close` - закрытие соединения
4. `BroadcastToNamespace` - рассылка всему namespace, 

### zishang
https://github.com/zishang520/socket.io

Разобраться, настроить, понять, насколько удобно
