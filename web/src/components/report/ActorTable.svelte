<!-- web/src/components/report/ActorTable.svelte -->
<!-- The Damage done, Healing and Damage taken Plate (logs spec 2026-11-07 6.2): Section "{tab}"
     ("Click a player to see their abilities") over one StatTable, a row per actor in the scope:
     the crest, the name, the sub line, a class-colour Meter (share of the top), Amount (Effective
     on Healing, its sub adding the overheal), DPS / HPS / Taken / s, Share and Active; Parse leads
     on a ranked boss kill; a Total row foots it. Every figure head sorts. An empty window reads
     the tab's own line inside the Plate. -->
<script lang="ts">
  import EmptyLine from './EmptyLine.svelte';
  import type { Placement } from '../../lib/report/percentile';
  import { formatAmount, formatDuration, formatPercent } from '../../lib/report/format';
  import { formatRate } from '../../lib/report/frame';
  import type { Actor, RosterRow } from '../../lib/report/types';
  import type { StatColumn, StatFigure, StatSort } from '../../lib/ui/stat-table';
  import Section from '../ui/Section.svelte';
  import StatRow from '../ui/StatRow.svelte';
  import StatTable from '../ui/StatTable.svelte';
  import ActorRow from './ActorRow.svelte';
  import CopyCsv from './CopyCsv.svelte';

  type ActorTab = 'damage-done' | 'healing' | 'damage-taken';

  let {
    tab,
    actors,
    durationMs,
    roster,
    players,
    percentiles = new Map<string, Placement>(),
    parseFallback = '',
    showParse = false,
    amountApproximate = false,
    approximate = false,
    absent = [],
    windowIsWhole = true,
    deadAt = new Map<string, number>(),
    selected = '',
    playerHref,
    onSelectPlayer,
  }: {
    tab: ActorTab;
    actors: readonly Actor[];
    durationMs: number;
    /** The pull's roster, by GUID: the spec, class and item level of a row's sub line. */
    roster: ReadonlyMap<string, RosterRow>;
    /** The players' GUIDs: their names open their own view; a pet or an enemy's does not. */
    players: ReadonlySet<string>;
    percentiles?: ReadonlyMap<string, Placement>;
    parseFallback?: string;
    /** A ranked boss kill: the Parse column leads (never on Damage taken). */
    showParse?: boolean;
    /** The amounts are prorated: a window met by a target or boss filter. */
    amountApproximate?: boolean;
    /** The splits are prorated (a window or a filter): the mitigated line cannot be summed bare. */
    approximate?: boolean;
    /** Players with no row in this window, and when they died if they were dead. */
    absent?: readonly { name: string; deadSince: number | null }[];
    windowIsWhole?: boolean;
    /** Players dead at the window's end, by GUID, with when they died. */
    deadAt?: ReadonlyMap<string, number>;
    /** The bar's Source, when it is one player. */
    selected?: string;
    playerHref: (guid: string) => string;
    onSelectPlayer: (guid: string) => void;
  } = $props();

  const TITLE: Readonly<Record<ActorTab, string>> = {
    'damage-done': 'Damage done',
    healing: 'Healing',
    'damage-taken': 'Damage taken',
  };
  const RATE: Readonly<Record<ActorTab, string>> = {
    'damage-done': 'DPS',
    healing: 'HPS',
    'damage-taken': 'Taken / s',
  };
  const amountLabel = $derived(tab === 'healing' ? 'Effective' : 'Amount');
  const parse = $derived(showParse && tab !== 'damage-taken');
  const columns = $derived<StatColumn[]>([
    ...(parse ? [{ id: 'parse', label: 'Parse' }] : []),
    { id: 'amount', label: amountLabel, sortable: true },
    { id: 'rate', label: RATE[tab], sortable: true },
    { id: 'share', label: 'Share', sortable: true },
    { id: 'active', label: 'Active', sortable: true },
  ]);

  let sort = $state<StatSort>({ column: 'amount', direction: 'desc' });
  function onsort(column: string): void {
    sort =
      sort.column === column
        ? { column, direction: sort.direction === 'desc' ? 'asc' : 'desc' }
        : { column, direction: 'desc' };
  }
  const total = $derived(actors.reduce((sum, actor) => sum + actor.effective, 0));
  const top = $derived(actors.reduce((most, actor) => Math.max(most, actor.effective), 0));
  const timeOf = (actor: Actor): number => actor.time_ms ?? durationMs;
  const activeOf = (actor: Actor): number =>
    timeOf(actor) === 0 ? 0 : Math.min((actor.active_ms / timeOf(actor)) * 100, 100);
  function sortValue(actor: Actor, column: string): number {
    if (column === 'rate') return timeOf(actor) === 0 ? 0 : actor.effective / timeOf(actor);
    if (column === 'active') return activeOf(actor);
    return actor.effective;
  }
  const rows = $derived.by(() => {
    const sign = sort.direction === 'desc' ? -1 : 1;
    return [...actors].sort((a, b) => sign * (sortValue(a, sort.column) - sortValue(b, sort.column)));
  });
  const mark = $derived(amountApproximate ? '~' : '');

  function figures(actor: Actor): StatFigure[] {
    const placed = percentiles.get(actor.guid);
    return [
      ...(parse
        ? [
            placed === undefined
              ? { value: parseFallback || '–', muted: true }
              : { value: String(Math.floor(placed.percentile)) },
          ]
        : []),
      { value: `${mark}${formatAmount(actor.effective)}` },
      { value: formatRate(timeOf(actor) === 0 ? 0 : actor.effective / (timeOf(actor) / 1000)) },
      { value: total === 0 ? '0%' : formatPercent((actor.effective / total) * 100) },
      actor.active_unknown === true ? { value: '–', muted: true } : { value: formatPercent(activeOf(actor)) },
    ];
  }

  function subLine(actor: Actor): string {
    const row = roster.get(actor.guid);
    const who = [row?.spec, row?.class ?? actor.class].filter(Boolean).join(' ');
    const overheal =
      tab === 'healing' && actor.overheal !== undefined && actor.overheal > 0 && actor.total > 0
        ? `${formatAmount(actor.overheal)} overheal (${Math.round((actor.overheal / actor.total) * 100)}%)`
        : '';
    const dead = deadAt.get(actor.guid);
    return [
      who,
      row?.item_level ? `item level ${row.item_level}` : '',
      overheal,
      dead === undefined ? '' : `dead since ${formatDuration(dead)}`,
    ]
      .filter((part) => part !== '')
      .join(' · ');
  }

  const totalSeconds = $derived(Math.max(durationMs, 1) / 1000);
  const totalFigures = $derived<StatFigure[]>([
    ...(parse ? [{ value: '' }] : []),
    { value: `${mark}${formatAmount(total)}` },
    { value: formatRate(total / totalSeconds) },
    { value: '100%' },
    { value: '' },
  ]);
  /** Damage taken's mitigated line (tanks read it first): what was absorbed, blocked and avoided. */
  const mitigatedLine = $derived.by(() => {
    if (tab !== 'damage-taken') return '';
    let absorbed = 0;
    let blocked = 0;
    let prorated = false;
    // eslint-disable-next-line svelte/prefer-svelte-reactivity
    const avoided = new Map<string, number>();
    const addMisses = (misses: Readonly<Record<string, number>> | undefined): void => {
      for (const [type, count] of Object.entries(misses ?? {}))
        avoided.set(type, (avoided.get(type) ?? 0) + count);
    };
    for (const actor of actors) {
      // A measured row carries its own figures; otherwise the per-ability ones, prorated with them.
      if (actor.mitigated !== undefined) {
        absorbed += actor.mitigated.absorbed;
        blocked += actor.mitigated.blocked;
        addMisses(actor.mitigated.misses);
        continue;
      }
      prorated ||= approximate;
      for (const ability of actor.abilities) {
        absorbed += ability.absorbed ?? 0;
        blocked += ability.blocked ?? 0;
        addMisses(ability.misses);
      }
    }
    if (prorated)
      return 'Mitigated: not split under this filter over this window; open a pull to measure it.';
    const hits = [...avoided.values()].reduce((sum, count) => sum + count, 0);
    if (absorbed === 0 && blocked === 0 && hits === 0) return '';
    // Alphabetical: the same numbers read in one order under every filter.
    const byType = [...avoided.entries()]
      .sort((a, b) => a[0].localeCompare(b[0]))
      .map(([type, count]) => `${count} ${type.toLowerCase()}`)
      .join(', ');
    const parts = [
      ...(absorbed > 0 ? [`${formatAmount(absorbed)} absorbed`] : []),
      ...(blocked > 0 ? [`${formatAmount(blocked)} blocked`] : []),
      ...(hits > 0 ? [`${hits} hits avoided (${byType})`] : []),
    ];
    return `Mitigated: ${parts.join(' · ')}`;
  });

  /** The table as it stands, for a spreadsheet: name, share, amount, overheal, per second, active. */
  function csvLines(): string[][] {
    const healing = tab === 'healing';
    const lines = [
      [
        'Rank',
        'Name',
        'Share',
        'Amount',
        ...(healing ? ['Overheal', 'Overheal %'] : []),
        'Per second',
        'Active %',
      ],
    ];
    rows.forEach((actor, index) => {
      const overheal = actor.overheal ?? 0;
      const gross = actor.effective + overheal;
      const overhealCells =
        actor.overheal === undefined
          ? ['', '']
          : [String(overheal), gross === 0 ? '0' : ((overheal / gross) * 100).toFixed(1)];
      lines.push([
        String(index + 1),
        actor.name,
        total === 0 ? '0' : ((actor.effective / total) * 100).toFixed(2),
        String(actor.effective),
        ...(healing ? overhealCells : []),
        timeOf(actor) === 0 ? '0' : (actor.effective / (timeOf(actor) / 1000)).toFixed(1),
        actor.active_unknown === true ? '' : activeOf(actor).toFixed(1),
      ]);
    });
    return lines;
  }

  const absentLine = $derived(
    absent.length === 0
      ? ''
      : `No row of their own ${windowIsWhole ? 'under this filter' : 'in this window'}: ${absent
          .map((entry) =>
            entry.deadSince === null
              ? entry.name
              : `${entry.name} (dead since ${formatDuration(entry.deadSince)})`,
          )
          .join(', ')}.`,
  );
</script>

{#snippet totalName()}Total, {actors.length} {actors.length === 1 ? 'player' : 'players'}{/snippet}
{#snippet absentSub()}{#if mitigatedLine !== ''}<span class="block" data-testid="actor-mitigated"
      >{mitigatedLine}</span
    >{/if}{#if absentLine !== ''}<span class="block" data-testid="actor-absent">{absentLine}</span
    >{/if}{/snippet}

<Section title={TITLE[tab]} testid="report-table-card" panel frame="framed" flush>
  {#snippet head()}{#if actors.length > 0}<CopyCsv lines={csvLines} inHead />{/if}{/snippet}
  <div>
    {#if actors.length === 0}
      <EmptyLine text="Nothing in this window." testid="table-empty" />
    {:else}
      <StatTable {columns} nameHead="Player" {sort} {onsort} label={TITLE[tab]} testid="actor-table">
        {#each rows as actor (actor.guid)}
          <ActorRow
            {actor}
            subLine={subLine(actor)}
            figures={figures(actor)}
            ofTop={top === 0 ? 0 : actor.effective / top}
            phone={`${formatRate(timeOf(actor) === 0 ? 0 : actor.effective / (timeOf(actor) / 1000))} ${RATE[tab]}`}
            href={players.has(actor.guid) ? playerHref(actor.guid) : undefined}
            selected={selected === actor.guid}
            onSelect={onSelectPlayer}
          />
        {/each}
        <StatRow
          figures={totalFigures}
          total
          name={totalName}
          sub={absentLine === '' && mitigatedLine === '' ? undefined : absentSub}
          testid="actor-total"
        />
      </StatTable>
    {/if}
  </div>
</Section>
