<!-- web/src/components/guild/dense/ReadinessBoard.svelte -->
<!-- Readiness as a matrix (dense spec 5.4): every graded raider on one grid, the short first, a
     column per check (attunement, item level against the median, the gain to BiS, enchants, flask and
     food, talent points, the days since the last log) and a fixed Nudge column. A failing cell names
     the shortfall in the wipe colour; a pass is a muted tick. The shortfall words also ride the
     raider's sub line, so a phone keeps them when the columns fold. -->
<script lang="ts">
  import type { GuildReadinessRow } from '../../../lib/guild/api';
  import { classColorVar } from '../../../lib/report/format';
  import { enchantSlotLabel } from '../../../lib/guild/readiness-view';
  import type { StatFigure } from '../../../lib/ui/stat-table';
  import ClassCrestRing from '../../character/ClassCrestRing.svelte';
  import StatRow from '../../ui/StatRow.svelte';
  import StatTable from '../../ui/StatTable.svelte';

  let {
    short,
    ready,
    officer,
    readOnly = false,
    describedBy,
    copiedKey = null,
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
    copiedKey?: string | null;
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
  const columns = $derived([
    { id: 'state', label: 'State' },
    { id: 'attuned', label: 'Attuned' },
    { id: 'ilvl', label: median === null ? 'Item level' : `Item level vs ${median}` },
    { id: 'bis', label: 'To BiS' },
    { id: 'enchants', label: 'Enchants' },
    { id: 'stock', label: 'Flask, food' },
    { id: 'talents', label: 'Talents' },
    ...(fireKeys === null ? [] : [{ id: 'fire', label: 'Fire res' }]),
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
    const stocked = row.flask_food == null ? null : row.flask_food.flask_or_elixir && row.flask_food.food;
    return [
      shortfalls.length > 0
        ? { value: `${shortfalls.length} short`, tone: 'wipe' }
        : { value: 'Ready', muted: true },
      check(row.attuned, '×'),
      delta === null
        ? UNKNOWN
        : { value: delta > 0 ? `+${delta}` : delta === 0 ? '±0' : `−${-delta}`, muted: true },
      row.gear_gap?.gain_dps == null
        ? UNKNOWN
        : {
            value: `+${Math.round(row.gear_gap.gain_dps)}${row.gear_gap.unit === null ? '' : ` ${row.gear_gap.unit.toUpperCase()}`}`,
          },
      row.enchants.checked ? check(missing.length === 0, missing.map(enchantSlotLabel).join(', ')) : UNKNOWN,
      check(stocked, '×'),
      row.talent_points_read === false
        ? UNKNOWN
        : check(row.talent_points_unspent === 0, String(row.talent_points_unspent)),
      ...(fireKeys === null
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

<StatTable {columns} nameHead="Raider" bar={false} dense label="Readiness" testid="guild-readiness-board">
  {#each entries as { row, shortfalls }, index (row.character_key)}
    {#if index === short.length && short.length > 0 && ready.length > 0}
      <StatRow total bar={false} figures={blank}>{#snippet name()}Ready {ready.length}{/snippet}</StatRow>
    {/if}
    <StatRow
      bar={false}
      figures={figures(row, shortfalls)}
      testid={`guild-readiness-row-${row.character_key}`}
    >
      {#snippet icon()}<ClassCrestRing characterClass={row.class} size={28} />{/snippet}
      {#snippet name()}<span style:color={classColorVar(row.class)}>{row.name}</span>{/snippet}
      {#snippet sub()}{row.spec} · {row.item_level ?? '–'} ilvl{#if shortfalls.length > 0}<div
            class="text-wipe whitespace-normal"
            data-testid="guild-readiness-shortfalls"
          >
            {shortfalls.join(', ')}{#if officer}<button
                type="button"
                class="nudge"
                disabled={readOnly}
                aria-describedby={describedBy}
                onclick={() => onNudge(row)}
                aria-label={`Nudge ${row.name}`}
                data-testid={`guild-readiness-nudge-${row.character_key}`}
                >{copiedKey === row.character_key ? 'Copied' : 'Nudge'}</button
              >{/if}
          </div>{/if}{/snippet}
    </StatRow>
  {/each}
</StatTable>

<style>
  .nudge {
    display: block;
    min-height: 32px;
    margin-top: 2px;
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
