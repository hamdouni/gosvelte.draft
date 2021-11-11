<script>
	import Login from './functions/Login.svelte';
	import Dashboard from './functions/Dashboard.svelte';
	import { getCookie } from './lib/cookie.js';
	import { callCheckConnexion } from "./lib/network.js";

	// pas de jeton => l'utilisateur n'est pas connecté
	// si on a un jeton, on vérifie auprès du serveur
	let connectedStatus = true;
	if(getCookie("jeton") === "") connectedStatus = false; 
	else callCheckConnexion().then(response => { connectedStatus = response; });	
</script>

{#if !connectedStatus}
	<Login bind:connectedStatus/>
{:else}
	<Dashboard/>
{/if}
