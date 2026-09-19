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
  import { battlenetStartUrl } from '../../lib/account/api';
  import { simCopy } from '../../lib/sim/copy';
  import { createSimStore } from '../../lib/sim/store.svelte';
  import { parseSimState } from '../../lib/sim/url';
  import { ENGINE_VERSION, engineLabel } from '../../lib/sim/version';
  import type { SimResult } from '../../lib/sim/types';
  import CharacterStrip from './CharacterStrip.svelte';
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
    const { source, ref } = parseSimState(search);
    // `code` is not yet part of SimState's own vocabulary (url.ts), so it is read straight
    // off the query string here rather than by editing that module for one field -- the
    // decoder it reaches (decodeFS1) already bounds and validates it.
    const code = new URLSearchParams(search).get('code') ?? undefined;
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
</script>

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

    {#if store.character === null}
      <p class="text-muted px-[18px] text-[14px] md:px-0" data-testid="sim-empty">
        {simCopy.emptyPrompt}
      </p>
    {/if}
  {/if}
</div>
