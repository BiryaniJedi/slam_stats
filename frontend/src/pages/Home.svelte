<script>
    import {sanitizeName} from '../lib/utils'
    let playerName = $state('')
    let debounceTimer
    let searchTerm = $state('')

    const debounce = () => {
        clearTimeout(debounceTimer)
        debounceTimer = setTimeout(() => {
            searchTerm = playerName
        }, 300)
    }

    const fetchPlayers = async () => {
        if (!searchTerm.trim()) return []
        const res = await fetch(`/api/players/${sanitizeName(searchTerm)}`)
        return await res.json().catch(() => [])
    }
</script>
<input bind:value={playerName} oninput={debounce} placeholder="Search players">
<h2>Results:</h2>
{#await fetchPlayers()}
    <p>Loading...</p>
{:then players}
    {#each players as player (player.id)}
        <div>
            <a href="#/player/{player.id}">
                <h2>{`${player.firstName} ${player.lastName}`}</h2>
            </a>
        </div>
    {/each}
{/await}
