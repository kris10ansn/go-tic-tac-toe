function createElement(htmlString) {
    const template = document.createElement("template");
    template.innerHTML = htmlString;
    return template.content;
}

function appendGameToList({ name, id }) {
    const gamesList = document.querySelector("table#games tbody");
    gamesList.appendChild(
        createElement(`
            <tr id="${id}">
                <td>${name}</td>
                <td>
                    <a href="/game?id=${id}">
                        <button>join game</button>
                    </a>
                </td>
            </tr>
        `)
    );
}

function handleCreateGameSubmits() {
    const createGameForm = document.forms.namedItem("create-game");

    createGameForm.addEventListener("submit", async (event) => {
        event.preventDefault();

        const data = Object.fromEntries(new FormData(createGameForm).entries());
        createGameForm.reset();

        const response = await fetchAddGame(data);

        location.href = `/game?id=${response.id}`;
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
            const message = "Websocket connection closed";

            window.confirm(message) && window.location.reload();
            console.log(message, event);
        });

        gamesSocket.addEventListener("error", (error) => {
            console.error("Websocket error", error);
        });
    });
}

async function fetchAddGame({ name }) {
    return fetch(`/game/add`, {
        method: "POST",
        body: JSON.stringify({ name }),
    }).then((it) => it.json());
}

async function fetchGames() {
    return await fetch(`/game/list`)
        .then((response) => response.json())
        .catch((error) => {
            console.error(error);
            return [];
        });
}

async function main() {
    const games = await fetchGames();
    games.forEach((game) => appendGameToList(game));

    connectGamesSocket();
    handleCreateGameSubmits();
}

main();
