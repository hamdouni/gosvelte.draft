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

  <h1 class="title">Bonjour</h1>

  <div class="field">
    <label class="label" for="name">Entrer un prénom</label>
    <div class="control has-icons-left">
      <input
        bind:value={nom}
        name="name"
        type="text"
        placeholder="Prénom..."
        class="input" />
      <span class="icon is-small is-left">
        <i class="fas fa-user" />
      </span>
    </div>
  </div>
  <div class="field">
    <button on:click={bonjour} class="button is-success">Bonjour</button>
  </div>
  {#if resultat}
    <hr />
    {resultat}
  {/if}
</div>
