<!-- web/src/components/report/FightSelector.svelte -->
<!-- The first control on the page. Encounters first-class, trash labelled and foldable,
     kills and wipes marked, a live fight marked as such. On phone it is a scrollable
     column of 44px rows rather than a dropdown: picking the right pull out of forty is the
     single most common thing anyone does here. -->
<script lang="ts">
  import { formatClock, formatDuration } from '../../lib/report/format';
  import type { FightEntry } from '../../lib/report/types';

  let {
    fights,
    selected,
    onSelect,
  }: { fights: FightEntry[]; selected: number; onSelect: (index: number) => void } = $props();

  let showTrash = $state(false);

  const encounters = $derived(fights.filter((fight) => fight.kind === 'encounter'));
  const shown = $derived(showTrash ? fights : encounters.length > 0 ? encounters : fights);
  const trashCount = $derived(fights.length - encounters.length);

  function outcome(fight: FightEntry): string {
    if (fight.in_progress) return 'Live';
    if (fight.kind !== 'encounter') return `${fight.npc_kills} killed`;
    return fight.kill ? 'Kill' : 'Wipe';
  }
</script>

<nav
  aria-label="Fights"
  class="border-line rounded-panel bg-raised flex flex-col border"
  data-testid="fight-selector"
>
  <ul class="flex max-h-[320px] flex-col overflow-y-auto md:max-h-[560px]">
    {#each shown as fight (fight.index)}
      <li>
        <button
          type="button"
          class="border-line-soft flex min-h-11 w-full items-center gap-3 border-b px-3 py-2 text-left text-[14px]"
          class:bg-card-top={fight.index === selected}
          aria-current={fight.index === selected ? 'true' : undefined}
          data-testid={`fight-${fight.index}`}
          onclick={() => onSelect(fight.index)}
        >
          <span class="text-muted tabular w-[68px] shrink-0 font-mono text-[12px]"
            >{formatClock(fight.start)}</span
          >
          <span class="min-w-0 flex-1 truncate {fight.kind === 'encounter' ? 'text-strong' : 'text-muted'}">
            {fight.name}
          </span>
          <span
            class="tabular shrink-0 font-mono text-[12px]"
            class:text-gold={fight.in_progress}
            data-testid={`fight-${fight.index}-outcome`}
          >
            {outcome(fight)}
          </span>
          <span class="text-muted tabular w-[56px] shrink-0 text-right font-mono text-[12px]">
            {formatDuration(fight.duration_ms)}
          </span>
        </button>
      </li>
    {/each}
  </ul>
  {#if trashCount > 0 && encounters.length > 0}
    <button
      type="button"
      class="text-muted inline-flex min-h-11 items-center px-3 text-[12px] font-bold tracking-[0.06em] uppercase"
      onclick={() => (showTrash = !showTrash)}
      data-testid="toggle-trash"
    >
      {showTrash ? 'Hide trash' : `Show ${trashCount} trash fights`}
    </button>
  {/if}
</nav>
