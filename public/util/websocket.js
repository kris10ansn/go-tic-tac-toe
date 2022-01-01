export function websocketSend(socket, object) {
    socket.send(JSON.stringify(object));
}
