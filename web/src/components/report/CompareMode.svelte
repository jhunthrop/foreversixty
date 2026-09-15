<!-- web/src/components/report/CompareMode.svelte -->
<!-- Two fights, one table, one row per player: what they did in each and the difference.
     This is the thing a raid leader does all night -- this pull against the one before --
     and doing it by opening two tabs is what everyone does today instead.

     Both sides are always the whole fight. The chart above can brush a time window, and
     that window scopes Analyze's tables (window.ts's scopeSummary) to a slice of the
     fight that is currently selected -- but a millisecond range from one fight's start
     has no defined meaning against a second fight of a different length, and there is no
     way to show a scaled figure from one side next to an exact figure from the other
     without the reader having to guess which is which. So Compare ignores the window
     entirely rather than inventing a per-fight rescoping, and says so, so a window left
     brushed on the chart is not mistaken for narrowing this table too. -->
<script lang="ts">
  import { splitUnitName } from '../../lib/characters';
  import { classColorVar, formatAmount, formatDuration, outcomeLabel } from '../../lib/report/format';
  import { fetchSummary } from '../../lib/report/load';
  import type { FightEntry, RosterRow, Summary } from '../../lib/report/types';
  import { clampWindow, scopeSummary, type TimeWindow } from '../../lib/report/window';

  let {
    fights,
    current,
    dataBaseUrl,
    left: leftWhole,
    window = null,
    rightIndex = null,
    metric: metricParam = '',
    onPatch,
  }: {
    fights: FightEntry[];
    current: number;
    dataBaseUrl: string;
    left: Summary;
    /** The Analyze window, applied to both sides; null compares whole fights. */
    window?: TimeWindow | null;
    /** The second fight, from the url; null until one is picked. */
    rightIndex?: number | null;
    /** The metric id from the url; '' means the default. */
    metric?: string;
    onPatch: (patch: { compareWith?: number | null; compareMetric?: string }) => void;
  } = $props();

  type CompareMetric = 'damage_done' | 'dps' | 'healing_done' | 'hps' | 'damage_taken' | 'dtps';
  const PER_SECOND = new Set<CompareMetric>(['dps', 'hps', 'dtps']);

  const options = $derived(fights.filter((fight) => fight.index !== current));
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
  const METRIC_IDS: CompareMetric[] = ['damage_done', 'dps', 'healing_done', 'hps', 'damage_taken', 'dtps'];
  const metric = $derived<CompareMetric>(
    (METRIC_IDS as string[]).includes(metricParam) ? (metricParam as CompareMetric) : 'dps',
  );

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
    void fetchSummary(dataBaseUrl, wanted)
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

  function rowMetric(row: RosterRow): number {
    return PER_SECOND.has(metric) ? Math.round(row[metric]) : row[metric];
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

  const lines = $derived.by<Line[]>(() => {
    if (right === null) return [];
    // A plain Map, not SvelteMap: this is a throwaway local built up once and then
    // turned into the plain array `$derived.by` returns, never read key-by-key outside
    // this function -- the same reasoning ReportView.svelte's own `summaries` cache and
    // `percentiles` builder give for their matching eslint-disable.
    // eslint-disable-next-line svelte/prefer-svelte-reactivity
    const byGuid = new Map<string, Line>();
    for (const row of left.roster) {
      byGuid.set(row.guid, { guid: row.guid, name: row.name, class: row.class, a: rowMetric(row), b: 0 });
    }
    for (const row of right.roster) {
      const found = byGuid.get(row.guid);
      const value = rowMetric(row);
      if (found === undefined) {
        byGuid.set(row.guid, { guid: row.guid, name: row.name, class: row.class, a: 0, b: value });
      } else {
        found.b = value;
      }
    }
    return [...byGuid.values()].sort((x, y) => y.a - y.b - (x.a - x.b));
  });

  const rightFight = $derived(fights.find((fight) => fight.index === rightIndex) ?? null);
  const currentFight = $derived(fights.find((fight) => fight.index === current) ?? null);
</script>

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
        {#each options as fight (fight.index)}
          <option value={String(fight.index)} selected={fight.index === rightIndex}
            >{fightLabel(fight)}</option
          >
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
        <option value="dps">DPS</option>
        <option value="damage_done">Damage done</option>
        <option value="hps">HPS</option>
        <option value="healing_done">Healing done</option>
        <option value="dtps">Damage taken per second</option>
        <option value="damage_taken">Damage taken</option>
      </select>
    </label>
  </div>

  {#if error !== ''}
    <p class="text-[14px]" role="alert">{error}</p>
  {:else if right === null}
    <p class="text-muted text-[14px]">Pick a second fight to see the difference per player.</p>
  {:else}
    <!-- Four columns of names and figures do not fit 360px, and `w-full` on a table is a
         floor, not a ceiling: the table grows to its min-content width and eats the page
         gutter. The scroller keeps it a real table -- caption, `th scope="col"`, the
         header association a card stack would lose -- and lets it be wider than the phone
         rather than narrower than its contents. QueriesView.svelte wraps its own result
         table the same way. -->
    <div class="overflow-x-auto">
      <table class="w-full border-collapse text-[14px]" data-testid="compare-table">
        <caption class="sr-only">
          Per-player {metric.replace('_', ' ')} in {currentFight?.name ?? 'this fight'}, compared with {rightFight?.name ??
            'the selected fight'}.
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
          </tr>
        </thead>
        <tbody>
          {#each lines as line (line.guid)}
            <tr class="border-line-soft min-h-11 border-b" data-testid={`compare-${line.guid}`}>
              <!-- Sticky, so a name stays beside its figures when the box scrolls on a phone. -->
              <td
                class="bg-bg sticky left-0 max-w-[40vw] truncate px-2 py-2 font-semibold md:max-w-none"
                style={`color: ${classColorVar(line.class)}`}
              >
                {splitUnitName(line.name).name}
              </td>
              <td class="tabular px-2 py-2 text-right font-mono">{formatAmount(line.a)}</td>
              <td class="text-muted tabular px-2 py-2 text-right font-mono">{formatAmount(line.b)}</td>
              <td
                class="tabular w-[96px] px-2 py-2 text-right font-mono"
                class:text-gold={line.a >= line.b}
                data-testid="compare-delta"
              >
                {line.a - line.b >= 0 ? '+' : ''}{formatAmount(line.a - line.b)}
              </td>
            </tr>
          {/each}
        </tbody>
      </table>
    </div>
  {/if}
</div>
