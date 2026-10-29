<!-- web/src/components/guild/GuildDiscordAdd.svelte -->
<!-- The ways to add the Forever Sixty bot (bot spec §3.1). `nothing` (state 5): the gold add link,
     its explanation, the reserved status line and `Connect a channel without the bot` with its
     note. `offer` (state 4): the sentence and an outline add link above a connected webhook's
     actions. Both links are navigations to the API (middle-click works) and, officer-read-only
     or frozen, go nowhere: no href, aria-disabled, out of the tab order, described by the gate. -->
<script lang="ts">
  import { LINK_BUTTON } from '../../lib/account/discord-styles';
  import { guildBotCopy } from '../../lib/guild/discord-bot-copy';
  import type { DiscordStatus } from '../../lib/guild/discord-settings';
  import { button } from '../../lib/ui/button';
  import DiscordConnectLink from '../account/discord/DiscordConnectLink.svelte';
  import DiscordStatusLine from '../account/discord/DiscordStatusLine.svelte';

  let {
    variant,
    installHref,
    webhookHref,
    gated,
    gateId,
    busy,
    status,
    onStuck,
  }: {
    variant: 'nothing' | 'offer';
    installHref: string;
    webhookHref: string;
    gated: boolean;
    gateId: string | undefined;
    busy: boolean;
    status: DiscordStatus | null;
    onStuck: () => void;
  } = $props();
</script>

{#if variant === 'offer'}
  <div class="flex flex-col items-start gap-2" data-testid="bot-offer">
    <p class="text-[14px]">{guildBotCopy.offer}</p>
    <DiscordConnectLink
      label={guildBotCopy.add}
      class={button('secondary')}
      testid="bot-add"
      href={installHref}
      {busy}
      {gated}
      describedBy={gateId}
      {onStuck}
    />
  </div>
{:else}
  <div class="flex flex-col items-start">
    <DiscordConnectLink
      label={guildBotCopy.add}
      class={button('primary')}
      testid="bot-add"
      href={installHref}
      {busy}
      {gated}
      describedBy={gateId}
      {onStuck}
    />
    <p class="text-muted mt-2 text-[13px]" data-testid="bot-explanation">{guildBotCopy.addExplanation}</p>
    <div id="guild-discord-status" tabindex="-1" class="w-full outline-none">
      <DiscordStatusLine {status} />
    </div>
    <div class="mt-3 flex flex-col">
      <DiscordConnectLink
        label={guildBotCopy.webhookPath}
        class="{LINK_BUTTON} w-fit"
        testid="bot-webhook"
        href={webhookHref}
        {busy}
        {gated}
        describedBy={gateId}
        {onStuck}
      />
      <p class="text-muted text-[13px]">{guildBotCopy.webhookPathNote}</p>
    </div>
  </div>
{/if}
