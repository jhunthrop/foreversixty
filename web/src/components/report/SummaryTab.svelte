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
  import { plannerLinkFor } from '../../lib/report/planner-link';
  import type { RosterRow, Summary } from '../../lib/report/types';

  let {
    summary,
    durationMs,
    percentiles = new Map<string, number>(),
    approximate = false,
    dataBuild = '',
    classOf = new Map<string, string>(),
    treeSizesFor = () => [],
  }: {
    summary: Summary;
    durationMs: number;
    percentiles?: Map<string, number>;
    approximate?: boolean;
    dataBuild?: string;
    classOf?: Map<string, string>;
    treeSizesFor?: (className: string | undefined) => number[];
  } = $props();

  const roster = $derived([...summary.roster].sort((a, b) => b.damage_done - a.damage_done));
  const missingBuffs = $derived(
    summary.combatants.filter((row) => row.missing_buffs.length > 0).length,
  );
  const activeMark = $derived(wholeFightMark(approximate));
  const activeTitle = $derived(wholeFightTitle(approximate));

  interface Figure {
    label: string;
    value: string;
    title?: string;
    ariaLabel?: string;
  }

  /**
   * The row's five figures with the words the column headings carry above `md`. Built
   * here rather than written out five times in the markup so the phone strip cannot drift
   * from the desktop columns it stands in for.
   */
  function figuresFor(row: RosterRow): Figure[] {
    return [
      { label: 'Damage', value: formatAmount(row.damage_done) },
      { label: 'Healing', value: formatAmount(row.healing_done) },
      { label: 'Taken', value: formatAmount(row.damage_taken) },
      {
        label: 'Active',
        value: `${activeMark}${formatPercent(row.activity_pct)}`,
        title: activeTitle,
        ariaLabel: wholeFightAriaLabel(approximate, formatPercent(row.activity_pct)),
      },
      { label: 'Deaths', value: String(row.deaths) },
    ];
  }
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
            class="font-mono tabular col-start-2 row-start-1 text-right text-[12px] md:col-auto md:row-auto md:text-left"
            style={percentile === null ? undefined : `color: ${percentileToken(percentile)}`}
          >
            {percentile === null ? '' : Math.round(percentile)}
          </span>
          <span class="truncate font-semibold" style={`color: ${classColorVar(row.class)}`}>
            {display.name}
          </span>
          <span class="text-muted col-span-2 text-[13px] md:col-span-1">
            {row.spec ?? row.class ?? 'Unknown'}
          </span>
          <span class="font-mono tabular hidden text-right md:inline">{formatAmount(row.damage_done)}</span>
          <span class="font-mono tabular hidden text-right md:inline">{formatAmount(row.healing_done)}</span>
          <span class="font-mono tabular hidden text-right md:inline">{formatAmount(row.damage_taken)}</span>
          <span
            class="text-muted font-mono tabular hidden text-right text-[13px] md:inline"
            title={activeTitle}
            aria-label={wholeFightAriaLabel(approximate, formatPercent(row.activity_pct))}
            data-testid="roster-active"
          >
            {activeMark}{formatPercent(row.activity_pct)}
          </span>
          <span class="font-mono tabular hidden text-right md:inline">{row.deaths}</span>

          <!-- The five figures above are columns under headings in the header row, which is
               `hidden` below `md`. A card has no headings, so on phone they are replaced by
               this strip, where each figure states what it is. Active keeps the whole-fight
               mark, its title and its composed accessible name: a card layout that dropped
               them would be claiming a precision the summary does not have. -->
          <span class="text-muted label col-span-2 flex flex-wrap gap-x-3 gap-y-1 md:hidden" data-testid="roster-figures">
            {#each figuresFor(row) as figure (figure.label)}
              <span>
                {figure.label}
                <span class="font-mono tabular" title={figure.title} aria-label={figure.ariaLabel}>{figure.value}</span>
              </span>
            {/each}
          </span>
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
          {@const link = plannerLinkFor({
            dataBuild,
            className: classOf.get(combatant.guid),
            treeSizes: treeSizesFor(classOf.get(combatant.guid)),
            combatant,
          })}
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
            {#if link}
              <a
                class="text-gold ml-auto inline-flex min-h-11 items-center text-[13px] md:min-h-0"
                href={link.href}
                data-testid="combatant-build-link"
              >
                {link.label}
              </a>
            {/if}
          </li>
        {/each}
      </ul>
    </div>
  {/if}
</section>
