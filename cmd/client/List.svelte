<script>
	import * as net from "./lib/network.js";
    import Error from "./components/Error.svelte";

	let users = []; 
	async function list() {
		let res = await net.callGetUsers()
		if (res.error != null) {
			console.log("error:", res.error);
			return
		}
		users = res.response;
	}
	list();

</script>

<Error message=""/>

<h1>Liste des utilisateurs</h1>

<ul>
{#each users as user}
<li> {user.Username}	{user.Role}	{user.Realm} </li>
{/each}
</ul>
