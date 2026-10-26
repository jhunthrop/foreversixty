<!-- web/src/components/guild/RaidCancelConfirm.svelte -->
<!-- The cancel confirm (raids spec 4 "Cancel confirm"): it replaces the officer row with the
     question, the remembered `Post the cancellation to Discord` box, `Cancel the raid` and `Keep it`.
     Focus lands on Keep it, which stays live while the cancel is in flight; Escape acts as Keep it. -->
<script lang="ts">
  import { focusOnMount } from '../../lib/guild/focus-now';
  import { raidCopy } from '../../lib/guild/raid-copy';

  let {
    postCancel,
    busy,
    buttonClass,
    onPostChange,
    onConfirm,
    onKeep,
  }: {
    postCancel: boolean;
    /** The cancel is in flight. */
    busy: boolean;
    buttonClass: string;
    onPostChange: (post: boolean) => void;
    onConfirm: () => void;
    onKeep: () => void;
  } = $props();

  function onKeydown(event: KeyboardEvent): void {
    if (event.key !== 'Escape') return;
    event.preventDefault();
    onKeep();
  }
</script>

<div
  class="flex flex-col gap-3"
  role="group"
  aria-labelledby="raid-cancel-question"
  aria-busy={busy}
  onkeydown={onKeydown}
  data-testid="raid-cancel-confirm"
>
  <p id="raid-cancel-question" class="text-strong text-[14px]">{raidCopy.cancelConfirm.question}</p>
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
    <button
      type="button"
      class={buttonClass}
      disabled={busy}
      onclick={onConfirm}
      data-testid="raid-cancel-confirm-button"
      >{busy ? raidCopy.cancelConfirm.confirming : raidCopy.cancelConfirm.confirm}</button
    >
    <button type="button" class={buttonClass} use:focusOnMount onclick={onKeep} data-testid="raid-cancel-keep"
      >{raidCopy.cancelConfirm.keep}</button
    >
  </div>
</div>
