<script>
    export let params
    
    const fetchPlayer = async (id) => {
        const res = await fetch(`/api/player/${id}`)
        const data = await res.json()
            .then((val) => {return val})
            .catch((err) => {return err})
        console.log(data)
        return data
    }
</script>

{#await fetchPlayer(params.id)}
    <p>Loading...</p>
{:then player}
    <h1>{player.firstName + ' ' + player.lastName}</h1>
    <img 
        src="https://img.mlbstatic.com/mlb-photos/image/upload/d_people:generic:headshot:67:current.png/w_213,q_auto:best/v1/people/{player.id}/headshot/67/current"
        alt={player.firstName + ' ' + player.lastName}
    />
    <h3>{player.currentAge} yrs</h3>
    <h3>Position: {player.primaryPosition.name}</h3>
    <h3>Active: {player.active}</h3>
{:catch err}
    <p>Error: {err}</p>
{/await}
