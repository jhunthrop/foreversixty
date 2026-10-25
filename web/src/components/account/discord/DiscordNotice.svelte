<!-- web/src/components/account/discord/DiscordNotice.svelte -->
<!-- The 2 px left-bar notice: failure colour for a switched-off webhook, gold for a lapsed plan. -->
<script lang="ts">
  import type { Snippet } from 'svelte';

  let {
    tone,
    title,
    detail = null,
    children,
  }: {
    tone: 'wipe' | 'gold';
    title: string;
    detail?: string | null;
    children?: Snippet;
  } = $props();
</script>

<div
  class="flex flex-col gap-1 border-l-2 pl-3 {tone === 'wipe' ? 'border-wipe' : 'border-gold'}"
  data-testid="discord-notice"
>
  <p class="text-strong text-[14px]">{title}</p>
  {#if detail !== null}<p class="text-muted text-[13px]">{detail}</p>{/if}
  {@render children?.()}
</div>
