<!-- web/src/components/report/ExchangeTable.svelte -->
<!-- Interrupts and Dispels are the same shape: someone did something to someone else's
     spell. One component, two tabs (logs spec 2026-11-07 section 6.2): the tab's one Plate, a
     StatTable of the spell's icon, `{name} · {spell}` in the class colour over `Interrupted
     {cast} on {target}` / `Removed {aura} on {target}`, and Count. What went through or ran its
     course follows as an open Section.

     summary.ExchangeRow carries no timestamp, so nothing in it can be cut to a brushed
     window: scopeSummary (window.ts) passes `interrupts` and `dispels` straight through,
     untouched. These counts are the whole fight's under every window, and this table says
     so outright rather than letting an unchanged count after a brush read as "nothing
     happened in this window" -- a mark and a note a reader could miss beat a raid leader
     concluding the wrong thing from a number that never moved. The `†` mark and the note
     below are unconditional: unlike CastTable's approximate marks, nothing here depends
     on whether the window happens to be whole. -->
<script lang="ts">
  import CharacterName from '../ui/CharacterName.svelte';
  import EmptyState from '../ui/EmptyState.svelte';
  import { splitUnitName } from '../../lib/characters';
  import { wholeFightMark } from '../../lib/report/format';
  import { SpellIcons } from '../../lib/report/spell-icons.svelte';
  import type { AuraTrack, CastRow, ExchangeRow } from '../../lib/report/types';
  import Section from '../ui/Section.svelte';
  import SpellIcon from '../ui/SpellIcon.svelte';
  import StatRow from '../ui/StatRow.svelte';
  import StatTable from '../ui/StatTable.svelte';
  import CopyCsv from './CopyCsv.svelte';

  let {
    rows,
    emptyText,
    casts = [],
    everyone = rows,
    auras = [],
    players = new Set<string>(),
    classOf = new Map<string, string>(),
    kind: kindProp,
    windowed = false,
  }: {
    rows: ExchangeRow[];
    emptyText: string;
    /** The fight's cast rows, for the "went through" count on the Interrupts tab. */
    casts?: CastRow[];
    /**
     * Every source's rows, unscoped: what went through or ran its course is counted
     * against the whole group's stops, whoever the table is scoped to.
     */
    everyone?: ExchangeRow[];
    /** The fight's aura tracks, for the "ran its course" count on the Dispels tab. */
    auras?: AuraTrack[];
    players?: ReadonlySet<string>;
    /** GUID to class, for the name's colour. */
    classOf?: ReadonlyMap<string, string>;
    /** Which tab this is; by default the rows' own kind, else Dispels when aura tracks came. */
    kind?: 'interrupt' | 'dispel';
    /** A brushed window: the counts stay the whole fight's, so they are marked and a line says so. */
    windowed?: boolean;
  } = $props();

  const kind = $derived(
    kindProp ??
      (rows[0]?.kind === 'dispel' || (rows.length === 0 && auras.length > 0) ? 'dispel' : 'interrupt'),
  );
  const title = $derived(kind === 'dispel' ? 'Dispels' : 'Interrupts');
  const verb = $derived(kind === 'dispel' ? 'Removed' : 'Interrupted');
  const icons = new SpellIcons();
  const SEPARATOR = ' · ';

  interface Uncured {
    spell_id: number;
    name: string;
    applied: number;
    dispelled: number;
    /** Dispels by the sources this table is scoped to, when that is not everyone. */
    own: number;
  }

  /**
   * For every debuff somebody dispelled at least once: how many times it landed on a
   * player against how many times it was dispelled. The rest ran its full course.
   */
  const uncured = $derived.by<Uncured[]>(() => {
    if (auras.length === 0) return [];
    // eslint-disable-next-line svelte/prefer-svelte-reactivity
    const bySpell = new Map<number, Uncured>();
    for (const row of everyone) {
      if (row.kind !== 'dispel') continue;
      const found = bySpell.get(row.extra_spell_id) ?? {
        spell_id: row.extra_spell_id,
        name: row.extra_spell_name,
        applied: 0,
        dispelled: 0,
        own: 0,
      };
      found.dispelled += row.count;
      bySpell.set(row.extra_spell_id, found);
    }
    for (const row of rows) {
      const found = bySpell.get(row.extra_spell_id);
      if (found !== undefined && row.kind === 'dispel') found.own += row.count;
    }
    for (const track of auras) {
      if (track.type !== 'DEBUFF' || !players.has(track.target_guid)) continue;
      const found = bySpell.get(track.spell_id);
      if (found !== undefined) found.applied += track.applications;
    }
    return [...bySpell.values()]
      .map((entry) => ({ ...entry, applied: Math.max(entry.applied, entry.dispelled) }))
      .filter((entry) => entry.applied > entry.dispelled)
      .sort((a, b) => b.applied - b.dispelled - (a.applied - a.dispelled));
  });

  interface Missed {
    spell_id: number;
    name: string;
    cast: number;
    stopped: number;
    /** Stops by the sources this table is scoped to, when that is not everyone. */
    own: number;
  }

  /**
   * For every enemy spell somebody interrupted at least once: how many times an enemy
   * cast it against how many times it was stopped. A spell nobody ever interrupted is not
   * listed, because the summary cannot tell an uninterruptible cast from an ignored one.
   */
  const missed = $derived.by<Missed[]>(() => {
    if (casts.length === 0) return [];
    // A plain Map: built once inside the derived, never read reactively by key.
    // eslint-disable-next-line svelte/prefer-svelte-reactivity
    const stopped = new Map<number, Missed>();
    for (const row of everyone) {
      if (row.kind !== 'interrupt') continue;
      const found = stopped.get(row.extra_spell_id) ?? {
        spell_id: row.extra_spell_id,
        name: row.extra_spell_name,
        cast: 0,
        stopped: 0,
        own: 0,
      };
      found.stopped += row.count;
      stopped.set(row.extra_spell_id, found);
    }
    for (const row of rows) {
      const found = stopped.get(row.extra_spell_id);
      if (found !== undefined && row.kind === 'interrupt') found.own += row.count;
    }
    for (const cast of casts) {
      if (players.has(cast.guid)) continue;
      const found = stopped.get(cast.spell_id);
      if (found !== undefined) found.cast += Math.max(cast.started, cast.succeeded);
    }
    return [...stopped.values()]
      .map((entry) => ({ ...entry, cast: Math.max(entry.cast, entry.stopped) }))
      .filter((entry) => entry.cast > entry.stopped)
      .sort((a, b) => b.cast - b.stopped - (a.cast - a.stopped));
  });

  /** Whether the table shows some sources only, so the group's counts name the scope's part. */
  const scopedToSome = $derived(everyone !== rows && everyone.length !== rows.length);
  const ordered = $derived(
    [...rows].sort((a, b) => b.count - a.count || a.source_name.localeCompare(b.source_name)),
  );
  /** The table as lines: who did it, with what, on whom, which spell, how often. */
  function csvLines(): string[][] {
    const credit = [
      ['By', 'With', 'On', 'Spell', 'Spell id', 'Count'],
      ...ordered.map((row) => [
        splitUnitName(row.source_name).name,
        row.spell_name,
        splitUnitName(row.target_name).name,
        row.extra_spell_name,
        String(row.extra_spell_id),
        String(row.count),
      ]),
    ];
    // The denominators under the table, as a second block: what went through or ran its
    // course is the half a roster is built from.
    const through = missed.map((entry) => [
      'Went through',
      entry.name,
      String(entry.spell_id),
      String(entry.cast),
      String(entry.stopped),
      ...(scopedToSome ? [String(entry.own)] : []),
      String(entry.cast - entry.stopped),
    ]);
    const course = uncured.map((entry) => [
      'Ran their course',
      entry.name,
      String(entry.spell_id),
      String(entry.applied),
      String(entry.dispelled),
      ...(scopedToSome ? [String(entry.own)] : []),
      String(entry.applied - entry.dispelled),
    ]);
    if (through.length === 0 && course.length === 0) return credit;
    return [
      ...credit,
      [],
      [
        'Block',
        'Spell',
        'Spell id',
        'Cast or landed',
        'Stopped or dispelled',
        ...(scopedToSome ? ['In this scope'] : []),
        'Went through or ran their course',
      ],
      ...through,
      ...course,
    ];
  }
  const mark = $derived(wholeFightMark(windowed));
  const total = $derived(rows.reduce((sum, row) => sum + row.count, 0));

  /** One line per source, most first: who did the interrupting or dispelling tonight. */
  const bySource = $derived.by(() => {
    // eslint-disable-next-line svelte/prefer-svelte-reactivity
    const totals = new Map<string, { name: string; count: number }>();
    for (const row of rows) {
      const found = totals.get(row.source_guid);
      totals.set(row.source_guid, { name: row.source_name, count: (found?.count ?? 0) + row.count });
    }
    return [...totals.values()].sort((a, b) => b.count - a.count || a.name.localeCompare(b.name));
  });
  const shortfall = $derived(
    kind === 'dispel'
      ? uncured.map((entry) => ({
          spellId: entry.spell_id,
          name: entry.name,
          line: `landed on the raid ${entry.applied} times · dispelled ${entry.dispelled}`,
          own: entry.own,
          left: entry.applied - entry.dispelled,
        }))
      : missed.map((entry) => ({
          spellId: entry.spell_id,
          name: entry.name,
          line: `cast ${entry.cast} · stopped ${entry.stopped}`,
          own: entry.own,
          left: entry.cast - entry.stopped,
        })),
  );
</script>

<Section {title} meta={total > 0 ? `${mark}${total}` : undefined} testid="exchange-section" panel flush>
  {#snippet head()}{#if ordered.length > 0}<CopyCsv lines={csvLines} inHead />{/if}{/snippet}
  <div>
    {#if ordered.length === 0}
      <EmptyState inset message={emptyText} testid="table-empty" />
    {:else}
      <StatTable
        columns={[{ id: 'count', label: 'Count' }]}
        nameHead="Who, with what"
        bar={false}
        label={title}
        testid="exchange-table"
      >
        {#each ordered as row (`${row.source_guid}-${row.target_guid}-${row.spell_id}-${row.extra_spell_id}`)}
          <StatRow figures={[{ value: `${mark}${row.count}` }]} bar={false} phone={`${row.count}×`}>
            {#snippet icon()}<SpellIcon
                socket
                url={icons.of(row.spell_id, row.spell_name)}
                tip={row.spell_id}
              />{/snippet}
            {#snippet name()}
              <!-- The anchor a rating moment's link lands on (lib/rating/moments.ts). -->
              <span id={`exchange-${row.source_guid}-${row.spell_id}`} class="scroll-mt-24"
                ><CharacterName
                  name={splitUnitName(row.source_name).name}
                  cls={classOf.get(row.source_guid)}
                  weight="inherit"
                />{SEPARATOR}{row.spell_name}</span
              >
            {/snippet}
            {#snippet sub()}{verb} {row.extra_spell_name} on {splitUnitName(row.target_name).name}{/snippet}
          </StatRow>
        {/each}
      </StatTable>
    {/if}
  </div>
  {#if ordered.length > 0 && (bySource.length > 1 || windowed)}
    <div class="text-muted flex flex-col gap-1 px-5 pt-3 pb-4 text-[13px] leading-[18px]">
      {#if bySource.length > 1}<p data-testid="exchange-by-source">
          {#each bySource as entry, index (entry.name)}{#if index > 0}{SEPARATOR}{/if}<span
              ><span class="text-text font-semibold">{splitUnitName(entry.name).name}</span>
              <span class="tabular-nums">{mark}{entry.count}</span></span
            >{/each}
        </p>{/if}
      {#if windowed}<p data-testid="exchange-wholefight-note">
          Count is marked {mark}: interrupts and dispels are the whole fight's totals, because the summary
          keeps no timestamp for them and a brushed window cannot cut them down.
        </p>{/if}
    </div>
  {/if}
</Section>

{#if shortfall.length > 0}
  <Section
    title={kind === 'dispel' ? 'Ran their course' : 'Went through'}
    meta={kind === 'dispel'
      ? 'Debuffs someone dispelled at least once'
      : 'Casts someone stopped at least once'}
    testid={kind === 'dispel' ? 'dispels-uncured' : 'interrupts-missed'}
    panel
  >
    <StatTable
      columns={[{ id: 'left', label: kind === 'dispel' ? 'Ran their course' : 'Went through' }]}
      bar={false}
      open
      label={kind === 'dispel' ? 'Ran their course' : 'Went through'}
    >
      {#each shortfall as entry (entry.spellId)}
        <StatRow figures={[{ value: String(entry.left) }]} bar={false}>
          {#snippet icon()}<SpellIcon
              socket
              tip={entry.spellId}
              url={icons.of(entry.spellId, entry.name)}
            />{/snippet}
          {#snippet name()}{entry.name}{/snippet}
          {#snippet sub()}{entry.line}{#if scopedToSome}<span>
                ({entry.own} in this scope)</span
              >{/if}{SEPARATOR}<span class="text-wipe"
              >{entry.left} {kind === 'dispel' ? 'ran their course' : 'went through'}</span
            >{/snippet}
        </StatRow>
      {/each}
    </StatTable>
  </Section>
{/if}
