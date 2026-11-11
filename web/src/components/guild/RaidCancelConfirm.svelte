<!-- web/src/components/guild/RaidCancelConfirm.svelte -->
<!-- The cancel confirm (raids spec 4 "Cancel confirm"): it replaces the officer row with the
     question, the `Post the cancellation to Discord` box (ticked when the raid posts by default), `Cancel the raid` and `Keep it`.
     ui/ConfirmButtons draws the pair: focus lands on Keep it and Escape acts as Keep it. -->
<script lang="ts">
  import ConfirmButtons from '../ui/ConfirmButtons.svelte';
  import type { GuildZoneKind } from '../../lib/guild/api';
  import { cancelConfirmOf, cancelQuestionOf, raidCopy } from '../../lib/guild/raid-copy';

  let {
    kind = 'raid',
    postCancel,
    busy,
    onPostChange,
    onConfirm,
    onKeep,
  }: {
    /** Whose words the question and the button use. */
    kind?: GuildZoneKind;
    postCancel: boolean;
    /** The cancel is in flight. */
    busy: boolean;
    onPostChange: (post: boolean) => void;
    onConfirm: () => void;
    onKeep: () => void;
  } = $props();
</script>

<div
  class="flex flex-col gap-3"
  role="group"
  aria-labelledby="raid-cancel-question"
  aria-busy={busy}
  data-testid="raid-cancel-confirm"
>
  <p id="raid-cancel-question" class="text-strong text-[14px]">{cancelQuestionOf(kind)}</p>
  <label class="flex min-h-11 items-center gap-3 text-[14px]">
    <input
      type="checkbox"
      checked={postCancel}
      disabled={busy}
      onchange={(event) => onPostChange(event.currentTarget.checked)}
      data-testid="raid-cancel-discord"
    />
    {raidCopy.cancelConfirm.postToDiscord}
  </label>
  <div class="flex flex-wrap gap-2">
    <ConfirmButtons
      size="md"
      cancelWhileBusy
      confirm={cancelConfirmOf(kind)}
      cancel={raidCopy.cancelConfirm.keep}
      {busy}
      busyLabel={raidCopy.cancelConfirm.confirming}
      onconfirm={onConfirm}
      oncancel={onKeep}
      testid="raid-cancel"
      confirmTestid="raid-cancel-confirm-button"
      cancelTestid="raid-cancel-keep"
    />
  </div>
</div>
