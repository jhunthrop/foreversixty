<!-- web/src/components/guild/GuildLoot.svelte -->
<!-- Guild loot spec 2026-11-05, on the control-centre spec §4.F/§12.1: the loot council helper. Each
     drop anyone can take is a card with up to five ranked candidates, then Show {n} more: tier 0 (on
     their BiS list) first, tier 1 (item level only) under its own label. The rest collapse into one
     line (LootUnranked). An awarded item says who got it, when and from whom once, on the item
     itself; every other candidate row reads "Already awarded" and has no button. Member sees every
     drop and ranking read-only (§10.1, proposed); only an officer sees the Award control. Award asks
     first ("Award Stormrage Cover to Rootbound?" with Award and Cancel, focus on Cancel) and an
     awarded item offers the officer Undo (lifecycle round 6, R2). Officer tools gate §3/§4: with the
     guild's officer tools off (officer-read-only) the plan sentence leads the tab, the controls stay
     on screen disabled, and a server refusal of a click says so once, in the item's card. -->
<script lang="ts">
  import type { GuildLootItem, GuildLootPage } from '../../lib/guild/api';
  import type { AttendanceSource, RaidLootAward } from '../../lib/guild/attendance-contract';
  import { guildGateCopy } from '../../lib/guild/copy';
  import type { GateFacts } from '../../lib/guild/plan-gate';
  import {
    LOOT_FOOTNOTE,
    NOBODY_AT_60,
    NO_CANDIDATES_AT_60,
    TIER_ONE_LABEL,
    TIER_ONE_SHORT,
    UNDO_AWARD,
    awardQuestion,
    awardedLine,
    bossPickerLabel,
    candidateAction,
    rankedDrops,
    showsAttendance,
    unrankedDrops,
  } from '../../lib/guild/loot-view';
  import { slotLine } from '../../lib/items/slot-words';
  import { button } from '../../lib/ui/button';
  import Card from '../ui/Card.svelte';
  import EmptyState from '../ui/EmptyState.svelte';
  import Select from '../ui/Select.svelte';
  import LootCandidateRow from './LootCandidateRow.svelte';
  import LootUnranked from './LootUnranked.svelte';
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
    ownKeys = [],
    busyItemId,
    refusal,
    gateFacts,
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
    /** The viewer's own characters: their rows carry the wash and `(you)`. */
    ownKeys?: readonly string[];
    busyItemId: number | null;
    /** The Award or Undo the server refused with `guild_plan_required` (§4), per boss and item. */
    refusal: GuildPlanRefusal | null;
    /** Where the plan sentence's link goes: the guild's settings tab. */
    /** The gate sentence's facts (spec 2026-11-05 §3.2). */
    gateFacts: GateFacts;
    /** `itemId` names the copy when it is not the row's own (a side's Head of Onyxia). */
    onAward: (item: GuildLootItem, characterKey: string, itemId?: number) => void;
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
  const ranked = $derived(loot === null ? [] : rankedDrops(loot.items));
  const unranked = $derived(loot === null ? [] : unrankedDrops(loot.items));
  const showAttendance = $derived(loot !== null && showsAttendance(loot.items));
  /** Candidates shown before `Show {n} more`. */
  const FIRST_ROWS = 5;
  /** The cards whose `Show {n} more` was pressed. */
  let expanded = $state<ReadonlySet<number>>(new Set());
  function showAll(itemId: number): void {
    expanded = new Set([...expanded, itemId]);
  }
  const shownRows = (item: GuildLootItem) =>
    expanded.has(item.item_id) ? item.candidates : item.candidates.slice(0, FIRST_ROWS);
  /** The first card showing a tier-1 row says the label in full; every later card marks its tier-1
   *  group with the short label, so the sentence is read once per tab (guild round 2). */
  const fullTierOneItem = $derived(
    ranked.find((item) => shownRows(item).some((candidate) => candidate.tier === 1))?.item_id,
  );
  const officerReadOnly = $derived(officer && !(officerTools && raidGuildTools));
  /** Native disabled plus the sentence as description (§3.4); undefined leaves no attribute. */
  const gateDescribedBy = $derived(officerReadOnly ? 'officer-tools-gate-loot' : undefined);
  const refusedLine = $derived(
    refusedHere?.kind === 'undo' ? guildGateCopy.undoRefused : guildGateCopy.awardRefused,
  );
</script>

{#snippet itemIcon(item: GuildLootItem)}
  <RaidLootItem award={awardOf(item)} iconSize={40} nameSize="" show="icon" />
{/snippet}

{#snippet itemName(item: GuildLootItem)}
  <RaidLootItem award={awardOf(item)} iconSize={40} nameSize="text-[24px]" show="name" />
{/snippet}

<section class="flex flex-col gap-[22px] lg:gap-6" data-testid="guild-loot-tab">
  {#if officerReadOnly}<GuildPlanGate tab="loot" facts={gateFacts} />{/if}

  <Card glyph="star" title="Loot" meta={memberReadOnly ? guildGateCopy.readOnly : undefined} body="none">
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
    {#if ranked.length === 0}
      <div class="border-line-soft border-t px-5 py-4">
        <EmptyState
          message={loot === null || loot.at_level_cap === 0 ? NOBODY_AT_60 : NO_CANDIDATES_AT_60}
          testid="guild-loot-empty"
        />
      </div>
    {/if}
  </Card>

  {#if ranked.length > 0}
    <div class="flex flex-col gap-[22px] lg:gap-6" data-testid="guild-loot-items">
      {#each ranked as item (item.item_id)}
        {@const firstTierOne = item.candidates.findIndex((candidate) => candidate.tier === 1)}
        {@const rows = shownRows(item)}
        <Card
          label={slotLine(item.inventory_type, item.item_class, item.item_subclass)}
          title={item.name}
          body="none"
          testid={`guild-loot-item-${item.item_id}`}
        >
          {#snippet lead()}{@render itemIcon(item)}{/snippet}
          {#snippet titleContent()}{@render itemName(item)}{/snippet}
          {#snippet aside()}
            {#if item.awarded_to !== null}
              <span
                class="inline-block rounded-md border border-[rgba(30,255,0,.25)] bg-[rgba(30,255,0,.10)] px-3 py-2 text-left text-[14px] leading-snug text-[#7bff5c]"
                data-testid={`guild-loot-awarded-${item.item_id}`}>{awardedLine(item.awarded_to)}</span
              >
            {/if}
          {/snippet}
          <div class="flex flex-col">
            {#each rows as candidate, index (candidate.character_key)}
              {@const awardState = candidateAction(item, candidate.character_key)}
              {#if index === firstTierOne}
                <p
                  class="border-line-soft text-muted border-t px-5 py-2 text-[12px]"
                  data-testid={`guild-loot-tier-one-${item.item_id}`}
                >
                  {item.item_id === fullTierOneItem ? TIER_ONE_LABEL : TIER_ONE_SHORT}
                </p>
              {/if}
              <LootCandidateRow
                {candidate}
                own={ownKeys.includes(candidate.character_key)}
                {showAttendance}
                {attendanceSource}
              >
                {#snippet action()}
                  {#if awardState.kind === 'awarded'}
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
                  {:else if awardState.kind === 'awarded-to-other'}
                    <span class="text-muted text-[12px]">{awardState.label}</span>
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
                {/snippet}
              </LootCandidateRow>
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
            {#if rows.length < item.candidates.length}
              <button
                class="border-line-soft min-h-11 border-t px-5 text-left text-[13px] underline"
                onclick={() => showAll(item.item_id)}
                data-testid={`guild-loot-show-more-${item.item_id}`}
              >
                Show {item.candidates.length - rows.length} more
              </button>
            {/if}
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

  {#if unranked.length > 0}
    <LootUnranked
      items={unranked}
      {officer}
      readOnly={officerReadOnly}
      {gateDescribedBy}
      {busyItemId}
      {onAward}
      {onUndo}
    />
  {/if}

  {#if loot !== null}
    <p class="text-muted max-w-[720px] text-[13px]" data-testid="guild-loot-footnote">{LOOT_FOOTNOTE}</p>
  {/if}
</section>
