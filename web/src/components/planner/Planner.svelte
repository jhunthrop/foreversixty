<!-- web/src/components/planner/Planner.svelte -->
<!-- The planner island. Mounted two ways: by src/pages/planner.astro with client:load, and
     by src/planner-island.ts on the API-rendered /b/:id page, where `record` is supplied and
     the build starts read-only. Desktop lays the trees out side by side with the gear panel
     under them; phone shows one panel at a time behind a tab switcher (Tasks 9 and 17). -->
<script lang="ts">
  import { untrack } from 'svelte';
  import type { WeightsFile } from '../../lib/addon/score';
  import activeBuild from '../../data/active-build.json';
  import { clearCurrent, readCurrent, type CurrentCharacter } from '../../lib/current-character';
  import { CHIP_HEIGHT } from '../../lib/current-character-layout';
  import { DEFAULT_CLASS_SLUG } from '../../lib/planner/config';
  import {
    decidePlannerLoad,
    isBarePlannerUrl,
    plannerAddonCode,
    writePlannerPointer,
  } from '../../lib/planner/current-character-planner';
  import { ranksByTalent } from '../../lib/planner/derive';
  import { encodeFS1, orderFromRanks } from '../../lib/planner/fs1';
  import { createLiveDps } from '../../lib/planner/live-dps.svelte';
  import { isConstrainedDevice, liveGate } from '../../lib/planner/live-gate';
  import {
    DATA_LOAD_FAILED,
    DataLoadError,
    loadItems,
    loadReference,
    loadSets,
    loadTalents,
    loadWeights,
  } from '../../lib/planner/load';
  import type { TalentIndex } from '../../lib/planner/rules';
  import { createPlannerStore } from '../../lib/planner/store.svelte';
  import { SECONDARY_BUTTON } from '../../lib/planner/styles';
  import { treeSourceNotice } from '../../lib/planner/tree-source';
  import type { BuildRecord, Gear, TalentFile } from '../../lib/planner/types';
  import { characterFromPlanner } from '../../lib/sim/character';
  import { defaultSimState, simSearch, withSimState } from '../../lib/sim/url';
  import CurrentCharacterChip from '../CurrentCharacterChip.svelte';
  import LoadError from '../ui/LoadError.svelte';
  import Skeleton from '../ui/Skeleton.svelte';
  import GearPanel from './GearPanel.svelte';
  import ImportBox from './ImportBox.svelte';
  import OrderStrip from './OrderStrip.svelte';
  import SharePanel from './SharePanel.svelte';
  import SummaryBar from './SummaryBar.svelte';
  import TreeGrid from './TreeGrid.svelte';

  let {
    treeVersion,
    classSlug = DEFAULT_CLASS_SLUG,
    raceSlug,
    record = null,
    gear,
    oncode,
    standalone = true,
  }: {
    treeVersion: string;
    classSlug?: string;
    raceSlug?: string;
    record?: BuildRecord | null;
    /**
     * Seeds the store's gear when there is no `record` (dps D39/D40) -- Top Gear's inline
     * "add a build" (TalentCandidates.svelte) passes the loaded character's gear so its live
     * DPS card sims the same equipment the comparison table does. `record?.gear` always wins
     * when a record is present: `/b/:id` and `/planner` must stay byte-identical to today.
     */
    gear?: Gear;
    /**
     * Called with the build's own FS1 code whenever it changes. Top Gear's "add a build"
     * (Task 14's TalentCandidates) mounts this component inline and reads the code back
     * through it; /planner and /b/:id pass nothing and the callback never fires.
     */
    oncode?: (code: string) => void;
    /** False only for Top Gear's inline "add a build" (Task 10): no chip, no restore, no write. */
    standalone?: boolean;
  } = $props();

  // The page is static, so ?class= and ?race= can only be read in the browser. A record
  // (from /b/:id) wins over the query string: that build already names its class and race.
  function fromQuery(name: string): string | undefined {
    if (typeof window === 'undefined') return undefined;
    return new URLSearchParams(window.location.search).get(name) ?? undefined;
  }

  // Current-character pointer (Task 10, spec section 1) -- decision in current-character-
  // planner.ts. `codeParam` is `?code=` itself, or a restored pointer's code, same path.
  const plannerLoad = untrack(() => {
    const urlCode = record ? null : (fromQuery('code') ?? null);
    const search = typeof window === 'undefined' ? '' : window.location.search;
    return decidePlannerLoad(urlCode, isBarePlannerUrl(search), record !== null, standalone, readCurrent());
  });
  const { codeParam, decoded } = plannerLoad;
  let restored = $state(plannerLoad.restored);
  let pointer = $state<CurrentCharacter | null>(plannerLoad.pointer);
  // A dead restored pointer is forgotten, not shown as an error. Runs once, on mount.
  $effect(() => {
    if (plannerLoad.deadPointer) clearCurrent();
  });

  function onForgetPointer(): void {
    clearCurrent();
    pointer = null;
    restored = false;
  }

  // A parsed count inside one of these renders in a tabular, monospace span, the same as every
  // other number the planner shows (OrderStrip, SummaryBar, GearPanel, TalentCell, ItemPicker).
  // This stays a small discriminated union rather than a plain string so the template can wrap
  // just the digits with a real element -- `decoded.message` (current-character-planner.ts's
  // own `decodeFS1` call, above) stays flat text (its exact wording is pinned by fs1.test.ts
  // and is the shape Task 14 reads), so the split happens once, here, rather than by
  // reformatting arbitrary text at render time.
  type CodeNote =
    | { kind: 'message'; text: string }
    | { kind: 'tree-count'; got: string; want: string }
    | { kind: 'reconstructed'; dropped: number; gearOnly: boolean };

  /** Set when the build came in as an FS1 code: either why it could not be read, or that its
   *  order is a reconstruction (nothing in the game records the order points were spent in). */
  let codeNote = $state<CodeNote | null>(null);

  // Matches the one message `decodeFS1` returns that carries two counts, so the digits can be
  // pulled out and wrapped here without `decodeFS1` itself having to return anything but a
  // flat, tested string.
  const TREE_COUNT_MESSAGE = /^That code has (\d+) talent trees; a build has (\d+)\.$/;

  function noteForMessage(message: string): CodeNote {
    const match = TREE_COUNT_MESSAGE.exec(message);
    return match ? { kind: 'tree-count', got: match[1], want: match[2] } : { kind: 'message', text: message };
  }

  // A code has exactly one turn on the class it names, tracked by two things together rather
  // than one flag armed early:
  //
  //  - Every `load()` run only ever considers applying the code when *this run's own class*
  //    matches the class the code names. A run for any other class leaves it alone no matter
  //    how the timing between two `load()` calls falls out, which is what stops a class switch
  //    (immediate, or after a failed load) from ever applying a code meant for a different
  //    class -- `orderFromRanks` only knows tab positions, not which class they belong to, so
  //    it would apply them onto the wrong tree without complaint if it were reachable at all.
  //
  //  - `codeApplied` gates *reuse within that one matching class*. It is armed only once the
  //    code has genuinely had its turn: applied successfully, or permanently refused because
  //    the class it names has no talent data (retrying that can never succeed). It is left
  //    unarmed by any other failure on that same class -- a network blip on the reference
  //    files, a 5xx on talents, sets or items -- so a same-class Retry still gets to apply the
  //    code. Arming it unconditionally on the first attempt, whatever the failure, would drop
  //    a perfectly good code the moment an unrelated transient error hit first and leave no
  //    trace once the retry quietly succeeded onto an empty default build.
  let codeApplied = false;

  // The store is seeded once, from the props as they arrive. `untrack` says so: without it
  // the compiler warns that these reads only capture the initial value, which is the point --
  // after mount the store owns the class, the race and the order, not the props.
  const store = untrack(() =>
    createPlannerStore({
      treeVersion: record?.tree_version ?? treeVersion,
      classSlug: record
        ? classSlug
        : decoded?.ok
          ? decoded.build.classSlug
          : (fromQuery('class') ?? classSlug),
      raceSlug: record
        ? (raceSlug ?? '')
        : decoded?.ok
          ? decoded.build.raceSlug
          : (fromQuery('race') ?? raceSlug ?? ''),
      order: record?.point_order,
      gear: record?.gear ?? gear,
      title: record?.title,
      sourceId: record?.id ?? null,
      readOnly: record !== null,
    }),
  );

  // The chip's "Copy addon code" link; shared with SharePanel's own button (Task 10).
  const addonCode = $derived(plannerAddonCode(store));

  // Every load source writes the pointer through here (Task 10); writePlannerPointer is a
  // no-op inline or before talent data has loaded.
  function writePointer(source: 'code' | 'addon' | 'build', ref: string, cls: string, title?: string): void {
    pointer = writePlannerPointer(store, standalone, source, ref, cls, title);
  }

  // The live DPS estimate (Task 20). Created once -- createLiveDps holds no pool until the
  // first request, so this costs nothing on mount and does not touch engine.ts until a talent
  // or slot actually changes. `untrack` for the same reason the store above needs it: the
  // factory call itself reads nothing reactive, and without it the compiler would warn that a
  // value read once here looks like a dependency.
  const live = untrack(() => createLiveDps());

  // Every edit to the build re-requests an estimate; live.request debounces the burst into
  // one run and cancels whatever was already in flight. `characterFromPlanner` returns null
  // until the talent file has loaded, which `request` treats as "off" rather than an error.
  // ...but only for a build worth simming, on a device that wants it (live-gate.ts): until
  // every point is spent nothing is asked of the engine, so nothing is downloaded either, and
  // a phone is asked first. Stopping keeps the last figure on screen, dimmed.
  const constrained = untrack(() => isConstrainedDevice());
  let dpsOptedIn = $state(false);
  const gate = $derived(liveGate({ spent: store.spent, constrained, optedIn: dpsOptedIn }));

  // A figure from another class is not a stale answer to this build. Declared BEFORE the
  // request effect below on purpose: effects flush in declaration order, and a build pasted
  // in for another class changes the class and the talents in one go -- clearing second
  // would cancel the very run that paste had just scheduled.
  $effect(() => {
    void store.classSlug;
    untrack(() => live.clear());
  });

  $effect(() => {
    void store.order;
    void store.gear;
    if (gate === 'run') live.request(characterFromPlanner(store), store.talentIndex);
    else live.request(null, null);
  });

  // No reactive reads of its own: this effect's body runs once, on mount, purely to register
  // the teardown that runs it returns -- which is the only thing that has to happen when the
  // planner unmounts, since `live.request` above already cancels and re-schedules on its own.
  $effect(() => () => live.dispose());

  /** The current build's talent ranks, one array per tree in tab order -- encodeFS1's shape. */
  function treeRanksFor(index: TalentIndex, order: number[]): number[][] {
    const ranks = ranksByTalent(order);
    return index.trees.map((tree) => tree.talents.map((talent) => ranks.get(talent.id) ?? 0));
  }

  /**
   * The build's own FS1 code -- the planner's export format. Derived once so `simHref`
   * (below) and an embedder's `oncode` (Top Gear's "add a build") always read the identical
   * encoding of the identical build, rather than each calling `encodeFS1` again and risking
   * the two drifting apart.
   */
  const liveCode = $derived(
    store.talentIndex === null
      ? ''
      : encodeFS1({
          dataBuild: store.treeVersion,
          classSlug: store.classSlug,
          raceSlug: store.raceSlug,
          treeRanks: treeRanksFor(store.talentIndex, store.order),
          gear: store.gear,
        }),
  );

  // A `$effect` rather than a call inside a derivation: a derivation must stay a pure read,
  // and calling `oncode` is a side effect that has to run again on every build change.
  $effect(() => {
    if (liveCode !== '') oncode?.(liveCode);
  });

  // "Sim this build": a saved build's own link when it has one, otherwise `liveCode` above.
  // The link is always present so a player can reach the full results whether or not the
  // live estimate has run, or could run at all. Both branches go through
  // `simSearch`/`withSimState` rather than building the query string by hand, so /sim's own
  // URL state (lib/sim/url.ts) is the one place that encodes it.
  const simHref = $derived(
    `/sim${simSearch(
      withSimState(
        defaultSimState(),
        store.sourceId !== null
          ? { source: 'build', ref: store.sourceId }
          : liveCode !== ''
            ? { code: liveCode }
            : {},
      ),
    )}`,
  );

  let status = $state<'loading' | 'ready' | 'failed'>('loading');
  let attempt = $state(0);
  /** The build's stat weights (Task 19). Empty until loaded, or on a build with none. */
  let weights = $state<WeightsFile>([]);
  // Reset asks in the toolbar rather than through window.confirm: a browser dialog cannot be
  // styled, cannot say what it is about to clear, and reads badly on phone.
  let confirmingReset = $state(false);
  // Which panel the phone shows: a tree by index, or the gear panel last. Desktop ignores it
  // and lays every tree out side by side with the gear panel under them.
  let activeTree = $state(0);

  // Gear is the optional half of a build: a class the build ships no item file for has an
  // empty index, so it gets no panel and no tab. The tab sits after the trees, which is both
  // where it belongs in the strip and why its index can never collide with a tree's.
  const hasGear = $derived(store.itemIndex.size > 0);
  const gearTabIndex = $derived(store.talentIndex ? store.talentIndex.trees.length : 0);

  // The roving tabindex keeps the unselected tabs out of the tab order, so arrow keys are the
  // only way to reach them: without this a keyboard could never open the second tree, or the
  // gear panel behind the last tab. Moving selects, which is the automatic-activation half of
  // the ARIA tabs pattern -- switching panel costs nothing, so there is no reason to make it
  // a second keypress. The tabs come off the event rather than a `bind:this`: the listener is
  // on the tablist itself, so currentTarget is already the element, and there is no reference
  // to go stale when the branch unmounts.
  function onTabKeys(event: KeyboardEvent & { currentTarget: HTMLDivElement }): void {
    const step = event.key === 'ArrowRight' ? 1 : event.key === 'ArrowLeft' ? -1 : 0;
    if (step === 0) return;
    event.preventDefault();
    const tabs = event.currentTarget.querySelectorAll<HTMLButtonElement>('[role="tab"]');
    activeTree = (activeTree + step + tabs.length) % tabs.length;
    tabs[activeTree].focus();
  }

  // Nothing cancels a request that is already in flight, so switching class twice in quick
  // succession leaves two runs of this racing to write to the same store. The class this run
  // was started for is its generation token: a run whose class has since moved on stops before
  // it writes anything, rather than leaving the store holding one class's trees under
  // another's slug -- which `toDraft` would then post as that class's talent ids under the
  // wrong class_id, for the API to refuse. SharePanel.svelte draws the same line for saves,
  // by comparing drafts.
  async function load(): Promise<void> {
    const slug = store.classSlug;
    const stale = (): boolean => store.classSlug !== slug;

    // A decode failure is not tied to any class -- the code will never decode differently no
    // matter which class loads or how many times -- so it is shown, and `codeApplied` armed,
    // the first chance `load()` gets, independent of `slug`. A well-formed code only ever gets
    // a turn on the run whose class matches the one it names; see the comment on `codeApplied`
    // above for why that alone (with no early, unconditional arming) is already enough to stop
    // a class switch from ever applying it to the wrong class.
    if (decoded !== null && !decoded.ok && !codeApplied) {
      codeApplied = true;
      codeNote = noteForMessage(decoded.message);
    }
    const codeForThisClass =
      decoded !== null && decoded.ok && !codeApplied && decoded.build.classSlug === slug ? decoded : null;

    status = 'loading';
    try {
      const reference = await loadReference(store.treeVersion);
      if (stale()) return;
      store.setReference(reference);

      let talents: TalentFile;
      try {
        talents = await loadTalents(store.treeVersion, slug);
      } catch (error) {
        // A class this planner has no talent data for will never load no matter how many
        // times this is retried, so that -- and only that -- earns the code its one turn even
        // though nothing was applied. Any other failure here (a network blip, a 5xx) leaves
        // `codeApplied` unarmed, so a same-class Retry still gets a real chance to apply it.
        // A person who followed a build link deserves to be told the class is why: the
        // generic panel below says nothing about the code, and its Retry button cannot
        // succeed without also changing class.
        if (codeForThisClass !== null && error instanceof DataLoadError && error.status === 404) {
          codeApplied = true;
          if (!stale()) {
            codeNote = {
              kind: 'message',
              text: `That code names a class this planner does not have: ${codeForThisClass.build.classSlug}.`,
            };
          }
        }
        throw error;
      }
      if (stale()) return;
      store.setTalents(talents);

      if (codeForThisClass !== null && store.talentIndex !== null) {
        codeApplied = true;
        const rebuilt = orderFromRanks(store.talentIndex, codeForThisClass.build.treeRanks);
        store.applyOrder(rebuilt.order, codeForThisClass.build.gear);
        // A build link from the deaths recap (src/lib/report/planner-link.ts) carries gear
        // alone when the log's talents field could not be read as ranks -- every tree comes
        // through as all zeros. The note says so rather than claiming talents loaded.
        const gearOnly = codeForThisClass.build.treeRanks.every((tree) => tree.every((rank) => rank === 0));
        codeNote = { kind: 'reconstructed', dropped: rebuilt.dropped.length, gearOnly };
      }

      // Task 10: a `?code=`/restored code just applied above, or this mount's `record`.
      if (record !== null) writePointer('build', record.id, slug, record.title);
      else if (codeForThisClass !== null && codeParam !== null) {
        writePointer('code', codeParam, codeForThisClass.build.classSlug);
      }

      const sets = await loadSets(store.treeVersion);
      if (stale()) return;
      store.setSets(sets);
      // Gear is optional in the same way sets are: a build whose item table did not
      // normalize ships no items/<class>.json, and the planner is complete without a gear
      // panel. Only a 404 means that. A 5xx, an unreachable network or a malformed file is a
      // broken build rather than an absent one, so it is rethrown into the failure state --
      // the same line loadSets draws, for the same reason.
      try {
        const items = await loadItems(store.treeVersion, slug);
        if (stale()) return;
        store.setItems(items);
      } catch (error) {
        if (!(error instanceof DataLoadError) || error.status !== 404) throw error;
        if (stale()) return;
        store.setItems({ build: store.treeVersion, class_slug: slug, items: [] });
      }
      // Weights are optional the same way sets and items are: a build the data lane has
      // not regenerated ships none, and loadWeights already returns [] for a 404. Any
      // other failure here is rethrown into the outer catch, the same line loadSets draws.
      weights = await loadWeights(store.treeVersion);
      if (stale()) return;
      status = 'ready';
    } catch {
      // A stale run's failure is not this class's failure: the run that replaced it owns the
      // status, and reporting this one would put a working planner behind a Retry button.
      if (stale()) return;
      status = 'failed';
    }
  }

  // Re-runs whenever the class changes (selectClass drops the loaded trees) or Retry bumps
  // `attempt`. Reading both synchronously here is what registers them as dependencies; the
  // writes `load` performs happen after the tracking window, so this cannot loop.
  $effect(() => {
    void store.classSlug;
    void attempt;
    void load();
  });

  // Switching class empties the build on its own, and the class selector stays reachable while
  // the confirm is open. Without this the planner comes back asking whether to clear a build the
  // switch already cleared. The phone tab goes back to the first tree for the same reason: the
  // open tree belonged to the class that just left, and a class with fewer trees than the last
  // one would leave no tab selected and no panel shown at all. Its own effect rather than a line
  // in the one above: that effect documents a careful no-loop invariant, and this has nothing to
  // do with loading.
  // Compared against the class this effect last saw, rather than cleared unconditionally, so
  // that the very first run -- which fires once on mount, in the same synchronous flush as the
  // load effect's own first run -- cannot wipe out a decode failure that load() may already
  // have written into `codeNote` moments earlier in that same flush (decode failures are
  // reported synchronously, before load()'s first `await`). Every run after the first is a
  // genuine class change, and only those should ever clear it.
  let classSlugForNoteReset = store.classSlug;
  $effect(() => {
    void store.classSlug;
    confirmingReset = false;
    activeTree = 0;
    if (store.classSlug !== classSlugForNoteReset) {
      // A code's note describes why the build looked the way it did on the class it was shown
      // under -- switching class already discards that build (selectClass), so a note left
      // behind (most visibly the "does not have" message: switching class is its entire
      // remedy) would sit under an unrelated, working build claiming something no longer true.
      // Every kind of note is cleared the same way; none of them describes anything about a
      // class the planner has since moved on from.
      codeNote = null;
      classSlugForNoteReset = store.classSlug;
    }
  });
</script>

<div class="flex flex-col gap-[22px] md:gap-8" data-testid="planner">
  {#if standalone}
    <!-- Task 10: fixed-height slot, mirrors ToolsView.svelte's own `sim-chip-slot`. -->
    <div class={CHIP_HEIGHT} data-testid="planner-chip-slot">
      <CurrentCharacterChip
        current={pointer}
        {restored}
        {addonCode}
        hasOwnPasteBox={!store.readOnly}
        onforget={onForgetPointer}
      />
    </div>
  {/if}
  <SummaryBar {store} {live} {simHref} {gate} {standalone} onshowdps={() => (dpsOptedIn = true)} />

  <p class="text-muted px-[18px] text-[13px] md:px-0">{treeSourceNotice(store.treeVersion)}</p>

  {#if codeNote !== null}
    <p class="text-muted px-[18px] text-[13px] md:px-0" data-testid="planner-code-note">
      {#if codeNote.kind === 'tree-count'}
        That code has <span class="tabular font-mono">{codeNote.got}</span> talent trees; a build has
        <span class="tabular font-mono">{codeNote.want}</span>.
      {:else if codeNote.kind === 'reconstructed'}
        {#if codeNote.gearOnly}
          Gear loaded from a character. The log did not record talent ranks for this build.
        {:else if codeNote.dropped === 0}
          Talents loaded from a character. The order points were spent in is not recorded in game, so this is
          the lowest-tier-first order that reaches the same tree.
        {:else}
          Talents loaded from a character, minus
          <span class="tabular font-mono">{codeNote.dropped}</span>
          that no legal order reaches. The order is a reconstruction: the game does not record the order points
          were spent in.
        {/if}
      {:else}
        {codeNote.text}
      {/if}
    </p>
  {/if}

  <!-- The three states below swap in place once the talent data arrives over the network, and
       they are wildly different heights: one line of status text against a planner several
       hundred pixels tall. Whatever is under the planner -- the footer, mainly -- moves by
       that difference, which Lighthouse measured as 0.185 of /planner.html's 0.186 CLS
       against the 0.05 lighthouserc.json budget. This min-height reserves the room up front
       so the swap moves nothing below it.

       What it does not do is make the three states identical, and it does not abolish the
       swap. The reserve is one number; the ready height is not. These two come from the
       loaded layout of the default class, so a class whose trees run to more tiers grows past
       them and still moves the footer -- by the difference rather than by the whole planner.
       Re-derive them by loading /planner, setting this element's min-height to 0, and reading
       its `getBoundingClientRect().height` below and above the md breakpoint. They measure
       1039.5 at 360px and 1379 from md up (the addon lane's Task 16 added `<ImportBox>` --
       a heading, a two-row textarea and a submit button -- inside this region on the
       non-read-only mount, plus SharePanel's "Copy addon code" button next to Share and the
       always-present hint paragraph under that row; together that grew the region by about
       246.5px at 360px and 300px from md up over the 793/1079 this measured before it. The
       gearless-desktop figure below quoting 616 is on the same non-read-only mount as this
       min-height, so it too now understates the ready height by roughly that same growth
       and was not re-measured (it stays far enough under either reserve for the conclusion
       below to hold regardless). The read-only figures quoting 646.5 and 962 are unaffected
       by this lane -- ImportBox and SharePanel's addon-code button are both `!readOnly`-only
       -- and remain as measured; both predate the Task 21 checkbox change and Task 11's tree
       header change before it, and were not re-measured, since neither feeds this min-height
       and re-deriving them needs a gear-mount scenario outside what the checked-in fixture
       data covers).

       The op-character lane (2026-09-21) added one more `<ImportBox>` line -- a "Get the
       addon" link, same non-read-only mount as Task 16's own addition -- growing the
       naturals from 1039.5/1379 to 1095.5 at 360px and 1435 from md up (the gearless-desktop
       and read-only figures below sit outside that mount, comfortably under either reserve
       either way). Rounded up per this comment's own convention: 1095.5 becomes 1096, 1435
       is already whole; under costs movement, over costs only dead space.

       Two things this comment used to have wrong, both settled by measurement. A reserve that
       is *too large* does not haul the footer up in the failed-to-load state: the min-height
       sits on the container wrapping all three branches, so a larger reserve binds identically
       in every one of them -- it buys dead space, never movement. Failed-to-load's own natural
       height is 171.5 at 360px and 144 from md up, far below any candidate reserve, so the
       reserve is what that branch measures whatever the reserve is. The only constraint on
       this number is `reserve >= the loaded natural`; overshooting costs blank space alone.
       And the Playwright *project* is irrelevant to every height here: Pixel 7 and Desktop
       Chrome return byte-identical heights at equal viewports, and
       tests/e2e/planner-phone.spec.ts sets its own
       `test.use({ viewport: { width: 360, height: 800 } })`, which overrides the mobile
       project's 412px -- so `--project=mobile` measures 360px, not 412. Only viewport width
       moves these numbers.

       One standing caveat, as true of the base reserve as of this one: every figure in this
       comment is measured against the *fixture* build (FOREVER_DATA=fixture, a two-tree
       warrior), so the reserve has only ever been sized to the fixture's loaded height -- at
       the base too, to the pixel. Real data has never been inside it, before this lane or
       after. Restoring the equality restores the invariant this lane broke and opens no new
       real-data gap; sizing for real data is a separate question from this fix.

       The phone figure fell from 1412.5 to 728 (predating this lane) when gear became the
       third tab: the gear panel used to stack under the trees there and now takes its turn
       in the same column. What is reserved for is the tab the planner lands on, which is the
       first tree, and that tab then measured 1039.5 (1095.5 as of the op-character lane
       above) with the import box and addon-code controls counted in -- both sit above the
       tab content, so they add the same height whichever tab is open. Opening Gear then
       measured 1516.5 (was 1143 before this lane) and
       pushes the footer down by the difference, and that is deliberate -- it is a tap rather
       than an unprompted shift, the same kind of movement showing the order strip or opening
       an item picker already makes, and none of it is what CLS measures. Reserving the gear
       height instead would buy that back at the price of 477.5px of dead space under every
       build that never opens the tab. Desktop is untouched by the tab strip (it is md:hidden,
       and the gear panel still sits under the order strip there), so the md figure carries
       the same ImportBox/addon-code growth as the rest of the toolbar.

       This still earns its keep even though the ready planner is now tall enough that the
       footer is below the fold in both states, and the numbers are here so the question does
       not have to be re-opened blind: removing it entirely takes /planner.html's median CLS
       from 0.0009 to 0.1324 and its performance score from 0.98 to 0.94, because without
       it the *loading* state is short enough to leave the footer on screen and the swap then
       hauls it down from inside the viewport. Nor is the answer a number tuned to the audit:
       Lighthouse emulates a 640px-tall fold, so anything over that scores well while still
       moving the footer on the 800px-tall phone tests/e2e/planner-phone.spec.ts drives --
       which is what those two footer assertions are for. The reserve has to cover the ready
       height, not the audit's viewport.

       The phone figure is deliberately the one measured at 360px, the narrowest width the
       site designs for and the width Lighthouse emulates (lighthouserc.json). It is the
       tallest: the toolbar row wraps one button further at 360 than it does from 390px up,
       which is 57px, so above 360 the ready planner comes in under this reserve and the
       reserve is what the region measures in all three states. That is dead space rather
       than movement, and it is the safe direction to err.

       It wraps the swapping branches only, not the planner as a whole, and that is what lets
       one number hold: the summary bar and the notice above are in all three states and
       reflow with the viewport width, so keeping them outside the reserve takes their
       wrapping out of the figure. Inside it every part is a fixed height -- the tab strip,
       the toolbar (including the always-visible title field and Share button), the order
       strip's reserved row, a tree grid sized by tier count rather than by width, and --
       from md up, where gear is part of the column rather than a tab -- a gear panel whose
       slot grid is a fixed count of fixed-height rows and whose totals and sets columns
       start on their one-line empty state.

       A class the build ships no item file for loses the gear panel, and with it the Gear
       tab. On a phone that changes nothing: the tree tab is what is reserved for, and it
       measures the same 1095.5 (op-character's own figure, above). On desktop the panel
       leaves the column and the ready planner comes in at 616 (stale, see above -- still
       comfortably under the 1435 md reserve either way), dead space rather than movement.

       Fork replaces Reset and drops the SharePanel section, but only on the read-only mount --
       the editable toolbar this measures is untouched. The read-only mount is the shorter one,
       646.5 and 962, so it sits well under the reserve and leaves that much space above
       the footer on the API's /b/:id. Reserving the taller figure in both is deliberate:
       Fork grows the toolbar back to the editable height, and a reserve that
       tracked `readOnly` would spend that growth shoving the footer down the moment it is
       pressed. /b/:id carries no CLS budget of its own -- it is server-rendered, so the
       island's whole planner arrives after first paint regardless of what this reserves. -->
  <div class="flex min-h-[1096px] flex-col gap-[22px] md:min-h-[1435px] md:gap-8">
    {#if status === 'loading'}
      <!-- The planner's own panel chrome rather than a bare line on a blank reserve: a
           viewport of empty space reads as a broken page, and the frame reads as the planner
           arriving. It cannot show the trees themselves -- their names, tiers and columns are
           the very thing still loading -- so it grows to fill the reserve and says so. -->
      <div class="border-line bg-raised rounded-panel mx-[18px] flex grow flex-col gap-3 border p-4 md:mx-0">
        <Skeleton lines={4} rowHeight="h-4" label="Loading talent data" testid="planner-talent-skeleton" />
      </div>
    {:else if status === 'failed'}
      <div class="border-line bg-raised rounded-panel mx-[18px] flex flex-col gap-3 border p-5 md:mx-0">
        <p class="text-strong text-[15px] font-semibold">{DATA_LOAD_FAILED}</p>
        <!-- Honest for every branch of `load`: the reference files fail into this state too,
             not just talents/<class>.json. -->
        <p class="text-muted text-[13px]">
          Build {store.treeVersion} did not return the files the planner needs.
        </p>
        <LoadError
          message="Talent data did not load."
          onRetry={() => (attempt += 1)}
          testid="planner-load-error"
        />
      </div>
    {:else if store.talentIndex}
      <!-- One panel at a time on a phone: three trees side by side do not fit 360px, and
           stacking them -- with the gear panel's seventeen slots under them -- puts the last
           one several screens down. Desktop keeps the columns and hides this. The roving
           tabindex lives on the tabs, as it does on TreeGrid's cells.
           The -1 on the container changes nothing about the keyboard order -- a bare div was
           never a tab stop -- and is there to satisfy the compiler's a11y rule that an element
           carrying an interactive role and a key handler declare a tabindex; -1 declares one
           without adding a stop, and it matches what TreeGrid's `role="grid"` does. -->
      <div
        role="tablist"
        tabindex={-1}
        aria-label="Planner sections"
        class="border-line-soft mx-[18px] flex gap-2 border-b pb-2 md:hidden"
        onkeydown={onTabKeys}
      >
        {#each store.talentIndex.trees as tree, i (tree.id)}
          <button
            type="button"
            role="tab"
            id={`tree-tab-${tree.id}`}
            aria-selected={activeTree === i}
            aria-controls={`tree-panel-${tree.id}`}
            tabindex={activeTree === i ? 0 : -1}
            class="{SECONDARY_BUTTON} flex-1 justify-center {activeTree === i
              ? 'border-gold text-gold'
              : 'border-line text-nav'}"
            onclick={() => (activeTree = i)}
          >
            <span>{tree.name}</span>
            <span class="tabular text-muted ml-2 font-mono">{store.split[i] ?? 0}</span>
          </button>
        {/each}
        <!-- Last, so `gearTabIndex` is the tree count and onTabKeys picks it up from the
             tablist's own DOM order without knowing gear exists. No count beside the name:
             the trees show the points spent in them because that number is otherwise only on
             the panel behind the tab, and the gear panel's own totals are not one number. -->
        {#if hasGear}
          <button
            type="button"
            role="tab"
            id="gear-tab"
            aria-selected={activeTree === gearTabIndex}
            aria-controls="gear-tabpanel"
            tabindex={activeTree === gearTabIndex ? 0 : -1}
            class="{SECONDARY_BUTTON} flex-1 justify-center {activeTree === gearTabIndex
              ? 'border-gold text-gold'
              : 'border-line text-nav'}"
            onclick={() => (activeTree = gearTabIndex)}
          >
            Gear
          </button>
        {/if}
      </div>

      <div class="grid grid-cols-1 gap-4 px-[18px] md:grid-cols-3 md:px-0" data-testid="tree-columns">
        {#each store.talentIndex.trees as tree, i (tree.id)}
          <!-- The inactive trees are hidden with a class, not the `hidden` attribute: the
               attribute would hide them on desktop too, where `md:flex` cannot override it. -->
          <div
            id={`tree-panel-${tree.id}`}
            role="tabpanel"
            aria-labelledby={`tree-tab-${tree.id}`}
            data-testid={`tree-panel-${tree.id}`}
            class="border-line-warm bg-raised rounded-panel flex-col gap-3 border p-4 md:flex {activeTree ===
            i
              ? 'flex'
              : 'hidden'}"
          >
            <!-- The game's tree header: name on the left, points in the tree on the
                 right, a rule under both. Warm border and gold number are the
                 design system's; the proportions are the client's. -->
            <header class="border-line-soft flex items-baseline justify-between border-b pb-2">
              <h2 class="section-title text-[15px]">{tree.name}</h2>
              <span class="tabular text-gold font-mono text-[15px]" data-testid={`tree-points-${tree.id}`}>
                {store.split[i] ?? 0}
              </span>
            </header>
            <TreeGrid {store} {tree} />
          </div>
        {/each}
      </div>

      <div class="flex flex-wrap items-center gap-3 px-[18px] md:px-0" data-testid="planner-toolbar">
        {#if store.readOnly}
          <!-- A build opened from a share link. Every edit is refused until Fork, so the
               toolbar says so up front rather than leaving the refusal message to explain it
               after the first click. Reset and Share are gone with it: there is nothing of
               one's own to clear, and re-sharing someone else's build under a new id is the
               one thing Fork is for. -->
          <p class="text-muted text-[13px]">
            This build was shared as a link. Fork it to spend points of your own.
          </p>
          <button
            type="button"
            class="{SECONDARY_BUTTON} border-line-warm-strong text-gold px-4"
            onclick={() => store.fork()}
          >
            Fork
          </button>
        {:else}
          <!-- Nested rather than a third arm of the branch above, so `readOnly` is asked once:
               Reset and Share belong to the same half of that decision, and SharePanel has to
               sit outside the confirm to survive it -- it holds the title being typed and the
               link of the last save, and re-mounting it when the confirm opens would throw
               both away. -->
          {#if confirmingReset}
            <span class="text-muted text-[13px]">Clear every point in this build?</span>
            <button
              type="button"
              class="{SECONDARY_BUTTON} border-line-warm-strong text-gold px-4"
              onclick={() => {
                store.reset();
                confirmingReset = false;
              }}
            >
              Clear all points
            </button>
            <!-- Reset leaves the DOM the moment it is pressed, so the keyboard lands on the
                 question it just asked rather than back at the top of the document. It lands
                 on the safe answer: a second Enter pressed out of habit keeps the build rather
                 than clearing it, which is the only reason the second step exists. -->
            <button
              type="button"
              class="{SECONDARY_BUTTON} border-line-warm text-text px-4"
              {@attach (node) => node.focus()}
              onclick={() => (confirmingReset = false)}
            >
              Keep the build
            </button>
          {:else}
            <button
              type="button"
              class="{SECONDARY_BUTTON} border-line-warm text-text px-4"
              onclick={() => (confirmingReset = true)}
            >
              Reset…
            </button>
          {/if}

          <SharePanel {store} {live} />
        {/if}
      </div>

      {#if !store.readOnly}
        <!-- A read-only build (opened from a share link) has nowhere for an imported build to
             go until it is forked, so the box only mounts once the toolbar above already shows
             Reset and Share rather than "Fork it to spend points of your own." -->
        <ImportBox
          talents={store.talentIndex}
          activeBuild={activeBuild.build}
          onimport={(build, pastedCode) => {
            store.loadImported(build);
            writePointer('addon', pastedCode, build.classSlug);
          }}
        />
      {/if}

      <OrderStrip {store} />

      <!-- The gear tab's panel. Hidden by a class rather than the `hidden` attribute for the
           same reason the tree panels are: the attribute would hide it on desktop too, where
           it belongs under the order strip and `md:flex` could not override it. -->
      {#if hasGear}
        <div
          id="gear-tabpanel"
          role="tabpanel"
          aria-labelledby="gear-tab"
          class="flex-col md:flex {activeTree === gearTabIndex ? 'flex' : 'hidden'}"
        >
          <GearPanel {store} {weights} />
        </div>
      {/if}
    {/if}
  </div>
</div>
