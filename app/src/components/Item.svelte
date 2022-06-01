<script>
    import Card from "./Card.svelte"
    import Modal from "./Modal.svelte"
    import {closeModal, Modals, openModal} from "svelte-modals"

    export let url
    let showCard = true

    async function update(data) {
        const json = {
            redirect: data.redirect,
            short: data.short,
            random: data.random,
            id: url.id
        }

        await fetch("http://localhost:3000/url", {
            method: "PATCH",
            headers: {"Content-Type": "application/json"},
            body: JSON.stringify(json)
        }).then(res => {
            console.log(res)
        })
    }

    function handleOpen(url) {
        openModal(Modal, {
            title: "Update Url Link",
            send: update,
            short: url.short,
            redirect: url.redirect,
            random: url.random
        })
    }

    async function deleteUrl() {
        if (confirm("Are you sure you wish to delete this url link (" + url.short + ")?")) {
            await fetch("http://localhost:3000/url/" + url.id, {
                method: "DELETE"
            }).then(response => {
                showCard = false
                console.log(response)
            })
        }
    }

</script>
{#if showCard }
    <Card>
        <p>ID: {url.id}</p>
        <p>ShortUrl: http://localhost:3000/r/{url.short} </p>
        <p>Redirect: {url.redirect}</p>
        <p>Clicked: {url.clicked}</p>
        <button class="update" on:click={ handleOpen(url) }>Update</button>
        <button class="delete" on:click={deleteUrl}>Delete</button>
    </Card>
{/if}
<Modals>
    <div class="backdrop"
         on:click={closeModal}
         slot="backdrop"></div>
</Modals>

<style>
    button {
        color: white;
        font-weight: bolder;
        border: none;
        padding: .60rem;
        border-radius: 5px;
    }

    .update {
        background-color: cornflowerblue;
    }

    .delete {
        background-color: red;
    }

    .backdrop {
        position: fixed;
        top: 0;
        bottom: 0;
        right: 0;
        left: 0;
        background: rgb(255, 255, 255)
    }
</style>