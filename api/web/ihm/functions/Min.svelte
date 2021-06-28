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
  <h1 class="title">Minuscule</h1>
  <div class="field">
    <label for="name" class="label">
      Entrer le mot ou la phrase à mettre en minuscule
    </label>
    <div class="control has-icons-left">
      <input
        bind:value={nom}
        name="name"
        type="text"
        placeholder="Entrer le mot ou la phrase à mettre en minuscule..."
        class="input" />
      <span class="icon is-small is-left">
        <i class="far fa-keyboard" />
      </span>
    </div>
  </div>
  <div class="field">
    <button on:click={min} class="button is-success">Minuscule</button>
  </div>
  {#if resultat}
    <hr />
    {resultat}
  {/if}
</div>
