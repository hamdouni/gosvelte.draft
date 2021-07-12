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
  <h1>Majuscule</h1>
  <div class="field icon">
    <label class="label" for="name">
      Entrer le mot ou la phrase à mettre en majuscule
    </label>
    <input bind:value={nom} name="name" type="text" placeholder="Mot ou phrase..." class="input" />
    <i class="las la-keyboard" />
  </div>
  <div class="field">
    <button on:click={maj}>Majuscule</button>
  </div>
  {#if resultat}
    <hr />
    {resultat}
  {/if}
</div>