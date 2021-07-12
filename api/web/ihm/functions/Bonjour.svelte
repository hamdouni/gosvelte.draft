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

  <h1>Bonjour</h1>

  <div class="field icon">
    <label for="name">Entrer un prénom</label>
    <input bind:value={nom} name="name" type="text" placeholder="Prénom..." class="input" />
    <i class="las la-user" />
  </div>
  <div class="field">
    <button on:click={bonjour}>Bonjour</button>
  </div>
  {#if resultat}
    <hr />
    {resultat}
  {/if}
</div>
