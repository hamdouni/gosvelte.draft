<script>
	import Bonjour from './Bonjour.svelte';
	import Maj from './Maj.svelte';
	import Min from './Min.svelte';
	import Historic from './Historic.svelte';

	let Menu = [
		{ id: "bonjour", component: Bonjour, icon: "fa-handshake", label: "Bonjour" },
		{ id: "maj", component: Maj, icon: "fa-chart-bar", label: "Majuscule" },
		{ id: "min", component: Min, icon: "fa-compass", label: "Minuscule" },
		{ id: "historic", component: Historic, icon: "fa-credit-card", label: "Historique" },
	];

	let showNavbarMenu = false;
	let actualMenu = Menu[0];

	function activate(menuID) {
		actualMenu = Menu.find(elem => elem.id === menuID);
		showNavbarMenu = !showNavbarMenu;
	}
</script>

<nav class="navbar is-fixed-top has-shadow is-light">
	<div class="navbar-brand">
		<div class="navbar-item">
			<span class="icon-text">
				<span>
					<i class="far fa-fw fa-gem"></i>
				</span>
				<span>Webtoolkit</span>
			</span>
		</div>
		<div class="navbar-burger is-hoverable" on:click="{() => showNavbarMenu=!showNavbarMenu}">
			<span></span>
			<span></span>
			<span></span>
		</div>
	</div>
	<div class="navbar-menu" class:is-active="{showNavbarMenu}">
		<div class="navbar-start">
			{#each Menu as item (item.id)}
			<a href="#{item.id}" class="navbar-item" on:click="{()=>activate(item.id)}">
				<span>
					<i class="far fa-fw {item.icon}"></i>
				</span>
				<span> {item.label} </span>
			</a>
			{/each}
		</div>
		<div class="navbar-end">
			<div class="navbar-item has-dropdown is-hoverable">
				<span class="navbar-link is-arrowless">
					<span class="icon"> <i class="far fa-user-circle"></i> </span>
				</span>
				<div class="navbar-dropdown is-right">
					<span class="navbar-item">Maximilien</span>
					<hr class="navbar-divider">
					<a class="navbar-item" href="/logout">Se déconnecter</a>
				</div>
			</div>
		</div>
	</div>
</nav>
<section class="section" is-main-content>
	<svelte:component this={actualMenu.component} />
</section>

<style>
	@media screen and (max-width: 1023px) {
		.navbar-menu {
			position: absolute;
			max-width: max-content;
			right: 0;
		}
	}
</style>