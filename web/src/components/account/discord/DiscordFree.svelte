<!-- web/src/components/account/discord/DiscordFree.svelte -->
<!-- Not on Premium and nothing saved: one sentence and See plans, no panel, no lock. -->
<script lang="ts">
  import { notificationsCopy } from '../../../lib/account/notifications-copy';
</script>

<p class="text-muted min-h-11 text-[13px]" data-testid="discord-free">
  {notificationsCopy.freeSentence}
  <a
    class="text-text hover:text-gold-hover inline-flex min-h-11 items-center underline"
    href="/premium"
    data-testid="discord-free-link">{notificationsCopy.seePlans}</a
  >
</p>
