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

<nav class="navbar is-fixed-top has-shadow is-light">
	<div class="navbar-brand">
		<div class="navbar-item">
			<span class="icon is-medium">
				<i class="far fa-gem"></i>
			</span>
			<span>Webtoolkit</span>
		</div>
  </div>
</nav>
<section class="hero is-primary is-fullheight-with-navbar">
  <div class="hero-body">
    <div class="container">
      <div class="columns is-centered">
        <div class="column is-5-tablet is-4-desktop is-3-widescreen">
          <p class="title">Web Tool Kit</p>
          <!-- svelte-ignore a11y-autofocus-->
          <form
            on:submit|preventDefault={login}
            action="/login"
            method="post"
            class="box ">
            <div class="field">
              <label for="username" class="label">Identifiant</label>
              <div class="control has-icons-left has-icons-right">
                <input
                  class="input"
                  type="text"
                  name="username"
                  bind:value={username}
                  autofocus="autofocus" />
                <span class="icon is-small is-left">
                  <i class="far fa-user" />
                </span>
              </div>
            </div>
            <div class="field">
              <label for="password" class="label">Mot de passe</label>
              <div class="control has-icons-left has-icons-right">
                <input
                  class="input"
                  type="password"
                  name="password"
                  bind:value={password} />
                <span class="icon is-small is-left">
                  <i class="far fa-unlock-alt" />
                </span>
              </div>
            </div>
            <div class="field">
              <button class="button is-success" on:click|preventDefault={login}>
                Connecter
              </button>
            </div>
            <div>
              {#if err}
                <article class="message is-danger">
                  <div class="message-body">
                    <strong>{err}</strong>
                  </div>
                </article>
              {/if}
            </div>
          </form>
        </div>
      </div>
    </div>
  </div>
</section>

<style>
  .title {
    font-family: 'Oleo Script Swash Caps', cursive;
    font-size: 3rem;
  }
</style>