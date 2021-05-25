<script>
import { createEventDispatcher } from 'svelte';
const dispatch = createEventDispatcher();

let username;
let password;

async function login() {
    let response = await fetch("/login",{
        method: "POST",
        body: 'username='+username+'&password='+password,
        headers: {"Content-Type": "application/x-www-form-urlencoded"}
    });
    if(response.ok) {
        dispatch('connected');
    }
}
</script>

<form on:submit|preventDefault={login} action="/login" method="post">
    <input type="text" name="username" placeholder="Utilisateur" bind:value={username}>
    <input type="password" name="password" placeholder="Mot de passe" bind:value={password}>
    <button type="submit">Connecter</button>
</form>