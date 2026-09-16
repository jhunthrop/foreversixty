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
  import {
    approximateAriaLabel,
    approximateMark,
    approximateTitle,
    classColorVar,
    formatAmount,
    formatDuration,
    formatPercent,
    tauntKey,
  } from '../../lib/report/format';
  import type { Taunt, ThreatPair, ThreatRow } from '../../lib/report/types';
  import { aroundWindow, type TimeWindow } from '../../lib/report/window';
  import CopyCsv from './CopyCsv.svelte';
  import TimeChart from './TimeChart.svelte';

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
  /** Every line on screen is measured: the window's figures came from the series, not a ratio. */
  const allMeasured = $derived(lines.length > 0 && lines.every((line) => line.measured === true));
  /**
   * Under a brush with nothing measured the figures are scaled totals, not a standing, so
   * the rows keep the roster's order (by name) rather than an order that reads as a
   * ranking. Measured lines are a ranking and keep it.
   */
  const lines = $derived<ThreatLine[]>(
    approximate && !(picked?.lines ?? totalLines).every((line) => line.measured === true)
      ? [...(picked?.lines ?? totalLines)].sort((a, b) => a.name.localeCompare(b.name))
      : [...(picked?.lines ?? totalLines)].sort((a, b) => b.threat - a.threat),
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
  const total = $derived(
    picked !== undefined
      ? pickedTotal
      : (totalThreat ??
          lines.filter((line) => isPlayer(line.guid)).reduce((sum, line) => sum + line.threat, 0)),
  );
  /** A row that has a share: every row on a picked enemy, only the players on the totals table. */
  const shares = (line: ThreatLine): boolean => picked !== undefined || isPlayer(line.guid);
  const incomplete = $derived(ordered.some((row) => !row.complete));
  const modelVersion = $derived(ordered[0]?.model_version ?? '');
  /**
   * The empty taunt line names what emptied it: the picked enemy, the source scope, or
   * the brush. A whole pull with no taunt says so and blames no window.
   */
  const noTaunts = $derived(
    `No taunts${sourceName === undefined ? '' : ` by ${splitUnitName(sourceName).name}`}${
      picked === undefined ? '' : ` on ${picked.name}`
    } in this ${approximate ? 'window' : scopeNoun}.`,
  );
  /**
   * Bars and shares draw a ranking. A measured window is a ranking — standing is exactly
   * the number that decides aggro — so they are drawn. A window with nothing measured
   * (a report parsed before engine 0.4.0) still holds them back.
   */
  const ranked = $derived(!approximate || allMeasured);
  /** The tilde is for a scaled figure only; a measured one never carries it. */
  const scaled = $derived(approximate && !allMeasured);
  const mark = $derived(approximateMark(scaled));
  const title = $derived(approximateTitle(scaled));
  /** The second figure exists only under a window: the whole fight has one number. */
  const showBuilt = $derived(allMeasured && approximate);
  const rowGrid = $derived(
    showBuilt
      ? 'md:grid-cols-[minmax(120px,1.2fr)_minmax(0,3fr)_96px_96px_64px]'
      : ranked
        ? 'md:grid-cols-[minmax(120px,1.2fr)_minmax(0,3fr)_96px_64px]'
        : 'md:grid-cols-[minmax(120px,1.2fr)_96px]',
  );
  // With an enemy picked, its taunts alone: the list answers "who took this one off me".
  const orderedTaunts = $derived(
    [...(taunts ?? [])]
      .filter((taunt) => picked === undefined || splitUnitName(taunt.target_name).name === picked.name)
      .sort((a, b) => a.at_ms - b.at_ms),
  );
  /** The units table spells a name; the event's own copy of it is the fallback. */
  const nameOf = (guid: string, recorded: string): string => splitUnitName(names.get(guid) ?? recorded).name;
  /**
   * The enemy the chart draws: the picked one, or — with "Every enemy" picked — the one
   * carrying the most threat, named under the chart so nobody reads it as the raid's total.
   */
  const charted = $derived(picked ?? everyoneGroups[0] ?? groups[0]);
  /** A running total, so the line's value at a second is that player's standing then. */
  function cumulative(series: number[]): number[] {
    let running = 0;
    return series.map((value) => (running += value));
  }
  /**
   * One line per player on the charted enemy, each in their class colour, longest first so
   * the legend reads top-down like the table. Every pair of a player on this enemy is
   * summed: six adds of one name are one enemy, and a player's line on it is the sum.
   */
  const chartLines = $derived.by(() => {
    if (charted === undefined) return [];
    // eslint-disable-next-line svelte/prefer-svelte-reactivity
    const byPlayer = new Map<string, number[]>();
    for (const pair of pairs) {
      if (pair.series === undefined || !charted.guids.includes(pair.target_guid)) continue;
      const found = byPlayer.get(pair.guid) ?? [];
      const merged = [...found];
      pair.series.forEach((value, index) => {
        merged[index] = (merged[index] ?? 0) + value;
      });
      byPlayer.set(pair.guid, merged);
    }
    return [...byPlayer.entries()]
      .map(([guid, series]) => ({
        label: splitUnitName(names.get(guid) ?? guid).name,
        series: cumulative(series.map((value) => value ?? 0)),
        token: classColorVar(classOf.get(guid)),
      }))
      .sort((a, b) => (b.series[b.series.length - 1] ?? 0) - (a.series[a.series.length - 1] ?? 0));
  });
  /** The taunts on the charted enemy, as marks on the chart's time axis. */
  const chartMarks = $derived(
    orderedTaunts.map((taunt) => ({
      atMs: taunt.at_ms,
      label: `${nameOf(taunt.source_guid, taunt.source_name)} · ${taunt.spell_name} on ${nameOf(taunt.target_guid, taunt.target_name)}`,
    })),
  );
  const showChart = $derived(!nightMode && chartLines.length > 0);
  const selectClass =
    'border-line-warm bg-raised rounded-control text-text h-11 w-full max-w-full min-w-0 px-2 text-[13px] md:h-9 md:w-auto';
</script>

<div class="flex flex-col gap-2" data-testid="threat-table">
  {#if incomplete}
    <p class="text-muted text-[12px]" data-testid="threat-incomplete">
      Threat model {modelVersion} does not yet carry every class's modifiers: a tank's stance, taunt and threat
      multipliers are not in it, so a tank can read below the damage dealers they were holding threat over. The
      per-class table lands with Forever's ability data. Under a brushed window, threat is the fight's total scaled
      by the window's share, not the window's own events.
    </p>
  {/if}
  {#if showChart}
    <div data-testid="threat-chart">
      <TimeChart
        series={[]}
        extra={chartLines}
        marks={chartMarks}
        perSecond={false}
        {durationMs}
        {window}
        deaths={[]}
        label={`Threat on ${charted?.name ?? ''}`}
        {onWindow}
      />
    </div>
  {/if}
  <p class="text-muted text-[12px]" data-testid="threat-chart-note">
    {#if nightMode}
      A night has no clock to draw threat on, so there is no chart here: the totals below are every pull's
      threat added up.
    {:else if showChart}
      One cumulative line per player, on {charted?.name}{picked === undefined
        ? ', the enemy carrying the most threat; pick another above'
        : ''}. A taunt in the game puts the taunter on top, and the base threat model does not: the lines are
      what damage and healing built, so a taunt mark is the moment the order stopped matching them.
    {:else}
      This report was parsed before threat was kept second by second, so there is no chart; the totals below
      are the whole fight's.
    {/if}
  </p>
  {#if groups.length > 0}
    <label class="label text-muted flex flex-wrap items-center gap-2" for="threat-target">
      Enemy
      <select
        id="threat-target"
        class={selectClass}
        data-testid="threat-target"
        value={picked?.id ?? ''}
        onchange={(event) => onPatch({ target: (event.currentTarget as HTMLSelectElement).value })}
      >
        <option value="">Every enemy</option>
        {#each groups as group (group.id)}
          <option value={group.id}
            >{group.name}{group.guids.length > 1 ? ` ×${group.guids.length}` : ''}</option
          >
        {/each}
      </select>
    </label>
  {/if}
  {#if lines.length === 0}
    <p class="text-muted text-[14px]" data-testid="table-empty">No threat in this window.</p>
  {:else}
    <div class={`text-muted label hidden gap-x-3 px-2 pb-1 md:grid ${rowGrid}`}>
      <span>{picked === undefined ? 'Unit' : 'Player'}</span>
      {#if ranked}<span title="Threat generated, against the highest row">Threat</span>{/if}
      <span
        class="text-right"
        title={showBuilt
          ? 'Cumulative threat at the window’s end: the number that decides who the enemy attacks'
          : picked === undefined
            ? 'Threat accumulated from damage and healing over this window'
            : `Threat this player built on ${picked.name} over this window`}
        >{showBuilt ? 'Standing' : 'Total'}</span
      >
      {#if showBuilt}
        <span class="text-right" title="Threat this player built inside the window">Built</span>
      {/if}
      {#if ranked}
        <span
          class="text-right"
          title={picked === undefined
            ? "Share of every player's threat in this window; an enemy's own row is outside it"
            : `Share of every player's threat on ${picked.name}`}>Share</span
        >
      {/if}
    </div>
    <ul
      class="flex flex-col"
      data-testid={picked === undefined ? undefined : 'threat-on-target'}
      data-ranked={ranked ? 'true' : 'false'}
    >
      {#each lines as row (row.guid)}
        <li
          class={`border-line-soft grid min-h-11 grid-cols-[minmax(0,1fr)_auto] items-center gap-x-3 gap-y-1 border-b px-2 py-2 text-[14px] ${rowGrid}`}
          data-testid={`threat-${row.guid}`}
        >
          <span class="truncate font-semibold" style={`color: ${classColorVar(classOf.get(row.guid))}`}>
            {splitUnitName(row.name).name}{#if picked === undefined && copies.has(row.guid)}
              <span class="text-muted ml-1 font-mono text-[11px]" title="Another unit of the same name"
                >#{copyOf.get(row.guid)}</span
              >{/if}
          </span>
          {#if ranked}
            <span class="bg-line-soft col-span-2 block h-[6px] w-full md:col-span-1">
              <span
                class="bg-gold block h-full"
                style={`width: ${peak === 0 ? 0 : (row.threat / peak) * 100}%`}
              ></span>
            </span>
          {/if}
          <!-- The phone has no header row, so the figure says what it is. -->
          <span
            class={`tabular text-right font-mono ${ranked ? '' : 'text-muted'}`}
            {title}
            data-testid="threat-standing"
            aria-label={approximateAriaLabel(scaled, formatAmount(Math.round(row.threat)))}
          >
            {mark}{formatAmount(Math.round(row.threat))}<span class="label font-body ml-1.5 md:hidden"
              >{showBuilt ? 'standing' : 'threat'}</span
            >
          </span>
          {#if showBuilt}
            <span
              class="text-muted tabular text-right font-mono"
              title="Threat this player built inside the window"
              data-testid="threat-built"
            >
              {formatAmount(Math.round(row.built ?? 0))}<span class="label font-body ml-1.5 md:hidden"
                >built</span
              >
            </span>
          {/if}
          {#if ranked && shares(row)}
            <span
              class="text-muted tabular col-span-2 text-right font-mono text-[13px] md:col-span-1"
              data-testid="threat-share"
              >{formatPercent(total === 0 ? 0 : (row.threat / total) * 100)}<span
                class="label font-body ml-1.5 md:hidden">share</span
              ></span
            >
          {:else if ranked}
            <span
              class="text-muted col-span-2 text-right text-[12px] md:col-span-1"
              title="An enemy's threat is the model run over its own damage; the raid's shares do not include it"
              data-testid="threat-no-share">—</span
            >
          {/if}
        </li>
      {/each}
    </ul>
    {#if ranked}
      <p class="text-muted text-[12px]" data-testid="threat-share-note">
        Share is of every player’s threat {picked === undefined ? 'in this window' : `on ${picked.name}`},
        whatever the source scope shows; the bar is drawn against the highest row.
      </p>
    {/if}
    <CopyCsv lines={csvLines} />
    {#if showBuilt}
      <p class="text-muted text-[12px]" data-testid="threat-window-note">
        Standing is the threat this player had built by the window’s end, which is the number that decides who
        the enemy attacks; Built is what they made inside the window. Both are measured from the fight’s own
        seconds, so neither is marked. The table sorts by Standing.
      </p>
    {:else if scaled}
      <p class="text-muted text-[12px]" data-testid="threat-approximate-note">
        This report was parsed before threat was kept second by second, so threat inside a window is not
        measured. Each figure is marked {mark} because it is the whole fight’s threat scaled to this window’s share
        of that player’s total, not the window’s own events, so it is not a standing and is not drawn as one: no
        bars, no shares, no order to read. Parse the report again for the measured figures. The whole pull is exact.
      </p>
    {/if}
  {/if}
  <div class="flex flex-col gap-1 pt-2" data-testid="threat-taunts">
    <h3 class="label text-muted">Taunts</h3>
    {#if taunts === undefined}
      <p class="text-muted text-[13px]">This report was parsed before taunts were kept; parse it again.</p>
    {:else if orderedTaunts.length === 0}
      <p class="text-muted text-[13px]" data-testid="threat-taunts-empty">{noTaunts}</p>
    {:else}
      <ul class="flex flex-col">
        {#each orderedTaunts as taunt (tauntKey(taunt))}
          <li
            class="border-line-soft flex min-h-11 flex-wrap items-center gap-x-2 gap-y-1 border-b px-2 py-2 text-[14px]"
          >
            <span
              class="text-muted tabular font-mono text-[13px]"
              title="On the pull's clock, whatever the window">{formatDuration(taunt.at_ms)}</span
            >
            <span aria-hidden="true" class="text-muted">·</span>
            <span>
              <span class="font-semibold" style={`color: ${classColorVar(classOf.get(taunt.source_guid))}`}
                >{nameOf(taunt.source_guid, taunt.source_name)}</span
              >
              taunted {nameOf(taunt.target_guid, taunt.target_name)}
              <span class="text-muted"
                >({taunt.spell_name}{taunt.pre_pull
                  ? ', cast before the pull, so Casts does not count it'
                  : ''})</span
              >
            </span>
            {#if taunt.label !== undefined}
              <span class="text-muted text-[12px]">{taunt.label}</span>
            {/if}
            <!-- The night has no window to set: a taunt there names its pull instead, and
                 the link that did nothing on the night is not offered. -->
            {#if scopeNoun !== 'night'}
              <button
                type="button"
                class="text-nav ml-auto inline-flex min-h-11 items-center text-[12px] font-bold tracking-[0.06em] uppercase md:min-h-0"
                data-testid="taunt-window"
                title="Set the window to the span around this taunt"
                onclick={() => onWindow(aroundWindow(taunt.at_ms, durationMs))}>Around it</button
              >
            {/if}
          </li>
        {/each}
      </ul>
    {/if}
  </div>
</div>
