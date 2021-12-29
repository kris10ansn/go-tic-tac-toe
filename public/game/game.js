function connectGameSocket() {
    const socket = new WebSocket(`ws:/${location.host}/socket/game`);

    socket.addEventListener("open", () => {
        console.log("Websocket connected!");

        const gameId = new URLSearchParams(location.search).get("id");

        socket.send(
            JSON.stringify({
                gameId,
            })
        );

        socket.addEventListener("message", (msg) => {
            const message = JSON.parse(msg.data);

            switch (message.type) {
                case "assign-tic": {
                    console.log(`Assigned tic ${message.data}`);
                    break;
                }

                case "error": {
                    window.alert(`Websocket error: ${message.data}`);
                    location.href = "/";
                    break;
                }
            }
        });

        socket.addEventListener("close", (event) => {
            console.log("Websocket connection closed", event);
        });

        socket.addEventListener("error", (error) => {
            console.error("Websocket error", error);
        });
    });
}

function main() {
    connectGameSocket();
}

main();
