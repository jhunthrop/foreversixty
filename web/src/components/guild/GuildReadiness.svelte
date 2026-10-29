<!-- web/src/components/guild/GuildReadiness.svelte -->
<!-- Guild control-centre spec §4.E: the raid leader's pre-pull board. Worst-first sort,
     own row pinned, Nudge (officer only) copies to the clipboard -- there is no network
     path from the site to the addon. Desktop renders an eight-column grid; phone replaces
     it with stacked labelled chips per row (§4.E/§7: the grid overflowed 390px). Officer
     tools gate §3: with the officer tools off (officer-read-only) the plan sentence sits under
     the h2 and every Nudge stays on screen, disabled. -->
<script lang="ts">
  import ClassCrestRing from '../character/ClassCrestRing.svelte';
  import { dayWord } from '../../lib/dates';
  import { classColorVar } from '../../lib/report/format';
  import type { GuildReadinessRow } from '../../lib/guild/api';
  import type { ReadinessRaid } from '../../lib/guild/attendance-contract';
  import { attendanceCopy, nextRaidToggle, showingParts, stampLabel } from '../../lib/guild/attendance-copy';
  import { FILTER_BTN, FILTER_BTN_ON } from '../../lib/guild/filter-styles';
  import { PILL_WARN } from '../../lib/guild/pill-styles';
  import { PHONE_HIT_AREA } from '../../lib/guild/field-styles';
  import { dateLabel } from '../../lib/guild/raid-time';
  import { syncCell } from '../../lib/guild/readiness-raid';
  import EmptyState from '../ui/EmptyState.svelte';
  import { guildGateCopy } from '../../lib/guild/copy';
  import {
    consumablesKind,
    consumablesLabel,
    enchantKind,
    enchantLabel,
    enchantPillText,
    gearGapLabel,
    isWithheldWord,
    itemLevelKind,
    itemLevelLabel,
    nudgeText,
    pinReadinessOwnRow,
    sortReadinessWorstFirst,
    talentPointsKind,
    talentPointsLabel,
    type ReadinessCellKind,
  } from '../../lib/guild/readiness-view';
  import GuildPlanGate from './GuildPlanGate.svelte';

  let {
    rows,
    officer,
    officerTools,
    raidGuildTools,
    myCharacterKey,
    settingsHref,
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
    settingsHref: string;
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
  const ordered = $derived(
    officer
      ? sortReadinessWorstFirst(rows)
      : pinReadinessOwnRow(sortReadinessWorstFirst(rows), myCharacterKey),
  );

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
  const NUDGE_BTN =
    'inline-flex h-8 items-center justify-center rounded-control border border-[#3a3326] bg-none px-3 text-[11px] font-bold tracking-[0.06em] text-strong uppercase disabled:opacity-50 disabled:cursor-default';
  const CHIP = 'flex flex-col gap-0.5 text-[12px] text-text';

  /** The pill treatment per cell kind (Enchants/Consumables/Talent points): `consent`
   *  reads as plain muted text (never a coloured pill -- there is nothing to flag), `ok`
   *  green, `warn` red/gold depending on the cell. */
  const PILL_BASE = 'pill';
  const PILL_OK = 'text-[#7bff5c] bg-[rgba(30,255,0,.10)] border-[rgba(30,255,0,.25)]';
  const PILL_GOLD = 'text-gold bg-[rgba(229,185,85,.14)] border-[rgba(229,185,85,.35)]';
  // Overrides .pill's fixed single-line 22px height for the Enchants pill alone (inline,
  // never a scoped <style> rule touching the shared .pill class every other pill on the
  // page still uses): a long missing-slot list ("Main hand, Chest +2 more") wraps onto a
  // second line instead of overflowing the 120px column, and clamps there -- the cell's
  // own title attribute carries the untruncated list for a reader who needs it.
  const ENCHANT_PILL_STYLE =
    'height:auto;min-height:22px;white-space:normal;line-height:1.3;display:-webkit-box;-webkit-line-clamp:2;-webkit-box-orient:vertical;overflow:hidden;text-align:left;';

  // The stale pill's text (`Synced 9 days before the raid`) is longer than the 130px Synced column:
  // it drops .pill's fixed 22px height and wraps instead of clipping.
  /** A signed character with no export has no gear to read (UX 16): Gear gap to Item level read a dash. */
  const NO_EXPORT_CELLS = ['Gear gap', 'Enchants', 'Consumables', 'Talent points', 'Item level'] as const;
  const DASH = '—';
  /** An expression, so its leading space survives (Svelte trims a block's edge whitespace; WP 20). */
  /** What `sortReadinessWorstFirst` sorts by (WP B). */
  const SORT_NOTE = 'Sorted by the most checks to fix, then the biggest gear gap.';
  const NUDGE_NOTE = ' Nudge copies a message you can paste to them.';
  const STALE_PILL = 'h-auto min-h-[22px] whitespace-normal py-0.5 text-left leading-[1.3]';

  function pillClass(kind: ReadinessCellKind, warnIsGold = false): string {
    if (kind === 'consent') return '';
    if (kind === 'ok') return `${PILL_BASE} ${PILL_OK}`;
    return `${PILL_BASE} ${warnIsGold ? PILL_GOLD : PILL_WARN}`;
  }

  /** A gear gap that says why it has no figure reads muted; a real one reads as plain text. */
  function gearGapTextClass(row: GuildReadinessRow): string {
    return isWithheldWord(gearGapLabel(row)) ? 'text-muted' : 'text-text';
  }

  /** The day the row's export last synced, as a word in the viewer's zone. */
  function syncedDay(row: GuildReadinessRow): string {
    return dayWord(new Date(row.logged_at));
  }

  /** Item level reads as plain coloured text, never a pill (spec: "signed and coloured
   *  -- green at or above, red below"). */
  function itemLevelTextClass(kind: ReadinessCellKind): string {
    if (kind === 'ok') return 'text-[#7bff5c]';
    if (kind === 'warn') return 'text-[#ff6b5c]';
    return 'text-muted';
  }
</script>

<!-- Last synced: the day, or (next-raid board) a sync older than the raid's cutoff in the warn pill,
     or a signed character with no export, muted. -->
{#snippet syncedCell(row: GuildReadinessRow, size: string, dayText: string)}
  {@const cell = cellOf(row)}
  {#if cell.kind === 'stale'}
    <span class={`${pillClass('warn', true)} ${STALE_PILL}`} data-testid="guild-readiness-stale"
      >{cell.text}</span
    >
  {:else}
    <span
      class={`text-muted ${size}`}
      data-testid={cell.kind === 'no_export' ? 'guild-readiness-no-export' : undefined}
      >{cell.kind === 'no_export' ? cell.text : dayText}</span
    >
  {/if}
{/snippet}

<section class="flex flex-col gap-3" data-testid="guild-readiness-tab">
  <div class="flex items-center justify-between gap-3">
    <h2 class="section-title">Readiness</h2>
    {#if officerReadOnly}<span class="text-muted text-[12px]">{guildGateCopy.readOnly}</span>{/if}
  </div>
  {#if officerReadOnly}<GuildPlanGate tab="readiness" {settingsHref} raidGuild={officerTools} />{/if}

  {#if raid !== null && showing !== null}
    <div class="flex flex-col gap-2">
      <div class="flex flex-wrap gap-2">
        <button
          class={`${FILTER_BTN} ${nextRaidOnly ? FILTER_BTN_ON : ''}`}
          aria-pressed={nextRaidOnly}
          onclick={onToggleNextRaid}
          data-testid="guild-readiness-next-raid"
        >
          {nextRaidToggle(dateLabel(new Date(raid.starts_at), guildTimezone ?? viewerTimezone))}
        </button>
      </div>
      {#if nextRaidOnly && rows.length > 0}
        <p class="text-muted text-[13px]" data-testid="guild-readiness-showing">
          {showing.lead}{#if raidHref !== null}<a href={raidHref} class="text-nav underline {PHONE_HIT_AREA}"
              >{showing.title}</a
            >{:else}{showing.title}{/if}{showing.tail}
        </p>
      {/if}
    </div>
  {/if}
  {#if filtered && rows.length === 0}
    <EmptyState message={attendanceCopy.readiness.nobodySigned} testid="guild-readiness-nobody" />
  {/if}

  {#if ordered.length > 0}<div
      class="hidden gap-2 px-1 pb-1 md:grid"
      style="grid-template-columns:170px 170px 120px 120px 110px 100px 130px 90px"
      data-testid="guild-readiness-header"
    >
      <span class="quiet-label">Raider</span>
      <span
        class="quiet-label"
        style="cursor:help"
        title="Slots with a better item in the best-in-slot list, and the DPS a simulation says they are worth"
        >Gear gap</span
      >
      <span
        class="quiet-label"
        style="cursor:help"
        title="Weapon, chest, cloak and boots checked for an enchant">Enchants</span
      >
      <span
        class="quiet-label"
        style="cursor:help"
        title="Flasks, food and potions in bags. Only members who share their bags are checked."
        >Consumables</span
      >
      <span class="quiet-label" style="cursor:help" title="Spent points vs 51 at level 60">Talent points</span
      >
      <span class="quiet-label" style="cursor:help" title="This character vs the roster's own median"
        >Item level</span
      >
      <span class="quiet-label">Last synced</span>
      <span></span>
    </div>{/if}

  <div class="flex flex-col" data-testid="guild-readiness-rows">
    {#each ordered as row (row.character_key)}
      {@const pinned = row.character_key === myCharacterKey}
      <!-- Desktop grid -->
      <div
        class={`border-line-soft hidden items-center gap-2 border-b px-1 py-2 md:grid ${pinned ? 'bg-[rgba(229,185,85,.06)]' : ''}`}
        style="grid-template-columns:170px 170px 120px 120px 110px 100px 130px 90px"
        data-testid={`guild-readiness-row-${row.character_key}`}
      >
        <span class="flex min-w-0 items-center gap-2">
          <ClassCrestRing characterClass={row.class} size={26} />
          <span class="truncate text-[13px] font-bold" style={`color:${classColorVar(row.class)}`}>
            {row.name}{pinned ? ' (you)' : ''}
          </span>
        </span>
        {#if row.no_export}
          {#each NO_EXPORT_CELLS as cell (cell)}<span
              class="text-muted text-[12px]"
              data-testid="guild-readiness-dash">{DASH}</span
            >{/each}
        {:else}
          <span class={`numeric text-[12px] ${gearGapTextClass(row)}`}>{gearGapLabel(row)}</span>
          {#if enchantKind(row) === 'consent'}
            <span class="text-muted text-[12px]">{enchantLabel(row)}</span>
          {:else}
            <span class={pillClass(enchantKind(row))} style={ENCHANT_PILL_STYLE} title={enchantLabel(row)}
              >{enchantPillText(row)}</span
            >
          {/if}
          {#if consumablesKind(row) === 'consent'}
            <span class="text-muted text-[12px]">{consumablesLabel(row)}</span>
          {:else}
            <span class={pillClass(consumablesKind(row))}>{consumablesLabel(row)}</span>
          {/if}
          {#if talentPointsKind(row) === 'warn'}
            <span class={pillClass('warn', true)}>{talentPointsLabel(row)}</span>
          {:else}
            <span class="text-muted text-[12px]">{talentPointsLabel(row)}</span>
          {/if}
          <span class={`numeric text-[12px] ${itemLevelTextClass(itemLevelKind(row))}`}
            >{itemLevelLabel(row)}</span
          >
        {/if}
        {@render syncedCell(row, 'text-[11px]', syncedDay(row))}
        {#if officer}
          <button
            class={NUDGE_BTN}
            onclick={() => void onNudge(row)}
            disabled={officerReadOnly}
            aria-describedby={gateDescribedBy}
            aria-label={`Nudge ${row.name}`}
            data-testid={`guild-readiness-nudge-${row.character_key}`}
          >
            {copiedKey === row.character_key ? 'Copied' : 'Nudge'}
          </button>
        {/if}
      </div>

      <!-- Phone stacked card: its own Tailwind `md:hidden` (not a scoped <style> block --
           see this component's own header comment) is the only thing controlling display
           at this breakpoint, so there is no cascade-order risk of it fighting the desktop
           grid's `hidden md:grid` above. -->
      <div
        class={`border-line-soft flex flex-col gap-2.5 border-b px-1 py-2.5 md:hidden ${pinned ? 'bg-[rgba(229,185,85,.06)]' : ''}`}
        data-testid={`guild-readiness-row-${row.character_key}`}
      >
        <div class="flex items-center justify-between gap-2">
          <span class="flex min-w-0 items-center gap-2">
            <ClassCrestRing characterClass={row.class} size={26} />
            <span class="truncate text-[13px] font-bold" style={`color:${classColorVar(row.class)}`}>
              {row.name}{pinned ? ' (you)' : ''}
            </span>
          </span>
          {@render syncedCell(row, 'text-[11px]', `Synced ${syncedDay(row)}`)}
        </div>
        {#if row.no_export}
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
            <span class={CHIP}>
              <span class="quiet-label">Enchants</span>
              {#if enchantKind(row) === 'consent'}
                {enchantLabel(row)}
              {:else}
                <span class={pillClass(enchantKind(row))} style={ENCHANT_PILL_STYLE} title={enchantLabel(row)}
                  >{enchantPillText(row)}</span
                >
              {/if}
            </span>
            <span class={CHIP}>
              <span class="quiet-label">Consumables</span>
              {#if consumablesKind(row) === 'consent'}
                {consumablesLabel(row)}
              {:else}
                <span class={pillClass(consumablesKind(row))}>{consumablesLabel(row)}</span>
              {/if}
            </span>
            <span class={CHIP}>
              <span class="quiet-label">Talent points</span>
              {#if talentPointsKind(row) === 'warn'}
                <span class={pillClass('warn', true)}>{talentPointsLabel(row)}</span>
              {:else}
                {talentPointsLabel(row)}
              {/if}
            </span>
            <span class={CHIP}>
              <span class="quiet-label">Item level</span>
              <span class={itemLevelTextClass(itemLevelKind(row))}>{itemLevelLabel(row)}</span>
            </span>
          </div>
        {/if}
        {#if officer}
          <button
            class={`${NUDGE_BTN} w-fit`}
            onclick={() => void onNudge(row)}
            disabled={officerReadOnly}
            aria-describedby={gateDescribedBy}
            aria-label={`Nudge ${row.name}`}
            data-testid={`guild-readiness-nudge-${row.character_key}`}
          >
            {copiedKey === row.character_key ? 'Copied' : 'Nudge'}
          </button>
        {/if}
      </div>
    {/each}
  </div>

  <p class="text-muted text-[12px]">
    {SORT_NOTE}{#if officer}{NUDGE_NOTE}{/if}
  </p>
</section>
