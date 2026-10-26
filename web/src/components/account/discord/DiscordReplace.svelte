<!-- web/src/components/account/discord/DiscordReplace.svelte -->
<!-- The replace link by source (spec 3.2 row 7). A pasted webhook: the link-button that opens the
     replace view. A connected one: `Connect a different channel`, straight to Discord, with the
     note that connecting again removes the old webhook from Discord. The guild section passes
     its own `href` and `gated` through to the connect link. -->
<script lang="ts">
  import type { DiscordSource } from '../../../lib/account/api';
  import { LINK_BUTTON } from '../../../lib/account/discord-styles';
  import { notificationsCopy } from '../../../lib/account/notifications-copy';
  import DiscordConnectLink from './DiscordConnectLink.svelte';

  let {
    source,
    disabled = false,
    href,
    gated = false,
    describedBy,
    onReplace = () => {},
    onStuck,
  }: {
    source: DiscordSource;
    disabled?: boolean;
    href?: string;
    gated?: boolean;
    describedBy?: string;
    /** The pasted source's replace view; a connected source navigates instead. */
    onReplace?: () => void;
    onStuck: () => void;
  } = $props();
</script>

{#if source === 'connected'}
  <div class="flex w-fit flex-col">
    <DiscordConnectLink
      label={notificationsCopy.reconnect}
      class="{LINK_BUTTON} w-fit"
      testid="discord-replace"
      busy={disabled}
      {href}
      {gated}
      {describedBy}
      {onStuck}
    />
    <p class="text-muted text-[13px]" data-testid="discord-replace-note">{notificationsCopy.reconnectNote}</p>
  </div>
{:else}
  <button class="{LINK_BUTTON} w-fit" {disabled} onclick={onReplace} data-testid="discord-replace">
    {notificationsCopy.replace}
  </button>
{/if}
