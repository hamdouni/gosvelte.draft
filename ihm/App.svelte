<script>
    import Login from "./Login.svelte";
    import * as net from "./network.js";

    let nom = "";
    let resultat = "";
    let historique = [];
    let isConnected = false;

    async function bonjour() {
        let res = await net.callBonjour(nom);
        if(res != null) resultat = res;
        reloadHistoric();
    }
    async function maj() {
        let res = await net.callMaj(nom);
        if(res != null) resultat = res;
        reloadHistoric();
    }
    async function min() {
        let res = await net.callMin(nom);
        if(res != null) resultat = res;
        reloadHistoric();
    }
    async function reloadHistoric() {
        if(isConnected) {
            let res = await net.callHistoric();
            if(res != null) historique = res;
        }
    }
    reloadHistoric();

    async function checkConnexion() {
        let res = await net.callCheckConnexion();
        if(res) {
            isConnected = true;
            console.log("HOURA");
        } else {
            isConnected = false;
            console.log("BOUHOU");
        }
    }

</script>

<pre class="toto">
    Bonjour depuis Svelte.
</pre>

<hr>
    |
    Connexion : {isConnected} (<a href="#check" on:click={checkConnexion}>Check</a>)

{#if !isConnected}
    <Login 
        on:connected={function(){
            isConnected=true;
        }}
        on:notconnected={function(){
            alert("Votre connexion a échoué");
        }}
    />
{:else}
    <a href="/logout">Logout</a>
    <hr>
    <input bind:value="{nom}" type="text" placeholder="Entrer votre prénom...">
    <button on:click="{bonjour}">
        Bonjour XXX
    </button>
    <button on:click="{maj}">
        Majuscule
    </button>
    <button on:click="{min}">
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
{/if}


<style>
    pre {
        background-color: #eee;
    }
</style>
