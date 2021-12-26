function createElement(htmlString) {
    const element = new DOMParser().parseFromString(string)
    return element.documentElement.querySelector("body").firstChild
}

function main() {
    const gamesSocket = new WebSocket(`ws:/${location.host}/socket/games`)

    gamesSocket.addEventListener("open", () => {
        console.log("Websocket connected!")
    }) 
}

main()