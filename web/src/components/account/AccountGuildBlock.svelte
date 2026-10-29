<!-- web/src/components/account/AccountGuildBlock.svelte -->
<!-- One guild in the account page's Guild card (polish spec 4.3 pattern): the name in the heading
     face with its plan and verification tags, then settings rows (share level, guild settings or
     the next step), then Leave at the block's end. Leave asks first, in place. The page owns the
     state and the calls; this draws one guild and reports what the viewer chose. -->
<script lang="ts">
  import type { MeGuild } from '../../lib/account/api';
  import { guildHref } from '../../lib/characters';
  import { guildConsentCopy } from '../../lib/guild/copy';
  import { button } from '../../lib/ui/button';
  import communityLinks from '../../data/links.json';
  import SettingsRow from '../ui/SettingsRow.svelte';
  import Select, { type SelectOption } from '../ui/Select.svelte';
  import AccountGuildPlanLine from './AccountGuildPlanLine.svelte';
  import AccountGuildRowStep from './AccountGuildRowStep.svelte';

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

  /** Focus lands on Cancel, so a stray Enter backs out of leaving rather than leaves. */
  function focusOnMount(node: HTMLElement): void {
    node.focus();
  }
</script>

<li class="border-line-soft border-t" data-testid="account-guild-row">
  <div class="flex flex-wrap items-center gap-x-3 gap-y-2 px-5 pt-4 pb-3">
    <h3 class="heading-name min-w-0 text-[20px] [overflow-wrap:anywhere]">
      <a
        class="text-strong hover:text-gold focus-visible:outline-gold rounded-control focus-visible:outline-2"
        href={guildHref(guild.region, guild.ruleset, guild.name)}>{guild.name}</a
      >
    </h3>
    <AccountGuildRowStep {guild} part="plan" />
    <AccountGuildRowStep {guild} part="mark" />
  </div>
  <SettingsRow glyph="user" label={guildConsentCopy.shareLabel}>
    {#snippet control()}
      <Select
        id={shareId}
        label={guildConsentCopy.shareLabel}
        value={guild.consent ?? 'gear'}
        options={SHARE_OPTIONS}
        {disabled}
        onchange={onconsent}
        testid="account-guild-consent"
      />
    {/snippet}
  </SettingsRow>
  {#if guild.plan}
    <AccountGuildPlanLine plan={guild.plan} region={guild.region} ruleset={guild.ruleset} name={guild.name} />
  {/if}
  <AccountGuildRowStep {guild} part="line" />
  {#if !asking}
    <div class="border-line-soft flex flex-wrap items-center gap-x-4 gap-y-1 border-t px-5 py-4">
      <button class={button('danger')} onclick={onask} {disabled} data-testid="account-guild-leave">
        {guildConsentCopy.leave}
      </button>
      <span class="text-muted min-h-[21px] text-[13px]" aria-live="polite">
        {#if consentSaved}
          <span data-testid="account-guild-consent-saved">
            {guildConsentCopy.consentSaved(characterCount)}
          </span>
        {/if}
      </span>
    </div>
  {:else}
    <div
      class="border-line-soft rounded-control mx-5 mb-4 flex flex-col gap-3 border p-4"
      data-testid="account-guild-leave-confirm"
    >
      <p class="text-[13px]">{guildConsentCopy.leaveConfirmLine(guild.name, characterCount)}</p>
      {#if guild.can?.includes('release')}
        <p class="text-[13px]" data-testid="account-guild-leave-releases">
          {guildConsentCopy.leaveReleasesClaim(guild.name)}
        </p>
        <p class="text-muted text-[13px]" data-testid="account-guild-leave-no-handover">
          {guildConsentCopy.leaveNoHandOverParts[0]}<a
            class="underline"
            href={communityLinks.discordInvite}
            rel="noopener"
            data-testid="account-guild-leave-discord">{guildConsentCopy.leaveNoHandOverParts[1]}</a
          >{guildConsentCopy.leaveNoHandOverParts[2]}
        </p>
      {:else if losesOfficerTools}
        <p class="text-[13px]" data-testid="account-guild-leave-officer">
          {guildConsentCopy.leaveLosesOfficer(guild.name)}
        </p>
      {/if}
      <div class="flex flex-wrap gap-3">
        <button
          class={button('danger', 'md', true)}
          onclick={onleave}
          {disabled}
          data-testid="account-guild-leave-confirm-button"
        >
          {guildConsentCopy.leaveConfirmButton}
        </button>
        <button
          class={button('secondary', 'md', true)}
          onclick={oncancel}
          {disabled}
          use:focusOnMount
          data-testid="account-guild-leave-cancel"
        >
          {guildConsentCopy.leaveCancel}
        </button>
      </div>
      {#if consentSaved}
        <span class="text-muted text-[13px]" data-testid="account-guild-consent-saved">
          {guildConsentCopy.consentSaved(characterCount)}
        </span>
      {/if}
    </div>
  {/if}
</li>
