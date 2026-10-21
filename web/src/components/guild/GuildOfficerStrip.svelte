<!-- web/src/components/guild/GuildOfficerStrip.svelte -->
<!-- Guild control-centre spec §4.B/§4.A: the officer tools glance -- claim state, who is
     waiting for approval, invite link rotate. Shared by the Overview tab (the glance) and
     the Settings tab (the same strip, restyled unchanged, spec §4.G: "the strip here is
     the glance, Settings is the editing surface"). -->
<script lang="ts">
  import GuildInviteLink from './GuildInviteLink.svelte';

  let {
    guildId,
    claimedByName,
    waitingCount,
    onSelectRoster,
  }: { guildId: number; claimedByName: string | null; waitingCount: number; onSelectRoster: () => void } =
    $props();
</script>

<div
  class="border-line-soft bg-raised flex flex-wrap items-center justify-between gap-4 rounded-md border p-4"
  data-testid="guild-officer-strip"
>
  <div class="flex items-center gap-3">
    <span class="pill pill-site">Claimed</span>
    {#if claimedByName !== null}<span class="text-muted text-[13px]">by {claimedByName}</span>{/if}
  </div>
  {#if waitingCount > 0}
    <button
      class="text-gold text-[13px] font-bold"
      onclick={onSelectRoster}
      data-testid="guild-officer-waiting"
    >
      {waitingCount} waiting for approval
    </button>
  {/if}
  <div class="flex flex-wrap items-center gap-3">
    <span class="label text-muted">Invite link</span>
    <span class="text-muted text-[12px]">Anyone with this link can join as a member.</span>
  </div>
  <GuildInviteLink {guildId} testidPrefix="guild-officer-invite" />
</div>
