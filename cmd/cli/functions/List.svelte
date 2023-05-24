<script>
	import * as net from "../lib/network.js";
    import Error from "../components/Error.svelte";

	let users = []; 
	async function list() {
		let res = await net.callGetUsers()
		if (res.error != null) {
			console.log("error:", res.error);
			return
		}
		users = res.data;
	}
	list();

</script>

<Error message=""/>

<h1>Liste des utilisateurs</h1>

<div class="grid grid-cols-3 gap-4">
{#each users as user}
	<div> {user.Username} </div>
	<div> {user.Role} </div>
	<div> {user.Realm} </div>
{/each}
</div>
