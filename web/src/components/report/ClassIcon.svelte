<!-- web/src/components/report/ClassIcon.svelte -->
<!-- A class's circular ringed HD crest beside a name (design/DESIGN-SYSTEM.md tenet 7, one
     crest language); a class this client has no crest for (a log from another game version)
     shows a class-coloured dot instead of nothing. -->
<script lang="ts">
  import { classSlugOf } from '../../lib/report/planner-link';
  import { classColorVar } from '../../lib/report/format';
  import ClassCrestRing from '../character/ClassCrestRing.svelte';

  let { className, size = 18 }: { className: string | undefined; size?: number } = $props();

  const slug = $derived(classSlugOf(className));
</script>

{#if slug !== null}
  <ClassCrestRing characterClass={slug} {size} />
{:else}
  <span
    class="inline-block shrink-0 rounded-full align-middle"
    style={`width: ${Math.round(size / 2)}px; height: ${Math.round(size / 2)}px; background: ${classColorVar(className)}`}
    aria-hidden="true"
  ></span>
{/if}
