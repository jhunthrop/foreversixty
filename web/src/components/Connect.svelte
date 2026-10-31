<!-- web/src/components/Connect.svelte -->
<!-- /connect: where the companion's "Sign in with Battle.net" lands. One centred card whose
     device panel is always 88px, so signing in never moves anything. A bare link cannot
     attach a stranger's companion: the code and the device name are shown, and a signed-in
     click on Connect is the only way the approve call is ever made. -->
<script lang="ts">
  import { onMount } from 'svelte';
  import {
    EMAIL_SENT,
    battlenetStartUrl,
    connectDeviceApprove,
    connectDeviceDeny,
    connectDeviceLookup,
    fetchMeOnce,
    forgetSignedInState,
    requestEmailLink,
    signOut,
    type ConnectDevice,
    type Me,
  } from '../lib/account/api';
  import { clearAccountReserve } from '../lib/account/layout';
  import { connectCopy } from '../lib/account/connect-copy';
  import { displayCode, normalizeCode, stateForError, type ConnectState } from '../lib/account/connect-flow';
  import { PRIMARY_BUTTON_FIXED, SECONDARY_BUTTON_FIXED } from '../lib/planner/styles';

  let state = $state<ConnectState>('loading');
  let code = $state('');
  let me = $state<Me | null>(null);
  let device = $state<ConnectDevice | null>(null);
  let busy = $state(false);
  let showEmail = $state(false);
  let email = $state('');
  let emailNote = $state('');
  // What "Try again" repeats: the load, or the approve that could not reach the site.
  let retry: () => void = () => {};

  const shownCode = $derived(displayCode(code));
  const who = $derived(me?.user.battletag ?? me?.user.email ?? 'your account');
  const nextPath = $derived(`/connect?code=${shownCode}`);

  async function load(): Promise<void> {
    state = 'loading';
    retry = () => void load();
    if (code === '') {
      state = 'expired';
      return;
    }
    try {
      me = await fetchMeOnce();
      if (me === null) {
        state = 'signedOut';
        return;
      }
      device = await connectDeviceLookup(shownCode);
      state = 'pending';
    } catch (error) {
      state = stateForError(error);
    }
  }

  async function onConnect(): Promise<void> {
    busy = true;
    retry = () => void onConnect();
    try {
      device = await connectDeviceApprove(shownCode);
      state = 'success';
    } catch (error) {
      state = stateForError(error);
    } finally {
      busy = false;
    }
  }

  async function onCancel(): Promise<void> {
    busy = true;
    retry = () => void onCancel();
    try {
      await connectDeviceDeny(shownCode);
      state = 'cancelled';
    } catch (error) {
      state = stateForError(error);
    } finally {
      busy = false;
    }
  }

  async function onSignOut(): Promise<void> {
    try {
      await signOut();
    } catch {
      // Signing out of a session that is already gone is the same outcome.
    }
    forgetSignedInState();
    clearAccountReserve();
    me = null;
    state = 'signedOut';
  }

  async function onEmail(event: SubmitEvent): Promise<void> {
    event.preventDefault();
    busy = true;
    try {
      await requestEmailLink(email.trim(), undefined, nextPath);
      emailNote = EMAIL_SENT;
    } catch (error) {
      emailNote = error instanceof Error ? error.message : '';
    } finally {
      busy = false;
    }
  }

  const onClose = (): void => {
    // A tab the companion opened can close itself; any other falls through to the link's href.
    window.close();
  };

  onMount(() => {
    code = normalizeCode(new URLSearchParams(window.location.search).get('code'));
    // The buttons are Barlow 700, the one face the layout does not preload, and nothing on this
    // page asks for it while it says "Checking". Without this the first state with a button
    // (expired, above all, which needs no request) paints it in the fallback face. Ask now and
    // wait for it, but never longer than a moment.
    const bold = document.fonts?.load('700 12px Barlow').catch(() => []) ?? Promise.resolve([]);
    void Promise.race([bold, new Promise((resolve) => setTimeout(resolve, 1500))]).then(load);
  });
</script>

<div class="flex flex-col gap-4" data-testid="connect" data-state={state}>
  {#if state === 'loading'}
    <h1 class="heading-page">{connectCopy.title}</h1>
    <p class="text-muted text-[14px]">{connectCopy.checking}</p>
    <div class="h-[88px]" aria-hidden="true"></div>
  {:else if state === 'signedOut'}
    <h1 class="heading-page">{connectCopy.title}</h1>
    <p class="text-[14px]">{connectCopy.signedOutLead}</p>
    <div
      class="border-line-warm bg-bg rounded-control flex h-[88px] flex-col justify-center gap-1 border px-4"
      data-testid="connect-panel"
    >
      <p class="tabular text-strong font-mono text-[20px]" data-testid="connect-code">
        {connectCopy.codeLabel(shownCode)}
      </p>
      <p class="text-muted text-[13px]">{connectCopy.codeCheck}</p>
    </div>
    <p class="text-muted text-[13px]">{connectCopy.consent}</p>
    <div class="flex flex-col gap-3">
      <a
        class="{PRIMARY_BUTTON_FIXED} justify-center px-5"
        href={battlenetStartUrl(nextPath)}
        data-testid="connect-battlenet"
      >
        {connectCopy.battlenet}
      </a>
      {#if !showEmail}
        <button
          class="{SECONDARY_BUTTON_FIXED} border-line-warm-strong text-strong justify-center px-4"
          onclick={() => (showEmail = true)}
          data-testid="connect-email-toggle"
        >
          {connectCopy.emailToggle}
        </button>
      {:else}
        <form class="flex flex-col gap-3" onsubmit={onEmail} data-testid="connect-email-form">
          <label class="sr-only" for="connect-email">{connectCopy.emailLabel}</label>
          <div class="flex gap-3">
            <input
              id="connect-email"
              class="border-line-warm bg-raised rounded-control text-text h-11 min-w-0 flex-1 px-3 text-[15px]"
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
              {connectCopy.emailSend}
            </button>
          </div>
          {#if emailNote !== ''}
            <p class="text-[13px]" role="status" data-testid="connect-email-note">{emailNote}</p>
          {/if}
        </form>
      {/if}
    </div>
  {:else if state === 'pending' && device !== null}
    <h1 class="heading-page">{connectCopy.pendingTitle(device.name, who)}</h1>
    <p class="text-[14px]">{connectCopy.pendingBody}</p>
    <dl
      class="border-line-warm bg-bg rounded-control grid h-[88px] grid-cols-[auto_1fr] content-center gap-x-4 gap-y-0.5 border px-4 text-[14px]"
      data-testid="connect-panel"
    >
      <dt class="text-muted">{connectCopy.deviceLabel}</dt>
      <dd class="text-strong truncate" data-testid="connect-device">{device.name}</dd>
      <dt class="text-muted">{connectCopy.platformLabel}</dt>
      <dd class="text-strong truncate">{device.platform}</dd>
      <dt class="text-muted">{connectCopy.codeRowLabel}</dt>
      <dd class="tabular text-strong font-mono" data-testid="connect-code">{shownCode}</dd>
    </dl>
    <p class="text-muted text-[13px]" data-testid="connect-code-check">{connectCopy.pendingCodeCheck}</p>
    <p class="text-muted text-[13px]">{connectCopy.consent}</p>
    <div class="flex gap-3">
      <button
        class="{PRIMARY_BUTTON_FIXED} flex-1 justify-center px-5 disabled:opacity-60"
        onclick={onConnect}
        disabled={busy}
        data-testid="connect-approve"
      >
        {busy ? connectCopy.connecting : connectCopy.connect}
      </button>
      <button
        class="{SECONDARY_BUTTON_FIXED} border-line-warm-strong text-strong px-4"
        onclick={onCancel}
        disabled={busy}
        data-testid="connect-deny"
      >
        {connectCopy.cancel}
      </button>
    </div>
    <p class="text-muted text-[13px]">
      {connectCopy.notYou}
      <button class="text-nav underline" onclick={onSignOut} data-testid="connect-signout">
        {connectCopy.signOut}.
      </button>
    </p>
  {:else if state === 'success'}
    <div class="flex items-center gap-3">
      <svg
        class="text-kill h-6 w-6 shrink-0"
        viewBox="0 0 24 24"
        fill="none"
        stroke="currentColor"
        stroke-width="2.5"
        aria-hidden="true"
        data-testid="connect-check"
      >
        <path d="M4 12.5l5 5 11-11" stroke-linecap="round" stroke-linejoin="round" />
      </svg>
      <h1 class="heading-page">{connectCopy.successTitle}</h1>
    </div>
    <p class="text-[14px]">{connectCopy.successBody}</p>
    <div class="flex flex-wrap gap-3">
      <a class="{PRIMARY_BUTTON_FIXED} justify-center px-5" href="/account" data-testid="connect-account">
        {connectCopy.yourAccount}
      </a>
      <a
        class="{SECONDARY_BUTTON_FIXED} border-line-warm-strong text-strong justify-center px-4"
        href="/setup"
        data-testid="connect-setup"
      >
        {connectCopy.getSetUp}
      </a>
    </div>
  {:else if state === 'cancelled'}
    <h1 class="heading-page">{connectCopy.cancelledTitle}</h1>
    <p class="text-[14px]">{connectCopy.cancelledBody}</p>
    <a
      class="{SECONDARY_BUTTON_FIXED} border-line-warm-strong text-strong justify-center self-start px-4"
      href="/"
      onclick={onClose}
      data-testid="connect-close"
    >
      {connectCopy.close}
    </a>
  {:else if state === 'expired'}
    <h1 class="heading-page">{connectCopy.expiredTitle}</h1>
    <p class="text-[14px]">{connectCopy.expiredBody}</p>
    <a
      class="{SECONDARY_BUTTON_FIXED} border-line-warm-strong text-strong justify-center self-start px-4"
      href="/"
      data-testid="connect-home"
    >
      {connectCopy.backHome}
    </a>
  {:else if state === 'used'}
    <h1 class="heading-page">{connectCopy.usedTitle}</h1>
    <p class="text-[14px]">{connectCopy.usedBody}</p>
    <a
      class="{PRIMARY_BUTTON_FIXED} justify-center self-start px-5"
      href="/logs"
      data-testid="connect-reports"
    >
      {connectCopy.openReports}
    </a>
  {:else}
    <h1 class="heading-page">{connectCopy.failedTitle}</h1>
    <p class="text-[14px]" role="alert">{connectCopy.failedBody}</p>
    <button
      class="{PRIMARY_BUTTON_FIXED} justify-center self-start px-5"
      onclick={() => retry()}
      data-testid="connect-retry"
    >
      {connectCopy.tryAgain}
    </button>
  {/if}
</div>
