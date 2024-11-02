Методичка с заданиями для лабараторных работ: 
https://docs.google.com/document/d/1ApTU6C1dMvFEM2OlxrktOCrC0EpM9DUjS7OXLEwWjeo/edit?tab=t.0

Чтобы запустить каждую работу необходимо перейти в директорию этой работы и написать go run . \ go run main.go
Для 7й лабароторной необходимо так же перейти в директорию сервера\клиента

Для того чтобы отправить что-то с клиента на клиент с помощью сервера на WebSocket'ах необходимо:
1. зайти на любую страницу браузера
2. прописать в консоль следующее: let socket = new WebSocket("ws://localhost:*PORT*/ws")
3. socket.onmessage = (event) => {console.log("from server: ", event.data)}
4. после чего можно отправить любое сообщение с помощью: socket.send("message")
