<!-- web/src/components/guild/planner/PlannerAssignments.svelte -->
<!-- The raid planner's Assignments (web layout spec 2026-11-05 §7.2): one Section per encounter of
     the raid, each slot a Row (AssignmentSlotRow), then Copy as raid warning (the /rw lines) and,
     for officers, Post to Discord (one message under the raid's post). Members read it, their own
     slots highlighted. A raid whose bosses the client does not carry yet says so. Manual only:
     the raid leader picks every name. -->
<script lang="ts">
  import { onDestroy } from 'svelte';
  import type { GuildRaidView } from '../../../lib/guild/api';
  import type { AssignmentActions } from '../../../lib/guild/assignment-actions.svelte';
  import { slotId } from '../../../lib/guild/assignment-actions.svelte';
  import type { EncounterAssignments } from '../../../lib/guild/assignments-api';
  import {
    anyPicked,
    assignmentsCopy,
    encountersOf,
    notInClientOf,
    ownJobs,
    raidWarningText,
  } from '../../../lib/guild/assignments-view';
  import { createCopyFeedback, type CopyFeedbackState } from '../../../lib/ui/copy-feedback';
  import Button from '../../ui/Button.svelte';
  import Section from '../../ui/Section.svelte';
  import Tag from '../../ui/Tag.svelte';
  import AssignmentSlotRow from './AssignmentSlotRow.svelte';

  let {
    raid,
    officer,
    officerTools,
    closed = false,
    actions,
  }: {
    raid: GuildRaidView;
    officer: boolean;
    officerTools: boolean;
    /** Cancelled or ended: the picks read, nothing writes. */
    closed?: boolean;
    actions: AssignmentActions;
  } = $props();

  const encounters = $derived(encountersOf(raid.assignments));
  const missing = $derived(notInClientOf(raid.assignments));
  const candidates = $derived(raid.assignments?.candidates ?? []);
  const editable = $derived(officer && officerTools && !closed);

  let editing = $state<string | null>(null);
  let copyState = $state<CopyFeedbackState>('idle');
  /** The encounter whose lines the failed copy shows, selectable. */
  let copyShown = $state<number | null>(null);

  const feedback = createCopyFeedback((next) => (copyState = next));
  onDestroy(() => feedback.dispose());

  function copyWarning(encounter: EncounterAssignments): void {
    copyShown = encounter.encounter_id;
    void feedback.copy(raidWarningText(encounter));
  }

  function toggle(id: string): void {
    editing = editing === id ? null : id;
  }

  function meta(encounter: EncounterAssignments): string | undefined {
    const jobs = ownJobs(encounter);
    return jobs.length === 0 ? undefined : `${assignmentsCopy.yours}: ${jobs.join(', ')}`;
  }
</script>

{#each missing as name (name)}
  <Section title={assignmentsCopy.title} panel testid="raid-assignments-missing">
    {#snippet head()}<Tag tone="muted" testid="raid-assignments-not-in-client"
        >{assignmentsCopy.notInClient}</Tag
      >{/snippet}
    <p class="text-muted text-[14px]">{assignmentsCopy.notInClientLine(name)}</p>
  </Section>
{/each}

{#each encounters as encounter (encounter.encounter_id)}
  <Section
    panel
    title={assignmentsCopy.encounterTitle(encounter.name)}
    meta={meta(encounter)}
    id={`raid-assignments-${encounter.encounter_id}`}
    testid="raid-assignments"
  >
    <ul>
      {#each encounter.slots as slot (slot.key)}
        {@const id = slotId(encounter.encounter_id, slot.key)}
        <AssignmentSlotRow
          id={`${encounter.encounter_id}-${slot.key}`}
          {slot}
          {candidates}
          {officer}
          {editable}
          editing={editing === id}
          busy={actions.busySlot === id}
          line={actions.slotLine?.slot === id ? actions.slotLine.message : null}
          onToggle={() => toggle(id)}
          onAdd={(key) => void actions.add(encounter.encounter_id, slot, key)}
          onRemove={(key) => void actions.remove(encounter.encounter_id, slot, key)}
          onClear={() => void actions.clear(encounter.encounter_id, slot)}
        />
      {/each}
    </ul>
    <div class="flex flex-wrap items-center gap-2 pt-3">
      <Button
        size="sm"
        disabled={!anyPicked(encounter)}
        onclick={() => copyWarning(encounter)}
        testid="raid-assignments-copy"
        >{copyState === 'copied' && copyShown === encounter.encounter_id
          ? assignmentsCopy.copied
          : assignmentsCopy.copyWarning}</Button
      >
      {#if officer}
        <Button
          size="sm"
          disabled={!editable || !anyPicked(encounter) || actions.posting}
          onclick={() => void actions.post()}
          testid="raid-assignments-post"
          >{actions.posting ? assignmentsCopy.posting : assignmentsCopy.post}</Button
        >
      {/if}
    </div>
    <span class="sr-only" aria-live="polite">
      {copyState === 'copied' && copyShown === encounter.encounter_id ? assignmentsCopy.copied : ''}
    </span>
    {#if officer && actions.postAnswer !== null}
      <p class="text-muted pt-2 text-[13px]" role="status" data-testid="raid-assignments-post-answer">
        {actions.postAnswer}
      </p>
    {/if}
    {#if copyState === 'failed' && copyShown === encounter.encounter_id}
      <p class="text-wipe pt-2 text-[13px]" role="alert">{assignmentsCopy.copyFailed}</p>
      <pre
        class="text-strong pt-1 text-[13px] whitespace-pre-wrap select-all"
        data-testid="raid-assignments-warning-text">{raidWarningText(encounter)}</pre>
    {/if}
  </Section>
{/each}
