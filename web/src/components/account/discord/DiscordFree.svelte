<!-- web/src/components/account/discord/DiscordFree.svelte -->
<!-- Not on Premium and nothing saved: the card's one sentence and See plans as an action link
     (polish spec 4.3), no lock and no tag. -->
<script lang="ts">
  import { notificationsCopy } from '../../../lib/account/notifications-copy';
  import { PREMIUM_LINKS } from '../../../lib/billing/premium-links';
  import ActionLink from '../../ui/ActionLink.svelte';
</script>

<div class="flex flex-col items-start gap-1 px-5 pb-4" data-testid="discord-free">
  <p class="text-muted text-[14px]">{notificationsCopy.freeSentence}</p>
  <ActionLink href={PREMIUM_LINKS.notifications} data-testid="discord-free-link"
    >{notificationsCopy.seePlans}</ActionLink
  >
</div>
