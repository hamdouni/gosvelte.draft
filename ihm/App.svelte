<pre class="toto">
    Bonjour depuis Svelte.
</pre>
<hr>
<input bind:value="{nom}" type="text" placeholder="Entrer votre prénom...">
<button on:click="{call}">
    Cliquez ici pour lancer un appel à l'api
</button>
<hr>
<h1>{resultat}</h1>

<script>
    let nom = "";
    let resultat = "";
    function call() {
        fetch("/hello?nom="+nom).then(function(response){
           return(response.json());
        }).then(function(data){
            resultat = data;
        })
    }
</script>

<style>
    pre {
        background-color: #eee;
    }
</style>
