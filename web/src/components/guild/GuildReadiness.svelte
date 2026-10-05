<!-- web/src/components/guild/GuildReadiness.svelte -->
<!-- Guild control-centre spec §4.E: the raid leader's pre-pull board. Worst-first sort,
     own row pinned, Nudge (officer only) copies to the clipboard -- there is no network
     path from the site to the addon. Desktop renders an eight-column grid; phone replaces
     it with stacked labelled chips per row (§4.E/§7: the fixed-pixel grid overflowed a
     390px viewport in the round-2 mock review, fixed before these boards shipped). -->
<script lang="ts">
  import ClassCrestRing from '../character/ClassCrestRing.svelte';
  import { classColorVar } from '../../lib/report/format';
  import type { GuildReadinessRow } from '../../lib/guild/api';
  import {
    consumablesKind,
    consumablesLabel,
    enchantKind,
    enchantLabel,
    gearGapLabel,
    itemLevelKind,
    itemLevelLabel,
    nudgeText,
    pinReadinessOwnRow,
    sortReadinessWorstFirst,
    talentPointsKind,
    talentPointsLabel,
    type ReadinessCellKind,
  } from '../../lib/guild/readiness-view';

  let {
    rows,
    officer,
    myCharacterKey,
  }: { rows: GuildReadinessRow[]; officer: boolean; myCharacterKey: string | null } = $props();

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
    'inline-flex h-8 items-center justify-center rounded-control border border-[#3a3326] bg-none px-3 text-[11px] font-bold tracking-[0.06em] text-strong uppercase';
  const CHIP = 'flex flex-col gap-0.5 text-[12px] text-text';

  /** The pill treatment per cell kind (Enchants/Consumables/Talent points): `consent`
   *  reads as plain muted text (never a coloured pill -- there is nothing to flag), `ok`
   *  green, `warn` red/gold depending on the cell. */
  const PILL_BASE = 'pill';
  const PILL_OK = 'text-[#7bff5c] bg-[rgba(30,255,0,.10)] border-[rgba(30,255,0,.25)]';
  const PILL_WARN = 'text-[#ff6b5c] bg-[rgba(255,107,92,.12)] border-[rgba(255,107,92,.32)]';
  const PILL_GOLD = 'text-gold bg-[rgba(229,185,85,.14)] border-[rgba(229,185,85,.35)]';

  function pillClass(kind: ReadinessCellKind, warnIsGold = false): string {
    if (kind === 'consent') return '';
    if (kind === 'ok') return `${PILL_BASE} ${PILL_OK}`;
    return `${PILL_BASE} ${warnIsGold ? PILL_GOLD : PILL_WARN}`;
  }

  /** Item level reads as plain coloured text, never a pill (spec: "signed and coloured
   *  -- green at or above, red below"). */
  function itemLevelTextClass(kind: ReadinessCellKind): string {
    if (kind === 'ok') return 'text-[#7bff5c]';
    if (kind === 'warn') return 'text-[#ff6b5c]';
    return 'text-muted';
  }
</script>

<section class="flex flex-col gap-3" data-testid="guild-readiness-tab">
  <h2 class="section-title text-[18px]">Readiness</h2>

  <div
    class="hidden gap-2 px-1 pb-1 md:grid"
    style="grid-template-columns:170px 170px 120px 120px 110px 100px 130px 90px"
    data-testid="guild-readiness-header"
  >
    <span class="label text-muted">Raider</span>
    <span class="label text-muted" style="cursor:help" title="Upgrades and DPS gain vs this band's BiS"
      >Gear gap</span
    >
    <span
      class="label text-muted"
      style="cursor:help"
      title="Weapon, chest, cloak, boots checked for a non-zero enchant id">Enchants</span
    >
    <span
      class="label text-muted"
      style="cursor:help"
      title="Flask/food/potion presence in bags -- gear_bags consent only">Consumables</span
    >
    <span class="label text-muted" style="cursor:help" title="Spent points vs 51 at level 60"
      >Talent points</span
    >
    <span class="label text-muted" style="cursor:help" title="This character vs the roster's own median"
      >Item level</span
    >
    <span class="label text-muted">Last synced</span>
    <span></span>
  </div>

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
        <span class="tabular text-text font-mono text-[12px]">{gearGapLabel(row)}</span>
        {#if enchantKind(row) === 'consent'}
          <span class="text-muted text-[12px]">{enchantLabel(row)}</span>
        {:else}
          <span class={pillClass(enchantKind(row))}>{enchantLabel(row)}</span>
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
        <span class={`tabular font-mono text-[12px] ${itemLevelTextClass(itemLevelKind(row))}`}
          >{itemLevelLabel(row)}</span
        >
        <span class="text-muted text-[11px]">{row.logged_at.slice(0, 10)}</span>
        {#if officer}
          <button
            class={NUDGE_BTN}
            onclick={() => void onNudge(row)}
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
          <span class="text-muted text-[11px]">Synced {row.logged_at.slice(0, 10)}</span>
        </div>
        <div class="flex flex-wrap gap-4">
          <span class={CHIP}><span class="label text-muted">Gear gap</span>{gearGapLabel(row)}</span>
          <span class={CHIP}>
            <span class="label text-muted">Enchants</span>
            {#if enchantKind(row) === 'consent'}
              {enchantLabel(row)}
            {:else}
              <span class={pillClass(enchantKind(row))}>{enchantLabel(row)}</span>
            {/if}
          </span>
          <span class={CHIP}>
            <span class="label text-muted">Consumables</span>
            {#if consumablesKind(row) === 'consent'}
              {consumablesLabel(row)}
            {:else}
              <span class={pillClass(consumablesKind(row))}>{consumablesLabel(row)}</span>
            {/if}
          </span>
          <span class={CHIP}>
            <span class="label text-muted">Talent pts</span>
            {#if talentPointsKind(row) === 'warn'}
              <span class={pillClass('warn', true)}>{talentPointsLabel(row)}</span>
            {:else}
              {talentPointsLabel(row)}
            {/if}
          </span>
          <span class={CHIP}>
            <span class="label text-muted">Item level</span>
            <span class={itemLevelTextClass(itemLevelKind(row))}>{itemLevelLabel(row)}</span>
          </span>
        </div>
        {#if officer}
          <button
            class={`${NUDGE_BTN} w-fit`}
            onclick={() => void onNudge(row)}
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
    Sorted worst-first. "Nudge" copies a message to the clipboard — there is no network path to the addon, so
    this is copy only, never a push.
  </p>
</section>
