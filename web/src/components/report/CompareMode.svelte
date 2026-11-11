<!-- web/src/components/report/CompareMode.svelte -->
<!-- Two fights, one table, one row per player: what they did in each and the difference.
     This is the thing a raid leader does all night -- this pull against the one before --
     and doing it by opening two tabs is what everyone does today instead.

     Compare asks two different questions, and both expand from the same per-player row.
     "What did the player above me do differently" is one player's abilities in two pulls,
     which is a row expanding. "What did I do differently from them" is two players in this
     pull, which is the second picker beside "Compare with" -- picking a player swaps the
     whole panel from two fights to two players, rather than adding a third column, since a
     reader asking one question is not usually asking the other at the same time.

     Both sides are always the whole fight. The chart above can brush a time window, and
     that window scopes Analyze's tables (window.ts's scopeSummary) to a slice of the
     fight that is currently selected -- but a millisecond range from one fight's start
     has no defined meaning against a second fight of a different length, and there is no
     way to show a scaled figure from one side next to an exact figure from the other
     without the reader having to guess which is which. So Compare ignores the window
     entirely rather than inventing a per-fight rescoping, and says so, so a window left
     brushed on the chart is not mistaken for narrowing this table too. -->
<script lang="ts">
  import {
    COMPARE_METRICS,
    METRIC_LABELS,
    PER_SECOND_METRICS,
    abilityDiff,
    metricTable,
    phaseWindow,
    playerAbilityDiff,
    sharedPhases,
    sideScaled,
    type AbilityDiff,
    type CompareMetric,
  } from '../../lib/report/compare';
  import { splitUnitName } from '../../lib/characters';
  import { classColorVar, joinFacts, outcomeLabel } from '../../lib/report/format';
  import { fetchSummary } from '../../lib/report/load';
  import type { FightEntry, RosterRow, Summary } from '../../lib/report/types';
  import { clampWindow, scopeSummary, type TimeWindow } from '../../lib/report/window';
  import { reportCopy } from '../../lib/report/copy';
  import { selectCopy } from '../../lib/ui/copy';
  import type { SelectOption } from '../../lib/ui/listbox-model';
  import { classOption } from '../../lib/ui/select-options';
  import { classSlugOf } from '../../lib/report/planner-link';
  import type { StatColumn, StatFigure } from '../../lib/ui/stat-table';
  import { clockDuration } from '../../lib/format/duration';
  import { compact } from '../../lib/format/number';
  import ClassCrestRing from '../character/ClassCrestRing.svelte';
  import LoadError from '../ui/LoadError.svelte';
  import Notice from '../ui/Notice.svelte';
  import Section from '../ui/Section.svelte';
  import Select from '../ui/Select.svelte';
  import StatRow from '../ui/StatRow.svelte';
  import StatTable from '../ui/StatTable.svelte';
  import CopyCsv from './CopyCsv.svelte';

  let {
    fights,
    current,
    dataBaseUrl,
    engineVersion = '',
    left: leftWhole,
    window = null,
    rightIndex = null,
    metric: metricParam = '',
    vs = '',
    source = '',
    controls = true,
    onPatch,
  }: {
    fights: FightEntry[];
    current: number;
    dataBaseUrl: string;
    /** The engine that wrote the files, on the summary url so a re-parse is a cache miss. */
    engineVersion?: string;
    left: Summary;
    /** The Analyze window, applied to both sides; null compares whole fights. */
    window?: TimeWindow | null;
    /** The second fight, from the url; null until one is picked. */
    rightIndex?: number | null;
    /** The metric id from the url; '' means the default. */
    metric?: string;
    /** The compared player's GUID from the url; '' compares two fights. */
    vs?: string;
    /** The page's source scope: a player's GUID when one is picked, else a scope word. */
    source?: string;
    /** Draws the Compare with, Or a player, Metric and Phase pickers above the table; the page
     *  turns it off once its ControlBar carries them (logs spec 2026-11-07 section 6.1). */
    controls?: boolean;
    onPatch: (patch: { compareWith?: number | null; compareMetric?: string; compareVs?: string }) => void;
  } = $props();

  const options = $derived(fights.filter((fight) => fight.index !== current));
  /** "pull 2 of 3" per boss pull, the way the fight list says it, so the picker reads the same. */
  const pullOf = $derived.by(() => {
    // eslint-disable-next-line svelte/prefer-svelte-reactivity
    const seen = new Map<string, number>();
    // eslint-disable-next-line svelte/prefer-svelte-reactivity
    const total = new Map<string, number>();
    for (const fight of fights)
      if (fight.kind === 'encounter') total.set(fight.name, (total.get(fight.name) ?? 0) + 1);
    // eslint-disable-next-line svelte/prefer-svelte-reactivity
    const out = new Map<number, string>();
    for (const fight of fights) {
      if (fight.kind !== 'encounter') continue;
      const n = (seen.get(fight.name) ?? 0) + 1;
      seen.set(fight.name, n);
      const of = total.get(fight.name) ?? 1;
      out.set(fight.index, of > 1 ? `pull ${n} of ${of}` : '');
    }
    return out;
  });
  const bossOptions = $derived(options.filter((fight) => fight.kind === 'encounter'));
  const trashOptions = $derived(options.filter((fight) => fight.kind !== 'encounter'));
  /** A fight as the Compare with list reads it: the name (and which pull), and under it the
   *  outcome and length, the way FightSelector shows them (§A10 row 8). */
  function fightChoice(fight: FightEntry, group: string): SelectOption {
    const pull = pullOf.get(fight.index) ?? '';
    return {
      value: String(fight.index),
      label: `${fight.name}${pull === '' ? '' : ` · ${pull}`}`,
      description: joinFacts(outcomeLabel(fight), clockDuration(fight.duration_ms)),
      group,
    };
  }
  const compareChoices = $derived([
    ...bossOptions.map((fight) => fightChoice(fight, reportCopy.bossPulls)),
    ...trashOptions.map((fight) => fightChoice(fight, reportCopy.trash)),
  ]);
  let rightWhole = $state<Summary | null>(null);

  /**
   * The phase both sides are read over, when one is picked: each side is scoped to its
   * own span for it, so a pull whose Phase 2 ran twice as long is still compared phase
   * against phase. Not in the url: it is a way of reading the two fights on screen, and
   * a link already carries which two they are.
   */
  let phase = $state('');
  const phases = $derived(sharedPhases(leftWhole, rightWhole));
  // A second fight picked while a phase is chosen may not have that phase.
  $effect(() => {
    if (phase !== '' && !phases.includes(phase)) phase = '';
  });

  /** Each side in its own phase's span when one is picked, else in the page's window. */
  const leftWindow = $derived(phase === '' ? window : phaseWindow(leftWhole, phase));
  const rightWindow = $derived(phase === '' ? window : phaseWindow(rightWhole, phase));
  /** Each side in its own window, clamped to that fight's own length. */
  const left = $derived(
    leftWindow === null ? leftWhole : scopeSummary(leftWhole, clampWindow(leftWindow, leftWhole.duration_ms)),
  );
  const right = $derived(
    rightWhole === null || rightWindow === null
      ? rightWhole
      : scopeSummary(rightWhole, clampWindow(rightWindow, rightWhole.duration_ms)),
  );
  let error = $state('');
  /** Bumped by Try again: the effect below reads it, so the same fight is asked for again. */
  let attempt = $state(0);
  const metric = $derived<CompareMetric>(
    (COMPARE_METRICS as readonly string[]).includes(metricParam) ? (metricParam as CompareMetric) : 'dps',
  );
  /** The metric select's option text: METRIC_LABELS with a capital first letter, since the
   *  caption below it reads mid-sentence and keeps the label's own lowercase form. */
  function metricOptionLabel(id: CompareMetric): string {
    const label = METRIC_LABELS[id];
    return label.charAt(0).toUpperCase() + label.slice(1);
  }

  // A fight picked elsewhere on the page (the fight selector, or the browser's own back
  // button) can leave `rightIndex` pointing at the fight that just became `current`.
  // Comparing a fight to itself is never useful, so the selection is dropped rather than
  // shown as a table of zero differences.
  $effect(() => {
    if (rightIndex === current) onPatch({ compareWith: null });
  });

  /**
   * Picking a second fight does not cancel the request for the one picked before it, so
   * two of these can be in flight at once and settle in either order -- the same hazard
   * ReportView.svelte's own loadFight guards against with `index === state.fight`.
   * `wanted === rightIndex` is the whole guard here, checked on both the success and the
   * rejection path: an answer for a fight nobody has selected any more is dropped rather
   * than painted under the current selection's name, and a stale failure does not blank a
   * table that has since loaded correctly.
   */
  $effect(() => {
    const wanted = rightIndex;
    void attempt;
    if (wanted === null) {
      rightWhole = null;
      error = '';
      return;
    }
    void fetchSummary(dataBaseUrl, wanted, engineVersion)
      .then((summary) => {
        if (wanted !== rightIndex) return;
        rightWhole = summary;
        error = '';
      })
      .catch((thrown: unknown) => {
        if (wanted !== rightIndex) return;
        rightWhole = null;
        error = thrown instanceof Error ? thrown.message : 'That fight did not load.';
      });
  });

  /**
   * Threat inside a window is the whole fight's total scaled, not measured (the Threat tab
   * greys it and says so); Compare marks it the same way rather than printing it as fact.
   * "Inside a window" is each side's own effective window -- the page's shared one, or
   * (Task 21) a picked phase's own span for that side -- so a phase-scoped threat figure
   * is marked exactly like a window-scoped one, never printed as fact either.
   */
  const threatScaled = $derived(
    (sideScaled(leftWhole, leftWindow) || sideScaled(rightWhole, rightWindow)) &&
      (metric === 'threat' || metric === 'tps'),
  );
  const mark = $derived(threatScaled ? '~' : '');
  /** A difference as sign, mark, magnitude: "+~13,550" and "−~26,473" read the same way round. */
  const signed = (value: number): string => `${value >= 0 ? '+' : '−'}${mark}${compact(Math.abs(value))}`;
  /** A player's figure for the picked metric: off the roster row, or the threat table. */
  function rowMetric(row: RosterRow, summary: Summary): number {
    // A window past the shorter pull's end scopes that pull to nothing: no seconds and
    // no rate, for any per-second metric, rather than a total over a clamped-to-zero length.
    const scopedOut = summary.duration_ms <= 0;
    if (scopedOut) return 0;
    if (metric === 'threat' || metric === 'tps') {
      const threat = summary.threat.find((line) => line.guid === row.guid)?.threat ?? 0;
      return Math.round(metric === 'tps' ? threat / (summary.duration_ms / 1000) : threat);
    }
    return PER_SECOND_METRICS.has(metric) ? Math.round(row[metric]) : row[metric];
  }

  function fightLabel(fight: FightEntry | null, own: TimeWindow | null): string {
    if (fight === null) return '';
    const stretch =
      own === null
        ? clockDuration(fight.duration_ms)
        : `${clockDuration(own.startMs)} to ${clockDuration(Math.min(own.endMs, fight.duration_ms))} of ${clockDuration(fight.duration_ms)}`;
    return joinFacts(fight.name, stretch, outcomeLabel(fight).toLowerCase());
  }

  interface Line {
    guid: string;
    name: string;
    class?: string;
    a: number;
    b: number;
    /** False when the player is not in the other pull: the figure and difference read a dash. */
    inB: boolean;
  }

  /**
   * Every player in this pull, this metric, against the second fight's if one is picked --
   * `b` is 0 until `right` loads, the same way a roster member who sat out one side reads
   * 0 rather than being dropped. This still runs with no second fight picked, because the
   * versus-player panel below reads its own top row (and the picker's options) off this
   * same list rather than a second one.
   */
  const lines = $derived.by<Line[]>(() => {
    // A plain Map, not SvelteMap: this is a throwaway local built up once and then
    // turned into the plain array `$derived.by` returns, never read key-by-key outside
    // this function -- the same reasoning ReportView.svelte's own `summaries` cache and
    // `percentiles` builder give for their matching eslint-disable.
    // eslint-disable-next-line svelte/prefer-svelte-reactivity
    const byGuid = new Map<string, Line>();
    for (const row of left.roster) {
      byGuid.set(row.guid, {
        guid: row.guid,
        name: row.name,
        class: row.class,
        a: rowMetric(row, left),
        b: 0,
        inB: false,
      });
    }
    if (right !== null) {
      for (const row of right.roster) {
        const found = byGuid.get(row.guid);
        const value = rowMetric(row, right);
        if (found === undefined) {
          byGuid.set(row.guid, {
            guid: row.guid,
            name: row.name,
            class: row.class,
            a: 0,
            b: value,
            inB: true,
          });
        } else {
          byGuid.set(row.guid, { ...found, b: value, inB: true });
        }
      }
    }
    return [...byGuid.values()].sort((x, y) => y.a - y.b - (x.a - x.b));
  });

  /**
   * The player the picked one is compared against. With a player picked the question is
   * "what did I do differently from them", and the "I" is the player the page is already
   * scoped to: Source names them, so Source is the base whenever it names a player. With
   * no player scope it falls back to the metric's leader, the row a reader lands on. Either
   * way it is never the picked player, so the panel always holds two different people.
   */
  const basePlayer = $derived.by(() => {
    if (source !== '' && source !== vs && left.roster.some((row) => row.guid === source)) return source;
    // The metric's leader in this pull: the largest own figure, not the first row of the
    // table, which is sorted by the difference to the other pull when one is picked.
    const leader = lines
      .filter((line) => line.guid !== vs)
      .reduce<(typeof lines)[number] | undefined>(
        (best, line) => (best === undefined || line.a > best.a ? line : best),
        undefined,
      );
    return leader?.guid ?? lines[0]?.guid ?? '';
  });
  /** The players this pull's roster offers as a second player, the picked one excluded. */
  const vsChoices = $derived([
    { value: '', label: reportCopy.compareFightsInstead },
    ...left.roster
      .filter((row) => row.guid !== basePlayer)
      .map((row) =>
        classOption({ value: row.guid, label: splitUnitName(row.name).name, className: row.class }),
      )
      .sort((a, b) => a.label.localeCompare(b.label)),
  ]);
  /** True while the mode is two players inside this pull rather than two pulls. */
  const versusPlayer = $derived(vs !== '' && vs !== basePlayer && left.roster.some((row) => row.guid === vs));
  /** The name a card and the scope line call a player by. */
  function playerName(guid: string): string {
    return splitUnitName(left.roster.find((row) => row.guid === guid)?.name ?? '').name;
  }
  /**
   * The player table. Comparing two pulls it is every player, each against their own
   * figure in the other pull. Comparing two players it is those two, each against the
   * other's figure in this same pull (and window): the question is what one did that the
   * other did not, so a card's "compared with" is the other player, never a second pull
   * that is not part of this question at all.
   */
  const shownLines = $derived.by<Line[]>(() => {
    if (!versusPlayer) return lines;
    const base = lines.find((line) => line.guid === basePlayer);
    const other = lines.find((line) => line.guid === vs);
    if (base === undefined || other === undefined) return [];
    return [
      { ...base, b: other.a, inB: true },
      { ...other, b: base.a, inB: true },
    ];
  });
  /** The ability rows the reader opened, by GUID: a row is opened, not a tab. */
  let openRows = $state<string[]>([]);
  function toggleRow(guid: string): void {
    openRows = openRows.includes(guid) ? openRows.filter((entry) => entry !== guid) : [...openRows, guid];
  }
  const splitTable = $derived(metricTable(metric));
  /** Two players inside this pull, ability by ability. */
  const versusRows = $derived(versusPlayer ? playerAbilityDiff(left, basePlayer, vs, metric) : []);
  function rowsFor(guid: string): AbilityDiff[] {
    return abilityDiff(left, right, guid, metric);
  }

  /**
   * An ability split inside a window is prorated: window.ts scales each ability by the
   * window's share of its actor's total, because the summary keeps no per-ability series.
   * So the ability table's own figures carry the tilde whenever a side has an effective
   * window -- the page's shared one, or a picked phase's own span for that side -- whatever
   * the metric -- which is a different mark from `threatScaled`, the whole-table one that
   * only applies to the threat metrics.
   */
  const splitScaled = $derived(sideScaled(leftWhole, leftWindow) || sideScaled(rightWhole, rightWindow));
  const splitMark = $derived(splitScaled ? '~' : '');
  /** A signed difference in the ability table, marked when the split it came from was scaled. */
  const splitSigned = (value: number): string =>
    `${value >= 0 ? '+' : '−'}${splitMark}${compact(Math.abs(value))}`;

  /** An ability diff as lines, in the same shape as every other table's CSV. */
  function abilityCsv(rows: AbilityDiff[], aHead: string, bHead: string): string[][] {
    return [
      ['Ability', 'Spell id', 'Via', aHead, bHead, 'Difference'],
      ...rows.map((row) => [
        row.name,
        row.id === undefined ? '' : String(row.id),
        row.via ?? '',
        row.a === null ? '' : String(row.a),
        row.b === null ? '' : String(row.b),
        String(row.delta),
      ]),
    ];
  }
  /** The player table as lines, for the CSV of what is on screen. */
  function playerCsv(): string[][] {
    return [
      ['Player', 'This fight', 'Compared with', 'Difference'],
      ...shownLines.map((line) => [
        splitUnitName(line.name).name,
        String(line.a),
        line.inB ? String(line.b) : '',
        line.inB ? String(line.a - line.b) : '',
      ]),
    ];
  }

  const rightFight = $derived(fights.find((fight) => fight.index === rightIndex) ?? null);
  const currentFight = $derived(fights.find((fight) => fight.index === current) ?? null);

  const DASH = '–';
  const DASH_FIGURE: StatFigure = { value: DASH, muted: true };
  const toneOf = (value: number): 'kill' | 'wipe' | undefined =>
    value > 0 ? 'kill' : value < 0 ? 'wipe' : undefined;
  /** The other pull's column head: its name and length, as the fight list reads it. */
  const rightHead = $derived(
    rightFight === null
      ? reportCopy.compareWith
      : joinFacts(rightFight.name, clockDuration(rightFight.duration_ms)),
  );
  const playerColumns = $derived<StatColumn[]>(
    versusPlayer
      ? [
          { id: 'a', label: 'In this pull' },
          { id: 'b', label: 'The other player' },
          { id: 'diff', label: 'Difference' },
        ]
      : [
          { id: 'a', label: 'This pull' },
          { id: 'b', label: rightHead },
          { id: 'diff', label: 'Difference' },
        ],
  );
  function lineFigures(line: Line, deltaTestid: string): StatFigure[] {
    if (!line.inB) return [{ value: mark + compact(line.a) }, DASH_FIGURE, DASH_FIGURE];
    const delta = line.a - line.b;
    return [
      { value: mark + compact(line.a) },
      { value: mark + compact(line.b), muted: true },
      { value: signed(delta), tone: toneOf(delta), testid: deltaTestid },
    ];
  }
  const abilityColumns = (aHead: string, bHead: string): StatColumn[] => [
    { id: 'a', label: aHead },
    { id: 'b', label: bHead },
    { id: 'diff', label: 'Difference' },
  ];
  function abilityFigures(row: AbilityDiff): StatFigure[] {
    return [
      row.a === null ? DASH_FIGURE : { value: splitMark + compact(row.a) },
      row.b === null ? DASH_FIGURE : { value: splitMark + compact(row.b), muted: true },
      { value: splitSigned(row.delta), tone: toneOf(row.delta), testid: 'ability-delta' },
    ];
  }
  const metricTitle = $derived(metricOptionLabel(metric));
  const sectionMeta = $derived(
    versusPlayer
      ? `${metricTitle}, ${playerName(basePlayer)} against ${playerName(vs)}`
      : right === null
        ? metricTitle
        : `${metricTitle}, ${fightLabel(currentFight, leftWindow) || 'this pull'} against ${fightLabel(rightFight, rightWindow) || 'the other'}`,
  );
</script>

{#snippet abilityTable(rows: AbilityDiff[], aHead: string, bHead: string, testid: string)}
  <div class="flex flex-col gap-1" data-testid={testid}>
    {#if splitTable === null}
      <p class="text-muted text-[13px]">
        Threat has no ability split: the engine keeps it per player, not per spell. Pick a damage or healing
        metric to read the abilities.
      </p>
    {:else if rows.length === 0}
      <p class="text-muted text-[13px]">Neither side used an ability of this kind here.</p>
    {:else}
      <StatTable columns={abilityColumns(aHead, bHead)} nameHead="Ability" bar={false} open label="Abilities">
        {#each rows as row (row.key)}
          <StatRow figures={abilityFigures(row)} bar={false} phone={splitSigned(row.delta)}>
            {#snippet name()}{row.name}{/snippet}
            {#snippet sub()}{#if row.via}{row.via}{/if}{#if row.id !== undefined}<span
                  data-testid="compare-ability-id"
                  title={row.id === 0
                    ? 'Two things share this name; this one is the auto-attack swing, which has no spell id'
                    : `Two spells share this name; this is spell id ${row.id}`}
                  >{row.via ? ' · ' : ''}{row.id === 0 ? 'swing' : `#${row.id}`}</span
                >{/if}{/snippet}
          </StatRow>
        {/each}
      </StatTable>
      <CopyCsv lines={() => abilityCsv(rows, aHead, bHead)} />
      {#if splitScaled}
        <p class="text-muted text-[12px]" data-testid="compare-split-note">
          Marked ~: inside a window the split by ability is each ability's whole-fight amount scaled to the
          window's share of that player's total, not the window's own events. The player totals above are
          exact.
        </p>
      {/if}
    {/if}
  </div>
{/snippet}

{#snippet abilitiesToggle(line: Line)}
  <button
    type="button"
    class="text-gold inline-flex min-h-11 items-center text-[13px] font-semibold underline-offset-2 hover:underline md:min-h-0"
    aria-expanded={openRows.includes(line.guid)}
    onclick={() => toggleRow(line.guid)}
    >{openRows.includes(line.guid) ? 'Hide abilities' : 'Abilities'}</button
  >
{/snippet}

{#snippet playerRow(line: Line, rowTestid: string, deltaTestid: string, panelTestid: string)}
  <StatRow
    figures={lineFigures(line, deltaTestid)}
    bar={false}
    phone={line.inB ? `${mark}${compact(line.b)} · ${signed(line.a - line.b)}` : DASH}
    testid={rowTestid}
  >
    {#snippet icon()}<ClassCrestRing characterClass={classSlugOf(line.class) ?? ''} size={28} />{/snippet}
    {#snippet name()}<span style:color={classColorVar(line.class)}>{splitUnitName(line.name).name}</span
      >{/snippet}
    {#snippet sub()}
      {#if versusPlayer}
        <span>{mark}{compact(line.a)} in this pull</span> ·
        <span>{mark}{compact(line.b)} for {playerName(line.guid === basePlayer ? vs : basePlayer)}</span>
        ·
      {/if}
      {@render abilitiesToggle(line)}
    {/snippet}
  </StatRow>
  {#if openRows.includes(line.guid)}
    <div class="ui-st-r" role="row">
      <div class="col-span-full" role="cell">
        {@render abilityTable(rowsFor(line.guid), 'This pull', rightHead, panelTestid)}
      </div>
    </div>
  {/if}
{/snippet}

{#snippet pickers()}
  <div class="flex flex-wrap items-end gap-x-4 gap-y-3" data-testid="compare-controls">
    <label class="flex w-full flex-col gap-1 md:w-auto">
      <span class="text-muted text-[13px]">{reportCopy.compareWith}</span>
      <Select
        id="compare-with"
        class="w-full md:w-[260px]"
        size="dense"
        value={rightIndex === null ? '' : String(rightIndex)}
        options={compareChoices}
        placeholder={reportCopy.pickFight}
        emptyLabel={selectCopy.noCompareFight}
        testid="compare-with"
        onchange={(event) => {
          const raw = (event.currentTarget as HTMLSelectElement).value;
          onPatch({ compareWith: raw === '' ? null : Number(raw) });
        }}
      />
    </label>
    <label class="flex w-full flex-col gap-1 md:w-auto">
      <span class="text-muted text-[13px]">{reportCopy.orAPlayer}</span>
      <Select
        id="compare-vs"
        class="w-full md:w-[220px]"
        size="dense"
        value={versusPlayer ? vs : ''}
        options={vsChoices}
        testid="compare-vs"
        onchange={(event) => onPatch({ compareVs: (event.currentTarget as HTMLSelectElement).value })}
      />
    </label>
    <label class="flex flex-col gap-1">
      <span class="text-muted text-[13px]">{reportCopy.metric}</span>
      <Select
        id="compare-metric"
        class="w-[160px]"
        size="dense"
        value={metric}
        options={COMPARE_METRICS.map((option) => ({ value: option, label: metricOptionLabel(option) }))}
        testid="compare-metric"
        onchange={(event) => onPatch({ compareMetric: (event.currentTarget as HTMLSelectElement).value })}
      />
    </label>
    {#if rightWhole !== null && phases.length > 0}
      <label class="flex flex-col gap-1">
        <span class="text-muted text-[13px]">{reportCopy.phase}</span>
        <Select
          id="compare-phase"
          class="w-[180px]"
          size="dense"
          bind:value={phase}
          options={[
            { value: '', label: reportCopy.wholePull },
            ...phases.map((name) => ({ value: name, label: name })),
          ]}
          testid="compare-phase"
        />
      </label>
    {/if}
  </div>
{/snippet}

<div class="flex flex-col gap-4" data-testid="compare-mode">
  <Section title="Compare" meta={sectionMeta} panel flush>
    {#if controls}<div class="px-5 pb-3">{@render pickers()}</div>{/if}
    <div>
      {#if error !== ''}
        <div class="px-4 py-4">
          <LoadError message={error} onRetry={() => (attempt += 1)} testid="compare-error" />
        </div>
      {:else if versusPlayer}
        <StatTable
          columns={playerColumns}
          nameHead="Player"
          bar={false}
          label="Two players in this pull"
          testid="compare-cards"
        >
          {#each shownLines as line (line.guid)}
            {@render playerRow(
              line,
              `compare-card-${line.guid}`,
              'compare-card-delta',
              `compare-abilities-card-${line.guid}`,
            )}
          {/each}
        </StatTable>
      {:else if right === null}
        <p class="text-muted px-4 py-4 text-[14px]">
          Pick a second fight, or a second player, to see the difference.
        </p>
      {:else}
        <StatTable
          columns={playerColumns}
          nameHead="Player"
          bar={false}
          label={`Per-player ${METRIC_LABELS[metric]} in ${fightLabel(currentFight, leftWindow) || 'this fight'}, compared with ${fightLabel(rightFight, rightWindow) || 'the selected fight'}`}
          testid="compare-table"
        >
          {#each shownLines as line (line.guid)}
            {@render playerRow(
              line,
              `compare-${line.guid}`,
              'compare-delta',
              `compare-abilities-${line.guid}`,
            )}
          {/each}
        </StatTable>
      {/if}
    </div>
    <div class="flex flex-col gap-2 px-5 pt-2 pb-4">
      {#if error === '' && versusPlayer}
        <p class="text-muted text-[13px]" data-testid="compare-players-scope">
          {playerName(basePlayer)} against
          {playerName(vs)}, in this pull{window === null
            ? ''
            : `'s ${clockDuration(window.startMs)} to ${clockDuration(window.endMs)}`}.
        </p>
        <h3 class="text-strong pt-2 text-[15px] font-bold">
          {playerName(basePlayer)} against {playerName(vs)}, ability by ability
        </h3>
        {@render abilityTable(versusRows, playerName(basePlayer), playerName(vs), 'compare-players')}
        <CopyCsv lines={playerCsv} label="Copy the two players as CSV" />
      {:else if error === '' && right !== null}
        <CopyCsv lines={playerCsv} />
        <Notice
          tone="neutral"
          text="Players missing from the other pull show a dash."
          testid="compare-missing-note"
        />
      {/if}
      <p class="text-muted text-[13px]" data-testid="compare-scope">
        {#if phase !== ''}
          Each side shows its own {phase}: {clockDuration(leftWindow?.startMs ?? 0)} to {clockDuration(
            leftWindow?.endMs ?? 0,
          )} of this pull against {clockDuration(rightWindow?.startMs ?? 0)} to {clockDuration(
            rightWindow?.endMs ?? 0,
          )} of the other, so a long phase and a short one are read phase against phase.
        {:else if window === null}
          Both sides show the whole fight. Set a Window above to compare the same stretch of each pull.
        {:else}
          Both sides show {clockDuration(window.startMs)} to {clockDuration(window.endMs)} of each fight, so a long
          wipe and a short one are read over the same stretch. Clear the Window above to compare whole fights.
        {/if}
      </p>
      {#if rightWhole !== null && phases.length === 0}
        <p class="text-muted text-[13px]" data-testid="compare-phase-note">
          These two pulls have no phase in common, so there is nothing to align by: either their bosses have
          no curated phases, or one pull did not reach the other's.
        </p>
      {/if}
      {#if metric === 'threat' || metric === 'tps'}
        <p class="text-muted text-[13px]" data-testid="compare-threat-note">
          Threat is under the base threat model, which has no tank stance, taunt or threat multipliers yet, so
          a tank can read below the damage dealers they held threat over; a longer pull reads higher for being
          longer, so read the difference against the lengths.{threatScaled
            ? ' Inside a window it is marked ~: the whole fight’s threat scaled to the window’s share, not measured from the window’s own events, the same figure the Threat tab greys.'
            : ''}
        </p>
      {/if}
    </div>
  </Section>
</div>
