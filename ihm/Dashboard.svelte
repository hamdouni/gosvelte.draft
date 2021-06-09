<nav class="navbar is-fixed-top is-wrapped has-shadow is-light">
	<div class="navbar-brand">
		<div class="navbar-item is-size-5">
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
			<a href="#bonjour" class="navbar-item" on:click="{activateBonjour}">
				<span class="icon is-medium"> <i class="fas fa-home"></i> </span>
				<span>Bonjour</span>
			</a>
			<a href="#maj" class="navbar-item" on:click="{activateMaj}">
				<span class="icon is-medium"> <i class="fas fa-chart-bar"></i> </span>
				<span>Majuscule</span>
			</a>
			<a href="#min" class="navbar-item" on:click="{activateMin}">
				<span class="icon is-medium"> <i class="fas fa-book"></i> </span>
				<span>Minuscule</span>
			</a>
			<a href="#historic" class="navbar-item" on:click="{activateHistoric}">
				<span class="icon is-medium"> <i class="fas fa-shopping-cart"></i> </span>
				<span>Historique</span>
			</a>
			
			<div class="navbar-item has-dropdown is-hoverable">
				<a href="#aides" class="navbar-link">
					<span class="icon is-medium"> <i class="fas fa-question"></i> </span>
					<span>Aides</span>
				</a>
				<div class="navbar-dropdown">
					<a href="#documentation" class="navbar-item">
						<span class="icon is-medium"> <i class="fas fa-book"></i> </span>
						<span>Documentation</span>
					</a>
					<a href="#faq" class="navbar-item">
						<span class="icon is-medium"> <i class="fas fa-sitemap"></i> </span>
						<span>F.A.Q</span>
					</a>
					<hr class="navbar-divider">
					<div class="navbar-item">
						Version 0.1.0
					</div>
				</div>
			</div>
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

<main class="content is-wrapped">
	{#if actualMenu == 'Historic'}
	<Historic />
	{/if}
	{#if actualMenu == 'Maj'}
	<Maj />
	{/if}
	{#if actualMenu == 'Min'}
	<Min />
	{/if}
	{#if actualMenu == 'Bonjour'}
	<Bonjour />
	{/if}
</main>

<script>
	import Historic from './functions/Historic.svelte';
	import Maj from './functions/Maj.svelte';
	import Min from './functions/Min.svelte';
	import Bonjour from './functions/Bonjour.svelte';

	let showNavbarMenu = false;
	let actualMenu = 'Bonjour';

	function activateBonjour() {
		actualMenu = 'Bonjour';
	}
	function activateMaj() {
		actualMenu = 'Maj';
	}
	function activateMin() {
		actualMenu = 'Min';
	}
	function activateHistoric() {
		actualMenu = 'Historic';
	}
</script>