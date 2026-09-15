<!-- web/src/components/report/ActorRow.svelte -->
<!-- One table row, exactly the shape spec section 4 lists: parse percentile, class-coloured
     name, amount bar segmented by ability, active time, per-second figure, expander. On
     phone the row becomes a card: the same fields, stacked, with the bar full width.

     Active time: `actor.active_ms` is exact even under a brushed window (window.ts measures
     it from the one-second series rather than scaling it), but that series is one-second
     grained, so a percentage of a short window quantises hard -- a 3s window can only ever
     read 0%, 33%, 67% or 100%. Below ACTIVITY_SECONDS_BELOW_MS this column reads seconds
     instead of a percentage: "2s active" is honest at that grain, "67%" implies precision
     the data does not have.

     The Amount column is `actor.effective`, which window.ts computes independently from
     the one-second series -- exact regardless of the window -- and which filters.ts leaves
     alone unless a target or boss filter is active (ReportView.svelte's `approximate`
     already folds that in). It never carries the `~` mark. The expanded Abilities and
     Targets tables below the row DO carry it: `ability.total`/`effective` and
     `target.total` are the per-ability and per-target splits that both the window and a
     target/boss filter scale by a ratio. -->
<script lang="ts">
  import { characterHref, splitUnitName } from '../../lib/characters';
  import {
    approximateAriaLabel,
    approximateMark,
    approximateTitle,
    classColorVar,
    formatAmount,
    formatPercent,
    formatPerSecond,
    parseTitle,
    percentileToken,
    schoolName,
  } from '../../lib/report/format';
  import type { Placement } from '../../lib/report/percentile';
  import type { Actor } from '../../lib/report/types';
  import AbilityBar from './AbilityBar.svelte';

  let {
    rank,
    actor,
    peak,
    durationMs,
    percentile = null,
    parseFallback = '',
    approximate = false,
    characterLink = null,
  }: {
    rank: number;
    actor: Actor;
    peak: number;
    durationMs: number;
    percentile?: Placement | null;
    parseFallback?: string;
    approximate?: boolean;
    characterLink?: { region: string; ruleset: string } | null;
  } = $props();

  /** Below ten one-second buckets, a percentage rounds to steps of 10% or coarser. */
  const ACTIVITY_SECONDS_BELOW_MS = 10_000;

  let open = $state(false);
  const color = $derived(classColorVar(actor.class));
  const display = $derived(splitUnitName(actor.name));
  const activitySeconds = $derived(Math.round(actor.active_ms / 1000));
  const activityPct = $derived(durationMs === 0 ? 0 : (actor.active_ms / durationMs) * 100);
  const showActivitySeconds = $derived(durationMs > 0 && durationMs < ACTIVITY_SECONDS_BELOW_MS);
  /** One source for the figure, which the desktop column and the phone card both read. */
  const activeText = $derived(showActivitySeconds ? `${activitySeconds}s` : formatPercent(activityPct));
  const mark = $derived(approximateMark(approximate));
  const title = $derived(approximateTitle(approximate));
  /** Overhealing as a share of the raw total, for healing rows; null where there is none. */
  const overhealPct = $derived(
    actor.overheal === undefined || actor.total <= 0 ? null : (actor.overheal / actor.total) * 100,
  );
  /**
   * Targets merged by name: a trash pack is six "Gluttonous Tick" GUIDs, and six rows of
   * the same name tell nobody anything the one row with a count does not.
   */
  /** Ability names two different spell ids share, shown with the id to tell them apart. */
  const sameName = $derived.by(() => {
    // eslint-disable-next-line svelte/prefer-svelte-reactivity
    const seen = new Map<string, number>();
    for (const ability of actor.abilities) seen.set(ability.name, (seen.get(ability.name) ?? 0) + 1);
    return new Set([...seen.entries()].filter(([, count]) => count > 1).map(([name]) => name));
  });

  const targetsByName = $derived.by(() => {
    // A plain Map: built once inside the derived and never read reactively by key.
    // eslint-disable-next-line svelte/prefer-svelte-reactivity
    const merged = new Map<string, { name: string; total: number; count: number }>();
    for (const target of actor.targets) {
      const name = splitUnitName(target.name).name;
      const found = merged.get(name);
      if (found === undefined) merged.set(name, { name, total: target.total, count: 1 });
      else {
        found.total += target.total;
        found.count += 1;
      }
    }
    return [...merged.values()].sort((a, b) => b.total - a.total);
  });
</script>

<li class="border-line-soft border-b" data-testid={`actor-${actor.guid}`}>
  <button
    type="button"
    class="grid min-h-11 w-full grid-cols-[28px_auto_minmax(0,1fr)_auto] items-center gap-x-3 gap-y-1 px-2 py-3 text-left text-[14px] md:grid-cols-[28px_40px_minmax(120px,1.4fr)_minmax(0,3fr)_92px_80px_64px] md:py-2"
    aria-expanded={open}
    onclick={() => (open = !open)}
  >
    <span class="text-muted tabular font-mono text-[12px]"
      ><span class="mr-1 inline-block w-3 transition-transform" class:rotate-90={open} aria-hidden="true"
        >›</span
      >{rank}</span
    >

    <span
      class="tabular font-mono text-[12px]"
      class:text-muted={percentile === null}
      style={percentile === null ? undefined : `color: ${percentileToken(percentile.percentile)}`}
      title={percentile === null
        ? parseTitle(parseFallback)
        : parseTitle(percentile.percentile, percentile.ranked)}
      data-testid="row-percentile"
    >
      {#if percentile === null}{parseFallback}{:else if percentile.ranked === 1}<span class="text-muted"
          >only</span
        >{:else}{Math.round(percentile.percentile)}{#if percentile.ranked > 0}<span
            class="text-muted ml-1 text-[10px]">of {percentile.ranked}</span
          >{/if}{/if}
    </span>

    <span class="truncate font-semibold" style={`color: ${color}`} data-testid="row-name">
      {#if characterLink}
        <a
          href={characterHref(characterLink.region, characterLink.ruleset, display.name)}
          style={`color: ${color}`}
        >
          {display.name}
        </a>
      {:else}
        {display.name}
      {/if}
    </span>

    <span class="col-span-4 row-start-2 md:col-span-1 md:row-auto">
      <AbilityBar abilities={actor.abilities} total={actor.effective} {peak} {color} />
    </span>

    <span class="tabular flex flex-col text-right font-mono leading-tight" data-testid="row-amount">
      <span>{formatAmount(actor.effective)}</span>
      {#if overhealPct !== null}
        <span
          class="text-muted text-[11px]"
          title={approximate
            ? 'Overhealing, scaled to the window in proportion to the total'
            : 'Healing that landed on a full health bar'}
          data-testid="row-overheal">{mark}{formatPercent(overhealPct)} over</span
        >
      {/if}
    </span>

    <!-- The card's third line, below the bar: Active on the left, Per sec on the right.
         It sits here in source order rather than after the two cells it replaces because a
         grid places its row-3 items in source order, and it is display:none above `md`, so
         where it sits costs the desktop grid nothing. -->
    <span class="text-muted label col-span-2 row-start-3 md:hidden" data-testid="phone-labels">
      <span class="tabular font-mono">{activeText}</span> active
    </span>
    <!-- One element, both layouts: a column under ActorTable's "Per sec" heading above
         `md`, and the same figure saying what it is once that heading is gone. -->
    <span
      class="text-muted tabular col-span-2 row-start-3 text-right font-mono text-[13px] md:col-span-1 md:row-auto"
      data-testid="row-per-second"
    >
      <span class="flex flex-col leading-tight">
        <span title={actor.time_ms === undefined ? undefined : 'Per second of the pulls this player was in'}
          >{formatPerSecond(actor.effective, actor.time_ms ?? durationMs)}<span
            class="label font-body ml-1.5 md:hidden">per sec</span
          ></span
        >
        {#if actor.active_ms > 0 && actor.active_ms < durationMs}
          <span
            class="text-[11px]"
            title="Per second over the time this row was active, not the whole window"
            data-testid="row-active-per-second"
            >{formatPerSecond(actor.effective, actor.active_ms)} active</span
          >
        {/if}
      </span>
    </span>
    <span
      class="text-muted tabular hidden text-right font-mono text-[13px] md:inline"
      data-testid="row-active"
    >
      {activeText}
    </span>
  </button>

  {#if open}
    <div
      class="bg-card-top flex flex-col gap-4 overflow-x-auto px-2 py-3 md:flex-row"
      data-testid="row-detail"
    >
      <table class="flex-1 text-[13px]">
        <caption class="label text-muted text-left">Abilities</caption>
        <tbody>
          {#each [...actor.abilities]
            .filter((ability) => ability.total > 0 || ability.hits + ability.ticks > 0)
            .sort((a, b) => b.total - a.total) as ability (ability.spell_id)}
            <tr class="border-line-soft border-b">
              <td class="py-1 pr-3"
                >{ability.name}{#if sameName.has(ability.name)}
                  <span
                    class="text-muted ml-1 font-mono text-[11px]"
                    title="Two spells share this name; this is spell id {ability.spell_id}"
                    >#{ability.spell_id}</span
                  >{/if}{#if schoolName(ability.school)}
                  <span class="text-muted ml-1 text-[11px]">{schoolName(ability.school)}</span>{/if}</td
              >
              <td
                class="tabular py-1 pr-3 text-right font-mono"
                {title}
                aria-label={approximateAriaLabel(approximate, formatAmount(ability.total))}
              >
                {mark}{formatAmount(ability.total)}
              </td>
              <td class="text-muted tabular py-1 pr-3 text-right font-mono"
                >{ability.hits + ability.ticks} hits</td
              >
              <td class="text-muted tabular py-1 pr-3 text-right font-mono" title="Largest single hit"
                >{#if ability.max > 0}max {formatAmount(ability.max)}{/if}</td
              >
              <td class="text-muted tabular py-1 pr-3 text-right font-mono"
                >{ability.crits} crits{#if ability.hits + ability.ticks > 0}
                  <span class="text-[11px]"
                    >({formatPercent((ability.crits / (ability.hits + ability.ticks)) * 100)})</span
                  >{/if}</td
              >
              {#if ability.overheal !== undefined && ability.total > 0}
                <td class="text-muted tabular py-1 text-right font-mono" title="Overhealing"
                  >{formatPercent((ability.overheal / ability.total) * 100)} over</td
                >
              {:else}
                <td class="text-muted tabular py-1 text-right font-mono text-[12px]">
                  {#if ability.absorbed}<span title="Absorbed by shields"
                      >{formatAmount(ability.absorbed)} absorbed</span
                    >{/if}
                  {#if ability.blocked}<span class="ml-2" title="Blocked"
                      >{formatAmount(ability.blocked)} blocked</span
                    >{/if}
                  {#if ability.misses !== undefined && Object.keys(ability.misses).length > 0}
                    <span class="ml-2" title="Avoided or fully absorbed, by type"
                      >{Object.entries(ability.misses)
                        .map(([type, count]) => `${count} ${type.toLowerCase()}`)
                        .join(', ')}</span
                    >
                  {/if}
                </td>
              {/if}
            </tr>
          {/each}
        </tbody>
      </table>
      <table class="flex-1 text-[13px]">
        <caption class="label text-muted text-left">Targets</caption>
        <tbody>
          {#each targetsByName as target (target.name)}
            <tr class="border-line-soft border-b">
              <td class="py-1 pr-3"
                >{target.name}{#if target.count > 1}
                  <span class="text-muted tabular font-mono text-[12px]">×{target.count}</span>{/if}</td
              >
              <td
                class="tabular py-1 text-right font-mono"
                {title}
                aria-label={approximateAriaLabel(approximate, formatAmount(target.total))}
              >
                {mark}{formatAmount(target.total)}
              </td>
            </tr>
          {/each}
        </tbody>
      </table>
    </div>
  {/if}
</li>
