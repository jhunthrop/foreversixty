<!-- web/src/components/claim/ClaimNextSteps.svelte -->
<!-- The `Next` list once the viewer holds the claim (spec 2026-10-27 §6.2). With a plan to
     subscribe, the subscribe button leads and is the page's one fill; the roster link is then a
     secondary. A Section of Buttons in sentence case (web layout spec 2026-11-05 §3A). Without one, the roster link keeps today's primary. Plan an event and Add the
     Discord bot close the row; a guild with no plan reads the gate sentence's part 1 under it,
     with no link (guild round 3). -->
<script lang="ts">
  import type { ClaimPlan } from '../../lib/guild/claim-plan';
  import { guildClaimCopy } from '../../lib/guild/copy';
  import Button from '../ui/Button.svelte';
  import Section from '../ui/Section.svelte';
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
</script>

<Section title={guildClaimCopy.nextHeading} panel testid="guild-claim-next">
  <!-- The first step is the one fill, the others outlines; full width on a phone. -->
  <div class="flex flex-col gap-3 md:flex-row md:flex-wrap [&>*]:max-md:w-full">
    {#if claimPlan !== null && subscribe !== null}
      <ClaimSubscribeButton guildId={subscribe.id} {claimPlan} label={subscribe.label} />
    {/if}
    <Button
      tone={subscribe === null ? 'primary' : 'secondary'}
      href={`${guildBase}#roster`}
      testid="guild-claim-next-roster">{guildClaimCopy.nextRoster}</Button
    >
    <Button href={settingsBase} testid="guild-claim-next-invite">{guildClaimCopy.nextInvite}</Button>
    <Button href={`${guildBase}#raids`} testid="guild-claim-next-plan">{guildClaimCopy.nextPlan}</Button>
    <Button href={`${settingsBase}#discord`} testid="guild-claim-next-discord"
      >{guildClaimCopy.nextDiscord}</Button
    >
  </div>
  {#if gated}
    <p class="text-muted pt-3 text-[13px]" data-testid="guild-claim-next-gated">{guildClaimCopy.nextGated}</p>
  {/if}
</Section>
