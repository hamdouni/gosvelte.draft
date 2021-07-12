<script>
  export let connectedStatus = "";

  let username;
  let password;
  let err;

  async function login() {
    let response = await fetch("/login", {
      method: "POST",
      body: "username=" + username + "&password=" + password,
      headers: { "Content-Type": "application/x-www-form-urlencoded" }
    });
    if (response.ok) {
      connectedStatus = true;
    } else {
      connectedStatus = false;
      err =
        "Echec de la connexion. Veuillez vérifier votre identifiant et votre mot de passe. Si le problème persiste, merci de contacter l'administrateur.";
    }
  }
</script>

<section >
  <main class="login">
    <h1>
      <span class="icon">
        <i class="las la-cube la-fw" />
      </span>
      <span>Webtoolkit</span>
    </h1>
    <!-- svelte-ignore a11y-autofocus-->
    <form on:submit|preventDefault={login} action="/login" method="post" class="box">
      <div class="field icon">
        <label for="username" class="label">E-mail</label>
        <input type="email" name="username" bind:value={username} autofocus="autofocus" />
        <i class="las la-user" />
      </div>
      <div class="field icon">
        <label for="password" class="label">Mot de passe</label>
          <input type="password" name="password" bind:value={password} />
            <i class="las la-unlock" />
      </div>
      <div class="field">
        <button on:click|preventDefault={login}>
          Connecter
        </button>
      </div>
      {#if err}
        <article class="message alert">
          {err}
        </article>
      {/if}
    </form>
  </main>
</section>
