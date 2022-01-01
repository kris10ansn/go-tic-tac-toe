export function confirmReload(message) {
    window.confirm(message) && window.location.reload();
}

export function confirmRedirect(message, path) {
    window.confirm(message) && navigate(path);
}

export function navigate(url) {
    window.location.href = url;
}
