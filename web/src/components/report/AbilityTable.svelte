<!-- web/src/components/report/AbilityTable.svelte -->
<!-- Damage by ability (sim spec 2026-11-06 §4.7, on spec 2026-11-04-sim-results 3.6): one StatTable
     row per damaging ability, its icon, name and rank, the outcome split on its sub line, the share
     Meter, then DPS, Share, Casts, Avg hit and Crit. A DoT's uptime joins its sub line (the Uptime
     column retired), and a spell that hits several enemies says its hits a cast. Biggest hit draws
     only once the engine records one (gap G14). Every figure comes from lib/report/ability-table.ts;
     this only draws them. Icons and rank arrive through `decorate` once their files load, inside a
     fixed box, so nothing moves. Below 1024 the figures join the sub line (StatRow's phone parts). -->
<script lang="ts" module>
  import ItemName from '../ui/ItemName.svelte';
  import SpellIcon from '../ui/SpellIcon.svelte';
  import type { AbilityRow } from '../../lib/report/ability-table';

  export interface AbilityWeapon {
    id: number;
    name: string;
    quality: number;
    iconUrl: string | null;
  }

  /** What a row needs beyond its figures: all of it optional until the build's files load. */
  export interface AbilityDecoration {
    iconUrl: string | null;
    /** 0 for a spell with no rank, or none known yet. */
    rank: number;
    /** A white swing's weapon. */
    weapon: AbilityWeapon | null;
    /** The Extra attacks row's sources. */
    sources: readonly { name: string; count: number | null }[];
    /** On the main-hand row: the swings Heroic Strike replaced. */
    swingsBecame: number;
    /** Whether the row's icon and name open a spell tooltip. */
    tip: boolean;
    /** A spell whose tooltip hits several enemies: the fight's target count, else null. */
    targets: number | null;
  }

  export type AbilityDecorate = (row: AbilityRow) => AbilityDecoration;
</script>

<script lang="ts">
  import {
    DEFAULT_SORT,
    nextSort,
    sortAbilityRows,
    type AbilitySort,
    type AbilityTable,
    type SortColumn,
  } from '../../lib/report/ability-table';
  import { damageCopy as copy } from '../../lib/sim/result-copy';
  import { CRIT_NOTE_TIP } from '../../lib/sim/result-tips';
  import { spellTipKey } from '../../lib/tooltips/game-tip';
  import type { StatColumn, StatFigure } from '../../lib/ui/stat-table';
  import { grouped, oneDecimal } from '../../lib/format/number';
  import Meter from '../ui/Meter.svelte';
  import StatRow from '../ui/StatRow.svelte';
  import StatTable from '../ui/StatTable.svelte';

  let { table, decorate }: { table: AbilityTable; decorate: AbilityDecorate } = $props();

  let sort = $state<AbilitySort>(DEFAULT_SORT);

  // Held in constants: a block's edge trims literal whitespace in the markup.
  const SEPARATOR = ' · ';
  const SPACE = ' ';

  const sorted = $derived(sortAbilityRows(table.rows, sort));
  const ownRows = $derived(sorted.filter((row) => row.pet === null));
  const petNames = $derived([...new Set(sorted.flatMap((row) => (row.pet === null ? [] : [row.pet])))]);

  const fixedOne = (value: number): string => value.toFixed(1);
  const peakShare = $derived(Math.max(0, ...table.rows.map((row) => row.share)));

  const columns = $derived<StatColumn[]>([
    { id: 'dps', label: copy.headers.dps, sortable: true },
    { id: 'share', label: copy.headers.share, sortable: true },
    { id: 'casts', label: copy.headers.casts, sortable: true },
    { id: 'avgHit', label: copy.headers.avgHit, sortable: true },
    { id: 'crit', label: copy.headers.crit, sortable: true, tip: CRIT_NOTE_TIP },
    ...(table.hasBiggestHit ? [{ id: 'biggestHit', label: copy.headers.biggestHit, sortable: true }] : []),
  ]);

  const DASH: StatFigure = { value: copy.dash, muted: true };

  function figuresOf(row: AbilityRow): StatFigure[] {
    return [
      { value: oneDecimal(row.dps) },
      { value: `${fixedOne(row.share)}%` },
      row.casts === null ? DASH : { value: grouped(row.casts) },
      { value: grouped(row.avgHit) },
      row.crit === null ? DASH : { value: `${fixedOne(row.crit)}%` },
      ...(table.hasBiggestHit ? [row.biggestHit > 0 ? { value: grouped(row.biggestHit) } : DASH] : []),
    ];
  }

  const totalFigures = $derived<StatFigure[]>([
    { value: oneDecimal(table.dps) },
    { value: '100%' },
    ...columns.slice(2).map(() => ({ value: '' })),
  ]);

  function onsort(column: string): void {
    sort = nextSort(sort, column as SortColumn);
  }
</script>

{#snippet separator()}<span aria-hidden="true">{SEPARATOR}</span>{/snippet}

{#snippet abilityRow(row: AbilityRow)}
  {@const extras = decorate(row)}
  {@const iconUrl = row.hand !== null && extras.weapon !== null ? extras.weapon.iconUrl : extras.iconUrl}
  <StatRow figures={figuresOf(row)} phone={`${fixedOne(row.share)}%`} testid={`ability-row-${row.key}`}>
    {#snippet icon()}
      <!-- The icon opens the same tooltip as the name: the spell's, or a swing row's weapon. -->
      <SpellIcon
        socket
        url={iconUrl}
        size={28}
        tip={extras.tip ? row.spellId : undefined}
        item={extras.weapon?.id}
        testid="ability-icon"
      />
    {/snippet}
    {#snippet name()}
      <!-- The name is the keyboard's stop for the tooltip; a phone gets the 44 px hit area. -->
      <!-- svelte-ignore a11y_no_noninteractive_tabindex -->
      <span
        class={[
          'relative',
          extras.tip &&
            'focus-visible:outline-gold rounded-[2px] after:absolute after:inset-x-0 after:-inset-y-3 hover:underline focus-visible:outline-2 md:after:hidden',
        ]}
        tabindex={extras.tip ? 0 : undefined}
        data-game-tip={extras.tip ? spellTipKey(row.spellId) : undefined}
        data-testid="ability-name">{row.name}</span
      >{#if extras.rank > 0}<span class="text-muted ml-1 text-[13px] font-normal"
          >{copy.rank(extras.rank)}</span
        >{/if}{#if extras.weapon !== null}{SPACE}<ItemName
          name={`· ${extras.weapon.name}`}
          quality={extras.weapon.quality}
          id={extras.weapon.id}
          class="text-[14px] font-semibold"
          testid="ability-weapon"
        />{/if}
    {/snippet}
    {#snippet sub()}
      <div data-testid="ability-subline">
        {#if row.periodic}
          {@const [ticks, tickCrit] = copy.ticks(grouped(row.ticks), fixedOne(row.crit ?? 0))}
          <span class="whitespace-nowrap">{ticks}</span>{@render separator()}<span class="whitespace-nowrap"
            >{tickCrit}</span
          >{#if row.uptime !== null}{@render separator()}<span class="whitespace-nowrap"
              >{copy.upFor(fixedOne(row.uptime))}</span
            >{/if}
        {:else}
          {#each row.outcomes as outcome, index (outcome.kind)}
            {#if index > 0}{@render separator()}{/if}
            <span class="whitespace-nowrap">{copy.outcome[outcome.kind]} {fixedOne(outcome.pct)}%</span>
          {/each}
        {/if}
        {#if extras.swingsBecame > 0}
          {#if row.outcomes.length > 0}{@render separator()}{/if}
          <span class="whitespace-nowrap">{copy.becameHeroicStrike(grouped(extras.swingsBecame))}</span>
        {/if}
        {#if extras.targets !== null && row.casts !== null && row.casts > 0}
          {#if row.outcomes.length > 0}{@render separator()}{/if}
          <span class="whitespace-nowrap" data-testid="ability-hits-a-cast"
            >{copy.hitsACast(fixedOne(row.attempts / row.casts), extras.targets)}</span
          >
        {/if}
        {#if extras.sources.length > 0}
          {#if row.outcomes.length > 0}{@render separator()}{/if}
          <span data-testid="ability-sources"
            >{copy.fromSources}{SPACE}{#each extras.sources as source, index (source.name)}
              {#if index > 0}{@render separator()}{/if}
              <span class="whitespace-nowrap"
                >{source.name}{#if source.count !== null}{SPACE}{grouped(source.count)}{/if}</span
              >
            {/each}</span
          >
        {/if}
        <!-- A phone has no figure columns: the figures join the sub line, every pair whole. -->
        {#if row.casts !== null}<span class="ui-st-phone" data-testid="ability-phone-casts"
            >{SEPARATOR}{copy.phoneCasts(grouped(row.casts), row.hand !== null)}</span
          >{/if}<span class="ui-st-phone">{SEPARATOR}{copy.phoneAvg(grouped(row.avgHit))}</span
        >{#if row.crit !== null && !row.periodic}<span class="ui-st-phone"
            >{SEPARATOR}{copy.phoneCrit(fixedOne(row.crit))}</span
          >{/if}
      </div>
    {/snippet}
    {#snippet meter()}<Meter value={peakShare === 0 ? 0 : row.share / peakShare} />{/snippet}
  </StatRow>
{/snippet}

<StatTable
  {columns}
  nameHead={copy.headers.ability}
  barHead={copy.headers.shareOfDamage}
  {sort}
  {onsort}
  label={copy.title}
  testid="ability-table"
>
  {#each ownRows as row (row.key)}
    {@render abilityRow(row)}
  {/each}
  {#each petNames as pet (pet)}
    <div class="ui-st-r" role="row" data-testid="ability-pet-heading">
      <span class="text-label text-muted col-span-full font-bold" role="cell">{copy.pet(pet)}</span>
    </div>
    {#each sorted.filter((row) => row.pet === pet) as row (`${pet}:${row.key}`)}
      {@render abilityRow(row)}
    {/each}
  {/each}
  <StatRow figures={totalFigures} total testid="ability-total">
    {#snippet name()}{copy.total(table.rows.length)}{SEPARATOR}{copy.totalDamage(
        grouped(table.total),
      )}{/snippet}
  </StatRow>
</StatTable>
