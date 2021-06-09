<script>
    import * as net from "../lib/network.js";
    let nom;
    let resultat;

    async function maj() {
        let res = await net.callMaj(nom);
        if (res != null) resultat = res;
    }
</script>

<div>
    <h1 class="title is-2">Entrer le mot ou la phrase à mettre en majuscule</h1>
    <input bind:value="{nom}" type="text" placeholder="Entrer le mot ou la phrase à mettre en majuscule..."
        class="input">
    <button on:click="{maj}">
        Majuscule
    </button>
    <hr>
    {#if resultat}
    {resultat}
    {/if}
</div>