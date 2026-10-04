<!-- web/src/components/guild/GuildOfficerStrip.svelte -->
<!-- Guild control-centre spec §4.B/§4.A: the officer tools glance -- claim state, who is
     waiting for approval, invite link rotate. Shared by the Overview tab (the glance) and
     the Settings tab (the same strip, restyled unchanged, spec §4.G: "the strip here is
     the glance, Settings is the editing surface"). -->
<script lang="ts">
  import { GuildApiError, rotateInvite } from '../../lib/guild/api';
  import { SECONDARY_BUTTON_FIXED } from '../../lib/planner/styles';
  import { BUSY_CLASS } from '../../lib/ui/busy';

  let {
    guildId,
    claimedByName,
    waitingCount,
    onSelectRoster,
  }: { guildId: number; claimedByName: string | null; waitingCount: number; onSelectRoster: () => void } =
    $props();

  let busy = $state(false);
  let rotated = $state<{ url: string } | null>(null);
  let error = $state('');

  async function onRotate(): Promise<void> {
    busy = true;
    error = '';
    try {
      const result = await rotateInvite(guildId);
      rotated = { url: result.url };
    } catch (thrown) {
      error = thrown instanceof GuildApiError ? thrown.message : 'That did not work; try again';
    } finally {
      busy = false;
    }
  }
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
    <button
      class={`${SECONDARY_BUTTON_FIXED} border-line-warm text-text px-3 ${busy ? BUSY_CLASS : ''}`}
      onclick={() => void onRotate()}
      disabled={busy}
      aria-busy={busy}
      data-testid="guild-officer-rotate-invite"
    >
      Rotate invite link
    </button>
  </div>
  {#if rotated !== null}
    <p class="w-full text-[12px]" data-testid="guild-officer-invite-token">
      This link is shown once. Copy it now. <code class="font-mono">{rotated.url}</code>
    </p>
  {/if}
  {#if error !== ''}<p class="w-full text-[13px]" role="alert">{error}</p>{/if}
</div>
