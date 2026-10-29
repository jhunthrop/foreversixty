<!-- web/src/components/account/discord/DiscordStatusLine.svelte -->
<!-- The result line under the Discord block's actions. Where a result must never move the buttons
     (the connect view, the guild section) the 21 px line is reserved; the account card passes
     `reserve={false}` and draws the line only when there is a result, so no dead band sits under
     its action row. Success is a polite status, a failure an alert in the failure colour. -->
<script lang="ts">
  import { notificationsCopy } from '../../../lib/account/notifications-copy';
  import { PREMIUM_LINKS } from '../../../lib/billing/premium-links';

  /** `target: 'plans'` is a refusal for want of Premium: the sentence is followed by See plans. */
  let {
    status,
    reserve = true,
  }: { status: { ok: boolean; message: string; target?: string } | null; reserve?: boolean } = $props();
</script>

{#if reserve || status !== null}
  <div class="mt-2 text-[13px] {reserve ? 'min-h-[21px]' : ''}" data-testid="discord-status">
    {#if status !== null}
      <p class={status.ok ? 'text-text' : 'text-wipe'} role={status.ok ? 'status' : 'alert'}>
        {status.message}
        {#if status.target === 'plans'}
          <a class="text-text underline" href={PREMIUM_LINKS.notifications}>{notificationsCopy.seePlans}</a>
        {/if}
      </p>
    {/if}
  </div>
{/if}
