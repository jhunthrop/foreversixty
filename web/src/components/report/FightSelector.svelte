<!-- web/src/components/report/FightSelector.svelte -->
<!-- The first control on the page. Encounters first-class, trash labelled and foldable,
     kills and wipes marked, a live fight marked as such. On phone it is a scrollable
     column of 44px rows rather than a dropdown: picking the right pull out of forty is the
     single most common thing anyone does here. -->
<script lang="ts">
  import { fightRows, fightTitle, pullNumbers } from '../../lib/report/fights';
  import { formatClock, formatDuration, outcomeLabel } from '../../lib/report/format';
  import type { FightEntry } from '../../lib/report/types';
  import { ALL_FIGHTS } from '../../lib/report/url';
  import Card from '../ui/Card.svelte';
  import Glyph from '../ui/Glyph.svelte';
  import { ROW_INTERACTIVE } from '../../lib/ui/row';

  let {
    fights,
    selected,
    onSelect,
    phaseOf = new Map<number, string>(),
    foldOnPhone = false,
  }: {
    /** On a phone, fold the list to the one row that names the fight on show (the sample report,
     *  whose table would otherwise start two screens down). */
    foldOnPhone?: boolean;
    fights: FightEntry[];
    selected: number;
    onSelect: (index: number) => void;
    /** Per fight index, the phase that pull reached; a wipe's row says which. */
    phaseOf?: ReadonlyMap<number, string>;
  } = $props();

  let showTrash = $state(false);
  /** On a phone the list is folded to the one row that names the fight on show, so the report's
   *  table starts above the second screen (W.9); a pick closes it again. From md it is always open. */
  let phoneOpen = $state(false);
  const current = $derived(fights.find((fight) => fight.index === selected) ?? null);
  const currentLabel = $derived(
    selected === ALL_FIGHTS || current === null
      ? 'All boss pulls'
      : `${fightTitle(current)} · ${outcome(current, current.kill ? '' : (phaseOf.get(current.index) ?? ''))}${current.kind === 'encounter' ? ` · ${formatDuration(current.duration_ms)}` : ''}`,
  );

  const encounters = $derived(fights.filter((fight) => fight.kind === 'encounter'));
  const shown = $derived(showTrash ? fights : encounters.length > 0 ? encounters : fights);
  const trashCount = $derived(fights.length - encounters.length);
  /** Roughly how many two-line rows the phone's 320px box shows before it scrolls. */
  const PHONE_ROWS_SHOWN = 5;
  let listAtEnd = $state(false);
  const pulls = $derived(pullNumbers(fights));
  const rows = $derived(fightRows(shown));
  /** Trash groups the reader opened, by the index of their first pull. */
  let openGroups = $state<number[]>([]);

  const outcome = outcomeLabel;

  function pick(index: number): void {
    phoneOpen = false;
    onSelect(index);
  }
</script>

{#snippet fightRow(fight: FightEntry, nested: boolean)}
  {@const pull = pulls.get(fight.index)}
  {@const isSelected = fight.index === selected}
  <!-- Two lines: the name on its own line so it is never cut to "Gener Kaal", and the
           clock, length, deaths and pull number under it. The selected row carries a gold
           edge and the raised card colour; a background shade alone was too close to the
           rest of the list to find at a glance. Kill and wipe each have their own hue. -->
  <li class={nested ? 'pl-3' : ''}>
    <button
      type="button"
      class="flex min-h-11 w-full flex-col gap-0.5 border-b border-l-[3px] px-3 py-2 text-left text-[14px] {isSelected
        ? 'border-l-gold bg-card-top'
        : 'border-line-soft hover:bg-card-top/60 border-l-transparent'}"
      aria-current={isSelected ? 'true' : undefined}
      data-testid={`fight-${fight.index}`}
      onclick={() => pick(fight.index)}
    >
      <span class="flex w-full items-baseline gap-2">
        <span
          class="min-w-0 flex-1 truncate leading-tight {fight.kind === 'encounter'
            ? 'text-strong'
            : 'text-muted'}"
          title={fightTitle(fight)}
        >
          {fightTitle(fight)}
        </span>
        <span
          class="tabular numeric shrink-0 text-[12px] font-semibold {fight.in_progress
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
      <span class="text-muted numeric flex w-full items-baseline gap-2 text-[11px]">
        <span>{formatClock(fight.start)}</span>
        <!-- A trash row's name already says its length ("Trash · 3:26"): once is enough. -->
        {#if fight.kind === 'encounter'}<span>{formatDuration(fight.duration_ms)}</span>{/if}
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
          <span class="ml-auto" data-testid={`fight-${fight.index}-pull`}>pull {pull.pull} of {pull.of}</span>
        {/if}
      </span>
    </button>
  </li>
{/snippet}

<nav aria-label="Fights" class="md:sticky md:top-4" data-testid="fight-selector">
  <Card
    glyph="list"
    title="Fights"
    meta={`${encounters.length} ${encounters.length === 1 ? 'pull' : 'pulls'}`}
    body="none"
    compact
    headClass={foldOnPhone ? 'max-md:hidden' : ''}
  >
    <!-- Folded on a phone, the toggle is the card's one row and stands in for its head: the
         head on top of it took the folded list from 46 to 95 px and pushed the sample
         report's first damage row below the first screen (onboarding-second-screen). -->
    <div class="border-line-soft flex flex-col {foldOnPhone ? 'md:border-t' : 'border-t'}">
      {#if foldOnPhone}
        <button
          type="button"
          class="flex min-h-11 w-full items-center justify-between gap-3 px-3 py-2 text-left text-[14px] md:hidden"
          aria-expanded={phoneOpen}
          data-testid="fight-selector-toggle"
          onclick={() => (phoneOpen = !phoneOpen)}
        >
          <span class="text-strong min-w-0 truncate" data-testid="fight-selector-current">{currentLabel}</span
          >
          <span class="text-muted shrink-0 text-[12px]" aria-hidden="true"
            >{phoneOpen ? 'Hide' : 'Change'}</span
          >
        </button>
      {/if}
      <div
        class="flex flex-col {phoneOpen || !foldOnPhone ? '' : 'max-md:hidden'}"
        data-testid="fight-selector-body"
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
                onclick={() => pick(ALL_FIGHTS)}
              >
                <span class="text-strong flex-1 font-semibold">All boss pulls</span>
                <span class="text-muted numeric text-[12px]">{encounters.length} pulls</span>
              </button>
            </li>
          {/if}
          {#each rows as row (row.type === 'fight' ? row.fight.index : `group-${row.key}`)}
            {#if row.type === 'fight'}
              {@render fightRow(row.fight, false)}
            {:else}
              {@const holdsSelected = row.fights.some((entry) => entry.index === selected)}
              {@const isOpen = holdsSelected || openGroups.includes(row.key)}
              <!-- Consecutive trash in one zone: one row, opened to reach each pull. -->
              <li>
                <button
                  type="button"
                  class="border-line-soft hover:bg-card-top/60 flex min-h-11 w-full flex-col gap-0.5 border-b border-l-[3px] border-l-transparent px-3 py-2 text-left text-[14px]"
                  aria-expanded={isOpen}
                  data-testid={`fight-group-${row.key}`}
                  onclick={() =>
                    (openGroups = openGroups.includes(row.key)
                      ? openGroups.filter((key) => key !== row.key)
                      : [...openGroups, row.key])}
                >
                  <span class="flex w-full items-baseline gap-2">
                    <span
                      class="text-muted min-w-0 flex-1 truncate leading-tight"
                      title={`Trash · ${row.zone}`}>Trash · {row.zone}</span
                    >
                    <span class="text-muted numeric shrink-0 text-[12px]" title="Trash: how many enemies died"
                      >{row.kills} {row.kills === 1 ? 'kill' : 'kills'}</span
                    >
                  </span>
                  <span class="text-muted numeric flex w-full items-baseline gap-2 text-[11px]">
                    <span>{formatDuration(row.duration_ms)}</span>
                    {#if row.deaths > 0}
                      <span
                        class="text-death"
                        title={`${row.deaths} ${row.deaths === 1 ? 'death' : 'deaths'}`}
                        >{row.deaths}<span aria-hidden="true">†</span></span
                      >
                    {/if}
                    <span class="ml-auto">{row.fights.length} pulls {isOpen ? '‹' : '›'}</span>
                  </span>
                </button>
              </li>
              {#if isOpen}
                {#each row.fights as fight (fight.index)}
                  {@render fightRow(fight, true)}
                {/each}
              {/if}
            {/if}
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
            class="{ROW_INTERACTIVE} border-line-soft text-strong flex min-h-11 w-full items-center gap-2 border-t px-3 text-left text-[13px] font-semibold"
            onclick={() => (showTrash = !showTrash)}
            aria-expanded={showTrash}
            data-testid="toggle-trash"
          >
            <span class="flex-1">{showTrash ? 'Hide trash' : `Show ${trashCount} trash fights`}</span>
            <Glyph
              name="chevron"
              size={16}
              class="text-muted shrink-0 {showTrash ? '-rotate-90' : 'rotate-90'}"
            />
          </button>
        {/if}
      </div>
    </div>
  </Card>
</nav>
