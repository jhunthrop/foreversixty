<!-- web/src/components/sim/tools/SourcePicker.svelte -->
<!-- Design 6.1. A raid card is one checkbox for every boss; each boss is its own. A source
     gated on a phase that has not opened is hidden until "show unreleased content", and
     then carries the date it opens rather than pretending it is live.
     There are no zone icons in the data (src/content/zones/*.md carries a title, a continent and
     a level range, contract 10.7), so a row leads with its best item's icon instead. Each kind is
     a Disclosure titled "n of m ticked", open when its list is short; its sources sit in columns
     (one at 390, up to three at 1440); a row names its usable items and their best item level
     once, a dungeon its boss count inline; a source this class can use nothing from is left out;
     reputation is one row per faction with its standings as inline ticks; PvP lists the
     character's own side and names each rank's best reward; a kind that has not opened says so in
     one line (lib/sim/source-rows.ts). -->
<script lang="ts">
  import { SvelteSet } from 'svelte/reactivity';
  import { bulkCopy, toolFixCopy } from '../../../lib/sim/copy';
  import { sourceGroupVisibility, triedCount } from '../../../lib/sim/drop-picks';
  import {
    bossName,
    groupSources,
    isOpen,
    itemsOfSource,
    professionSplit,
    sourceLabel,
    SOURCE_KIND_LABELS,
    type LootFile,
    type LootSource,
  } from '../../../lib/sim/loot';
  import { PHASE_LATER, openDateLabel, phaseLabel, type PhaseRow } from '../../../lib/sim/phase';
  import {
    gatedLine,
    groupByFaction,
    pvpForFaction,
    showsSource,
    sourceItemStats,
    tickedSources,
  } from '../../../lib/sim/source-rows';
  import { rarityClassFor } from '../../../lib/planner/items';
  import type { Faction } from '../../../lib/bis/types';
  import type { Item } from '../../../lib/planner/types';
  import Disclosure from '../../ui/Disclosure.svelte';
  import ItemIcon from '../../ui/ItemIcon.svelte';
  import ToolSection from './ToolSection.svelte';

  let {
    loot,
    phases,
    shownKinds,
    showUpcoming,
    picked,
    professions,
    items,
    known,
    now,
    treeVersion = '',
    faction = null,
    ontogglekind,
    ontoggleupcoming,
    ontoggle,
    plate = false,
  }: {
    loot: LootFile;
    phases: readonly PhaseRow[];
    shownKinds: readonly string[];
    showUpcoming: boolean;
    picked: readonly string[];
    professions: readonly string[];
    /** This class's item database, for `triedCount` -- the picker's badges. */
    items: ReadonlyMap<number, Item>;
    /** The engine's `simitems.json` ids, `null` when the build ships none. */
    known: ReadonlySet<number> | null;
    now: Date;
    /** The build whose icons the rows draw. */
    treeVersion?: string;
    /** The character's side: PvP ranks of the other side are left out. */
    faction?: Faction | null;
    ontogglekind: (kind: string) => void;
    ontoggleupcoming: (value: boolean) => void;
    ontoggle: (sourceId: string, bossId?: string) => void;
    /** The picker sits on the page's one Plate (/sim/drops). */
    plate?: boolean;
  } = $props();

  const groups = $derived(groupSources(loot));
  /** One column at 390, two at 768, three at 1280. */
  const COLUMNS = 'grid grid-cols-1 gap-x-8 md:grid-cols-2 xl:grid-cols-3';
  /** A kind with this many rows or fewer opens by default; a longer list folds to its summary. */
  const OPEN_UP_TO = 8;
  /** The dungeons whose boss list the player opened. */
  const openBosses = new SvelteSet<string>();

  function isPicked(sourceId: string, bossId = ''): boolean {
    return picked.includes(`${sourceId}|${bossId}`);
  }

  /**
   * "First raids, 9 December 2026" for a dated phase, and the no-date for the
   * literal `"later"` -- contract 10.4's "shows as unreleased without a date".
   */
  function gateLabel(source: LootSource): string {
    if (source.opens === undefined || isOpen(phases, source, now)) return '';
    if (source.opens === PHASE_LATER) return bulkCopy.opensLater;
    const date = openDateLabel(phases, source.opens);
    return date === '' ? bulkCopy.notOpenYet : bulkCopy.opensOn(phaseLabel(source.opens), date);
  }

  /** A kind's rows as the picker shows them: PvP cut to the character's side, empty sources left out. */
  function rowsOf(kind: string, shown: readonly LootSource[]): LootSource[] {
    const sided = kind === 'pvp' ? pvpForFaction(shown, faction) : [...shown];
    return sided.filter((source) =>
      showsSource(kind, sourceItemStats(itemsOfSource(source), items, known).count),
    );
  }

  /** "Honored" from "Timbermaw Hold: Honored": the faction is the row's own heading. */
  function standingOf(source: LootSource, factionName: string): string {
    const label = sourceLabel(source);
    return label.startsWith(`${factionName}: `) ? label.slice(factionName.length + 2) : label;
  }

  function toggleBosses(sourceId: string): void {
    if (openBosses.has(sourceId)) openBosses.delete(sourceId);
    else openBosses.add(sourceId);
  }
</script>

{#snippet bosses(source: LootSource)}
  {#each source.bosses ?? [] as boss (boss.id)}
    <label class="text-text flex min-h-11 items-center gap-2 pl-7 text-[13px]">
      <input
        type="checkbox"
        class="accent-gold h-5 w-5"
        data-testid={`sim-source-${boss.id}`}
        checked={isPicked(source.id, boss.id)}
        onchange={() => ontoggle(source.id, boss.id)}
      />
      <!-- Contract 10.4: "a boss with no name in either database is emitted
                       with an empty name". dps D35: the id is not a name a player should
                       read, so an unresolved boss is named by its zone instead
                       (`loot.ts`'s `bossName`) -- the one thing both databases agree on,
                       without inventing a kind ("Chest") neither one states. -->
      <span class="flex-1">{bossName(source, boss)}</span>
      <span class="text-muted text-[13px] tabular-nums">
        {triedCount(boss.items, items, known)}
      </span>
    </label>
  {/each}
{/snippet}

{#snippet sourceRow(source: LootSource, kind: string)}
  {@const stats = sourceItemStats(itemsOfSource(source), items, known)}
  {@const bossCount = (source.bosses ?? []).length}
  <div class="flex flex-col">
    <div class="flex min-h-11 items-center gap-2">
      <label class="text-text flex min-w-0 flex-1 items-center gap-2 text-[14px]">
        <input
          type="checkbox"
          class="accent-gold h-5 w-5 shrink-0"
          data-testid={`sim-source-${source.id}`}
          checked={isPicked(source.id)}
          onchange={() => ontoggle(source.id)}
        />
        {#if stats.best !== null}<ItemIcon
            icon={stats.best.icon}
            quality={stats.best.quality}
            build={treeVersion}
            size={24}
          />{/if}
        <span class="min-w-0 flex-1 truncate">
          <span class="text-strong font-bold">{sourceLabel(source)}</span>
          {#if kind === 'pvp' && stats.best !== null}<span
              class="text-[13px] {rarityClassFor(stats.best.quality)}"
            >
              {stats.best.name}</span
            >{/if}
        </span>
      </label>
      {#if kind === 'dungeon' && bossCount > 0}
        <!-- A dungeon's bosses fold under its row (a real build lists 900): the count is the toggle. -->
        <button
          type="button"
          class="text-muted hover:text-text min-h-11 shrink-0 text-[13px] tabular-nums"
          aria-expanded={openBosses.has(source.id)}
          data-testid={`sim-source-bosses-${source.id}`}
          onclick={() => toggleBosses(source.id)}>{bulkCopy.sourceBosses(bossCount)}</button
        >
      {/if}
      <span class="text-muted shrink-0 text-[13px] tabular-nums">
        {bulkCopy.sourceItems(stats.count)}{#if stats.bestIlvl > 0}<span class="pl-2"
            >{bulkCopy.sourceBestIlvl(stats.bestIlvl)}</span
          >{/if}
      </span>
    </div>
    {#if gateLabel(source) !== ''}
      <p class="text-muted m-0 pl-7 text-[13px]">{gateLabel(source)}</p>
    {/if}
    {#if kind !== 'dungeon' || openBosses.has(source.id)}
      {@render bosses(source)}
    {/if}
  </div>
{/snippet}

{#snippet factionRow(name: string, standings: readonly LootSource[])}
  <div class="flex flex-col py-1">
    <span class="text-strong pt-2 text-[14px] leading-5 font-bold">{name}</span>
    <span class="flex flex-wrap gap-x-4">
      {#each standings as source (source.id)}
        <label class="text-text flex min-h-11 items-center gap-2 text-[13px]">
          <input
            type="checkbox"
            class="accent-gold h-5 w-5"
            data-testid={`sim-source-${source.id}`}
            checked={isPicked(source.id)}
            onchange={() => ontoggle(source.id)}
          />
          {standingOf(source, name)}
          <span class="text-muted tabular-nums"
            >{sourceItemStats(itemsOfSource(source), items, known).count}</span
          >
        </label>
      {/each}
    </span>
  </div>
{/snippet}

{#snippet groupBody(kind: string, shown: readonly LootSource[])}
  {#if kind === 'rep'}
    <div class={COLUMNS}>
      {#each groupByFaction(shown) as rep (rep.name)}
        {@render factionRow(rep.name, rep.standings)}
      {/each}
    </div>
  {:else}
    {#if kind === 'crafted' && professionSplit(shown, professions).mine.length > 0}
      <p class="text-muted m-0 text-[13px]">{bulkCopy.sourcesMyProfessions}</p>
    {/if}
    <div class={COLUMNS}>
      {#each shown as source (source.id)}
        {@render sourceRow(source, kind)}
      {/each}
    </div>
  {/if}
{/snippet}

<ToolSection title={bulkCopy.sourcesTitle} testid="sim-source-picker" {plate}>
  <!-- The kind ticks choose which lists show, not what is simmed: they sit after "Show". -->
  <div class="flex flex-wrap items-center gap-x-5" role="group" aria-labelledby="sim-sources-show">
    <span id="sim-sources-show" class="text-muted text-[14px]">{bulkCopy.sourcesShow}</span>
    {#each groups as group (group.kind)}
      <label class="text-text flex min-h-11 items-center gap-2 text-[14px]">
        <input
          type="checkbox"
          class="accent-gold h-5 w-5"
          data-testid={`sim-kind-${group.kind}`}
          checked={shownKinds.includes(group.kind)}
          onchange={() => ontogglekind(group.kind)}
        />
        {group.label}
      </label>
    {/each}
    <label class="text-text flex min-h-11 items-center gap-2 text-[14px]">
      <input
        type="checkbox"
        class="accent-gold h-5 w-5"
        data-testid="sim-upcoming"
        checked={showUpcoming}
        onchange={(event) => ontoggleupcoming(event.currentTarget.checked)}
      />
      {bulkCopy.showUpcoming}
    </label>
  </div>

  {#each groups as group (group.kind)}
    {@const visibility = sourceGroupVisibility(
      group.kind,
      group.sources,
      shownKinds,
      phases,
      showUpcoming,
      now,
    )}
    {@const shown = rowsOf(group.kind, visibility.shown)}
    {#if visibility.ticked && group.sources.length > 0}
      <div class="pt-2" data-testid={`sim-kind-${group.kind}-group`}>
        {#if visibility.allGated}
          <!-- dps D33 (BLOCKER) / newcomer MAJOR "ticking Raids produces an empty void": a
               ticked kind whose every source is gated names the gate, one line per date; the
               "show unreleased content" tick above opens them. -->
          {@const gated = gatedLine(group.sources, gateLabel, bulkCopy.opensLater)}
          <h3 class="text-strong m-0 pt-2 text-[15px] leading-5 font-bold">
            {SOURCE_KIND_LABELS[group.kind]}
          </h3>
          <div
            class="text-text flex flex-col pt-1 text-[14px] leading-6"
            data-testid={`sim-kind-${group.kind}-gated`}
          >
            {#each gated.dated as line (line)}<span>{line}</span>{/each}
            {#if gated.undated.length > 0}<span class="text-muted"
                >{bulkCopy.sourcesNoDate(gated.undated)}</span
              >{/if}
            <span class="sr-only">{toolFixCopy.sourcesAllGated}</span>
          </div>
        {:else}
          <Disclosure
            summary={SOURCE_KIND_LABELS[group.kind]}
            meta={tickedSources(shown, picked) > 0
              ? bulkCopy.sourcesTicked(tickedSources(shown, picked), shown.length)
              : bulkCopy.sourcesCount(shown.length)}
            open={shown.length <= OPEN_UP_TO}
            testid={`sim-kind-${group.kind}-list`}
          >
            {@render groupBody(group.kind, shown)}
          </Disclosure>
        {/if}
      </div>
    {/if}
  {/each}
</ToolSection>
