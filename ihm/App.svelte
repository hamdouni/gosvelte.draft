<script>
    let nom = "";
    let resultat = "";
    let historique = [];
    async function callApi(endpoint) {
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
    async function callBonjour() {
        let res = await callApi("/hello?nom="+nom);
        if(res != null) resultat = res;
        reloadHistoric();
    }
    async function callMaj() {
        let res = await callApi("/upper?nom="+nom);
        if(res != null) resultat = res;
        reloadHistoric();
    }
    async function callMin() {
        let res = await callApi("/lower?nom="+nom);
        if(res != null) resultat = res;
        reloadHistoric();
    }
    async function reloadHistoric() {
        let res = await callApi("/historic");
        if(res != null) historique = res;
    }
    reloadHistoric();
</script>

<pre class="toto">
    Bonjour depuis Svelte.
</pre>
<hr>
<input bind:value="{nom}" type="text" placeholder="Entrer votre prénom...">
<button on:click="{callBonjour}">
    Bonjour
</button>
<button on:click="{callMaj}">
    Majuscule
</button>
<button on:click="{callMin}">
    Minuscule
</button>
<hr>
<h1>{resultat}</h1>
<hr>
<ul>
    {#each historique as evenement}
        <li>{evenement}</li>
    {/each}
</ul>

<style>
    pre {
        background-color: #eee;
    }
</style>
