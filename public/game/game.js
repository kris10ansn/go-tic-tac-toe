import { confirmRedirect } from "../util/navigation.js";
import { copyMatrix, searchParam } from "../util/util.js";
import { websocketSend } from "../util/websocket.js";

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
    if (tic === X_TIC) return "x";
    if (tic === O_TIC) return "o";
    if (tic === EMPTY_TIC) return " ";

    throw new Error(`Unknown tic ${tic}!`);
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

function sendJoinMessage(socket, gameId) {
    websocketSend(socket, { type: "join", data: { gameId } });
}

function sendMoveMessage(socket, move) {
    websocketSend(socket, { type: "move", data: move });
}

function connectGameSocket() {
    const socket = new WebSocket(`ws:/${location.host}/socket/game`);

    socket.addEventListener("open", () => {
        sendJoinMessage(socket, searchParam("id"));

        socket.addEventListener("message", ({ data: dataString }) => {
            const { type, data } = JSON.parse(dataString);

            if (type === "assign-tic") globalState.tic = data;
            if (type === "present-board") {
                copyMatrix(data, globalState.board);
                renderBoard(globalState.board, gameTable);
            }

            if (type === "await-move") {
                const game = document.querySelector("table#game");

                game.addEventListener("click", function handler(event) {
                    const move = moveFromClickEvent(event);

                    if (move != null) {
                        sendMoveMessage(socket, move);
                        game.removeEventListener("click", handler);
                    }
                });
            }

            if (type === "end-game") {
            }

            if (type === "error") confirmRedirect(`Error: ${data}`, "/");
        });

        socket.addEventListener("close", () =>
            confirmRedirect("Websocket connection closed", "/")
        );

        socket.addEventListener("error", (error) =>
            console.error("Websocket error", error)
        );
    });
}

function main() {
    connectGameSocket();
}

main();
