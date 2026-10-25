<!-- web/src/components/account/discord/DiscordConnect.svelte -->
<!-- The Connect state (spec 3.2, connect lane): the intro, the one gold `Connect a Discord
     channel` link, what Discord will ask and what we keep, the reserved status line for a return
     outcome, and the swap to the paste view. In the replace view the intro is omitted and
     `Cancel` ends the column. The paste view is the other view of the same panel. -->
<script lang="ts">
  import { LINK_BUTTON } from '../../../lib/account/discord-styles';
  import { notificationsCopy } from '../../../lib/account/notifications-copy';
  import { PRIMARY_BUTTON_FIXED } from '../../../lib/planner/styles';
  import DiscordConnectLink from './DiscordConnectLink.svelte';
  import DiscordStatusLine from './DiscordStatusLine.svelte';

  let {
    status,
    replacing = false,
    onPaste,
    onCancel,
  }: {
    status: { ok: boolean; message: string } | null;
    replacing?: boolean;
    onPaste: () => void;
    onCancel: () => void;
  } = $props();
</script>

<div class="flex flex-col gap-4" data-testid="discord-connect-view">
  {#if !replacing}<p class="text-muted text-[13px]">{notificationsCopy.intro}</p>{/if}
  <div class="flex flex-col items-start">
    <DiscordConnectLink
      label={notificationsCopy.connect}
      class="{PRIMARY_BUTTON_FIXED} px-5 whitespace-nowrap"
      testid="discord-connect"
    />
    <p class="text-muted mt-2 text-[13px]">{notificationsCopy.connectExplanation}</p>
    <div class="w-full"><DiscordStatusLine {status} /></div>
    <button type="button" class={LINK_BUTTON} onclick={onPaste} data-testid="discord-paste-instead">
      {notificationsCopy.pasteInstead}
    </button>
    {#if replacing}
      <button type="button" class={LINK_BUTTON} onclick={onCancel} data-testid="discord-cancel">
        {notificationsCopy.cancel}
      </button>
    {/if}
  </div>
</div>
