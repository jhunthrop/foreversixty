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
  import type { AttendanceSource, RaidLootAward } from '../../lib/guild/attendance-contract';
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
  import { button } from '../../lib/ui/button';
  import Card from '../ui/Card.svelte';
  import Select from '../ui/Select.svelte';
  import RaidLootItem from './RaidLootItem.svelte';
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
    onBoss,
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
    /** The boss chooser changed: read that encounter's loot. */
    onBoss: (encounterId: number) => void;
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

  const bossOptions = $derived(
    (loot?.encounters ?? []).map((e) => ({ value: String(e.encounter_id), label: bossPickerLabel(e) })),
  );
  /** The item as the shared item hover reads it: the candidates' class opens the tooltip's class file. */
  const awardOf = (item: GuildLootItem): RaidLootAward => ({
    item_id: item.item_id,
    item_name: item.name,
    quality: item.quality,
    icon: item.icon,
    character_key: '',
    name: '',
    class: item.candidates[0]?.class ?? '',
    by_name: '',
    at: '',
  });
  const notSimmed = $derived(loot !== null && anyNotSimmed(loot.items));
  const officerReadOnly = $derived(officer && !(officerTools && raidGuildTools));
  /** Native disabled plus the sentence as description (§3.4); undefined leaves no attribute. */
  const gateDescribedBy = $derived(officerReadOnly ? 'officer-tools-gate-loot' : undefined);
  const refusedLine = $derived(
    refusedHere?.kind === 'undo' ? guildGateCopy.undoRefused : guildGateCopy.awardRefused,
  );
</script>

{#snippet itemIcon(item: GuildLootItem)}
  <RaidLootItem award={awardOf(item)} iconSize={36} nameSize="" show="icon" />
{/snippet}

{#snippet itemName(item: GuildLootItem)}
  <RaidLootItem award={awardOf(item)} iconSize={36} nameSize="text-[24px]" show="name" />
{/snippet}

<section class="flex flex-col gap-[22px] lg:gap-6" data-testid="guild-loot-tab">
  {#if officerReadOnly}<GuildPlanGate tab="loot" {settingsHref} raidGuild={officerTools} />{/if}

  <Card
    glyph="star"
    title="Loot"
    meta={memberReadOnly || officerReadOnly ? guildGateCopy.readOnly : undefined}
    body="none"
  >
    {#snippet aside()}
      {#if loot !== null && loot.items.length > 0}
        <div class="w-56 text-left" data-testid="guild-loot-boss-picker">
          <Select
            id="guild-loot-boss"
            size="compact"
            label="Boss"
            options={bossOptions}
            value={String(loot.selected ?? '')}
            onchange={(event) => onBoss(Number((event.currentTarget as HTMLSelectElement).value))}
          />
        </div>
      {/if}
    {/snippet}
    {#if loot === null || loot.items.length === 0}
      <p class="border-line-soft text-muted border-t px-5 py-4 text-[14px]" data-testid="guild-loot-empty">
        No bosses killed yet this tier — loot ranking starts after the first kill
      </p>
    {:else if notSimmed}
      <p
        class="border-line-soft text-muted border-t px-5 py-4 text-[13px]"
        data-testid="guild-loot-not-simmed"
      >
        {NOT_SIMMED_NOTE}
      </p>
    {/if}
    {#if loot !== null && loot.items.length > 0}
      <div class="border-line-soft text-muted border-t px-5 py-4 text-[13px]">{LOOT_FOOTNOTE}</div>
    {/if}
  </Card>

  {#if loot !== null && loot.items.length > 0}
    <div class="flex flex-col gap-[22px] lg:gap-6" data-testid="guild-loot-items">
      {#each loot.items as item (item.item_id)}
        <Card
          label={item.slot === '' ? undefined : item.slot}
          title={item.name}
          body="none"
          testid={`guild-loot-item-${item.item_id}`}
        >
          {#snippet lead()}{@render itemIcon(item)}{/snippet}
          {#snippet titleContent()}{@render itemName(item)}{/snippet}
          {#snippet titleSnippet()}
            {#if (item.copies ?? 1) > 1}
              <span class="text-muted text-[12px]" data-testid={`guild-loot-copies-${item.item_id}`}
                >x{item.copies}</span
              >
            {/if}
          {/snippet}
          {#snippet aside()}
            {#if item.awarded_to !== null}
              <span
                class="inline-block rounded-md border border-[rgba(30,255,0,.25)] bg-[rgba(30,255,0,.10)] px-3 py-2 text-left text-[14px] leading-snug text-[#7bff5c]"
                data-testid={`guild-loot-awarded-${item.item_id}`}>{awardedLine(item.awarded_to)}</span
              >
            {/if}
          {/snippet}
          {#if item.candidates.length === 0}
            <p
              class="border-line-soft text-muted border-t px-5 py-4 text-[13px]"
              data-testid={`guild-loot-no-candidate-${item.item_id}`}
            >
              {NO_CANDIDATE}
            </p>
          {/if}
          <div class="flex flex-col">
            {#each item.candidates as candidate (candidate.character_key)}
              {@const action = candidateAction(item, candidate.character_key)}
              <div
                class="border-line-soft flex min-h-14 flex-wrap items-center gap-x-3 gap-y-1 border-t px-5 py-2"
              >
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
                      class={button('secondary', 'sm')}
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
                  class="border-line-soft bg-surface flex flex-wrap items-center gap-x-4 gap-y-2 border-t px-5 py-3"
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
                    class={button('primary', 'sm')}
                    onclick={() => confirmAward(item, candidate.character_key)}
                    data-testid={`guild-loot-award-yes-${item.item_id}`}
                  >
                    Award
                  </button>
                  <button
                    class="{button('ghost', 'sm')} focus:outline-gold focus:outline-2 focus:outline-offset-2"
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
              class="border-line-soft text-wipe border-t px-5 py-3 text-[13px] leading-5"
              role="alert"
              tabindex="-1"
              use:focusRefusalOnce
              data-testid={`guild-loot-plan-refused-${item.item_id}`}
            >
              {refusedLine}
            </p>
          {/if}
        </Card>
      {/each}
    </div>
  {/if}
</section>
