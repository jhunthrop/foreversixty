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
  import { createLazyComponent, type LazyLoadState } from '../../../lib/report/lazy-component.svelte';
  import { TOOL_SKELETONS } from '../../../lib/sim/bulk-skeleton';
  import { createBulkStore, type SimTool } from '../../../lib/sim/bulk-store.svelte';
  import type { Origin } from '../../../lib/sim/candidates';
  import { bulkCopy, simCopy } from '../../../lib/sim/copy';
  import { parseSimState } from '../../../lib/sim/url';
  import CharacterStrip from '../CharacterStrip.svelte';
  import SourceSwitcher from '../SourceSwitcher.svelte';

  let { tool }: { tool: SimTool } = $props();

  const bootstrap = untrack(() => {
    const { source, ref } = parseSimState(window.location.search);
    const params = new URLSearchParams(window.location.search);
    const mount = document.getElementById('sim-tools');
    return {
      source,
      ref,
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

  $effect(() => {
    if (store.character !== null) switcherOpen = false;
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

  onMount(() => {
    void fetchMe()
      .then((result) => {
        me = result;
        store.setPremium(result?.user.premium === true);
      })
      .catch(() => {});
    void store.loadSpecs();
    // The URL's own bootstrap, once, here rather than in an effect: a "sim this build" link
    // arrives as ?source=&ref=.
    if (bootstrap.source !== '' && bootstrap.ref !== '') {
      if (bootstrap.source === 'addon') void store.loadAddon(bootstrap.ref);
      else if (bootstrap.source === 'build') void store.loadBuild(bootstrap.ref);
      else if (bootstrap.source === 'fight') void store.loadFight(bootstrap.ref);
    }
    return () => store.dispose();
  });

  function onSignIn(): void {
    window.location.href = battlenetStartUrl(`${window.location.pathname}${window.location.search}`);
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
