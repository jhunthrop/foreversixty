<!-- web/src/components/sim/SimView.svelte -->
<!-- The simulator island's root (spec 2026-11-06-sim). It owns the URL state and the page frame:
     /sim with no character (SimLandingPage), /sim with a character (the head with its answer, the
     run line, then the gear and talents before a run or the report after one), /sim/<id>
     (SavedSim) and /sim/specs (SpecGrid). Every section below is a component that takes data and
     raises events, so this file stays a composition.

     The store reads its own URL bootstrap (a "Sim this build" link's `?code=`, or a saved build's
     `?source=build&ref=`) from init, not from a `$effect` here: `createSimStore` stays DOM-free
     and unit-testable, and this component's only job is to read `window.location` once. -->
<script lang="ts">
  import { onMount, tick, untrack, type Component, type ComponentProps } from 'svelte';
  import type SpecGrid from './SpecGrid.svelte';
  import activeBuild from '../../data/active-build.json';
  import { battlenetStartUrl, effectiveServerSims, fetchMeOnce, type Me } from '../../lib/account/api';
  import { heroCharacter } from '../../lib/account/hero-character';
  import { mainCharacter } from '../../lib/account/main-character';
  import { sessionHinted } from '../../lib/data/query';
  import { createQueryState } from '../../lib/data/query.svelte';
  import { API_BASE_URL } from '../../lib/planner/config';
  import type { CharacterPath } from '../../lib/characters';
  import { parseCharacterPath } from '../../lib/characters';
  import { onCurrentCharacterChange, readCurrent, type CurrentCharacter } from '../../lib/current-character';
  import { createLazyComponent } from '../../lib/report/lazy-component.svelte';
  import { fetchReportMeta, fetchSummary } from '../../lib/report/load';
  import type { Summary } from '../../lib/report/types';
  import { simAnswer, type StoppedSnapshot } from '../../lib/sim/answer';
  import { fetchSpecs, listMySims } from '../../lib/sim/api';
  import type { BisSetClass } from '../../lib/sim/bis-set-rail';
  import type { BisBenchmarks } from '../../lib/sim/bis-benchmark';
  import type { PrerenderedNote } from '../../lib/sim/prerender-notes';
  import { codeForCharacterSpec } from '../../lib/sim/character';
  import { followCurrentCharacter, simHoldsWork } from '../../lib/sim/follow-current-character';
  import { decideBootstrap, RESTORE_BUSY_KEY, runBootstrapRestore } from '../../lib/sim/character-bootstrap';
  import { compareSummaries } from '../../lib/sim/compare';
  import { simCopy } from '../../lib/sim/copy';
  import { characterMeta, characterTitle } from '../../lib/sim/head-meta';
  import { averageItemLevel } from '../../lib/gear/paperdoll';
  import { slotsFromWorn } from '../../lib/gear/paperdoll-slots';
  import { landingCopy } from '../../lib/sim/landing-copy';
  import { createServerRunsController } from '../../lib/sim/server-runs-controller.svelte';
  import { browserNotifier, enableNotifications, notifyFinished } from '../../lib/sim/notify';
  import { askToFocusRun, takeFocusRun } from '../../lib/sim/focus-run';
  import { simPageCopy as pageCopy } from '../../lib/sim/page-copy';
  import { createSavedSimState } from '../../lib/sim/saved-sim-state.svelte';
  import { SIM_SAVED_SKELETON_HTML } from '../../lib/sim/skeleton';
  import { parseBuildInput } from '../../lib/sim/build-input';
  import { fightRefOf, parseFightRef } from '../../lib/sim/sources';
  import { isSimulatedSpec, specLabel } from '../../lib/sim/spec-label';
  import { mergeSpecRows } from '../../lib/sim/spec-state';
  import { createSimStore } from '../../lib/sim/store.svelte';
  import { SIM_TABS, tabHref, tabStateFor } from '../../lib/sim/tabs';
  import {
    decodeRequestParam,
    defaultSimState,
    encodeRequestParam,
    parseSimState,
    simSearch,
    withSimState,
  } from '../../lib/sim/url';
  import type { SimListRow, SimRequest, SimResult, SpecFidelity } from '../../lib/sim/types';
  import LoadError from '../ui/LoadError.svelte';
  import Notice from '../ui/Notice.svelte';
  import Skeleton from '../ui/Skeleton.svelte';
  import RunReport from './RunReport.svelte';
  import SavedSim from './SavedSim.svelte';
  import SimAnswer from './SimAnswer.svelte';
  import SimCharacterHead from './SimCharacterHead.svelte';
  import SimFidelityCaption from './SimFidelityCaption.svelte';
  import SimLandingPage from './SimLandingPage.svelte';
  import SimRacePick from './SimRacePick.svelte';
  import SimRunBar from './SimRunBar.svelte';
  import SimSaveActions from './SimSaveActions.svelte';
  import SimToolTabs from './SimToolTabs.svelte';

  import { grouped } from '../../lib/format/number';
  let {
    simId = '',
    inlineResult = null,
    bisSets = [],
    bisBenchmarks = {},
    initialNotes,
    specGrid = null,
  }: {
    simId?: string;
    inlineResult?: SimResult | null;
    bisSets?: readonly BisSetClass[];
    /** The published sets a run is read against (lib/sim/bis-benchmark.ts). */
    bisBenchmarks?: BisBenchmarks;
    /** A prerendered saved run's Worth knowing, read at build time (sim/[id].astro). */
    initialNotes?: readonly PrerenderedNote[];
    /** /sim/specs's grid, fetched by the entry before it replaced the shell (sim-island.ts). */
    specGrid?: Component<ComponentProps<typeof SpecGrid>> | null;
  } = $props();

  // /sim/<sim_id>, or the prerendered fixture page that inlines its result.
  const hasSavedSimId = untrack(() => simId !== '' || inlineResult !== null);

  // The page's one-shot bootstrap values, read once (not bindings kept in sync with the URL).
  const bootstrap = untrack(() => {
    const { source, ref, code, req, mode, run, keep } = parseSimState(window.location.search);
    const mount = document.getElementById('sim');
    const treeVersion = mount?.dataset.treeVersion ?? activeBuild.build;
    // specs.astro stamps `data-sim-view="specs"`; every other page is the ordinary simulator.
    const view: 'sim' | 'specs' = mount?.dataset.simView === 'specs' ? 'specs' : 'sim';
    return { treeVersion, source, ref, code, request: decodeRequestParam(req), mode, view, run, keep };
  });

  // Compare mode loads its character through `enterCompare`, never through the store's own URL
  // bootstrap: a second bootstrap racing it could land after `store.run()` and null the result.
  const store = untrack(() =>
    createSimStore({
      treeVersion: bootstrap.treeVersion,
      request: bootstrap.request ?? undefined,
      source: bootstrap.mode === 'compare' ? undefined : bootstrap.source,
      ref: bootstrap.mode === 'compare' ? undefined : bootstrap.ref,
      code: bootstrap.code,
      keep: bootstrap.keep,
    }),
  );

  let actual = $state<Summary | null>(null);
  let comparing = $state(false);

  /** Compare mode's whole entry path, called once (never from an effect). */
  async function enterCompare(ref: string): Promise<void> {
    const parsed = parseFightRef(ref);
    if (parsed === null) {
      store.setMessage(simCopy.fightRefInvalid);
      return;
    }
    comparing = true;
    await store.loadFight(ref);
    try {
      const meta = await fetchReportMeta(parsed.reportId);
      actual = await fetchSummary(meta.data_base_url, parsed.fightIndex);
    } catch {
      actual = null;
      comparing = false;
      store.setMessage(simCopy.fightNoCombatant);
      return;
    }
    // The one place a sim starts without a press: the player already asked for the comparison.
    // A logged fight records no race, so the strip asks for one first (character.ts
    // PENDING_RACE); the comparison then starts as soon as it is answered.
    if (store.character === null) return;
    if (store.needsRace) compareAwaitsRace = true;
    else await store.run();
  }

  /** Compare mode stopped at the race question: the run it owes starts once a race is picked. */
  let compareAwaitsRace = $state(false);
  $effect(() => {
    if (!compareAwaitsRace || store.needsRace || store.character === null) return;
    compareAwaitsRace = false;
    void store.run();
  });

  if (bootstrap.mode === 'compare') void enterCompare(bootstrap.ref);

  const savedSim = createSavedSimState(untrack(() => inlineResult));
  untrack(() => {
    if (hasSavedSimId && savedSim.result === null && simId !== '') savedSim.load(simId);
  });

  /** "Open in the sim": the saved run's link (lib/sim/rerun-href.ts) carries the setup; the
   *  click only asks /sim to focus Run once it has loaded. */
  function onRerunSaved(): void {
    askToFocusRun();
  }

  // The other half of "Open in the sim": the setup is loaded, so Enter starts the run.
  let pendingRunFocus = typeof window !== 'undefined' && !hasSavedSimId && takeFocusRun();
  $effect(() => {
    if (!pendingRunFocus || store.character === null) return;
    pendingRunFocus = false;
    void tick().then(() => document.querySelector<HTMLElement>('[data-testid="sim-run-button"]')?.focus());
  });

  /** A share URL of an edited request is a full reproduction. Null past the URL budget. */
  function shareUrlFor(request: SimRequest): string | null {
    const encoded = encodeRequestParam(request);
    if (encoded === null) return null;
    return `${window.location.origin}/sim${simSearch(withSimState(defaultSimState(), { req: encoded }))}`;
  }

  const comparison = $derived(
    comparing && actual !== null && store.result !== null && store.character !== null
      ? compareSummaries(store.result.summary, actual, store.character.name, store.actionNames)
      : null,
  );
  const compareViewLazy = createLazyComponent(() => import('./CompareView.svelte'));
  // Chunks a page loads only when it needs them, so the island's boot bundle stays what the
  // landing and a saved run draw first (scripts/check-island-size.mjs): More settings on its first
  // open, the before-a-run body once a character is loaded, the spec grid on /sim/specs.
  const moreSheetLazy = createLazyComponent(() => import('./MoreSettingsSheet.svelte'));
  const preRunLazy = createLazyComponent(() => import('./SimPreRun.svelte'));
  const specGridLazy = createLazyComponent(
    () => import('./SpecGrid.svelte'),
    untrack(() => specGrid),
  );
  // The saved sim's failure panel ships as its own chunk: only a page whose load failed fetches it.
  // Your sims (the rail's tail) is a signed-in reader's alone and fills from its own fetch: its chunk
  // starts loading as soon as a session is known or hinted, beside that fetch.
  const yourSimsLazy = createLazyComponent(() => import('./SimYourSims.svelte'));
  const savedFailedLazy = createLazyComponent(() => import('./SavedSimLoadFailed.svelte'));
  $effect(() => {
    if (comparison !== null) compareViewLazy.load();
  });

  // One `/v1/me` read, shared with every other island through the client cache.
  const session = createQueryState<Me | null>(`${API_BASE_URL}/v1/me`, () => fetchMeOnce(), {
    scope: 'private',
    ttlMs: 10 * 60 * 1000,
  });
  const me = $derived(session.data);
  const sessionPending = $derived(session.status === 'loading' && session.data === null && sessionHinted());
  const signedIn = $derived(me !== null);

  // The header selector's character (pointer, then main): the page loads the same one.
  let currentPointer = $state<CurrentCharacter | null>(null);
  $effect(() => {
    const read = (): void => {
      currentPointer = readCurrent();
    };
    read();
    return onCurrentCharacterChange(read);
  });
  const headerCharacter = $derived.by(() => {
    if (me === null) return null;
    const pointerCharacter = heroCharacter(currentPointer, me.characters);
    const main = mainCharacter(me.characters, me.main_character_key);
    return pointerCharacter ?? (currentPointer === null ? main : null);
  });
  const headerPath = $derived(
    headerCharacter === null ? null : parseCharacterPath(`/character/${headerCharacter.key}`),
  );

  let historyRows = $state<SimListRow[] | null>(null);
  let historyError = $state<string | null>(null);

  async function loadHistory(): Promise<void> {
    historyError = null;
    try {
      historyRows = (await listMySims(1, undefined, 'all')).rows;
    } catch (error) {
      historyRows = null;
      historyError = error instanceof Error ? error.message : simCopy.loadFailed;
    }
  }

  // A payer's server runs (spec 2026-10-24 §3), folded into the rail's Your sims.
  const serverRuns = createServerRunsController();
  const hasServerSims = $derived(effectiveServerSims(me));
  const showsServerRuns = $derived(hasServerSims && !hasSavedSimId && bootstrap.view !== 'specs');
  $effect(() => {
    if (!showsServerRuns) return;
    return serverRuns.start();
  });
  let wasServerRunning = false;
  $effect(() => {
    const running = store.serverRunning;
    if (running && !wasServerRunning) serverRuns.runStarted();
    if (!running && wasServerRunning) serverRuns.runEnded();
    wasServerRunning = running;
  });

  // The finish notification (design 5.4): asked for only from More settings, server runs only.
  const notifier = untrack(() => browserNotifier());
  let notifyWanted = $state(false);
  let notifiedFor = $state('');
  async function toggleNotify(wanted: boolean): Promise<void> {
    notifyWanted = wanted && (await enableNotifications(notifier));
  }
  $effect(() => {
    const finished = store.result;
    if (finished === null || finished.lane !== 'server') return;
    const key = `${finished.duration_ms}-${finished.iterations_run}`;
    if (key === notifiedFor) return;
    notifiedFor = key;
    notifyFinished(notifier, notifyWanted, store.reportTitle, simCopy.notifyBody(grouped(finished.dps.mean)));
  });

  // The session decides the server lane and the history read; guarded against re-running for
  // the same `me` object (the 2026-09-23 account-page loop).
  let lastMe: Me | null = null;
  $effect(() => {
    if (me === lastMe) return;
    lastMe = me;
    store.setPremium(effectiveServerSims(me));
    if (me !== null) void loadHistory();
  });

  onMount(() => {
    if (!hasSavedSimId) void restoreFromPointer();
    // The header selector moves this page along: a character chosen there loads in place.
    const stopFollowing = followCurrentCharacter({
      loaders: store,
      characterLoaded: () => store.character !== null,
      pinned: () => hasSavedSimId || bootstrap.request !== null,
      holdsWork: () => simHoldsWork(store, bootstrap.code),
      onBusy: (busy) => (landingBusyKey = busy ? RESTORE_BUSY_KEY : null),
    });
    return () => {
      stopFollowing();
      store.dispose();
    };
  });

  // The spec support list: /sim/specs renders it, /sim reads the loaded spec's fidelity from it.
  let specRows = $state<SpecFidelity[] | null>(null);
  let specsError = $state<string | null>(null);
  async function loadSpecs(): Promise<void> {
    specsError = null;
    try {
      specRows = await fetchSpecs();
    } catch (error) {
      specRows = null;
      specsError = error instanceof Error ? error.message : simCopy.specsFailed;
    }
  }
  $effect(() => {
    void loadSpecs();
  });
  /** The spec list's row for a spec, or null before the list loads. */
  function specRowFor(spec: string | undefined): SpecFidelity | null {
    if (spec === undefined || specRows === null) return null;
    return mergeSpecRows(specRows).find((row) => row.spec === spec) ?? null;
  }
  const characterSpecRow = $derived(store.character === null ? null : specRowFor(store.character.spec));
  const savedSpecRow = $derived(savedSim.result === null ? null : specRowFor(savedSim.result.request.spec));

  /** An FS1 code for a character with no ref (a paste), so the tool tabs carry it. */
  const fallbackCode = $derived.by(() => {
    if (store.character === null || store.character.source.ref !== '') return null;
    const request = store.buildRequest();
    return request === null ? null : codeForCharacterSpec(request.character, bootstrap.treeVersion);
  });
  const tabSource = $derived(store.character === null ? null : store.character.source);
  const topGearHref = $derived(
    tabHref(
      SIM_TABS.find((tab) => tab.id === 'gear')?.href ?? '/sim/gear',
      tabStateFor(tabSource, fallbackCode, bootstrap.keep),
    ),
  );

  function onSignIn(): void {
    window.location.href = battlenetStartUrl(`${window.location.pathname}${window.location.search}`);
  }

  /** The paste box's Run: load the export, then run it if the sim runs its spec. */
  async function runPastedInput(code: string): Promise<void> {
    await store.loadAddon(code);
    runLoadedIfSimulated();
  }

  /** "From a build": a saved link or id loads the saved build; an unsaved planner link (its
   *  ?code=) or a bare code loads that code. */
  function loadBuildInput(value: string): void {
    const input = parseBuildInput(value);
    if (input === null) return;
    if (input.kind === 'code') void store.loadAddon(input.code);
    else void store.loadBuild(input.id);
  }

  function runLoadedIfSimulated(): void {
    if (store.character !== null && isSimulatedSpec(store.character.spec)) void store.run();
  }

  // The key of the character a load is in flight for, and of the one whose load failed for want
  // of a build (the sim-input 404, `landingCopy.buildMissingFallback`).
  let landingBusyKey = $state<string | null>(null);
  let landingFailedKey = $state<string | null>(null);

  async function pickCharacter(path: CharacterPath): Promise<void> {
    const key = `${path.region}/${path.ruleset}/${path.slug}`;
    landingBusyKey = key;
    landingFailedKey = null;
    await store.loadStored(path);
    landingBusyKey = null;
    landingFailedKey = store.message === landingCopy.buildMissingFallback ? key : null;
  }

  /** The URL's own bootstrap first; the stored pointer only when that settled with nothing. */
  async function restoreFromPointer(): Promise<void> {
    await store.ready;
    const url = {
      code: bootstrap.code,
      source: bootstrap.source,
      ref: bootstrap.ref,
      hasRequest: bootstrap.request !== null,
    };
    const stored = readCurrent();
    if (!decideBootstrap(url, stored).restored) return;
    landingBusyKey = RESTORE_BUSY_KEY;
    try {
      await runBootstrapRestore(store, url, stored, () => store.character !== null);
    } finally {
      landingBusyKey = null;
    }
  }

  // Signed in with characters and nothing loaded: the header's character loads (spec §4.3, the
  // landing is that character, not a list). Once per character key, so a failed load is not
  // retried in a loop; the selector or Try again asks again.
  let headerTried = $state<string | null>(null);
  $effect(() => {
    if (hasSavedSimId || bootstrap.view === 'specs' || comparing) return;
    if (
      store.character !== null ||
      landingBusyKey !== null ||
      headerPath === null ||
      headerCharacter === null
    )
      return;
    if (headerTried === headerCharacter.key) return;
    headerTried = headerCharacter.key;
    const path = headerPath;
    void store.ready.then(() => {
      if (store.character === null) void pickCharacter(path);
    });
  });
  const headerLoadFailed = $derived(
    headerCharacter !== null &&
      headerTried === headerCharacter.key &&
      store.character === null &&
      landingBusyKey === null,
  );

  // `run=1` (the tier list's and BiS page's "View ranking sim"): run once the code's character
  // has loaded, under the damage-spec guard. Once only.
  let pendingAutoRun = untrack(() => bootstrap.run && !hasSavedSimId && bootstrap.mode !== 'compare');
  $effect(() => {
    if (!pendingAutoRun || store.character === null) return;
    pendingAutoRun = false;
    runLoadedIfSimulated();
  });

  // The run's own page state: the lane the last run went to (a failure's label), the partial
  // number kept when the player pressed Stop (§5.1), the More settings sheet.
  let lastLane = $state<'browser' | 'server'>('browser');
  let stopped = $state<StoppedSnapshot | null>(null);
  let moreOpen = $state(false);
  $effect(() => {
    if (moreOpen) moreSheetLazy.load();
  });
  $effect(() => {
    if (signedIn || sessionPending) yourSimsLazy.load();
  });
  $effect(() => {
    if (savedSim.error !== null && savedSim.error !== simCopy.notFound) savedFailedLazy.load();
  });
  $effect(() => {
    if (bootstrap.view === 'specs') specGridLazy.load();
    else if (store.character !== null && store.result === null) preRunLazy.load();
  });
  let savedUrl = $state<string | null>(null);

  function runBrowser(): void {
    stopped = null;
    lastLane = 'browser';
    void store.run();
  }
  function runServer(): void {
    stopped = null;
    lastLane = 'server';
    void store.runOnServer();
  }
  function stop(): void {
    stopped = { estimate: store.estimate, done: store.iterationsDone, total: store.iterationsTotal };
    store.stop();
  }

  const busy = $derived(store.phase === 'running' || store.phase === 'loading-engine' || store.serverRunning);
  const simulated = $derived(store.character !== null && isSimulatedSpec(store.character.spec));
  const answerModel = $derived(
    simAnswer({
      phase: store.phase,
      simulated,
      specName: store.character === null ? '' : specLabel(store.character.spec),
      estimate: store.estimate,
      iterationsDone: store.iterationsDone,
      iterationsTotal: store.iterationsTotal,
      serverRunning: store.serverRunning,
      result: store.result,
      message: store.message === simCopy.stopped ? null : store.message,
      detail: store.detail,
      stopped: store.message === simCopy.stopped ? stopped : null,
      lastLane,
    }),
  );

  function onPrimary(): void {
    const { primary } = answerModel;
    if (primary === 'stop') stop();
    else if (primary === 'run' || primary === 'run-again' || primary === 'try-again') runBrowser();
  }

  const request = $derived(store.character === null ? null : store.buildRequest());
  const headLines = $derived.by(() => {
    const character = store.character;
    if (character === null || request === null) return [];
    const itemLevel = averageItemLevel(
      slotsFromWorn(character.gear, { lookup: (id) => store.items.get(id) }),
    );
    return characterMeta(
      {
        race: character.race_slug,
        spec: character.spec,
        level: character.level,
        talents: request.character.talents,
        source: character.source,
      },
      itemLevel,
    );
  });
</script>

<!-- The rail's tail: Your sims, a payer's server runs among them (spec 2026-11-06-sim). -->
{#snippet yourSims()}
  {#if signedIn && yourSimsLazy.current}<yourSimsLazy.current
      rows={historyRows}
      error={historyError}
      onretry={() => void loadHistory()}
      serverRows={hasServerSims ? serverRuns.rows : null}
      serverFailed={hasServerSims && serverRuns.failed}
      now={new Date(serverRuns.nowMs)}
      onserverretry={() => void serverRuns.reload()}
      onstarted={() => serverRuns.runStarted()}
      pageSpec={store.character?.spec ?? ''}
    />{/if}
{/snippet}

<div class="flex flex-col" data-testid="sim-view">
  {#if hasSavedSimId}
    {#if savedSim.result !== null}
      <!-- No fade here: the saved head is the page's largest paint, and an element first painted at
           opacity 0 is never an LCP candidate (Lighthouse reported no LCP on /sim/<id>). -->
      <div>
        <SavedSim
          result={savedSim.result}
          onrerun={onRerunSaved}
          {simId}
          fidelity={savedSpecRow}
          {initialNotes}
          {bisBenchmarks}
        />
      </div>
    {:else if savedSim.error === simCopy.notFound}
      <div class="sim-pad">
        <LoadError message={savedSim.error} onRetry={() => savedSim.load(simId)} testid="sim-saved-error" />
      </div>
    {:else if savedSim.error !== null}
      {#if savedFailedLazy.current}
        <savedFailedLazy.current onRetry={() => savedSim.load(simId)} />
      {:else if savedFailedLazy.error}
        <!-- The panel's own chunk failed to load: the same retry, as the shared load error. -->
        <div class="sim-pad">
          <LoadError message={savedSim.error} onRetry={() => savedSim.load(simId)} testid="sim-saved-error" />
        </div>
      {:else}
        <div class="sim-pad"><Skeleton lines={4} testid="sim-lazy-skeleton" /></div>
      {/if}
    {:else}
      <!-- The same static string sim/[id].astro renders before hydration (skeleton.ts). -->
      <!-- eslint-disable-next-line svelte/no-at-html-tags -->
      {@html SIM_SAVED_SKELETON_HTML}
    {/if}
  {:else if bootstrap.view === 'specs'}
    {#if specGridLazy.current}
      <specGridLazy.current
        rows={specRows}
        error={specsError}
        onretry={() => void loadSpecs()}
        benchmarks={bisBenchmarks}
      />
    {:else}
      <div class="sim-pad"><Skeleton lines={6} testid="sim-lazy-skeleton" /></div>
    {/if}
  {:else if store.character === null}
    {#if sessionPending || (landingBusyKey !== null && signedIn)}
      <div class="sim-pad">
        <Skeleton lines={4} minHeight="280px" testid="sim-landing-skeleton" />
      </div>
    {:else}
      <SimLandingPage
        {bisSets}
        busy={store.phase === 'loading-character'}
        message={store.message}
        onpaste={(code) => void runPastedInput(code)}
        onbuild={loadBuildInput}
        onfight={(ref) => void store.loadFight(fightRefOf(ref))}
        onsignin={onSignIn}
      >
        {#snippet notice()}
          {#if me !== null && me.characters.length === 0}
            <div class="sim-pad" data-testid="sim-no-characters">
              <Notice
                text={pageCopy.noCharacters}
                tone="neutral"
                action={{ label: pageCopy.noCharactersLink, href: pageCopy.noCharactersHref }}
              />
            </div>
          {:else if headerLoadFailed && headerCharacter !== null}
            <div class="sim-pad" data-testid="sim-landing-build-missing">
              {#if landingFailedKey !== null}
                <Notice
                  text={`${landingCopy.buildMissingLead(headerCharacter.name)} ${landingCopy.buildMissingPasteLink}, ${landingCopy.buildMissingMiddle}`}
                  tone="neutral"
                  action={{
                    label: landingCopy.buildMissingAccountLink,
                    href: landingCopy.buildMissingAccountHref,
                  }}
                />
              {:else}
                <Notice text={simCopy.characterFailed} tone="neutral" />
              {/if}
            </div>
          {/if}
        {/snippet}
        {#snippet railTail()}{@render yourSims()}{/snippet}
      </SimLandingPage>
    {/if}
  {:else}
    {@const character = store.character}
    <SimCharacterHead
      name={characterTitle(character.name, character.spec)}
      classSlug={character.class_slug}
      lines={headLines}
    >
      {#snippet tabs()}<SimToolTabs source={tabSource} {fallbackCode} keep={bootstrap.keep} />{/snippet}
      {#snippet answer()}
        <SimAnswer model={answerModel}>
          {#snippet badge()}
            {#if answerModel.state !== 'not-simulated'}<SimFidelityCaption fidelity={characterSpecRow} />{/if}
          {/snippet}
          {#snippet extra()}
            {#if store.message !== null && answerModel.state !== 'failed' && store.message !== simCopy.stopped}
              <span class="text-muted text-[13px] leading-5" role="alert" data-testid="sim-message"
                >{store.message}</span
              >
            {/if}
          {/snippet}
        </SimAnswer>
      {/snippet}
      {#snippet actions()}
        {#if answerModel.state === 'result' && store.result !== null}
          {#if savedUrl !== null}<span class="text-kill text-[13px] font-semibold"
              >{pageCopy.answer.saved}</span
            >{/if}
          <SimSaveActions
            result={store.result}
            reportTitle={store.reportTitle}
            {signedIn}
            {topGearHref}
            onsave={(title) => store.save(title)}
            onsignin={onSignIn}
            onsaved={(url) => (savedUrl = url)}
          />
        {:else if answerModel.state === 'not-simulated'}
          <a class="text-gold text-[14px] font-semibold" href={pageCopy.answer.specSupportHref}
            >{pageCopy.answer.specSupport}</a
          >
        {/if}
      {/snippet}
    </SimCharacterHead>
    <SimRunBar
      settings={store.settings}
      precisionId={store.precisionId}
      spec={character.spec}
      level={character.level}
      primary={store.needsRace && !busy ? 'disabled' : answerModel.primary}
      {busy}
      onsettings={(next) => store.setSettings(next)}
      onprecision={(value) => store.setPrecisionId(value)}
      onprimary={onPrimary}
      onmore={() => (moreOpen = true)}
    />
    {#if moreSheetLazy.current}
      <moreSheetLazy.current
        open={moreOpen}
        settings={store.settings}
        spec={character.spec}
        level={character.level}
        build={character.tree_version}
        classSlug={character.class_slug}
        names={store.buffNames}
        {request}
        disabled={busy}
        notifier={notifier !== null && store.premium}
        premium={store.premium}
        {signedIn}
        onserver={() => {
          moreOpen = false;
          runServer();
        }}
        {notifyWanted}
        onclose={() => (moreOpen = false)}
        onchange={(next) => store.setSettings(next)}
        onnotify={(wanted) => void toggleNotify(wanted)}
        onvalidate={(json) => store.validateRequest(json)}
        onapply={(next) => void store.applyRequest(next)}
        onrun={(next) => {
          moreOpen = false;
          stopped = null;
          lastLane = 'browser';
          void store.runRequest(next);
        }}
        onshare={(next) => shareUrlFor(next)}
      />
    {/if}
    {#if store.needsRace}
      <div class="sim-pad">
        <SimRacePick {character} races={store.races} onrace={(slug) => store.setRace(slug)} />
      </div>
    {/if}
    {#if comparing}
      <div class="sim-pad pt-7">
        {#if actual === null}
          <p class="text-muted text-[14px]" data-testid="compare-loading">{simCopy.compareLoading}</p>
        {:else if comparison !== null}
          {#if compareViewLazy.current}
            <compareViewLazy.current
              {comparison}
              simDuration={store.result?.summary.duration_ms ?? 0}
              actualDuration={actual.duration_ms}
            />
          {:else}
            <Skeleton lines={6} testid="sim-lazy-skeleton" />
          {/if}
        {/if}
      </div>
    {:else if store.result !== null && answerModel.state === 'result'}
      <div class="reveal">
        <RunReport result={store.result} live railTail={yourSims} {bisBenchmarks} />
      </div>
    {:else if request !== null}
      {#if preRunLazy.current}
        <preRunLazy.current {request} running={busy} railTail={yourSims} />
      {:else}
        <!-- The doll's Plate at its ready height (17 rows of 58 over two columns, the heads). -->
        <div class="sim-pad pt-7"><Skeleton lines={9} minHeight="560px" testid="sim-lazy-skeleton" /></div>
      {/if}
    {/if}
  {/if}
</div>

<style>
  .sim-pad {
    padding-inline: var(--page-inset);
    padding-top: 16px;
  }
</style>
