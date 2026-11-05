<!-- web/src/components/claim/ClaimNextSteps.svelte -->
<!-- The `Next` list once the viewer holds the claim (spec 2026-10-27 §6.2). With a plan to
     subscribe, the subscribe button leads and is the page's one fill; the roster link is then a
     secondary. Without one, the roster link keeps today's primary. Plan an event and Add the
     Discord bot close the row; a guild with no plan reads the gate sentence's part 1 under it,
     with no link (guild round 3). -->
<script lang="ts">
  import type { ClaimPlan } from '../../lib/guild/claim-plan';
  import { guildClaimCopy } from '../../lib/guild/copy';
  import { PRIMARY_BUTTON_FIXED, SECONDARY_BUTTON_FIXED } from '../../lib/planner/styles';
  import ClaimSubscribeButton from './ClaimSubscribeButton.svelte';

  let {
    subscribe,
    claimPlan,
    guildBase,
    settingsBase,
    gated = false,
  }: {
    /** Set when the guild is ready to subscribe: the id to check out and the button's label. */
    subscribe: { id: number; label: string } | null;
    claimPlan: ClaimPlan | null;
    guildBase: string;
    settingsBase: string;
    /** The guild has no plan: events and the Discord bot are gated. */
    gated?: boolean;
  } = $props();

  const LINK = 'w-full justify-center px-4 md:w-auto';
  const GHOST = `${SECONDARY_BUTTON_FIXED} border-line-warm text-text ${LINK}`;
  const rosterClass = $derived(subscribe === null ? `${PRIMARY_BUTTON_FIXED} ${LINK}` : GHOST);
</script>

<section class="flex flex-col gap-2" data-testid="guild-claim-next">
  <h2 class="section-title">{guildClaimCopy.nextHeading}</h2>
  <!-- The first step is the one fill, the others ghosts; full width on a phone. -->
  <div class="flex flex-col gap-3 md:flex-row">
    {#if claimPlan !== null && subscribe !== null}
      <ClaimSubscribeButton guildId={subscribe.id} {claimPlan} label={subscribe.label} />
    {/if}
    <a class={rosterClass} href={`${guildBase}#roster`} data-testid="guild-claim-next-roster">
      {guildClaimCopy.nextRoster}
    </a>
    <a class={GHOST} href={settingsBase} data-testid="guild-claim-next-invite">
      {guildClaimCopy.nextInvite}
    </a>
    <a class={GHOST} href={`${guildBase}#raids`} data-testid="guild-claim-next-plan">
      {guildClaimCopy.nextPlan}
    </a>
    <a class={GHOST} href={`${settingsBase}#discord`} data-testid="guild-claim-next-discord">
      {guildClaimCopy.nextDiscord}
    </a>
  </div>
  {#if gated}
    <p class="text-muted text-[13px]" data-testid="guild-claim-next-gated">{guildClaimCopy.nextGated}</p>
  {/if}
</section>
