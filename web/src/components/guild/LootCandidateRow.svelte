<!-- web/src/components/guild/LootCandidateRow.svelte -->
<!-- One candidate on a loot card (guild loot spec 2026-11-05 §3.3): who, what the drop replaces, the
     gain, attendance and the officer's action, in that order from 768. Below it the row stacks:
     crest, name and gain, then Replaces, then attendance and the action, 120 px at most with the
     44 px Award (6 px padding, no row gap). Attendance is left out, cell and column, while the guild
     has no counted raid night. -->
<script lang="ts">
  import type { Snippet } from 'svelte';
  import type { GuildLootCandidate } from '../../lib/guild/api';
  import {
    EMPTY_SLOT,
    attendanceLabel,
    candidateGainLabel,
    candidateSubline,
  } from '../../lib/guild/loot-view';
  import { classColorVar } from '../../lib/report/format';
  import { ROW_BAND_OWN } from '../../lib/ui/row';
  import ClassCrestRing from '../character/ClassCrestRing.svelte';
  import RaidLootItem from './RaidLootItem.svelte';

  let {
    candidate,
    own,
    showAttendance,
    action,
  }: {
    candidate: GuildLootCandidate;
    /** The viewer's own row: the own-row wash and `(you)`. */
    own: boolean;
    showAttendance: boolean;
    /** The officer's Award or the awarded state. */
    action: Snippet;
  } = $props();

  /** Readiness's own-row suffix, space included: a span's leading space is trimmed by Svelte. */
  const OWN_SUFFIX = ' (you)';

  const grid = $derived(
    showAttendance
      ? 'md:grid-cols-[minmax(160px,1.2fr)_minmax(0,1.6fr)_120px_112px_auto]'
      : 'md:grid-cols-[minmax(160px,1.2fr)_minmax(0,1.6fr)_120px_auto]',
  );
</script>

<div
  class="border-line-soft hover:bg-row-hover grid min-h-11 grid-cols-[minmax(0,1fr)_auto] items-center gap-x-4 gap-y-0 border-t px-5 py-1.5 md:py-2 {grid} {own
    ? ROW_BAND_OWN
    : ''}"
  data-testid="guild-loot-candidate"
>
  <div class="flex min-w-0 items-center gap-2">
    <ClassCrestRing characterClass={candidate.class} size={24} />
    <div class="min-w-0">
      <p class="truncate text-[14px] font-semibold" style={`color:${classColorVar(candidate.class)}`}>
        {candidate.name}{#if own}<span class="text-muted font-normal">{OWN_SUFFIX}</span>{/if}
      </p>
      <p class="text-muted truncate text-[12px]" data-testid="guild-loot-candidate-spec">
        {candidateSubline(candidate)}
      </p>
    </div>
  </div>
  <div
    class="col-span-2 row-start-2 min-w-0 text-[13px] md:col-span-1 md:row-start-auto"
    data-testid="guild-loot-candidate-replaces"
  >
    {#if candidate.replaces_item !== null}
      <span class="inline-flex max-w-full flex-wrap items-center gap-x-1">
        <span>Replaces</span>
        <RaidLootItem
          award={{
            item_id: candidate.replaces_item.item_id,
            item_name: candidate.replaces_item.name,
            quality: candidate.replaces_item.quality,
            icon: candidate.replaces_item.icon,
            character_key: '',
            name: '',
            class: candidate.class,
            by_name: '',
            at: '',
          }}
          iconSize={16}
          nameSize="text-[13px]"
        />
      </span>
    {:else if candidate.fills_empty_slot !== true}
      <!-- An empty slot the gain line already names ("Fills an empty slot") leaves this cell blank. -->
      <span class="text-muted">{EMPTY_SLOT}</span>
    {/if}
  </div>
  <div
    class="col-start-2 row-start-1 text-right whitespace-nowrap md:col-start-auto md:row-start-auto md:text-left {candidate.tier ===
    0
      ? 'numeric text-[14px]'
      : 'numeric text-muted text-[13px]'}"
    data-testid="guild-loot-candidate-gain"
  >
    {candidateGainLabel(candidate)}
  </div>
  {#if showAttendance}
    <div
      class="text-muted row-start-3 text-[13px] md:row-start-auto"
      data-testid="guild-loot-candidate-attendance"
    >
      {attendanceLabel(candidate.attendance.present, candidate.attendance.nights)}
    </div>
  {/if}
  <div class="col-start-2 row-start-3 justify-self-end md:col-start-auto md:row-start-auto">
    {@render action()}
  </div>
</div>
