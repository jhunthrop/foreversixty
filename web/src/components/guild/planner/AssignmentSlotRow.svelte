<!-- web/src/components/guild/planner/AssignmentSlotRow.svelte -->
<!-- One assignment slot (web layout spec 2026-11-05 §7.2): a Row with the slot's label, how full it
     is and what it does, its names as Tags under it (the viewer's own in gold and the row selected,
     an off-role pick in amber). An officer's Edit opens the picker under the row: the signed
     raiders who fit the slot first, a Remove per name and Clear. Manual only. -->
<script lang="ts">
  import type { AssignmentRaider, AssignmentSlot } from '../../../lib/guild/assignments-api';
  import {
    assignmentsCopy,
    holdsSlot,
    isFull,
    nameTone,
    pickerOptions,
    slotSub,
  } from '../../../lib/guild/assignments-view';
  import { ROW_ITEM } from '../../../lib/ui/primitives';
  import Button from '../../ui/Button.svelte';
  import Row from '../../ui/Row.svelte';
  import Select from '../../ui/Select.svelte';
  import Tag from '../../ui/Tag.svelte';

  let {
    id,
    slot,
    candidates,
    officer,
    editable,
    editing,
    busy,
    line,
    onToggle,
    onAdd,
    onRemove,
    onClear,
  }: {
    /** Unique per slot on the page: the picker's id and the test ids. */
    id: string;
    slot: AssignmentSlot;
    candidates: readonly AssignmentRaider[];
    officer: boolean;
    /** Officer tools on and the raid still open. */
    editable: boolean;
    editing: boolean;
    busy: boolean;
    line: string | null;
    onToggle: () => void;
    onAdd: (key: string) => void;
    onRemove: (key: string) => void;
    onClear: () => void;
  } = $props();

  const options = $derived(pickerOptions(slot, candidates));
  const full = $derived(isFull(slot));
  let picked = $state('');

  function choose(event: Event): void {
    const key = (event.currentTarget as HTMLSelectElement).value;
    picked = '';
    if (key !== '') onAdd(key);
  }
</script>

<li class={ROW_ITEM} data-testid={`assign-slot-${slot.key}`} data-own={holdsSlot(slot) ? 'true' : undefined}>
  <Row title={slot.label} sub={slotSub(slot)} selected={holdsSlot(slot)} open>
    {#snippet trail()}
      {#if officer}
        <Button
          size="sm"
          tone="ghost"
          disabled={!editable}
          onclick={onToggle}
          testid={`assign-edit-${slot.key}`}>{editing ? assignmentsCopy.done : assignmentsCopy.edit}</Button
        >
      {/if}
    {/snippet}
    {#snippet foot()}
      {#if slot.names.length === 0}
        <span class="text-muted text-[13px]" data-testid={`assign-open-${slot.key}`}
          >{assignmentsCopy.open}</span
        >
      {:else}
        {#each slot.names as name (name.character_key)}
          <Tag
            tone={nameTone(name)}
            title={name.off_role ? assignmentsCopy.offRole(name.name) : undefined}
            testid={`assign-name-${slot.key}`}>{name.name}</Tag
          >
        {/each}
      {/if}
    {/snippet}
  </Row>
  {#if officer && editing}
    <div class="flex flex-col gap-2 pt-2 pb-3" data-testid={`assign-picker-${slot.key}`}>
      <Select
        id={`assign-${id}`}
        bind:value={picked}
        {options}
        size="compact"
        label={assignmentsCopy.addLabel(slot.label)}
        placeholder={full ? assignmentsCopy.full : assignmentsCopy.add}
        emptyLabel={assignmentsCopy.noOneSigned}
        disabled={!editable || busy || full}
        {busy}
        testid={`assign-add-${slot.key}`}
        onchange={choose}
      />
      {#if slot.names.length > 0}
        <div class="flex flex-wrap gap-2">
          {#each slot.names as name (name.character_key)}
            <Button
              size="sm"
              tone="ghost"
              disabled={!editable || busy}
              onclick={() => onRemove(name.character_key)}
              testid={`assign-remove-${slot.key}`}>{assignmentsCopy.remove(name.name)}</Button
            >
          {/each}
          <Button
            size="sm"
            tone="danger"
            disabled={!editable || busy}
            onclick={onClear}
            testid={`assign-clear-${slot.key}`}>{assignmentsCopy.clear}</Button
          >
        </div>
      {/if}
      {#if line !== null}
        <p class="text-wipe text-[13px]" role="alert" data-testid={`assign-line-${slot.key}`}>{line}</p>
      {/if}
    </div>
  {/if}
</li>
