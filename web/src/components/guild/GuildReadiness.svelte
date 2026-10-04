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
    consumablesLabel,
    enchantLabel,
    gearGapLabel,
    itemLevelLabel,
    nudgeText,
    pinReadinessOwnRow,
    sortReadinessWorstFirst,
    talentPointsLabel,
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
        class="guild-readiness-row hidden md:grid"
        class:is-pinned={pinned}
        style="grid-template-columns:170px 170px 120px 120px 110px 100px 130px 90px"
        data-testid={`guild-readiness-row-${row.character_key}`}
      >
        <span class="flex min-w-0 items-center gap-2">
          <ClassCrestRing characterClass={row.class} size={26} />
          <span class="truncate text-[13px] font-bold" style={`color:${classColorVar(row.class)}`}>
            {row.name}{pinned ? ' (you)' : ''}
          </span>
        </span>
        <span class="tabular font-mono text-[12px]">{gearGapLabel(row)}</span>
        <span class="text-[12px]">{enchantLabel(row)}</span>
        <span class="text-[12px]">{consumablesLabel(row)}</span>
        <span class="text-[12px]">{talentPointsLabel(row)}</span>
        <span class="tabular font-mono text-[12px]">{itemLevelLabel(row)}</span>
        <span class="text-muted text-[11px]">{row.logged_at.slice(0, 10)}</span>
        {#if officer}
          <button
            class="guild-nudge-btn"
            onclick={() => void onNudge(row)}
            aria-label={`Nudge ${row.name}`}
            data-testid={`guild-readiness-nudge-${row.character_key}`}
          >
            {copiedKey === row.character_key ? 'Copied' : 'Nudge'}
          </button>
        {/if}
      </div>

      <!-- Phone stacked card. `flex flex-col` is a Tailwind utility, not the scoped
           <style> block below, on purpose: a scoped `display:flex` here previously beat
           Tailwind's own `md:hidden` at desktop widths (both are single-class selectors,
           and this component's own <style> block loads after Tailwind's utilities in the
           built stylesheet), so the phone card and the desktop grid were BOTH visible
           above the md breakpoint -- found when e2e's Nudge-button count came back double
           the roster's own verified count. Tailwind's responsive utilities keep their own
           cascade order consistent, so display now lives there exclusively. -->
      <div
        class="guild-readiness-card flex flex-col md:hidden"
        class:is-pinned={pinned}
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
          <span class="guild-readiness-chip"
            ><span class="label text-muted">Gear gap</span>{gearGapLabel(row)}</span
          >
          <span class="guild-readiness-chip"
            ><span class="label text-muted">Enchants</span>{enchantLabel(row)}</span
          >
          <span class="guild-readiness-chip"
            ><span class="label text-muted">Consumables</span>{consumablesLabel(row)}</span
          >
          <span class="guild-readiness-chip"
            ><span class="label text-muted">Talent pts</span>{talentPointsLabel(row)}</span
          >
          <span class="guild-readiness-chip"
            ><span class="label text-muted">Item level</span>{itemLevelLabel(row)}</span
          >
        </div>
        {#if officer}
          <button
            class="guild-nudge-btn w-fit"
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

<style>
  .guild-readiness-row,
  .guild-readiness-card {
    align-items: center;
    gap: 8px;
    padding: 8px 4px;
    border-bottom: 1px solid var(--color-border-soft);
  }
  .guild-readiness-card {
    gap: 10px;
    padding: 10px 4px;
  }
  .guild-readiness-row.is-pinned,
  .guild-readiness-card.is-pinned {
    background: rgba(229, 185, 85, 0.06);
  }
  .guild-readiness-chip {
    display: flex;
    flex-direction: column;
    gap: 2px;
    font-size: 12px;
    color: var(--color-text);
  }
  .guild-nudge-btn {
    display: inline-flex;
    align-items: center;
    justify-content: center;
    height: 32px;
    padding: 0 12px;
    border: 1px solid #3a3326;
    border-radius: var(--radius-control);
    background: none;
    color: var(--color-strong);
    font-size: 11px;
    font-weight: 700;
    letter-spacing: 0.06em;
    text-transform: uppercase;
  }
</style>
