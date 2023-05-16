<script>
  import * as net from "./lib/network.js";
  import Error from "./components/Error.svelte";

  let resultat;
  let nom = "";
  let errMessage = "";

  async function bonjour() {
	if(nom == "") {
		errMessage = "Un prénom doit être renseigné";
		return;
	}
    let res = await net.callBonjour(nom);
	if (res.error != null) {
		console.log("error:", res.error);
		return
	}
	resultat = res.response;
  }
</script>

<div>

  <h1>Bonjour</h1>

  {#if errMessage}
  <Error message={errMessage}/>
  {/if}

  <label for="name">Entrer un prénom</label>
  <div class="field icon">
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
