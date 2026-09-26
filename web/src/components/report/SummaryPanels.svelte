<!-- web/src/components/report/SummaryPanels.svelte -->
<!-- The summary as a dashboard: the things a reader wants at a glance, side by side, before
     any tab is clicked. Damage taken by ability and the deaths, each a short bar list that
     links to the tab that goes deeper. Damage and healing by source used to sit here too;
     design review 2026-09-26 finding 6 dropped them -- the Summary table two sections above
     already breaks damage and healing down by source, in the same numbers, so this repeated
     them rather than adding anything. -->
<script lang="ts">
  /** The log's null unit: a fall, a hazard, or a source the log did not name (as DeathsTab reads it). */
  const NULL_GUID = /^0+$/;
  import { splitUnitName } from '../../lib/characters';
  import {
    classColorVar,
    formatAmount,
    formatDuration,
    formatPerSecond,
    schoolToken,
  } from '../../lib/report/format';
  import type { Actor, Summary } from '../../lib/report/types';

  let {
    summary,
    everyone,
    approximate = false,
    durationMs,
    players,
    onTab,
    onSelectPlayer = undefined,
  }: {
    summary: Summary;
    durationMs: number;
    players: ReadonlySet<string>;
    /** The unscoped window: a source scope narrows the rows, never the total they share. */
    everyone?: Summary;
    /**
     * A brushed window: the by-ability split is the whole fight's, prorated by the window's
     * share, not measured from the window's events the way the Damage Taken tab measures
     * it, and a panel that reads 10,016 of a spell the window never saw is not a panel to
     * trust bare.
     */
    approximate?: boolean;
    onTab: (tab: 'damage-taken' | 'deaths') => void;
    /** Narrows the page to one player, the way the roster's names do. */
    onSelectPlayer?: (guid: string) => void;
  } = $props();

  const ROWS = 8;

  /** Every ability that hit a player, summed over the players it hit, largest first. */
  function abilityTotals(table: Actor[]): { name: string; school?: number; total: number }[] {
    // eslint-disable-next-line svelte/prefer-svelte-reactivity
    const totals = new Map<string, { name: string; school?: number; total: number }>();
    for (const actor of table) {
      if (players.size > 0 && !players.has(actor.guid)) continue;
      for (const ability of actor.abilities) {
        const key = ability.name === '' ? 'Melee' : ability.name;
        const found = totals.get(key);
        if (found === undefined)
          totals.set(key, { name: key, school: ability.school, total: ability.effective });
        else found.total += ability.effective;
      }
    }
    return [...totals.values()].sort((a, b) => b.total - a.total);
  }
  // The share is of everything every player took: a source scope narrows the rows to one
  // player's hits, never the total they share.
  const takenByAbility = $derived.by(() => {
    const rows = abilityTotals(summary.damage_taken);
    const total = abilityTotals((everyone ?? summary).damage_taken).reduce((sum, row) => sum + row.total, 0);
    return rows.slice(0, ROWS).map((row) => ({ ...row, share: total === 0 ? 0 : (row.total / total) * 100 }));
  });

  const deaths = $derived([...summary.deaths].sort((a, b) => a.at_ms - b.at_ms));
  const panel = 'border-line rounded-panel bg-raised flex flex-col gap-2 border p-3';
  const heading = 'label text-muted flex items-center justify-between';
  const more =
    'text-gold inline-flex min-h-11 items-center text-[11px] normal-case tracking-normal md:min-h-0';
</script>

<div class="grid grid-cols-1 gap-4 md:grid-cols-2" data-testid="summary-panels">
  <section class={panel} data-testid="panel-taken">
    <h2 class={heading}>
      Damage taken by ability
      <button type="button" class={more} onclick={() => onTab('damage-taken')}>Damage Taken tab</button>
    </h2>
    <ul class="flex flex-col">
      <li
        class="text-muted label hidden min-h-6 grid-cols-[minmax(150px,1.6fr)_44px_minmax(0,2fr)_64px_56px] items-center gap-x-2 md:grid"
        aria-hidden="true"
      >
        <span>Ability</span>
        <span class="text-right" title="Share of all damage taken">Share</span>
        <span class="hidden md:inline"></span>
        <span class="text-right" title="Damage taken in this window">Amount</span>
        <span class="text-right" title="Amount divided by the window's length">Per sec</span>
      </li>
      {#each takenByAbility as row (row.name)}
        <li
          class="border-line-soft grid min-h-8 grid-cols-[minmax(0,1fr)_44px_64px] items-center gap-x-2 gap-y-1 border-b py-1 text-[13px] md:grid-cols-[minmax(150px,1.6fr)_44px_minmax(0,2fr)_64px_56px] md:border-0 md:py-0"
        >
          <span class="truncate font-semibold" title={row.name}>{row.name}</span>
          <span class="text-muted tabular text-right font-mono text-[12px] whitespace-nowrap"
            >{row.share.toFixed(1)}%<span class="label font-body ml-1 md:hidden">share</span></span
          >
          <span class="bg-line-soft col-span-3 block h-[8px] w-full md:col-span-1"
            ><span class="block h-full" style={`width: ${row.share}%; background: ${schoolToken(row.school)}`}
            ></span></span
          >
          <span
            class="tabular text-right font-mono whitespace-nowrap"
            title={approximate
              ? 'Prorated inside the window from the whole fight’s split; the Damage Taken tab measures it'
              : undefined}
            >{approximate ? '~' : ''}{formatAmount(row.total)}<span class="label font-body ml-1 md:hidden"
              >amount</span
            ></span
          >
          <span class="text-muted tabular text-right font-mono text-[12px] whitespace-nowrap"
            >{formatPerSecond(row.total, durationMs)}<span class="label font-body ml-1 md:hidden"
              >per sec</span
            ></span
          >
        </li>
      {/each}
    </ul>
    <p class="text-muted text-[11px]" data-testid="panel-share-note">
      Share is of everything every player took in this window, whatever the source scope shows.
    </p>
    {#if approximate}
      <p class="text-muted text-[11px]" data-testid="panel-taken-approximate">
        Inside a window this split is the whole fight’s, prorated by the window’s share (~); the Damage Taken
        tab measures the window’s own split from the fight’s events.
      </p>
    {/if}
  </section>

  <section class={panel} data-testid="panel-deaths">
    <h2 class={heading}>
      Deaths
      <button type="button" class={more} onclick={() => onTab('deaths')}>Deaths tab</button>
    </h2>
    {#if deaths.length === 0}
      <p class="text-muted text-[13px]">Nobody died.</p>
    {:else}
      <ul class="flex flex-col">
        {#each deaths as death, i (`${death.guid}-${death.at_ms}-${i}`)}
          <li
            class="grid min-h-8 grid-cols-[52px_minmax(0,1fr)] items-center gap-x-2 gap-y-0.5 py-1 text-[13px] md:grid-cols-[52px_minmax(0,1fr)_minmax(0,1.6fr)]"
          >
            <span class="text-muted tabular font-mono text-[12px]">{formatDuration(death.at_ms)}</span>
            <span class="truncate font-semibold" style={`color: ${classColorVar(death.class)}`}
              >{#if onSelectPlayer}<button
                  type="button"
                  class="inline-flex min-h-11 items-center underline-offset-2 hover:underline md:min-h-0"
                  style={`color: ${classColorVar(death.class)}`}
                  title="Show only this player"
                  onclick={() => onSelectPlayer(death.guid)}>{splitUnitName(death.name).name}</button
                >{:else}{splitUnitName(death.name).name}{/if}</span
            >
            <!-- The killing hit wraps under the name on a phone rather than clipping its number. -->
            <span
              class="text-muted col-start-2 text-[12px] md:col-start-3"
              title="The hit that killed them, and how hard it landed"
            >
              {#if death.killing_blow}
                {NULL_GUID.test(death.killing_blow.source_guid) && death.killing_blow.spell_name === ''
                  ? 'a fall, a hazard or an untracked source'
                  : death.killing_blow.spell_name === ''
                    ? 'Melee'
                    : death.killing_blow.spell_name} ·
                {splitUnitName(death.killing_blow.source_name).name} ·
                <span class="tabular font-mono whitespace-nowrap"
                  >{formatAmount(death.killing_blow.amount)}</span
                >
              {/if}
            </span>
          </li>
        {/each}
      </ul>
    {/if}
  </section>
</div>
