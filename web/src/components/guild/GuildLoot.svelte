<!-- web/src/components/guild/GuildLoot.svelte -->
<!-- Guild control-centre spec §4.F/§12.1: the loot council helper. Boss picker defaults to
     the next unkilled named boss; a boss that is down reads "on farm." Each drop shows up to
     four ranked candidates. An awarded item says who got it, when and from whom once, on the
     item itself; every other candidate row reads "Already awarded" and has no button, and an
     unawarded item keeps the Award button on every row (§12.1's round-2 fix: one awardee per
     item, not one per candidate row). Member sees every drop and ranking read-only (§10.1,
     proposed); only officer sees the Award control. Award asks first ("Award Stormrage Cover to
     Rootbound?" with Award and Cancel, focus on Cancel) and an awarded item offers the officer
     Undo, which removes the award (lifecycle round 6, R2). Officer tools gate §3/§4: with the
     guild's officer tools off (officer-read-only) the plan sentence leads the tab, the controls
     stay on screen disabled, and a server refusal of a click says so once, in the item's card. -->
<script lang="ts">
  import ClassCrestRing from '../character/ClassCrestRing.svelte';
  import { classColorVar } from '../../lib/report/format';
  import type { GuildLootItem, GuildLootPage } from '../../lib/guild/api';
  import type { AttendanceSource } from '../../lib/guild/attendance-contract';
  import { guildGateCopy } from '../../lib/guild/copy';
  import {
    LOOT_FOOTNOTE,
    NOT_SIMMED_NOTE,
    NO_CANDIDATE,
    UNDO_AWARD,
    anyNotSimmed,
    awardQuestion,
    awardedLine,
    bossPickerLabel,
    candidateAction,
    candidateGainLabel,
    candidateLine,
  } from '../../lib/guild/loot-view';
  import { PRIMARY_BUTTON_FIXED } from '../../lib/planner/styles';
  import EmptyState from '../ui/EmptyState.svelte';
  import type { GuildPlanRefusal } from '../../lib/guild/officer-tools';
  import GuildPlanGate from './GuildPlanGate.svelte';

  let {
    loot,
    officer,
    officerTools,
    raidGuildTools,
    memberReadOnly,
    attendanceSource = 'reports',
    busyItemId,
    refusal,
    settingsHref,
    onAward,
    onUndo,
  }: {
    loot: GuildLootPage | null;
    officer: boolean;
    /** False when the guild's officer tools are off (gate spec §3.1, officer-read-only). */
    officerTools: boolean;
    /** False for an officer of a Guild-tier guild: Loot is part of Raid guild (pricing spec §3.2.8). */
    raidGuildTools: boolean;
    memberReadOnly: boolean;
    /** `raids`: the candidate line counts raids (`7 of 8 raids`). */
    attendanceSource?: AttendanceSource;
    busyItemId: number | null;
    /** The Award or Undo the server refused with `guild_plan_required` (§4), per boss and item. */
    refusal: GuildPlanRefusal | null;
    /** Where the plan sentence's link goes: the guild's settings tab. */
    settingsHref: string;
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
  /** Focus lands on Cancel, so a stray Enter backs out rather than awards; the §4 refused
   *  line takes focus the same way, since the button that held it is now disabled. */
  function focusOnMount(node: HTMLElement): void {
    node.focus();
  }

  /** The refusal that already took focus: switching boss and back remounts the line, not the focus. */
  let focusedRefusal: GuildPlanRefusal | null = null;
  function focusRefusalOnce(node: HTMLElement): void {
    if (focusedRefusal === refusal) return;
    focusedRefusal = refusal;
    node.focus();
  }
  const refusedHere = $derived(
    refusal !== null && loot !== null && loot.selected === refusal.encounterId ? refusal : null,
  );

  const selectedEncounter = $derived(loot?.encounters.find((e) => e.encounter_id === loot.selected));
  const notSimmed = $derived(loot !== null && anyNotSimmed(loot.items));
  const officerReadOnly = $derived(officer && !(officerTools && raidGuildTools));
  /** Native disabled plus the sentence as description (§3.4); undefined leaves no attribute. */
  const gateDescribedBy = $derived(officerReadOnly ? 'officer-tools-gate-loot' : undefined);
  const refusedLine = $derived(
    refusedHere?.kind === 'undo' ? guildGateCopy.undoRefused : guildGateCopy.awardRefused,
  );
</script>

<section class="flex flex-col gap-4" data-testid="guild-loot-tab">
  <div class="flex items-center justify-between gap-3">
    <h2 class="section-title">Loot</h2>
    {#if memberReadOnly || officerReadOnly}<span class="text-muted text-[12px]">{guildGateCopy.readOnly}</span
      >{/if}
  </div>
  {#if officerReadOnly}<GuildPlanGate tab="loot" {settingsHref} raidGuild={officerTools} />{/if}

  {#if loot === null || loot.items.length === 0}
    <EmptyState
      message="No bosses killed yet this tier — loot ranking starts after the first kill"
      testid="guild-loot-empty"
    />
  {:else}
    <div class="flex items-center gap-3" data-testid="guild-loot-boss-picker">
      <span class="quiet-label">Boss</span>
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
          <!-- Below lg the awarded line is its own full-width row under the title (a chip that
               grows with its text); at lg it sits right of it. Undo is the winner row's alone. -->
          <div class="flex flex-col gap-2 lg:flex-row lg:items-start lg:justify-between lg:gap-3">
            <span class="heading-name min-w-0 text-[15px]"
              >{item.name}{#if (item.copies ?? 1) > 1}
                <span class="text-muted text-[12px]" data-testid={`guild-loot-copies-${item.item_id}`}
                  >x{item.copies}</span
                >{/if}</span
            >
            {#if item.awarded_to !== null}
              <div class="flex min-w-0 flex-wrap items-center gap-3 lg:justify-end">
                <span
                  class="min-w-0 basis-full rounded-md border border-[rgba(30,255,0,.25)] bg-[rgba(30,255,0,.10)] px-3 py-2 text-[14px] leading-snug text-[#7bff5c] lg:basis-auto"
                  data-testid={`guild-loot-awarded-${item.item_id}`}>{awardedLine(item.awarded_to)}</span
                >
              </div>
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
              <div class="border-line-soft flex flex-wrap items-center gap-x-3 gap-y-1 border-b py-1.5">
                <ClassCrestRing characterClass={candidate.class} size={26} />
                <span
                  class="min-w-[120px] text-[13px] font-semibold"
                  style={`color:${classColorVar(candidate.class)}`}
                >
                  {candidate.name}
                </span>
                <span
                  class="min-w-0 basis-full text-[13px] sm:basis-auto"
                  data-testid="guild-loot-candidate-line"
                >
                  {candidateLine(candidate, attendanceSource)}
                  <!-- One unbreakable chunk, and nothing at all for an empty slot. -->
                  {#if candidateGainLabel(candidate) !== ''}<span
                      class="text-muted numeric text-[11px] whitespace-nowrap"
                      data-testid="guild-loot-candidate-gain">{candidateGainLabel(candidate)}</span
                    >{/if}
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
                    <span class="text-muted text-[12px]" data-testid="guild-loot-winner-slot">
                      Awarded{#if officer}
                        &nbsp;·
                        <button
                          class="inline-flex h-11 items-center underline disabled:cursor-default disabled:no-underline lg:h-7"
                          onclick={() => onUndo(item)}
                          disabled={officerReadOnly || busyItemId === item.item_id}
                          aria-describedby={gateDescribedBy}
                          aria-label={`Undo the award of ${item.name} to ${candidate.name}`}
                          data-testid={`guild-loot-row-undo-${item.item_id}`}>{UNDO_AWARD}</button
                        >{/if}
                    </span>
                  {:else if action.kind === 'awarded-to-other'}
                    <span class="text-muted text-[12px]">{action.label}</span>
                  {:else if officer}
                    <button
                      class="rounded-control text-strong inline-flex h-11 items-center justify-center border border-[#3a3326] bg-none px-3 text-[11px] font-bold tracking-[0.06em] uppercase disabled:cursor-default disabled:opacity-50 md:h-7"
                      onclick={() => (confirming = { itemId: item.item_id, key: candidate.character_key })}
                      disabled={officerReadOnly || busyItemId === item.item_id || confirming !== null}
                      aria-describedby={gateDescribedBy}
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
              {#if officer && !officerReadOnly && confirming !== null && confirming.itemId === item.item_id && confirming.key === candidate.character_key}
                <div
                  class="border-line-soft bg-surface flex flex-wrap items-center gap-x-4 gap-y-2 border-b py-2 pl-[38px]"
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
                    class="{PRIMARY_BUTTON_FIXED} justify-center px-4 md:h-7"
                    onclick={() => confirmAward(item, candidate.character_key)}
                    data-testid={`guild-loot-award-yes-${item.item_id}`}
                  >
                    Award
                  </button>
                  <button
                    class="rounded-control text-muted inline-flex h-11 items-center justify-center bg-none px-4 text-[11px] font-bold tracking-[0.06em] uppercase focus:outline-2 focus:outline-offset-2 focus:outline-[var(--color-gold)] md:h-7"
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
          {#if refusedHere !== null && refusedHere.itemId === item.item_id}
            <p
              class="text-wipe text-[13px] leading-5"
              role="alert"
              tabindex="-1"
              use:focusRefusalOnce
              data-testid={`guild-loot-plan-refused-${item.item_id}`}
            >
              {refusedLine}
            </p>
          {/if}
        </div>
      {/each}
    </div>
    <p class="text-muted text-[12px]">{LOOT_FOOTNOTE}</p>
  {/if}
</section>
