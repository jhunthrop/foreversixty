<!-- web/src/components/guild/RaidClassRow.svelte -->
<!-- The band's class row (raid comp spec 3.1 line 6): every class with someone In or Late, in the
     game's order, its crest, name and that count. Not interactive and not a tab stop: each item's
     specs are its `title` for a pointer and its accessible name for touch and screen readers. Below
     768 the class name drops, crest and count stay. Every class once the raid is cancelled or under
     way reads muted with the crest at 45%. Absent with nobody In or Late. -->
<script lang="ts">
  import { compCopy } from '../../lib/guild/raid-copy';
  import type { ClassRowItem } from '../../lib/guild/raid-comp';
  import { classColorVar } from '../../lib/report/format';
  import ClassCrestRing from '../character/ClassCrestRing.svelte';

  const CREST_PX = 20;

  let {
    items,
    quiet,
  }: {
    items: readonly ClassRowItem[];
    /** Cancelled or started: every count muted, every crest at 45%. */
    quiet: boolean;
  } = $props();
</script>

{#if items.length > 0}
  <ul
    class="flex flex-wrap gap-x-2 gap-y-1.5"
    aria-label={compCopy.classRowLabel}
    data-testid="raid-class-row"
  >
    {#each items as item (item.classSlug)}
      {@const muted = quiet}
      <li
        class="flex h-6 items-center"
        title={item.specs === '' ? undefined : item.specs}
        data-testid={`raid-class-${item.classSlug}`}
      >
        <span class="sr-only">{item.accessibleName}</span>
        <span class="flex items-center" aria-hidden="true">
          <span class="flex shrink-0 {muted ? 'opacity-45' : ''}"
            ><ClassCrestRing characterClass={item.classSlug} size={CREST_PX} /></span
          >
          <span
            class="ml-1.5 hidden text-[13px] font-semibold md:inline {muted ? 'text-muted' : ''}"
            style={muted ? undefined : `color:${classColorVar(item.classSlug)}`}>{item.name}</span
          >
          <span class="numeric ml-1 text-[13px] font-bold {muted ? 'text-muted' : 'text-strong'}"
            >{item.count}</span
          >
        </span>
      </li>
    {/each}
  </ul>
{/if}
