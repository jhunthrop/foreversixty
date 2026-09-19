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
  import { battlenetStartUrl, fetchMe, type Me } from '../../lib/account/api';
  import type { CharacterPath } from '../../lib/characters';
  import { encodeFS1V2 } from '../../lib/planner/fs1';
  import type { Slot } from '../../lib/planner/types';
  import { createLazyComponent, type LazyLoadState } from '../../lib/report/lazy-component.svelte';
  import { fetchReportMeta, fetchSummary } from '../../lib/report/load';
  import type { Summary } from '../../lib/report/types';
  import { fetchSim, fetchSpecs, listMySims } from '../../lib/sim/api';
  import { gearFromSlots, ranksFromTalentsString } from '../../lib/sim/character';
  import { compareSummaries } from '../../lib/sim/compare';
  import { simCopy } from '../../lib/sim/copy';
  import { browserNotifier, enableNotifications, notifyFinished } from '../../lib/sim/notify';
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

  let { simId = '', inlineResult = null }: { simId?: string; inlineResult?: SimResult | null } = $props();

  // True on /sim/<id> -- a saved sim, `sim-island.ts`'s own `simIdFor` already resolved off
  // the mount's data or the path -- and on the prerendered fixture page, which inlines its
  // result for Lighthouse. Task 17 renders that view; the character strip, the source
  // switcher and the empty prompt below are /sim's own UI, not /sim/<id>'s, so they stay out
  // of the way rather than showing a builder under content that has not landed yet.
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
  // store's own URL bootstrap: both call the identical `fromLoggedFight(ref)`, and
  // `adopt()` unconditionally clears `result` on every successful load, so a second,
  // redundant bootstrap racing `enterCompare`'s own bootstrap could land after
  // `store.run()` and null out the sim result `comparison` was just built from. One
  // loader for one entry path avoids that race rather than trusting the two to agree on
  // an order they are never sequenced to keep.
  const store = untrack(() =>
    createSimStore({
      treeVersion: bootstrap.treeVersion,
      request: bootstrap.request ?? undefined,
      source: bootstrap.mode === 'compare' ? undefined : bootstrap.source,
      ref: bootstrap.mode === 'compare' ? undefined : bootstrap.ref,
      code: bootstrap.code,
    }),
  );

  // Compare mode's own state: `actual` is the logged fight's raw summary, read straight
  // off the report lane's own loader (no second parse, no new API), and `comparing` gates
  // every compare-only branch below so plain /sim never has to think about either.
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
  let savedResult = $state<SimResult | null>(untrack(() => inlineResult));
  let savedError = $state<string | null>(null);

  // One-time init reads of savedResult and simId, the same reason bootstrap and store
  // above are wrapped: this runs once, at setup, never again as either changes. The
  // .then/.catch callbacks below run later, as ordinary reactive writes -- untrack only
  // covers the synchronous read that kicks the fetch off.
  untrack(() => {
    if (hasSavedSimId && savedResult === null && simId !== '') {
      void fetchSim(simId)
        .then((result) => (savedResult = result))
        .catch((error) => {
          savedError = error instanceof Error ? error.message : simCopy.loadFailed;
        });
    }
  });

  /**
   * "Run this yourself", the stale-result remedy and the ordinary way off a saved page:
   * opens /sim with a character loaded, so the player can change something and run it
   * themselves. It is a callback rather than an `<a href>` SavedSim builds itself so the
   * URL is built with the same `sim/url.ts` vocabulary this file already owns for its own
   * bootstrap, in one place. It is never triggered automatically -- only this handler,
   * from the player's own click.
   *
   * `source.kind`+`ref` round-trips correctly for `build` and `fight`: both always carry a
   * non-empty `ref` (fromPlannerBuild, fromLoggedFight) that `bootstrapSource` can look
   * back up. `addon` (a pasted or pushed FS1 export) and `manual` (a build adopted from the
   * planner) never persist a `ref` at all -- `sources.ts`'s `fromAddonExport` and
   * `character.ts`'s `characterFromPlanner` both write `ref: ''` -- so navigating with
   * their `source`/`ref` landed on an empty `/sim?source=addon` with no character and no
   * message (H4, final whole-branch review). The saved result's own `request.character`
   * carries everything an FS1 code does, so this builds one and hands it to `/sim`'s
   * existing `?code=` bootstrap -- the same path "Sim this build" already uses -- instead
   * of a ref nothing wrote.
   */
  function onRerunSaved(): void {
    if (savedResult === null) return;
    const { kind, ref } = savedResult.request.source;
    const target =
      ref !== ''
        ? withSimState(defaultSimState(), { source: kind, ref })
        : withSimState(defaultSimState(), {
            // encodeFS1V2, not encodeFS1: the saved result's own gear list carries any
            // enchant or suffix (contract 10.5), and dropping them here would lose exactly
            // what characterFromFs1 now keeps.
            code: encodeFS1V2({
              dataBuild: bootstrap.treeVersion,
              classSlug: savedResult.request.character.class,
              raceSlug: savedResult.request.character.race,
              treeRanks: ranksFromTalentsString(savedResult.request.character.talents),
              gear: gearFromSlots(savedResult.request.character.gear),
              gearSlots: savedResult.request.character.gear.map((slot) => ({
                slot: slot.slot as Slot,
                itemId: slot.item_id,
                ...(slot.enchant === undefined ? {} : { enchant: slot.enchant }),
                ...(slot.suffix === undefined ? {} : { suffix: slot.suffix }),
              })),
              bags: [],
              bank: [],
              sets: [],
              loadouts: [],
              professions: [...(savedResult.request.character.professions ?? [])],
              ignored: [],
            }),
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

  // Set from `onMount`'s own `fetchMe` below, the same answer that already gates the
  // premium control -- one fetch, read by the history panel (Task 17), the landing state
  // and the source switcher's signed-in card (Task 18), rather than a signed-in flag each
  // of them would otherwise need its own copy of.
  let me = $state<Me | null>(null);
  // The history panel (Task 17), for a signed-in player on plain /sim only.
  const signedIn = $derived(me !== null);
  let historyRows = $state<SimListRow[] | null>(null);
  let historyError = $state<string | null>(null);

  async function loadHistory(): Promise<void> {
    historyError = null;
    try {
      const page = await listMySims();
      historyRows = page.rows;
    } catch (error) {
      historyRows = null;
      historyError = error instanceof Error ? error.message : simCopy.loadFailed;
    }
  }

  // Loaded lazily (Task 17): the history panel is empty weight for every signed-out
  // visitor and for /sim/<id>, where it never renders at all, and this file's compiled
  // bundle is shared by all three /sim routes.
  const simHistoryLazy = createLazyComponent(() => import('./SimHistory.svelte'));
  $effect(() => {
    if (signedIn) simHistoryLazy.load();
  });

  // Design 5.4: the finish notification. `notifier` is the real Notification API, or null
  // where the browser has none (notify.ts's own seam) -- read once, since the API itself
  // never appears mid-session. Permission is asked for only from `toggleNotify`, the
  // player's own click on the checkbox; nothing here asks on mount.
  const notifier = untrack(() => browserNotifier());
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

  // The save form under the results (Task 17). `saveOpen`/`saveTitle` are the inline
  // form; `savedUrl` is what replaces it on success, exactly as `SharePanel.svelte`'s own
  // save flow does for a build. Any new run invalidates whatever the form was showing --
  // a fresh result is a different sim to save, and an old saved link would be pointing at
  // the wrong one.
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

  onMount(() => {
    // `user.premium` on GET /v1/me, per the simulator contract -- the server lane renders
    // only once this answers true. A signed-out visitor and an unreachable API read the
    // same way here (fetchMe resolves null, or the promise rejects and is swallowed): both
    // mean "no premium control", the way Account.svelte's own `load()` already treats a
    // failed fetchMe as "not signed in" rather than an error banner.
    //
    // The same answer also gates the history panel (Task 17), the landing state and the
    // source switcher's signed-in card (Task 18): `signedIn` above is `me !== null`, and
    // the history panel's own `GET /v1/sims?mine=1` fires only then -- a signed-out visitor
    // gets no second request for a list that would come back empty anyway.
    void fetchMe()
      .then((result) => {
        me = result;
        store.setPremium(result?.user.premium === true);
        if (result !== null) void loadHistory();
      })
      .catch(() => {});
    return () => store.dispose();
  });

  // The spec support list: `/sim/specs` renders it as a grid, and `/sim` needs it too, to
  // know whether the loaded character's own spec is one the engine models at all. One
  // small, cacheable, credential-free GET, so both views share it rather than each fetching
  // their own copy.
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

  // No reactive read inside, so this fires once, on mount, the way an `onMount` fetch would
  // -- the effect form is what Task 15's brief calls for, since `/sim/specs`' own grid needs
  // exactly this same one-shot load.
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

  // False until the player explicitly asks for the switcher -- the strip's "Change source",
  // the landing state's "Sim something else", or the no-characters card's account link.
  // Design 4.6: a signed-in member with characters opens on the landing state instead of
  // the switcher, so defaulting this to `store.character === null` (as it read before the
  // landing state existed) would show the switcher on every first paint and the landing
  // state would never appear. Once a character *is* on screen, the effect below keeps this
  // false regardless, the same way it always has.
  let switcherOpen = $state(false);

  $effect(() => {
    if (store.character !== null) switcherOpen = false;
  });

  function onSignIn(): void {
    window.location.href = battlenetStartUrl(`${window.location.pathname}${window.location.search}`);
  }

  // The landing state's own busy key (Task 18): the row a pick is in flight for, so its
  // button reads "Loading…" while every other row disables rather than reads it too.
  // `store.loadStored`'s own `adopt()` is what sets `store.message` on a refusal -- the
  // "no race recorded" case sources.ts's `fromStoredCharacter` returns when the API has not
  // recorded one yet -- so the landing-state message below reads that field rather than a
  // second one this function would have to keep in step with it.
  let landingBusyKey = $state<string | null>(null);

  async function pickCharacter(path: CharacterPath): Promise<void> {
    landingBusyKey = `${path.region}/${path.ruleset}/${path.slug}`;
    await store.loadStored(path);
    landingBusyKey = null;
  }

  // The stale-engine banner and pill describe a *settled* result, not one that is being
  // re-run right now: while either lane is actively running, `store.result` still holds
  // the old, stale result (neither lane clears it until a fresh one lands), so without this
  // gate the stale banner's own "Run again" would render right alongside a live "Loading
  // engine…"/"Stop" or the server-lane's own in-flight state -- two different "run this
  // again" affordances on screen, one of them describing a run that already started.
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

{#snippet lazyFallback(lazy: LazyLoadState)}
  {#if lazy.error !== ''}
    <p class="text-muted px-[18px] text-[13px] md:px-0" role="alert" data-testid="sim-results-error">
      {lazy.error}
      <button
        type="button"
        class="text-strong ml-1 inline-flex min-h-11 items-center underline"
        onclick={() => lazy.load()}>{simCopy.tryAgain}</button
      >
    </p>
  {/if}
{/snippet}

<div class="flex flex-col gap-[22px] md:gap-8" data-testid="sim-view">
  {#if hasSavedSimId}
    <!-- /sim/<sim_id>: read-only, and not the sim page with a result in it -- no switcher,
         no settings bar, no run control. SavedSim composes its own heading. -->
    {#if savedResult !== null}
      <SavedSim result={savedResult} onrerun={onRerunSaved} />
    {:else if savedError !== null}
      <p class="text-muted px-[18px] text-[14px] md:px-0" role="alert" data-testid="sim-saved-error">
        {savedError}
      </p>
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
    <div class="flex flex-wrap items-baseline justify-end gap-x-4 gap-y-1 px-[18px] md:px-0">
      <a class="tabular text-muted font-mono text-[12px]" href="/sim/specs" data-testid="sim-engine-version"
        >{engineLabel(ENGINE_VERSION)}</a
      >
    </div>

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
      {:else if me !== null && me.characters.length > 0 && !switcherOpen}
        <!-- Design 4.6: a signed-in member sees their characters and one button each, and
             no form until they ask for one -- so this replaces the switcher entirely rather
             than sitting above it. -->
        <LandingState
          characters={me.characters}
          busyKey={landingBusyKey}
          onpick={(path) => void pickCharacter(path)}
          onother={() => (switcherOpen = true)}
        />
        {#if store.message !== null}
          <!-- The only failure a stored-character pick raises today is sources.ts's own
               "no race recorded" refusal (a combat log carries none, and the API has not
               started sending one for a stored character either) -- but whatever the
               message, the remedy is the same: the addon export is the one source that
               always carries a race, so the hint follows every refusal here rather than
               only the one the copy names. -->
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

      {#if signedIn && simHistoryLazy.current}
        <simHistoryLazy.current rows={historyRows} error={historyError} />
      {/if}

      {#if store.character !== null}
        <SettingsBar
          settings={store.settings}
          spec={store.character.spec}
          disabled={store.phase === 'running' || store.serverRunning}
          onchange={(next) => store.setSettings(next)}
        />
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
          phase={store.phase}
          estimate={store.estimate}
          iterationsDone={store.iterationsDone}
          iterationsTotal={store.iterationsTotal}
          precisionId={store.precisionId}
          relativeError={store.relativeError}
          lane={store.lane}
          premium={store.premium}
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
              {@render lazyFallback(compareViewLazy)}
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
            />
          {:else}
            {@render lazyFallback(simResultsLazy)}
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
            <input
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
              class="border-line-warm-strong rounded-control bg-card-top text-strong label min-h-11 border px-5 disabled:opacity-50"
              disabled={saving}
              onclick={() => void confirmSave()}
              data-testid="sim-save-confirm"
            >
              {saving ? simCopy.savingAction : simCopy.saveAction}
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
              title={store.result !== null && !canSave ? simCopy.saveAbortedDisabled : undefined}
              onclick={openSaveForm}
              data-testid="sim-save-open"
            >
              {simCopy.saveThisSim}
            </button>
          {/if}
        </div>
      {:else}
        <p class="text-muted px-[18px] text-[14px] md:px-0" data-testid="sim-empty">
          {simCopy.emptyPrompt}
        </p>
      {/if}
    {/if}
  {/if}
</div>
