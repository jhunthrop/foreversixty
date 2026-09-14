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
    ACCOUNT_FAILED, EMAIL_SENT, battlenetStartUrl, fetchMe, listDevices, pairDevice,
    revokeDevice, requestEmailLink, setAnonymize, signOut, type Device, type Me, type PairingCode,
  } from '../lib/account/api';
  import { characterHref, rulesetLabel } from '../lib/characters';
  import { SECONDARY_BUTTON_FIXED } from '../lib/planner/styles';
  import MyReports from './MyReports.svelte';

  let { mode, next = '/logs' }: { mode: 'nav' | 'login' | 'account' | 'pairing' | 'reports'; next?: string } = $props();

  let me = $state<Me | null>(null);
  let status = $state<'loading' | 'ready' | 'failed'>('loading');
  let error = $state('');
  let notice = $state('');
  let email = $state('');
  let devices = $state<Device[]>([]);
  let pairing = $state<PairingCode | null>(null);
  let busy = $state(false);

  const signedIn = $derived(me !== null);
  const displayName = $derived(me?.user.battletag ?? me?.user.email ?? 'Your account');

  async function load(): Promise<void> {
    status = 'loading';
    error = '';
    try {
      me = await fetchMe();
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

  const onAnonymize = (event: Event): void => {
    const wanted = (event.currentTarget as HTMLInputElement).checked;
    void run(async () => {
      await setAnonymize(wanted);
      if (me !== null) me = { ...me, user: { ...me.user, anonymize: wanted } };
    });
  };
</script>

{#if mode === 'nav'}
  <div class="flex items-center gap-3 text-[13px]" data-testid="session-nav">
    {#if status === 'loading'}
      <span class="text-muted">&nbsp;</span>
    {:else if signedIn}
      <a href="/account" class="text-nav hover:text-strong inline-flex min-h-11 items-center md:min-h-0">
        {displayName}
      </a>
    {:else}
      <a href="/login" class="text-nav hover:text-strong inline-flex min-h-11 items-center md:min-h-0">
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
          href={battlenetStartUrl(next)}
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
            required
            bind:value={email}
          />
          <button class="{SECONDARY_BUTTON_FIXED} border-line-warm-strong text-strong px-4" type="submit" disabled={busy}>
            Send link
          </button>
        </div>
      </form>
    {/if}
    {#if notice !== ''}<p class="text-[14px]" data-testid="account-notice">{notice}</p>{/if}
    {#if error !== ''}<p class="text-[14px]" role="alert" data-testid="account-error">{error}</p>{/if}
  </div>
{:else if mode === 'reports'}
  <MyReports signedIn={signedIn} />
{:else if mode === 'pairing'}
  <div class="flex flex-col gap-3" data-testid="pairing">
    {#if !signedIn}
      <p class="text-[14px]"><a href="/login">Sign in</a> to pair the companion.</p>
    {:else if pairing === null}
      <button class="{SECONDARY_BUTTON_FIXED} border-line-warm-strong text-strong px-4" onclick={onPair} disabled={busy}>
        Show pairing code
      </button>
    {:else}
      <p class="font-mono tabular text-strong text-[24px]" data-testid="pairing-code">{pairing.code}</p>
      <p class="text-muted text-[13px]">
        Type this into the companion within {Math.round(pairing.expires_in / 60)} minutes. It pairs one
        device and cannot be reused.
      </p>
    {/if}
    {#if error !== ''}<p class="text-[14px]" role="alert" data-testid="account-error">{error}</p>{/if}
  </div>
{:else}
  <div class="flex flex-col gap-8" data-testid="account">
    {#if status === 'loading'}
      <p class="text-muted text-[14px]">Loading your account.</p>
    {:else if !signedIn}
      <p class="text-[14px]"><a href="/login">Sign in</a> to see your devices and characters.</p>
    {:else}
      <section class="flex flex-col gap-3">
        <h2 class="section-title text-[18px]">Account</h2>
        <p class="text-[14px]">{displayName}</p>
        <button class="{SECONDARY_BUTTON_FIXED} border-line-warm text-text w-fit px-4" onclick={onSignOut} disabled={busy}>
          Sign out
        </button>
      </section>

      <section class="flex flex-col gap-3">
        <h2 class="section-title text-[18px]">Devices</h2>
        {#if devices.length === 0}
          <p class="text-muted text-[14px]">No devices paired.</p>
        {:else}
          <ul class="flex flex-col">
            {#each devices as device (device.id)}
              <li class="border-line-soft flex min-h-11 items-center justify-between gap-4 border-b py-2">
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
          <button class="{SECONDARY_BUTTON_FIXED} border-line-warm-strong text-strong w-fit px-4" onclick={onPair} disabled={busy}>
            Pair a device
          </button>
        {:else}
          <p class="font-mono tabular text-strong text-[24px]" data-testid="pairing-code">{pairing.code}</p>
          <p class="text-muted text-[13px]">
            Type this into the companion within {Math.round(pairing.expires_in / 60)} minutes.
          </p>
        {/if}
      </section>

      <section class="flex flex-col gap-3">
        <h2 class="section-title text-[18px]">Characters</h2>
        {#if me!.characters.length === 0}
          <p class="text-muted text-[14px]">No characters linked yet. Sign in with Battle.net to link them.</p>
        {:else}
          <ul class="flex flex-col">
            {#each me!.characters as character (character.key)}
              <li class="border-line-soft flex min-h-11 items-center gap-3 border-b py-2 text-[14px]">
                <a href={characterHref(character.region, character.ruleset, character.name)}>{character.name}</a>
                <span class="text-muted">
                  {rulesetLabel(character.ruleset)} {character.region.toUpperCase()}
                </span>
              </li>
            {/each}
          </ul>
        {/if}
      </section>

      <section class="flex flex-col gap-3">
        <h2 class="section-title text-[18px]">Name</h2>
        <label class="flex min-h-11 items-center gap-3 text-[14px]">
          <input type="checkbox" checked={me!.user.anonymize} onchange={onAnonymize} disabled={busy} data-testid="anonymize" />
          Show a pseudonym instead of my character names
        </label>
        <p class="text-muted text-[13px]">
          Applies everywhere your characters appear, on reports and rankings alike. Reports themselves are
          never deleted or rewritten.
        </p>
      </section>

      <MyReports signedIn={signedIn} />
    {/if}
    {#if error !== ''}<p class="text-[14px]" role="alert" data-testid="account-error">{error}</p>{/if}
  </div>
{/if}
