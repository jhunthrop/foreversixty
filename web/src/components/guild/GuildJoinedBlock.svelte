<!-- web/src/components/guild/GuildJoinedBlock.svelte -->
<!-- The invite's joined state (spec 2026-10-27 §6.3): the joined sentence, the guild's plan line,
     then `Open the guild`. The live region is mounted with the block and the plan line reserves its
     height, so the /v1/me re-read landing neither goes unannounced nor shifts the layout. Focus
     moves to the sentence because the join button it replaces has just unmounted. -->
<script lang="ts">
  import { guildJoinCopy } from '../../lib/guild/copy';
  import { SECONDARY_BUTTON_FIXED } from '../../lib/planner/styles';

  let { planLine, openHref }: { planLine: string | null; openHref: string } = $props();

  let joinedEl = $state<HTMLElement | null>(null);
  $effect(() => {
    joinedEl?.focus();
  });
</script>

<div role="status" class="flex flex-col gap-4">
  <p bind:this={joinedEl} tabindex="-1" class="text-[14px]" data-testid="guild-join-joined">
    {guildJoinCopy.joined}
  </p>
  <p class="text-muted min-h-[21px] text-[13px]" data-testid="guild-join-plan">{planLine ?? ''}</p>
</div>
<a
  class="{SECONDARY_BUTTON_FIXED} border-line-warm-strong text-strong w-fit px-4"
  href={openHref}
  data-testid="guild-join-open"
>
  {guildJoinCopy.openGuild}
</a>
