<!-- web/src/components/account/discord/DiscordStatusLine.svelte -->
<!-- The reserved 21 px result line under the Discord block's actions: a result never moves the
     buttons. Success is a polite status, a failure an alert in the failure colour. -->
<script lang="ts">
  import { notificationsCopy } from '../../../lib/account/notifications-copy';

  /** `target: 'plans'` is a refusal for want of Premium: the sentence is followed by See plans. */
  let { status }: { status: { ok: boolean; message: string; target?: string } | null } = $props();
</script>

<div class="min-h-[21px] text-[13px]" data-testid="discord-status">
  {#if status !== null}
    <p class={status.ok ? 'text-text' : 'text-wipe'} role={status.ok ? 'status' : 'alert'}>
      {status.message}
      {#if status.target === 'plans'}
        <a class="text-text underline" href="/premium">{notificationsCopy.seePlans}</a>
      {/if}
    </p>
  {/if}
</div>
