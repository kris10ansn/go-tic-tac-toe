export function extractFormData(form) {
    return Object.fromEntries(new FormData(form).entries());
}

export function createElement(htmlString) {
    const template = document.createElement("template");
    template.innerHTML = htmlString;
    return template.content;
}

export function searchParam(name) {
    return new URLSearchParams(location.search).get(name);
}

export function copyMatrix(matrixFrom, matrixTo) {
    for (let y = 0; y < 3; y++) {
        for (let x = 0; x < 3; x++) {
            matrixTo[y][x] = matrixFrom[y][x];
        }
    }
}
