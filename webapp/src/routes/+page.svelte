<script lang="ts">
  import type { PageData } from "./$types";

    export let data: PageData;
</script>
<h1>From server: {data.msg}</h1>
