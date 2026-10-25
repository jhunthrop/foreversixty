<!-- web/src/components/account/discord/DiscordReplace.svelte -->
<!-- The replace link by source (spec 3.2 row 7). A pasted webhook: the link-button that opens the
     replace view. A connected one: `Connect a different channel`, straight to Discord, with the
     note that connecting again removes the old webhook from Discord. -->
<script lang="ts">
  import type { DiscordNotifications } from '../../../lib/account/api';
  import { LINK_BUTTON } from '../../../lib/account/discord-styles';
  import { sourceOf } from '../../../lib/account/notifications';
  import { notificationsCopy } from '../../../lib/account/notifications-copy';
  import DiscordConnectLink from './DiscordConnectLink.svelte';

  let {
    discord,
    disabled = false,
    onReplace,
  }: { discord: DiscordNotifications; disabled?: boolean; onReplace: () => void } = $props();
</script>

{#if sourceOf(discord) === 'connected'}
  <div class="flex w-fit flex-col">
    <DiscordConnectLink
      label={notificationsCopy.reconnect}
      class="{LINK_BUTTON} w-fit"
      testid="discord-replace"
    />
    <p class="text-muted text-[13px]" data-testid="discord-replace-note">{notificationsCopy.reconnectNote}</p>
  </div>
{:else}
  <button class="{LINK_BUTTON} w-fit" {disabled} onclick={onReplace} data-testid="discord-replace">
    {notificationsCopy.replace}
  </button>
{/if}
