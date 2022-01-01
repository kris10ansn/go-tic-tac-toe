export function fetchAddGame({ name }) {
    return fetch(`/game/add`, {
        method: "POST",
        body: JSON.stringify({ name }),
    }).then((it) => it.json());
}

export function fetchGames() {
    return fetch(`/game/list`)
        .then((response) => response.json())
        .catch((error) => {
            console.error(error);
            return [];
        });
}
