<!-- web/src/components/guild/GuildLoot.svelte -->
<!-- Guild control-centre spec §4.F/§12.1: the loot council helper. Boss picker defaults to
     the next unkilled named boss; Onyxia today reads "next unkilled: none, farm." Each
     drop shows up to four ranked candidates; the awarded item's own top candidate carries
     the "Awarded" tag, every other candidate shows a muted "Awarded to {name}" note, and an
     unawarded item keeps the Award button on every row (§12.1's round-2 fix: one awardee
     per item, not one per candidate row). Member sees every drop and ranking read-only
     (§10.1, proposed); only officer sees the Award control. -->
<script lang="ts">
  import ClassCrestRing from '../character/ClassCrestRing.svelte';
  import { classColorVar } from '../../lib/report/format';
  import type { GuildLootItem, GuildLootPage } from '../../lib/guild/api';
  import { bossPickerLabel, candidateAction, candidateGainLabel } from '../../lib/guild/loot-view';
  import EmptyState from '../ui/EmptyState.svelte';

  let {
    loot,
    officer,
    memberReadOnly,
    busyItemId,
    onAward,
  }: {
    loot: GuildLootPage | null;
    officer: boolean;
    memberReadOnly: boolean;
    busyItemId: number | null;
    onAward: (item: GuildLootItem, characterKey: string) => void;
  } = $props();

  const selectedEncounter = $derived(loot?.encounters.find((e) => e.encounter_id === loot.selected));
</script>

<section class="flex flex-col gap-4" data-testid="guild-loot-tab">
  <div class="flex items-center justify-between gap-3">
    <h2 class="section-title text-[18px]">Loot</h2>
    {#if memberReadOnly}<span class="text-muted text-[12px]">Read-only</span>{/if}
  </div>

  {#if loot === null || loot.items.length === 0}
    <EmptyState
      message="No bosses killed yet this tier — loot ranking starts after the first kill"
      testid="guild-loot-empty"
    />
  {:else}
    <div class="flex items-center gap-3" data-testid="guild-loot-boss-picker">
      <span class="label text-muted">Boss</span>
      <span
        class="border-gold text-gold inline-flex h-9 items-center rounded-md border px-3 text-[13px] font-bold"
      >
        {bossPickerLabel(selectedEncounter)}
      </span>
    </div>

    <div class="flex flex-col gap-4" data-testid="guild-loot-items">
      {#each loot.items as item (item.item_id)}
        <div
          class="border-line-soft bg-raised flex flex-col gap-2 rounded-md border p-4"
          data-testid={`guild-loot-item-${item.item_id}`}
        >
          <div class="flex items-center justify-between gap-2">
            <span class="font-display text-[15px] font-bold">{item.name}</span>
            {#if item.awarded_to !== null}
              <span
                class="pill"
                style="color:#7bff5c;background:rgba(30,255,0,.10);border-color:rgba(30,255,0,.25)"
                >Awarded</span
              >
            {/if}
          </div>
          <div class="flex flex-col">
            {#each item.candidates as candidate (candidate.character_key)}
              {@const action = candidateAction(item, candidate.character_key)}
              <!-- flex-wrap (live-fix round, defect 5): every real candidate carries
                   `not_sim_checked` today (the tier-1 fallback, CONTROL_CENTRE.md), and
                   that extra badge no longer fits this row's fixed-width items on one
                   line at 390px -- wraps onto a second line there instead of overflowing
                   the viewport sideways. -->
              <div class="border-line-soft flex flex-wrap items-center gap-x-3 gap-y-1 border-b py-1.5">
                <ClassCrestRing characterClass={candidate.class} size={26} />
                <span
                  class="min-w-[120px] text-[13px] font-semibold"
                  style={`color:${classColorVar(candidate.class)}`}
                >
                  {candidate.name}
                </span>
                <span class="tabular font-mono text-[12px]">{candidateGainLabel(candidate)}</span>
                {#if candidate.not_sim_checked}
                  <span class="text-muted text-[11px]">not sim-checked</span>
                {/if}
                <span class="text-muted tabular font-mono text-[12px]">
                  {candidate.attendance.present}/{candidate.attendance.nights} nights
                </span>
                {#if candidate.already_equivalent}
                  <span
                    class="pill"
                    style="color:var(--color-muted);background:rgba(154,148,132,.12);border-color:rgba(154,148,132,.3)"
                  >
                    Already holds equivalent
                  </span>
                {/if}
                <span class="ml-auto">
                  {#if action.kind === 'awarded'}
                    <span
                      class="pill"
                      style="color:#7bff5c;background:rgba(30,255,0,.10);border-color:rgba(30,255,0,.25)"
                      >Awarded</span
                    >
                  {:else if action.kind === 'awarded-to-other'}
                    <span class="text-muted text-[12px]">{action.label}</span>
                  {:else if officer}
                    <button
                      class="rounded-control text-strong inline-flex h-7 items-center justify-center border border-[#3a3326] bg-none px-3 text-[11px] font-bold tracking-[0.06em] uppercase"
                      onclick={() => onAward(item, candidate.character_key)}
                      disabled={busyItemId === item.item_id}
                      aria-label={`Award ${item.name} to ${candidate.name}`}
                      data-testid={`guild-loot-award-${item.item_id}-${candidate.character_key}`}
                    >
                      Award
                    </button>
                  {/if}
                </span>
              </div>
            {/each}
          </div>
        </div>
      {/each}
    </div>
    <p class="text-muted text-[12px]">
      Drops are the raid's own real loot table; ranked gain reuses the planner's own gain rule. "Awarded"
      needs a new table this round proposes — a drop stays unawarded until an officer acts.
    </p>
  {/if}
</section>
