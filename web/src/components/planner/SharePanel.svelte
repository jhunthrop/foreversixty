<!-- web/src/components/planner/SharePanel.svelte -->
<!-- Saving a build and showing its link. The build stays in the store whatever the API
     says, so a failed save costs nothing but a retry. The preview card is the same PNG
     Discord and Twitter unfurl.

     Task 21: a saved build's card can carry a simmed DPS figure. The card is rendered by
     the API, which joins it in from the newest sim whose source is this build's own id --
     no new field on BuildDraft or the card route. The panel's job is only to run that sim,
     after the build has an id, and to never let a failed sim take the save down with it. -->
<script lang="ts">
  import { tick } from 'svelte';
  import { addonCopy } from '../../lib/addon/copy';
  import { plannerCopy } from '../../lib/planner/copy';
  import { plannerAddonCode, unsavedPlannerHref } from '../../lib/planner/current-character-planner';
  import type { LiveDps } from '../../lib/planner/live-dps.svelte';
  import {
    cardUrlFor,
    saveBuild,
    sharedFields,
    UNSAVABLE_BUILD_MESSAGE,
    type SaveOutcome,
    type SavedBuild,
  } from '../../lib/planner/share';
  import { MAX_TITLE_LENGTH, type PlannerStore } from '../../lib/planner/store.svelte';
  import { SECONDARY_BUTTON } from '../../lib/planner/styles';
  import type { BuildDraft } from '../../lib/planner/types';
  import { saveSim } from '../../lib/sim/api';
  import { characterFromPlanner, toCharacterSpec } from '../../lib/sim/character';
  import { simCopy } from '../../lib/sim/copy';
  import { runSim } from '../../lib/sim/run';
  import { defaultSettings } from '../../lib/sim/settings';
  import { referenceStatOf } from '../../lib/sim/spec-label';
  import { ITERATIONS } from '../../lib/sim/types';
  import { createPool, type SimPool } from '../../lib/sim/worker';

  // Retry and Copy link are both the neutral (non-gold) secondary button; only the layout
  // around them differs.
  const NEUTRAL_BUTTON = `${SECONDARY_BUTTON} border-line-warm text-text px-4`;

  let { store, live }: { store: PlannerStore; live: LiveDps } = $props();

  // Per-instance, not a literal id: Top Gear mounts a second, inline Planner (and so a
  // second SharePanel) on the same page, and two static `id="share-confirm-heading"`
  // elements would make `aria-labelledby` ambiguous for whichever one is not first in the
  // DOM. `$props.id()` is Svelte's own per-component unique id -- a different problem from
  // `SearchBox.svelte`'s own `search-title-${i}` ids, which disambiguate multiple results
  // inside one component instance by loop index, not one component instance from another.
  const uid = $props.id();
  const shareConfirmHeadingId = `share-confirm-heading-${uid}`;

  let saving = $state(false);
  let outcome = $state<SaveOutcome | null>(null);
  let cardBroken = $state(false);
  /**
   * Which button last copied, or null. One flag for every copy button rather than one each:
   * only one thing can be on the clipboard, so only one button should say so, and every
   * button then behaves identically instead of the panel having three copy idioms.
   *
   * No timer. "Copied" stands until the build changes, which the $effect below already
   * watches for -- and it is true for exactly that long, because the text on the clipboard
   * goes stale at exactly that moment. A timed revert would need a cleared handle, a named
   * duration and a cleanup on destroy to say something less accurate.
   */
  let copiedFrom = $state<'link' | 'addon' | 'unsaved' | null>(null);

  // Derived, not computed on click: the button is disabled while there is nothing to
  // copy, and a $derived keeps that in step with every talent and gear edit for free.
  // talentIndex is null until the talent data has loaded (or has failed to); the addon
  // code is the empty string then, rather than a cast that would crash on that failure.
  // `plannerAddonCode` (Task 10) is the one derivation, shared with Planner.svelte's own
  // current-character chip so the two "Copy addon code" surfaces can never disagree.
  const addonCode = $derived(plannerAddonCode(store));

  // Same story for the confirm step's "unsaved link" alternative: path and query only, no
  // origin (`window.location.origin` is only ever read inside the click handler below, so
  // this stays safe to evaluate during SSR). '' before talent data has loaded.
  const unsavedHref = $derived(unsavedPlannerHref(store));

  async function copyToClipboard(text: string, source: 'link' | 'addon' | 'unsaved'): Promise<void> {
    try {
      await navigator.clipboard.writeText(text);
      copiedFrom = source;
    } catch {
      copiedFrom = null;
    }
  }

  function copyUnsavedLink(): void {
    if (unsavedHref === '') return;
    void copyToClipboard(`${window.location.origin}${unsavedHref}`, 'unsaved');
  }

  /** Checked by default when the planner has an estimate for this build; see the panel. */
  let includeSim = $state(true);
  let cardState = $state<'idle' | 'running' | 'done' | 'skipped'>('idle');
  let cardDps = $state('');
  let cardVersion = $state('');

  /**
   * A talent edit and a re-save while the previous save's card sim is still running starts
   * a second `attachSim` call over the same three `$state` variables above. Each call is
   * stamped with the generation `share()` was on when it started; a call whose stamp no
   * longer matches -- or whose build is no longer the one on screen -- writes nothing. The
   * draft comparison in `share()` already drops a stale *build* response; this is that same
   * discipline applied to the slower sim response layered on top of it.
   */
  let saveGeneration = 0;

  const saved = $derived(outcome !== null && outcome.ok ? outcome.build : null);
  const failure = $derived(outcome !== null && !outcome.ok ? outcome : null);

  // The checkbox reads true only while a live estimate is actually ready: a run that never
  // started, is mid-debounce, or failed (most often an unsupported spec) has nothing worth
  // running the card sim from, so the box shows unchecked and disabled rather than checked
  // against a build the engine cannot run.
  const canIncludeSim = $derived(live.state === 'ready');
  const includeSimChecked = $derived(canIncludeSim && includeSim);

  /**
   * One pool for the life of the panel, created on the first save and not before: the
   * planner's own Lighthouse budget is the reason the engine is never touched on mount.
   * `live` already holds one for the inline estimate, but it is private to that factory
   * and a shared handle would couple two features that cancel each other's runs.
   */
  let pool: SimPool | null = null;
  function poolOnce(): SimPool {
    pool ??= createPool();
    return pool;
  }

  // Any edit invalidates the link that was shown: builds are immutable, so a changed build
  // is a different build. The card sim's own status resets with it, for the same reason,
  // and the confirm step closes too: it asked about a build that no longer exists.
  //
  // Reading `draftOrNull()` -- not `store.order.length`/`store.gear` as before -- is what
  // makes this catch every field `toDraft()` actually sends, not just two of them: it reads
  // `classRow`/`raceRow` (so a class or race change counts), the whole `order` array via
  // its `[...order]` spread (so reordering points between trees at the same total count
  // counts, which watching only `.length` missed), `gear` (unchanged), and `title` (which
  // the old effect did not watch at all -- a title-only edit after a failed save used to
  // let Retry post the new title under a confirm that had shown the old one).
  //
  // `includeSim` -- the checkbox's own state, read raw rather than through
  // `includeSimChecked` -- is watched the same way: unticking "Include a sim on the card"
  // is treated as an edit too, so a failed save's Retry (which would otherwise reuse
  // whatever `confirmedIncludeSim` the earlier confirm captured) disappears, and the next
  // Share opens a fresh confirm that names what will actually happen. `includeSimChecked`
  // would also flip -- silently, with no click of the box at all -- the moment a pending
  // live estimate becomes ready, which would close a confirm the visitor had just opened
  // to read; `includeSim` only ever changes from the checkbox's own `onchange` below, so
  // that false trigger cannot happen. This still runs once on mount, harmlessly, since
  // `outcome`/`confirmOpen` already hold these exact values then.
  $effect(() => {
    void draftOrNull();
    void includeSim;
    outcome = null;
    copiedFrom = null;
    cardState = 'idle';
    cardDps = '';
    cardVersion = '';
    confirmOpen = false;
  });

  /** True once a newer save has started, or the build on screen is no longer this one. */
  function supersededSince(generation: number, build: SavedBuild): boolean {
    return generation !== saveGeneration || saved?.id !== build.id;
  }

  async function attachSim(build: SavedBuild, generation: number): Promise<void> {
    const character = characterFromPlanner(store);
    if (!confirmedIncludeSim || character === null || store.talentIndex === null) return;
    cardState = 'running';
    // The card sim runs at the settings the sim page opens on -- raid-buffed, three
    // minutes, single target, split physical or caster by this build's own spec. The
    // planner has no settings bar of its own, and inventing one here would put two
    // vocabularies on one page; `/sim` is where settings are chosen.
    const settings = defaultSettings(referenceStatOf(character.spec));
    try {
      const result = await runSim(
        poolOnce(),
        {
          spec: character.spec,
          // The saved build is what the card is for, so that is the sim's source and the
          // id the API's card renderer joins on. No new field on either shape.
          source: { kind: 'build', ref: build.id, captured_at: new Date().toISOString() },
          character: toCharacterSpec(character, store.talentIndex, settings.buffs, settings.consumables),
          encounter: settings.encounter,
          iterations: ITERATIONS.normal,
        },
        () => {},
      ).result;
      await saveSim(result);
      // A re-save started (and possibly finished) while this run was in flight. Its own
      // attachSim call owns cardState/cardDps/cardVersion now; this one writes nothing.
      if (supersededSince(generation, build)) return;
      cardDps = Math.round(result.dps.mean).toLocaleString('en-US');
      cardVersion = result.engine_version;
      cardState = 'done';
    } catch {
      // The build is saved and its link works. A background sim is not worth losing that.
      if (supersededSince(generation, build)) return;
      cardState = 'skipped';
    }
  }

  /**
   * The draft, or null when the store cannot compose one. `toDraft()` throws when the class
   * or race does not resolve against the loaded reference data; unguarded, that throw
   * escapes `share()` before `saving = false` and leaves the button reading "Saving" for
   * the rest of the session with nothing said. The store repairs both fields on every write
   * it owns, so this should be unreachable -- a permanently stuck button is too bad a
   * failure mode to leave resting on that.
   */
  function draftOrNull(): BuildDraft | null {
    try {
      return store.toDraft();
    } catch {
      return null;
    }
  }

  /**
   * Share asks before it writes (one-product spec section 1): the button opens this confirm
   * step rather than saving straight away, naming what `share()`'s own `saveBuild` call is
   * about to post publicly. `draftOrNull() === null` (the unsavable-build edge case) lists
   * nothing rather than guessing at fields that were never going to be sent.
   */
  let confirmOpen = $state(false);
  let shareButtonEl: HTMLButtonElement | undefined;
  let confirmHeadingEl: HTMLParagraphElement | undefined = $state();
  const shareItems = $derived.by(() => {
    const draft = draftOrNull();
    return draft === null ? [] : sharedFields(draft, includeSimChecked);
  });

  /**
   * Whether a sim will be attached, captured the instant "Share anyway" runs rather than
   * re-read live afterward. `attachSim` reads this, not `includeSim`/`live.state` directly:
   * a live estimate that leaves 'ready' between the confirm and the save resolving (or a
   * later Retry of the same failed save) can then neither attach a sim the confirm never
   * listed nor skip one it did -- what the visitor confirmed is what actually happens.
   * Retry reuses it unchanged, for the same reason: it re-attempts the save the visitor
   * already confirmed, not a fresh one.
   */
  let confirmedIncludeSim = $state(false);

  /**
   * Opens the confirm and, once its markup has rendered, moves focus onto its own heading --
   * not the "Share anyway" button, so Tab from there reaches every option in order rather
   * than skipping the first. `tick()`, not a `$effect` reading `confirmHeadingEl`: the
   * element is only ever read here, imperatively, the same as every other `bind:this` focus
   * move in this codebase (Planner.svelte's tab arrow keys, SearchBox.svelte's own input).
   */
  async function requestShare(): Promise<void> {
    confirmOpen = true;
    await tick();
    confirmHeadingEl?.focus();
  }

  /** Closes the confirm and returns focus to the button that opened it. */
  function cancelShare(): void {
    confirmOpen = false;
    shareButtonEl?.focus();
  }

  async function confirmedShare(): Promise<void> {
    confirmOpen = false;
    confirmedIncludeSim = includeSimChecked;
    await share();
  }

  /**
   * Escape cancels while the confirm is open, from anywhere on the page -- the same
   * `<svelte:window>` idiom `Disclosure.svelte` already uses for the identical reason: the
   * key can fire while focus is still on a button inside the confirm, not a keydown handler
   * on the (non-interactive) panel itself.
   */
  function onConfirmKeydown(event: KeyboardEvent): void {
    if (event.key !== 'Escape') return;
    cancelShare();
  }

  async function share(): Promise<void> {
    saving = true;
    cardBroken = false;
    copiedFrom = null;
    // Stamped onto this save's own attachSim call, below, before anything about it can be
    // stale -- a later share() bumps this past whatever a still-running sim was stamped
    // with, which is how that sim's eventual result recognises it no longer owns the card.
    saveGeneration += 1;
    const generation = saveGeneration;
    const draft = draftOrNull();
    if (draft === null) {
      saving = false;
      outcome = { ok: false, message: UNSAVABLE_BUILD_MESSAGE, fields: {} };
      return;
    }
    const result = await saveBuild(draft);
    saving = false;
    // The grid stays editable while a save is in flight -- a slow API should not lock the
    // planner -- so the build this response describes may no longer be the current one.
    // Comparing the drafts rather than disabling input catches every kind of edit (points,
    // gear, title, even a class switch) and discards a response that arrives stale.
    if (JSON.stringify(draftOrNull()) !== JSON.stringify(draft)) return;
    outcome = result;
    // The sim is saved after the build, because it needs the build's id, and it never
    // blocks the link from showing: the share URL above is already on screen the moment
    // this line starts, and the running/done/skipped line below fills in beside it.
    if (result.ok) void attachSim(result.build, generation);
  }

  $effect(() => () => pool?.terminate());
</script>

<svelte:window onkeydown={confirmOpen ? onConfirmKeydown : undefined} />

<section class="flex w-full flex-col gap-3" data-testid="share-panel">
  <div class="flex flex-wrap items-end gap-3">
    <label class="flex flex-col gap-1">
      <span class="label text-muted">Title</span>
      <input
        type="text"
        maxlength={MAX_TITLE_LENGTH}
        value={store.title}
        oninput={(event) => store.setTitle(event.currentTarget.value)}
        placeholder="Optional"
        class="border-line-warm rounded-control bg-raised text-text placeholder:text-muted h-11 w-[260px] border px-3 text-[14px]"
      />
    </label>
    <button
      type="button"
      class="{SECONDARY_BUTTON} border-line-warm-strong text-gold px-4"
      disabled={saving || store.spent === 0}
      data-testid="share-open"
      bind:this={shareButtonEl}
      onclick={() => void requestShare()}
    >
      {saving ? plannerCopy.saving : plannerCopy.share}
    </button>
    <button
      type="button"
      class={NEUTRAL_BUTTON}
      data-testid="copy-addon-code"
      disabled={addonCode === ''}
      onclick={() => copyToClipboard(addonCode, 'addon')}
    >
      {copiedFrom === 'addon' ? addonCopy.copiedAddonCode : addonCopy.copyAddonCode}
    </button>
  </div>
  <p class="text-muted text-[13px]">{addonCopy.addonCodeHint}</p>

  <label
    class="text-muted flex min-h-11 w-fit items-center gap-2 text-[13px] md:min-h-0"
    title={canIncludeSim ? undefined : simCopy.buildSimUnavailable}
  >
    <input
      type="checkbox"
      checked={includeSimChecked}
      disabled={!canIncludeSim}
      onchange={(event) => (includeSim = event.currentTarget.checked)}
    />
    {simCopy.includeSimOnCard}
  </label>

  {#if confirmOpen}
    <div
      class="border-line bg-raised rounded-panel flex flex-col gap-3 border p-3"
      role="group"
      aria-labelledby={shareConfirmHeadingId}
      data-testid="share-confirm"
    >
      <p
        id={shareConfirmHeadingId}
        tabindex="-1"
        bind:this={confirmHeadingEl}
        class="text-strong text-[14px] outline-none"
      >
        {plannerCopy.shareConfirmTitle}
      </p>
      <p class="text-muted text-[13px]">{plannerCopy.shareConfirmIntro}</p>
      <ul class="text-muted list-disc pl-5 text-[13px]">
        {#each shareItems as item (item)}
          <li>{item}</li>
        {/each}
      </ul>
      <div class="flex flex-wrap gap-3">
        <button
          type="button"
          class="{SECONDARY_BUTTON} border-line-warm-strong text-gold px-4"
          data-testid="share-confirm-proceed"
          onclick={() => void confirmedShare()}
        >
          {plannerCopy.shareConfirmProceed}
        </button>
        <button type="button" class={NEUTRAL_BUTTON} data-testid="share-confirm-cancel" onclick={cancelShare}>
          {plannerCopy.shareConfirmCancel}
        </button>
        <button
          type="button"
          class={NEUTRAL_BUTTON}
          disabled={addonCode === ''}
          data-testid="share-confirm-copy-code"
          onclick={() => void copyToClipboard(addonCode, 'addon')}
        >
          {copiedFrom === 'addon' ? addonCopy.copiedAddonCode : plannerCopy.shareConfirmCopyCode}
        </button>
        <button
          type="button"
          class={NEUTRAL_BUTTON}
          disabled={unsavedHref === ''}
          data-testid="share-confirm-copy-unsaved"
          onclick={copyUnsavedLink}
        >
          {copiedFrom === 'unsaved' ? plannerCopy.copiedUnsavedLink : plannerCopy.shareConfirmCopyUnsaved}
        </button>
      </div>
    </div>
  {/if}

  {#if failure}
    <div class="border-line bg-raised rounded-panel flex flex-col gap-2 border p-3" role="alert">
      <p class="text-strong text-[14px]">{failure.message}</p>
      {#each Object.entries(failure.fields) as [field, message] (field)}
        <p class="text-muted text-[13px]"><span class="font-mono">{field}</span>: {message}</p>
      {/each}
      <button type="button" class="{NEUTRAL_BUTTON} w-fit" onclick={share}> Retry </button>
    </div>
  {/if}

  {#if saved}
    <div class="border-line bg-raised rounded-panel flex flex-col gap-3 border p-3">
      <div class="flex flex-wrap items-center gap-3">
        <a href={saved.url} class="font-mono text-[14px]" data-testid="share-link">{saved.url}</a>
        <button type="button" class={NEUTRAL_BUTTON} onclick={() => copyToClipboard(saved.url, 'link')}>
          {copiedFrom === 'link' ? 'Copied' : 'Copy link'}
        </button>
      </div>
      {#if cardState !== 'idle'}
        <p class="text-muted text-[13px]" data-testid="build-sim-status">
          {#if cardState === 'running'}
            {simCopy.buildSimRunning}
          {:else if cardState === 'done'}
            {simCopy.buildSimDone(cardDps, cardVersion)}
          {:else}
            {simCopy.buildSimSkipped}
          {/if}
        </p>
      {/if}
      {#if !cardBroken}
        <img
          src={cardUrlFor(saved.url)}
          alt={`Preview card for build ${saved.id}`}
          width="600"
          height="315"
          loading="lazy"
          decoding="async"
          class="rounded-panel border-line w-full max-w-[600px] border"
          onerror={() => (cardBroken = true)}
        />
      {/if}
    </div>
  {/if}
</section>
