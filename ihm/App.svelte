<script>
    let nom = "";
    let resultat = "";
    let historique = [];
    function callApi(endpoint) {
        fetch("/"+endpoint+"?nom="+nom).then(function(response){
           return(response.json());
        }).then(function(data){
            resultat = data;
        });
        reloadHistoric();
    }
    function callBonjour() {
        callApi("hello");
    }
    function callMaj() {
        callApi("upper");
    }
    function callMin() {
        callApi("lower");
    }
    function reloadHistoric() {
        fetch("/historic").then(function(response){
           return(response.json());
        }).then(function(data){
            if(data != null) historique = data;
        })
    }
    reloadHistoric();
</script>

<pre class="toto">
    Bonjour depuis Svelte.
</pre>
<hr>
<input bind:value="{nom}" type="text" placeholder="Entrer votre prénom...">
<button on:click="{callBonjour}">
    Bonjour
</button>
<button on:click="{callMaj}">
    Majuscule
</button>
<button on:click="{callMin}">
    Minuscule
</button>
<hr>
<h1>{resultat}</h1>
<hr>
<ul>
    {#each historique as evenement}
        <li>{evenement}</li>
    {/each}
</ul>

<style>
    pre {
        background-color: #eee;
    }
</style>
