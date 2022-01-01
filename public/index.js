import { fetchAddGame, fetchGames } from "./util/fetch.js";
import { confirmReload, navigate } from "./util/navigation.js";
import { extractFormData, createElement } from "./util/util.js";

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
    const formCreateGame = document.forms.namedItem("create-game");

    formCreateGame.addEventListener("submit", (event) => {
        event.preventDefault();
        const { name } = extractFormData(formCreateGame);

        fetchAddGame({ name })
            .then(({ id }) => navigate(`/game?id=${id}`))
            .catch(console.error);

        formCreateGame.reset();
    });
}

function connectGamesSocket() {
    const gamesSocket = new WebSocket(`ws:/${location.host}/socket/games`);

    gamesSocket.addEventListener("open", () => {
        gamesSocket.addEventListener("message", ({ data: dataString }) => {
            const { type, data } = JSON.parse(dataString);

            if (type === "add-game") appendGameToList(data);
        });

        gamesSocket.addEventListener("close", () =>
            confirmReload("Websocket connection closed")
        );

        gamesSocket.addEventListener("error", (error) =>
            console.error("Websocket error", error)
        );
    });
}

async function main() {
    fetchGames().then((games) => games.forEach(appendGameToList));

    connectGamesSocket();
    handleCreateGameSubmits();
}

main();
