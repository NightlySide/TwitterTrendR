<script>
	export let name;

	const fetchTrends = async() => {
		const results = await fetch("https://twittertrendr.herokuapp.com/trends");
		return results.json();
	}
</script>

<main>
	<h1>Hello!</h1>
	<p>Visit the <a href="https://svelte.dev/tutorial">Svelte tutorial</a> to learn how to build Svelte apps.</p>
	{#await fetchTrends()}
		<p>Loading...</p>
	{:then trends}
		{#each trends as trend}
		<li>{trend.name}</li>
		{/each}
	{:catch error}
		{error}
	{/await}
</main>

<style>
	main {
		text-align: center;
		padding: 1em;
		max-width: 240px;
		margin: 0 auto;
	}

	h1 {
		color: #ff3e00;
		text-transform: uppercase;
		font-size: 4em;
		font-weight: 100;
	}

	@media (min-width: 640px) {
		main {
			max-width: none;
		}
	}
</style>