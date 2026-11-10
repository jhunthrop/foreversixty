<!-- web/src/components/account/AccountDiscord.svelte -->
<!-- /account's Discord section in the rail (premium spec 2026-11-08 section 7.6, `#discord`): the one
     place for everything Discord. The link row and the Premium role row (while the deployment has
     the bot on), then the Notifications block in its state. Draws nothing when Discord is off on
     both counts. -->
<script lang="ts">
  import type { Me } from '../../lib/account/api';
  import { blockState } from '../../lib/account/notifications';
  import { notificationsCopy } from '../../lib/account/notifications-copy';
  import Section from '../ui/Section.svelte';
  import AccountDiscordBlock from './AccountDiscordBlock.svelte';
  import AccountDiscordLink from './AccountDiscordLink.svelte';
  import AccountDiscordRole from './AccountDiscordRole.svelte';

  let { me, now = new Date() }: { me: Me; now?: Date } = $props();

  const linkOn = $derived(me.discord_link_available === true);
  const shown = $derived(linkOn || blockState(me).state !== 'off');
</script>

{#if shown}
  <Section title={notificationsCopy.label} anchor="discord" testid="account-discord-section" panel>
    <div class="flex flex-col gap-3">
      {#if linkOn}
        <div>
          <AccountDiscordLink {me} />
          <AccountDiscordRole {me} {now} />
        </div>
      {/if}
      <AccountDiscordBlock {me} quiet={(me.seat_offer ?? null) !== null} />
    </div>
  </Section>
{/if}
