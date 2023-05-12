<script>
	import * as net from "./lib/network.js";

	let users = "rien"; 
	async function list() {
		let res = await fetch("/list")
		if(res.ok) {
			users = await res.json();
		}
	}
	list();

</script>

<h1>Liste des utilisateurs</h1>

<ul>
{#each users as user}
<li> {user.Username}	{user.Role}	{user.Realm} </li>
{/each}
</ul>
