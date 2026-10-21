<!-- web/src/components/guild/GuildLoot.svelte -->
<!-- Guild control-centre spec §4.F/§12.1: the loot council helper. Boss picker defaults to
     the next unkilled named boss; a boss that is down reads "on farm." Each drop shows up to
     four ranked candidates. An awarded item says who got it, when and from whom once, on the
     item itself; every other candidate row reads "Already awarded" and has no button, and an
     unawarded item keeps the Award button on every row (§12.1's round-2 fix: one awardee per
     item, not one per candidate row). Member sees every drop and ranking read-only (§10.1,
     proposed); only officer sees the Award control. Award asks first ("Award Stormrage Cover to
     Rootbound?" with Award and Cancel, focus on Cancel) and an awarded item offers the officer
     Undo, which removes the award (lifecycle round 6, R2). -->
<script lang="ts">
  import ClassCrestRing from '../character/ClassCrestRing.svelte';
  import { classColorVar } from '../../lib/report/format';
  import type { GuildLootItem, GuildLootPage } from '../../lib/guild/api';
  import {
    LOOT_FOOTNOTE,
    NOT_SIMMED_NOTE,
    NO_CANDIDATE,
    UNDO_AWARD,
    anyNotSimmed,
    attendanceLabel,
    awardQuestion,
    awardedLine,
    bossPickerLabel,
    candidateAction,
    candidateGainLabel,
  } from '../../lib/guild/loot-view';
  import EmptyState from '../ui/EmptyState.svelte';

  let {
    loot,
    officer,
    memberReadOnly,
    busyItemId,
    onAward,
    onUndo,
  }: {
    loot: GuildLootPage | null;
    officer: boolean;
    memberReadOnly: boolean;
    busyItemId: number | null;
    onAward: (item: GuildLootItem, characterKey: string) => void;
    /** Remove the item's award (the award's id is on `item.awarded_to`). */
    onUndo: (item: GuildLootItem) => void;
  } = $props();

  /** The candidate whose Award is asking first; one at a time. */
  let confirming = $state<{ itemId: number; key: string } | null>(null);

  function confirmAward(item: GuildLootItem, key: string): void {
    confirming = null;
    onAward(item, key);
  }
  /** Focus lands on Cancel, so a stray Enter backs out rather than awards. */
  function focusOnMount(node: HTMLElement): void {
    node.focus();
  }

  const selectedEncounter = $derived(loot?.encounters.find((e) => e.encounter_id === loot.selected));
  const notSimmed = $derived(loot !== null && anyNotSimmed(loot.items));
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

    {#if notSimmed}
      <p class="text-muted text-[12px]" data-testid="guild-loot-not-simmed">{NOT_SIMMED_NOTE}</p>
    {/if}

    <div class="flex flex-col gap-4" data-testid="guild-loot-items">
      {#each loot.items as item (item.item_id)}
        <div
          class="border-line-soft bg-raised flex flex-col gap-2 rounded-md border p-4"
          data-testid={`guild-loot-item-${item.item_id}`}
        >
          <div class="flex items-center justify-between gap-2">
            <span class="font-display text-[15px] font-bold"
              >{item.name}{#if (item.copies ?? 1) > 1}
                <span class="text-muted text-[12px]" data-testid={`guild-loot-copies-${item.item_id}`}
                  >x{item.copies}</span
                >{/if}</span
            >
            {#if item.awarded_to !== null}
              <span class="flex flex-wrap items-center justify-end gap-x-1">
                <span
                  class="pill"
                  style="color:#7bff5c;background:rgba(30,255,0,.10);border-color:rgba(30,255,0,.25)"
                  data-testid={`guild-loot-awarded-${item.item_id}`}>{awardedLine(item.awarded_to)}</span
                >
                {#if officer}
                  <span class="text-muted" aria-hidden="true">·</span>
                  <button
                    class="text-muted inline-flex h-11 items-center px-2 text-[12px] underline md:h-7"
                    onclick={() => onUndo(item)}
                    disabled={busyItemId === item.item_id}
                    aria-label={`Undo the award of ${item.name} to ${item.awarded_to.name}`}
                    data-testid={`guild-loot-undo-${item.item_id}`}
                  >
                    {UNDO_AWARD}
                  </button>
                {/if}
              </span>
            {/if}
          </div>
          {#if item.candidates.length === 0}
            <p class="text-muted text-[13px]" data-testid={`guild-loot-no-candidate-${item.item_id}`}>
              {NO_CANDIDATE}
            </p>
          {/if}
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
                <span class="tabular font-mono text-[12px]">{candidateGainLabel(candidate, item.slot)}</span>
                <span class="text-muted tabular font-mono text-[12px]">
                  {attendanceLabel(candidate.attendance.present, candidate.attendance.nights)}
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
                    <!-- The item header already says who got it, when and from whom. -->
                  {:else if action.kind === 'awarded-to-other'}
                    <span class="text-muted text-[12px]">{action.label}</span>
                  {:else if officer}
                    <button
                      class="rounded-control text-strong inline-flex h-11 items-center justify-center border border-[#3a3326] bg-none px-3 text-[11px] font-bold tracking-[0.06em] uppercase md:h-7"
                      onclick={() => (confirming = { itemId: item.item_id, key: candidate.character_key })}
                      disabled={busyItemId === item.item_id}
                      aria-label={`Award ${item.name} to ${candidate.name}`}
                      data-testid={`guild-loot-award-${item.item_id}-${candidate.character_key}`}
                    >
                      Award
                    </button>
                  {/if}
                </span>
              </div>
              <!-- Escape backs out of the confirm from either button; the group itself takes no action. -->
              <!-- svelte-ignore a11y_no_noninteractive_element_interactions -->
              {#if officer && confirming !== null && confirming.itemId === item.item_id && confirming.key === candidate.character_key}
                <div
                  class="border-line-soft bg-surface flex flex-wrap items-center gap-x-4 gap-y-2 border-b py-2"
                  role="group"
                  aria-label={awardQuestion(item.name, candidate.name)}
                  onkeydown={(event) => {
                    if (event.key === 'Escape') confirming = null;
                  }}
                  data-testid={`guild-loot-award-confirm-${item.item_id}`}
                >
                  <span class="basis-full text-[13px] md:basis-auto"
                    >{awardQuestion(item.name, candidate.name)}</span
                  >
                  <button
                    class="rounded-control text-strong inline-flex h-11 items-center justify-center border border-[#3a3326] bg-none px-4 text-[11px] font-bold tracking-[0.06em] uppercase md:h-7"
                    onclick={() => confirmAward(item, candidate.character_key)}
                    data-testid={`guild-loot-award-yes-${item.item_id}`}
                  >
                    Award
                  </button>
                  <button
                    class="rounded-control text-muted inline-flex h-11 items-center justify-center border border-[#3a3326] bg-none px-4 text-[11px] font-bold tracking-[0.06em] uppercase md:h-7"
                    onclick={() => (confirming = null)}
                    use:focusOnMount
                    data-testid={`guild-loot-award-cancel-${item.item_id}`}
                  >
                    Cancel
                  </button>
                </div>
              {/if}
            {/each}
          </div>
        </div>
      {/each}
    </div>
    <p class="text-muted text-[12px]">{LOOT_FOOTNOTE}</p>
  {/if}
</section>
