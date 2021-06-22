<script>
	import Bonjour  from './Bonjour.svelte';
	import Maj      from './Maj.svelte';
	import Min      from './Min.svelte';
	import Historic from './Historic.svelte';

	let Menu = [
		{id: "bonjour",  component: Bonjour,  icon: "fa-home",          label: "Bonjour"},
		{id: "maj",      component: Maj,      icon: "fa-chart-bar",     label: "Majuscule"},
		{id: "min",      component: Min,      icon: "fa-book",          label: "Minuscule"},
		{id: "historic", component: Historic, icon: "fa-shopping-cart", label: "Historique"},
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
			<span class="icon is-medium">
				<i class="fas fa-gem"></i>
			</span>
			<span>Webtoolkit</span>
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
					<span class="icon is-medium"> <i class="fas {item.icon}"></i> </span>
					<span>{item.label}</span>
				</a>
			{/each}
		</div>
		<div class="navbar-end">
			<div class="navbar-item">
				<div class="buttons">
					<a href="/logout" class="button">Déconnexion</a>
				</div>
			</div>
		</div>
	</div>
</nav>
<section class="section" is-main-content>
	<svelte:component this={actualMenu.component}/>
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