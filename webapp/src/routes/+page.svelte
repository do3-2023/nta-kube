<script lang="ts">
  import type { PageData } from "./$types";
  import type { Drink } from '$lib/types';
  import autoAnimate from "@formkit/auto-animate"
  export let data: PageData;

  const addDrink = async (): Promise<void> => {
      let response: Response = await fetch(`http://localhost:3000/drinks`, {
          method: 'POST'
      });
      if (response.status !== 201) {
          throw new Error('Failed to add drink');
      };

      const drink: Drink = await response.json();
      data.drinks = [...data.drinks, drink];
  };

</script>

<h1>🍻 Bar counter</h1>
<div class="row flex-spaces child-borders">
  <button on:click={addDrink}>Add random drink</button>
</div>
<ul use:autoAnimate={{ duration: 150 }}>
  {#each data.drinks as drink}
    <li>
      {drink.Emoji} {drink.Name}
    </li>
  {/each}
</ul>
