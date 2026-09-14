<!-- web/src/components/report/ActorRow.svelte -->
<!-- One table row, exactly the shape spec section 4 lists: parse percentile, class-coloured
     name, amount bar segmented by ability, active time, per-second figure, expander. On
     phone the row becomes a card: the same fields, stacked, with the bar full width.

     Active time: `actor.active_ms` is exact even under a brushed window (window.ts measures
     it from the one-second series rather than scaling it), but that series is one-second
     grained, so a percentage of a short window quantises hard -- a 3s window can only ever
     read 0%, 33%, 67% or 100%. Below ACTIVITY_SECONDS_BELOW_MS this column reads seconds
     instead of a percentage: "2s active" is honest at that grain, "67%" implies precision
     the data does not have. -->
<script lang="ts">
  import { characterHref, splitUnitName } from '../../lib/characters';
  import {
    approximateMark, approximateTitle, classColorVar, formatAmount, formatPercent,
    formatPerSecond, percentileToken,
  } from '../../lib/report/format';
  import type { Actor } from '../../lib/report/types';
  import AbilityBar from './AbilityBar.svelte';

  let {
    rank,
    actor,
    peak,
    durationMs,
    percentile = null,
    approximate = false,
    characterLink = null,
  }: {
    rank: number;
    actor: Actor;
    peak: number;
    durationMs: number;
    percentile?: number | null;
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
  const mark = $derived(approximateMark(approximate));
  const title = $derived(approximateTitle(approximate));
</script>

<li class="border-line-soft border-b" data-testid={`actor-${actor.guid}`}>
  <button
    type="button"
    class="grid min-h-11 w-full grid-cols-[28px_minmax(0,1fr)_auto] items-center gap-x-3 gap-y-1 px-2 py-2 text-left text-[14px] md:grid-cols-[28px_40px_minmax(120px,1.4fr)_minmax(0,3fr)_92px_80px_64px]"
    aria-expanded={open}
    onclick={() => (open = !open)}
  >
    <span class="text-muted font-mono tabular text-[12px]">{rank}</span>

    <span
      class="font-mono tabular text-[12px]"
      style={percentile === null ? undefined : `color: ${percentileToken(percentile)}`}
      data-testid="row-percentile"
    >
      {percentile === null ? '' : Math.round(percentile)}
    </span>

    <span class="truncate font-semibold" style={`color: ${color}`} data-testid="row-name">
      {#if characterLink}
        <a href={characterHref(characterLink.region, characterLink.ruleset, display.name)} style={`color: ${color}`}>
          {display.name}
        </a>
      {:else}
        {display.name}
      {/if}
    </span>

    <span class="col-span-3 md:col-span-1">
      <AbilityBar abilities={actor.abilities} total={actor.effective} {peak} {color} />
    </span>

    <span class="font-mono tabular text-right" data-testid="row-amount" {title}>
      {mark}{formatAmount(actor.effective)}
    </span>
    <span class="text-muted font-mono tabular text-right text-[13px]" data-testid="row-per-second">
      {formatPerSecond(actor.effective, durationMs)}
    </span>
    <span class="text-muted font-mono tabular text-right text-[13px]" data-testid="row-active">
      {showActivitySeconds ? `${activitySeconds}s` : formatPercent(activityPct)}
    </span>
  </button>

  {#if open}
    <div class="bg-card-top flex flex-col gap-4 px-2 py-3 md:flex-row" data-testid="row-detail">
      <table class="flex-1 text-[13px]">
        <caption class="label text-muted text-left">Abilities</caption>
        <tbody>
          {#each [...actor.abilities].sort((a, b) => b.total - a.total) as ability (ability.spell_id)}
            <tr class="border-line-soft border-b">
              <td class="py-1 pr-3">{ability.name}</td>
              <td class="font-mono tabular py-1 pr-3 text-right">{mark}{formatAmount(ability.total)}</td>
              <td class="text-muted font-mono tabular py-1 pr-3 text-right">{ability.hits + ability.ticks} hits</td>
              <td class="text-muted font-mono tabular py-1 text-right">{ability.crits} crits</td>
            </tr>
          {/each}
        </tbody>
      </table>
      <table class="flex-1 text-[13px]">
        <caption class="label text-muted text-left">Targets</caption>
        <tbody>
          {#each actor.targets as target (target.guid)}
            <tr class="border-line-soft border-b">
              <td class="py-1 pr-3">{splitUnitName(target.name).name}</td>
              <td class="font-mono tabular py-1 text-right">{mark}{formatAmount(target.total)}</td>
            </tr>
          {/each}
        </tbody>
      </table>
    </div>
  {/if}
</li>
