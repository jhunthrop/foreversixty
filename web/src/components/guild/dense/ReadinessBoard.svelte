<!-- web/src/components/guild/dense/ReadinessBoard.svelte -->
<!-- Readiness as a matrix (dense spec 5.4): every graded raider on one grid, the short first, a
     column per check (attunement, item level against the median, the gain to BiS, enchants, flask and
     food, talent points, the days since the last log) and a fixed Nudge column. A failing cell names
     the shortfall in the wipe colour; a pass is a muted tick. The shortfall words also ride the
     raider's sub line, so a phone keeps them when the columns fold. -->
<script lang="ts">
  import type { GuildReadinessRow } from '../../../lib/guild/api';
  import { classColorVar } from '../../../lib/report/format';
  import { enchantSlotLabel, missingCount } from '../../../lib/guild/readiness-view';
  import type { StatFigure } from '../../../lib/ui/stat-table';
  import ClassCrestRing from '../../character/ClassCrestRing.svelte';
  import StatRow from '../../ui/StatRow.svelte';
  import StatTable from '../../ui/StatTable.svelte';

  /** The idle word: the shared CopyState swaps in Copied or Copy failed. */
  const NUDGE = 'Nudge';

  let {
    short,
    ready,
    officer,
    readOnly = false,
    describedBy,
    nudgeLabel = () => NUDGE,
    median = null,
    fireKeys = null,
    now,
    onNudge,
  }: {
    short: readonly { row: GuildReadinessRow; shortfalls: string[] }[];
    ready: readonly GuildReadinessRow[];
    officer: boolean;
    /** Officer tools off: Nudge shows, disabled. */
    readOnly?: boolean;
    /** The plan sentence's id that describes a disabled Nudge. */
    describedBy?: string;
    /** The Nudge button's words for a row: Nudge, then Copied or Copy failed (lib/ui/copy-state). */
    nudgeLabel?: (characterKey: string) => string;
    median?: number | null;
    /** The raid asks for fire resistance: these signed tanks' totals fill a Fire res column. */
    fireKeys?: ReadonlySet<string> | null;
    now: Date;
    onNudge: (row: GuildReadinessRow) => void;
  } = $props();

  const STALE_DAYS = 7;
  const DAY_MS = 86_400_000;
  /** From this many days without a log the Synced cell shows the age. */
  const SEEN_FROM_DAYS = 2;
  const daysOf = (row: GuildReadinessRow): number =>
    Math.floor((now.getTime() - new Date(row.logged_at).getTime()) / DAY_MS);
  const entries = $derived([...short, ...ready.map((row) => ({ row, shortfalls: [] as string[] }))]);
  /** The Synced column shows only when someone has not logged for SEEN_FROM_DAYS. */
  const anyStale = $derived(entries.some(({ row }) => daysOf(row) >= SEEN_FROM_DAYS));
  /** The ready raiders fold behind their count while someone is short. */
  let showReady = $state(false);
  /** Fire resistance shows when a shown row is a signed tank. */
  const fireShown = $derived(
    fireKeys !== null &&
      [...short.map((e) => e.row), ...(showReady || short.length === 0 ? ready : [])].some((r) =>
        fireKeys.has(r.character_key),
      ),
  );
  const columns = $derived([
    { id: 'state', label: 'Ready' },
    { id: 'attuned', label: 'Attuned' },
    { id: 'ilvl', label: median === null ? 'Item level' : `Ilvl vs ${median}` },
    { id: 'bis', label: 'To BiS' },
    { id: 'enchants', label: 'Enchants' },
    { id: 'flask', label: 'Flask' },
    { id: 'food', label: 'Food' },
    { id: 'talents', label: 'Talents' },
    ...(fireShown ? [{ id: 'fire', label: 'Fire res' }] : []),
    ...(anyStale ? [{ id: 'synced', label: 'Synced' }] : []),
  ]);
  const blank = $derived(columns.map(() => ({ value: '' })));
  const PASS: StatFigure = { value: '✓', muted: true };
  const UNKNOWN: StatFigure = { value: '–', muted: true };

  function check(pass: boolean | null | undefined, fail: string): StatFigure {
    if (pass === null || pass === undefined) return UNKNOWN;
    return pass ? PASS : { value: fail, tone: 'wipe' };
  }

  function figures(row: GuildReadinessRow, shortfalls: readonly string[]): StatFigure[] {
    const delta = row.item_level_delta;
    const days = daysOf(row);
    const missing = row.enchants.missing_slots;
    return [
      shortfalls.length > 0
        ? { value: `${missingCount(row, shortfalls)} missing`, tone: 'wipe' }
        : { value: '', muted: true },
      check(row.attuned, '×'),
      delta === null
        ? UNKNOWN
        : { value: delta > 0 ? `+${delta}` : delta === 0 ? '±0' : `−${-delta}`, muted: true },
      row.gear_gap?.gain_dps == null || Math.round(row.gear_gap.gain_dps) === 0
        ? UNKNOWN
        : {
            value: `+${Math.round(row.gear_gap.gain_dps)}${row.gear_gap.unit === null ? '' : ` ${row.gear_gap.unit.toUpperCase()}`}`,
          },
      row.enchants.checked ? check(missing.length === 0, missing.map(enchantSlotLabel).join(', ')) : UNKNOWN,
      check(row.flask_food == null ? null : row.flask_food.flask_or_elixir, '×'),
      check(row.flask_food == null ? null : row.flask_food.food, '×'),
      row.talent_points_read === false
        ? UNKNOWN
        : check(row.talent_points_unspent === 0, `${row.talent_points_unspent} pts`),
      ...(!fireShown || fireKeys === null
        ? []
        : [
            fireKeys.has(row.character_key) && row.resistances != null
              ? { value: String(row.resistances.fire) }
              : { value: '–', muted: true },
          ]),
      ...(anyStale
        ? [
            days >= STALE_DAYS
              ? { value: `${days} d`, tone: 'wipe' as const }
              : { value: days < SEEN_FROM_DAYS ? '' : `${days} d`, muted: true },
          ]
        : []),
    ];
  }
</script>

<div class="matrix">
  <StatTable
    {columns}
    nameHead="Raider"
    bar={false}
    dense
    action={officer}
    label="Readiness"
    testid="guild-readiness-board"
  >
    {#each entries as { row, shortfalls }, index (row.character_key)}
      {#snippet nudge()}<button
          type="button"
          class="nudge"
          disabled={readOnly}
          aria-describedby={describedBy}
          onclick={() => onNudge(row)}
          data-testid={`guild-readiness-nudge-${row.character_key}`}
          ><span aria-live="polite">{nudgeLabel(row.character_key)}</span><span class="sr-only">
            {row.name}</span
          ></button
        >{/snippet}
      {#if index === short.length && short.length > 0 && ready.length > 0}
        <StatRow total bar={false} figures={blank}
          >{#snippet name()}<button
              type="button"
              class="fold"
              aria-expanded={showReady}
              onclick={() => (showReady = !showReady)}
              data-testid="guild-readiness-ready-toggle">Ready {ready.length} {showReady ? '▴' : '▾'}</button
            >{/snippet}</StatRow
        >
      {/if}
      {#if index < short.length || showReady || short.length === 0}
        <StatRow
          bar={false}
          action={officer && shortfalls.length > 0 ? nudge : undefined}
          figures={figures(row, shortfalls)}
          testid={`guild-readiness-row-${row.character_key}`}
        >
          {#snippet icon()}<ClassCrestRing characterClass={row.class} size={28} />{/snippet}
          {#snippet name()}<span style:color={classColorVar(row.class)}>{row.name}</span>{/snippet}
          {#snippet sub()}<span class="sub-line"><span>{row.spec} · {row.item_level ?? '–'} ilvl</span></span
            >{#if shortfalls.length > 0}<span
                class="reasons text-wipe"
                data-testid="guild-readiness-shortfalls">{shortfalls.join(', ')}</span
              >{/if}{/snippet}
        </StatRow>
      {/if}
    {/each}
  </StatTable>
</div>

<style>
  /* One column grid for every row: the figures take fixed tracks, so a long cell never shifts the next. */
  .matrix :global(.ui-st) {
    --st-fig: 92px;
  }
  .fold {
    font: inherit;
    color: var(--color-gold);
    cursor: pointer;
  }
  .sub-line {
    display: flex;
    align-items: center;
    justify-content: space-between;
    gap: 12px;
  }
  .reasons {
    display: block;
    white-space: normal;
  }
  @media (min-width: 1024px) {
    .reasons {
      display: none;
    }
  }
  .nudge {
    display: inline-block;
    min-height: 32px;
    flex: none;
    padding: 0;
    font: 600 13px/32px var(--font-body);
    color: var(--color-gold);
    cursor: pointer;
  }
  @media (pointer: coarse) {
    .nudge {
      min-height: 44px;
      line-height: 44px;
    }
  }
  .nudge:hover:not(:disabled) {
    text-decoration: underline;
  }
  .nudge:disabled {
    color: var(--color-muted);
    cursor: default;
  }
</style>
