<script>
    import Item from "./Item.svelte"
    import {onMount} from "svelte"

    let urls = []

    onMount(async () => {
        const res = await fetch("http://localhost:3000/url")
        urls = await res.json()
    })
</script>

{#each urls as url}
    <Item url={url}/>
{/each}
