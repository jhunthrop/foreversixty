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
  import { battlenetStartUrl, fetchMe } from '../../lib/account/api';
  import { createLazyComponent, type LazyLoadState } from '../../lib/report/lazy-component.svelte';
  import { fetchReportMeta, fetchSummary } from '../../lib/report/load';
  import type { Summary } from '../../lib/report/types';
  import { fetchSim, fetchSpecs, listMySims } from '../../lib/sim/api';
  import { compareSummaries } from '../../lib/sim/compare';
  import { simCopy } from '../../lib/sim/copy';
  import { settingsLabel } from '../../lib/sim/settings';
  import { parseFightRef } from '../../lib/sim/sources';
  import { mergeSpecRows } from '../../lib/sim/spec-state';
  import { createSimStore } from '../../lib/sim/store.svelte';
  import { defaultSimState, parseSimState, simSearch, withSimState } from '../../lib/sim/url';
  import { ENGINE_VERSION, engineLabel, isStale } from '../../lib/sim/version';
  import type { SimListRow, SimResult, SpecFidelity } from '../../lib/sim/types';
  import CharacterStrip from './CharacterStrip.svelte';
  import RunControl from './RunControl.svelte';
  import SavedSim from './SavedSim.svelte';
  import SettingsBar from './SettingsBar.svelte';
  import SourceSwitcher from './SourceSwitcher.svelte';
  import SpecCard from './SpecCard.svelte';
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
    const { source, ref, code, mode } = parseSimState(search);
    // The mount element the shell can stamp a build id onto, the way planner-island.ts
    // reads `data-tree-version` off its own mount -- no /sim page stamps one yet, so this
    // falls back to the site's active build rather than an empty string no fetch would
    // resolve.
    const mount = document.getElementById('sim');
    const treeVersion = mount?.dataset.treeVersion ?? activeBuild.build;
    // specs.astro stamps `data-sim-view="specs"`; sim.astro and [id].astro stamp neither,
    // so an absent or unrecognised value reads as the ordinary simulator.
    const view: 'sim' | 'specs' = mount?.dataset.simView === 'specs' ? 'specs' : 'sim';
    return { treeVersion, source, ref, code, mode, view };
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

  if (hasSavedSimId && savedResult === null && simId !== '') {
    void fetchSim(simId)
      .then((result) => (savedResult = result))
      .catch((error) => {
        savedError = error instanceof Error ? error.message : simCopy.loadFailed;
      });
  }

  /**
   * "Run this yourself", the stale-result remedy and the ordinary way off a saved page:
   * opens /sim with the stored request's own source and ref, so the player can change
   * something and run it themselves. It is a callback rather than an `<a href>` SavedSim
   * builds itself so the URL is built with the same `sim/url.ts` vocabulary this file
   * already owns for its own bootstrap, in one place. It is never triggered automatically
   * -- only this handler, from the player's own click.
   */
  function onRerunSaved(): void {
    if (savedResult === null) return;
    const target = withSimState(defaultSimState(), {
      source: savedResult.request.source.kind,
      ref: savedResult.request.source.ref,
    });
    window.location.href = `/sim${simSearch(target)}`;
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

  // The history panel (Task 17), for a signed-in player on plain /sim only -- set from
  // `onMount`'s own `fetchMe` below, the same answer that already gates the premium
  // control, so there is one signed-in check for the page rather than two.
  let signedIn = $state(false);
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
    saveTitle = settingsLabel(store.settings);
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
    // The same answer also gates the history panel (Task 17): `signedIn` is `me !== null`,
    // and the panel's own `GET /v1/sims?mine=1` fires only then -- a signed-out visitor
    // gets no second request for a list that would come back empty anyway.
    void fetchMe()
      .then((me) => {
        store.setPremium(me?.user.premium ?? false);
        signedIn = me !== null;
        if (signedIn) void loadHistory();
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
  const specUnsupported = $derived(characterSpecRow?.state === 'unsupported');

  // Open until a character is on screen, or reopened by "Change source". The store's own
  // URL bootstrap (above) can land a character before this component's first render, so
  // this reads `store.character` rather than defaulting to a fixed value.
  let switcherOpen = $state(store.character === null);

  $effect(() => {
    if (store.character !== null) switcherOpen = false;
  });

  function onSignIn(): void {
    window.location.href = battlenetStartUrl(`${window.location.pathname}${window.location.search}`);
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
      <button type="button" class="text-strong ml-1 underline" onclick={() => lazy.load()}>Try again</button>
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
    {/if}
  {:else}
    <div class="flex flex-wrap items-baseline gap-x-4 gap-y-1 px-[18px] md:px-0">
      <h1 class="section-title">Simulator</h1>
      <a
        class="tabular text-muted ml-auto font-mono text-[12px]"
        href="/sim/specs"
        data-testid="sim-engine-version">{engineLabel(ENGINE_VERSION)}</a
      >
    </div>

    {#if bootstrap.view === 'specs'}
      <p class="text-muted px-[18px] text-[14px] md:px-0" data-testid="specs-intro">{simCopy.specsIntro}</p>
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
      {:else}
        <SourceSwitcher
          busy={store.phase === 'loading-character'}
          message={store.message}
          signedIn={false}
          onaddon={(code) => void store.loadAddon(code)}
          onbuild={(id) => void store.loadBuild(id)}
          onfight={(ref) => void store.loadFight(ref)}
          onsignin={onSignIn}
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
        {#if specUnsupported && characterSpecRow !== null}
          <!-- Instead of the run control, the sentence and the results -- not above them.
               The settings bar above still says what would be simulated; this says why it
               cannot be, with the same card /sim/specs shows for this spec. -->
          <p class="text-muted px-[18px] text-[14px] md:px-0" data-testid="spec-unsupported-lead">
            {simCopy.specUnsupportedLead}
          </p>
          <div class="px-[18px] md:px-0">
            <SpecCard row={characterSpecRow} actionNames={store.actionNames} compact />
          </div>
        {:else}
          <RunControl
            phase={store.phase}
            estimate={store.estimate}
            iterationsDone={store.iterationsDone}
            iterationsTotal={store.iterationsTotal}
            precision={store.precision}
            premium={store.premium}
            message={store.message}
            detail={store.detail}
            racePending={store.needsRace}
            {staleVersion}
            serverRunning={store.serverRunning}
            onrun={() => void store.run()}
            onstop={() => store.stop()}
            onprecision={(value) => store.setPrecision(value)}
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
              />
            {:else}
              {@render lazyFallback(simResultsLazy)}
            {/if}
          {/if}

          <!-- The save form (Task 17): disabled until there is a result, an inline
               title field pre-filled with the settings clause rather than a dialog, and
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
        {/if}
      {:else}
        <p class="text-muted px-[18px] text-[14px] md:px-0" data-testid="sim-empty">
          {simCopy.emptyPrompt}
        </p>
      {/if}
    {/if}
  {/if}
</div>
