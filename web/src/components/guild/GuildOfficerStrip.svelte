<!-- web/src/components/guild/GuildOfficerStrip.svelte -->
<!-- Guild control-centre spec §4.B/§4.A: the officer tools glance -- claim state, who is
     waiting for approval, invite link rotate. Shared by the Overview tab (the glance) and
     the Settings tab (the same strip, restyled unchanged, spec §4.G: "the strip here is
     the glance, Settings is the editing surface").

     At every width (P.1, UX S2-1), on every tab but Settings (`collapsible`), the card folds to one 44 px line --
     "Claimed by hunthrop#1894 · Invite link" and a Show toggle -- so the tab's own answer is
     not pushed down by ~190 px of claim tools (round 8, F8-12). The full card stays mounted
     while folded, so a rotated link is not lost by folding it. -->
<script lang="ts">
  import GuildInviteLink from './GuildInviteLink.svelte';

  let {
    guildId,
    claimedByName,
    waitingCount,
    onSelectRoster,
    collapsible = false,
  }: {
    guildId: number;
    claimedByName: string | null;
    waitingCount: number;
    onSelectRoster: () => void;
    /** Fold to the one-line summary below lg (every tab but Settings). */
    collapsible?: boolean;
  } = $props();

  let open = $state(false);
  const folded = $derived(collapsible && !open);
</script>

<div class="border-line-soft bg-raised flex flex-col rounded-md border" data-testid="guild-officer-strip">
  {#if collapsible}
    <div
      class="flex h-11 items-center justify-between gap-3 pr-1 pl-4"
      data-testid="guild-officer-strip-summary"
    >
      <span class="text-muted min-w-0 truncate text-[13px]"
        >{claimedByName !== null ? `Claimed by ${claimedByName}` : 'Claimed'} · Invite link</span
      >
      <button
        class="text-gold inline-flex h-11 shrink-0 items-center px-3 text-[12px] font-bold tracking-[0.06em] uppercase"
        onclick={() => (open = !open)}
        aria-expanded={open}
        data-testid="guild-officer-strip-toggle"
      >
        {open ? 'Hide' : 'Show'}
      </button>
    </div>
  {/if}
  <div
    class={`flex flex-wrap items-center justify-between gap-4 p-4 ${folded ? 'hidden' : ''} ${collapsible && open ? 'border-line-soft border-t' : ''}`}
    data-testid="guild-officer-strip-body"
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
      <span class="quiet-label">Invite link</span>
      <span class="text-muted text-[12px]">Anyone with this link can join as a member.</span>
    </div>
    <GuildInviteLink {guildId} testidPrefix="guild-officer-invite" />
  </div>
</div>
