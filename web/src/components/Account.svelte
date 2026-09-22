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
    revokeDevice,
    requestEmailLink,
    setAnonymize,
    signOut,
    type Device,
    type Me,
    type PairingCode,
  } from '../lib/account/api';
  import { safeNextPath } from '../lib/account/safe-next';
  import { openPortal } from '../lib/billing/api';
  import { billingBlockCopy } from '../lib/billing/copy';
  import { characterListCopy } from '../lib/account/character-list-copy';
  import { accountPageCopy } from '../lib/account/account-page-copy';
  import { characterDescriptor } from '../lib/account/character-descriptor';
  import { heroCharacter } from '../lib/account/hero-character';
  import { characterHref, guildHref, parseCharacterPath } from '../lib/characters';
  import { CURRENT_CHARACTER_CHANGED, readCurrent, type CurrentCharacter } from '../lib/current-character';
  import { classColorVar } from '../lib/report/format';
  import { leaveGuild, updateConsent, type GuildConsent } from '../lib/guild/api';
  import { guildConsentCopy } from '../lib/guild/copy';
  import { SECONDARY_BUTTON_FIXED } from '../lib/planner/styles';
  import { relativeTime } from '../lib/dates';
  import CharacterHandoffLinks from './CharacterHandoffLinks.svelte';
  import CharacterList from './account/CharacterList.svelte';
  import CurrentCharacterBar from './CurrentCharacterBar.svelte';
  import MyReports from './MyReports.svelte';
  import SignInPrompt from './SignInPrompt.svelte';
  import Skeleton from './ui/Skeleton.svelte';
  import LoadError from './ui/LoadError.svelte';
  import StatePanel from './ui/StatePanel.svelte';
  import {
    CHARACTERS_SKELETON_MIN_H,
    IDENTITY_SKELETON_MIN_H,
    MORE_SKELETON_MIN_H,
    SIGNED_OUT_MIN_H,
  } from '../lib/account/layout';
  import { accountSignInCopy } from '../lib/account/signin-copy';

  let { mode, next = '/logs' }: { mode: 'nav' | 'login' | 'account' | 'pairing' | 'reports'; next?: string } =
    $props();

  let me = $state<Me | null>(null);
  let status = $state<'loading' | 'ready' | 'failed'>('loading');
  let error = $state('');
  let notice = $state('');
  let email = $state('');
  let devices = $state<Device[]>([]);
  let pairing = $state<PairingCode | null>(null);
  let busy = $state(false);
  let guildBusy = $state<number | null>(null);

  const signedIn = $derived(me !== null);
  const displayName = $derived(me?.user.battletag ?? me?.user.email ?? 'Your account');
  // Optional chaining all the way through: `entitlements` itself may be absent on an
  // older/stubbed /v1/me response (see the `Me.entitlements` doc comment in account/api.ts).
  const billing = $derived(me?.entitlements?.billing ?? null);
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
    window.addEventListener(CURRENT_CHARACTER_CHANGED, read);
    return () => window.removeEventListener(CURRENT_CHARACTER_CHANGED, read);
  });
  const hero = $derived(me === null ? null : heroCharacter(currentCharacter, me.characters));
  const heroPath = $derived(hero === null ? null : parseCharacterPath(`/character/${hero.key}`));

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

  async function load(): Promise<void> {
    status = 'loading';
    error = '';
    try {
      me = await fetchMeOnce();
      if (me !== null && (mode === 'account' || mode === 'pairing')) devices = await listDevices();
      status = 'ready';
    } catch {
      status = 'failed';
      error = ACCOUNT_FAILED;
    }
  }

  // One load on mount. $effect rather than onMount so the component works identically
  // whether Astro hydrates it or the report island mounts it by hand.
  $effect(() => {
    void load();
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
      pairing = await pairDevice();
    });

  const onRevoke = (id: string): void =>
    void run(async () => {
      await revokeDevice(id);
      devices = devices.filter((device) => device.id !== id);
    });

  const onSignOut = (): void =>
    void run(async () => {
      await signOut();
      window.location.assign('/');
    });

  const onManageBilling = (): void =>
    void run(async () => {
      const result = await openPortal(undefined);
      window.location.assign(result.portal_url);
    });

  const onAnonymize = (event: Event): void => {
    const wanted = (event.currentTarget as HTMLInputElement).checked;
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
      } finally {
        guildBusy = null;
      }
    });
  };

  const onLeaveGuild = (guildId: number): void => {
    guildBusy = guildId;
    void run(async () => {
      try {
        await leaveGuild(guildId);
        if (me !== null) me = { ...me, guilds: me.guilds.filter((g) => g.id !== guildId) };
      } finally {
        guildBusy = null;
      }
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
      <p class="text-[14px]">
        Signed in as {displayName}. <a href="/account">Your account</a>.
      </p>
    {:else}
      <div class="flex flex-col gap-3">
        <a
          class="{SECONDARY_BUTTON_FIXED} border-line-warm-strong text-strong px-4"
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
        <label class="label text-muted" for="account-email">Or a sign-in link by email</label>
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
         `load()` call above landing (or a `run()` action failing) never shoves the page
         down after first paint, the same reason `mode === 'nav'` renders `&nbsp;` while
         `status === 'loading'`. Repeated for the pairing and account modes below. -->
    <div class="min-h-[21px]">
      {#if error !== ''}<p class="text-[14px]" role="alert" data-testid="account-error">{error}</p>{/if}
    </div>
  </div>
{:else if mode === 'reports'}
  <!-- /logs only, inside that page's own "Your reports" panel, which is why MyReports is
       told to leave its heading off. The floor is one row tall so the panel does not jump
       when the session answers. -->
  <div class="min-h-[88px]">
    {#if status === 'loading'}
      <p class="text-muted text-[14px]">Loading your reports.</p>
    {:else}
      <MyReports {signedIn} heading={false} />
    {/if}
  </div>
{:else if mode === 'pairing'}
  <!-- Signed out and signed in are the same shape on purpose, a line over a button, so the
       block is the same height either way and the steps under it never move when the
       session answers. The floor holds that height while it is still being asked for. -->
  <div class="flex min-h-[110px] flex-col items-start gap-3" data-testid="pairing">
    {#if status === 'loading'}
      <p class="text-muted text-[14px]">Checking whether you are signed in.</p>
    {:else if !signedIn}
      <SignInPrompt line="Sign in to pair the companion with your account." testid="pairing-signin" />
    {:else if pairing === null}
      <p class="text-[14px]">Signed in as {displayName}. Show a code, then type it into the companion.</p>
      <button
        class="{SECONDARY_BUTTON_FIXED} border-line-warm-strong text-strong px-4"
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
      <Skeleton lines={3} minHeight={CHARACTERS_SKELETON_MIN_H} testid="account-characters-skeleton" />
      <Skeleton lines={4} minHeight={IDENTITY_SKELETON_MIN_H} testid="account-identity-skeleton" />
      <Skeleton lines={3} minHeight={MORE_SKELETON_MIN_H} testid="account-more-skeleton" />
    {:else if status === 'failed'}
      <LoadError message={error} onRetry={() => void load()} testid="account-load-error" />
    {:else if !signedIn}
      <div class={SIGNED_OUT_MIN_H}>
        <SignInPrompt line={accountSignInCopy.reason} next="/account" testid="account-signin" />
      </div>
    {:else}
      <div class="reveal flex flex-col gap-8">
        <!-- Page header (brief 2026-09-22 §B1): title left, identity line + Sign out right
             at lg, stacked under the title on phone. -->
        <div class="flex flex-col gap-3 lg:flex-row lg:items-start lg:justify-between">
          <h1 class="section-title text-[18px]">{accountPageCopy.title}</h1>
          <div class="flex flex-col items-start gap-2 lg:items-end" data-testid="account-identity">
            <p class="text-strong text-[14px]">{displayName}</p>
            <p class="text-muted text-[13px]">{signedInMethod}</p>
            <button
              class="{SECONDARY_BUTTON_FIXED} border-line-warm text-text px-4"
              onclick={onSignOut}
              disabled={busy}
            >
              {accountPageCopy.signOut}
            </button>
          </div>
        </div>

        <CurrentCharacterBar compact />

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

        <div class="flex flex-col gap-8 lg:grid lg:grid-cols-12 lg:items-start lg:gap-8">
          <div class="flex flex-col gap-8 lg:col-span-8">
            {#if hero !== null}
              <!-- The current character's render (brief §B4): shown only when the compact
                   chip's pointer matches a listed character that has one -- nothing invented. -->
              <div
                class="border-line-soft flex flex-col-reverse items-start gap-4 overflow-hidden border-b bg-[var(--color-bg)] pb-6 lg:flex-row lg:items-end lg:justify-between"
                data-testid="account-hero"
              >
                <div class="flex flex-col gap-1">
                  <a
                    class="w-fit [font-family:var(--font-display)] text-[15px] font-semibold"
                    style:color={classColorVar(hero.class)}
                    href={characterHref(hero.region, hero.ruleset, hero.name)}
                  >
                    {hero.name}
                  </a>
                  <span class="text-muted text-[13px]">{characterDescriptor(hero)}</span>
                  {#if heroPath !== null}
                    <CharacterHandoffLinks path={heroPath} />
                  {/if}
                </div>
                <img
                  class="max-h-[280px] w-auto object-contain lg:max-h-[360px]"
                  src={hero.render_url}
                  alt=""
                  loading="lazy"
                  data-testid="account-hero-render"
                />
              </div>
            {/if}

            <CharacterList characters={me!.characters} bnetImportedAt={me!.bnet_imported_at} />

            <StatePanel label="Your reports" testid="account-reports">
              <MyReports {signedIn} />
            </StatePanel>
          </div>

          <div class="flex flex-col gap-8 lg:col-span-4">
            <StatePanel
              label={accountPageCopy.devicesLabel}
              updated={devicesUpdated}
              testid="account-devices"
            >
              {#if devices.length === 0}
                <p class="text-muted text-[14px]">{accountPageCopy.noDevices}</p>
              {:else}
                <ul class="flex flex-col">
                  {#each devices as device (device.id)}
                    <li
                      class="border-line-soft flex min-h-11 items-center justify-between gap-4 border-b py-2"
                    >
                      <span class="text-[14px]">
                        {device.name}
                        <span class="text-muted">· {device.platform}</span>
                      </span>
                      <button
                        class="{SECONDARY_BUTTON_FIXED} border-line-warm text-text px-3"
                        onclick={() => onRevoke(device.id)}
                        disabled={busy}
                      >
                        Revoke
                      </button>
                    </li>
                  {/each}
                </ul>
              {/if}
              {#if pairing === null}
                <button
                  class="{SECONDARY_BUTTON_FIXED} border-line-warm-strong text-strong w-fit px-4"
                  onclick={onPair}
                  disabled={busy}
                >
                  {accountPageCopy.pairADevice}
                </button>
              {:else}
                <p class="tabular text-strong font-mono text-[24px]" data-testid="pairing-code">
                  {pairing.code}
                </p>
                <p class="text-muted text-[13px]">
                  Type this into the companion within {Math.round(pairing.expires_in / 60)} minutes.
                </p>
              {/if}
            </StatePanel>

            <StatePanel label={accountPageCopy.youLabel} testid="account-you">
              <label class="flex min-h-11 items-center gap-3 text-[14px]">
                <input
                  type="checkbox"
                  checked={me!.user.anonymize}
                  onchange={onAnonymize}
                  disabled={busy}
                  data-testid="anonymize"
                />
                {accountPageCopy.pseudonymLabel}
              </label>
              <p class="text-muted text-[13px]">{accountPageCopy.pseudonymNote}</p>
            </StatePanel>

            <StatePanel label={accountPageCopy.guildsAndPlanLabel} testid="account-guilds-plan">
              {#if me!.guilds.length > 0}
                <ul class="flex flex-col" data-testid="account-guilds">
                  {#each me!.guilds as guild (guild.id)}
                    <li
                      class="border-line-soft flex min-h-11 flex-wrap items-center gap-3 border-b py-2 text-[14px]"
                    >
                      <a href={guildHref(guild.region, guild.ruleset, guild.name)}>{guild.name}</a>
                      <select
                        class="border-line-warm bg-raised rounded-control text-text h-11 px-3 text-[13px] md:h-9"
                        value={guild.consent ?? 'gear'}
                        onchange={(event) => onConsentChange(guild.id, event)}
                        disabled={busy || guildBusy === guild.id}
                        data-testid="account-guild-consent"
                      >
                        <option value="roster">{guildConsentCopy.roster}</option>
                        <option value="gear">{guildConsentCopy.gear}</option>
                        <option value="gear_bags">{guildConsentCopy.gearBags}</option>
                      </select>
                      <button
                        class="{SECONDARY_BUTTON_FIXED} border-line-warm text-text px-3"
                        onclick={() => onLeaveGuild(guild.id)}
                        disabled={busy || guildBusy === guild.id}
                        data-testid="account-guild-leave"
                      >
                        {guildConsentCopy.leave}
                      </button>
                    </li>
                  {/each}
                </ul>
              {/if}

              <div
                class={me!.guilds.length > 0
                  ? 'border-line-soft flex flex-col gap-3 border-t pt-4'
                  : 'flex flex-col gap-3'}
              >
                {#if billing === null}
                  <p class="text-[14px]" data-testid="account-billing-row">
                    <span class="text-muted">{accountPageCopy.planKey}</span> · {billingBlockCopy.notSubscribed}
                    <a class="text-text underline" href="/premium">{billingBlockCopy.seePlans}</a>
                  </p>
                {:else}
                  <p class="text-[14px]">
                    {billing.plan} —
                    {billing.cancel_at_period_end ? billingBlockCopy.ends : billingBlockCopy.renews}
                    {billing.current_period_end
                      ? new Date(billing.current_period_end).toLocaleDateString()
                      : ''}
                  </p>
                  {#if billing.status === 'past_due'}
                    <p class="text-strong text-[13px]" role="alert">{billingBlockCopy.pastDueBanner}</p>
                  {/if}
                  <button
                    class="{SECONDARY_BUTTON_FIXED} border-line-warm text-text w-fit px-4"
                    onclick={onManageBilling}
                    disabled={busy}
                  >
                    {billingBlockCopy.manageBilling}
                  </button>
                {/if}
              </div>
            </StatePanel>
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
