<!-- web/src/components/account/discord/DiscordFree.svelte -->
<!-- Not on Premium and nothing saved: the block's one sentence and See plans (premium spec
     2026-11-08 section 7.6), no lock and no tag. -->
<script lang="ts">
  import { notificationsCopy } from '../../../lib/account/notifications-copy';
  import { PREMIUM_LINKS } from '../../../lib/billing/premium-links';
  import { LINK } from '../../../lib/ui/primitives';
</script>

<div class="flex flex-col gap-[6px]" data-testid="discord-free">
  <p class="text-small text-muted">{notificationsCopy.freeSentence}</p>
  <p>
    <a class="{LINK} text-[15px]" href={PREMIUM_LINKS.notifications} data-testid="discord-free-link"
      >{notificationsCopy.seePlans} ›</a
    >
  </p>
</div>
