<!-- web/src/components/report/AuraTable.svelte -->
<!-- Buffs and Debuffs (logs spec 2026-11-07 section 6.2): the tab's one Plate, a StatTable of
     the aura's icon, its name over `on {target} · from {applier}`, an uptime Meter (muted on
     Debuffs), Up and Applied. Over the night a second, open Section reads each spell's uptime
     across every target that had it.

     Every figure here is exact under any window, not scaled: window.ts's scopeAuraTrack clips
     each segment to the window's own bounds and recomputes `uptime_ms` and `applications`
     from those clipped segments, rather than taking the whole fight's numbers and scaling
     them by a ratio the way the per-ability and per-target splits are. So this table carries
     neither the `~` nor the `†` mark. The one field on AuraTrack that IS whole-fight,
     `max_stacks`, is not rendered here. -->
<script lang="ts">
  import { splitUnitName } from '../../lib/characters';
  import { formatPercent, schoolName } from '../../lib/report/format';
  import activeBuild from '../../data/active-build.json';
  import { loadSharedSpellNames, nameAuras, type SpellNames } from '../../lib/report/aura-names';
  import { SpellIcons } from '../../lib/report/spell-icons.svelte';
  import type { AuraTrack } from '../../lib/report/types';
  import type { StatColumn, StatSort } from '../../lib/ui/stat-table';
  import Meter from '../ui/Meter.svelte';
  import Plate from '../ui/Plate.svelte';
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

  const pct = (ms: number): number => (durationMs === 0 ? 0 : (ms / durationMs) * 100);
  /**
   * Uptime over the time the track's target was in: the figure the row shows. Clamped to
   * 100 here, once, as a last line of defence (2026-09-21 result-page review round 3, E8: a
   * combined sim result's aura rows read up to 100.9%), so a future source of the same shape
   * of bug reads as an honest 100%, never a number over it.
   */
  const shareOf = (track: AuraTrack): number =>
    Math.min(
      track.time_ms === undefined ? pct(track.uptime_ms) : (track.uptime_ms / track.time_ms) * 100,
      100,
    );

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
  const keyOf = (track: AuraTrack, column: AuraSortColumn): number =>
    column === 'up' ? shareOf(track) : track.applications;

  const rows = $derived(
    nameAuras(tracks, spellNames)
      .filter((track) => track.type === kind)
      .sort((a, b) => {
        const column = sort.column as AuraSortColumn;
        const order = keyOf(b, column) - keyOf(a, column);
        return (sort.direction === 'desc' ? order : -order) || a.name.localeCompare(b.name);
      }),
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
    for (const track of rows) {
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
    for (const track of rows) {
      const key = `${track.target_guid}|${track.name}`;
      seen.set(key, new Set([...(seen.get(key) ?? []), track.spell_id]));
    }
    return new Set([...seen.entries()].filter(([, ids]) => ids.size > 1).map(([key]) => key));
  });

  const applierNames = (track: AuraTrack): string =>
    track.appliers
      .map((guid) => splitUnitName(names.get(guid) ?? '').name)
      .filter(Boolean)
      .slice(0, 2)
      .join(', ') || 'an unnamed source';
  /** A buff the target gave itself says nothing new by naming its applier. */
  const showApplier = (track: AuraTrack): boolean =>
    track.appliers.length > 0 &&
    (kind === 'DEBUFF' || track.appliers.some((guid) => guid !== track.target_guid));

  const over = $derived((windowed ?? startMs > 0) ? 'this window' : 'the whole fight');
  const meta = $derived(`${rows.length} on players · uptime over ${over}`);

  /** The table as lines: one per aura on one target, uptime in seconds and as a share. */
  function csvLines(): string[][] {
    return [
      ['Aura', 'Spell id', 'On', 'Uptime s', 'Uptime %', 'Applied', 'From'],
      ...rows.map((track) => [
        track.name,
        String(track.spell_id),
        splitUnitName(track.target_name).name,
        (track.uptime_ms / 1000).toFixed(1),
        shareOf(track).toFixed(1),
        String(track.applications),
        track.appliers.map((guid) => splitUnitName(names.get(guid) ?? guid).name).join('; '),
      ]),
    ];
  }

  const SEPARATOR = ' · ';
</script>

<Section {title} {meta} testid="aura-section">
  <Plate label={title}>
    {#if rows.length === 0}
      <p class="text-muted px-4 py-4 text-[14px]" data-testid="table-empty">
        {emptyMessage ?? `No ${kind === 'BUFF' ? 'buffs' : 'debuffs'} in this window.`}
      </p>
    {:else}
      <StatTable {columns} nameHead="Aura" barHead="Uptime" {sort} {onsort} label={title} testid="aura-table">
        {#each rows as track (`${track.target_guid}-${track.spell_id}`)}
          <StatRow
            figures={[
              { value: formatPercent(shareOf(track)), testid: 'aura-uptime' },
              { value: String(track.applications), testid: 'aura-applied' },
            ]}
            testid={`aura-${track.spell_id}-${track.target_guid}`}
          >
            {#snippet icon()}<SpellIcon
                url={icons.of(track.spell_id)}
                spellId={track.spell_id}
                tip
              />{/snippet}
            {#snippet name()}
              <!-- The anchor a rating moment's link lands on (lib/rating/moments.ts). -->
              <span id={`aura-${track.target_guid}-${track.spell_id}`} class="scroll-mt-24">{track.name}</span
              >{#if track.name.startsWith('Spell #')}<span class="text-muted ml-1 text-[13px] font-normal"
                  >unnamed in the log</span
                >{/if}{#if ambiguous.has(`${track.target_guid}|${track.name}`)}<span
                  class="text-muted ml-1 text-[13px] font-normal tabular-nums"
                  title="Two spells share this name; this is spell id {track.spell_id}"
                  >#{track.spell_id}</span
                >{/if}
            {/snippet}
            {#snippet sub()}
              <span>on {splitUnitName(track.target_name).name}</span>{#if showApplier(track)}{SEPARATOR}<span
                  >from {applierNames(track)}</span
                >{/if}{#if kind === 'DEBUFF' && schoolName(track.school)}{SEPARATOR}<span
                  title="The spell's school. Whether a debuff can be dispelled is not in the log; the Dispels tab shows which ones were."
                  data-testid="aura-school">{schoolName(track.school)}</span
                >{/if}<span class="ui-st-phone">{SEPARATOR}applied {track.applications}</span>
            {/snippet}
            {#snippet meter()}<Meter
                value={shareOf(track) / 100}
                tone={kind === 'BUFF' ? 'gold' : 'muted'}
              />{/snippet}
          </StatRow>
        {/each}
      </StatTable>
    {/if}
  </Plate>
  {#if rows.length > 0}<CopyCsv lines={csvLines} />{/if}
</Section>

{#if bySpell.length > 0}
  <Section
    title={kind === 'DEBUFF' ? 'On the bosses, across the night' : 'Across the night'}
    meta={kind === 'DEBUFF' ? 'Uptime over the time the bosses were up' : 'Over everyone who had it'}
    testid="aura-by-spell"
  >
    <StatTable columns={[{ id: 'up', label: 'Up' }]} open label="Uptime across the night">
      {#each bySpell as entry (entry.name)}
        <StatRow figures={[{ value: formatPercent(entry.pct) }]}>
          {#snippet icon()}<SpellIcon url={icons.of(entry.spellId)} spellId={entry.spellId} />{/snippet}
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
