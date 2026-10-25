<!-- web/src/components/account/discord/DiscordFree.svelte -->
<!-- Not on Premium and nothing saved: one sentence and See plans, no panel, no lock. The 44 px
     row is the paragraph's; the link is inline so its focus ring hugs the words. -->
<script lang="ts">
  import { notificationsCopy } from '../../../lib/account/notifications-copy';
</script>

<p class="text-muted flex min-h-11 items-center text-[13px]" data-testid="discord-free">
  <span>
    {notificationsCopy.freeSentence}
    <a class="text-text hover:text-gold-hover underline" href="/premium" data-testid="discord-free-link"
      >{notificationsCopy.seePlans}</a
    >
  </span>
</p>
