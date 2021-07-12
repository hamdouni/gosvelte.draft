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
  <h1>Minuscule</h1>
  <label for="name" class="label">
    Entrer le mot ou la phrase à mettre en minuscule
  </label>
  <div class="field icon">
    <input bind:value={nom} name="name" type="text" placeholder="Entrer le mot ou la phrase à mettre en minuscule..." class="input" />
    <i class="las la-keyboard" />
  </div>
  <div class="field">
    <button on:click={min}>Minuscule</button>
  </div>
  {#if resultat}
    <hr />
    {resultat}
  {/if}
</div>
