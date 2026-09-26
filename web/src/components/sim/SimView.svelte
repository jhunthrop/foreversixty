<!-- web/src/components/sim/SimView.svelte -->
<!-- The simulator island's root. It owns the URL state and the page frame; every section
     below it is a component that takes data and raises events, so this file stays a
     composition and never a place where a layout decision hides.

     The store reads its own URL bootstrap (a "Sim this build" link's `?code=`, or a saved
     build's `?source=build&ref=`) from init, not from a `$effect` here: `createSimStore`
     stays DOM-free and unit-testable, and this component's only job is to read
     `window.location` once and hand the values across. -->
<script lang="ts">
  import { onMount, untrack } from 'svelte';
  import activeBuild from '../../data/active-build.json';
  import { battlenetStartUrl, effectiveServerSims, fetchMeOnce, type Me } from '../../lib/account/api';
  import { sessionHinted } from '../../lib/data/query';
  import { createQueryState } from '../../lib/data/query.svelte';
  import { API_BASE_URL } from '../../lib/planner/config';
  import { SIM_LANDING_SKELETON_MIN_H } from '../../lib/sim/layout';
  import type { CharacterPath } from '../../lib/characters';
  import { readCurrent } from '../../lib/current-character';
  import { VIEW_GAP } from '../../lib/current-character-layout';
  import { readLastUpgrade, type LastUpgrade } from '../../lib/sim/last-upgrade';
  import { createLazyComponent, type LazyLoadState } from '../../lib/report/lazy-component.svelte';
  import { fetchReportMeta, fetchSummary } from '../../lib/report/load';
  import type { Summary } from '../../lib/report/types';
  import { fetchSpecs, listMySims } from '../../lib/sim/api';
  import { codeForCharacterSpec } from '../../lib/sim/character';
  import { decideBootstrap, RESTORE_BUSY_KEY, runBootstrapRestore } from '../../lib/sim/character-bootstrap';
  import { compareSummaries } from '../../lib/sim/compare';
  import { simCopy } from '../../lib/sim/copy';
  import type { KindFilter } from '../../lib/sim/history';
  import { SIM_LAZY_MIN_H } from '../../lib/sim/lazy-layout';
  import { browserNotifier, enableNotifications, notifyFinished } from '../../lib/sim/notify';
  import { createSavedSimState } from '../../lib/sim/saved-sim-state.svelte';
  import { SIM_SAVED_SKELETON_HTML } from '../../lib/sim/skeleton';
  import { parseFightRef } from '../../lib/sim/sources';
  import {
    mergeSpecRows,
    needsFidelityNote,
    specPillClass,
    specStateLabel,
    specStateNote,
  } from '../../lib/sim/spec-state';
  import { createSimStore } from '../../lib/sim/store.svelte';
  import { syncTabHrefs } from '../../lib/sim/tabs';
  import { landingCopy } from '../../lib/sim/landing-copy';
  import {
    decodeRequestParam,
    defaultSimState,
    encodeRequestParam,
    parseSimState,
    simSearch,
    withSimState,
  } from '../../lib/sim/url';
  import { ENGINE_VERSION, engineLabel, isStale } from '../../lib/sim/version';
  import type { SimListRow, SimRequest, SimResult, SpecFidelity } from '../../lib/sim/types';
  import BuffPanel from './BuffPanel.svelte';
  import CharacterStrip from './CharacterStrip.svelte';
  import CurrentCharacterBar from '../CurrentCharacterBar.svelte';
  import DetailsCard from './DetailsCard.svelte';
  import LandingState from './LandingState.svelte';
  import ReportOptions from './ReportOptions.svelte';
  import RequestDrawer from './RequestDrawer.svelte';
  import RotationCard from './RotationCard.svelte';
  import RunControl from './RunControl.svelte';
  import SavedSim from './SavedSim.svelte';
  import SettingsBar from './SettingsBar.svelte';
  import SourceSwitcher from './SourceSwitcher.svelte';
  import SpecGrid from './SpecGrid.svelte';
  import LoadError from '../ui/LoadError.svelte';
  import Skeleton from '../ui/Skeleton.svelte';
  import { BUSY_CLASS } from '../../lib/ui/busy';

  let { simId = '', inlineResult = null }: { simId?: string; inlineResult?: SimResult | null } = $props();

  // True on /sim/<id> -- a saved sim, `sim-island.ts`'s own `simIdFor` already resolved off
  // the mount's data or the path -- and on the prerendered fixture page, which inlines its
  // result. Task 17 renders that view; the character strip, the source switcher and the
  // empty prompt below stay out of the way rather than showing a builder under it.
  const hasSavedSimId = untrack(() => simId !== '' || inlineResult !== null);

  // Read once, like simId and inlineResult above: these are the page's own one-shot
  // bootstrap values, not bindings this island keeps synced against a changing URL.
  const bootstrap = untrack(() => {
    const search = window.location.search;
    const { source, ref, code, req, mode } = parseSimState(search);
    // The mount element the shell can stamp a build id onto, the way planner-island.ts
    // reads `data-tree-version` off its own mount -- no /sim page stamps one yet, so this
    // falls back to the site's active build rather than an empty string no fetch would
    // resolve.
    const mount = document.getElementById('sim');
    const treeVersion = mount?.dataset.treeVersion ?? activeBuild.build;
    // specs.astro stamps `data-sim-view="specs"`; sim.astro and [id].astro stamp neither,
    // so an absent or unrecognised value reads as the ordinary simulator.
    const view: 'sim' | 'specs' = mount?.dataset.simView === 'specs' ? 'specs' : 'sim';
    return { treeVersion, source, ref, code, request: decodeRequestParam(req), mode, view };
  });

  // Compare mode loads its character through `enterCompare` below, never through the
  // store's own URL bootstrap: a second, redundant bootstrap racing `enterCompare`'s own
  // could land after `store.run()` and null out the result `comparison` was built from.
  const store = untrack(() =>
    createSimStore({
      treeVersion: bootstrap.treeVersion,
      request: bootstrap.request ?? undefined,
      source: bootstrap.mode === 'compare' ? undefined : bootstrap.source,
      ref: bootstrap.mode === 'compare' ? undefined : bootstrap.ref,
      code: bootstrap.code,
    }),
  );

  // Compare mode's own state: `actual` is the logged fight's raw summary, and `comparing`
  // gates every compare-only branch below so plain /sim never has to think about either.
  let actual = $state<Summary | null>(null);
  let comparing = $state(false);

  /**
   * Compare mode's whole entry path, called once from the same place the other `?source=`
   * loaders are called -- never from a `$effect`, which would re-fetch the fight whenever
   * anything unrelated in the store changed.
   */
  async function enterCompare(ref: string): Promise<void> {
    const parsed = parseFightRef(ref);
    if (parsed === null) {
      store.setMessage(simCopy.fightRefInvalid);
      return;
    }
    comparing = true;
    // loadFight gives us the character; fromLoggedFight is what it calls underneath.
    await store.loadFight(ref);
    try {
      const meta = await fetchReportMeta(parsed.reportId);
      actual = await fetchSummary(meta.data_base_url, parsed.fightIndex);
    } catch {
      // A dead page helps nobody: keep the character, drop back to plain sim mode.
      actual = null;
      comparing = false;
      store.setMessage(simCopy.fightNoCombatant);
      return;
    }
    // The one place a sim starts without a press: the player already pressed something to
    // get here, and a compare with nothing to compare against is not a page.
    if (store.character !== null) await store.run();
  }

  if (bootstrap.mode === 'compare') void enterCompare(bootstrap.ref);

  // /sim/<sim_id>'s own entry path (Task 17), the saved-sim counterpart to `enterCompare`
  // above: a prerendered fixture page already carries its result (`inlineResult`), so only
  // the id-only case fetches. Called once, here, rather than from an `$effect` -- the same
  // reason `enterCompare` is not one.
  const savedSim = createSavedSimState(untrack(() => inlineResult));

  // One-time init read, the same reason bootstrap and store above are wrapped: the
  // .then/.catch callbacks run later, as ordinary reactive writes.
  untrack(() => {
    if (hasSavedSimId && savedSim.result === null && simId !== '') savedSim.load(simId);
  });

  /**
   * "Run this yourself": opens /sim with a character loaded, via the same `sim/url.ts`
   * vocabulary this file's own bootstrap uses. `addon`/`manual` sources persist no `ref`
   * (H4, final whole-branch review), so those fall back to the saved result's own
   * `request.character`, converted to an FS1 code and handed to `?code=` instead.
   */
  function onRerunSaved(): void {
    if (savedSim.result === null) return;
    const { kind, ref } = savedSim.result.request.source;
    const target =
      ref !== ''
        ? withSimState(defaultSimState(), { source: kind, ref })
        : withSimState(defaultSimState(), {
            // codeForCharacterSpec carries the saved result's own gear list, enchants and
            // suffixes included (contract 10.5) -- character.ts's own reason.
            code: codeForCharacterSpec(savedSim.result.request.character, bootstrap.treeVersion),
          });
    window.location.href = `/sim${simSearch(target)}`;
  }

  /** Design 8: a share URL of an edited request is a full reproduction. Null past the budget. */
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

  // CompareView is never needed for /sim's first paint -- it only exists once a compare
  // link is followed -- so it ships as its own chunk, the same way SimResults does below.
  const compareViewLazy = createLazyComponent(() => import('./CompareView.svelte'));
  $effect(() => {
    if (comparison !== null) compareViewLazy.load();
  });

  // One `/v1/me` read, shared with every other island through the client cache
  // (web/src/lib/data/query.ts) -- the same call Account.svelte, SessionNav.svelte and
  // HomeAccountPanel.svelte make. Read by the history panel, the landing state and the
  // source switcher's signed-in card.
  const session = createQueryState<Me | null>(`${API_BASE_URL}/v1/me`, () => fetchMeOnce(), {
    scope: 'private',
    ttlMs: 10 * 60 * 1000,
  });
  const me = $derived(session.data);
  // The landing skeleton shows only for a visitor who has a session to load (the cookie's
  // readable half is present) and no snapshot yet; a signed-out visitor gets the source
  // switcher at once, as before, rather than a skeleton for a read that answers "nobody".
  const sessionPending = $derived(session.status === 'loading' && session.data === null && sessionHinted());
  // The history panel (Task 17), for a signed-in player on plain /sim only.
  const signedIn = $derived(me !== null);
  let historyRows = $state<SimListRow[] | null>(null);
  let historyError = $state<string | null>(null);
  // The history filter (Task 18). "all" sends no `kind=` at all -- see api.ts's listMySims.
  let historyKind = $state<KindFilter>('all');

  async function loadHistory(): Promise<void> {
    historyError = null;
    try {
      const page = await listMySims(1, undefined, historyKind);
      historyRows = page.rows;
    } catch (error) {
      historyRows = null;
      historyError = error instanceof Error ? error.message : simCopy.loadFailed;
    }
  }

  function setHistoryKind(next: KindFilter): void {
    historyKind = next;
    historyRows = null;
    void loadHistory();
  }

  // Loaded lazily (Task 17): the history panel is empty weight for every signed-out
  // visitor and for /sim/<id>, where it never renders at all, and this file's compiled
  // bundle is shared by all three /sim routes.
  const simHistoryLazy = createLazyComponent(() => import('./SimHistory.svelte'));
  $effect(() => {
    if (signedIn) simHistoryLazy.load();
  });

  // Design 5.4: the finish notification. `notifier` is the real Notification API, or null
  // where the browser has none, read once. Permission is asked only from `toggleNotify`.
  const notifier = untrack(() => browserNotifier());

  // Spec 2026-09-25 §6: the after-sim sentence's data source, read once -- same one-shot
  // pattern as `notifier`/`bootstrap`/`store` above. No re-read on later state changes: a
  // Droptimizer run made *during this same /sim session* (impossible today -- they are
  // separate page loads) would need a fresh read, but nothing on this page can produce one
  // without a navigation, which remounts the island anyway.
  const topUpgrade = untrack<LastUpgrade | null>(() => readLastUpgrade());
  let notifyWanted = $state(false);
  // The id of the last result a notification was raised for, so a re-render never raises a
  // second one for the same run.
  let notifiedFor = $state('');

  async function toggleNotify(wanted: boolean): Promise<void> {
    notifyWanted = wanted && (await enableNotifications(notifier));
  }

  // Server runs only, per design 5.4: a browser run finishes on the tab you are looking at.
  $effect(() => {
    const finished = store.result;
    if (finished !== null && finished.lane === 'server') {
      // The key is the wall clock plus the figure, which no two runs of one session share.
      const key = `${finished.duration_ms}-${finished.iterations_run}`;
      if (key !== notifiedFor) {
        notifiedFor = key;
        notifyFinished(
          notifier,
          notifyWanted,
          store.reportTitle,
          simCopy.notifyBody(Math.round(finished.dps.mean).toLocaleString('en-US')),
        );
      }
    }
  });

  // The save form: `saveOpen`/`saveTitle` are the inline form, `savedUrl` replaces it on success.
  let saveOpen = $state(false);
  let saveTitle = $state('');
  let saving = $state(false);
  let saveFailed = $state(false);
  let savedUrl = $state<string | null>(null);
  let savedLinkCopied = $state(false);

  $effect(() => {
    void store.result;
    saveOpen = false;
    saveFailed = false;
    savedUrl = null;
    savedLinkCopied = false;
  });

  // A result the run loop reports as stopped rather than finished (`sim/api`'s additive
  // `aborted`) has nothing complete to save -- the button stays disabled and says why,
  // rather than saving a partial run under a title the player chose for a real result.
  const canSave = $derived(store.result !== null && store.result.aborted !== true);

  function openSaveForm(): void {
    saveTitle = store.reportTitle;
    saveFailed = false;
    savedUrl = null;
    saveOpen = true;
  }

  function cancelSave(): void {
    saveOpen = false;
    saveFailed = false;
  }

  async function confirmSave(): Promise<void> {
    saving = true;
    saveFailed = false;
    const id = await store.save(saveTitle);
    saving = false;
    if (id === null) {
      saveFailed = true;
      return;
    }
    saveOpen = false;
    savedUrl = `${window.location.origin}/sim/${id}`;
  }

  async function copySavedLink(): Promise<void> {
    if (savedUrl === null) return;
    try {
      await navigator.clipboard.writeText(savedUrl);
      savedLinkCopied = true;
    } catch {
      savedLinkCopied = false;
    }
  }

  // effectiveServerSims(me) on GET /v1/me -- the server lane renders only once this answers
  // true. A signed-out visitor and an unreachable API read the same way (`me` stays null
  // whether the session read answered null or failed; the sim never shows a session error,
  // spec 2026-09-23 §3). The same answer gates the history panel, the landing state and the
  // source switcher's signed-in card: `signedIn` above is `me !== null`, and the history
  // panel's own `GET /v1/sims?mine=1` fires only then. Keyed on `me`'s reference and guarded
  // against re-running for the same object (the 2026-09-23 account-page loop); it never
  // starts a session read itself, only reacts to the one `createQueryState` made above.
  let lastMe: Me | null = null;
  $effect(() => {
    if (me === lastMe) return;
    lastMe = me;
    store.setPremium(effectiveServerSims(me));
    if (me !== null) void loadHistory();
  });

  onMount(() => {
    // Never on /sim/<id>: a saved sim never restores the visitor's own current character.
    if (!hasSavedSimId) void restoreFromPointer();
    return () => store.dispose();
  });

  // The spec support list: /sim/specs renders it as a grid, and /sim needs it too, to know
  // whether the loaded character's own spec is one the engine models -- both share it.
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

  // No reactive read inside, so this fires once, on mount -- the effect form Task 15's
  // brief calls for, since /sim/specs' own grid needs exactly this one-shot load.
  $effect(() => {
    void loadSpecs();
  });

  /**
   * The loaded character's own row, once the spec list has answered. Null while the list is
   * still loading or failed, or before a character is on screen -- `null` reads as "unknown
   * yet", not as "unsupported", so the run control stays up rather than flashing the
   * unsupported card ahead of the real answer.
   */
  const characterSpecRow = $derived.by((): SpecFidelity | null => {
    if (store.character === null || specRows === null) return null;
    return mergeSpecRows(specRows).find((row) => row.spec === store.character?.spec) ?? null;
  });

  // False until the player explicitly asks for the switcher. Design 4.6: a signed-in
  // member with characters opens on the landing state instead, so defaulting this to
  // `store.character === null` would show the switcher on every first paint instead.
  let switcherOpen = $state(false);

  $effect(() => {
    if (store.character !== null) switcherOpen = false;
  });

  // `restored` is only ever set by `restoreFromPointer`; the spine bar shows the note.
  let restored = $state(false);

  /**
   * A fresh FS1 v2 code for the loaded character, the same conversion "Run this yourself"
   * above falls back to. Only reached when `store.character.source.ref` is empty (fix
   * round 1, Finding A) -- an addon paste or a `?code=` link -- so `store.buildRequest()`
   * runs only then, not on every character.
   */
  function fallbackTabCode(): string | null {
    if (store.character === null || store.character.source.ref !== '') return null;
    const request = store.buildRequest();
    return request === null ? null : codeForCharacterSpec(request.character, bootstrap.treeVersion);
  }

  // Keeps the loaded character on every tab in SimTabs.astro's strip (task-1-brief.md):
  // whenever the character this island holds changes -- loaded, changed source, or cleared
  // -- every tab's href is rewritten to carry the same query its own destination can
  // actually bootstrap from (`?source=&ref=`, or the `?code=` fallback every tab now reads,
  // `SIM_TABS`' own `supportsCode` -- current-character spec, 2026-09-21). The strip lives
  // above this island's own mount point (SimTabs.astro's own comment), so `syncTabHrefs`
  // reaches it through `document` rather than this component's own root.
  $effect(() => {
    const source = store.character === null ? null : store.character.source;
    syncTabHrefs(source, fallbackTabCode());
  });

  function onSignIn(): void {
    window.location.href = battlenetStartUrl(`${window.location.pathname}${window.location.search}`);
  }

  // The landing state's own busy key (Task 18): the row a pick is in flight for, so its
  // button reads "Loading…" while every other row disables rather than reads it too.
  let landingBusyKey = $state<string | null>(null);
  // Finding 2, 2026-09-24 landing pass: the key of the character a pick just failed for,
  // want of a build (the sim-input 404) -- set only then, never for the race refusal, which
  // keeps its own hint below (the `sim-landing-message` paragraph, gated on this being
  // null). `store.message` carries no status code, so this compares it against
  // `fetchSimInput`'s own 404 sentinel (`sim/api.ts`) rather than guessing from the text.
  let landingFailedKey = $state<string | null>(null);

  async function pickCharacter(path: CharacterPath): Promise<void> {
    const key = `${path.region}/${path.ruleset}/${path.slug}`;
    landingBusyKey = key;
    landingFailedKey = null;
    await store.loadStored(path);
    landingBusyKey = null;
    landingFailedKey = store.message === landingCopy.buildMissingFallback ? key : null;
  }

  // `store.ready` already ran the URL's own bootstrap; `runBootstrapRestore` (shared with
  // ToolsView.svelte) only fires the stored-pointer fallback once that settled with
  // nothing. `landingBusyKey` is held for the same window (fix round 1): `LandingState`
  // disables its picks off that prop alone, never off `store.phase`, and `adopt()` has no
  // per-load generation guard, so a fast pick could otherwise race this background load.
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
      restored = await runBootstrapRestore(store, url, stored, () => store.character !== null);
    } finally {
      landingBusyKey = null;
    }
  }

  // The stale-engine banner describes a *settled* result: while either lane is running,
  // `store.result` still holds the old result, so this gate keeps its "Run again" from
  // rendering alongside a live "Loading engine…"/"Stop".
  const staleVersion = $derived(
    store.result !== null &&
      isStale(store.result.engine_version) &&
      store.phase !== 'running' &&
      store.phase !== 'loading-engine' &&
      !store.serverRunning
      ? store.result.engine_version
      : null,
  );

  // SimResults pulls in five report components and is never needed for /sim's first
  // paint -- the page opens on an empty state with no result at all -- so it ships as its
  // own chunk, loaded the moment a result actually lands, the same way ReportView.svelte
  // lazy-loads Compare, Mechanics, Rankings, Timelines, Events and Queries.
  const simResultsLazy = createLazyComponent(() => import('./SimResults.svelte'));
  $effect(() => {
    if (store.result !== null) simResultsLazy.load();
  });
</script>

{#snippet lazyFallback(lazy: LazyLoadState, minHeight: string)}
  {#if lazy.error !== ''}
    <!-- LoadError has no class-passthrough prop, so the phone gutter the paragraph it
         replaces carried (px-[18px] md:px-0) lives on this wrapper instead of the shared
         primitive -- the same fix the report island's top-level error needed. -->
    <div class="px-[18px] md:px-0">
      <LoadError message={lazy.error} onRetry={() => lazy.load()} testid="sim-results-error" />
    </div>
  {:else}
    <Skeleton {minHeight} testid="sim-lazy-skeleton" />
  {/if}
{/snippet}

<div class={`flex flex-col ${VIEW_GAP}`} data-testid="sim-view">
  <!-- The spine bar is the one current-character band on this page (the old chip said the
       same character a second time, 44px below it); it reserves CHIP_HEIGHT itself. -->
  <CurrentCharacterBar spine currentDoor="sim" {restored} />
  {#if hasSavedSimId}
    <!-- /sim/<sim_id>: read-only, and not the sim page with a result in it -- no switcher,
         no settings bar, no run control. SavedSim composes its own heading. -->
    {#if savedSim.result !== null}
      <!-- `.reveal` goes on the wrapper of the ready state only (design 2026-09-22 spec
           section 1.4), and SavedSim takes no class prop -- so this one-child div carries
           it. A single stretched flex item in place of the component's own root: the
           column's gap sits outside it either way, so nothing moves. -->
      <div class="reveal">
        <SavedSim result={savedSim.result} onrerun={onRerunSaved} />
      </div>
    {:else if savedSim.error !== null}
      <!-- Same reason as the lazyFallback wrapper above: LoadError has no class-passthrough prop. -->
      <div class="px-[18px] md:px-0">
        <LoadError message={savedSim.error} onRetry={() => savedSim.load(simId)} testid="sim-saved-error" />
      </div>
    {:else}
      <!-- Between mounting and a real, non-prerendered id's fetch resolving. Static,
           trusted markup of our own (skeleton.ts): no data goes into it, and it is the
           same string sim/[id].astro's own shell renders before hydration, so nothing
           shifts between the two moments -- report/skeleton.ts's own reason. -->
      <!-- eslint-disable-next-line svelte/no-at-html-tags -->
      {@html SIM_SAVED_SKELETON_HTML}
    {/if}
  {:else}
    <!-- No <h1> here: sim.astro and sim/specs.astro each carry their own, in the static
         HTML, ahead of this island entirely -- an island-mounted heading is invisible to
         Lighthouse's LCP measurement and, before this, was "Simulator" on both pages (M7,
         final whole-branch review), which the spec-support page's own title disagreed
         with. -->
    <!-- Finding 6, 2026-09-24 landing pass: the engine-version hash used to render here,
         unconditionally, ahead of every state below -- including the landing state, which
         it preceded with a lone right-aligned hash before a signed-in member had even
         loaded a character. It now renders once a character is loaded, right after
         SettingsBar below -- not inside SettingsBar.svelte itself, which TopGear.svelte's
         bulk tool pages (/sim/gear, /sim/talents, /sim/drops, /sim/weights) also render:
         putting it there leaked a sub-44px control onto pages that never budgeted for it
         (sim-tools-phone.spec.ts's own audit). This keeps it scoped to plain /sim, where
         "the engine matters" actually means something. -->

    {#if bootstrap.view === 'specs'}
      <!-- Round 2 (Lighthouse): specs-intro is static now, in sim/specs.astro, ahead of
           this island -- it was the measured LCP element and this branch used to render a
           second, redundant copy of the exact same paragraph. Nothing here replaces it. -->
      <SpecGrid rows={specRows} error={specsError} onretry={() => void loadSpecs()} />
    {:else}
      {#if store.character !== null && !switcherOpen}
        <CharacterStrip
          character={store.character}
          items={store.items}
          races={store.races}
          onchange={() => (switcherOpen = true)}
          onrace={(slug) => store.setRace(slug)}
        />
      {:else if sessionPending && !switcherOpen}
        <!-- Spec 2026-09-23 §3: a cold cache shows a reserved skeleton while the session read
             is in flight; a returning signed-in visitor's snapshot skips this entirely. -->
        <Skeleton lines={4} minHeight={SIM_LANDING_SKELETON_MIN_H} testid="sim-landing-skeleton" />
      {:else if me !== null && me.characters.length > 0 && !switcherOpen}
        <!-- Design 4.6: a signed-in member sees their characters and one button each, and
             no form until they ask for one -- so this replaces the switcher entirely rather
             than sitting above it. -->
        <LandingState
          characters={me.characters}
          busyKey={landingBusyKey}
          failedKey={landingFailedKey}
          onpick={(path) => void pickCharacter(path)}
          onother={() => (switcherOpen = true)}
        />
        {#if store.message !== null && landingFailedKey === null}
          <!-- The race refusal (sources.ts's "no race recorded", a combat log carries none
               and the API has not started sending one for a stored character either) is the
               one failure this paragraph still shows -- the sim-input 404 (Finding 2, 2026-
               09-24 landing pass) has its own alert inside LandingState now, and
               `landingFailedKey` gates this one off whenever that is the failure on
               screen, so the two never both render for the same pick. -->
          <p class="text-muted px-[18px] text-[14px] md:px-0" role="alert" data-testid="sim-landing-message">
            {store.message}
            {simCopy.landingNoRace}
          </p>
        {/if}
      {:else if me !== null && me.characters.length === 0}
        <p class="text-muted px-[18px] text-[14px] md:px-0" data-testid="sim-no-characters">
          {simCopy.noCharactersYet}
          <a class="text-nav underline" href="/logs">{simCopy.noCharactersYetLink}</a>
        </p>
        <SourceSwitcher
          busy={store.phase === 'loading-character'}
          message={store.message}
          signedIn
          hasCharacters={false}
          onaddon={(code) => void store.loadAddon(code)}
          onbuild={(id) => void store.loadBuild(id)}
          onfight={(ref) => void store.loadFight(ref)}
          onsignin={onSignIn}
          onback={() => (switcherOpen = false)}
        />
      {:else}
        <SourceSwitcher
          busy={store.phase === 'loading-character'}
          message={store.message}
          signedIn={me !== null}
          onaddon={(code) => void store.loadAddon(code)}
          onbuild={(id) => void store.loadBuild(id)}
          onfight={(ref) => void store.loadFight(ref)}
          onsignin={onSignIn}
          onback={() => (switcherOpen = false)}
        />
      {/if}

      <p class="text-muted px-[18px] text-[14px] md:px-0" data-testid="sim-scope-note">
        {landingCopy.scopeCaveat}
      </p>

      {#if signedIn && simHistoryLazy.current}
        <simHistoryLazy.current
          rows={historyRows}
          error={historyError}
          kind={historyKind}
          onkind={setHistoryKind}
        />
      {/if}

      {#if store.character !== null}
        <!-- Everything a loaded character brings with it, under the one wrapper `.reveal`
             is allowed to sit on (design 2026-09-22 spec section 1.4). The wrapper repeats
             the view's own `flex flex-col VIEW_GAP` so these panels keep the exact spacing
             they had as direct children of `sim-view`; none of them sets an `align-self`
             or a flex ratio, so nesting them one level deeper changes nothing else. -->
        <div class={`reveal flex flex-col ${VIEW_GAP}`}>
          <SettingsBar
            settings={store.settings}
            spec={store.character.spec}
            disabled={store.phase === 'running' || store.serverRunning}
            onchange={(next) => store.setSettings(next)}
            names={store.buffNames}
          />
          <!-- Finding 6: the engine hash, right after the settings bar it used to sit above
               every state -- same anchor, same classes, same test id and href, now visible
               only once a character has actually reached the engine. -->
          <div class="flex justify-end px-[18px] md:px-0">
            <a
              class="tabular text-muted font-mono text-[12px]"
              href="/sim/specs"
              data-testid="sim-engine-version">{engineLabel(ENGINE_VERSION)}</a
            >
          </div>
          {#if store.settings.preset === 'custom'}
            <BuffPanel
              settings={store.settings}
              build={store.character.tree_version}
              names={store.buffNames}
              disabled={store.phase === 'running' || store.serverRunning}
              onchange={(next) => store.setSettings(next)}
            />
          {/if}
          <RequestDrawer
            request={store.buildRequest()}
            disabled={store.phase === 'running' || store.serverRunning}
            onvalidate={(json) => store.validateRequest(json)}
            onapply={(request) => void store.applyRequest(request)}
            onrun={(request) => void store.runRequest(request)}
            onshare={(request) => shareUrlFor(request)}
          />
          {#if characterSpecRow !== null && needsFidelityNote(characterSpecRow)}
            <!-- A fidelity state labels, it never blocks: the run control below always
               renders once a character is loaded, and this is the one-line footnote
               linking to the full card on /sim/specs. -->
            <p class="text-muted px-[18px] text-[13px] md:px-0" data-testid="spec-fidelity-note">
              <a href="/sim/specs" class={specPillClass(characterSpecRow.state)}
                >{specStateLabel(characterSpecRow.state)}</a
              >
              {specStateNote(characterSpecRow.state)}
            </p>
          {/if}
          <RunControl
            spec={store.character.spec}
            phase={store.phase}
            estimate={store.estimate}
            iterationsDone={store.iterationsDone}
            iterationsTotal={store.iterationsTotal}
            precisionId={store.precisionId}
            relativeError={store.relativeError}
            lane={store.lane}
            premium={store.premium}
            signedIn={me !== null}
            message={store.message}
            detail={store.detail}
            racePending={store.needsRace}
            {staleVersion}
            serverRunning={store.serverRunning}
            onrun={() => void store.run()}
            onstop={() => store.stop()}
            onprecision={(value) => store.setPrecisionId(value)}
            onserver={() => void store.runOnServer()}
            onrerun={() => void store.run()}
          />
          {#if comparing}
            <!-- Replaces the sentence and the results, per Design 4.2 -- the strip, the
                 settings bar and the run control above stay exactly where they are. -->
            {#if actual === null}
              <p class="text-muted px-[18px] text-[14px] md:px-0" data-testid="compare-loading">
                {simCopy.compareLoading}
              </p>
            {:else if comparison !== null}
              {#if compareViewLazy.current}
                <compareViewLazy.current
                  {comparison}
                  simDuration={store.result?.summary.duration_ms ?? 0}
                  actualDuration={actual.duration_ms}
                />
              {:else}
                {@render lazyFallback(compareViewLazy, SIM_LAZY_MIN_H.compare)}
              {/if}
            {/if}
          {:else if store.result !== null}
            {#if simResultsLazy.current}
              <simResultsLazy.current
                summary={store.result.summary}
                estimate={store.result.dps}
                iterationsRun={store.result.iterations_run}
                actionNames={store.actionNames}
                sample={store.result.sample}
                {topUpgrade}
              />
            {:else}
              {@render lazyFallback(simResultsLazy, SIM_LAZY_MIN_H.results)}
            {/if}
          {/if}

          {#if store.result !== null && !comparing}
            <DetailsCard result={store.result} />
            <RotationCard spec={store.result.request.spec} fidelity={characterSpecRow} />
          {/if}

          {#if store.result !== null}
            <!-- Design 5.4: the report title and the finish notification. The title feeds
               the save form, the notification and the saved link (openSaveForm reads
               store.reportTitle below); the notification checkbox only appears where the
               browser actually has a Notification API to ask (`notifier`, owned by this
               file since the $effect above needs it in component scope). -->
            <ReportOptions
              title={store.reportTitle}
              notifierAvailable={notifier !== null}
              {notifyWanted}
              ontitlechange={(value) => store.setReportTitle(value)}
              onnotifychange={(wanted) => void toggleNotify(wanted)}
            />
          {/if}

          <!-- The save form (Task 17): disabled until there is a result, an inline
               title field pre-filled with the report title rather than a dialog, and
               the saved link shown in place -- the page never navigates away from the
               result it just saved. -->
          <div class="mx-[18px] flex flex-wrap items-center gap-3 md:mx-0" data-testid="sim-save">
            {#if savedUrl !== null}
              <label class="sr-only" for="sim-save-link">{simCopy.savedLinkLabel}</label>
              <input
                id="sim-save-link"
                type="text"
                readonly
                value={savedUrl}
                class="border-line-warm rounded-control bg-raised text-text h-11 min-w-0 flex-1 border px-3 text-[14px] md:max-w-[420px]"
                data-testid="sim-save-link"
                onclick={(event) => event.currentTarget.select()}
              />
              <button
                type="button"
                class="border-line-warm rounded-control text-nav label min-h-11 border px-4"
                onclick={() => void copySavedLink()}
                data-testid="sim-save-copy"
              >
                {savedLinkCopied ? simCopy.copied : simCopy.copyLink}
              </button>
              <a
                class="border-line-warm rounded-control text-nav label inline-flex min-h-11 items-center border px-4"
                href={savedUrl}
                target="_blank"
                rel="noopener"
                data-testid="sim-open-new-tab"
              >
                {simCopy.openInNewTab}
              </a>
            {:else if saveOpen}
              <label class="flex flex-col gap-1">
                <span class="label text-muted">{simCopy.saveTitleLabel}</span>
                <input
                  type="text"
                  bind:value={saveTitle}
                  class="border-line-warm rounded-control bg-raised text-text h-11 w-[260px] border px-3 text-[14px]"
                  data-testid="sim-save-title"
                />
              </label>
              <button
                type="button"
                class={`border-line-warm-strong rounded-control bg-card-top text-strong label min-h-11 border px-5 disabled:opacity-50 ${saving ? BUSY_CLASS : ''}`}
                disabled={saving}
                aria-busy={saving}
                onclick={() => void confirmSave()}
                data-testid="sim-save-confirm"
              >
                {simCopy.saveAction}
              </button>
              <button
                type="button"
                class="border-line-warm rounded-control text-nav label min-h-11 border px-4"
                onclick={cancelSave}
                data-testid="sim-save-cancel"
              >
                {simCopy.cancel}
              </button>
              {#if saveFailed}
                <span role="alert" class="text-strong text-[13px]" data-testid="sim-save-error">
                  {simCopy.saveFailed}
                </span>
              {/if}
            {:else}
              <button
                type="button"
                class="border-line-warm-strong rounded-control bg-card-top text-strong label min-h-11 border px-5 disabled:opacity-50"
                disabled={!canSave}
                onclick={openSaveForm}
                data-testid="sim-save-open"
              >
                {simCopy.saveThisSim}
              </button>
              <!-- Task 7: the disabled reason, said plainly (Task 3's pattern), never a title=. -->
              {#if store.result !== null && !canSave}
                <p class="text-muted text-[12px]" data-testid="sim-save-disabled-note">
                  {simCopy.saveAbortedDisabled}
                </p>
              {/if}
            {/if}
          </div>
        </div>
      {:else if me === null || me.characters.length === 0}
        <p class="text-muted px-[18px] text-[14px] md:px-0" data-testid="sim-empty">
          {simCopy.emptyPrompt}
        </p>
      {/if}
    {/if}
  {/if}
</div>
