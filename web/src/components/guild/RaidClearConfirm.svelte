<!-- web/src/components/guild/RaidClearConfirm.svelte -->
<!-- The clear confirm (attendance spec 3.1 "Clear confirm"): it replaces line 2's `Not this log`
     with the question, `Clear the match` (the cancel confirm's failure outline) and `Keep it`.
     Focus lands on Keep it, which stays live while the clear is in flight; Escape acts as Keep it. -->
<script lang="ts">
  import { attendanceCopy } from '../../lib/guild/attendance-copy';
  import { focusOnMount } from '../../lib/guild/focus-now';
  import { RAID_BUTTON, RAID_WIPE_BUTTON } from '../../lib/guild/raid-buttons';

  let {
    busy,
    onConfirm,
    onKeep,
  }: {
    /** The clear is in flight. */
    busy: boolean;
    onConfirm: () => void;
    onKeep: () => void;
  } = $props();

  const copy = attendanceCopy.clearConfirm;

  function onKeydown(event: KeyboardEvent): void {
    if (event.key !== 'Escape') return;
    event.preventDefault();
    onKeep();
  }
</script>

<div
  class="flex flex-col gap-3"
  role="group"
  aria-labelledby="raid-clear-question"
  aria-busy={busy}
  onkeydown={onKeydown}
  data-testid="raid-clear-confirm"
>
  <p id="raid-clear-question" class="text-strong text-[14px]">{copy.prompt}</p>
  <div class="flex flex-wrap gap-2">
    <button
      type="button"
      class={RAID_WIPE_BUTTON}
      disabled={busy}
      onclick={onConfirm}
      data-testid="raid-clear-confirm-button">{busy ? attendanceCopy.match.clearing : copy.confirm}</button
    >
    <button type="button" class={RAID_BUTTON} use:focusOnMount onclick={onKeep} data-testid="raid-clear-keep"
      >{copy.keep}</button
    >
  </div>
</div>
