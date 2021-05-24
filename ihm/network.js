export async function callApi(endpoint) {
    let url = endpoint;
    try {
        let response = await fetch(url);
        if(response.ok) {
            return response.json();
        } else {
            console.log("Erreur http " + response.status + " sur url " + url);
        }
    } catch (error) {
        console.log("Erreur réseau " + error);
    }
}

export async function callCheckConnexion() {
    let response = await fetch("check");
    return response.ok;
}

export async function callBonjour(nom) {
    return await callApi("/hello?nom="+nom);
}
export async function callMaj(nom) {
    return await callApi("/upper?nom="+nom);
}
export async function callMin(nom) {
    return await callApi("/lower?nom="+nom);
}
export async function callHistoric() {
    return await callApi("/historic");
}
