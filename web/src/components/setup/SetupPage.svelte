<!-- web/src/components/setup/SetupPage.svelte -->
<!-- /setup (setup spec 2026-11-07 sections 4 and 5; directions mock b4-setup) as one island, because
     the head's download, the framed plate's three steps, The addon's and The companion's quiet links
     and the pairing block all follow one account read: the PageHeader (its answer the companion's
     download), the claim's way back, then a split Page of panels whose framed Plate is "Your setup"
     (lib/setup/steps.ts's rows), then the panels The companion, The addon and What it does in the
     game, in the plate's order; the rail's panels hold Try an export, the commands, the handshake and the two flows.

     The server markup is the signed-out page, with both companion builds offered: the island
     reads the visitor's system on mount and the Download step becomes their own build (or, on a
     phone, the link to copy). Until the island mounts, the plate draws the rows the
     pre-paint hints on <html> point at (data-session: the loading rows; data-pointer: the addon
     ticked), so nothing flips after paint (W.10). -->
<script lang="ts">
  import { onMount } from 'svelte';
  import { invalidateMe, battlenetStartUrl, type Device, type Me } from '../../lib/account/api';
  import { characterListCopy } from '../../lib/account/character-list-copy';
  import {
    compareSnapshot,
    saveSnapshot,
    snapshotOf,
    takeSnapshot,
  } from '../../lib/account/refresh-snapshot';
  import { addonCopy } from '../../lib/addon/copy';
  import { readCurrent } from '../../lib/current-character';
  import { onCurrentCharacterChange } from '../../lib/current-character-events';
  import { sessionHinted } from '../../lib/data/query';
  import { fetchAddonVersion, fetchCompanionRelease, type CompanionRelease } from '../../lib/setup/release';
  import { SETUP_REFRESHED } from '../../lib/setup/links';
  import { loadSetupState } from '../../lib/setup/state';
  import { setupSteps, type SessionState, type StepAction } from '../../lib/setup/steps';
  import { visitorPlatform, type VisitorPlatform } from '../../lib/setup/visitor-platform';
  import SetupBackToClaim from '../claim/SetupBackToClaim.svelte';
  import Page from '../ui/Page.svelte';
  import PageHeader from '../ui/PageHeader.svelte';
  import Plate from '../ui/Plate.svelte';
  import SetupAddon from './SetupAddon.svelte';
  import CompanionDownload from './CompanionDownload.svelte';
  import SetupCompanion from './SetupCompanion.svelte';
  import SetupRail from './SetupRail.svelte';
  import SetupStepList from './SetupStepList.svelte';
  import SetupSurfaces from './SetupSurfaces.svelte';

  /** While Battle.net's import runs, `/v1/me` is read again this often (setup spec G1). */
  const IMPORT_POLL_MS = 5000;
  /** How long Copy link says "Link copied". */
  const COPIED_MS = 2000;

  let mounted = $state(false);
  let session = $state<SessionState>('out');
  let me = $state<Me | null>(null);
  let devices = $state<readonly Device[] | null>([]);
  let platform = $state<VisitorPlatform>('unknown');
  let release = $state<CompanionRelease | null>(null);
  let addonLatest = $state<string | null>(null);
  let copied = $state(false);
  let pointer = $state(false);
  let signinFailed = $state(false);
  let refreshFailed = $state(false);
  let refreshing = $state(false);
  let pairingCode = $state<string | null>(null);
  let now = $state(new Date());

  /** One read of the account and its companions. A visitor with no session hint who cannot be
   *  read is signed out, never "failed": there is no session to fail. */
  async function read(hinted: boolean): Promise<Me | null> {
    const result = await loadSetupState();
    now = new Date();
    session = result.session === 'error' && !hinted ? 'out' : result.session;
    me = result.me;
    devices = result.devices;
    return result.me;
  }

  /** Drops the URL parameters this page reads once (`signin`, `refreshed`), keeping `next`. */
  function dropOnceParams(params: URLSearchParams): void {
    params.delete('signin');
    params.delete('refreshed');
    const query = params.toString();
    window.history.replaceState(
      {},
      '',
      `${window.location.pathname}${query === '' ? '' : `?${query}`}${window.location.hash}`,
    );
  }

  // Once, on mount: never re-run when a value it writes changes (onMount does not track).
  onMount(() => {
    mounted = true;
    platform = visitorPlatform(navigator);
    void fetchCompanionRelease().then((fresh) => (release = fresh));
    void fetchAddonVersion().then((fresh) => (addonLatest = fresh));
    pointer = readCurrent() !== null;
    const params = new URLSearchParams(window.location.search);
    signinFailed = params.get('signin') === 'failed';
    const refreshed = params.get('refreshed') === '1';
    if (signinFailed || refreshed) dropOnceParams(params);
    const before = refreshed ? takeSnapshot(sessionStorage) : null;
    const hinted = sessionHinted();
    if (hinted) session = 'loading';
    void read(hinted).then((fresh) => {
      if (refreshed && fresh !== null) refreshFailed = compareSnapshot(before, fresh).kind === 'error';
    });
    const unsubscribe = onCurrentCharacterChange(() => {
      pointer = readCurrent() !== null;
    });
    // Back from Battle.net with the browser's Back: the refresh is not starting any more.
    const reset = (event: PageTransitionEvent): void => {
      if (event.persisted) refreshing = false;
    };
    window.addEventListener('pageshow', reset);
    return () => {
      unsubscribe();
      window.removeEventListener('pageshow', reset);
    };
  });

  // The Battle.net import is still running: read the account again until it is not.
  $effect(() => {
    if (me?.bnet_import?.state !== 'running') return;
    const timer = setTimeout(() => {
      invalidateMe();
      void read(true);
    }, IMPORT_POLL_MS);
    return () => clearTimeout(timer);
  });

  /** Copy link: this page's address, to open on the computer the companion runs on. */
  async function copyLink(): Promise<void> {
    try {
      await navigator.clipboard.writeText(`${window.location.origin}${window.location.pathname}`);
    } catch {
      return; // no clipboard (an insecure context, a refused permission): the address bar still has it
    }
    copied = true;
    setTimeout(() => (copied = false), COPIED_MS);
  }

  function onAction(action: StepAction): void {
    if (action.id === 'copylink') {
      void copyLink();
      return;
    }
    if (action.id === 'retry') {
      session = 'loading';
      invalidateMe();
      void read(true);
      return;
    }
    if (action.id === 'refresh' && !refreshing && me !== null) {
      refreshing = true;
      saveSnapshot(snapshotOf(me), sessionStorage);
      window.location.assign(battlenetStartUrl(SETUP_REFRESHED));
    }
  }

  const rows = $derived(
    setupSteps({
      session,
      me,
      devices,
      platform,
      signinFailed,
      refreshFailed,
      pointer,
      pairingCode,
      addonLatest,
      now,
    }),
  );
  /** Before the island mounts, the server draws the signed-out rows as the live list (kept, not
   *  replaced, on mount: the first paint's text stays the page's largest paint) beside the rows the
   *  other pre-paint hints point at, swapped by CSS. */
  const HINT_OUT_ROWS = setupSteps({ session: 'out', me: null, devices: [], platform: 'unknown' });
  const hintRows = [
    {
      key: 'pointer',
      rows: setupSteps({ session: 'out', me: null, devices: [], platform: 'unknown', pointer: true }),
    },
    {
      key: 'in',
      rows: setupSteps({ session: 'loading', me: null, devices: null, platform: 'unknown' }),
    },
  ];
  const settled = $derived(mounted && session !== 'loading');
  const signedIn = $derived(session === 'in' && me !== null);
  const addonDone = $derived(rows.find((row) => row.id === 'addon')?.mark === 'ok');
</script>

<!-- The companion's download at the top, in every state (owner 2026-11-08), never gold (the plate
     keeps the page's one primary): the head's actions alone (ui/PageHeader); the signed-in
     character is the header's and the Sign in step's. -->
{#snippet headDownload(shown: VisitorPlatform)}
  <CompanionDownload
    platform={shown}
    {release}
    primary={false}
    {copied}
    onCopyLink={() => void copyLink()}
    idPrefix="setup-head"
    end
  />
{/snippet}

<!-- Before the island mounts, the server draws both the phone's download and the one for a
     system it cannot name, and the pre-paint data-platform hint (Base.astro) keeps the right one,
     so mounting does not change the head's height (the Lighthouse CLS budget at 360). -->
{#snippet downloadContent()}
  <div class="setup-head-download">
    {#if mounted}
      {@render headDownload(platform)}
    {:else}
      <div class="setup-hint" data-hint="platform-other">{@render headDownload('unknown')}</div>
      <div class="setup-hint" data-hint="platform-mobile">{@render headDownload('mobile')}</div>
    {/if}
  </div>
{/snippet}

<Page layout="split" panels testid="setup-grid">
  {#snippet header()}
    <PageHeader title={addonCopy.setupPageTitle} actions={downloadContent} testid="setup-head"></PageHeader>
    <SetupBackToClaim />
  {/snippet}

  <Plate label="Your setup" frame="framed" testid="setup-plate">
    <div class={mounted ? undefined : 'setup-hint'} data-hint={mounted ? undefined : 'out'}>
      <SetupStepList
        rows={mounted ? rows : HINT_OUT_ROWS}
        {platform}
        {release}
        {refreshing}
        {copied}
        {onAction}
        testid="setup-steps"
      />
    </div>
    {#if !mounted}
      {#each hintRows as hint (hint.key)}
        <div class="setup-hint" data-hint={hint.key}>
          <SetupStepList rows={hint.rows} platform="unknown" />
        </div>
      {/each}
    {/if}
  </Plate>
  <p class="sr-only" role="status">{refreshing ? characterListCopy.refreshingAnnounce : ''}</p>
  <SetupCompanion
    {platform}
    {release}
    {copied}
    onCopyLink={() => void copyLink()}
    {signedIn}
    devices={devices ?? []}
    {now}
    onCode={(code) => (pairingCode = code)}
    onPaired={(device) => {
      pairingCode = null;
      devices = [...(devices ?? []), device];
    }}
  />
  <SetupAddon {addonDone} {platform} />
  <SetupSurfaces />

  {#snippet rail()}
    <SetupRail signedIn={settled ? signedIn : null} />
  {/snippet}
</Page>

<style>
  /* The pre-paint swap: signed out (with or without a stored character) or session hinted, and
     the head's download for a phone or not. */
  :global(html[data-session='1']) .setup-hint:not([data-hint='in']),
  :global(html:not([data-session='1'])) .setup-hint[data-hint='in'],
  :global(html:not([data-session='1'])[data-pointer='1']) .setup-hint[data-hint='out'],
  :global(html:not([data-pointer='1'])) .setup-hint[data-hint='pointer'],
  :global(html[data-platform='mobile']) .setup-hint[data-hint='platform-other'],
  :global(html:not([data-platform='mobile'])) .setup-hint[data-hint='platform-mobile'] {
    display: none;
  }
  /* The Sections' shared body: a column, reading lines, quiet link rows. */
  :global(.setup-body) {
    display: flex;
    flex-direction: column;
    gap: 14px;
  }
  /* Every sentence in a Section body reads at 15 / 24 (the mock's .body p), quiet ones muted. */
  :global(.setup-body p) {
    max-width: 72ch;
    font-size: 15px;
    line-height: 24px;
    text-wrap: pretty;
  }
  :global(.setup-acts) {
    display: flex;
    flex-wrap: wrap;
    align-items: center;
    gap: 12px;
  }
</style>
