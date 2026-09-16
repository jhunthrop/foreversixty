<!-- web/src/components/report/SummaryPanels.svelte -->
<!-- The summary as a dashboard: the four things a reader wants at a glance, side by side,
     before any tab is clicked. Damage and healing by source, damage taken by ability, and
     the deaths, each a short bar list that links to the tab that goes deeper. -->
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
  import ClassIcon from './ClassIcon.svelte';

  let {
    summary,
    everyone,
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
    onTab: (tab: 'damage-done' | 'healing' | 'damage-taken' | 'deaths') => void;
    /** Narrows the page to one player, the way the roster's names do. */
    onSelectPlayer?: (guid: string) => void;
  } = $props();

  const ROWS = 8;

  /** The players' rows of a table, largest first, with each row's share of every player's total. */
  function bySource(table: Actor[], all: Actor[] = table): { actor: Actor; share: number }[] {
    const ofPlayers = (actor: Actor): boolean => players.size === 0 || players.has(actor.guid);
    const rows = table.filter(ofPlayers);
    const total = all.filter(ofPlayers).reduce((sum, actor) => sum + actor.effective, 0);
    return rows
      .sort((a, b) => b.effective - a.effective)
      .slice(0, ROWS)
      .map((actor) => ({ actor, share: total === 0 ? 0 : (actor.effective / total) * 100 }));
  }

  const damage = $derived(bySource(summary.damage_done, (everyone ?? summary).damage_done));
  const healing = $derived(bySource(summary.healing, (everyone ?? summary).healing));

  /** Every ability that hit a player, summed over the players it hit, largest first. */
  const takenByAbility = $derived.by(() => {
    // eslint-disable-next-line svelte/prefer-svelte-reactivity
    const totals = new Map<string, { name: string; school?: number; total: number }>();
    for (const actor of summary.damage_taken) {
      if (players.size > 0 && !players.has(actor.guid)) continue;
      for (const ability of actor.abilities) {
        const key = ability.name === '' ? 'Melee' : ability.name;
        const found = totals.get(key);
        if (found === undefined)
          totals.set(key, { name: key, school: ability.school, total: ability.effective });
        else found.total += ability.effective;
      }
    }
    const rows = [...totals.values()].sort((a, b) => b.total - a.total);
    const total = rows.reduce((sum, row) => sum + row.total, 0);
    return rows.slice(0, ROWS).map((row) => ({ ...row, share: total === 0 ? 0 : (row.total / total) * 100 }));
  });

  const deaths = $derived([...summary.deaths].sort((a, b) => a.at_ms - b.at_ms));
  const panel = 'border-line rounded-panel bg-raised flex flex-col gap-2 border p-3';
  const heading = 'label text-muted flex items-center justify-between';
  const more =
    'text-gold inline-flex min-h-11 items-center text-[11px] normal-case tracking-normal md:min-h-0';
</script>

<div class="grid grid-cols-1 gap-4 md:grid-cols-2" data-testid="summary-panels">
  <section class={panel} data-testid="panel-damage">
    <h2 class={heading}>
      Damage done by source
      <button type="button" class={more} onclick={() => onTab('damage-done')}>Damage Done tab</button>
    </h2>
    <ul class="flex flex-col">
      <li
        class="text-muted label hidden min-h-6 grid-cols-[minmax(150px,1.6fr)_44px_minmax(0,2fr)_64px_56px] items-center gap-x-2 md:grid"
        aria-hidden="true"
      >
        <span>Name</span>
        <span
          class="text-right"
          title="Share of every player's total in this window, whatever the source scope">Share</span
        >
        <span class="hidden md:inline"></span>
        <span class="text-right" title="Amount in this window">Amount</span>
        <span class="text-right" title="Amount divided by the window's length">Per sec</span>
      </li>
      {#each damage as { actor, share } (actor.guid)}
        <li
          class="grid min-h-8 grid-cols-[minmax(0,1fr)_44px_64px] items-center gap-x-2 gap-y-1 text-[13px] md:grid-cols-[minmax(150px,1.6fr)_44px_minmax(0,2fr)_64px_56px]"
        >
          <span
            class="flex min-w-0 items-center gap-1.5 truncate font-semibold"
            style={`color: ${classColorVar(actor.class)}`}
            title={splitUnitName(actor.name).name}
            ><ClassIcon className={actor.class} size={16} />{#if onSelectPlayer}<button
                type="button"
                class="inline-flex min-h-11 items-center truncate underline-offset-2 hover:underline md:min-h-0"
                style={`color: ${classColorVar(actor.class)}`}
                title="Show only this player"
                onclick={() => onSelectPlayer(actor.guid)}>{splitUnitName(actor.name).name}</button
              >{:else}{splitUnitName(actor.name).name}{/if}</span
          >
          <span class="text-muted tabular text-right font-mono text-[12px]">{share.toFixed(1)}%</span>
          <span class="bg-line-soft col-span-3 block h-[8px] w-full md:col-span-1"
            ><span class="block h-full" style={`width: ${share}%; background: ${classColorVar(actor.class)}`}
            ></span></span
          >
          <span class="tabular text-right font-mono whitespace-nowrap"
            >{formatAmount(actor.effective)}<span class="label font-body ml-1 md:hidden">amount</span></span
          >
          <span class="text-muted tabular text-right font-mono text-[12px] whitespace-nowrap"
            >{formatPerSecond(actor.effective, durationMs)}<span class="label font-body ml-1 md:hidden"
              >per sec</span
            ></span
          >
        </li>
      {/each}
    </ul>
  </section>

  <section class={panel} data-testid="panel-healing">
    <h2 class={heading}>
      Healing done by source
      <button type="button" class={more} onclick={() => onTab('healing')}>Healing tab</button>
    </h2>
    <ul class="flex flex-col">
      <li
        class="text-muted label hidden min-h-6 grid-cols-[minmax(150px,1.6fr)_44px_minmax(0,2fr)_64px_56px] items-center gap-x-2 md:grid"
        aria-hidden="true"
      >
        <span>Name</span>
        <span
          class="text-right"
          title="Share of every player's total in this window, whatever the source scope">Share</span
        >
        <span class="hidden md:inline"></span>
        <span class="text-right" title="Amount in this window">Amount</span>
        <span class="text-right" title="Amount divided by the window's length">Per sec</span>
      </li>
      {#each healing as { actor, share } (actor.guid)}
        <li
          class="grid min-h-8 grid-cols-[minmax(0,1fr)_44px_64px] items-center gap-x-2 gap-y-1 text-[13px] md:grid-cols-[minmax(150px,1.6fr)_44px_minmax(0,2fr)_64px_56px]"
        >
          <span
            class="flex min-w-0 items-center gap-1.5 truncate font-semibold"
            style={`color: ${classColorVar(actor.class)}`}
            title={splitUnitName(actor.name).name}
            ><ClassIcon className={actor.class} size={16} />{#if onSelectPlayer}<button
                type="button"
                class="inline-flex min-h-11 items-center truncate underline-offset-2 hover:underline md:min-h-0"
                style={`color: ${classColorVar(actor.class)}`}
                title="Show only this player"
                onclick={() => onSelectPlayer(actor.guid)}>{splitUnitName(actor.name).name}</button
              >{:else}{splitUnitName(actor.name).name}{/if}</span
          >
          <span class="text-muted tabular text-right font-mono text-[12px]">{share.toFixed(1)}%</span>
          <span class="bg-line-soft col-span-3 block h-[8px] w-full md:col-span-1"
            ><span class="block h-full" style={`width: ${share}%; background: ${classColorVar(actor.class)}`}
            ></span></span
          >
          <span class="tabular text-right font-mono whitespace-nowrap"
            >{formatAmount(actor.effective)}<span class="label font-body ml-1 md:hidden">amount</span></span
          >
          <span class="text-muted tabular text-right font-mono text-[12px] whitespace-nowrap"
            >{formatPerSecond(actor.effective, durationMs)}<span class="label font-body ml-1 md:hidden"
              >per sec</span
            ></span
          >
        </li>
      {/each}
    </ul>
  </section>

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
          class="grid min-h-8 grid-cols-[minmax(0,1fr)_44px_64px] items-center gap-x-2 gap-y-1 text-[13px] md:grid-cols-[minmax(150px,1.6fr)_44px_minmax(0,2fr)_64px_56px]"
        >
          <span class="truncate font-semibold" title={row.name}>{row.name}</span>
          <span class="text-muted tabular text-right font-mono text-[12px]">{row.share.toFixed(1)}%</span>
          <span class="bg-line-soft col-span-3 block h-[8px] w-full md:col-span-1"
            ><span class="block h-full" style={`width: ${row.share}%; background: ${schoolToken(row.school)}`}
            ></span></span
          >
          <span class="tabular text-right font-mono">{formatAmount(row.total)}</span>
          <span class="text-muted tabular text-right font-mono text-[12px]"
            >{formatPerSecond(row.total, durationMs)}<span class="label font-body ml-1 md:hidden"
              >per sec</span
            ></span
          >
        </li>
      {/each}
    </ul>
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
