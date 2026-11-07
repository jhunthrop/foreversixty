<!-- web/src/components/guild/GuildReadiness.svelte -->
<!-- Guild control-centre spec §4.E: the raid leader's pre-pull board, in the flat system (web
     layout spec 2026-11-05 §3A; GUILD-ALL 2026-11-07): an open Section whose one Plate holds the
     board. Worst-first sort, own row pinned, Nudge (officer only) copies to the clipboard -- there
     is no network path from the site to the addon. Desktop renders an eight-column grid; phone
     replaces it with stacked labelled facts per row (§4.E/§7: the grid overflowed 390px). A cell's
     state is its text colour (the Grid's ok, warn and miss tones), never a pill. Officer tools gate
     §3: with the officer tools off (officer-read-only) the plan sentence sits under the heading and
     every Nudge stays on screen, disabled. -->
<script lang="ts">
  import ClassCrestRing from '../character/ClassCrestRing.svelte';
  import { syncedWord } from '../../lib/dates';
  import { classColorVar } from '../../lib/report/format';
  import type { GuildReadinessRow } from '../../lib/guild/api';
  import type { ReadinessRaid } from '../../lib/guild/attendance-contract';
  import {
    attendanceCopy,
    nextRaidToggle,
    showingParts,
    stampLabel,
    syncedLine,
  } from '../../lib/guild/attendance-copy';
  import { FILTER_BTN, FILTER_BTN_ON } from '../../lib/guild/filter-styles';
  import { PHONE_HIT_AREA } from '../../lib/guild/field-styles';
  import { dateLabel } from '../../lib/guild/raid-time';
  import { syncCell } from '../../lib/guild/readiness-raid';
  import EmptyState from '../ui/EmptyState.svelte';
  import Page from '../ui/Page.svelte';
  import Plate from '../ui/Plate.svelte';
  import Ruled from '../ui/Ruled.svelte';
  import Section from '../ui/Section.svelte';
  import Tag from '../ui/Tag.svelte';
  import { button } from '../../lib/ui/button';
  import { CELL_TEXT } from '../../lib/ui/cells';
  import { guildGateCopy } from '../../lib/guild/copy';
  import type { GateFacts } from '../../lib/guild/plan-gate';
  import {
    consumablesKind,
    consumablesLabel,
    enchantKind,
    enchantLabel,
    gearGapLabel,
    AVG_ITEM_LEVEL,
    AVG_ITEM_LEVEL_TOOLTIP,
    LEVELING_GROUP,
    LEVELING_SUFFIX,
    READINESS_EMPTY,
    READINESS_READY,
    isLeveling,
    isWithheldWord,
    itemLevelChip,
    itemLevelLabel,
    levelingLabel,
    medianCaption,
    nudgeText,
    passesEveryCheck,
    pinReadinessOwnRow,
    readinessGroups,
    sortReadinessWorstFirst,
    talentPointsKind,
    talentPointsLabel,
    type ReadinessCellKind,
  } from '../../lib/guild/readiness-view';
  import GuildPlanGate from './GuildPlanGate.svelte';
  import { ROW_RULE, ownRowBandInside } from '../../lib/ui/row';

  let {
    rows,
    officer,
    officerTools,
    raidGuildTools,
    myCharacterKey,
    gateFacts,
    raid = null,
    nextRaidOnly = false,
    onToggleNextRaid = () => {},
    raidHref = null,
    guildTimezone = null,
    viewerTimezone = new Intl.DateTimeFormat().resolvedOptions().timeZone,
  }: {
    rows: GuildReadinessRow[];
    officer: boolean;
    /** False when the guild's officer tools are off (gate spec §3.1, officer-read-only). */
    officerTools: boolean;
    /** False for an officer of a Guild-tier guild (pricing spec §3.2.8). */
    raidGuildTools: boolean;
    myCharacterKey: string | null;
    /** Where the plan sentence's link goes: the guild's settings tab. */
    /** The gate sentence's facts (spec 2026-11-05 §3.2). */
    gateFacts: GateFacts;
    /** The next planned or locked raid's filtered board header; null shows no toggle. */
    raid?: ReadinessRaid | null;
    /** The toggle's state, held by the page so it survives a re-read. */
    nextRaidOnly?: boolean;
    onToggleNextRaid?: () => void;
    /** The raid page the `Showing ...` line links. */
    raidHref?: string | null;
    guildTimezone?: string | null;
    viewerTimezone?: string;
  } = $props();

  const filtered = $derived(raid !== null && nextRaidOnly);
  const showing = $derived(
    raid === null
      ? null
      : showingParts(
          rows.length,
          raid.title,
          stampLabel(new Date(raid.starts_at), viewerTimezone, guildTimezone ?? viewerTimezone),
        ),
  );
  const cellOf = (row: GuildReadinessRow) => syncCell(row, filtered ? raid : null);

  const officerReadOnly = $derived(officer && !(officerTools && raidGuildTools));
  /** Native disabled plus the sentence as description (§3.4); undefined leaves no attribute. */
  const gateDescribedBy = $derived(officerReadOnly ? 'officer-tools-gate-readiness' : undefined);

  // Spec §4.0's visibility matrix: Readiness pins the viewer's own row first for a MEMBER
  // only ("own row pinned first, no Nudge button"); an officer's own matrix cell reads
  // "Nudge on every row" with no pin clause -- the officer board stays strictly
  // worst-first throughout, exactly the raid leader's own pre-pull triage order.
  // The leveling characters follow the raiders under their own rule (week-one spec §3.5); a
  // member's own row pins first within whichever group it is in.
  const groups = $derived(readinessGroups(sortReadinessWorstFirst(rows)));
  const graded = $derived(officer ? groups.graded : pinReadinessOwnRow(groups.graded, myCharacterKey));
  const levelingRows = $derived(
    officer ? groups.leveling : pinReadinessOwnRow(groups.leveling, myCharacterKey),
  );
  const ordered = $derived([...graded, ...levelingRows]);

  let copiedKey = $state<string | null>(null);

  async function onNudge(row: GuildReadinessRow): Promise<void> {
    const text = nudgeText(row);
    try {
      await navigator.clipboard.writeText(text);
      copiedKey = row.character_key;
      setTimeout(() => {
        if (copiedKey === row.character_key) copiedKey = null;
      }, 2000);
    } catch {
      // Clipboard access can be refused (no permission, insecure context) -- the button's
      // own label never promises success beyond "copied," so a refusal just leaves no
      // visible confirmation rather than a thrown error surfacing to the viewer.
    }
  }

  // Tailwind utilities only (no scoped <style>): see GuildTabs.svelte's own header
  // comment -- this tab's rows only ever mount after the guild's async readiness fetch
  // resolves, so a scoped <style> block's CSS is silently dropped from this build's
  // SSR-critical-CSS inlining pass.
  const NUDGE_BTN = button('secondary', 'sm');
  const CHIP = 'flex flex-col gap-0.5 text-[13px] text-text';
  /** Fits the 1024 px content width, where the board's grid starts: the name and gear gap columns give. */
  // Last synced is 128px: a stale sync (`Synced 27 Oct`) stays on one line (guild loop round 3).
  // Enchants grows with the name and gear gap to 170px, so its missing-slot list reads in full
  // (guild loop round 3, web4: `MAIN HAND, CHEST AND 2…` was clipped at 1440).
  const BOARD_COLUMNS =
    'minmax(120px,170px) minmax(120px,170px) minmax(120px,170px) 100px 100px 140px 128px 84px';

  /** A signed character with no export has no gear to read (UX 16): Gear gap to Item level read a dash. */
  const NO_EXPORT_CELLS = ['Gear gap', 'Enchants', 'Consumables', 'Talent points', 'Item level'] as const;
  const DASH = '—';
  /** An expression, so its leading space survives (Svelte trims a block's edge whitespace; WP 20). */
  /** What `sortReadinessWorstFirst` sorts by (WP B). */
  const SORT_NOTE = 'Sorted by the most checks to fix, then the biggest gear gap.';
  const NUDGE_NOTE = ' Nudge copies a message you can paste to them.';

  /** A cell's state as its text colour: met green, to fix red (gold where it only needs a look),
   *  and a cell that has nothing to flag in the muted body text. */
  function cellClass(kind: ReadinessCellKind, warnIsGold = false): string {
    if (kind === 'consent' || kind === 'neutral') return 'text-muted text-[13px]';
    if (kind === 'ok') return `${CELL_TEXT.ok} text-[13px] font-semibold`;
    return `${warnIsGold ? CELL_TEXT.warn : CELL_TEXT.miss} text-[13px] font-semibold`;
  }

  /** A gear gap that says why it has no figure reads muted; a real one reads as plain text. */
  function gearGapTextClass(row: GuildReadinessRow): string {
    return isWithheldWord(gearGapLabel(row)) ? 'text-muted' : 'text-text';
  }

  /** The day the row's export last synced, in the column's one format (`3 days ago`, `27 Oct`). */
  function syncedDay(row: GuildReadinessRow): string {
    return syncedWord(new Date(row.logged_at));
  }

  /** The median caption's colour: above green, below red, at the median muted (week-one spec §3.5). */
  function medianClass(row: GuildReadinessRow): string {
    const delta = row.item_level_delta ?? 0;
    if (delta > 0) return 'text-kill';
    return delta < 0 ? 'text-wipe' : 'text-muted';
  }
</script>

{#snippet raiderName(row: GuildReadinessRow, pinned: boolean)}
  <span class="flex min-w-0 items-center gap-2">
    <ClassCrestRing characterClass={row.class} size={26} />
    <span class="truncate text-[14px] font-bold" style={`color:${classColorVar(row.class)}`}>
      {row.name}{pinned ? ' (you)' : ''}
    </span>
    {#if pinned && passesEveryCheck(row)}<Tag tone="success" testid="guild-readiness-ready"
        >{READINESS_READY}</Tag
      >{/if}
  </span>
{/snippet}

<!-- Last synced: the day, or (next-raid board) a sync older than the raid's cutoff in gold,
     or a signed character with no export, muted. -->
{#snippet syncedCell(row: GuildReadinessRow, dayText: string)}
  {@const cell = cellOf(row)}
  {#if cell.kind === 'stale'}
    <span class="{cellClass('warn', true)} whitespace-nowrap" data-testid="guild-readiness-stale"
      >{cell.text}</span
    >
  {:else}
    <span
      class="text-muted text-[13px]"
      data-testid={cell.kind === 'no_export' ? 'guild-readiness-no-export' : undefined}
      >{cell.kind === 'no_export' ? cell.text : dayText}</span
    >
  {/if}
{/snippet}

{#snippet nudge(row: GuildReadinessRow, extra: string)}
  <button
    class="{NUDGE_BTN} {extra}"
    onclick={() => void onNudge(row)}
    disabled={officerReadOnly}
    aria-describedby={gateDescribedBy}
    aria-label={`Nudge ${row.name}`}
    data-testid={`guild-readiness-nudge-${row.character_key}`}
  >
    {copiedKey === row.character_key ? 'Copied' : 'Nudge'}
  </button>
{/snippet}

{#snippet headCell(label: string, title?: string)}
  <span class="quiet-label" style={title === undefined ? undefined : 'cursor:default'} {title}>{label}</span>
{/snippet}

<div data-testid="guild-readiness-tab">
  <Page layout="ledger">
    <!-- Members only, as on Loot: a gated officer reads the plan sentence instead (guild round 2). -->
    <Section
      title="Readiness"
      meta={officer ? undefined : guildGateCopy.readOnly}
      metaTestid="guild-readiness-read-only"
    >
      {#snippet head()}
        {#if raid !== null && showing !== null}
          <button
            class={`${FILTER_BTN} ${nextRaidOnly ? FILTER_BTN_ON : ''}`}
            aria-pressed={nextRaidOnly}
            onclick={onToggleNextRaid}
            data-testid="guild-readiness-next-raid"
          >
            {nextRaidToggle(dateLabel(new Date(raid.starts_at), guildTimezone ?? viewerTimezone))}
          </button>
        {/if}
      {/snippet}
      <div class="flex flex-col gap-3">
        {#if officerReadOnly}<GuildPlanGate tab="readiness" facts={gateFacts} />{/if}
        {#if raid !== null && showing !== null && nextRaidOnly && rows.length > 0}
          <p class="text-muted text-[14px]" data-testid="guild-readiness-showing">
            {showing.lead}{#if raidHref !== null}<a
                href={raidHref}
                class="text-nav underline {PHONE_HIT_AREA}">{showing.title}</a
              >{:else}{showing.title}{/if}{showing.tail}
          </p>
        {/if}
        {#if filtered && rows.length === 0}
          <EmptyState message={attendanceCopy.readiness.nobodySigned} testid="guild-readiness-nobody" />
        {:else if rows.length === 0}
          <EmptyState
            message={READINESS_EMPTY}
            action={officer ? { label: 'Open the roster', href: '#roster' } : undefined}
            testid="guild-readiness-empty"
          />
        {/if}

        {#if ordered.length > 0}
          <Plate label="Readiness board" testid="guild-readiness-board">
            <div
              class="hidden gap-2 px-4 pt-3 pb-2 lg:grid"
              style:grid-template-columns={BOARD_COLUMNS}
              data-testid="guild-readiness-header"
            >
              {@render headCell('Raider')}
              {@render headCell(
                'Gear gap',
                'Slots with a better item in the best-in-slot list, and for a damage spec the DPS a simulation says they are worth',
              )}
              {@render headCell('Enchants', 'Weapon, chest, cloak and boots checked for an enchant')}
              {@render headCell(
                'Consumables',
                'Flasks, food and potions in bags. Only members who share their bags are checked.',
              )}
              {@render headCell('Talent points', 'Spent points vs 51 at level 60')}
              {@render headCell(AVG_ITEM_LEVEL, AVG_ITEM_LEVEL_TOOLTIP)}
              {@render headCell('Last synced')}
              <span></span>
            </div>

            <Ruled testid="guild-readiness-rows">
              {#each ordered as row, index (row.character_key)}
                {@const pinned = row.character_key === myCharacterKey}
                {@const leveling = isLeveling(row)}
                {#if leveling && index === graded.length && graded.length > 0}
                  <div class="px-4 pt-3 pb-1" data-testid="guild-readiness-leveling-group">
                    <span class="quiet-label">{LEVELING_GROUP}</span>
                  </div>
                {/if}
                <!-- One box per raider holds the rule and the own band (lib/ui/row.ts), so the band and
                     the rules beside it are the same at every width whichever layout below is shown. -->
                <div class={`mx-1 ${ROW_RULE} ${ownRowBandInside(pinned)}`}>
                  <!-- Desktop grid -->
                  <div
                    class="hidden items-center gap-2 py-2 lg:grid"
                    style:grid-template-columns={BOARD_COLUMNS}
                    data-testid={`guild-readiness-row-${row.character_key}`}
                  >
                    {@render raiderName(row, pinned)}
                    {#if leveling}
                      <span class="text-text col-span-5 text-[14px]" data-testid="guild-readiness-leveling">
                        {levelingLabel(row.level ?? 0)}<span class="text-muted text-[13px]"
                          >{LEVELING_SUFFIX}</span
                        >
                      </span>
                    {:else if row.no_export}
                      {#each NO_EXPORT_CELLS as cell (cell)}<span
                          class="text-muted text-[13px]"
                          data-testid="guild-readiness-dash">{DASH}</span
                        >{/each}
                    {:else}
                      <span class={`text-[13px] tabular-nums ${gearGapTextClass(row)}`}
                        >{gearGapLabel(row)}</span
                      >
                      <span class={cellClass(enchantKind(row))} data-testid="guild-readiness-enchants"
                        >{enchantLabel(row)}</span
                      >
                      <span class={cellClass(consumablesKind(row))}>{consumablesLabel(row)}</span>
                      <span class={cellClass(talentPointsKind(row), true)}>{talentPointsLabel(row)}</span>
                      <span class="flex flex-col" data-testid="guild-readiness-item-level">
                        <span
                          class={row.item_level === null
                            ? 'text-muted text-[13px]'
                            : 'text-text text-[14px] tabular-nums'}>{itemLevelLabel(row)}</span
                        >
                        {#if medianCaption(row) !== ''}<span class="text-[12px] {medianClass(row)}"
                            >{medianCaption(row)}</span
                          >{/if}
                      </span>
                    {/if}
                    {@render syncedCell(row, syncedDay(row))}
                    {#if officer && leveling}
                      <span></span>
                    {:else if officer}
                      {@render nudge(row, '')}
                    {/if}
                  </div>

                  <!-- Phone: the raider's facts stacked under their name. -->
                  <div
                    class="flex flex-col gap-2.5 py-2.5 lg:hidden"
                    data-testid={`guild-readiness-row-${row.character_key}`}
                  >
                    <div class="flex items-center justify-between gap-2">
                      {@render raiderName(row, pinned)}
                      {@render syncedCell(row, syncedLine(syncedDay(row)))}
                    </div>
                    {#if leveling}
                      <p class="text-text text-[14px]" data-testid="guild-readiness-leveling">
                        {levelingLabel(row.level ?? 0)}
                      </p>
                    {:else if row.no_export}
                      <div class="flex flex-wrap gap-4">
                        {#each NO_EXPORT_CELLS as cell (cell)}<span class={CHIP}
                            ><span class="quiet-label">{cell}</span><span
                              class="text-muted"
                              data-testid="guild-readiness-dash">{DASH}</span
                            ></span
                          >{/each}
                      </div>
                    {:else}
                      <div class="flex flex-wrap gap-4">
                        <span class={CHIP}
                          ><span class="quiet-label">Gear gap</span><span class={gearGapTextClass(row)}
                            >{gearGapLabel(row)}</span
                          ></span
                        >
                        <span class={CHIP}
                          ><span class="quiet-label">Enchants</span><span
                            class={cellClass(enchantKind(row))}
                            data-testid="guild-readiness-enchants-phone">{enchantLabel(row)}</span
                          ></span
                        >
                        <span class={CHIP}
                          ><span class="quiet-label">Consumables</span><span
                            class={cellClass(consumablesKind(row))}>{consumablesLabel(row)}</span
                          ></span
                        >
                        <span class={CHIP}
                          ><span class="quiet-label">Talent points</span><span
                            class={cellClass(talentPointsKind(row), true)}>{talentPointsLabel(row)}</span
                          ></span
                        >
                        <span class={CHIP}
                          ><span class="quiet-label">{AVG_ITEM_LEVEL}</span><span class={medianClass(row)}
                            >{itemLevelChip(row)}</span
                          ></span
                        >
                      </div>
                    {/if}
                    {#if officer && !leveling}{@render nudge(row, 'w-fit')}{/if}
                  </div>
                </div>
              {/each}
            </Ruled>
          </Plate>
        {/if}
        <p class="text-muted text-[13px]">
          {SORT_NOTE}{#if officer}{NUDGE_NOTE}{/if}
        </p>
      </div>
    </Section>
  </Page>
</div>
