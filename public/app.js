function createElement(htmlString) {
    const element = new DOMParser().parseFromString(string);
    return element.documentElement.querySelector("body").firstChild;
}

function main() {
    const gamesSocket = new WebSocket(`ws:/${location.host}/socket/games`);

    gamesSocket.addEventListener("open", () => {
        console.log("Websocket connected!");

        gamesSocket.addEventListener("message", (message) => {
            console.log("Websocket message", message);
        });

        gamesSocket.addEventListener("close", (event) => {
            console.log("Websocket connection closed", event);
        });

        gamesSocket.addEventListener("error", (error) => {
            console.error("Websocket error", error);
        });
    });
}

main();
