<script>
  import Error from "./components/Error.svelte";

  export let connectedStatus = "";

  let username;
  let password;
  let errMessage;

  async function login() {
	try {
		let response = await fetch("/login", {
		  method: "POST",
		  body: "username=" + username + "&password=" + password,
		  headers: { "Content-Type": "application/x-www-form-urlencoded" }
		});
		if (response.ok) {
		  connectedStatus = true;
		} else {
		  connectedStatus = false;
		  errMessage =
			"Echec de la connexion. Veuillez vérifier votre identifiant et votre mot de passe. Si le problème persiste, merci de contacter l'administrateur.";
		}
	}
	catch(error) {
		errMessage = "Problème réseau: impossible de contacter le serveur";
	}
  }
</script>

<section >
  <main class="login">

{#if errMessage}
<Error message="{errMessage}"/>
{/if}

    <h1>
      <span class="icon">
        <i class="las la-cube la-fw" />
      </span>
      <span>Webtoolkit</span>
    </h1>
    <!-- svelte-ignore a11y-autofocus-->
    <form on:submit|preventDefault={login} action="/login" method="post" class="box">
      <label for="username" class="label">E-mail</label>
      <div class="field icon">
        <input type="email" name="username" bind:value={username} autofocus="autofocus" />
        <i class="las la-user" />
      </div>
      <label for="password" class="label">Mot de passe</label>
      <div class="field icon">
          <input type="password" name="password" bind:value={password} />
            <i class="las la-unlock" />
      </div>
      <div class="field">
        <button on:click|preventDefault={login}>
          Connecter
        </button>
      </div>
    </form>
  </main>
</section>
