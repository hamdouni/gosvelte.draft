{#if !connectedStatus}
	<Login bind:connectedStatus/>
{:else}
	<Dashboard/>
{/if}

<script>
	import Login from './functions/Login.svelte';
	import Dashboard from './functions/Dashboard.svelte';
	import * as cookie from './lib/cookie.js';
	import * as net from "./lib/network.js";

	let connectedStatus = false;
	let idCookieValue = cookie.get("jeton");

	if(idCookieValue === "") {
		// pas de jeton => l'utilisateur n'est pas connecté
		connectedStatus = false;
	} else {
		// si on a un jeton, on vérifie auprès du serveur
		net.callCheckConnexion().then(response => {
			connectedStatus = response;
		});
	}
</script>
