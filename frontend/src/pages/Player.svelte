<script>
    export let params

    const fetchPlayer = async (id) => {
        const res = await fetch(`/api/player/${id}`)
        if (!res.ok) throw new Error('Player not found.')
        return await res.json()
    }
</script>

<div class="page">
    <a href="/" class="back">← Search</a>

    {#await fetchPlayer(params.id)}
        <p class="status">Loading…</p>
    {:then player}
        <div class="profile">
            <img
                class="headshot"
                src="https://img.mlbstatic.com/mlb-photos/image/upload/d_people:generic:headshot:67:current.png/w_213,q_auto:best/v1/people/{player.id}/headshot/67/current"
                alt={player.firstName + ' ' + player.lastName}
            />
            <div class="info">
                <h1 class="name">{player.firstName} {player.lastName}</h1>
                <div class="tags">
                    <span class="tag">{player.primaryPosition.name}</span>
                    <span class="tag">{player.currentAge} yrs</span>
                    <span class="tag {player.active ? 'active' : 'inactive'}">
                        {player.active ? 'Active' : 'Inactive'}
                    </span>
                </div>
            </div>
        </div>
    {:catch err}
        <p class="error">{err.message}</p>
    {/await}
</div>

<style>
    .page {
        max-width: 700px;
        margin: 0 auto;
        padding: 40px 24px 48px;
    }

    .back {
        display: inline-block;
        font-size: 0.875rem;
        color: var(--text-muted);
        margin-bottom: 36px;
        transition: color 0.15s;
    }

    .back:hover {
        color: var(--text);
    }

    .profile {
        display: flex;
        align-items: flex-start;
        gap: 32px;
    }

    .headshot {
        width: 120px;
        height: 120px;
        border-radius: 50%;
        object-fit: cover;
        background: var(--surface-2);
        border: 2px solid var(--border);
        flex-shrink: 0;
    }

    .info {
        padding-top: 8px;
    }

    .name {
        font-size: 2rem;
        letter-spacing: -0.5px;
        margin-bottom: 14px;
    }

    .tags {
        display: flex;
        flex-wrap: wrap;
        gap: 8px;
    }

    .tag {
        font-size: 0.8rem;
        padding: 4px 10px;
        border-radius: 20px;
        background: var(--surface);
        border: 1px solid var(--border);
        color: var(--text-muted);
    }

    .tag.active {
        color: #3fb950;
        border-color: rgba(63, 185, 80, 0.4);
        background: rgba(63, 185, 80, 0.1);
    }

    .tag.inactive {
        color: var(--error);
        border-color: rgba(248, 81, 73, 0.4);
        background: rgba(248, 81, 73, 0.1);
    }

    .status {
        color: var(--text-muted);
        padding: 48px 0;
        text-align: center;
    }

    .error {
        color: var(--error);
        padding: 48px 0;
        text-align: center;
    }
</style>
