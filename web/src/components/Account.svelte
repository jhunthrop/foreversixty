<!-- web/src/components/Account.svelte -->
<!-- One component, four modes, because they are four views of one thing and splitting them
     would mean four copies of the same "who is signed in" fetch:
       nav      the header's sign-in state, on every logs-product page
       login    /login: the Battle.net button and the email form
       account  /account: devices, pairing, characters, anonymize, sign out
       pairing  the pairing code on its own, embedded in /logs
     Each mode renders one branch; the shared parts are the `me` load and the error line. -->
<script lang="ts">
  import {
    ACCOUNT_FAILED,
    EMAIL_SENT,
    battlenetStartUrl,
    fetchMeOnce,
    listDevices,
    pairDevice,
    refreshDevices,
    revokeDevice,
    requestEmailLink,
    setAnonymize,
    setMainCharacter,
    signOut,
    type Device,
    type Me,
    type PairingCode,
  } from '../lib/account/api';
  import AccountGuildBlock from './account/AccountGuildBlock.svelte';
  import { safeNextPath } from '../lib/account/safe-next';
  import { openPortal } from '../lib/billing/api';
  import { premiumRowLinked, premiumRowValue, premiumRowVisible } from '../lib/account/notifications';
  import AccountDiscordLink from './account/AccountDiscordLink.svelte';
  import AccountDiscordBlock from './account/AccountDiscordBlock.svelte';
  import AccountPremiumBlock from './account/AccountPremiumBlock.svelte';
  import AccountSeatOffer from './account/AccountSeatOffer.svelte';
  import type { MeSeatOffer } from '../lib/billing/contract';
  import { accountPlanLine, onPremium } from '../lib/account/plan-line';
  import AccountLoading from './account/AccountLoading.svelte';
  import { createQueryState } from '../lib/data/query.svelte';
  import { characterListCopy } from '../lib/account/character-list-copy';
  import { accountPageCopy } from '../lib/account/account-page-copy';
  import { accountNameOf } from '../lib/character-selector/rows';
  import { heroCharacter } from '../lib/account/hero-character';
  import { characterHref, parseCharacterPath } from '../lib/characters';
  import {
    CURRENT_CHARACTER_CHANGED,
    onCurrentCharacterChange,
    readCurrent,
    writeCurrent,
    type CurrentCharacter,
  } from '../lib/current-character';
  import { mainCharacter, pointerForCharacter } from '../lib/account/main-character';
  import { leaveGuild, updateConsent, type GuildConsent } from '../lib/guild/api';
  import { devicesCopy, pairingEntryLabel } from '../lib/account/devices-copy';
  import { pairingCopy } from '../lib/reports/copy';
  import { createPairingPoll, type PairingPoll } from '../lib/reports/pairing-poll';
  import { API_BASE_URL } from '../lib/planner/config';
  import { PRIMARY_BUTTON_FIXED, SECONDARY_BUTTON_FIXED } from '../lib/planner/styles';
  import { button } from '../lib/ui/button';
  import { ACTION_LINK, ACTION_LINK_CHEVRON } from '../lib/ui/link';
  import { ROW } from '../lib/ui/row';
  import { relativeTime, viewerZone } from '../lib/dates';
  import CharacterHandoffLinks from './CharacterHandoffLinks.svelte';
  import CharacterList from './account/CharacterList.svelte';
  import CharacterIdentity from './character/CharacterIdentity.svelte';
  import CharacterRatingPanel from './CharacterRatingPanel.svelte';
  import MyReports from './MyReports.svelte';
  import SignInPrompt from './SignInPrompt.svelte';
  import AccountTitle from './account/AccountTitle.svelte';
  import CardHead from './ui/CardHead.svelte';
  import Glyph from './ui/Glyph.svelte';
  import IconTile from './ui/IconTile.svelte';
  import LoadError from './ui/LoadError.svelte';
  import SettingsRow from './ui/SettingsRow.svelte';
  import Switch from './ui/Switch.svelte';
  import StatePanel from './ui/StatePanel.svelte';
  import {
    SIGNED_OUT_MIN_H,
    clearAccountReserve,
    recordAccountReserve,
    scheduleAccountReserveWrites,
  } from '../lib/account/layout';
  import { accountSignInCopy } from '../lib/account/signin-copy';

  let { mode, next = '/logs' }: { mode: 'nav' | 'login' | 'account' | 'pairing'; next?: string } = $props();

  // One `/v1/me` read, shared with every other island through the client cache
  // (web/src/lib/data/query.ts) -- see SessionNav.svelte and HomeAccountPanel.svelte's own
  // copies of this same call.
  const session = createQueryState<Me | null>(`${API_BASE_URL}/v1/me`, () => fetchMeOnce(), {
    scope: 'private',
    ttlMs: 10 * 60 * 1000,
  });

  // A writable $derived (Svelte 5): tracks `session.data`, but the guild consent/leave
  // handlers below (onConsentChange, onLeaveGuild) also reassign it directly so the page
  // changes in the same frame; lib/guild/api.ts's updateConsent/leaveGuild drop the cached
  // `/v1/me` (invalidateGuild) and the handlers then re-read it with `session.refresh()`.
  // Any such local override lasts only until `session.data` next changes, at which point
  // this reverts to tracking it again -- setMain/onAnonymize's own edits are harmless duplicates
  // of that: setMainCharacter/setAnonymize (lib/account/api.ts) already push the server's
  // fresh `me` into this same cache key via setQueryData.
  let me = $derived(session.data);

  // Whether the devices list (fetched only for 'account'/'pairing', same gate as before) has
  // loaded, mirroring the granularity `session.status` alone cannot give this component: the
  // old `load()` awaited both the `/v1/me` and the devices read before flipping to 'ready'.
  let devicesStatus = $state<'idle' | 'loading' | 'ready' | 'failed'>('idle');
  let error = $state('');
  let notice = $state('');
  let email = $state('');
  let devices = $state<Device[]>([]);
  let pairing = $state<PairingCode | null>(null);
  let busy = $state(false);
  let guildBusy = $state<number | null>(null);
  // Which guild's Leave is asking "are you sure", and the "Saved." line a consent change
  // prints (shown for three seconds, the same place the guild settings print theirs).
  let leaveAsking = $state<number | null>(null);
  let consentSaved = $state<number | null>(null);
  let consentTimer: ReturnType<typeof setTimeout> | undefined;
  /** How many of this account's characters sit in the guild: consent and Leave cover all of them. */
  const guildCharacterCount = (guildId: number): number =>
    me?.characters.filter((c) => c.guild?.id === guildId).length ?? 0;

  // Logs landing spec (2026-10-04) §4.D.2: once a code is shown, poll for the companion
  // actually using it. `pairedDevice` is the poll's own success state, distinct from
  // `devices` (which it also appends to, so the devices panel's own list stays current
  // without a second fetch); `activePoll` is not `$state` -- it is plumbing, never read by
  // a template, and `$state`-wrapping a mutable timer handle would gain nothing.
  let pairedDevice = $state<Device | null>(null);
  /** The device whose Revoke is asking "Revoke <name>?"; null when no row is asking. */
  let revoking = $state<string | null>(null);
  let activePoll: PairingPoll | null = null;

  $effect(() => {
    return () => activePoll?.stop();
  });

  /** Guilds left on this page. The character cards' guild line goes with the GUILD section
   *  in the same frame; the refetched `/v1/me` then confirms it. */
  let leftGuildIds = $state<number[]>([]);
  const cardCharacters = $derived(
    me === null
      ? []
      : me.characters.map((c) =>
          c.guild !== undefined && leftGuildIds.includes(c.guild.id) ? { ...c, guild: undefined } : c,
        ),
  );

  const signedIn = $derived(me !== null);
  const displayName = $derived(me?.user.battletag ?? me?.user.email ?? 'Your account');
  // Optional chaining all the way through: `entitlements` itself may be absent on an
  // older/stubbed /v1/me response (see the `Me.entitlements` doc comment in account/api.ts).
  const billing = $derived(me?.entitlements?.billing ?? null);
  // The offer just answered stays on screen as its status line after /v1/me stops sending it.
  let answeredOffer = $state<MeSeatOffer | null>(null);
  const onOfferAnswered = (): void => {
    answeredOffer = me?.seat_offer ?? answeredOffer;
    session.refresh();
  };
  const shownOffer = $derived(answeredOffer ?? me?.seat_offer ?? null);
  const planLine = $derived(me === null ? undefined : accountPlanLine(me, new Date(), viewerZone()));
  const isSupporter = $derived(me?.entitlements?.supporter_mark === true);
  const signedInMethod = $derived(
    me?.user.battletag !== undefined && me?.user.battletag !== null
      ? accountPageCopy.signedInWithBattlenet
      : accountPageCopy.signedInByEmail,
  );

  // The hero band's own read of the current-character pointer (brief 2026-09-22 §B4/B5) --
  // the same localStorage read CurrentCharacterBar.svelte makes for the compact chip, kept
  // separate because that component stays a pure render of its own `current` prop and
  // never exposes it. Re-read on CURRENT_CHARACTER_CHANGED so Forget clears the band too.
  let currentCharacter = $state<CurrentCharacter | null>(null);
  $effect(() => {
    if (mode !== 'account') return;
    const read = (): void => {
      currentCharacter = readCurrent();
    };
    read();
    return onCurrentCharacterChange(read);
  });
  // The hero is the current character when one is pointed at, else the account's main
  // (chosen, or the site's guess): the main is the default context everywhere.
  const hero = $derived(
    me === null
      ? null
      : (heroCharacter(currentCharacter, me.characters) ??
          mainCharacter(me.characters, me.main_character_key)),
  );
  /** "Your ratings" shows only once the hero character has a rated fight. */
  let ratingsShown = $state(false);
  const heroPath = $derived(hero === null ? null : parseCharacterPath(`/character/${hero.key}`));
  // spec 2026-09-22 §3.1: "the sentence that Blizzard serves no data for this realm type
  // when no character has a build" -- account-wide, not just the hero, since it is telling
  // the player why the handoff links below are stuck on the paste fallback for every
  // character, not only this one.
  const noBattlenetData = $derived(me !== null && me.characters.every((c) => c.build === undefined));

  /** Devices panel's Updated stamp: the most recently seen device, or '' when none has
   *  ever reported in (brief 2026-09-22 §B2: "an Updated stamp ... where a timestamp
   *  exists"). Guilds carries no verified-at timestamp in the /v1/me contract today (only
   *  a boolean, auth/store.go's CharacterGuild/Guild), so that panel shows none -- nothing
   *  invented. */
  const devicesUpdated = $derived.by(() => {
    const seen = devices.map((d) => d.last_seen_at).filter((at): at is string => at !== null);
    if (seen.length === 0) return '';
    const latest = seen.reduce((a, b) => (a > b ? a : b));
    return relativeTime(new Date(latest));
  });

  // The combined status every mode below reads: 'loading'/'failed' first reflect the
  // `/v1/me` read itself, then -- only for 'account'/'pairing', the same gate `load()` used
  // to apply around its own `listDevices()` call -- the devices read alongside it, so this
  // stays 'loading' until both reads that used to be one sequential `await` chain have
  // settled, and 'failed' if either one did.
  // A `session.refresh()` after an action (a seat offer answered, consent saved, a guild left)
  // re-reads `/v1/me` with the account already on screen: that read keeps the page as it is,
  // so the panel that acted (the seat offer's status line) is not unmounted under the reader.
  const rereading = $derived(session.status === 'loading' && session.data !== null);
  const status = $derived.by<'loading' | 'ready' | 'failed'>(() => {
    if (session.status === 'failed') return 'failed';
    if (session.status !== 'ready' && !rereading) return 'loading';
    const needsDevices = (mode === 'account' || mode === 'pairing') && me !== null;
    if (!needsDevices) return 'ready';
    if (devicesStatus === 'failed') return 'failed';
    if (devicesStatus === 'ready') return 'ready';
    return 'loading';
  });

  // Devices, for 'account'/'pairing' only, once `me` is known -- the second half of the old
  // `load()`'s sequential await chain, now driven by `me` changing instead.
  // Keyed on whether `me` exists, not on the object: an in-place patch of it (the Discord block
  // writes its saved state into the cache) must not drop the page back to its loading state.
  const hasMe = $derived(me !== null);

  // The ready page's two column heights, stored so the next visit's loading state reserves them
  // (spec 2026-10-25-discord-notifications 9.1; Base.astro reads them before first paint).
  // Stops the reserve writer; set while the ready account page is mounted, called by sign-out.
  let stopReserveWrites: (() => void) | null = null;
  const readyAccount = $derived(mode === 'account' && status === 'ready' && signedIn);
  $effect(() => {
    if (!readyAccount) return;
    const stop = scheduleAccountReserveWrites(() =>
      recordAccountReserve(
        document.querySelector('[data-testid="account-left-col"]'),
        document.querySelector('[data-testid="account-rail-col"]'),
      ),
    );
    stopReserveWrites = stop;
    return () => {
      stop();
      stopReserveWrites = null;
    };
  });
  $effect(() => {
    if (!hasMe || (mode !== 'account' && mode !== 'pairing')) {
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

  // `error` here is the *load* failure's message; run()'s own actions (below) overwrite it
  // with their own message once the page is past 'loading'/'failed', exactly as `load()`
  // used to reset it at the top of every attempt and set it in its own catch.
  $effect(() => {
    if (status === 'failed') {
      error = ACCOUNT_FAILED;
    } else if (status === 'loading') {
      error = '';
    }
  });

  // The static build has no per-request server, so Astro frontmatter never sees a real
  // visitor's query string -- it only ever runs once, at build time. `next` (the prop) is
  // therefore always the caller's hardcoded fallback. Read the real `?next=` here instead,
  // client-side after hydration, which is the one place in this architecture that runs on
  // the visitor's own request. $effect does not run during SSR, so this is safe without a
  // `typeof window` guard.
  let resolvedNext = $state(next);

  $effect(() => {
    if (mode !== 'login') return;
    const params = new URLSearchParams(window.location.search);
    resolvedNext = safeNextPath(params.get('next'), next);
  });

  // Spec 2026-09-22 §7.3: a second Battle.net login (the refresh link) redirects back to
  // ?refreshed=1. Shown once, then the query string is dropped with replaceState so a
  // reload or a shared link never re-shows a stale toast. Same "$effect, no typeof window
  // guard" reasoning as the `next`-param effect above.
  let toast = $state('');

  $effect(() => {
    if (mode !== 'account') return;
    const params = new URLSearchParams(window.location.search);
    if (params.get('refreshed') !== '1') return;
    toast = characterListCopy.refreshedToast;
    window.history.replaceState({}, '', window.location.pathname);
  });

  // Hub arrival (spec 2026-09-22 §3.1): a fresh sign-in lands on ?signed_in=1, picks the
  // main character, writes it as the site's current-character pointer, and shows a one-line
  // banner naming it. Same "read the query string, act once, then replaceState it away"
  // idiom as the ?refreshed=1 toast above -- replaceState is what keeps this from re-firing
  // on a later `me` reassignment, not extra state.
  let signedInBanner = $state('');

  $effect(() => {
    if (mode !== 'account' || me === null) return;
    const params = new URLSearchParams(window.location.search);
    if (params.get('signed_in') !== '1') return;
    const main = mainCharacter(me.characters, me.main_character_key);
    if (main !== null) {
      const pointer = pointerForCharacter(main);
      writeCurrent(pointer);
      window.dispatchEvent(new Event(CURRENT_CHARACTER_CHANGED));
      signedInBanner = accountPageCopy.signedInBanner(main.name);
    } else {
      signedInBanner = accountPageCopy.signedInBannerEmpty(accountNameOf(me));
    }
    window.history.replaceState({}, '', window.location.pathname);
  });

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

  async function run(action: () => Promise<void>): Promise<void> {
    busy = true;
    error = '';
    notice = '';
    try {
      await action();
    } catch (thrown) {
      error = thrown instanceof Error ? thrown.message : ACCOUNT_FAILED;
    } finally {
      busy = false;
    }
  }

  const onEmail = (event: SubmitEvent): void => {
    event.preventDefault();
    void run(async () => {
      await requestEmailLink(email.trim());
      notice = EMAIL_SENT;
    });
  };

  const onPair = (): void =>
    void run(async () => {
      activePoll?.stop();
      pairedDevice = null;
      const beforeIds = new Set(devices.map((device) => device.id));
      const code = await pairDevice();
      pairing = code;
      activePoll = createPairingPoll(
        () => refreshDevices(),
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
    void run(async () => {
      await revokeDevice(id);
      devices = devices.filter((device) => device.id !== id);
    });

  const onSignOut = (): void =>
    void run(async () => {
      await signOut();
      // The next account may have a different page: forget the heights, and stop pagehide
      // from storing them again on the way out.
      stopReserveWrites?.();
      clearAccountReserve();
      window.location.assign('/');
    });

  const onManageBilling = (): void =>
    void run(async () => {
      const result = await openPortal(undefined);
      window.location.assign(result.portal_url);
    });

  const onAnonymize = (wanted: boolean): void => {
    void run(async () => {
      await setAnonymize(wanted);
      if (me !== null) me = { ...me, user: { ...me.user, anonymize: wanted } };
    });
  };

  const onConsentChange = (guildId: number, event: Event): void => {
    const value = (event.currentTarget as HTMLSelectElement).value as GuildConsent;
    guildBusy = guildId;
    void run(async () => {
      try {
        await updateConsent(guildId, value);
        if (me !== null) {
          me = {
            ...me,
            guilds: me.guilds.map((g) => (g.id === guildId ? { ...g, consent: value } : g)),
          };
        }
        consentSaved = guildId;
        clearTimeout(consentTimer);
        consentTimer = setTimeout(() => (consentSaved = null), 3000);
      } finally {
        guildBusy = null;
      }
      session.refresh();
    });
  };

  const onLeaveGuild = (guildId: number): void => {
    guildBusy = guildId;
    void run(async () => {
      try {
        await leaveGuild(guildId);
        // The character cards read the same source: the guild line goes with the guild.
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
        leftGuildIds = [...leftGuildIds, guildId];
        leaveAsking = null;
      } finally {
        guildBusy = null;
      }
      session.refresh();
    });
  };
</script>

{#if mode === 'nav'}
  <div class="flex items-center gap-3 text-[13px]" data-testid="session-nav">
    {#if status === 'loading'}
      <!--
        The placeholder is the signed-out label, hidden rather than absent, so the slot is
        exactly the width and height it will have once the session resolves. A narrower
        placeholder lets the header's 1fr brand column stretch, and the brand then wraps to
        two lines and un-wraps on hydration: a 16px header growth that moves the whole page
        and was the entirety of /logs' layout shift.
      -->
      <span class="invisible inline-flex min-h-11 items-center px-2 md:min-h-0 md:px-0" aria-hidden="true">
        Sign in
      </span>
    {:else if signedIn}
      <a href="/account" class="text-nav hover:text-strong inline-flex min-h-11 items-center md:min-h-0">
        {displayName}
      </a>
    {:else}
      <!-- "Sign in" is short enough (38px) to miss the 44px hit-target floor on its width
           alone; the padding here has to match the invisible placeholder above exactly (its
           own comment explains why) so the swap from placeholder to real link never shifts
           the header. -->
      <a
        href="/login"
        class="text-nav hover:text-strong inline-flex min-h-11 items-center px-2 md:min-h-0 md:px-0"
      >
        Sign in
      </a>
    {/if}
  </div>
{:else if mode === 'login'}
  <div class="flex flex-col gap-6" data-testid="login">
    {#if signedIn}
      <p class="text-[14px]">Signed in as {displayName}.</p>
      <a class="{PRIMARY_BUTTON_FIXED} self-start px-5" href="/account" data-testid="login-go-account">
        Your account
      </a>
    {:else}
      <p class="text-muted text-[14px]">{accountSignInCopy.reason}</p>
      <div class="flex flex-col gap-3">
        <a
          class="{PRIMARY_BUTTON_FIXED} self-start px-5"
          href={battlenetStartUrl(resolvedNext)}
          data-testid="battlenet"
        >
          Sign in with Battle.net
        </a>
        <p class="text-muted text-[13px]">
          Battle.net tells the site your BattleTag and your characters. It never sees a password here.
        </p>
      </div>
      <form class="flex flex-col gap-3" onsubmit={onEmail} data-testid="email-form">
        <label class="quiet-label" for="account-email">Or a sign-in link by email</label>
        <div class="flex flex-col gap-3 sm:flex-row">
          <input
            id="account-email"
            class="border-line-warm bg-raised rounded-control text-text h-11 flex-1 px-3 text-[15px]"
            type="email"
            autocomplete="email"
            placeholder="you@example.com"
            required
            bind:value={email}
          />
          <button
            class="{SECONDARY_BUTTON_FIXED} border-line-warm-strong text-strong px-4"
            type="submit"
            disabled={busy}
          >
            Send link
          </button>
        </div>
      </form>
    {/if}
    {#if notice !== ''}<p class="text-[14px]" data-testid="account-notice">{notice}</p>{/if}
    <!-- 21px is the measured line height of the text-[14px] error line: reserved so the
         session load landing (or a `run()` action failing) never shoves the page down after
         first paint, the same reason `mode === 'nav'` renders `&nbsp;` while
         `status === 'loading'`. Repeated for the pairing and account modes below. -->
    <div class="min-h-[21px]">
      {#if error !== ''}<p class="text-[14px]" role="alert" data-testid="account-error">{error}</p>{/if}
    </div>
  </div>
{:else if mode === 'pairing'}
  <!-- Signed out and signed in are the same shape on purpose, a line over a button, so the
       block is the same height either way and the steps under it never move when the
       session answers. The floor holds that height while it is still being asked for. -->
  <div class="flex min-h-[150px] flex-col items-start gap-3" data-testid="pairing">
    {#if status === 'loading'}
      <p class="text-muted text-[14px]">Checking whether you are signed in.</p>
    {:else if !signedIn}
      <p class="text-[14px]" data-testid="pairing-signed-out">{pairingCopy.signedOut}</p>
    {:else if pairedDevice !== null}
      <!-- §4.D.2: replaces the code block in place once the companion actually uses it;
           the expiry line is answered and removed with it. -->
      <p class="text-[14px]" data-testid="pairing-success">{pairingCopy.success(pairedDevice.name)}</p>
    {:else if pairing === null}
      <p class="text-muted text-[13px]" data-testid="pairing-companion-first">{pairingCopy.companionFirst}</p>
      <p class="text-[14px]">
        Signed in as {displayName}. For the fallback, show a code and type it into the companion.
      </p>
      <button
        class="text-nav inline-flex min-h-11 items-center text-[13px] underline md:min-h-0"
        onclick={onPair}
        disabled={busy}
      >
        Show pairing code
      </button>
    {:else}
      <p class="tabular text-strong font-mono text-[24px]" data-testid="pairing-code">{pairing.code}</p>
      <p class="text-muted text-[13px]">
        Type this into the companion within {Math.round(pairing.expires_in / 60)} minutes. It pairs one device and
        cannot be reused.
      </p>
    {/if}
    <div class="min-h-[21px]">
      {#if error !== ''}<p class="text-[14px]" role="alert" data-testid="account-error">{error}</p>{/if}
    </div>
  </div>
{:else}
  <div class="flex flex-col gap-8" data-testid="account">
    {#if status === 'loading'}
      <AccountLoading />
    {:else if status === 'failed'}
      <div class="account-header-reserve">
        <AccountTitle />
      </div>
      <LoadError message={error} onRetry={() => session.refresh()} testid="account-load-error" />
    {:else if !signedIn}
      <div class="account-header-reserve">
        <AccountTitle />
      </div>
      <div class="session-reserve {SIGNED_OUT_MIN_H}">
        <SignInPrompt line={accountSignInCopy.reason} testid="account-signin" />
      </div>
    {:else}
      <div class="reveal flex flex-col gap-8">
        <!-- Page header (polish spec 4.3): eyebrow and title left; at lg the identity tile, the
             BattleTag over the sign-in method, and Sign out right, stacked under the title on phone. -->
        <AccountTitle aside={identity} />
        {#snippet identity()}
          <div class="flex items-center gap-3" data-testid="account-identity">
            <IconTile glyph="user" />
            <div class="flex min-w-0 flex-col">
              <p class="text-strong text-[15px] leading-5 font-semibold [overflow-wrap:anywhere]">
                {displayName}
              </p>
              <p class="text-muted text-[13px] leading-5">{signedInMethod}</p>
            </div>
            <button class="{button('secondary', 'sm')} ml-2" onclick={onSignOut} disabled={busy}>
              <Glyph name="signout" />{accountPageCopy.signOut}
            </button>
          </div>
        {/snippet}

        {#if signedInBanner !== ''}
          <div
            class="bg-raised border-gold flex flex-col gap-1 border-l-2 px-4 py-3"
            data-testid="account-signed-in-banner"
          >
            <p class="text-[14px]">{signedInBanner}</p>
          </div>
        {/if}

        {#if toast !== ''}
          <div
            class="bg-raised border-gold flex flex-col gap-1 border-l-2 px-4 py-3"
            data-testid="account-toast"
          >
            <p class="text-[14px]">{toast}</p>
            {#if me!.bnet_imported_at !== undefined}
              <p class="text-muted text-[13px]">
                {characterListCopy.importedFrom(relativeTime(new Date(me!.bnet_imported_at)))}
              </p>
            {/if}
          </div>
        {/if}

        <div class="flex flex-col gap-8 lg:grid lg:grid-cols-[minmax(0,1fr)_440px] lg:items-start lg:gap-8">
          <div class="flex flex-col gap-8" data-testid="account-left-col">
            {#if hero !== null}
              <!-- The current character (brief §B4, loosened §3.1): shown whenever the
                   compact chip's pointer matches a listed character; the render below is
                   shown only when that character has one -- an addon-only import still gets
                   a hero band, just without art. -->
              <section
                class="card-surface flex flex-col-reverse items-start gap-4 overflow-hidden px-5 py-5 md:px-6 lg:flex-row lg:items-end lg:justify-between"
                data-testid="account-hero"
              >
                <div class="flex min-w-0 flex-1 flex-col gap-4 self-stretch">
                  <h2 class="text-muted text-[11px] leading-none font-bold tracking-[.14em] uppercase">
                    {accountPageCopy.currentCharacterLabel}
                  </h2>
                  <CharacterIdentity
                    character={hero}
                    size="xxl"
                    descriptor="full"
                    supporter={isSupporter}
                    href={characterHref(hero.region, hero.ruleset, hero.name)}
                    testid="account-hero"
                  />
                  {#if hero.build === undefined}
                    <div
                      class="border-line-card flex flex-wrap items-center justify-between gap-3 rounded-[var(--radius-control)] border border-dashed px-4 py-3"
                      data-testid="account-hero-no-export"
                    >
                      <p class="text-text min-w-0 flex-1 text-[14px]">{accountPageCopy.noExportPrompt}</p>
                      <a
                        class={button('primary')}
                        href={characterListCopy.pasteHref}
                        data-testid="account-hero-paste"
                      >
                        {characterListCopy.pasteAnExport}
                      </a>
                    </div>
                  {:else if heroPath !== null}
                    <div class="flex min-h-11 flex-wrap items-center gap-3 md:min-h-0">
                      <CharacterHandoffLinks path={heroPath} buttons />
                      <a
                        class="inline-flex min-h-11 items-center text-[13px] font-semibold md:min-h-0"
                        href="/logs"
                      >
                        {accountPageCopy.heroLogs}
                      </a>
                    </div>
                  {/if}
                  {#if noBattlenetData}
                    <p class="text-muted text-[13px]" data-testid="account-hero-no-bnet-data">
                      {accountPageCopy.noBattlenetDataForRealm}
                    </p>
                  {/if}
                </div>
                {#if hero.render_url !== undefined}
                  <img
                    class="max-h-[280px] w-auto object-contain lg:max-h-[360px]"
                    src={hero.render_url}
                    alt=""
                    loading="lazy"
                    data-testid="account-hero-render"
                  />
                {/if}
              </section>
            {/if}

            <div id="characters">
              <CharacterList
                characters={cardCharacters}
                bnetImportedAt={me!.bnet_imported_at}
                mainKey={mainCharacter(me!.characters, me!.main_character_key)?.key}
                currentKey={hero?.key}
                onSetMain={setMain}
                logCoach={me?.entitlements?.log_coach === true}
              />
            </div>

            <section class="card-surface overflow-hidden" data-testid="account-reports">
              <CardHead glyph="chart" title={accountPageCopy.reportsLabel} />
              <div class="border-line-soft border-t px-5 py-4">
                <MyReports heading={false} nightDates />
              </div>
            </section>
          </div>

          <div class="flex flex-col gap-8" data-testid="account-rail-col">
            {#if heroPath !== null}
              <!-- Hidden (not unmounted: the panel's own fetch decides) until a rating exists. -->
              <div hidden={!ratingsShown} data-testid="account-ratings-slot">
                <StatePanel label={accountPageCopy.yourRatingsLabel} testid="account-ratings">
                  <div class="p-[18px]">
                    <CharacterRatingPanel
                      path={heroPath}
                      heading={false}
                      onshown={(shown) => (ratingsShown = shown)}
                    />
                  </div>
                </StatePanel>
              </div>
            {/if}

            <div id="devices">
              <section class="card-surface overflow-hidden" data-testid="account-devices">
                <CardHead glyph="device" title={accountPageCopy.devicesLabel}>
                  {#snippet aside()}
                    {#if devicesUpdated !== ''}
                      <span class="text-muted" data-testid="account-devices-updated">
                        Updated {devicesUpdated}
                      </span>
                    {/if}
                  {/snippet}
                </CardHead>
                {#if devices.length === 0}
                  <p class="text-muted border-line-soft border-t px-5 py-4 text-[14px]">
                    {accountPageCopy.noDevices}
                  </p>
                {:else}
                  <ul class="m-0 flex list-none flex-col p-0">
                    {#each devices as device (device.id)}
                      <li class="{ROW} justify-between gap-y-2">
                        {#if revoking === device.id}
                          <span class="min-w-0 flex-1 text-[14px]" data-testid="revoke-question">
                            {devicesCopy.revokeQuestion(device.name)}
                          </span>
                          <span class="flex gap-2">
                            <button
                              class={button('danger', 'sm')}
                              onclick={() => {
                                revoking = null;
                                onRevoke(device.id);
                              }}
                              disabled={busy}
                              data-testid="revoke-confirm"
                            >
                              {devicesCopy.revoke}
                            </button>
                            <button
                              class={button('ghost', 'sm')}
                              onclick={() => (revoking = null)}
                              data-testid="revoke-cancel"
                            >
                              {devicesCopy.cancel}
                            </button>
                          </span>
                        {:else}
                          <span class="min-w-0 flex-1 text-[15px] font-semibold">
                            {device.name}
                            <span class="text-muted text-[13px] font-normal">· {device.platform}</span>
                          </span>
                          <button
                            class={button('secondary', 'sm')}
                            onclick={() => (revoking = device.id)}
                            disabled={busy}
                            data-testid="revoke"
                          >
                            {devicesCopy.revoke}
                          </button>
                        {/if}
                      </li>
                    {/each}
                  </ul>
                {/if}
                <div class="border-line-soft flex flex-col items-start gap-1 border-t px-5 py-4">
                  {#if pairedDevice !== null}
                    <p class="text-[14px]" data-testid="pairing-success">
                      {pairingCopy.success(pairedDevice.name)}
                    </p>
                  {:else if pairing === null}
                    <!-- No companion yet: the download is the button, the code a text link under it. -->
                    {#if devices.length === 0}
                      <a
                        class={button('secondary')}
                        href={accountPageCopy.downloadCompanionHref}
                        data-testid="account-download-companion"
                      >
                        {accountPageCopy.downloadCompanion}
                      </a>
                    {/if}
                    <button
                      class={ACTION_LINK}
                      onclick={onPair}
                      disabled={busy}
                      data-testid="account-get-pairing-code"
                    >
                      {pairingEntryLabel(devices.length, accountPageCopy.pairADevice)}<Glyph
                        name="chevron"
                        class={ACTION_LINK_CHEVRON}
                      />
                    </button>
                  {:else}
                    <p class="text-muted text-[13px]" data-testid="pairing-fallback-line">
                      {devicesCopy.codeFallbackLine}
                    </p>
                    <p class="tabular text-strong font-mono text-[24px]" data-testid="pairing-code">
                      {pairing.code}
                    </p>
                    <p class="text-muted text-[13px]">
                      Type this into the companion within {Math.round(pairing.expires_in / 60)} minutes.
                    </p>
                  {/if}
                </div>
              </section>
            </div>

            <section class="card-surface scroll-mt-20 overflow-hidden" id="you" data-testid="account-you">
              <CardHead glyph="user" title={accountPageCopy.youLabel} />
              <SettingsRow label={accountPageCopy.pseudonymLabel} sub={accountPageCopy.pseudonymNote}>
                {#snippet control(labelId)}
                  <Switch
                    checked={me!.user.anonymize}
                    labelledby={labelId}
                    disabled={busy}
                    onchange={onAnonymize}
                    testid="anonymize"
                  />
                {/snippet}
              </SettingsRow>
              <div class="px-5 [&:not(:empty)]:pt-4 [&:not(:empty)]:pb-5"><AccountDiscordLink {me} /></div>
            </section>

            <div id="plan" class="flex flex-col gap-4">
              {#if me!.guilds.length > 0}
                <section class="card-surface overflow-hidden" data-testid="account-guilds-panel">
                  <CardHead glyph="guild" title={accountPageCopy.guildLabel} />
                  <ul class="flex flex-col" data-testid="account-guilds">
                    {#each me!.guilds as guild (guild.id)}
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
                  </ul>
                </section>
              {/if}
              <AccountPremiumBlock
                plan={planLine}
                pastDue={billing?.status === 'past_due' && planLine?.billingButton === 'manage'}
                {busy}
                onBilling={onManageBilling}
                seePlans={!onPremium(me)}
                entitlements={me?.entitlements}
                characters={me?.characters ?? []}
                mainKey={me?.main_character_key}
                serverRunsThisMonth={me?.server_sims?.runs_this_month}
                notifications={me !== null && premiumRowVisible(me)
                  ? { value: premiumRowValue(me), linked: premiumRowLinked(me) }
                  : null}
              >
                {#snippet offer()}
                  {#if shownOffer !== null}
                    <div class="px-5 pb-4">
                      <AccountSeatOffer offer={shownOffer} zone={viewerZone()} onAnswered={onOfferAnswered} />
                    </div>
                  {/if}
                {/snippet}
              </AccountPremiumBlock>
              {#if me !== null}
                <AccountDiscordBlock {me} quiet={(me.seat_offer ?? null) !== null} />
              {/if}
            </div>
          </div>
        </div>
      </div>
    {/if}
    <div class="min-h-[21px]">
      {#if status !== 'failed' && error !== ''}<p
          class="text-[14px]"
          role="alert"
          data-testid="account-error"
        >
          {error}
        </p>{/if}
    </div>
  </div>
{/if}
