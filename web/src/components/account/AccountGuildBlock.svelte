<!-- web/src/components/account/AccountGuildBlock.svelte -->
<!-- One guild in /account's Guild section (premium spec 2026-11-08 section 7.4): a Row with the
     guild's emblem, its name with the plan, Not verified and Claim pending Tags, the rank and Open;
     then slim rows: what the account shares with it, its settings with the plan line, the pending
     claim's hand-off link, the next step; then Leave guild, which asks first in place (focus on
     Cancel, Escape cancels). The page owns the state and the calls; this draws one guild and
     reports what the viewer chose. -->
<script lang="ts">
  import { smallButton } from '../../lib/account/small-button';
  import type { MeGuild } from '../../lib/account/api';
  import { characterListCopy } from '../../lib/account/character-list-copy';
  import { accountPlanCopy } from '../../lib/billing/account-plan-copy';
  import { fittingGuildPlan } from '../../lib/billing/checkout-params';
  import { checkoutGuildCopy } from '../../lib/billing/checkout-guild-copy';
  import { handoffUrl } from '../../lib/billing/subscribe-step';
  import { guildHref } from '../../lib/characters';
  import { FACTION_BAR_COLOR, factionLogoSrc } from '../../lib/faction-mark';
  import { guildConsentCopy } from '../../lib/guild/copy';
  import { COPIED_LONG_MS, createCopyFeedback, type CopyFeedbackState } from '../../lib/ui/copy-feedback';
  import { LINK } from '../../lib/ui/primitives';
  import communityLinks from '../../data/links.json';
  import type { SelectOption } from '../../lib/ui/listbox-model';
  import Row from '../ui/Row.svelte';
  import Select from '../ui/Select.svelte';
  import Tag from '../ui/Tag.svelte';
  import AccountGuildPlanLine from './AccountGuildPlanLine.svelte';
  import AccountGuildRowStep from './AccountGuildRowStep.svelte';
  import ConfirmButtons from '../ui/ConfirmButtons.svelte';

  let {
    guild,
    characterCount,
    disabled,
    asking,
    consentSaved,
    onconsent,
    onask,
    oncancel,
    onleave,
  }: {
    guild: MeGuild;
    /** How many of the account's characters sit in this guild (consent and Leave cover all of them). */
    characterCount: number;
    disabled: boolean;
    /** This guild's Leave is asking "are you sure". */
    asking: boolean;
    /** The "Saved." line is showing for this guild. */
    consentSaved: boolean;
    onconsent: (event: Event) => void;
    onask: () => void;
    oncancel: () => void;
    onleave: () => void;
  } = $props();

  const SHARE_OPTIONS: readonly SelectOption[] = [
    { value: 'roster', label: guildConsentCopy.roster },
    { value: 'gear', label: guildConsentCopy.gear },
    { value: 'gear_bags', label: guildConsentCopy.gearBags },
  ];
  const shareId = $derived(`account-guild-share-${guild.id}`);
  const losesOfficerTools = $derived(
    guild.can?.some((word) => ['approve', 'award', 'settings'].includes(word)) ?? false,
  );
  const faction = $derived(guild.faction === 'alliance' || guild.faction === 'horde' ? guild.faction : null);
  /** Spec 7.4: the account's own pending claim waits for a second officer; it sends them the link. */
  const claimPending = $derived(guild.subscribe?.step === 'claim_pending_yours');
  const claimLink = $derived(
    claimPending && guild.slug !== undefined
      ? handoffUrl(
          { region: guild.region, ruleset: guild.ruleset, slug: guild.slug },
          'claim_pending_yours',
          fittingGuildPlan(guild.verified_count),
          'yearly',
        )
      : null,
  );

  let copyState = $state<CopyFeedbackState>('idle');
  const feedback = createCopyFeedback(
    (text) => navigator.clipboard.writeText(text),
    (next) => (copyState = next),
    COPIED_LONG_MS,
  );
  $effect(() => () => feedback.dispose());
</script>

<div class="flex flex-col gap-3" data-testid="account-guild-row">
  <div>
    <Row title={guild.name} sub={guild.rank} wrapTrail open>
      {#snippet lead()}
        {#if faction !== null}<img
            class="h-8 w-8 rounded-full object-contain p-[5px]"
            style:box-shadow="0 0 0 2px {FACTION_BAR_COLOR[faction]}"
            src={factionLogoSrc(faction)}
            alt=""
            width="32"
            height="32"
          />{/if}
      {/snippet}
      {#snippet titleContent()}
        <span class="inline-flex flex-wrap items-center gap-2"
          ><span data-testid="account-guild-name">{guild.name}</span>
          <AccountGuildRowStep {guild} part="plan" />
          <AccountGuildRowStep {guild} part="mark" />
          {#if claimPending}<Tag testid="account-guild-claim-pending"
              >{accountPlanCopy.guildRow.claimPending}</Tag
            >{/if}
        </span>
      {/snippet}
      {#snippet trail()}
        <a class="{LINK} text-[15px]" href={guildHref(guild.region, guild.ruleset, guild.name)}
          >{characterListCopy.open} ›</a
        >
      {/snippet}
    </Row>
    <Row title={guildConsentCopy.shareLabel} density="slim" open wrapTrail>
      {#snippet trail()}
        <Select
          id={shareId}
          size="dense"
          label={guildConsentCopy.shareLabel}
          value={guild.consent ?? 'gear'}
          options={SHARE_OPTIONS}
          {disabled}
          onchange={onconsent}
          testid="account-guild-consent"
        />
      {/snippet}
    </Row>
    {#if guild.plan}
      <AccountGuildPlanLine
        plan={guild.plan}
        region={guild.region}
        ruleset={guild.ruleset}
        name={guild.name}
      />
    {/if}
    {#if claimPending}
      <Row
        title={checkoutGuildCopy.claimPendingYours(guild.name)}
        open
        wrapTrail
        testid="account-guild-claim"
      >
        {#snippet trail()}
          {#if claimLink !== null}
            <button
              type="button"
              class={smallButton()}
              onclick={() => void feedback.copy(claimLink)}
              data-testid="account-guild-claim-copy"
              >{copyState === 'copied'
                ? checkoutGuildCopy.copyField.copied
                : checkoutGuildCopy.copyField.copy}</button
            >
          {/if}
        {/snippet}
      </Row>
    {/if}
    <AccountGuildRowStep {guild} part="line" />
  </div>
  {#if copyState === 'failed' && claimLink !== null}
    <p class="text-wipe text-[13px]" role="alert">{checkoutGuildCopy.copyField.failed} {claimLink}</p>
  {/if}
  {#if asking}
    <div class="flex flex-col gap-3" data-testid="account-guild-leave-ask">
      <p class="max-w-[72ch] text-[15px] leading-6">
        {guildConsentCopy.leaveConfirmLine(guild.name, characterCount)}
      </p>
      {#if guild.can?.includes('release')}
        <p class="max-w-[72ch] text-[15px] leading-6" data-testid="account-guild-leave-releases">
          {guildConsentCopy.leaveReleasesClaim(guild.name)}
        </p>
        <p class="text-small text-muted" data-testid="account-guild-leave-no-handover">
          {guildConsentCopy.leaveNoHandOverParts[0]}<a
            class={LINK}
            href={communityLinks.discordInvite}
            rel="noopener"
            data-testid="account-guild-leave-discord">{guildConsentCopy.leaveNoHandOverParts[1]}</a
          >{guildConsentCopy.leaveNoHandOverParts[2]}
        </p>
      {:else if losesOfficerTools}
        <p class="max-w-[72ch] text-[15px] leading-6" data-testid="account-guild-leave-officer">
          {guildConsentCopy.leaveLosesOfficer(guild.name)}
        </p>
      {/if}
      <div class="flex flex-wrap items-center gap-x-4 gap-y-[10px]">
        <ConfirmButtons
          confirm={guildConsentCopy.leaveConfirmButton}
          cancel={guildConsentCopy.leaveCancel}
          size="md"
          busy={disabled}
          onconfirm={onleave}
          {oncancel}
          testid="account-guild-leave"
        />
      </div>
    </div>
  {:else}
    <div class="flex flex-wrap items-center gap-x-4 gap-y-[10px]">
      <button
        type="button"
        class={smallButton('danger')}
        onclick={onask}
        {disabled}
        data-testid="account-guild-leave">{guildConsentCopy.leave}</button
      >
      <span class="text-small text-muted" aria-live="polite"
        >{#if consentSaved}<span data-testid="account-guild-consent-saved"
            >{guildConsentCopy.consentSaved(characterCount)}</span
          >{/if}</span
      >
    </div>
  {/if}
</div>
