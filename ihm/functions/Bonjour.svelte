<script>
    import * as net from "../lib/network.js";
    let nom;
    let resultat;

    async function bonjour() {
        let res = await net.callBonjour(nom);
        if (res != null) resultat = res;
    }
</script>

<div>
    <h1 class="title is-2">Entrer un prénom</h1>
    <input bind:value="{nom}" name="name" type="text" placeholder="Entrer un prénom..."
        class="input">
    <button on:click="{bonjour}">
        Bonjour
    </button>
    <hr>
    {#if resultat}
    {resultat}
    {/if}
</div>