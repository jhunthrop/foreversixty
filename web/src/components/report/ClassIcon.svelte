<!-- web/src/components/report/ClassIcon.svelte -->
<!-- A class's circular ringed HD crest beside a name (design/DESIGN-SYSTEM.md tenet 7, one
     crest language); a unit the log could not place in a class shows the empty ring at the same
     size, never a dot or another class's crest. -->
<script lang="ts">
  import { classSlugOf } from '../../lib/report/planner-link';
  import Crest from '../ui/Crest.svelte';

  let { className, size = 18 }: { className: string | undefined; size?: number } = $props();

  const slug = $derived(classSlugOf(className));
</script>

{#if slug !== null}
  <Crest cls={slug} {size} />
{:else}
  <Crest cls="" {size} />
{/if}
