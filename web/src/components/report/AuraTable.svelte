<!-- web/src/components/report/AuraTable.svelte -->
<!-- Buffs and Debuffs (logs spec 2026-11-07 section 6.2): the tab's one Plate, a StatTable of
     the aura's icon, its name over `on {target} · from {applier}`, an uptime Meter (muted on
     Debuffs), Up and Applied. A debuff reads once by spell (logs lane review 2026-11-09): on how
     many players, its uptime over the time they were in, every application. Over the night a
     second, open Section reads each spell's uptime across every target that had it.

     Every figure here is exact under any window, not scaled: window.ts's scopeAuraTrack clips
     each segment to the window's own bounds and recomputes `uptime_ms` and `applications`
     from those clipped segments, rather than taking the whole fight's numbers and scaling
     them by a ratio the way the per-ability and per-target splits are. So this table carries
     neither the `~` nor the `†` mark. The one field on AuraTrack that IS whole-fight,
     `max_stacks`, is not rendered here. -->
<script lang="ts">
  import EmptyState from '../ui/EmptyState.svelte';
  import { splitUnitName } from '../../lib/characters';
  import { schoolName } from '../../lib/report/format';
  import activeBuild from '../../data/active-build.json';
  import { loadSharedSpellNames, nameAuras, type SpellNames } from '../../lib/report/aura-names';
  import { SpellIcons } from '../../lib/report/spell-icons.svelte';
  import type { AuraTrack } from '../../lib/report/types';
  import type { StatColumn, StatSort } from '../../lib/ui/stat-table';
  import { percent } from '../../lib/format/number';
  import Meter from '../ui/Meter.svelte';
  import Section from '../ui/Section.svelte';
  import SpellIcon from '../ui/SpellIcon.svelte';
  import StatRow from '../ui/StatRow.svelte';
  import StatTable from '../ui/StatTable.svelte';
  import CopyCsv from './CopyCsv.svelte';

  let {
    tracks,
    durationMs,
    startMs = 0,
    kind,
    bossNames = new Set<string>(),
    names = new Map<string, string>(),
    friendly = new Set<string>(),
    tauntSpells = new Set<number>(),
    emptyMessage,
    windowed,
  }: {
    tracks: AuraTrack[];
    durationMs: number;
    /** Where the window starts on the fight's clock. */
    startMs?: number;
    kind: 'BUFF' | 'DEBUFF';
    /** GUID to unit name, so a debuff can say who applied it. */
    names?: ReadonlyMap<string, string>;
    /** The encounter bosses' unit names, for the per-spell line across the night. */
    bossNames?: ReadonlySet<string>;
    /** The players, their pets and totems: Debuffs splits what was on them from what they put on the enemies. */
    friendly?: ReadonlySet<string>;
    /** The taunts' spell ids: a taunt on an enemy reads on Threat and the Timeline, not as a debuff uptime. */
    tauntSpells?: ReadonlySet<number>;
    /**
     * Overrides the generic "No buffs/debuffs in this window" text. A real fight's empty
     * DEBUFFS tab genuinely means none were cast, so ReportView.svelte's calls leave this
     * unset; a caller whose empty list means something else passes its own honest message.
     */
    emptyMessage?: string;
    /** Whether a window is brushed (the meta says what the uptime is over); by default a window
     *  that does not start at the pull. */
    windowed?: boolean;
  } = $props();

  const title = $derived(kind === 'BUFF' ? 'Buffs' : 'Debuffs');
  const icons = new SpellIcons();

  /** The build's spell names, for auras the log never named (B2); empty until read. */
  let spellNames = $state<SpellNames>({});
  $effect(() => {
    void loadSharedSpellNames(activeBuild.build).then((loaded) => {
      spellNames = loaded;
    });
  });

  /**
   * Uptime over the time the track's target was in: the figure the row shows. Clamped to
   * 100 here, once, as a last line of defence (2026-09-21 result-page review round 3, E8: a
   * combined sim result's aura rows read up to 100.9%), so a future source of the same shape
   * of bug reads as an honest 100%, never a number over it.
   */
  const shareOf = (row: AuraRow): number =>
    Math.min(row.time_ms === 0 ? 0 : (row.uptime_ms / row.time_ms) * 100, 100);

  type AuraSortColumn = 'up' | 'applied';
  let sort = $state<StatSort>({ column: 'up', direction: 'desc' });
  const columns: readonly StatColumn[] = [
    { id: 'up', label: 'Up', sortable: true },
    { id: 'applied', label: 'Applied', sortable: true },
  ];
  function onsort(column: string): void {
    sort =
      sort.column === column
        ? { column, direction: sort.direction === 'desc' ? 'asc' : 'desc' }
        : { column, direction: 'desc' };
  }
  const keyOf = (row: AuraRow, column: AuraSortColumn): number =>
    column === 'up' ? shareOf(row) : row.applications;

  const isUnnamed = (name: string): boolean => name.startsWith('Spell #');
  /** Carried in and never up, or carried in under no name the log or the build knows: nothing to read. */
  const isDead = (track: { applications: number; uptime_ms: number; name: string }): boolean =>
    track.applications === 0 && (track.uptime_ms === 0 || isUnnamed(track.name));
  /** The tracks the table reads: this kind's, named, alive. */
  const shown = $derived(
    nameAuras(tracks, spellNames).filter((track) => track.type === kind && !isDead(track)),
  );

  /** One table row: one buff on one player, or one debuff across every player it was on. */
  interface AuraRow {
    key: string;
    name: string;
    spell_id: number;
    school?: number;
    tracks: AuraTrack[];
    uptime_ms: number;
    /** The time the row's targets were in: the share's denominator. */
    time_ms: number;
    applications: number;
  }
  const rowOf = (group: AuraTrack[]): AuraRow => ({
    key:
      group.length === 1
        ? `${group[0].target_guid}-${group[0].spell_id}`
        : `${group[0].spell_id}-${group[0].name}`,
    name: group[0].name,
    spell_id: group[0].spell_id,
    school: group[0].school,
    tracks: group,
    uptime_ms: group.reduce((sum, track) => sum + track.uptime_ms, 0),
    time_ms: group.reduce((sum, track) => sum + (track.time_ms ?? durationMs), 0),
    applications: group.reduce((sum, track) => sum + track.applications, 0),
  });
  /** Buffs a row per player; debuffs a row per spell. */
  function grouped(list: readonly AuraTrack[]): AuraRow[] {
    if (kind === 'BUFF') return list.map((track) => rowOf([track]));
    // eslint-disable-next-line svelte/prefer-svelte-reactivity
    const bySpell = new Map<string, AuraTrack[]>();
    for (const track of list) {
      const key = `${track.spell_id}|${track.name}`;
      bySpell.set(key, [...(bySpell.get(key) ?? []), track]);
    }
    return [...bySpell.values()].map(rowOf);
  }

  function sorted(list: readonly AuraTrack[]): AuraRow[] {
    return grouped(list).sort((a, b) => {
      // The log's unnamed spells ("Spell #1") go last whatever the sort: a raider reads the named ones.
      const unnamed = Number(isUnnamed(a.name)) - Number(isUnnamed(b.name));
      if (unnamed !== 0) return unnamed;
      const column = sort.column as AuraSortColumn;
      const order = keyOf(b, column) - keyOf(a, column);
      return (sort.direction === 'desc' ? order : -order) || a.name.localeCompare(b.name);
    });
  }
  /** On Debuffs with the players known, the enemies' side: what the group kept on them. */
  const splitSides = $derived(kind === 'DEBUFF' && friendly.size > 0);
  const onEnemy = (track: AuraTrack): boolean => splitSides && !friendly.has(track.target_guid);
  const rows = $derived(sorted(shown.filter((track) => !onEnemy(track))));
  const enemyRows = $derived(
    sorted(shown.filter((track) => onEnemy(track) && !tauntSpells.has(track.spell_id))),
  );

  /**
   * Over the night, one line per spell across every target: uptime over the time the
   * targets were up. Only for folded tracks (those carrying their own time), and for
   * debuffs only those on a boss, since adds come and go.
   */
  const bySpell = $derived.by(() => {
    if (kind === 'DEBUFF' && bossNames.size === 0) return [];
    // eslint-disable-next-line svelte/prefer-svelte-reactivity
    const totals = new Map<
      string,
      { name: string; spellId: number; uptime: number; time: number; targets: number }
    >();
    for (const track of shown) {
      if (track.time_ms === undefined || (kind === 'DEBUFF' && !bossNames.has(track.target_name))) continue;
      const found = totals.get(track.name) ?? {
        name: track.name,
        spellId: track.spell_id,
        uptime: 0,
        time: 0,
        targets: 0,
      };
      totals.set(track.name, {
        ...found,
        uptime: found.uptime + track.uptime_ms,
        time: found.time + track.time_ms,
        targets: found.targets + 1,
      });
    }
    return [...totals.values()]
      .filter((entry) => entry.targets > 1)
      .map((entry) => ({ ...entry, pct: entry.time === 0 ? 0 : (entry.uptime / entry.time) * 100 }))
      .sort((a, b) => b.pct - a.pct);
  });

  /** Names that two different spells share on one target: shown with their spell id. */
  const ambiguous = $derived.by(() => {
    // eslint-disable-next-line svelte/prefer-svelte-reactivity
    const seen = new Map<string, Set<number>>();
    for (const track of shown) {
      const key = `${track.target_guid}|${track.name}`;
      seen.set(key, new Set([...(seen.get(key) ?? []), track.spell_id]));
    }
    return new Set([...seen.entries()].filter(([, ids]) => ids.size > 1).map(([key]) => key));
  });

  const appliersOf = (row: AuraRow): string[] => [...new Set(row.tracks.flatMap((track) => track.appliers))];
  const applierNames = (row: AuraRow): string =>
    appliersOf(row)
      .map((guid) => splitUnitName(names.get(guid) ?? '').name)
      .filter(Boolean)
      .slice(0, 2)
      .join(', ') || 'an unnamed source';
  /** A buff the target gave itself says nothing new by naming its applier. */
  const showApplier = (row: AuraRow): boolean =>
    appliersOf(row).length > 0 &&
    (kind === 'DEBUFF' || appliersOf(row).some((guid) => guid !== row.tracks[0].target_guid));
  /** Who a row's aura was on: the unit, or how many. */
  const onWhom = (row: AuraRow): string =>
    row.tracks.length === 1
      ? splitUnitName(row.tracks[0].target_name).name
      : `${row.tracks.length} ${row.tracks.some(onEnemy) ? 'enemies' : 'players'}`;
  /** A row on several units names each with its own uptime, up to this many; the rest count on. */
  const EACH_MAX = 5;
  const eachTarget = (row: AuraRow): AuraTrack[] =>
    [...row.tracks].sort((a, b) => shareOf(rowOf([b])) - shareOf(rowOf([a])));
  /** A unit's name in a row's list; two units of one name (two Elder Crocolisks) are numbered by GUID. */
  function unitLabel(row: AuraRow, one: AuraTrack): string {
    const name = splitUnitName(one.target_name).name;
    const same = row.tracks.filter((other) => splitUnitName(other.target_name).name === name);
    if (same.length < 2) return name;
    const order = same.map((other) => other.target_guid).sort();
    return `${name} ${order.indexOf(one.target_guid) + 1}`;
  }

  const over = $derived((windowed ?? startMs > 0) ? 'this window' : 'the whole fight');
  /** What the count counts: buffs on players, or the debuffs (each a spell, however many it was on). */
  /** The window the uptime is over, said only when it is not the whole fight the filter already names. */
  const overLine = $derived(over === 'the whole fight' ? '' : ` · uptime over ${over}`);
  const meta = $derived(
    `${rows.length} ${kind === 'BUFF' ? 'on players' : rows.length === 1 ? 'debuff' : 'debuffs'}${overLine}`,
  );
  const enemyMeta = $derived(
    `${enemyRows.length} ${enemyRows.length === 1 ? 'debuff' : 'debuffs'}${overLine}`,
  );

  /** The table as lines: one per aura on one target, uptime in seconds and as a share. */
  function csvLines(): string[][] {
    return [
      ['Aura', 'Spell id', 'On', 'Uptime s', 'Uptime %', 'Applied', 'From'],
      ...shown.map((track) => [
        track.name,
        String(track.spell_id),
        splitUnitName(track.target_name).name,
        (track.uptime_ms / 1000).toFixed(1),
        shareOf(rowOf([track])).toFixed(1),
        String(track.applications),
        track.appliers.map((guid) => splitUnitName(names.get(guid) ?? guid).name).join('; '),
      ]),
    ];
  }

  const SEPARATOR = ' · ';
</script>

{#snippet auraRows(list: AuraRow[], label: string, testid: string)}
  <StatTable {columns} nameHead="Aura" barHead="Uptime" {sort} {onsort} {label} {testid}>
    {#each list as row (row.key)}
      {@const track = row.tracks[0]}
      {@const each = row.tracks.length > 1 ? eachTarget(row) : []}
      <StatRow
        figures={[
          { value: percent(shareOf(row)), testid: 'aura-uptime' },
          { value: String(row.applications), testid: 'aura-applied' },
        ]}
        testid={row.tracks.length === 1
          ? `aura-${track.spell_id}-${track.target_guid}`
          : `aura-${track.spell_id}-${onEnemy(track) ? 'enemies' : 'players'}`}
      >
        {#snippet icon()}<SpellIcon
            socket
            url={icons.of(track.spell_id, track.name)}
            tip={track.spell_id}
          />{/snippet}
        {#snippet name()}
          <!-- The anchor a rating moment's link lands on (lib/rating/moments.ts); a row on several
               units carries one per unit, on that unit's own uptime in the sub line. -->
          <span
            id={row.tracks.length === 1 ? `aura-${track.target_guid}-${track.spell_id}` : undefined}
            class="scroll-mt-24">{track.name}</span
          >{#each each.slice(EACH_MAX) as other (other.target_guid)}<span
              id={`aura-${other.target_guid}-${other.spell_id}`}
              class="scroll-mt-24"
            ></span>{/each}{#if track.name.startsWith('Spell #')}<span
              class="text-muted ml-1 text-[13px] font-normal">unnamed in the log</span
            >{/if}{#if row.tracks.some((one) => ambiguous.has(`${one.target_guid}|${one.name}`))}<span
              class="text-muted ml-1 text-[13px] font-normal tabular-nums"
              title="Two spells share this name; this is spell id {track.spell_id}">#{track.spell_id}</span
            >{/if}
        {/snippet}
        {#snippet sub()}
          <span>on {onWhom(row)}</span
          >{#if showApplier(row)}{SEPARATOR}{#each applierNames(row).split(', ') as who, i (i)}{i > 0
                ? ', '
                : ''}<span>{i === 0 ? 'from ' : ''}{who}</span
              >{/each}{/if}{#if kind === 'DEBUFF' && schoolName(track.school)}{SEPARATOR}<span
              title="The spell's school. Whether a debuff can be dispelled is not in the log; the Dispels tab shows which ones were."
              data-testid="aura-school">{schoolName(track.school)}</span
            >{/if}<span class="ui-st-phone">{SEPARATOR}applied {row.applications}</span
          >{#if each.length > 0}<div data-testid="aura-each">
              {#each each.slice(0, EACH_MAX) as one, i (one.target_guid)}{i > 0 ? ' ' : ''}<span
                  id={`aura-${one.target_guid}-${one.spell_id}`}
                  class="scroll-mt-24 whitespace-nowrap"
                  >{i > 0 ? '· ' : ''}{unitLabel(row, one)}
                  <b class="text-text font-semibold tabular-nums">{percent(shareOf(rowOf([one])))}</b></span
                >{/each}{#if each.length > EACH_MAX}
                <span class="whitespace-nowrap">· {each.length - EACH_MAX} more</span>{/if}
            </div>{/if}
        {/snippet}
        {#snippet meter()}<Meter
            value={shareOf(row) / 100}
            tone={kind === 'BUFF' ? 'gold' : 'muted'}
          />{/snippet}
      </StatRow>
    {/each}
  </StatTable>
{/snippet}

<Section {title} {meta} testid="aura-section" panel flush>
  {#snippet head()}{#if rows.length > 0}<CopyCsv lines={csvLines} inHead />{/if}{/snippet}
  <div>
    {#if rows.length === 0}
      <EmptyState
        inset
        message={emptyMessage ?? `No ${kind === 'BUFF' ? 'buffs' : 'debuffs'} in this window.`}
        testid="table-empty"
      />
    {:else}
      {@render auraRows(rows, title, 'aura-table')}
    {/if}
  </div>
</Section>

{#if enemyRows.length > 0}
  <Section title="On the enemies" meta={enemyMeta} testid="aura-section-enemies" panel flush>
    <div>{@render auraRows(enemyRows, 'Debuffs on the enemies', 'aura-table-enemies')}</div>
  </Section>
{/if}

{#if bySpell.length > 0}
  <Section
    title={kind === 'DEBUFF' ? 'On the bosses, across the night' : 'Across the night'}
    meta={kind === 'DEBUFF' ? 'Uptime over the time the bosses were up' : 'Over everyone who had it'}
    testid="aura-by-spell"
    panel
  >
    <StatTable columns={[{ id: 'up', label: 'Up' }]} open label="Uptime across the night">
      {#each bySpell as entry (entry.name)}
        <StatRow figures={[{ value: percent(entry.pct) }]}>
          {#snippet icon()}<SpellIcon
              socket
              tip={entry.spellId}
              url={icons.of(entry.spellId, entry.name)}
            />{/snippet}
          {#snippet name()}{entry.name}{/snippet}
          {#snippet sub()}<span>on {entry.targets} {kind === 'DEBUFF' ? 'bosses' : 'players'}</span>{/snippet}
          {#snippet meter()}<Meter
              value={entry.pct / 100}
              tone={kind === 'BUFF' ? 'gold' : 'muted'}
            />{/snippet}
        </StatRow>
      {/each}
    </StatTable>
  </Section>
{/if}
