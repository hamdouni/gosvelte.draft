<script>
    import * as net from "../lib/network.js";

    let historique = [];

    async function reloadHistoric() {
        let res = await net.callHistoric();
        if(res != null) historique = res;
    }
    reloadHistoric();
</script>

<div>
    <h1>Historique</h1>
    <ul>
        {#each historique as evenement}
            <li>{evenement}</li>
        {/each}
    </ul>
</div>