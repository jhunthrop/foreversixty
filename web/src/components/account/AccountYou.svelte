<!-- web/src/components/account/AccountYou.svelte -->
<!-- /account's You panel at the top of the rail (premium spec 2026-11-08 section 7.7; directions mock
     b4-account): who is signed in (the BattleTag, or the email) and how, Sign out (which asks first),
     then the pseudonym Switch and what it does. Deleting the account is its own panel (AccountDelete). -->
<script lang="ts">
  import { accountPageCopy } from '../../lib/account/account-page-copy';
  import { devicesCopy } from '../../lib/account/devices-copy';
  import { smallButton } from '../../lib/account/small-button';
  import type { Fact } from '../../lib/ui/facts';
  import Facts from '../ui/Facts.svelte';
  import Section from '../ui/Section.svelte';
  import Switch from '../ui/Switch.svelte';
  import ConfirmButtons from '../ui/ConfirmButtons.svelte';

  let {
    battletag,
    email,
    anonymize,
    busy,
    error = '',
    signOutError = '',
    onAnonymize,
    onSignOut,
  }: {
    battletag: string | null | undefined;
    email: string | null | undefined;
    anonymize: boolean;
    busy: boolean;
    /** The switch's save failed. */
    error?: string;
    /** Sign out failed. */
    signOutError?: string;
    onAnonymize: (wanted: boolean) => void;
    onSignOut: () => void;
  } = $props();

  const labelId = 'account-pseudonym-label';
  let asking = $state(false);
  const viaBattlenet = $derived(battletag !== undefined && battletag !== null);
  const facts = $derived<Fact[]>([
    {
      id: 'who',
      label: viaBattlenet ? accountPageCopy.battletagLabel : accountPageCopy.emailLabel,
      value: (viaBattlenet ? battletag : email) ?? '',
      testid: 'account-battletag',
    },
    {
      id: 'how',
      label: accountPageCopy.signInLabel,
      value: viaBattlenet ? accountPageCopy.signInBattlenet : accountPageCopy.signInEmail,
      muted: true,
    },
  ]);
</script>

<Section title={accountPageCopy.youLabel} anchor="you" testid="account-you" panel>
  <div class="flex flex-col gap-3">
    <Facts items={facts} />
    {#if asking}
      <p class="text-[15px] leading-6" data-testid="account-sign-out-question">
        {accountPageCopy.signOutAsk}
      </p>
      <div class="flex flex-wrap gap-[10px]">
        <ConfirmButtons
          confirm={accountPageCopy.signOut}
          cancel={devicesCopy.cancel}
          {busy}
          onconfirm={onSignOut}
          oncancel={() => (asking = false)}
          testid="account-sign-out"
        />
      </div>
    {:else}
      <div>
        <button
          type="button"
          class={smallButton()}
          disabled={busy}
          onclick={() => (asking = true)}
          data-testid="account-sign-out">{accountPageCopy.signOut}</button
        >
      </div>
    {/if}
    {#if signOutError !== ''}<p class="text-wipe text-[13px]" role="alert">{signOutError}</p>{/if}
    <!-- pr-2 takes back the Switch's -mr-2 hit area, so the track ends on the panel's text edge. -->
    <div class="flex min-h-11 items-center justify-between gap-3 pr-2">
      <span id={labelId} class="text-[15px] leading-5">{accountPageCopy.pseudonymLabel}</span>
      <Switch
        checked={anonymize}
        labelledby={labelId}
        disabled={busy}
        onchange={onAnonymize}
        testid="anonymize"
      />
    </div>
    <p class="text-small text-muted">{accountPageCopy.pseudonymNote}</p>
    {#if error !== ''}<p class="text-wipe text-[13px]" role="alert" data-testid="account-you-error">
        {error}
      </p>{/if}
  </div>
</Section>
