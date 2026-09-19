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
  import { simCopy } from '../../lib/sim/copy';
  import { createSimStore } from '../../lib/sim/store.svelte';
  import { parseSimState } from '../../lib/sim/url';
  import { ENGINE_VERSION, engineLabel, isStale } from '../../lib/sim/version';
  import type { SimResult } from '../../lib/sim/types';
  import CharacterStrip from './CharacterStrip.svelte';
  import RunControl from './RunControl.svelte';
  import SettingsBar from './SettingsBar.svelte';
  import SourceSwitcher from './SourceSwitcher.svelte';

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
    const { source, ref, code } = parseSimState(search);
    // The mount element the shell can stamp a build id onto, the way planner-island.ts
    // reads `data-tree-version` off its own mount -- no /sim page stamps one yet, so this
    // falls back to the site's active build rather than an empty string no fetch would
    // resolve.
    const treeVersion = document.getElementById('sim')?.dataset.treeVersion ?? activeBuild.build;
    return { treeVersion, source, ref, code };
  });

  const store = untrack(() =>
    createSimStore({
      treeVersion: bootstrap.treeVersion,
      source: bootstrap.source,
      ref: bootstrap.ref,
      code: bootstrap.code,
    }),
  );

  onMount(() => {
    // `user.premium` on GET /v1/me, per the simulator contract -- the server lane renders
    // only once this answers true. A signed-out visitor and an unreachable API read the
    // same way here (fetchMe resolves null, or the promise rejects and is swallowed): both
    // mean "no premium control", the way Account.svelte's own `load()` already treats a
    // failed fetchMe as "not signed in" rather than an error banner.
    void fetchMe()
      .then((me) => store.setPremium(me?.user.premium ?? false))
      .catch(() => {});
    return () => store.dispose();
  });

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
  <div class="flex flex-wrap items-baseline gap-x-4 gap-y-1 px-[18px] md:px-0">
    <h1 class="section-title">Simulator</h1>
    <a
      class="tabular text-muted ml-auto font-mono text-[12px]"
      href="/sim/specs"
      data-testid="sim-engine-version">{engineLabel(ENGINE_VERSION)}</a
    >
  </div>

  {#if !hasSavedSimId}
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

    {#if store.character !== null}
      <SettingsBar
        settings={store.settings}
        spec={store.character.spec}
        disabled={store.phase === 'running' || store.serverRunning}
        onchange={(next) => store.setSettings(next)}
      />
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
      {#if store.result !== null}
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
    {:else}
      <p class="text-muted px-[18px] text-[14px] md:px-0" data-testid="sim-empty">
        {simCopy.emptyPrompt}
      </p>
    {/if}
  {/if}
</div>
