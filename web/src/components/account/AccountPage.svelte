<!-- web/src/components/account/AccountPage.svelte -->
<!-- /account (premium spec 2026-11-08 sections 7, 7A) as one island over one `/v1/me` read: the head
     (title, BattleTag, Sign out, the main character as the answer), then a split page. Main: the
     Characters Plate, the Plan, the Guilds; rail: Devices, Discord, You. Signed out it is a form
     page with the Battle.net button; a failed read is a Plate with Try again; while it loads, the
     ready page's shape (AccountLoading). The page owns the account and every call; the sections
     draw it and report what the viewer chose. -->
<script lang="ts">
  import {
    ACCOUNT_FAILED,
    battlenetStartUrl,
    companionsOf,
    fetchMeOnce,
    listDevices,
    pairDevice,
    refetchMe,
    refreshDevices,
    revokeDevice,
    setAnonymize,
    setMainCharacter,
    signOut,
    type Device,
    type Me,
    type PairingCode,
  } from '../../lib/account/api';
  import { accountPageCopy } from '../../lib/account/account-page-copy';
  import { characterListCopy } from '../../lib/account/character-list-copy';
  import { selectedCharacter } from '../../lib/account/hero-character';
  import {
    clearAccountReserve,
    recordAccountReserve,
    scheduleAccountReserveWrites,
  } from '../../lib/account/layout';
  import { mainCharacter, pointerForCharacter } from '../../lib/account/main-character';
  import { premiumRowLinked, premiumRowValue, premiumRowVisible } from '../../lib/account/notifications';
  import { accountPlanLine, onPremium } from '../../lib/account/plan-line';
  import { compareSnapshot, takeSnapshot, type RefreshResult } from '../../lib/account/refresh-snapshot';
  import { accountSignInCopy } from '../../lib/account/signin-copy';
  import { openPortal } from '../../lib/billing/api';
  import { checkoutCopy } from '../../lib/billing/copy';
  import type { MeSeatOffer } from '../../lib/billing/contract';
  import { accountNameOf } from '../../lib/character-selector/rows';
  import {
    CURRENT_CHARACTER_CHANGED,
    onCurrentCharacterChange,
    readCurrent,
    writeCurrent,
    type CurrentCharacter,
  } from '../../lib/current-character';
  import { createQueryState } from '../../lib/data/query.svelte';
  import { viewerZone } from '../../lib/dates';
  import { leaveGuild, updateConsent, type GuildConsent } from '../../lib/guild/api';
  import { API_BASE_URL } from '../../lib/planner/config';
  import { createPairingPoll, type PairingPoll } from '../../lib/reports/pairing-poll';
  import { myReportsCopy } from '../../lib/reports/my-reports-copy';
  import { button } from '../../lib/ui/button';
  import Notice from '../ui/Notice.svelte';
  import Page from '../ui/Page.svelte';
  import PageHeader from '../ui/PageHeader.svelte';
  import Plate from '../ui/Plate.svelte';
  import Section from '../ui/Section.svelte';
  import AccountDevices from './AccountDevices.svelte';
  import AccountDiscord from './AccountDiscord.svelte';
  import AccountGuildBlock from './AccountGuildBlock.svelte';
  import AccountHead from './AccountHead.svelte';
  import AccountLoading from './AccountLoading.svelte';
  import AccountPremiumBlock from './AccountPremiumBlock.svelte';
  import AccountSeatOffer from './AccountSeatOffer.svelte';
  import AccountYou from './AccountYou.svelte';
  import CharacterList from './CharacterList.svelte';

  /** Where a sign-in from this page comes back to. */
  const SIGN_IN_NEXT = '/account?signed_in=1';
  const CONSENT_SAVED_MS = 3000;

  // One `/v1/me` read, shared with every other island through the client cache (lib/data/query.ts).
  const session = createQueryState<Me | null>(`${API_BASE_URL}/v1/me`, () => fetchMeOnce(), {
    scope: 'private',
    ttlMs: 10 * 60 * 1000,
  });

  // A writable $derived: tracks the cache, and the guild and setting handlers below also write it
  // so the page changes in the same frame; the next `session.data` change takes over again.
  let me = $derived(session.data);
  let now = $state(new Date());

  /** Which section an action failed in, and why. */
  type Area = 'plan' | 'devices' | 'you' | 'guild' | 'head';
  let failure = $state<{ area: Area; message: string } | null>(null);
  const errorIn = (area: Area): string => (failure?.area === area ? failure.message : '');
  let busy = $state(false);
  let billingBusy = $state(false);

  // ---------------------------------------------------------------- devices
  let devicesStatus = $state<'idle' | 'loading' | 'ready' | 'failed'>('idle');
  let devices = $state<Device[]>([]);
  let pairing = $state<PairingCode | null>(null);
  let pairedDevice = $state<Device | null>(null);
  let activePoll: PairingPoll | null = null;
  $effect(() => () => activePoll?.stop());

  const hasMe = $derived(me !== null);
  $effect(() => {
    if (!hasMe) {
      devicesStatus = 'idle';
      return;
    }
    devicesStatus = 'loading';
    void listDevices()
      .then((result) => {
        devices = result;
        devicesStatus = 'ready';
      })
      .catch(() => {
        devicesStatus = 'failed';
      });
  });

  // A refresh after an action re-reads `/v1/me` with the account on screen: the page stays put.
  const rereading = $derived(session.status === 'loading' && session.data !== null);
  const status = $derived.by<'loading' | 'ready' | 'failed'>(() => {
    if (session.status === 'failed') return 'failed';
    if (session.status !== 'ready' && !rereading) return 'loading';
    if (me === null) return 'ready';
    if (devicesStatus === 'failed') return 'failed';
    return devicesStatus === 'ready' ? 'ready' : 'loading';
  });
  const signedIn = $derived(me !== null);
  const viaBattlenet = $derived(me?.user.battletag !== undefined && me?.user.battletag !== null);

  // ---------------------------------------------------------------- the loading reserve
  // The ready page's two column heights, stored so the next visit's loading state reserves them
  // (spec 2026-10-25-discord-notifications 9.1; Base.astro reads them before first paint).
  let stopReserveWrites: (() => void) | null = null;
  const readyAccount = $derived(status === 'ready' && signedIn);
  $effect(() => {
    if (!readyAccount) return;
    const stop = scheduleAccountReserveWrites(() =>
      recordAccountReserve(
        document.querySelector('[data-testid="account-grid"] > .ui-page-main'),
        document.querySelector('[data-testid="account-grid"] > .ui-page-rail'),
      ),
    );
    stopReserveWrites = stop;
    return () => {
      stop();
      stopReserveWrites = null;
    };
  });

  // ---------------------------------------------------------------- the header's character
  let current = $state<CurrentCharacter | null>(null);
  $effect(() => {
    const read = (): void => {
      current = readCurrent();
    };
    read();
    return onCurrentCharacterChange(read);
  });
  const currentKey = $derived(me === null ? undefined : selectedCharacter(current, me)?.key);
  const mainKey = $derived(
    me === null ? undefined : mainCharacter(me.characters, me.main_character_key)?.key,
  );

  // ---------------------------------------------------------------- arrivals (read once, then dropped)
  // Back from a Battle.net refresh (?refreshed=1): read the account afresh and say what changed.
  let refreshResult = $state<RefreshResult | null>(null);
  $effect(() => {
    const params = new URLSearchParams(window.location.search);
    if (params.get('refreshed') !== '1') return;
    window.history.replaceState({}, '', window.location.pathname + window.location.hash);
    const before = takeSnapshot(sessionStorage);
    refetchMe(
      (fresh) => {
        if (fresh !== null) refreshResult = compareSnapshot(before, fresh);
      },
      undefined,
      () => (refreshResult = { kind: 'error' }),
    );
  });

  // A fresh sign-in (?signed_in=1): the main becomes the current character, and a banner says so.
  let signedInBanner = $state('');
  $effect(() => {
    if (me === null) return;
    const params = new URLSearchParams(window.location.search);
    if (params.get('signed_in') !== '1') return;
    const main = mainCharacter(me.characters, me.main_character_key);
    if (main !== null) {
      writeCurrent(pointerForCharacter(main));
      window.dispatchEvent(new Event(CURRENT_CHARACTER_CHANGED));
      signedInBanner = accountPageCopy.signedInBanner(main.name);
    } else {
      signedInBanner = accountPageCopy.signedInBannerEmpty(accountNameOf(me), viaBattlenet);
    }
    window.history.replaceState({}, '', window.location.pathname + window.location.hash);
  });

  // ---------------------------------------------------------------- actions
  async function run(area: Area, action: () => Promise<void>): Promise<void> {
    busy = true;
    failure = null;
    try {
      await action();
    } catch (thrown) {
      failure = { area, message: thrown instanceof Error ? thrown.message : ACCOUNT_FAILED };
    } finally {
      busy = false;
    }
  }

  /** Records the chosen main and makes it the current character, so every tool follows. */
  async function setMain(key: string): Promise<void> {
    await setMainCharacter(key);
    if (me === null) return;
    me = { ...me, main_character_key: key };
    const chosen = me.characters.find((c) => c.key === key);
    if (chosen !== undefined) {
      writeCurrent(pointerForCharacter(chosen));
      window.dispatchEvent(new Event(CURRENT_CHARACTER_CHANGED));
    }
  }

  const onPair = (): void =>
    void run('devices', async () => {
      activePoll?.stop();
      pairedDevice = null;
      const beforeIds = new Set(devices.map((device) => device.id));
      const code = await pairDevice();
      pairing = code;
      activePoll = createPairingPoll(
        async () => companionsOf(await refreshDevices()),
        beforeIds,
        (device) => {
          pairedDevice = device;
          pairing = null;
          devices = [...devices, device];
        },
        { expiresInMs: code.expires_in * 1000 },
      );
      activePoll.start();
    });

  const onRevoke = (id: string): void =>
    void run('devices', async () => {
      await revokeDevice(id);
      devices = devices.filter((device) => device.id !== id);
    });

  const onSignOut = (): void =>
    void run('head', async () => {
      await signOut();
      // The next account may have a different page: forget the heights, and stop pagehide from
      // storing them again on the way out.
      stopReserveWrites?.();
      clearAccountReserve();
      window.location.assign('/');
    });

  async function onManageBilling(): Promise<void> {
    billingBusy = true;
    failure = null;
    try {
      const result = await openPortal(undefined);
      window.location.assign(result.portal_url);
    } catch {
      failure = { area: 'plan', message: ACCOUNT_FAILED };
      billingBusy = false;
    }
  }

  const onAnonymize = (wanted: boolean): void =>
    void run('you', async () => {
      await setAnonymize(wanted);
      if (me !== null) me = { ...me, user: { ...me.user, anonymize: wanted } };
    });

  // ---------------------------------------------------------------- guilds
  let guildBusy = $state<number | null>(null);
  let leaveAsking = $state<number | null>(null);
  let consentSaved = $state<number | null>(null);
  let consentTimer: ReturnType<typeof setTimeout> | undefined;
  $effect(() => () => clearTimeout(consentTimer));
  /** How many of this account's characters sit in the guild: consent and Leave cover all of them. */
  const guildCharacterCount = (guildId: number): number =>
    me?.characters.filter((c) => c.guild?.id === guildId).length ?? 0;

  const onConsentChange = (guildId: number, event: Event): void => {
    const value = (event.currentTarget as HTMLSelectElement).value as GuildConsent;
    guildBusy = guildId;
    void run('guild', async () => {
      try {
        await updateConsent(guildId, value);
        if (me !== null) {
          me = { ...me, guilds: me.guilds.map((g) => (g.id === guildId ? { ...g, consent: value } : g)) };
        }
        consentSaved = guildId;
        clearTimeout(consentTimer);
        consentTimer = setTimeout(() => (consentSaved = null), CONSENT_SAVED_MS);
      } finally {
        guildBusy = null;
      }
      session.refresh();
    });
  };

  const onLeaveGuild = (guildId: number): void => {
    guildBusy = guildId;
    void run('guild', async () => {
      try {
        await leaveGuild(guildId);
        // The character rows read the same account: the guild line goes with the guild.
        if (me !== null) {
          me = {
            ...me,
            guilds: me.guilds.filter((g) => g.id !== guildId),
            characters: me.characters.map((c) => {
              if (c.guild?.id !== guildId) return c;
              const { guild: _gone, ...rest } = c;
              return rest;
            }),
          };
        }
        leaveAsking = null;
      } finally {
        guildBusy = null;
      }
      session.refresh();
    });
  };

  // ---------------------------------------------------------------- the plan
  // The offer just answered stays on screen as its status line after /v1/me stops sending it.
  let answeredOffer = $state<MeSeatOffer | null>(null);
  const onOfferAnswered = (): void => {
    answeredOffer = me?.seat_offer ?? answeredOffer;
    session.refresh();
  };
  const shownOffer = $derived(answeredOffer ?? me?.seat_offer ?? null);
  const planLine = $derived(me === null ? undefined : accountPlanLine(me, now, viewerZone()));
  const pastDue = $derived(
    me?.entitlements?.billing?.status === 'past_due' && planLine?.billingButton === 'manage',
  );
</script>

<div class="flex flex-col" data-testid="account">
  {#if status === 'loading'}
    <AccountLoading />
  {:else if status === 'failed'}
    <Page layout="split" testid="account-grid">
      {#snippet header()}<PageHeader title={accountPageCopy.title} testid="account-head" />{/snippet}
      <Section title={characterListCopy.heading}>
        <Plate label={characterListCopy.heading}>
          <div class="flex flex-col items-start gap-3 px-5 pt-4 pb-5" data-testid="account-load-error">
            <p class="text-strong text-[15px] leading-6" role="alert">{ACCOUNT_FAILED}</p>
            <button
              type="button"
              class={button('primary')}
              onclick={() => session.refresh()}
              data-testid="account-load-error-retry">{checkoutCopy.retry}</button
            >
          </div>
        </Plate>
      </Section>
    </Page>
  {:else if me === null}
    <Page layout="form" testid="account-form">
      {#snippet header()}<PageHeader title={accountPageCopy.title} testid="account-head" />{/snippet}
      <Plate label="Sign in">
        <div class="flex flex-col items-start gap-3 px-5 pt-4 pb-5" data-testid="account-signin">
          <p class="text-[15px] leading-6">{accountSignInCopy.reason}</p>
          <a class={button('primary')} href={battlenetStartUrl(SIGN_IN_NEXT)} data-testid="battlenet"
            >{myReportsCopy.signInAction}</a
          >
        </div>
      </Plate>
    </Page>
  {:else}
    <Page layout="split" testid="account-grid">
      {#snippet header()}
        <AccountHead {me} {now} busy={busy && failure === null} {onSignOut} />
        {#if errorIn('head') !== ''}<p class="text-wipe px-[var(--page-inset)] text-[13px]" role="alert">
            {errorIn('head')}
          </p>{/if}
        {#if signedInBanner !== ''}
          <div class="px-[var(--page-inset)]">
            <Notice text={signedInBanner} tone="neutral" testid="account-signed-in-banner" />
          </div>
        {/if}
      {/snippet}

      <CharacterList
        characters={me.characters}
        bnetImportedAt={me.bnet_imported_at}
        {mainKey}
        {currentKey}
        onSetMain={setMain}
        logCoach={me.entitlements?.log_coach === true}
        {viaBattlenet}
        {refreshResult}
        {now}
      />
      <AccountPremiumBlock
        plan={planLine}
        {pastDue}
        busy={billingBusy}
        error={errorIn('plan')}
        onBilling={() => void onManageBilling()}
        seePlans={!onPremium(me, now)}
        entitlements={me.entitlements}
        characters={me.characters}
        mainKey={me.main_character_key}
        serverRunsThisMonth={me.server_sims?.runs_this_month}
        notifications={premiumRowVisible(me)
          ? { value: premiumRowValue(me), linked: premiumRowLinked(me) }
          : null}
      >
        {#snippet offer()}
          {#if shownOffer !== null}
            <AccountSeatOffer offer={shownOffer} zone={viewerZone()} onAnswered={onOfferAnswered} />
          {/if}
        {/snippet}
      </AccountPremiumBlock>
      {#if me.guilds.length > 0}
        <Section title={accountPageCopy.guildLabel} anchor="guild" testid="account-guilds-panel">
          <div class="flex flex-col gap-8" data-testid="account-guilds">
            {#each me.guilds as guild (guild.id)}
              <AccountGuildBlock
                {guild}
                characterCount={guildCharacterCount(guild.id)}
                disabled={busy || guildBusy === guild.id}
                asking={leaveAsking === guild.id}
                consentSaved={consentSaved === guild.id}
                onconsent={(event) => onConsentChange(guild.id, event)}
                onask={() => (leaveAsking = guild.id)}
                oncancel={() => (leaveAsking = null)}
                onleave={() => onLeaveGuild(guild.id)}
              />
            {/each}
          </div>
          {#if errorIn('guild') !== ''}<p class="text-wipe pt-2 text-[13px]" role="alert">
              {errorIn('guild')}
            </p>{/if}
        </Section>
      {/if}

      {#snippet rail()}
        <AccountDevices
          {devices}
          {pairing}
          {pairedDevice}
          {busy}
          error={errorIn('devices')}
          {now}
          {onPair}
          {onRevoke}
        />
        <AccountDiscord me={me!} {now} />
        <AccountYou anonymize={me!.user.anonymize} {busy} error={errorIn('you')} {onAnonymize} />
      {/snippet}
    </Page>
  {/if}
</div>

<style>
  /* The page starts 28 px under the head at every width (the mock's split page); the primitive
     tightens to 20 below 1024. */
  :global(.ui-page[data-testid='account-grid']),
  :global(.ui-page[data-testid='account-loading']) {
    padding-top: 28px;
  }
</style>
