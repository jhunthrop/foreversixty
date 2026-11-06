<!-- web/src/components/guild/loot/AwardConfirm.svelte -->
<!-- The award's small confirm (web layout spec 2026-11-05 §7.2): the question, main or off spec
     (a gear drop only; it starts at the candidate's MS/OS), an optional reason of at most 200
     characters, then Award and Cancel. Focus lands on Cancel so a stray Enter backs out; Escape
     backs out from anywhere in it. -->
<script lang="ts">
  import type { MsOs } from '../../../lib/guild/loot-council-api';
  import { REASON_MAX, awardAsk, awardReason, reasonError } from '../../../lib/guild/loot-council-view';
  import Button from '../../ui/Button.svelte';
  import Field from '../../ui/Field.svelte';
  import Segmented from '../../ui/Segmented.svelte';

  let {
    itemName,
    name,
    gear,
    initial,
    busy,
    onConfirm,
    onCancel,
  }: {
    itemName: string;
    name: string;
    /** A slotted drop: the MS/OS choice shows. */
    gear: boolean;
    initial: MsOs | null;
    busy: boolean;
    onConfirm: (msOs: MsOs | null, reason: string | null) => void;
    onCancel: () => void;
  } = $props();

  const MS_OS_OPTIONS = [
    { value: 'ms', label: 'Main spec' },
    { value: 'os', label: 'Off spec' },
  ] as const;

  // The confirm opens fresh for each candidate, so the starting value is read once.
  // svelte-ignore state_referenced_locally
  let msOs = $state<MsOs>(initial ?? 'ms');
  let reason = $state('');
  const error = $derived(reasonError(reason));

  function focusOnMount(node: HTMLElement): void {
    node.querySelector('button')?.focus();
  }
</script>

<!-- svelte-ignore a11y_no_noninteractive_element_interactions -->
<div
  class="flex w-full flex-col gap-3"
  role="group"
  aria-label={awardAsk(itemName, name)}
  onkeydown={(event) => {
    if (event.key === 'Escape') onCancel();
  }}
  data-testid="guild-loot-award-confirm"
>
  <p class="text-strong text-[14px] leading-5 font-semibold">{awardAsk(itemName, name)}</p>
  {#if gear}
    <Segmented
      options={MS_OS_OPTIONS}
      value={msOs}
      label="Main or off spec"
      size="sm"
      onChange={(value) => (msOs = value)}
      testid="guild-loot-award-msos"
    />
  {/if}
  <Field
    id={`guild-loot-reason-${name}`}
    label="Reason (optional)"
    bind:value={reason}
    maxlength={REASON_MAX}
    placeholder="Best upgrade, full attendance"
    {error}
    testid="guild-loot-award-reason"
  />
  <div class="flex flex-wrap gap-3">
    <Button
      tone="primary"
      size="sm"
      disabled={busy || error !== ''}
      onclick={() => onConfirm(gear ? msOs : null, awardReason(reason))}
      testid="guild-loot-award-yes">Award</Button
    >
    <span class="contents" use:focusOnMount>
      <Button tone="ghost" size="sm" onclick={onCancel} testid="guild-loot-award-cancel">Cancel</Button>
    </span>
  </div>
</div>
