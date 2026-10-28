<!-- web/src/components/checkout/GuildStepBlock.svelte -->
<!-- Spec 2026-10-27-subscribe-and-join §5.3 / §5.4: the step block of a guild card, the right
     column. It draws the step's one sentence, then the one action for the steps that hand off or
     only read (a CopyLinkField, `Check again`, `Open the guild`), then the verify sub-line, the
     other-character line (ask_leader) and the moderator fallback (ask_leader, claimed_ask).

     Two slots belong to other tasks and are marked below: Task 17 draws the claim button and its
     confirm block for claim / start_claim / confirm_claim / claim_link_bnet; Task 18 draws the
     ready button (with the switch note) and the on_plan payer's Manage billing / Move to Raid
     guild. Both arrive as the `action` snippet, drawn where this block would put its own. -->
<script lang="ts">
  import type { Snippet } from 'svelte';
  import type { MeGuild } from '../../lib/account/me-types';
  import { checkoutGuildCopy as copy, type GuildTier } from '../../lib/billing/checkout-guild-copy';
  import {
    claimKindOf,
    openerLabel,
    type CardNotice,
    type ClaimControls,
  } from '../../lib/billing/claim-flow';
  import { seatLine, stepSentence } from '../../lib/billing/guild-card-model';
  import {
    copyFieldLabel,
    handoffStepOf,
    handoffUrl,
    showsModeratorFallback,
    showsOtherCharacter,
    verifySubLine,
    type StepGuild,
  } from '../../lib/billing/subscribe-step';
  import { battlenetStartUrl } from '../../lib/account/api';
  import { INLINE_LINK } from '../../lib/guild/field-styles';
  import { SECONDARY_BUTTON_FIXED } from '../../lib/planner/styles';
  import communityLinks from '../../data/links.json';
  import CopyLinkField from '../ui/CopyLinkField.svelte';
  import ClaimConfirmBlock from './ClaimConfirmBlock.svelte';

  const CLAIM_OPENER = `${SECONDARY_BUTTON_FIXED} border-gold text-gold w-full justify-center px-4`;
  const NOTICE_RESERVE = 'min-h-[20px]';

  const TEXT_LINK =
    'text-text hover:text-strong focus-visible:outline-gold inline-flex min-h-11 w-fit items-center underline focus-visible:outline-2 focus-visible:outline-offset-2';
  const [FALLBACK_BEFORE = '', FALLBACK_AFTER = ''] = copy.moderatorFallback.split(copy.moderatorLinkText);

  let {
    guild,
    plan,
    interval,
    selfUrl,
    viewerIsOfficer,
    guildPageHref,
    onCheckAgain,
    action = undefined,
    claim = undefined,
    notice = null,
  }: {
    guild: MeGuild & StepGuild;
    plan: GuildTier;
    interval: string;
    /** This checkout's own path and query: the `next` of the other-character link. */
    selfUrl: string;
    viewerIsOfficer: boolean;
    guildPageHref: string;
    /** Re-reads /v1/me; the `checkout_open` card's `Check again`. */
    onCheckAgain: () => void;
    /** Tasks 17 and 18: the claim and ready / on_plan-payer controls. */
    action?: Snippet;
    /** Task 17: the claim openers' state; absent, claim steps draw nothing. */
    claim?: ClaimControls;
    notice?: CardNotice | null;
  } = $props();

  const step = $derived(guild.subscribe?.step ?? null);
  const sentence = $derived(stepSentence(guild, plan, viewerIsOfficer));
  const handoff = $derived(handoffStepOf(guild));
  const verifyLine = $derived(verifySubLine(guild, guild.verified));
  const seats = $derived(step === 'on_plan' ? seatLine(guild.covered) : null);
  const otherCharacterHref = $derived(`/setup?next=${encodeURIComponent(`${selfUrl}#guild-${guild.id}`)}`);
  const claimKind = $derived(claimKindOf(step));
  const slotStep = $derived(step === 'ready' || (step === 'on_plan' && guild.covered?.you_pay === true));
  const opensGuild = $derived(
    step === null ||
      step === 'frozen' ||
      step === 'claim_blocked' ||
      (step === 'on_plan' && guild.covered?.you_pay !== true),
  );
</script>

{#if sentence !== null}
  <p class="text-text m-0 text-[14px]" data-testid="guild-step-sentence">{sentence}</p>
{/if}
{#if seats !== null}
  <p class="text-text m-0 text-[14px]">{seats}</p>
{/if}

{#if handoff !== null}
  <CopyLinkField
    url={handoffUrl(guild, handoff, plan, interval)}
    label={copyFieldLabel(guild, handoff, viewerIsOfficer)}
    shareTitle={copy.share.title(guild.name)}
    shareText={handoff === 'claimed_ask'
      ? copy.share.settingsText(guild.name, plan)
      : copy.share.claimText(guild.name, plan)}
    testid={`guild-${guild.id}-handoff`}
  />
{:else if step === 'checkout_open'}
  <button type="button" class={TEXT_LINK} onclick={onCheckAgain} data-testid="guild-check-again">
    {copy.checkAgain}
  </button>
{:else if claimKind !== null}
  {#if claim}
    <button
      type="button"
      id={`guild-${guild.id}-claim-open`}
      class={CLAIM_OPENER}
      disabled={claim.locked}
      aria-expanded={claim.open}
      onclick={claim.onOpen}
      data-testid={`guild-${guild.id}-claim-open`}
    >
      {openerLabel(claimKind, guild.name)}
    </button>
    {#if claim.open}
      <ClaimConfirmBlock
        kind={claimKind}
        {guild}
        io={claim.io}
        error={claim.error}
        onDone={claim.onDone}
        onRefused={claim.onRefused}
        onCancel={claim.onCancel}
      />
    {/if}
  {/if}
{:else if step === 'claim_link_bnet'}
  <a
    href={battlenetStartUrl(`${selfUrl}#guild-${guild.id}`)}
    class={`${SECONDARY_BUTTON_FIXED} w-full justify-center px-4`}
    data-testid={`guild-${guild.id}-link-bnet`}
  >
    {copy.claimLinkBattlenet.button}
  </a>
{:else if slotStep}
  <!-- SLOT (Task 18: ready button, switch note, Manage billing / Move to Raid guild). -->
  {@render action?.()}
{/if}
{#if opensGuild}
  <a href={guildPageHref} class={TEXT_LINK}>{copy.openGuild}</a>
{/if}

{#if notice !== null}
  <p
    id={`guild-${guild.id}-notice`}
    tabindex="-1"
    role={notice.role}
    class={`text-muted m-0 text-[13px] outline-none ${NOTICE_RESERVE}`}
    data-testid={`guild-${guild.id}-notice`}
  >
    {notice.text}
  </p>
{/if}
{#if verifyLine !== null}
  <p class="text-muted m-0 text-[13px]" data-testid="guild-verify-line">{verifyLine}</p>
{/if}
{#if step !== null && showsOtherCharacter(step)}
  <p class="text-muted m-0 text-[13px]" data-testid="guild-other-character">
    {copy.otherCharacter(guild.name)}
    <a href={otherCharacterHref} class={INLINE_LINK}>{copy.getSetUp}</a>
  </p>
{/if}
{#if showsModeratorFallback(guild)}
  <p class="text-muted m-0 text-[13px]" data-testid="guild-moderator-fallback">
    {FALLBACK_BEFORE}<a href={communityLinks.discordInvite} target="_blank" rel="noopener" class={INLINE_LINK}
      >{copy.moderatorLinkText}</a
    >{FALLBACK_AFTER}
  </p>
{/if}
