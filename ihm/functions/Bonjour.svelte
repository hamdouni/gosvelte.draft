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
    <input bind:value="{nom}" type="text" placeholder="Entrer un prénom..."
        class="input">
    <button on:click="{bonjour}">
        Bonjour
    </button>
    <hr>
    {#if resultat}
    {resultat}
    {/if}
</div>