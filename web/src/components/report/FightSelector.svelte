<!-- web/src/components/report/FightSelector.svelte -->
<!-- The first control on the page. Encounters first-class, trash labelled and foldable,
     kills and wipes marked, a live fight marked as such. On phone it is a scrollable
     column of 44px rows rather than a dropdown: picking the right pull out of forty is the
     single most common thing anyone does here. -->
<script lang="ts">
  import { pullNumbers } from '../../lib/report/fights';
  import { formatClock, formatDuration, outcomeLabel } from '../../lib/report/format';
  import type { FightEntry } from '../../lib/report/types';
  import { ALL_FIGHTS } from '../../lib/report/url';

  let {
    fights,
    selected,
    onSelect,
    phaseOf = new Map<number, string>(),
  }: {
    fights: FightEntry[];
    selected: number;
    onSelect: (index: number) => void;
    /** Per fight index, the phase that pull reached; a wipe's row says which. */
    phaseOf?: ReadonlyMap<number, string>;
  } = $props();

  let showTrash = $state(false);

  const encounters = $derived(fights.filter((fight) => fight.kind === 'encounter'));
  const shown = $derived(showTrash ? fights : encounters.length > 0 ? encounters : fights);
  const trashCount = $derived(fights.length - encounters.length);
  /** Roughly how many two-line rows the phone's 320px box shows before it scrolls. */
  const PHONE_ROWS_SHOWN = 5;
  let listAtEnd = $state(false);
  const pulls = $derived(pullNumbers(fights));

  const outcome = outcomeLabel;
</script>

<nav
  aria-label="Fights"
  class="border-line rounded-panel bg-raised flex flex-col border"
  data-testid="fight-selector"
>
  <ul
    class="flex max-h-[320px] flex-col overflow-y-auto md:max-h-[calc(100vh-180px)]"
    onscroll={(event) => {
      const list = event.currentTarget;
      listAtEnd = list.scrollTop + list.clientHeight >= list.scrollHeight - 4;
    }}
  >
    {#if encounters.length > 0}
      <li>
        <button
          type="button"
          class="flex min-h-11 w-full items-center gap-3 border-b border-l-[3px] px-3 py-2 text-left text-[14px] {selected ===
          ALL_FIGHTS
            ? 'border-l-gold bg-card-top'
            : 'border-line-soft hover:bg-card-top/60 border-l-transparent'}"
          aria-current={selected === ALL_FIGHTS ? 'true' : undefined}
          data-testid="fight-all"
          onclick={() => onSelect(ALL_FIGHTS)}
        >
          <span class="text-strong flex-1 font-semibold">All boss pulls</span>
          <span class="text-muted tabular font-mono text-[12px]">{encounters.length} pulls</span>
        </button>
      </li>
    {/if}
    {#each shown as fight (fight.index)}
      {@const pull = pulls.get(fight.index)}
      {@const isSelected = fight.index === selected}
      <!-- Two lines: the name on its own line so it is never cut to "Gener Kaal", and the
           clock, length, deaths and pull number under it. The selected row carries a gold
           edge and the raised card colour; a background shade alone was too close to the
           rest of the list to find at a glance. Kill and wipe each have their own hue. -->
      <li>
        <button
          type="button"
          class="flex min-h-11 w-full flex-col gap-0.5 border-b border-l-[3px] px-3 py-2 text-left text-[14px] {isSelected
            ? 'border-l-gold bg-card-top'
            : 'border-line-soft hover:bg-card-top/60 border-l-transparent'}"
          aria-current={isSelected ? 'true' : undefined}
          data-testid={`fight-${fight.index}`}
          onclick={() => onSelect(fight.index)}
        >
          <span class="flex w-full items-baseline gap-2">
            <span
              class="min-w-0 flex-1 truncate leading-tight {fight.kind === 'encounter'
                ? 'text-strong'
                : 'text-muted'}"
              title={fight.name}
            >
              {fight.name}
            </span>
            <span
              class="tabular shrink-0 font-mono text-[12px] font-semibold {fight.in_progress
                ? 'text-gold'
                : fight.kind !== 'encounter'
                  ? 'text-muted'
                  : fight.kill
                    ? 'text-kill'
                    : 'text-wipe'}"
              title={fight.kind !== 'encounter'
                ? 'Trash: how many enemies died'
                : fight.kill
                  ? 'The boss died'
                  : 'The percentage is the boss’s health when the pull ended'}
              data-testid={`fight-${fight.index}-outcome`}
            >
              {outcome(fight, fight.kill ? '' : (phaseOf.get(fight.index) ?? ''))}
            </span>
          </span>
          <span class="text-muted tabular flex w-full items-baseline gap-2 font-mono text-[11px]">
            <span>{formatClock(fight.start)}</span>
            <span>{formatDuration(fight.duration_ms)}</span>
            {#if fight.deaths > 0}
              <span
                class="text-death"
                title={`${fight.deaths} ${fight.deaths === 1 ? 'death' : 'deaths'}`}
                data-testid={`fight-${fight.index}-deaths`}
              >
                {fight.deaths}<span aria-hidden="true">†</span>
              </span>
            {/if}
            {#if pull !== undefined && pull.of > 1}
              <span class="ml-auto" data-testid={`fight-${fight.index}-pull`}
                >pull {pull.pull} of {pull.of}</span
              >
            {/if}
          </span>
        </button>
      </li>
    {/each}
  </ul>
  {#if encounters.length > PHONE_ROWS_SHOWN && !listAtEnd}
    <!-- The phone list is a scroll box that ends on a whole row, so it looks complete when it
         is not: this says how much is below the fold until the reader gets there. -->
    <span
      class="text-muted border-line-soft border-t px-3 py-1 text-[12px] md:hidden"
      data-testid="fights-more">Scroll the list for all {encounters.length} pulls</span
    >
  {/if}
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
