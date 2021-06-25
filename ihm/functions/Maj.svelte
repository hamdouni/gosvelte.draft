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
  <h1 class="title">Majuscule</h1>
  <div class="field">
    <label class="label" for="name">
      Entrer le mot ou la phrase à mettre en majuscule
    </label>
    <div class="control has-icons-left">
      <input
        bind:value={nom}
        name="name"
        type="text"
        placeholder="Mot ou phrase..."
        class="input" />
      <span class="icon is-small is-left">
        <i class="far fa-keyboard" />
      </span>
    </div>
  </div>
  <div class="field">
    <button on:click={maj} class="button is-success">Majuscule</button>
  </div>
  {#if resultat}
    <hr />
    {resultat}
  {/if}
</div>