<script lang="ts">
    import DarkmodeToggle from "$lib/components/darkmodeToggle.svelte";
    import * as Card from "$lib/components/ui/card";
    import { ModeWatcher } from "mode-watcher";

    let menuPromise = getTodaysMenu();
    async function getTodaysMenu() {
        const dbUrl = import.meta.env.VITE_DB_URL;
        const endpoint = dbUrl + "getTodaysMenu";
        let response = await fetch(endpoint);

        console.log(response);
        let data = await response.json();
        console.log(data);
        return data;
    }
</script>

<ModeWatcher />

<main class="w-screen h-screen flex justify-center items-center">
    <div class="fixed top-10 right-10"><DarkmodeToggle /></div>
    {#await menuPromise}
        Loading Menu
    {:then menu}
        <Card.Root class="w-fit ">
            <Card.Header>
                <Card.Title>Today's Menu</Card.Title>
            </Card.Header>
            <Card.Content>
                <p>{menu.Breakfast}</p>
                <p>{menu.Lunch}</p>
                <p>{menu.Dinner}</p>
            </Card.Content>
        </Card.Root>
    {:catch someError}
        System error: {someError.message}.
    {/await}
</main>
