<script lang="ts">
	const fetchTrends: Promise<any[]> = (async() => {
		const results = await fetch("https://twittertrendr.herokuapp.com/trends");
		return results.json();
	})();
</script>

<main>
	<h1>Hello!</h1>
	<p>Visit the <a href="https://svelte.dev/tutorial">Svelte tutorial</a> to learn how to build Svelte apps.</p>
	{#await fetchTrends}
		<p>Loading...</p>
	{:then trends}
		{#each trends as trend}
		<li>{trend.name}</li>
		{/each}
	{:catch error}
		{error}
	{/await}
</main>