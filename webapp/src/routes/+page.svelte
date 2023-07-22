<script lang="ts">
  import type { Drink } from "$lib/types";
  import type { PageData } from "./$types";
  import autoAnimate from "@formkit/auto-animate"

  export let data: PageData;

  const handleAddDrink = async () => {
    const drink : Drink = await fetch(`/drinks`, {method: "POST"}).then(res => res.json());
    data.drinks = [...data.drinks, drink];
  };
</script>

<h1>🍻 Bar counter</h1>
<div class="row flex-spaces child-borders">
  <button on:click={handleAddDrink}>Add random drink</button>
</div>
<ul use:autoAnimate={{ duration: 150 }}>
  {#each data.drinks as drink}
    <li>
      {drink.Emoji} {drink.Name}
    </li>
  {/each}
</ul>
