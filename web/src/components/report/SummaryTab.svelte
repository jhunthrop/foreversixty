<!-- web/src/components/report/SummaryTab.svelte -->
<!-- The tab the report opens on: who was there, what they did, and what they brought.
     Task 14 adds the build link to each combatant row.

     `summary.roster[].activity_pct` and `.active_ms` are not part of scopeRoster's rewrite
     (web/src/lib/report/window.ts) -- under a brushed window they are silently still the
     whole fight's, because the summary never tracked activity over time in the roster row
     the way an Actor's own series does. Damage, healing, taken and deaths ARE rescoped
     exactly. Marking Active with the same `~` the amount tables use would claim a scaling
     accuracy it does not have, so it gets the whole-fight mark (`wholeFightMark`) instead,
     with a note explaining why -- the pattern Task 12 should reuse for interrupts, dispels,
     resource totals and aura max_stacks, which are whole-fight for the same reason. -->
<script lang="ts">
  import { splitUnitName } from '../../lib/characters';
  import {
    classColorVar, formatAmount, formatDuration, formatPercent, percentileToken,
    wholeFightAriaLabel, wholeFightMark, wholeFightTitle,
  } from '../../lib/report/format';
  import type { Summary } from '../../lib/report/types';

  let {
    summary,
    durationMs,
    percentiles = new Map<string, number>(),
    approximate = false,
  }: {
    summary: Summary;
    durationMs: number;
    percentiles?: Map<string, number>;
    approximate?: boolean;
  } = $props();

  const roster = $derived([...summary.roster].sort((a, b) => b.damage_done - a.damage_done));
  const missingBuffs = $derived(
    summary.combatants.filter((row) => row.missing_buffs.length > 0).length,
  );
  const activeMark = $derived(wholeFightMark(approximate));
  const activeTitle = $derived(wholeFightTitle(approximate));
</script>

<section class="flex flex-col gap-4" data-testid="summary-tab">
  <p class="text-muted text-[13px]">
    <span class="font-mono tabular">{roster.length}</span> players ·
    <span class="font-mono tabular">{summary.deaths.length}</span> deaths ·
    <span class="font-mono tabular">{formatDuration(durationMs)}</span>
    {#if missingBuffs > 0}
      · <span class="font-mono tabular">{missingBuffs}</span> missing a raid buff at pull
    {/if}
  </p>

  <div class="flex flex-col">
    <div class="text-muted label hidden grid-cols-[40px_minmax(120px,1.4fr)_88px_96px_96px_96px_72px_56px] gap-x-3 px-2 pb-1 md:grid">
      <span>Parse</span>
      <span>Name</span>
      <span>Spec</span>
      <span class="text-right">Damage</span>
      <span class="text-right">Healing</span>
      <span class="text-right">Taken</span>
      <span class="text-right">Active</span>
      <span class="text-right">Deaths</span>
    </div>
    <ul class="flex flex-col">
      {#each roster as row (row.guid)}
        {@const display = splitUnitName(row.name)}
        {@const percentile = percentiles.get(row.guid) ?? null}
        <li
          class="border-line-soft grid min-h-11 grid-cols-[minmax(0,1fr)_auto] items-center gap-x-3 gap-y-1 border-b px-2 py-2 text-[14px] md:grid-cols-[40px_minmax(120px,1.4fr)_88px_96px_96px_96px_72px_56px]"
          data-testid={`roster-${row.guid}`}
        >
          <span
            class="font-mono tabular order-2 text-[12px] md:order-none"
            style={percentile === null ? undefined : `color: ${percentileToken(percentile)}`}
          >
            {percentile === null ? '' : Math.round(percentile)}
          </span>
          <span class="truncate font-semibold" style={`color: ${classColorVar(row.class)}`}>
            {display.name}
          </span>
          <span class="text-muted text-[13px]">{row.spec ?? row.class ?? 'Unknown'}</span>
          <span class="font-mono tabular text-right">{formatAmount(row.damage_done)}</span>
          <span class="font-mono tabular text-right">{formatAmount(row.healing_done)}</span>
          <span class="font-mono tabular text-right">{formatAmount(row.damage_taken)}</span>
          <span
            class="text-muted font-mono tabular text-right text-[13px]"
            title={activeTitle}
            aria-label={wholeFightAriaLabel(approximate, formatPercent(row.activity_pct))}
            data-testid="roster-active"
          >
            {activeMark}{formatPercent(row.activity_pct)}
          </span>
          <span class="font-mono tabular text-right">{row.deaths}</span>
        </li>
      {/each}
    </ul>
  </div>

  {#if approximate}
    <p class="text-muted text-[12px]" data-testid="summary-wholefight-note">
      Active is marked {activeMark} because the summary does not track it over time: it is
      the whole fight's figure even inside a shorter window. Damage, healing, taken and
      deaths above are this window's own numbers.
    </p>
  {/if}

  {#if summary.combatants.length > 0}
    <div class="flex flex-col gap-2">
      <h3 class="label text-muted">At pull</h3>
      <ul class="flex flex-col" data-testid="combatants">
        {#each summary.combatants as combatant (combatant.guid)}
          {@const display = splitUnitName(combatant.name)}
          <li class="border-line-soft flex min-h-11 flex-wrap items-center gap-x-3 gap-y-1 border-b px-2 py-2 text-[14px]">
            <span class="font-semibold">{display.name}</span>
            <span class="text-muted text-[13px]">
              {combatant.spec ?? 'Unknown spec'} · item level
              <span class="font-mono tabular">{combatant.item_level ?? 0}</span>
            </span>
            <span class="text-muted text-[13px]">
              <span class="font-mono tabular">{combatant.gear.filter((item) => item.ID > 0).length}</span> items
            </span>
            {#if combatant.missing_buffs.length > 0}
              <span class="pill pill-sample">
                missing <span class="font-mono tabular">{combatant.missing_buffs.length}</span> buffs
              </span>
            {/if}
          </li>
        {/each}
      </ul>
    </div>
  {/if}
</section>
