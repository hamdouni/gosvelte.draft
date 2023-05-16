<script>
export let message = "une erreur inconnue s'est produite";
</script>

<article class="alert">{message}</article>

<style>
.alert {
    max-width: 15rem;
    color: var(--text-alert-color);
    font-weight: bold;
}
</style>
