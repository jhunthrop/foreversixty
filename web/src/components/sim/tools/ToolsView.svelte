<!-- web/src/components/sim/tools/ToolsView.svelte -->
<!-- The tools island's root: the character strip, the source switcher and the one tool
     view this page is. Every tool view is a lazy chunk, so /sim/weights downloads neither
     the slot grid nor the source picker.

     There is no <h1> here: each .astro shell carries its own, in static HTML, ahead of this
     island -- an island-mounted heading is invisible to Lighthouse's LCP measurement, which
     is the same reason SimView.svelte has none. -->
<script lang="ts">
  import { onMount, untrack } from 'svelte';
  import activeBuild from '../../../data/active-build.json';
  import { battlenetStartUrl, fetchMe, type Me } from '../../../lib/account/api';
  import { clearCurrent, readCurrent } from '../../../lib/current-character';
  import { currentCharacterCopy } from '../../../lib/current-character-copy';
  import { rowLink } from '../../../lib/report/format';
  import { createLazyComponent, type LazyLoadState } from '../../../lib/report/lazy-component.svelte';
  import { TOOL_SKELETONS } from '../../../lib/sim/bulk-skeleton';
  import { createBulkStore, type SimTool } from '../../../lib/sim/bulk-store.svelte';
  import type { Origin } from '../../../lib/sim/candidates';
  import { bulkCopy, simCopy } from '../../../lib/sim/copy';
  import { syncTabHrefs } from '../../../lib/sim/tabs';
  import { decideToolsBootstrap, sourceIdForInstance } from '../../../lib/sim/tools-bootstrap';
  import { parseSimState } from '../../../lib/sim/url';
  import CharacterStrip from '../CharacterStrip.svelte';
  import SourceSwitcher from '../SourceSwitcher.svelte';

  let { tool }: { tool: SimTool } = $props();

  const bootstrap = untrack(() => {
    const { source, ref, code } = parseSimState(window.location.search);
    const params = new URLSearchParams(window.location.search);
    const mount = document.getElementById('sim-tools');
    return {
      source,
      ref,
      code,
      // /sim/drops's own preselect (Task 4): a Droptimizer's zone slug, matched against a
      // LootSource.id once the loot file has loaded (the instance $effect below).
      instance: params.get('instance') ?? '',
      treeVersion: mount?.dataset.treeVersion ?? activeBuild.build,
      // A Droptimizer pin (Task 18): the item, and the `drop:<source-id>` origin and boss
      // name it carried in, so the row lands on Top Gear with its real provenance rather
      // than a plain search hit.
      pin: params.get('pin') ?? '',
      pinOrigin: params.get('pinOrigin') ?? '',
      pinName: params.get('pinName') ?? '',
    };
  });

  const store = untrack(() =>
    createBulkStore({
      tool,
      treeVersion: bootstrap.treeVersion,
      source: bootstrap.source,
      ref: bootstrap.ref,
      hardwareConcurrency: navigator.hardwareConcurrency,
    }),
  );

  let me = $state<Me | null>(null);
  let switcherOpen = $state(false);
  let pinApplied = false;
  let instanceApplied = false;
  // True once the load this mount kicked off came from the stored current-character
  // pointer rather than the URL (Task 4, current-character spec section 1) -- drives the
  // "Restored your last character. Forget" line above the strip.
  let restored = $state(false);

  $effect(() => {
    if (store.character !== null) switcherOpen = false;
  });

  // Keeps the loaded character on every tab in SimTabs.astro's strip (task-1-brief.md),
  // the same rewrite SimView.svelte's own effect performs for /sim, /sim/[id] and
  // /sim/specs: whenever this island's own character changes, every tab's href is
  // rewritten to carry the same query its own destination can actually bootstrap from --
  // `?source=&ref=`, or `store.characterCode` (bulk-store.svelte.ts) as the fallback
  // `?code=` every tab now reads (Task 4; `SIM_TABS`' own `supportsCode` is `true` on all
  // six as of the current-character spec, 2026-09-21).
  $effect(() => {
    const source = store.character === null ? null : store.character.source;
    // `store.characterCode` costs a talent-index rebuild (bulk-store-request.ts's own
    // `characterSpecOrNull`), read only when `source.ref` is empty and there is actually a
    // fallback to try -- the same guard SimView.svelte's own effect applies.
    const fallbackCode = source !== null && source.ref === '' ? store.characterCode : null;
    syncTabHrefs(source, fallbackCode);
  });

  /**
   * A Droptimizer pin (Task 18), applied the first time a character is truly settled on
   * screen -- whether it arrived with this navigation (`?source=&ref=` alongside `?pin=`)
   * or the player picked one afterwards through the switcher.
   *
   * Gated on `phase === 'idle'`, not `character !== null` alone (fix round 1, Finding 2):
   * `adopt()` sets `character` well before `items` is populated -- `phase` is deliberately
   * held at `'loading-character'` until `loadDataFor`/`seedRows` finish (bulk-store.svelte.ts's
   * own comment on that exact race). Gating on `character` alone let this effect fire into
   * an empty item map, where `addSearchItem` would silently find nothing and no-op -- exactly
   * the window `adopt()` was written to close for `seedRows`, reopened here for the pin.
   * `phase === 'idle'` is the same signal `adopt()` itself waits for before considering a
   * load "settled".
   */
  $effect(() => {
    if (pinApplied || store.phase !== 'idle' || store.character === null) return;
    pinApplied = true;
    if (bootstrap.pin === '') return;
    const itemId = Number.parseInt(bootstrap.pin, 10);
    if (!Number.isInteger(itemId) || itemId <= 0) return;
    const origin: Origin = bootstrap.pinOrigin.startsWith('drop:')
      ? (bootstrap.pinOrigin as Origin)
      : 'search';
    store.addSearchItem(itemId, origin, bootstrap.pinName);
  });

  /**
   * `/sim/drops`'s own preselect (Task 4): a Droptimizer link's `?instance=` zone slug,
   * matched against the loaded loot file's own sources once it has settled. Modelled on
   * the pin effect above -- same `phase === 'idle'` gate, same apply-once guard -- so a
   * player who switches sources afterward is never re-forced back onto this pick.
   */
  $effect(() => {
    if (instanceApplied || store.phase !== 'idle' || store.character === null) return;
    instanceApplied = true;
    if (tool !== 'drops' || bootstrap.instance === '') return;
    const matchId = sourceIdForInstance(store.loot.sources, bootstrap.instance);
    if (matchId !== null) store.toggleSource(matchId);
  });

  onMount(() => {
    void fetchMe()
      .then((result) => {
        me = result;
        store.setPremium(result?.user.premium === true);
      })
      .catch(() => {});
    void store.loadSpecs();
    // The URL's own bootstrap wins over the stored current-character pointer, which wins
    // over nothing (Task 4, current-character spec section 1) -- decideToolsBootstrap is
    // the one place that precedence lives, so it is testable without mounting this island.
    const decision = decideToolsBootstrap(
      { code: bootstrap.code, source: bootstrap.source, ref: bootstrap.ref },
      readCurrent(),
    );
    restored = decision.restored;
    if (decision.kind === 'code') void store.loadCode(decision.code);
    else if (decision.kind === 'addon') void store.loadAddon(decision.code);
    else if (decision.kind === 'build') void store.loadBuild(decision.id);
    else if (decision.kind === 'fight') void store.loadFight(decision.ref);
    else if (decision.kind === 'stored') void store.loadStored(decision.path);
    return () => store.dispose();
  });

  function onSignIn(): void {
    window.location.href = battlenetStartUrl(`${window.location.pathname}${window.location.search}`);
  }

  function onForgetRestored(): void {
    clearCurrent();
    restored = false;
  }

  // One chunk per tool, resolved from a closed map: `tool` is validated against TOOLS
  // before it reaches here, and a map rather than a template literal keeps the bundler's
  // own analysis exact.
  const views = {
    gear: () => import('./TopGear.svelte'),
    talents: () => import('./TopGear.svelte'),
    drops: () => import('./Droptimizer.svelte'),
    weights: () => import('./StatWeights.svelte'),
  } as const;

  // `tool` never changes after mount -- a different tool page is a different navigation,
  // not a prop update -- so this reads it once rather than as a reactive dependency, the
  // same one-shot pattern `bootstrap` and `store` above use.
  const viewLazy = createLazyComponent(untrack(() => views[tool]));
  viewLazy.load();
</script>

{#snippet lazyFallback(lazy: LazyLoadState)}
  {#if lazy.error !== ''}
    <p class="text-muted px-[18px] text-[13px] md:px-0" role="alert" data-testid="sim-tool-error">
      {lazy.error}
      <button
        type="button"
        class="text-strong ml-1 inline-flex min-h-11 items-center underline"
        onclick={() => lazy.load()}>{simCopy.tryAgain}</button
      >
    </p>
  {:else}
    <!-- eslint-disable-next-line svelte/no-at-html-tags -->
    {@html TOOL_SKELETONS[tool]}
  {/if}
{/snippet}

<div class="flex flex-col gap-[22px] md:gap-8" data-testid="sim-tools-view">
  {#if restored}
    <p
      class="text-muted flex min-h-11 items-center gap-3 px-[18px] text-[13px] md:px-0"
      data-testid="sim-restored-note"
    >
      {currentCharacterCopy.restoredNote}
      <button
        type="button"
        class={`${rowLink} text-nav`}
        onclick={onForgetRestored}
        data-testid="sim-restored-forget"
      >
        {currentCharacterCopy.forget}
      </button>
    </p>
  {/if}
  {#if store.character !== null && !switcherOpen}
    <CharacterStrip
      character={store.character}
      items={store.items}
      races={[]}
      onchange={() => (switcherOpen = true)}
      onrace={(slug) => store.setRace(slug)}
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

  {#if store.character !== null || tool === 'weights'}
    <!-- Design 7: /sim/weights opens with the caution before any character is loaded --
         the whole reason the page exists is a warning, so it cannot wait behind a load.
         The other three tools are unaffected: they still need a character before their
         view (a slot grid, a boss picker) means anything. -->
    {#if viewLazy.current}
      <viewLazy.current {store} {me} />
    {:else}
      {@render lazyFallback(viewLazy)}
    {/if}
  {:else}
    <p class="text-muted px-[18px] text-[14px] md:px-0" data-testid="sim-tools-empty">
      {bulkCopy.needCharacter}
    </p>
  {/if}
</div>
