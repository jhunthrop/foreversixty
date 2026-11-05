<!-- web/src/components/report/AbilityTable.svelte -->
<!-- Damage by ability (spec 2026-11-04-sim-results 3.6): one row per damaging ability with its
     icon, rank, DPS, share bar, casts, average per landed hit, crit share and the outcome split
     on the row, the way Details! and Warcraft Logs read. Every figure comes from
     lib/report/ability-table.ts; this only draws them. Icons and rank arrive through `decorate`
     once their files load, inside a fixed box, so nothing moves.
     On a phone the table is two columns (Ability, DPS) and the figures join the sub-line. -->
<script lang="ts" module>
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
  import { rarityClassFor } from '../../lib/planner/items';
  import { damageCopy as copy, grouped, oneDecimal } from '../../lib/sim/result-copy';
  import { CRIT_NOTE_TIP } from '../../lib/sim/result-tips';
  import { spellTipKey } from '../../lib/tooltips/game-tip';

  let { table, decorate }: { table: AbilityTable; decorate: AbilityDecorate } = $props();

  let sort = $state<AbilitySort>(DEFAULT_SORT);

  // Held in constants: a block's edge trims literal whitespace in the markup.
  const SEPARATOR = ' · ';
  const SPACE = ' ';

  const SORT_ORDER: Readonly<Record<AbilitySort['direction'], 'descending' | 'ascending'>> = {
    desc: 'descending',
    asc: 'ascending',
  };

  const sorted = $derived(sortAbilityRows(table.rows, sort));
  const ownRows = $derived(sorted.filter((row) => row.pet === null));
  const petNames = $derived([...new Set(sorted.flatMap((row) => (row.pet === null ? [] : [row.pet])))]);

  const ariaSort = (column: SortColumn): 'descending' | 'ascending' | 'none' =>
    sort.column === column ? SORT_ORDER[sort.direction] : 'none';

  const fixedOne = (value: number): string => value.toFixed(1);
  const peakShare = $derived(Math.max(0, ...table.rows.map((row) => row.share)));
  const barWidth = (share: number): number => (peakShare === 0 ? 0 : (share / peakShare) * 100);

  /** Columns after Ability, for the total row's span: DPS, Share, then what this table draws. */
  const figureColumns = $derived(3 + (table.hasBiggestHit ? 1 : 0) + (table.hasUptime ? 1 : 0));
</script>

{#snippet sortHeader(column: SortColumn, label: string, classes: string, tip?: string)}
  <th scope="col" class={`py-2.5 pr-3 text-right align-middle ${classes}`} aria-sort={ariaSort(column)}>
    <button
      type="button"
      class={`quiet-label hover:text-strong focus-visible:outline-gold ml-auto flex min-h-11 items-center gap-1 whitespace-nowrap focus-visible:outline-2 md:min-h-5 ${sort.column === column ? 'text-strong' : ''}`}
      data-game-tip={tip}
      data-testid={`ability-sort-${column}`}
      onclick={() => (sort = nextSort(sort, column))}
    >
      {label}{#if sort.column === column}<span class="text-gold" aria-hidden="true"
          >{sort.direction === 'desc' ? '▾' : '▴'}</span
        >{/if}
    </button>
  </th>
{/snippet}

{#snippet icon(url: string | null)}
  <span
    class="relative block h-7 w-7 shrink-0 overflow-hidden rounded-[6px] border border-black bg-white/[.05] shadow-[0_0_0_1px_rgba(229,185,85,.18)]"
  >
    {#if url !== null}
      <img
        src={url}
        alt=""
        width="28"
        height="28"
        loading="lazy"
        decoding="async"
        class="block h-full w-full object-cover"
      />
    {/if}
  </span>
{/snippet}

{#snippet abilityRow(row: AbilityRow, indented: boolean)}
  {@const extras = decorate(row)}
  <tr
    class="border-line-soft border-t [@media(hover:hover)_and_(pointer:fine)]:hover:bg-[var(--row-hover)]"
    data-testid={`ability-row-${row.key}`}
  >
    <td class={`py-2.5 pr-3 pl-5 ${indented ? 'md:pl-11' : ''}`}>
      <div class="flex items-center gap-3">
        <!-- The icon and name open one tooltip: the spell's, or a swing row's weapon. The name is
             the keyboard's stop, with a 44 px hit area on a phone that takes no height. -->
        <span
          class="relative shrink-0"
          data-game-tip={extras.tip ? spellTipKey(row.spellId) : undefined}
          data-bis-item={extras.weapon?.id}
          data-testid="ability-icon"
        >
          {@render icon(row.hand !== null && extras.weapon !== null ? extras.weapon.iconUrl : extras.iconUrl)}
        </span>
        <div class="min-w-0">
          <p class="text-[15px] leading-5 font-semibold">
            {#if extras.tip}
              <span
                class="text-strong focus-visible:outline-gold relative rounded-[2px] after:absolute after:-inset-y-3 after:inset-x-0 hover:underline focus-visible:outline-2 md:after:hidden"
                tabindex="0"
                role="button"
                data-game-tip={spellTipKey(row.spellId)}
                data-testid="ability-name">{row.name}</span
              >
            {:else}
              <span class="text-strong" data-testid="ability-name">{row.name}</span>
            {/if}
            {#if extras.rank > 0}
              <span class="text-muted text-[13px] font-normal">{copy.rank(extras.rank)}</span>
            {/if}
            {#if extras.weapon !== null}
              <span class="text-muted font-normal" aria-hidden="true">·</span>
              <span
                class={`focus-visible:outline-gold relative rounded-[2px] text-[13px] font-normal hover:underline focus-visible:outline-2 ${rarityClassFor(extras.weapon.quality)}`}
                tabindex="0"
                role="button"
                data-bis-item={extras.weapon.id}
                data-testid="ability-weapon">{extras.weapon.name}</span
              >
            {/if}
          </p>
          <p class="text-muted text-[13px] leading-[18px]" data-testid="ability-subline">
            {#if row.periodic}
              {@const [ticks, tickCrit] = copy.ticks(grouped(row.ticks), fixedOne(row.crit ?? 0))}
              <span class="whitespace-nowrap">{ticks}</span><span aria-hidden="true">{SEPARATOR}</span><span
                class="whitespace-nowrap">{tickCrit}</span
              >
            {:else}
              {#each row.outcomes as outcome, index (outcome.kind)}
                {#if index > 0}<span aria-hidden="true">{SEPARATOR}</span>{/if}
                <span class="whitespace-nowrap"
                  >{copy.outcome[outcome.kind]}
                  <span class="text-text">{fixedOne(outcome.pct)}%</span></span
                >
              {/each}
            {/if}
            {#if extras.swingsBecame > 0}
              {#if row.outcomes.length > 0}<span aria-hidden="true">{SEPARATOR}</span>{/if}
              <span class="whitespace-nowrap">{copy.becameHeroicStrike(grouped(extras.swingsBecame))}</span>
            {/if}
            {#if extras.sources.length > 0}
              {#if row.outcomes.length > 0}<span aria-hidden="true">{SEPARATOR}</span>{/if}
              <span data-testid="ability-sources"
                >{copy.fromSources}{SPACE}{#each extras.sources as source, index (source.name)}
                  {#if index > 0}<span aria-hidden="true">{SEPARATOR}</span>{/if}
                  <span class="whitespace-nowrap"
                    >{source.name}{#if source.count !== null}{SPACE}<span class="text-text"
                        >{grouped(source.count)}</span
                      >{/if}</span
                  >
                {/each}</span
              >
            {/if}
          </p>
          <!-- A phone has no figure columns: the figures join the sub-line, every pair whole. -->
          <p class="text-muted text-[13px] leading-[18px] md:hidden" data-testid="ability-phone-line">
            {#if row.casts !== null}
              <span class="whitespace-nowrap">{copy.phoneCasts(grouped(row.casts), row.hand !== null)}</span>
              <span aria-hidden="true">·</span>
            {/if}
            <span class="whitespace-nowrap">{copy.phoneAvg(grouped(row.avgHit))}</span>
            {#if row.crit !== null && !row.periodic}
              <span aria-hidden="true">·</span>
              <span class="whitespace-nowrap">{copy.phoneCrit(fixedOne(row.crit))}</span>
            {/if}
          </p>
        </div>
      </div>
    </td>
    <td class="py-2.5 pr-5 text-right align-middle md:pr-3">
      <span class="text-strong numeric block text-[15px] font-semibold">{oneDecimal(row.dps)}</span>
      <span class="text-muted numeric block text-[13px] md:hidden">{fixedOne(row.share)}%</span>
    </td>
    <td class="hidden py-2.5 pr-3 align-middle md:table-cell">
      <div class="flex items-center justify-end gap-2.5">
        <span class="bg-line-soft block h-2 w-[120px] overflow-hidden rounded-[4px]" aria-hidden="true">
          <span
            class="block h-full rounded-[4px] bg-[linear-gradient(90deg,var(--color-gold-deep),var(--color-gold))]"
            style:width={`${barWidth(row.share)}%`}
          ></span>
        </span>
        <span class="numeric text-text w-12 text-right text-[15px]">{fixedOne(row.share)}%</span>
      </div>
    </td>
    <td class="numeric text-text hidden py-2.5 pr-3 text-right text-[15px] align-middle md:table-cell">
      {#if row.casts === null}<span class="text-muted">{copy.dash}</span>{:else}{grouped(row.casts)}{/if}
    </td>
    <td class="numeric text-text hidden py-2.5 pr-3 text-right text-[15px] align-middle md:table-cell">
      {grouped(row.avgHit)}
    </td>
    <td class="numeric text-text hidden py-2.5 pr-3 text-right text-[15px] align-middle md:table-cell">
      {row.crit === null ? copy.dash : `${fixedOne(row.crit)}%`}
    </td>
    {#if table.hasBiggestHit}
      <td class="numeric text-text hidden py-2.5 pr-3 text-right text-[15px] align-middle md:table-cell">
        {#if row.biggestHit > 0}{grouped(row.biggestHit)}{:else}<span class="text-muted">{copy.dash}</span>{/if}
      </td>
    {/if}
    {#if table.hasUptime}
      <td class="numeric text-text hidden py-2.5 pr-5 text-right text-[15px] align-middle lg:table-cell">
        {#if row.uptime === null}<span class="text-muted">{copy.dash}</span>{:else}{fixedOne(row.uptime)}%{/if}
      </td>
    {/if}
  </tr>
{/snippet}

<table class="w-full border-collapse text-left" data-testid="ability-table">
  <thead>
    <tr>
      <th scope="col" class="quiet-label py-2.5 pr-3 pl-5 text-left align-middle">{copy.headers.ability}</th>
      {@render sortHeader('dps', copy.headers.dps, 'pr-5 md:pr-3')}
      {@render sortHeader('share', copy.headers.share, 'hidden md:table-cell')}
      {@render sortHeader('casts', copy.headers.casts, 'hidden md:table-cell')}
      {@render sortHeader('avgHit', copy.headers.avgHit, 'hidden md:table-cell')}
      {@render sortHeader('crit', copy.headers.crit, 'hidden md:table-cell', CRIT_NOTE_TIP)}
      {#if table.hasBiggestHit}
        {@render sortHeader('biggestHit', copy.headers.biggestHit, 'hidden md:table-cell')}
      {/if}
      {#if table.hasUptime}
        <th scope="col" class="quiet-label hidden py-2.5 pr-5 text-right align-middle lg:table-cell"
          >{copy.headers.uptime}</th
        >
      {/if}
    </tr>
  </thead>
  <tbody>
    {#each ownRows as row (row.key)}
      {@render abilityRow(row, false)}
    {/each}
    {#each petNames as pet (pet)}
      <tr class="border-line-soft border-t" data-testid="ability-pet-heading">
        <th scope="colgroup" colspan={figureColumns + 3} class="quiet-label py-2.5 pr-3 pl-5 text-left">
          {copy.pet(pet)}
        </th>
      </tr>
      {#each sorted.filter((row) => row.pet === pet) as row (`${pet}:${row.key}`)}
        {@render abilityRow(row, true)}
      {/each}
    {/each}
  </tbody>
  <tfoot>
    <tr class="border-line-soft text-muted border-t text-[13px]" data-testid="ability-total">
      <td class="py-3 pr-3 pl-5">{copy.total(table.rows.length)}</td>
      <td class="text-strong numeric py-3 pr-5 text-right text-[15px] font-semibold md:pr-3">
        {oneDecimal(table.dps)}
      </td>
      <td class="numeric hidden py-3 pr-3 text-right md:table-cell">100%</td>
      <td colspan={figureColumns} class="numeric hidden py-3 pr-5 text-right md:table-cell">
        {copy.totalDamage(grouped(table.total))}
      </td>
    </tr>
  </tfoot>
</table>
