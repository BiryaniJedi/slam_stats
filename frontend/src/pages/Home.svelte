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
        if (!res.ok) throw new Error('Search failed. Try again.')
        return await res.json()
    }
</script>

<div class="page">
    <header>
        <h1>Slam Stats</h1>
        <p class="subtitle">MLB player analytics</p>
    </header>

    <input
        class="search-input"
        bind:value={playerName}
        oninput={debounce}
        placeholder="Search players…"
        autofocus
    />

    {#await fetchPlayers()}
        {#if searchTerm.trim()}
            <p class="status">Searching…</p>
        {/if}
    {:then players}
        {#if players.length > 0}
            <ul class="results">
                {#each players as player (player.id)}
                    <li>
                        <a href="#/player/{player.id}" class="player-card">
                            <img
                                class="thumb"
                                src="https://img.mlbstatic.com/mlb-photos/image/upload/d_people:generic:headshot:67:current.png/w_80,q_auto:best/v1/people/{player.id}/headshot/67/current"
                                alt={player.firstName + ' ' + player.lastName}
                            />
                            <div class="player-info">
                                <span class="player-name">{player.firstName} {player.lastName}</span>
                                <span class="player-meta">{player.primaryPosition.name}</span>
                            </div>
                        </a>
                    </li>
                {/each}
            </ul>
        {:else if searchTerm.trim()}
            <p class="status">No players found.</p>
        {/if}
    {:catch err}
        <p class="error">{err.message}</p>
    {/await}
</div>

<style>
    .page {
        max-width: 600px;
        margin: 0 auto;
        padding: 72px 24px 48px;
    }

    header {
        text-align: center;
        margin-bottom: 36px;
    }

    header h1 {
        font-size: 2.25rem;
        letter-spacing: -0.5px;
    }

    .subtitle {
        color: var(--text-muted);
        font-size: 0.9rem;
        margin-top: 4px;
    }

    .search-input {
        width: 100%;
        padding: 12px 16px;
        font-size: 1rem;
        background: var(--surface);
        border: 1px solid var(--border);
        border-radius: 8px;
        color: var(--text);
        outline: none;
        transition: border-color 0.15s;
        margin-bottom: 16px;
    }

    .search-input::placeholder {
        color: var(--text-muted);
    }

    .search-input:focus {
        border-color: var(--accent);
    }

    .results {
        list-style: none;
        padding: 0;
        margin: 0;
        display: flex;
        flex-direction: column;
        gap: 6px;
    }

    .player-card {
        display: flex;
        align-items: center;
        gap: 14px;
        padding: 10px 14px;
        background: var(--surface);
        border: 1px solid var(--border);
        border-radius: 8px;
        color: var(--text);
        transition: border-color 0.15s, background 0.15s;
    }

    .player-card:hover {
        border-color: var(--accent);
        background: var(--surface-2);
        color: var(--text);
    }

    .thumb {
        width: 40px;
        height: 40px;
        border-radius: 50%;
        object-fit: cover;
        flex-shrink: 0;
        background: var(--surface-2);
    }

    .player-name {
        display: block;
        font-weight: 500;
        font-size: 0.95rem;
    }

    .player-meta {
        display: block;
        font-size: 0.8rem;
        color: var(--text-muted);
        margin-top: 2px;
    }

    .status {
        color: var(--text-muted);
        text-align: center;
        padding: 32px 0;
    }

    .error {
        color: var(--error);
        text-align: center;
        padding: 32px 0;
    }
</style>
