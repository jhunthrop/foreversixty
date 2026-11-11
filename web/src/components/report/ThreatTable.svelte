<!-- web/src/components/report/ThreatTable.svelte -->
<!-- Threat, with the model's own honesty on the page. The engine's BaseThreat reports
     itself incomplete until the per-class table lands, and a threat number nobody can
     trust is worse than one labelled untrustworthy.

     `row.threat` is a second, independent lie: window.ts's scopeThreat has no threat
     model in the browser to recompute from, so it scales each row by the same ratio the
     actor rows' damage-plus-healing split moved by. That is a `~`, the same mark
     ActorRow's per-ability and per-target splits carry, not the model-incompleteness note
     above -- the two are unrelated and both can apply to the same row at once. The
     per-target rows are scaled the same way and carry the same mark.

     The picker turns one question into two. "Every enemy" is the totals table: every
     unit's threat over the window, enemies folded by name. Pick an enemy and the table
     becomes that enemy's own ranking -- who it is looking at, and by how much -- which is
     the question a tank actually asks. It rides in the url's `target`, the same key the
     damage tabs' target filter uses, so one link carries one target. Over a whole night
     that target is the enemy's NAME rather than a GUID -- an add is a new GUID on every
     pull -- which is why url.ts reads `target` as any printable string and not as a GUID
     pattern. -->
<script lang="ts">
  import { splitUnitName } from '../../lib/characters';
  import { approximateMark, classColorVar } from '../../lib/report/format';
  import { classSlugOf } from '../../lib/report/planner-link';
  import { chartCopy } from '../../lib/chart/copy';
  import type { ThreatPlayer } from '../../lib/chart/threat';
  import type { ChartEvent, ChartPhase } from '../../lib/chart/types';
  import type { Actor, Taunt, TargetChange, ThreatPair, ThreatRow } from '../../lib/report/types';
  import type { TimeWindow } from '../../lib/report/window';
  import type { StatColumn, StatFigure } from '../../lib/ui/stat-table';
  import { compact, percent } from '../../lib/format/number';
  import Crest from '../ui/Crest.svelte';
  import Meter from '../ui/Meter.svelte';
  import Section from '../ui/Section.svelte';
  import Select from '../ui/Select.svelte';
  import StatRow from '../ui/StatRow.svelte';
  import StatTable from '../ui/StatTable.svelte';
  import CopyCsv from './CopyCsv.svelte';
  import ThreatChart from './ThreatChart.svelte';
  import ThreatRail from './ThreatRail.svelte';

  let {
    rows,
    pairs = [],
    everyonePairs = undefined,
    taunts = undefined,
    names = new Map<string, string>(),
    classOf = new Map<string, string>(),
    players = undefined,
    approximate = false,
    totalThreat = undefined,
    target = '',
    durationMs,
    window = { startMs: 0, endMs: durationMs },
    nightMode = false,
    sourceName = undefined,
    scopeNoun = 'pull',
    onPatch,
    onWindow,
    chart = undefined,
    rail = true,
  }: {
    rows: ThreatRow[];
    /** One player's threat on one enemy. Absent from summaries before engine 0.3.1. */
    pairs?: ThreatPair[];
    /** The same pairs before the source scope, so a scoped row still shares against everyone. */
    everyonePairs?: ThreatPair[];
    /** Undefined means the report was parsed before taunts were kept; [] means none. */
    taunts?: Taunt[];
    /** GUID to unit name, the units table being the canonical spelling of a name. */
    names?: ReadonlyMap<string, string>;
    /** The whole window's threat, so a source scope's rows still share against everyone. */
    totalThreat?: number;
    classOf?: Map<string, string>;
    /** The player GUIDs; a roster row without a class is still a player. Defaults to classOf's keys. */
    players?: ReadonlySet<string>;
    /** True when the window is brushed, so threat is a scaled share, not measured. */
    approximate?: boolean;
    /** The picked enemy, from the url's `target`; '' is every enemy. */
    target?: string;
    /** The one player the source scope narrows to, by name; undefined when it does not. */
    sourceName?: string;
    /** What the unbrushed scope is called in the empty taunt line: a pull or the night. */
    scopeNoun?: 'pull' | 'night';
    /** The whole fight's length, which a taunt's window link is clamped against. */
    durationMs: number;
    /** The page's window, so the chart's brush is the report's window. */
    window?: TimeWindow;
    /** True on the night, where there is no one clock and the chart is not drawn. */
    nightMode?: boolean;
    onPatch: (patch: { target?: string }) => void;
    onWindow: (window: TimeWindow | null) => void;
    /** What the estimated-threat chart needs beyond the table: undefined draws no chart. */
    chart?: {
      /** Whole-fight pairs, every player, so a line runs across the whole fight. */
      pairs: readonly ThreatPair[];
      players: ReadonlyMap<string, ThreatPlayer>;
      /** Whole-fight taunts, every player. */
      taunts: readonly Taunt[];
      targetChanges?: readonly TargetChange[];
      damageTaken: readonly Actor[];
      events: (enemyGuids: readonly string[]) => ChartEvent[];
      phases: readonly ChartPhase[];
      sourcePlayer?: string;
      youGuid?: string;
      chosen: readonly string[] | null;
      onChoose: (chosen: string[] | null) => void;
      brushShape: readonly number[];
      syncKey: string;
      onOpenDeath?: (deaths: readonly ChartEvent[]) => void;
    };
    /** Draws the Taunts list under the table; the frame passes false once ThreatRail is in its rail. */
    rail?: boolean;
  } = $props();

  /** A row of the table, whichever question it is answering. */
  interface ThreatLine {
    guid: string;
    name: string;
    threat: number;
    /** The threat built inside the window; undefined when the whole fight is on screen. */
    built?: number;
    /** True when threat and built came from the pair's own series rather than a ratio. */
    measured?: boolean;
  }

  /** One enemy, by name, and the players who built threat on it. */
  interface TargetGroup {
    /** The url value: the group's first GUID, which over a night is the enemy's name. */
    id: string;
    name: string;
    /** Every GUID folded under this name, so a link to any one of them picks the group. */
    guids: string[];
    lines: ThreatLine[];
    total: number;
  }

  /** The table as lines: each row's threat and its share of whatever the table totals. */
  function csvLines(): string[][] {
    const share = (line: ThreatLine): string[] =>
      ranked ? [(total === 0 ? 0 : (line.threat / total) * 100).toFixed(1)] : [];
    const amount = showBuilt ? 'Standing' : 'Threat';
    return [
      [
        ...(picked === undefined ? ['Unit', amount] : ['Player', `${amount} on ${picked.name}`]),
        ...(showBuilt ? ['Built in the window'] : []),
        ...(ranked ? ['Share %'] : []),
      ],
      ...lines.map((line) => [
        splitUnitName(line.name).name,
        String(Math.round(line.threat)),
        ...(showBuilt ? [String(Math.round(line.built ?? 0))] : []),
        ...(shares(line) ? share(line) : ranked ? [''] : []),
      ]),
    ];
  }
  /** The environment (a fall, a fire) has the null GUID and holds no threat. */
  const NULL_GUID = /^0+$/;
  const isPlayer = (guid: string): boolean => (players ?? new Set(classOf.keys())).has(guid);
  /**
   * One row per enemy name: six "General Kaal" units are one boss to the reader, the way
   * Damage Taken folds them; players keep their own rows.
   */
  const ordered = $derived.by(() => {
    // eslint-disable-next-line svelte/prefer-svelte-reactivity
    const byName = new Map<string, ThreatRow>();
    const out: ThreatRow[] = [];
    for (const row of rows) {
      if (NULL_GUID.test(row.guid)) continue;
      if (isPlayer(row.guid)) {
        out.push(row);
        continue;
      }
      const key = splitUnitName(row.name).name;
      const found = byName.get(key);
      if (found === undefined) {
        const merged = { ...row, name: key };
        byName.set(key, merged);
        out.push(merged);
      } else {
        const merged = {
          ...found,
          threat: found.threat + row.threat,
          complete: found.complete && row.complete,
        };
        byName.set(key, merged);
        out[out.indexOf(found)] = merged;
      }
    }
    return out.sort((a, b) => b.threat - a.threat);
  });
  /**
   * The picker's options: one per enemy name, biggest first, each folding every unit of
   * that name -- an add is a new GUID on every pull, and six of them are one enemy to the
   * reader. A player's threat on all of them adds up into one row.
   */
  function groupPairs(list: ThreatPair[]): TargetGroup[] {
    // eslint-disable-next-line svelte/prefer-svelte-reactivity
    const byName = new Map<string, { guids: string[]; players: Map<string, ThreatLine> }>();
    for (const pair of list) {
      // A row is a player's threat on an enemy: an enemy's own threat on another enemy (or
      // on itself) is not a row, and not in the denominator either.
      if (!isPlayer(pair.guid)) continue;
      const key = splitUnitName(pair.target_name).name;
      let found = byName.get(key);
      if (found === undefined) {
        found = { guids: [], players: new Map<string, ThreatLine>() };
        byName.set(key, found);
      }
      if (!found.guids.includes(pair.target_guid)) found.guids.push(pair.target_guid);
      const have = found.players.get(pair.guid);
      found.players.set(pair.guid, {
        guid: pair.guid,
        name: splitUnitName(pair.name).name,
        threat: (have?.threat ?? 0) + pair.threat,
        built: pair.built === undefined ? have?.built : (have?.built ?? 0) + pair.built,
        measured: pair.measured === true || have?.measured === true,
      });
    }
    const out: TargetGroup[] = [...byName.entries()].map(([name, group]) => {
      const lines = [...group.players.values()].sort((a, b) => b.threat - a.threat);
      return {
        id: group.guids[0],
        name,
        guids: group.guids,
        lines,
        total: lines.reduce((sum, line) => sum + line.threat, 0),
      };
    });
    return out.sort((a, b) => b.total - a.total);
  }
  const groups = $derived(groupPairs(pairs));
  /** The same fold over the unscoped pairs, computed once rather than per read. */
  const everyoneGroups = $derived(everyonePairs === undefined ? groups : groupPairs(everyonePairs));
  /** The url's target names one of this table's enemies, or it names none of them. */
  const picked = $derived(target === '' ? undefined : groups.find((group) => group.guids.includes(target)));
  /**
   * The picked enemy's threat from every player, not only the ones in scope -- the same
   * job `totalThreat` does for the totals table, so narrowing the source to one player
   * shows their real share of what the boss is looking at rather than 100% of themselves.
   */
  const pickedTotal = $derived(
    picked === undefined
      ? 0
      : (everyoneGroups.find((group) => group.name === picked.name)?.total ?? picked.total),
  );
  /**
   * Per player, standing and built across every enemy: the totals table's own measured
   * figures. An enemy's own threat row has no pair and keeps the summary's scaled total,
   * which is what it has always been and what the share already leaves out.
   */
  const measuredTotals = $derived.by(() => {
    // eslint-disable-next-line svelte/prefer-svelte-reactivity
    const out = new Map<string, { threat: number; built: number }>();
    for (const pair of pairs) {
      if (pair.measured !== true) continue;
      const found = out.get(pair.guid) ?? { threat: 0, built: 0 };
      out.set(pair.guid, {
        threat: found.threat + (pair.standing ?? pair.threat),
        built: found.built + (pair.built ?? 0),
      });
    }
    return out;
  });
  /** The totals table's lines, with a player's measured standing standing in for the scaled total. */
  const totalLines = $derived<ThreatLine[]>(
    ordered.map((row) => {
      const found = measuredTotals.get(row.guid);
      return found === undefined
        ? { guid: row.guid, name: row.name, threat: row.threat }
        : { guid: row.guid, name: row.name, threat: found.threat, built: found.built, measured: true };
    }),
  );
  /** The table on screen: the picked enemy's pairs, or every unit's totals. */
  const onScreen = $derived<ThreatLine[]>(picked?.lines ?? totalLines);
  /** Every line on screen is measured: the window's figures came from the series, not a ratio. */
  const allMeasured = $derived(onScreen.length > 0 && onScreen.every((line) => line.measured === true));
  /**
   * Under a brush with nothing measured the figures are scaled totals, not a standing, so
   * the rows keep the roster's order (by name) rather than an order that reads as a
   * ranking. Measured lines are a ranking and keep it.
   */
  const lines = $derived<ThreatLine[]>(
    approximate && !allMeasured
      ? [...onScreen].sort((a, b) => a.name.localeCompare(b.name))
      : [...onScreen].sort((a, b) => b.threat - a.threat),
  );
  /** Six units named "General Kaal" are six rows; each after the first says which copy it is. */
  const copyOf = $derived.by(() => {
    // eslint-disable-next-line svelte/prefer-svelte-reactivity
    const seen = new Map<string, number>();
    // eslint-disable-next-line svelte/prefer-svelte-reactivity
    const out = new Map<string, number>();
    for (const row of ordered) {
      const n = (seen.get(row.name) ?? 0) + 1;
      seen.set(row.name, n);
      out.set(row.guid, n);
    }
    return out;
  });
  const copies = $derived(new Set([...copyOf.entries()].filter(([, n]) => n > 1).map(([guid]) => guid)));
  const peak = $derived(lines.reduce((highest, line) => Math.max(highest, line.threat), 0));
  // Both branches share against everyone, never against the rows in scope: a picked enemy
  // against every player's threat on it, the totals table against every player's threat in
  // the window. The enemies' own rows are outside the share: their threat is the model run
  // over their damage, and a raid that reads 52% of itself because the boss holds the rest
  // is a raid misled.
  // A measured table shares against the players' measured standings: the page's
  // `totalThreat` is the prorated window total from the whole-fight rows, which a standing
  // (cumulative to the window's end) outgrows many times over, and dividing by it read as
  // shares past 100%.
  const playersTotal = $derived(
    lines.filter((line) => isPlayer(line.guid)).reduce((sum, line) => sum + line.threat, 0),
  );
  const total = $derived(
    picked !== undefined ? pickedTotal : allMeasured ? playersTotal : (totalThreat ?? playersTotal),
  );
  /** A row that has a share: every row on a picked enemy, only the players on the totals table. */
  const shares = (line: ThreatLine): boolean => picked !== undefined || isPlayer(line.guid);
  /**
   * Bars and shares draw a ranking. A measured window is a ranking, standing is exactly
   * the number that decides aggro, so they are drawn. A window with nothing measured
   * (a report parsed before engine 0.4.0) still holds them back.
   */
  const ranked = $derived(!approximate || allMeasured);
  /** The tilde is for a scaled figure only; a measured one never carries it. */
  const scaled = $derived(approximate && !allMeasured);
  const mark = $derived(approximateMark(scaled));
  /** The second figure exists only under a window: the whole fight has one number. */
  const showBuilt = $derived(allMeasured && approximate);
  /**
   * The enemies the chart draws: the picked one, or with "Every enemy" picked every enemy the
   * players built threat on, summed per player. The title says which (fix round 2): the field,
   * the title and the tooltip name the same thing.
   */
  const chartedGuids = $derived(
    picked !== undefined
      ? picked.guids
      : [
          ...new Set(
            (chart?.pairs ?? [])
              .filter((pair) => isPlayer(pair.guid) && !isPlayer(pair.target_guid))
              .map((pair) => pair.target_guid),
          ),
        ],
  );
  /** The threat chart draws once the pairs carry their per-second series (engine 0.4.0). */
  const showChart = $derived(
    !nightMode &&
      chart !== undefined &&
      chartedGuids.length > 0 &&
      chart.pairs.some((pair) => pair.series !== undefined),
  );
  const enemyOptions = $derived([
    { value: '', label: 'Every enemy' },
    // The group count is the row's meta, not a suffix on the name (§A10 row 4).
    ...groups.map((group) => ({
      value: group.id,
      label: group.name,
      meta: group.guids.length > 1 ? `×${group.guids.length}` : undefined,
    })),
  ]);
  const standingAtMs = $derived(window.endMs);
  /** The table's columns (spec 6.2): Threat (Standing under a measured window, then Built), Share
   *  and Of the top. A scaled window draws the figure alone: no bar, no shares, no ranking. */
  const columns = $derived<StatColumn[]>([
    { id: 'threat', label: showBuilt ? chartCopy.standingColumn : 'Threat' },
    ...(showBuilt ? [{ id: 'built', label: 'Built' }] : []),
    ...(ranked
      ? [
          { id: 'share', label: 'Share' },
          { id: 'top', label: 'Of the top' },
        ]
      : []),
  ]);
  const DASH = '–';
  const whole = (value: number): string => compact(Math.round(value));
  function figuresOf(line: ThreatLine): StatFigure[] {
    const share = total === 0 ? 0 : (line.threat / total) * 100;
    return [
      { value: `${mark}${whole(line.threat)}`, testid: 'threat-standing' },
      ...(showBuilt ? [{ value: whole(line.built ?? 0), testid: 'threat-built' }] : []),
      ...(ranked
        ? [
            shares(line)
              ? { value: percent(share), testid: 'threat-share' }
              : { value: DASH, muted: true, testid: 'threat-no-share' },
            { value: `${peak === 0 ? 0 : Math.round((line.threat / peak) * 100)}%` },
          ]
        : []),
    ];
  }
  /** Below 1024 the primary figure's line says what it is and carries the others. */
  function phoneOf(line: ThreatLine): string {
    const noun = showBuilt ? 'standing' : 'threat';
    const built = showBuilt ? ` · ${whole(line.built ?? 0)} built` : '';
    const share =
      ranked && shares(line) ? ` · ${percent(total === 0 ? 0 : (line.threat / total) * 100)}` : '';
    return `${noun}${built}${share}`;
  }
  const subOf = (line: ThreatLine): string =>
    isPlayer(line.guid) ? (classOf.get(line.guid) ?? 'Player') : 'Enemy';
</script>

{#snippet enemyPicker()}
  {#if groups.length > 0}
    <span class="inline-flex items-center gap-2 max-md:w-full">
      <label class="text-small text-muted max-md:sr-only" for="threat-target">{chartCopy.enemy}</label>
      <span class="w-[180px] max-md:flex-1">
        <Select
          id="threat-target"
          class="w-full"
          size="dense"
          align="end"
          testid="threat-target"
          options={enemyOptions}
          value={picked?.id ?? ''}
          onchange={(event) => onPatch({ target: (event.currentTarget as HTMLSelectElement).value })}
        />
      </span>
    </span>
  {/if}
{/snippet}

<div class="flex flex-col gap-[var(--section-gap,32px)]" data-testid="threat-table">
  <Section
    title={chartCopy.threatTableTitle}
    meta={picked === undefined || lines.length === 0
      ? undefined
      : chartCopy.threatTableMeta(picked.name, standingAtMs, lines.length)}
    testid="threat-table-card"
    panel
  >
    {#snippet head()}{#if lines.length > 0}<CopyCsv lines={csvLines} inHead />{/if}{/snippet}
    {#if !showChart}
      <div class="flex flex-col gap-2 pb-3">
        <p class="text-small text-muted" data-testid="threat-chart-note">
          {#if nightMode}
            A night has no clock to draw threat on, so there is no chart here: the totals below are every
            pull's threat added up.
          {:else}
            This report was parsed before threat was kept second by second, so there is no chart; the totals
            below are the whole fight's.
          {/if}
        </p>
        {@render enemyPicker()}
      </div>
    {/if}
    <div>
      {#if lines.length === 0}
        <p class="text-muted px-4 py-4 text-[14px]" data-testid="table-empty">No threat in this window.</p>
      {:else}
        <StatTable
          {columns}
          nameHead={picked === undefined ? 'Unit' : 'Player'}
          barHead={ranked ? chartCopy.threatColumn : ''}
          bar={ranked}
          label={chartCopy.threatTableTitle}
          testid={picked === undefined ? 'threat-totals' : 'threat-on-target'}
        >
          {#each lines as row (row.guid)}
            <StatRow figures={figuresOf(row)} phone={phoneOf(row)} bar={ranked} testid={`threat-${row.guid}`}>
              {#snippet icon()}
                <Crest
                  cls={isPlayer(row.guid) ? (classSlugOf(classOf.get(row.guid)) ?? '') : ''}
                  size={28}
                />
              {/snippet}
              {#snippet name()}
                <span style:color={classColorVar(classOf.get(row.guid))} data-testid="threat-name"
                  >{splitUnitName(row.name).name}</span
                >{#if picked === undefined && copies.has(row.guid)}<span
                    class="text-muted ml-1 text-[13px] font-normal tabular-nums"
                    title="Another unit of the same name">#{copyOf.get(row.guid)}</span
                  >{/if}
              {/snippet}
              {#snippet sub()}{subOf(row)}{/snippet}
              {#snippet meter()}
                <Meter
                  value={peak === 0 ? 0 : row.threat / peak}
                  tone={isPlayer(row.guid) ? 'class' : 'muted'}
                  color={isPlayer(row.guid) ? classColorVar(classOf.get(row.guid)) : undefined}
                />
              {/snippet}
            </StatRow>
          {/each}
        </StatTable>
      {/if}
    </div>
    {#if lines.length > 0}
      <div class="flex flex-col gap-1 pt-2">
        {#if ranked}
          <p class="text-small text-muted" data-testid="threat-share-note">
            Share is of every player’s threat {picked === undefined ? 'in this window' : `on ${picked.name}`},
            whatever the source scope shows; the bar is drawn against the highest row.
          </p>
        {/if}
        {#if showBuilt}
          <p class="text-small text-muted" data-testid="threat-window-note">
            Standing is the threat this player had built by the window’s end, which is the number that decides
            who the enemy attacks; Built is what they made inside the window. Both are measured from the
            fight’s own seconds, so neither is marked. The table sorts by Standing.
          </p>
        {:else if scaled}
          <p class="text-small text-muted" data-testid="threat-approximate-note">
            This report was parsed before threat was kept second by second, so threat inside a window is not
            measured. Each figure is marked {mark} because it is the whole fight’s threat scaled to this window’s
            share of that player’s total, not the window’s own events, so it is not a standing and is not drawn
            as one: no bars, no shares, no order to read. Parse the report again for the measured figures. The whole
            pull is exact.
          </p>
        {/if}
      </div>
    {/if}
  </Section>
  {#if showChart && chart !== undefined}
    <ThreatChart
      enemyName={picked?.name ?? null}
      enemyGuids={chartedGuids}
      pairs={chart.pairs}
      players={chart.players}
      taunts={chart.taunts}
      targetChanges={chart.targetChanges}
      damageTaken={chart.damageTaken}
      events={chart.events(chartedGuids)}
      phases={chart.phases}
      {durationMs}
      {window}
      {onWindow}
      sourcePlayer={chart.sourcePlayer}
      youGuid={chart.youGuid}
      chosen={chart.chosen}
      onChoose={chart.onChoose}
      brushShape={chart.brushShape}
      syncKey={chart.syncKey}
      onOpenDeath={chart.onOpenDeath}
      {enemyPicker}
    />
  {/if}
  {#if rail}
    <ThreatRail
      {taunts}
      enemy={picked?.name}
      {names}
      {classOf}
      {sourceName}
      {approximate}
      {scopeNoun}
      {durationMs}
      {onWindow}
    />
  {/if}
</div>
