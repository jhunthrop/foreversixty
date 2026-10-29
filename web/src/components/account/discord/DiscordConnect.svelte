<!-- web/src/components/account/discord/DiscordConnect.svelte -->
<!-- The Connect state (spec 3.2, connect lane): the intro, the one `Connect a Discord channel`
     link (the card's next step: primary, or secondary where another block holds the page's one
     fill), what Discord will ask and what we keep, the reserved status line for a return outcome,
     and the swap to the paste view as an action link. In the replace view the intro is omitted
     and `Cancel` ends the column. The paste view is the other view of the same card. -->
<script lang="ts">
  import { notificationsCopy } from '../../../lib/account/notifications-copy';
  import { button } from '../../../lib/ui/button';
  import { ACTION_LINK, ACTION_LINK_CHEVRON } from '../../../lib/ui/link';
  import Glyph from '../../ui/Glyph.svelte';
  import DiscordConnectLink from './DiscordConnectLink.svelte';
  import DiscordStatusLine from './DiscordStatusLine.svelte';

  let {
    status,
    replacing = false,
    quiet = false,
    onStuck,
    onPaste,
    onCancel,
  }: {
    status: { ok: boolean; message: string } | null;
    replacing?: boolean;
    /** Another block holds the page's one fill (a seat offer): Connect is a secondary button. */
    quiet?: boolean;
    onStuck: () => void;
    onPaste: () => void;
    onCancel: () => void;
  } = $props();
</script>

<div class="flex flex-col gap-4 px-5 pt-1 pb-5" data-testid="discord-connect-view">
  {#if !replacing}<p class="text-text text-[14px]">{notificationsCopy.intro}</p>{/if}
  <div class="flex flex-col items-start">
    <DiscordConnectLink
      label={notificationsCopy.connect}
      class="{button(quiet ? 'secondary' : 'primary')} w-fit"
      testid="discord-connect"
      {onStuck}
    />
    <p class="text-muted mt-2 text-[13px]">{notificationsCopy.connectExplanation}</p>
    <div class="w-full"><DiscordStatusLine {status} /></div>
    <button type="button" class={ACTION_LINK} onclick={onPaste} data-testid="discord-paste-instead">
      {notificationsCopy.pasteInstead}<Glyph name="chevron" class={ACTION_LINK_CHEVRON} />
    </button>
    {#if replacing}
      <button type="button" class={button('ghost')} onclick={onCancel} data-testid="discord-cancel">
        {notificationsCopy.cancel}
      </button>
    {/if}
  </div>
</div>
