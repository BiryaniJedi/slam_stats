<script>
    import Router from 'svelte-spa-router'
    import Home from './pages/Home.svelte'
    import Player from './pages/Player.svelte'

    const routes = {
        '/': Home,
        '/player/:id': Player,
    }
</script>

<Router {routes} />
