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
  import { clearCurrent, readCurrent, type CurrentCharacter } from '../../../lib/current-character';
  import { CHIP_HEIGHT } from '../../../lib/current-character-layout';
  import { createLazyComponent, type LazyLoadState } from '../../../lib/report/lazy-component.svelte';
  import { TOOL_SKELETONS } from '../../../lib/sim/bulk-skeleton';
  import { createBulkStore, type SimTool } from '../../../lib/sim/bulk-store.svelte';
  import type { Origin } from '../../../lib/sim/candidates';
  import { bulkCopy, simCopy } from '../../../lib/sim/copy';
  import { syncTabHrefs } from '../../../lib/sim/tabs';
  import {
    decideToolsBootstrap,
    settleToolsRestore,
    sourceIdForInstance,
    type ToolsBootstrapDecision,
  } from '../../../lib/sim/tools-bootstrap';
  import { parseSimState } from '../../../lib/sim/url';
  import CurrentCharacterChip from '../../CurrentCharacterChip.svelte';
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
  // pointer rather than the URL, AND that load actually produced a character (fix round 1,
  // Task 4's review, Important: a dead pointer must not claim "restored") -- passed to the
  // chip below, which shows "Restored your last character" beside the label only then.
  let restored = $state(false);
  // The chip's own prop (fix round 1, Critical: the chip, not an ad-hoc row here, is what
  // shows the current-character pointer -- CurrentCharacterChip.svelte's own header
  // comment). Refreshed from storage inside the tab-sync effect below, the same moment
  // every loader has already written it (sources.ts's own `recordCurrentCharacter`, called
  // before `adopt()` assigns `character`).
  let pointer = $state<CurrentCharacter | null>(null);

  $effect(() => {
    if (store.character !== null) switcherOpen = false;
  });

  // Keeps the loaded character on every tab in SimTabs.astro's strip (task-1-brief.md),
  // the same rewrite SimView.svelte's own effect performs for /sim, /sim/[id] and
  // /sim/specs: whenever this island's own character changes, every tab's href is
  // rewritten to carry the same query its own destination can actually bootstrap from --
  // `?source=&ref=`, or `store.characterCode` (bulk-store.svelte.ts) as the fallback
  // `?code=` every tab now reads (Task 4; `SIM_TABS`' own `supportsCode` is `true` on all
  // six as of the current-character spec, 2026-09-21). Also refreshes `pointer` (fix round
  // 1) from the same trigger: every loader writes the stored pointer before `character` is
  // assigned, so by the time this effect reacts to that change, `readCurrent()` already
  // reads what the load just wrote.
  $effect(() => {
    const source = store.character === null ? null : store.character.source;
    // `store.characterCode` costs a talent-index rebuild (bulk-store-request.ts's own
    // `characterSpecOrNull`), read only when `source.ref` is empty and there is actually a
    // fallback to try -- the same guard SimView.svelte's own effect applies.
    const fallbackCode = source !== null && source.ref === '' ? store.characterCode : null;
    syncTabHrefs(source, fallbackCode);
    pointer = readCurrent();
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
   *
   * Never runs on `/sim/drops` (fix round 1, Minor, Task 4's review): that page's own
   * instance `$effect`, below, rebuilds `rows` wholesale from `pickedBosses`
   * (`rowsFromPicks`), which would wipe a row this effect just added.
   */
  $effect(() => {
    if (pinApplied || store.phase !== 'idle' || store.character === null || tool === 'drops') return;
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

  /** `decision.kind`'s own loader, or null for `'none'` -- the one place that maps a
   *  bootstrap decision onto the store's loaders, so `bootstrapCharacter` below reads as
   *  "decide, load, settle" rather than a second copy of this switch. */
  function startLoad(decision: ToolsBootstrapDecision): Promise<void> | null {
    if (decision.kind === 'code') return store.loadCode(decision.code);
    if (decision.kind === 'addon') return store.loadAddon(decision.code);
    if (decision.kind === 'build') return store.loadBuild(decision.id);
    if (decision.kind === 'fight') return store.loadFight(decision.ref);
    if (decision.kind === 'stored') return store.loadStored(decision.path);
    return null;
  }

  /**
   * The URL's own bootstrap wins over the stored current-character pointer, which wins
   * over nothing (Task 4, current-character spec section 1) -- `decideToolsBootstrap` is
   * the one place that precedence lives, so it is testable without mounting this island.
   *
   * Fix round 1, Task 4's review (Important): the load this kicks off is awaited, not
   * `void`-ed, so `restored` is only ever set once it has actually settled --
   * `settleToolsRestore` (tools-bootstrap.ts) decides what "settled" means: a restore that
   * produced no character forgets the pointer and clears the store's own refusal message,
   * rather than showing "Restored your last character" beside an error for a load the
   * player never asked for. A URL-driven load keeps today's behaviour on failure.
   */
  async function bootstrapCharacter(): Promise<void> {
    const decision = decideToolsBootstrap(
      { code: bootstrap.code, source: bootstrap.source, ref: bootstrap.ref },
      readCurrent(),
    );
    const load = startLoad(decision);
    if (load !== null) await load;
    const outcome = settleToolsRestore(decision.restored, store.character !== null);
    restored = outcome.restored;
    if (outcome.clearPointer) clearCurrent();
    if (outcome.clearMessage) store.setMessage(null);
  }

  onMount(() => {
    void fetchMe()
      .then((result) => {
        me = result;
        store.setPremium(result?.user.premium === true);
      })
      .catch(() => {});
    void store.loadSpecs();
    void bootstrapCharacter();
    return () => store.dispose();
  });

  function onSignIn(): void {
    window.location.href = battlenetStartUrl(`${window.location.pathname}${window.location.search}`);
  }

  function onForgetPointer(): void {
    clearCurrent();
    pointer = null;
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
  <!-- Fix round 1, Task 4's review (Critical): a reserved, always-present slot -- never
       conditionally rendered -- so its height never changes and nothing below it ever
       shifts, whether the chip has a character to show or not. bulk-skeleton.ts's own
       `chipSlot` reserves the identical band before hydration. -->
  <div class={CHIP_HEIGHT} data-testid="sim-chip-slot">
    <CurrentCharacterChip current={pointer} {restored} hasOwnPasteBox onforget={onForgetPointer} />
  </div>
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
