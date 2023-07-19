<script lang="ts">
  import type { PageData } from "./$types";
  import type { Drink } from '$lib/types';
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
<button style="margin:auto;display:flex" on:click={addDrink}>Add random drink</button>
<ul>
  {#each data.drinks as drink}
    <li>
      {drink.Emoji} {drink.Name}
    </li>
  {/each}
</ul>
