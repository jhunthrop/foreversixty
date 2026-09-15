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
    classColorVar,
    formatAmount,
    formatDuration,
    formatPerSecond,
    formatPercent,
    parseTitle,
    percentileToken,
    wholeFightAriaLabel,
    wholeFightMark,
    wholeFightTitle,
  } from '../../lib/report/format';
  import { plannerLinkFor } from '../../lib/report/planner-link';
  import GearList from './GearList.svelte';
  import SummaryPanels from './SummaryPanels.svelte';
  import ClassIcon from './ClassIcon.svelte';
  import type { Placement } from '../../lib/report/percentile';
  import type { RosterRow, Summary } from '../../lib/report/types';

  let {
    summary,
    durationMs,
    percentiles = new Map<string, Placement>(),
    parseFallback = '',
    approximate = false,
    dataBuild = '',
    classOf = new Map<string, string>(),
    treeSizesFor = () => [],
    onSelectPlayer = undefined,
    players = new Set<string>(),
    onTab = undefined,
  }: {
    summary: Summary;
    durationMs: number;
    percentiles?: Map<string, Placement>;
    /** What an empty Parse cell shows: '' on trash, 'wipe', or a dash for not ranked yet. */
    parseFallback?: string;
    approximate?: boolean;
    dataBuild?: string;
    classOf?: Map<string, string>;
    treeSizesFor?: (className: string | undefined) => number[];
    /** Narrows the page to one player; the name becomes a button when this is given. */
    onSelectPlayer?: (guid: string) => void;
    /** The players' GUIDs, so the panels show the raid and not the trash. */
    players?: ReadonlySet<string>;
    /** Opens one of the deeper tabs from a panel's heading. */
    onTab?: (tab: 'damage-done' | 'healing' | 'damage-taken' | 'deaths') => void;
  } = $props();

  const roster = $derived([...summary.roster].sort((a, b) => b.damage_done - a.damage_done));
  /** The combatant whose gear list is open. */
  let gearOpen = $state<string | null>(null);
  const missingBuffs = $derived(summary.combatants.filter((row) => row.missing_buffs.length > 0).length);
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
    const percentile = percentiles.get(row.guid);
    return [
      ...(percentile === undefined
        ? []
        : [
            {
              label: 'Parse',
              value: String(Math.round(percentile.percentile)),
            },
          ]),
      { label: 'DPS', value: formatPerSecond(row.damage_done, durationMs) },
      { label: 'HPS', value: formatPerSecond(row.healing_done, durationMs) },
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
    <span class="tabular font-mono">{roster.length}</span>
    {roster.length === 1 ? 'player' : 'players'} ·
    <span class="tabular font-mono">{summary.deaths.length}</span> deaths ·
    <span class="tabular font-mono">{formatDuration(durationMs)}</span>
    {#if missingBuffs > 0}
      · <span class="tabular font-mono">{missingBuffs}</span> missing a raid buff at pull
    {/if}
  </p>

  <div class="flex flex-col">
    <div
      class="text-muted label hidden grid-cols-[40px_minmax(120px,1.4fr)_88px_96px_96px_96px_72px_56px] gap-x-3 px-2 pb-1 md:grid"
    >
      <span
        title="Percentile among ranked kills of the same boss by this spec: DPS for damage, HPS for healers, damage taken for tanks. Empty on a wipe or while nothing is ranked yet."
        >Parse</span
      >
      <span>Name</span>
      <span>Spec</span>
      <span class="text-right" title="Damage done, and per second over this window">Damage</span>
      <span class="text-right" title="Healing done, and per second over this window">Healing</span>
      <span class="text-right" title="Damage taken, and per second over this window">Taken</span>
      <span class="text-right" title="Share of the fight spent casting or attacking">Active</span>
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
            class="tabular col-start-2 row-start-1 text-right font-mono text-[12px] md:col-auto md:row-auto md:text-left"
            class:text-muted={percentile === null}
            style={percentile === null ? undefined : `color: ${percentileToken(percentile.percentile)}`}
            title={percentile === null
              ? parseTitle(parseFallback)
              : parseTitle(percentile.percentile, percentile.ranked)}
            data-testid="roster-percentile"
          >
            {#if percentile === null}{parseFallback}{:else}{Math.round(percentile.percentile)}{/if}
          </span>
          <span
            class="flex min-w-0 items-center gap-2 truncate font-semibold"
            style={`color: ${classColorVar(row.class)}`}
          >
            <ClassIcon className={row.class} />
            {#if onSelectPlayer}
              <button
                type="button"
                class="inline-flex min-h-11 items-center underline-offset-2 hover:underline md:min-h-0"
                style={`color: ${classColorVar(row.class)}`}
                title="Show only this player"
                data-testid="roster-name"
                onclick={() => onSelectPlayer(row.guid)}
              >
                {display.name}
              </button>
            {:else}
              {display.name}
            {/if}
          </span>
          <span class="text-muted col-span-2 text-[13px] md:col-span-1">
            {row.spec ?? row.class ?? 'Unknown'}
          </span>
          <!-- Total on top, per-second underneath: a raid leader compares totals across
               the roster, and a player compares their per-second figure with their
               parse. Both are the window's own numbers. -->
          <span
            class="tabular hidden flex-col text-right font-mono leading-tight md:flex"
            data-testid="roster-damage"
          >
            <span>{formatAmount(row.damage_done)}</span>
            <span class="text-muted text-[11px]">{formatPerSecond(row.damage_done, durationMs)}/s</span>
          </span>
          <span
            class="tabular hidden flex-col text-right font-mono leading-tight md:flex"
            data-testid="roster-healing"
          >
            <span>{formatAmount(row.healing_done)}</span>
            <span class="text-muted text-[11px]">{formatPerSecond(row.healing_done, durationMs)}/s</span>
          </span>
          <span
            class="tabular hidden flex-col text-right font-mono leading-tight md:flex"
            data-testid="roster-taken"
          >
            <span>{formatAmount(row.damage_taken)}</span>
            <span class="text-muted text-[11px]">{formatPerSecond(row.damage_taken, durationMs)}/s</span>
          </span>
          <span
            class="text-muted tabular hidden text-right font-mono text-[13px] md:inline"
            title={activeTitle}
            aria-label={wholeFightAriaLabel(approximate, formatPercent(row.activity_pct))}
            data-testid="roster-active"
          >
            {activeMark}{formatPercent(row.activity_pct)}
          </span>
          <span class="tabular hidden text-right font-mono md:inline" class:text-death={row.deaths > 0}
            >{row.deaths}</span
          >

          <!-- The five figures above are columns under headings in the header row, which is
               `hidden` below `md`. A card has no headings, so on phone they are replaced by
               this strip, where each figure states what it is. Active keeps the whole-fight
               mark, its title and its composed accessible name: a card layout that dropped
               them would be claiming a precision the summary does not have. -->
          <span
            class="text-muted label col-span-2 flex flex-wrap gap-x-3 gap-y-1 md:hidden"
            data-testid="roster-figures"
          >
            {#each figuresFor(row) as figure (figure.label)}
              <span>
                {figure.label}
                <span class="tabular font-mono" title={figure.title} aria-label={figure.ariaLabel}
                  >{figure.value}</span
                >
              </span>
            {/each}
          </span>
        </li>
      {/each}
    </ul>
  </div>

  {#if approximate}
    <p class="text-muted text-[12px]" data-testid="summary-wholefight-note">
      Active is marked {activeMark} because the summary does not track it over time: it is the whole fight's figure
      even inside a shorter window. Damage, healing, taken and deaths above are this window's own numbers.
    </p>
  {/if}

  {#if onTab}
    <SummaryPanels {summary} {durationMs} {players} {onTab} {onSelectPlayer} />
  {/if}

  {#if summary.combatants.length > 0}
    <div class="flex flex-col gap-2">
      <h2 class="label text-muted">At pull</h2>
      <ul class="flex flex-col" data-testid="combatants">
        {#each summary.combatants as combatant (combatant.guid)}
          {@const display = splitUnitName(combatant.name)}
          {@const link = plannerLinkFor({
            dataBuild,
            className: classOf.get(combatant.guid),
            treeSizes: treeSizesFor(classOf.get(combatant.guid)),
            combatant,
          })}
          <li
            class="border-line-soft flex min-h-11 flex-wrap items-center gap-x-3 gap-y-1 border-b px-2 py-2 text-[14px]"
          >
            <span class="font-semibold">{display.name}</span>
            <span class="text-muted text-[13px]">
              {combatant.spec ?? 'Unknown spec'} · item level
              <span class="tabular font-mono">{combatant.item_level ?? 0}</span>
            </span>
            <button
              type="button"
              class="text-muted inline-flex min-h-11 items-center text-[13px] underline-offset-2 hover:underline md:min-h-0"
              aria-expanded={gearOpen === combatant.guid}
              data-testid="combatant-gear"
              onclick={() => (gearOpen = gearOpen === combatant.guid ? null : combatant.guid)}
            >
              <span class="tabular font-mono">{combatant.gear.filter((item) => item.ID > 0).length}</span
              >&nbsp;items
            </button>
            {#if combatant.missing_buffs.length > 0}
              <span class="pill pill-sample">
                missing <span class="tabular font-mono">{combatant.missing_buffs.length}</span> buffs
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
            {#if gearOpen === combatant.guid}
              <div class="basis-full pt-2">
                <GearList gear={combatant.gear} className={classOf.get(combatant.guid)} {dataBuild} />
              </div>
            {/if}
          </li>
        {/each}
      </ul>
    </div>
  {/if}
</section>
