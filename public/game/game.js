const gameTable = document.querySelector("table#game tbody");

const X_TIC = 11;
const O_TIC = 13;
const EMPTY_TIC = 0;

const globalState = {
    tic: null,
    board: [
        [EMPTY_TIC, EMPTY_TIC, EMPTY_TIC],
        [EMPTY_TIC, EMPTY_TIC, EMPTY_TIC],
        [EMPTY_TIC, EMPTY_TIC, EMPTY_TIC],
    ],
};

function ticToString(tic) {
    switch (tic) {
        case X_TIC:
            return "x";
        case O_TIC:
            return "o";
        case EMPTY_TIC:
            return " ";
        default:
            throw new Error(`Unknown tic ${tic}!`);
    }
}

function renderBoard(board, tableElement) {
    for (let y = 0; y < 3; y++) {
        for (let x = 0; x < 3; x++) {
            const cell = tableElement.children[y].children[x];
            cell.textContent = ticToString(board[y][x]);
        }
    }
}

function moveFromClickEvent(event) {
    if (event.target instanceof HTMLTableCellElement) {
        const cell = event.target;
        const row = cell.parentElement;

        const y = [...row.parentElement.children].indexOf(row);
        const x = [...cell.parentElement.children].indexOf(cell);

        return { x, y };
    }

    return null;
}

function connectGameSocket() {
    const socket = new WebSocket(`ws:/${location.host}/socket/game`);

    socket.addEventListener("open", () => {
        console.log("Websocket connected!");

        const gameId = new URLSearchParams(location.search).get("id");

        socket.send(
            JSON.stringify({
                type: "join",
                data: { gameId },
            })
        );

        socket.addEventListener("message", (msg) => {
            const message = JSON.parse(msg.data);

            console.log(message.type, message.data);

            switch (message.type) {
                case "assign-tic": {
                    globalState.tic = message.data;
                    break;
                }

                case "present-board": {
                    for (let y = 0; y < 3; y++) {
                        for (let x = 0; x < 3; x++) {
                            globalState.board[y][x] = message.data[y][x];
                        }
                    }

                    renderBoard(globalState.board, gameTable);
                    break;
                }

                case "await-move": {
                    const handler = (event) => {
                        const move = moveFromClickEvent(event);

                        if (move != null) {
                            socket.send(
                                JSON.stringify({
                                    type: "move",
                                    data: move,
                                })
                            );

                            game.removeEventListener("click", handler);
                        }
                    };
                    game.addEventListener("click", handler);
                    break;
                }

                case "end-game": {
                    break;
                }

                case "error": {
                    window.confirm(`Websocket error: ${message.data}`) &&
                        (location.href = "/");
                    break;
                }
            }
        });

        socket.addEventListener("close", (event) => {
            window.confirm("Websocket connection closed") &&
                (location.href = "/");
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
