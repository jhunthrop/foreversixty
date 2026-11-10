<!-- web/src/components/Connect.svelte -->
<!-- /connect: where the companion's "Sign in with Battle.net" lands (setup spec 2026-11-07
     section 6). Page layout "form": the PageHeader says the state (its H1 is the state's title,
     so a screen reader hears it), then one framed Plate, "The companion": the check line and the
     code at display size (directions mock b4-connect), Device and Platform under it, the consent
     line, the actions. The check line, the code slot and the two Facts rows hold from loading
     through signed out to pending, so signing in never moves the code. A bare link cannot attach a stranger's companion: the code and the device
     are shown, and a signed-in click on Connect is the only way the approve call is ever made. -->
<script lang="ts">
  import { onMount } from 'svelte';
  import {
    battlenetStartUrl,
    connectDeviceApprove,
    connectDeviceDeny,
    connectDeviceLookup,
    fetchMeOnce,
    forgetSignedInState,
    signOut,
    type ConnectDevice,
    type Me,
  } from '../lib/account/api';
  import { clearAccountReserve } from '../lib/account/layout';
  import { connectCopy } from '../lib/account/connect-copy';
  import { displayCode, normalizeCode, stateForError, type ConnectState } from '../lib/account/connect-flow';
  import { platformLabel } from '../lib/account/platform-label';
  import { addonCopy } from '../lib/addon/copy';
  import { button } from '../lib/ui/button';
  import type { Fact } from '../lib/ui/facts';
  import { LINK } from '../lib/ui/primitives';
  import Button from './ui/Button.svelte';
  import Code from './ui/Code.svelte';
  import Facts from './ui/Facts.svelte';
  import Page from './ui/Page.svelte';
  import PageHeader from './ui/PageHeader.svelte';
  import Plate from './ui/Plate.svelte';
  import Tag from './ui/Tag.svelte';

  let state = $state<ConnectState>('loading');
  let code = $state('');
  let me = $state<Me | null>(null);
  let device = $state<ConnectDevice | null>(null);
  let busy = $state(false);
  // What "Try again" repeats: the load, or the approve that could not reach the site.
  let retry: () => void = () => {};

  const shownCode = $derived(displayCode(code));
  const who = $derived(me?.user.battletag ?? me?.user.email ?? 'your account');
  const nextPath = $derived(`/connect?code=${shownCode}`);

  /** The state's title and its one body sentence: the PageHeader's H1 and meta line. */
  const HEAD: Readonly<Record<Exclude<ConnectState, 'pending'>, readonly [string, string]>> = {
    loading: [connectCopy.title, connectCopy.checking],
    signedOut: [connectCopy.title, connectCopy.signedOutLead],
    success: [connectCopy.successTitle, connectCopy.successBody],
    cancelled: [connectCopy.cancelledTitle, connectCopy.cancelledBody],
    expired: [connectCopy.expiredTitle, connectCopy.expiredBody],
    used: [connectCopy.usedTitle, connectCopy.usedBody],
    failed: [connectCopy.failedTitle, connectCopy.failedBody],
  };

  const head = $derived.by((): readonly [string, string] => {
    if (state !== 'pending') return HEAD[state];
    const title = device === null ? connectCopy.title : connectCopy.pendingTitle(device.name, who);
    return [title, connectCopy.pendingBody];
  });

  /** The line over the code, while there is a code to check: the code itself never moves (rulings MUST 3). */
  const codeCheck = $derived.by((): string => {
    if (state === 'signedOut') return connectCopy.codeCheck;
    if (state === 'pending' || state === 'loading') return connectCopy.pendingCodeCheck;
    return '';
  });

  /** Device and Platform under the code, once a device is known (or what signing in will show). */
  const facts = $derived.by((): Fact[] => {
    if (state === 'loading')
      return [connectCopy.deviceLabel, connectCopy.platformLabel].map((label) => ({
        id: label,
        label,
        content: skeletonValue,
      }));
    if (state === 'expired') return [];
    if (state === 'signedOut')
      return [
        { id: 'device', label: connectCopy.deviceLabel, content: afterSignIn },
        { id: 'platform', label: connectCopy.platformLabel, content: afterSignIn },
      ];
    if (device === null) return [];
    return [
      { id: 'device', label: connectCopy.deviceLabel, value: device.name, testid: 'connect-device' },
      { id: 'platform', label: connectCopy.platformLabel, value: platformLabel(device.platform) },
    ];
  });

  /** The consent line under the Facts (Small). */
  const consent = $derived(state === 'signedOut' || state === 'pending');

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

  async function onSignOut(event: MouseEvent): Promise<void> {
    // The link's href reloads this page, signed out, where script never ran; here it signs out in place.
    event.preventDefault();
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

  const onClose = (): void => {
    // A tab the companion opened can close itself; any other falls through to the link's href.
    window.close();
  };

  onMount(() => {
    code = normalizeCode(new URLSearchParams(window.location.search).get('code'));
    // The buttons are Open Sans 600 (preloaded), and nothing on this page asks for it while it
    // says "Checking". Without this the first state with a button (expired, above all, which
    // needs no request) could paint it in the fallback face. Ask now and wait for it, but never
    // longer than a moment.
    const bold = document.fonts?.load('600 15px "Open Sans"').catch(() => []) ?? Promise.resolve([]);
    void Promise.race([bold, new Promise((resolve) => setTimeout(resolve, 1500))]).then(load);
  });
</script>

{#snippet skeletonValue()}<span class="skeleton-block my-[3px] ml-auto h-[14px] w-[120px]" aria-hidden="true"
  ></span>{/snippet}

{#snippet afterSignIn()}<span class="text-small text-muted font-normal">{connectCopy.afterSignIn}</span
  >{/snippet}

{#snippet okMark()}<span
    class="bg-kill text-bg inline-flex h-10 w-10 flex-none items-center justify-center rounded-full text-[20px] leading-10 font-bold"
    role="img"
    aria-label={connectCopy.connectedTag}
    data-testid="connect-check">&#10003;</span
  >{/snippet}

{#snippet headMeta()}<span role={state === 'failed' ? 'alert' : undefined} data-testid="connect-body"
    >{head[1]}</span
  >{/snippet}

{#snippet header()}<PageHeader
    title={head[0]}
    crest={state === 'success' ? okMark : undefined}
    meta={headMeta}
  />{/snippet}

<div data-testid="connect" data-state={state}>
  <Page layout="form" {header}>
    <Plate label={addonCopy.setupCompanionTitle} frame="framed">
      <div class="connect-body" aria-busy={state === 'loading' ? 'true' : undefined}>
        <div class="flex flex-col gap-2">
          {#if codeCheck !== ''}<p class="text-[15px] leading-6" data-testid="connect-code-check">
              {codeCheck}
            </p>{/if}
          <div class="flex flex-wrap items-center gap-3">
            {#if state === 'loading'}
              <span class="skeleton-block h-[30px] w-[150px]" aria-hidden="true"></span>
            {:else}
              <span class="sr-only">{connectCopy.codeRowLabel}</span>
              <Code text={shownCode} big testid="connect-code" />
            {/if}
            {#if state === 'success'}<Tag tone="success" testid="connect-connected"
                >{connectCopy.connectedTag}</Tag
              >{/if}
          </div>
        </div>
        {#if facts.length > 0}<Facts items={facts} testid="connect-facts" />{/if}
        {#if consent}<p class="text-muted">{connectCopy.consent}</p>{/if}
        <div class="flex flex-wrap items-center gap-3">
          {#if state === 'loading'}
            <span class="skeleton-block h-11 w-[200px]" aria-hidden="true"></span>
            <span class="skeleton-block h-11 w-[96px]" aria-hidden="true"></span>
          {:else if state === 'signedOut'}
            <Button tone="primary" href={battlenetStartUrl(nextPath)} testid="connect-battlenet">
              {connectCopy.battlenet}
            </Button>
          {:else if state === 'pending'}
            <Button tone="primary" disabled={busy} onclick={onConnect} testid="connect-approve">
              {busy ? connectCopy.connecting : connectCopy.connect}
            </Button>
            <Button disabled={busy} onclick={onCancel} testid="connect-deny">{connectCopy.cancel}</Button>
          {:else if state === 'success'}
            <Button tone="primary" href="/setup" testid="connect-setup">{connectCopy.getSetUp}</Button>
            <Button href="/account" testid="connect-account">{connectCopy.yourAccount}</Button>
          {:else if state === 'cancelled'}
            <a class={button('secondary')} href="/" onclick={onClose} data-testid="connect-close">
              {connectCopy.close}
            </a>
          {:else if state === 'expired'}
            <Button href="/" testid="connect-home">{connectCopy.backHome}</Button>
          {:else if state === 'used'}
            <Button tone="primary" href="/logs" testid="connect-reports">{connectCopy.openReports}</Button>
          {:else}
            <Button tone="primary" onclick={() => retry()} testid="connect-retry"
              >{connectCopy.tryAgain}</Button
            >
          {/if}
        </div>
        {#if state === 'pending'}
          <p class="text-muted">
            {connectCopy.notYou}
            <a class={LINK} href={nextPath} onclick={onSignOut} data-testid="connect-signout"
              >{connectCopy.signOut}.</a
            >
          </p>
        {/if}
      </div>
    </Plate>
  </Page>
</div>

<style>
  .connect-body {
    display: flex;
    flex-direction: column;
    gap: 14px;
    padding: 18px 20px 20px;
  }
  .connect-body > p {
    font-size: 15px;
    line-height: 24px;
    max-width: 72ch;
  }
</style>
