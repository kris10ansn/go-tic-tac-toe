function createElement(htmlString) {
    const element = new DOMParser().parseFromString(htmlString, "text/html");
    return element.documentElement.querySelector("body").firstChild;
}

function appendGameToList({ id, name }) {
    const gamesList = document.querySelector("ul#games");
    gamesList.appendChild(
        createElement(`
            <li id="${id}">
                ${name} <button>join game</button>
            </li>
        `)
    );
}

function handleCreateGameSubmits() {
    const createGameForm = document.forms.namedItem("create-game");

    createGameForm.addEventListener("submit", (event) => {
        event.preventDefault();

        const data = Object.fromEntries(new FormData(createGameForm).entries());

        fetch(`http://${location.host}/game/create`, {
            method: "POST",
            body: JSON.stringify(data),
        });

        createGameForm.reset();
    });
}

function connectGamesSocket() {
    const gamesSocket = new WebSocket(`ws:/${location.host}/socket/games`);

    gamesSocket.addEventListener("open", () => {
        console.log("Websocket connected!");

        gamesSocket.addEventListener("message", (msg) => {
            const message = JSON.parse(msg.data);

            switch (message.type) {
                case "add-game": {
                    appendGameToList(message.data);
                    break;
                }
                default: {
                    console.log(
                        "Websocket message with unknown message type received",
                        message
                    );
                }
            }
        });

        gamesSocket.addEventListener("close", (event) => {
            console.log("Websocket connection closed", event);
        });

        gamesSocket.addEventListener("error", (error) => {
            console.error("Websocket error", error);
        });
    });
}

async function main() {
    const games = await fetch(`http://${location.host}/game/list`).then(
        (response) => response.json()
    );

    games.forEach((game) => appendGameToList(game));

    connectGamesSocket();
    handleCreateGameSubmits();
}

main();
