<script>
    import * as net from "../lib/network.js";
    let nom;
    let resultat;

    async function min() {
        let res = await net.callMin(nom);
        if (res != null) resultat = res;
    }
</script>

<div>
    <h1 class="title is-2">Entrer le mot ou la phrase à mettre en minuscule</h1>
    <input bind:value="{nom}" type="text" placeholder="Entrer le mot ou la phrase à mettre en minuscule..."
        class="input">
    <button on:click="{min}">
        Minuscule
    </button>
    <hr>
    {#if resultat}
    {resultat}
    {/if}
</div>