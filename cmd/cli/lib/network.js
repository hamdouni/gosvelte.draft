export async function callApi(url) {
    try {
        let response = await fetch(url);
        if(response.ok) {
            return {
				error: null,
				response: await response.json()
			};
        }
		return {
				error: response.status,
				response: null
		};
    } catch (error) {
		return {
				error: error,
				response: null
		};
    }
}

export async function callCheckConnexion() {
	let response;
	try {
		response = await fetch("check");
		return response.ok;
	} catch (error) {
        console.log("Erreur réseau " + error);
		return null;
	}
}

export async function callBonjour(nom) {
    return await callApi("/hello?nom="+nom);
}

export async function callGetUsers() {
	return await callApi("/list");
}
