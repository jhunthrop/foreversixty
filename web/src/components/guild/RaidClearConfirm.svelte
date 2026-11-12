<!-- web/src/components/guild/RaidClearConfirm.svelte -->
<!-- The clear confirm (attendance spec 3.1 "Clear confirm"): it replaces line 2's `Not this log`
     with the question, `Clear the match` (the cancel confirm's failure outline) and `Keep it`.
     ui/ConfirmButtons draws the pair: focus lands on Keep it and Escape acts as Keep it. -->
<script lang="ts">
  import { attendanceCopy } from '../../lib/guild/attendance-copy';
  import ConfirmButtons from '../ui/ConfirmButtons.svelte';

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
</script>

<div
  class="flex flex-col gap-3"
  role="group"
  aria-labelledby="raid-clear-question"
  aria-busy={busy}
  data-testid="raid-clear-confirm"
>
  <p id="raid-clear-question" class="text-strong text-[14px]">{copy.prompt}</p>
  <div class="flex flex-wrap gap-2">
    <ConfirmButtons
      size="md"
      cancelWhileBusy
      confirm={copy.confirm}
      cancel={copy.keep}
      {busy}
      busyLabel={attendanceCopy.match.clearing}
      onconfirm={onConfirm}
      oncancel={onKeep}
      testid="raid-clear"
      confirmTestid="raid-clear-confirm-button"
      cancelTestid="raid-clear-keep"
    />
  </div>
</div>
