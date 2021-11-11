<script>
    import { onMount } from "svelte";

    let loading = false;
    let users = [];

    const titleize = (s) => s.replace(/^([a-z])/, (_, r) => r.toUpperCase());

    async function update() {
        loading = true;

        users = [];
        users = (
            await (await fetch("https://randomuser.me/api/?results=10")).json()
        ).results;

        loading = false;
    }

    onMount(() => update());
</script>

<div>
    <button class="button is-primary" on:click={update}>Update</button>
    <button class="button is-primary" on:click={() => (users = [])}
        >No Data</button
    >
</div>

<br />
<br />

<table class="table is-fullwidth">
    <thead>
        <tr>
            <th />
            <th>First Name</th>
            <th>Last Name</th>
            <th>City</th>
            <th>State</th>
        </tr>
    </thead>
    <tbody>
        {#each users as user}
            <tr>
                <td>
                    <figure class="image">
                        <img
                            class="is-rounded"
                            src={user.picture.thumbnail}
                            alt=""
                        />
                    </figure>
                </td>
                <td>{titleize(user.name.first)}</td>
                <td>{titleize(user.name.last)}</td>
                <td>{titleize(user.location.city)}</td>
                <td>{titleize(user.location.state)}</td>
            </tr>
        {:else}
            {#if !loading}
                <tr>
                    <td colspan="5">
                        <section class="section">
                            <div class="content has-text-grey has-text-centered">
                                <p><i class="far fa-3x fa-frown" /></p>
                                <p>No data</p>
                            </div>
                        </section>
                    </td>
                </tr>
            {/if}
        {/each}
    </tbody>
</table>
