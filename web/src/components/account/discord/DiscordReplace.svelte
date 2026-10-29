<!-- web/src/components/account/discord/DiscordReplace.svelte -->
<!-- The replace control by source (spec 3.2 row 7). A pasted webhook: the button that opens the
     replace view. A connected one: `Connect a different channel`, straight to Discord, with the
     note that connecting again removes the old webhook from Discord. `variant="row"` (the account
     card) draws it as a settings row with the refresh tile, the note as its sub and a chevron;
     the default is the guild section's text link. The guild section passes its own `href` and
     `gated` through to the connect link. -->
<script lang="ts">
  import type { DiscordSource } from '../../../lib/account/api';
  import { LINK_BUTTON } from '../../../lib/account/discord-styles';
  import { notificationsCopy } from '../../../lib/account/notifications-copy';
  import { ROW, ROW_INTERACTIVE, ROW_LABEL, ROW_SUB } from '../../../lib/ui/row';
  import Glyph from '../../ui/Glyph.svelte';
  import IconTile from '../../ui/IconTile.svelte';
  import SettingsRow from '../../ui/SettingsRow.svelte';
  import DiscordConnectLink from './DiscordConnectLink.svelte';

  let {
    source,
    variant = 'link',
    disabled = false,
    href,
    gated = false,
    describedBy,
    onReplace = () => {},
    onStuck,
  }: {
    source: DiscordSource;
    variant?: 'link' | 'row';
    disabled?: boolean;
    href?: string;
    gated?: boolean;
    describedBy?: string;
    /** The pasted source's replace view; a connected source navigates instead. */
    onReplace?: () => void;
    onStuck: () => void;
  } = $props();
</script>

{#if source === 'connected' && variant === 'row'}
  <DiscordConnectLink
    label={notificationsCopy.reconnect}
    class="{ROW} {ROW_INTERACTIVE}"
    testid="discord-replace"
    busy={disabled}
    {href}
    {gated}
    {describedBy}
    {onStuck}
  >
    {#snippet content(text)}
      <IconTile glyph="refresh" size={32} />
      <span class="flex min-w-0 flex-1 flex-col">
        <span class="text-strong {ROW_LABEL}">{text}</span>
        <span class={ROW_SUB} data-testid="discord-replace-note">{notificationsCopy.reconnectNote}</span>
      </span>
      <Glyph name="chevron" class="text-muted" />
    {/snippet}
  </DiscordConnectLink>
{:else if source === 'connected'}
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
{:else if variant === 'row'}
  <SettingsRow
    label={notificationsCopy.replace}
    glyph="refresh"
    {disabled}
    onclick={onReplace}
    testid="discord-replace"
  />
{:else}
  <button class="{LINK_BUTTON} w-fit" {disabled} onclick={onReplace} data-testid="discord-replace">
    {notificationsCopy.replace}
  </button>
{/if}
