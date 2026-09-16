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
    playerAbilityDiff,
    type AbilityDiff,
    type CompareMetric,
  } from '../../lib/report/compare';
  import { splitUnitName } from '../../lib/characters';
  import { classColorVar, formatAmount, formatDuration, outcomeLabel } from '../../lib/report/format';
  import { fetchSummary } from '../../lib/report/load';
  import type { FightEntry, RosterRow, Summary } from '../../lib/report/types';
  import { clampWindow, scopeSummary, type TimeWindow } from '../../lib/report/window';
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
  function optionLabel(fight: FightEntry): string {
    const pull = pullOf.get(fight.index) ?? '';
    return `${fight.name}${pull === '' ? '' : ` · ${pull}`} · ${formatDuration(fight.duration_ms)} · ${outcomeLabel(fight).toLowerCase()}`;
  }
  let rightWhole = $state<Summary | null>(null);
  /** Each side in the window, clamped to that fight's own length. */
  const left = $derived(
    window === null ? leftWhole : scopeSummary(leftWhole, clampWindow(window, leftWhole.duration_ms)),
  );
  const right = $derived(
    rightWhole === null || window === null
      ? rightWhole
      : scopeSummary(rightWhole, clampWindow(window, rightWhole.duration_ms)),
  );
  let error = $state('');
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
   */
  const threatScaled = $derived(window !== null && (metric === 'threat' || metric === 'tps'));
  const mark = $derived(threatScaled ? '~' : '');
  /** A difference as sign, mark, magnitude: "+~13,550" and "−~26,473" read the same way round. */
  const signed = (value: number): string =>
    `${value >= 0 ? '+' : '−'}${mark}${formatAmount(Math.abs(value))}`;
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

  function fightLabel(fight: FightEntry | null): string {
    if (fight === null) return '';
    const stretch =
      window === null
        ? formatDuration(fight.duration_ms)
        : `${formatDuration(window.startMs)} to ${formatDuration(Math.min(window.endMs, fight.duration_ms))} of ${formatDuration(fight.duration_ms)}`;
    return `${fight.name} · ${stretch} · ${outcomeLabel(fight).toLowerCase()}`;
  }

  interface Line {
    guid: string;
    name: string;
    class?: string;
    a: number;
    b: number;
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
      });
    }
    if (right !== null) {
      for (const row of right.roster) {
        const found = byGuid.get(row.guid);
        const value = rowMetric(row, right);
        if (found === undefined) {
          byGuid.set(row.guid, { guid: row.guid, name: row.name, class: row.class, a: 0, b: value });
        } else {
          found.b = value;
        }
      }
    }
    return [...byGuid.values()].sort((x, y) => y.a - y.b - (x.a - x.b));
  });

  /**
   * The player the picked one is compared against. With a player picked the question is
   * "what did I do differently from them", so the first side is the pull's own top row of
   * the metric — the player a reader is looking at — unless the url named it.
   */
  const topPlayer = $derived(lines[0]?.guid ?? '');
  /** The players this pull's roster offers as a second player, the picked one excluded. */
  const vsOptions = $derived(
    left.roster
      .filter((row) => row.guid !== topPlayer)
      .map((row) => ({ guid: row.guid, name: splitUnitName(row.name).name }))
      .sort((a, b) => a.name.localeCompare(b.name)),
  );
  /** True while the mode is two players inside this pull rather than two pulls. */
  const versusPlayer = $derived(vs !== '' && left.roster.some((row) => row.guid === vs));
  /** The player table, cut to the two players when one is picked. */
  const shownLines = $derived(
    versusPlayer ? lines.filter((line) => line.guid === topPlayer || line.guid === vs) : lines,
  );
  /** The ability rows the reader opened, by GUID: a row is opened, not a tab. */
  let openRows = $state<string[]>([]);
  function toggleRow(guid: string): void {
    openRows = openRows.includes(guid) ? openRows.filter((entry) => entry !== guid) : [...openRows, guid];
  }
  const splitTable = $derived(metricTable(metric));
  /** Two players inside this pull, ability by ability. */
  const versusRows = $derived(versusPlayer ? playerAbilityDiff(left, topPlayer, vs, metric) : []);
  function rowsFor(guid: string): AbilityDiff[] {
    return abilityDiff(left, right, guid, metric);
  }

  /** `leftWindow` and `rightWindow` are both the one `window` prop until Task 21's phase
   *  picker gives each side its own; kept as two names now so `splitScaled` is written
   *  once and does not change when that lands. */
  const leftWindow = $derived(window);
  const rightWindow = $derived(window);
  /**
   * An ability split inside a window is prorated: window.ts scales each ability by the
   * window's share of its actor's total, because the summary keeps no per-ability series.
   * So the ability table's own figures carry the tilde whenever a window is set, whatever
   * the metric -- which is a different mark from `threatScaled`, the whole-table one that
   * only applies to the threat metrics.
   */
  const splitScaled = $derived(leftWindow !== null || rightWindow !== null);
  const splitMark = $derived(splitScaled ? '~' : '');
  /** A signed difference in the ability table, marked when the split it came from was scaled. */
  const splitSigned = (value: number): string =>
    `${value >= 0 ? '+' : '−'}${splitMark}${formatAmount(Math.abs(value))}`;

  /** An ability diff as lines, in the same shape as every other table's CSV. */
  function abilityCsv(rows: AbilityDiff[], aHead: string, bHead: string): string[][] {
    return [
      ['Ability', 'Via', aHead, bHead, 'Difference'],
      ...rows.map((row) => [
        row.name,
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
        String(line.b),
        String(line.a - line.b),
      ]),
    ];
  }

  const rightFight = $derived(fights.find((fight) => fight.index === rightIndex) ?? null);
  const currentFight = $derived(fights.find((fight) => fight.index === current) ?? null);
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
      <!-- A grid, not a table: three numbers and a name at 360px read better stacked than
           they do as columns that must each keep a header. -->
      <div class="text-muted label hidden grid-cols-[minmax(0,1fr)_96px_96px_96px] gap-x-3 px-2 pb-1 md:grid">
        <span>Ability</span>
        <span class="text-right">{aHead}</span>
        <span class="text-right">{bHead}</span>
        <span class="text-right">Difference</span>
      </div>
      <ul class="flex flex-col">
        {#each rows as row (row.key)}
          <li
            class="border-line-soft grid min-h-11 grid-cols-[minmax(0,1fr)_auto] items-center gap-x-3 gap-y-0.5 border-b px-2 py-2 text-[14px] md:grid-cols-[minmax(0,1fr)_96px_96px_96px]"
          >
            <span class="truncate"
              >{row.name}{#if row.via}<span class="text-muted ml-1 text-[11px]">· {row.via}</span>{/if}</span
            >
            <span class="tabular text-right font-mono md:col-start-2"
              >{row.a === null ? '—' : splitMark + formatAmount(row.a)}<span
                class="label font-body text-muted ml-1.5 md:hidden">{aHead}</span
              ></span
            >
            <span class="text-muted tabular col-start-1 text-left font-mono md:col-start-3 md:text-right"
              >{row.b === null ? '—' : splitMark + formatAmount(row.b)}<span
                class="label font-body ml-1.5 md:hidden">{bHead}</span
              ></span
            >
            <span
              class="tabular text-right font-mono"
              class:text-gold={row.delta >= 0}
              data-testid="ability-delta">{splitSigned(row.delta)}</span
            >
          </li>
        {/each}
      </ul>
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

{#snippet playerCard(line: Line)}
  <li
    class="border-line-soft grid grid-cols-[minmax(0,1fr)_auto] gap-x-3 gap-y-0.5 border-b py-2 text-[14px]"
    data-testid={`compare-card-${line.guid}`}
  >
    <span class="truncate font-semibold" style={`color: ${classColorVar(line.class)}`}
      >{splitUnitName(line.name).name}</span
    >
    <span
      class="tabular text-right font-mono"
      class:text-gold={line.a >= line.b}
      title="This fight less the compared fight"
      data-testid="compare-card-delta">{signed(line.a - line.b)}</span
    >
    <span class="text-muted col-span-2 text-[12px]"
      ><span class="tabular font-mono">{mark}{formatAmount(line.a)}</span> this fight ·
      <span class="tabular font-mono">{mark}{formatAmount(line.b)}</span> compared with</span
    >
    <span class="col-span-2">
      <button
        type="button"
        class="text-nav inline-flex min-h-11 items-center text-[12px] font-bold tracking-[0.06em] uppercase"
        aria-expanded={openRows.includes(line.guid)}
        onclick={() => toggleRow(line.guid)}
        >{openRows.includes(line.guid) ? 'Hide abilities' : 'Abilities'}</button
      >
    </span>
    {#if openRows.includes(line.guid)}
      <span class="col-span-2">
        <!-- The card and the desktop row below both hold this player's expansion at once
             (one hidden by breakpoint, not removed from the DOM), so this copy gets its
             own testid rather than the desktop row's `compare-abilities-${guid}` --
             sharing one would resolve two elements for a single lookup. -->
        {@render abilityTable(
          rowsFor(line.guid),
          'This fight',
          'Compared with',
          `compare-abilities-card-${line.guid}`,
        )}
      </span>
    {/if}
  </li>
{/snippet}

<div class="flex flex-col gap-3" data-testid="compare-mode">
  <p class="text-muted text-[12px]" data-testid="compare-scope">
    {#if window === null}
      Both sides show the whole fight. Set a window in Analyze to compare the same stretch of each pull.
    {:else}
      Both sides show {formatDuration(window.startMs)} to {formatDuration(window.endMs)} of each fight, so a long
      wipe and a short one are read over the same stretch. Clear the window in Analyze to compare whole fights.
    {/if}
  </p>

  <div class="flex flex-wrap items-center gap-x-4 gap-y-2">
    <label class="label text-muted flex w-full flex-wrap items-center gap-2 md:w-auto" for="compare-with">
      Compare with
      <select
        id="compare-with"
        class="border-line-warm bg-raised rounded-control text-text h-11 w-full max-w-full min-w-0 px-2 text-[13px] md:h-9 md:w-auto"
        data-testid="compare-with"
        value={rightIndex === null ? '' : String(rightIndex)}
        onchange={(event) => {
          const raw = (event.currentTarget as HTMLSelectElement).value;
          onPatch({ compareWith: raw === '' ? null : Number(raw) });
        }}
      >
        <option value="">Pick a fight</option>
        <optgroup label="Boss pulls">
          {#each bossOptions as fight (fight.index)}
            <option value={String(fight.index)} selected={fight.index === rightIndex}
              >{optionLabel(fight)}</option
            >
          {/each}
        </optgroup>
        {#if trashOptions.length > 0}
          <optgroup label="Trash">
            {#each trashOptions as fight (fight.index)}
              <option value={String(fight.index)} selected={fight.index === rightIndex}
                >{optionLabel(fight)}</option
              >
            {/each}
          </optgroup>
        {/if}
      </select>
    </label>
    <label class="label text-muted flex w-full flex-wrap items-center gap-2 md:w-auto" for="compare-vs">
      Or a player
      <select
        id="compare-vs"
        class="border-line-warm bg-raised rounded-control text-text h-11 w-full max-w-full min-w-0 px-2 text-[13px] md:h-9 md:w-auto"
        data-testid="compare-vs"
        value={versusPlayer ? vs : ''}
        onchange={(event) => onPatch({ compareVs: (event.currentTarget as HTMLSelectElement).value })}
      >
        <option value="">Compare fights instead</option>
        {#each vsOptions as option (option.guid)}
          <option value={option.guid}>{option.name}</option>
        {/each}
      </select>
    </label>
    <label class="label text-muted flex items-center gap-2" for="compare-metric">
      Metric
      <select
        id="compare-metric"
        class="border-line-warm bg-raised rounded-control text-text h-11 px-2 text-[13px] md:h-9"
        value={metric}
        onchange={(event) => onPatch({ compareMetric: (event.currentTarget as HTMLSelectElement).value })}
        data-testid="compare-metric"
      >
        {#each COMPARE_METRICS as option (option)}
          <option value={option}>{metricOptionLabel(option)}</option>
        {/each}
      </select>
    </label>
    {#if metric === 'threat' || metric === 'tps'}
      <p class="text-muted text-[12px]" data-testid="compare-threat-note">
        Threat is under the base threat model, which has no tank stance, taunt or threat multipliers yet, so a
        tank can read below the damage dealers they held threat over; a longer pull reads higher for being
        longer, so read the difference against the lengths.{threatScaled
          ? ' Inside a window it is marked ~: the whole fight’s threat scaled to the window’s share, not measured from the window’s own events, the same figure the Threat tab greys.'
          : ''}
      </p>
    {/if}
  </div>

  {#if error !== ''}
    <p class="text-[14px]" role="alert">{error}</p>
  {:else if versusPlayer}
    <p class="text-muted text-[12px]" data-testid="compare-players-scope">
      {splitUnitName(left.roster.find((row) => row.guid === topPlayer)?.name ?? '').name} against
      {splitUnitName(left.roster.find((row) => row.guid === vs)?.name ?? '').name}, in this pull{window ===
      null
        ? ''
        : `'s ${formatDuration(window.startMs)} to ${formatDuration(window.endMs)}`}.
    </p>
    <ul class="flex flex-col" data-testid="compare-cards">
      {#each shownLines as line (line.guid)}{@render playerCard(line)}{/each}
    </ul>
    {@render abilityTable(
      versusRows,
      splitUnitName(left.roster.find((row) => row.guid === topPlayer)?.name ?? '').name,
      splitUnitName(left.roster.find((row) => row.guid === vs)?.name ?? '').name,
      'compare-players',
    )}
    <CopyCsv lines={playerCsv} label="Copy the two players as CSV" />
  {:else if right === null}
    <p class="text-muted text-[14px]">Pick a second fight, or a second player, to see the difference.</p>
  {:else}
    <!-- Four columns of names and figures do not fit 360px, and `w-full` on a table is a
         floor, not a ceiling: the table grows to its min-content width and eats the page
         gutter. The scroller keeps it a real table -- caption, `th scope="col"`, the
         header association a card stack would lose -- and lets it be wider than the phone
         rather than narrower than its contents. QueriesView.svelte wraps its own result
         table the same way. -->
    <!-- A phone gets one block per player: a four-column table in a 354px box showed five
         names and no numbers at rest. -->
    <ul class="flex flex-col md:hidden" data-testid="compare-cards">
      {#each shownLines as line (line.guid)}{@render playerCard(line)}{/each}
    </ul>
    <div class="hidden overflow-x-auto md:block">
      <table class="w-full border-collapse text-[14px]" data-testid="compare-table">
        <caption class="sr-only">
          Per-player {METRIC_LABELS[metric]} in {fightLabel(currentFight) || 'this fight'}, compared with {fightLabel(
            rightFight,
          ) || 'the selected fight'}.
        </caption>
        <thead>
          <tr class="border-line-soft border-b text-left">
            <th scope="col" class="label text-muted bg-bg sticky left-0 px-2 py-2 font-bold">Player</th>
            <th scope="col" class="label text-muted px-2 py-2 text-right font-bold">
              This fight
              <span class="block truncate text-[11px] normal-case">{fightLabel(currentFight)}</span>
            </th>
            <th scope="col" class="label text-muted px-2 py-2 text-right font-bold">
              Compared with
              <span class="block truncate text-[11px] normal-case">{fightLabel(rightFight)}</span>
            </th>
            <th scope="col" class="label text-muted px-2 py-2 text-right font-bold">Difference</th>
            <th scope="col"><span class="sr-only">Abilities</span></th>
          </tr>
        </thead>
        <tbody>
          {#each shownLines as line (line.guid)}
            <tr class="border-line-soft min-h-11 border-b" data-testid={`compare-${line.guid}`}>
              <!-- Sticky, so a name stays beside its figures when the box scrolls on a phone. -->
              <td
                class="bg-bg sticky left-0 max-w-[40vw] truncate px-2 py-2 font-semibold md:max-w-none"
                style={`color: ${classColorVar(line.class)}`}
              >
                {splitUnitName(line.name).name}
              </td>
              <td class="tabular px-2 py-2 text-right font-mono">{mark}{formatAmount(line.a)}</td>
              <td class="text-muted tabular px-2 py-2 text-right font-mono">{mark}{formatAmount(line.b)}</td>
              <td
                class="tabular w-[96px] px-2 py-2 text-right font-mono"
                class:text-gold={line.a >= line.b}
                data-testid="compare-delta"
              >
                {signed(line.a - line.b)}
              </td>
              <td class="px-2 py-2 text-right">
                <button
                  type="button"
                  class="text-nav inline-flex min-h-11 items-center text-[12px] font-bold tracking-[0.06em] uppercase md:min-h-0"
                  aria-expanded={openRows.includes(line.guid)}
                  onclick={() => toggleRow(line.guid)}
                  >{openRows.includes(line.guid) ? 'Hide abilities' : 'Abilities'}</button
                >
              </td>
            </tr>
            {#if openRows.includes(line.guid)}
              <tr class="bg-card-top">
                <td colspan="5" class="px-2 py-3">
                  {@render abilityTable(
                    rowsFor(line.guid),
                    'This fight',
                    'Compared with',
                    `compare-abilities-${line.guid}`,
                  )}
                </td>
              </tr>
            {/if}
          {/each}
        </tbody>
      </table>
    </div>
    <CopyCsv lines={playerCsv} />
  {/if}
</div>
