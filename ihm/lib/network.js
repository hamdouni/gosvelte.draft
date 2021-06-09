export async function callApi(url) {
    try {
        let response = await fetch(url);
        if(response.ok) {
            console.log("debug");
            console.log(response);
            return response.json();
        }
        console.log("Erreur http " + response.status + " sur url " + url);
        return null;
    } catch (error) {
        console.log("Erreur réseau " + error);
        return null;
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
