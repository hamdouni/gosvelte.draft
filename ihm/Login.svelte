<script>
export let connectedStatus = '';

let username;
let password;
let err;

async function login() {
    let response = await fetch("/login",{
        method: "POST",
        body: 'username='+username+'&password='+password,
        headers: {"Content-Type": "application/x-www-form-urlencoded"}
    });
    if(response.ok) {
        connectedStatus = true;
    } else {
        connectedStatus = false;
        err = "Echec de la connexion. Veuillez vérifier votre identifiant et votre mot de passe. Si le problème persiste, merci de contacter l'administrateur.";
    }
}
</script>

<login class="card">
    <!-- svelte-ignore a11y-autofocus-->
    <form on:submit|preventDefault={login} action="/login" method="post" class="card-content">
        <div class="field">
            <label for="username" class="label">Identifiant</label>
            <div class="control has-icons-left has-icons-right">
                <input class="input" type="text" name="username" bind:value={username} autofocus>
                <span class="icon is-small is-left">
                    <i class="fas fa-user"></i>
                </span>
            </div>
        </div>
        <div class="field">
            <label for="password" class="label">Mot de passe</label>
            <div class="control has-icons-left has-icons-right">
                <input class="input" type="password" name="password" bind:value={password}>
                <span class="icon is-small is-left">
                    <i class="fas fa-unlock-alt"></i>
                </span>
            </div>
        </div>
        <div class="buttons">
            <button class="button is-primary is-large" on:click|preventDefault={login}>
                <strong>
                    Connecter
                </strong>
            </button>
        </div>
        <div>
            {#if err}
            <article class="message is-danger">
                <div class="message-body">
                    <strong>
                        {err}
                    </strong>
                </div>
            </article>
            {/if}
        </div>
    </form>
</login>