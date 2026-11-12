<!-- web/src/components/planner/SharePanel.svelte -->
<!-- Share this build (spec 2026-11-08 section 4.8): the title, Share (which asks first, naming
     what goes public), Copy for chat (the build as one line of text), and the switch that sims a DPS
     onto the card. Saved: the link, Copy link, the card line and the 600 x 315 card. The build
     stays in the store whatever the API says, so a failed save costs nothing but a retry. Copy
     for the game left for In the game. Task 21: a saved build's card can carry a simmed DPS; the
     API joins it in from the newest sim whose source is this build's id, so the panel only runs
     that sim after the build has an id, and never lets a failed sim take the save down. -->
<script lang="ts">
  import Switch from '../ui/Switch.svelte';
  import { tick, untrack } from 'svelte';
  import { plannerCopy } from '../../lib/planner/copy';
  import { CopyState } from '../../lib/ui/copy-state.svelte';
  import { ensureItems } from '../../lib/planner/items-ready';
  import { unsavedPlannerHref } from '../../lib/planner/current-character-planner';
  import {
    cardUrlFor,
    saveBuild,
    sharedFields,
    UNSAVABLE_BUILD_MESSAGE,
    type SaveOutcome,
    type SavedBuild,
  } from '../../lib/planner/share';
  import { MAX_TITLE_LENGTH, type PlannerStore } from '../../lib/planner/store.svelte';
  import type { BuildDraft } from '../../lib/planner/types';
  import { saveSim } from '../../lib/sim/api';
  import { characterFromPlanner, toCharacterSpec } from '../../lib/sim/character';
  import { plannerSimCopy as simCopy } from '../../lib/sim/planner-copy';
  import { runSim } from '../../lib/sim/run';
  import { loadPresetSets } from '../../lib/sim/preset-sets';
  import { defaultSettings } from '../../lib/sim/settings';
  import { ITERATIONS } from '../../lib/sim/types';
  import { createPool, type SimPool } from '../../lib/sim/worker';
  import Button from '../ui/Button.svelte';
  import Field from '../ui/Field.svelte';
  import Section from '../ui/Section.svelte';
  import PlannerPart from './PlannerPart.svelte';

  import { grouped } from '../../lib/format/number';
  let { store, part = false }: { store: PlannerStore; part?: boolean } = $props();

  // Per-instance, not a literal id: Top Gear mounts a second, inline Planner (and so a
  // second SharePanel) on the same page, and two static `id="share-confirm-heading"`
  // elements would make `aria-labelledby` ambiguous for whichever one is not first in the
  // DOM. `$props.id()` is Svelte's own per-component unique id, disambiguating one
  // component instance from another -- a different problem from disambiguating multiple
  // items within one instance by loop index.
  const uid = $props.id();
  /** The title field: the store keeps it trimmed and capped. */
  let title = $state(untrack(() => store.title));
  $effect(() => {
    store.setTitle(title);
  });
  const shareConfirmHeadingId = `share-confirm-heading-${uid}`;
  const includeSimLabelId = `share-include-sim-label-${uid}`;

  let saving = $state(false);
  let outcome = $state<SaveOutcome | null>(null);
  let cardBroken = $state(false);
  /**
   * Which button last copied: one CopyState for every copy button, since only one thing can be
   * on the clipboard, so only one button says so. "Copied" clears after two seconds (owner
   * ruling 2026-11-10), and sooner when the build changes and the copied text goes stale.
   */
  const copier = new CopyState();
  $effect(() => () => copier.dispose());

  // Same story for the confirm step's "unsaved link" alternative: path and query only, no
  // origin (`window.location.origin` is only ever read inside the click handler below, so
  // this stays safe to evaluate during SSR). '' before talent data has loaded.
  const unsavedHref = $derived(unsavedPlannerHref(store));

  /**
   * Review finding 8 (spec §4.H): a plain-text summary for someone with no account and no
   * addon. Computed, never templated with a placeholder -- `''` before talent/reference
   * data has loaded, which the button below disables on, the same rule every other export
   * action in this panel already follows.
   */
  const textSummary = $derived.by(() => {
    const index = store.talentIndex;
    if (index === null || store.classRow === null || store.raceRow === null) return '';
    const trees = index.trees.map((tree, i) => ({ name: tree.name, points: store.split[i] ?? 0 }));
    return plannerCopy.shareTextSummary(store.raceRow.name, store.classRow.name, trees, store.level);
  });

  function copyToClipboard(text: string, source: 'link' | 'unsaved' | 'text'): Promise<void> {
    return copier.copy(text, source);
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

  /** A finished build can carry a simmed DPS on its card (the sim runs after the save, below). */
  const canIncludeSim = $derived(
    store.talentIndex !== null && store.spent > 0 && store.spent >= store.pointsAvailable,
  );
  const includeSimChecked = $derived(canIncludeSim && includeSim);

  /**
   * One pool for the life of the panel, created on the first save and not before: the
   * planner's own Lighthouse budget is the reason the engine is never touched on mount.
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
  // would also flip -- silently, with no click of the box at all -- the moment the last
  // point is spent, which would close a confirm the visitor had just opened to read; `includeSim` only ever changes from the checkbox's own `onchange` below, so
  // that false trigger cannot happen. This still runs once on mount, harmlessly, since
  // `outcome`/`confirmOpen` already hold these exact values then.
  $effect(() => {
    void draftOrNull();
    void includeSim;
    outcome = null;
    copier.reset();
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
    // minutes, single target, with this build's spec and level's own raid consumables. The
    // planner has no settings bar of its own, and inventing one here would put two
    // vocabularies on one page; `/sim` is where settings are chosen.
    try {
      await loadPresetSets();
      const settings = defaultSettings(character.spec, character.level);
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
      cardDps = grouped(result.dps.mean);
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
  let confirmHeadingEl: HTMLParagraphElement | undefined = $state();
  const shareItems = $derived.by(() => {
    const draft = draftOrNull();
    return draft === null ? [] : sharedFields(draft, includeSimChecked);
  });

  /**
   * Whether a sim will be attached, captured the instant "Share anyway" runs rather than
   * re-read afterward. `attachSim` reads this, not `includeSim` directly: a build that
   * changes between the confirm and the save resolving (or a
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
   * move in this codebase (Planner.svelte's tab arrow keys).
   */
  async function requestShare(): Promise<void> {
    confirmOpen = true;
    await tick();
    confirmHeadingEl?.focus();
  }

  /** Closes the confirm and returns focus to the button that opened it. */
  function cancelShare(): void {
    confirmOpen = false;
    document.querySelector<HTMLElement>('[data-testid="share-open"]')?.focus();
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

  // The shared gear is the part of the worn gear the class catalogue lists, so a build that wears
  // anything asks for the catalogue up front, before the confirm names what it will publish.
  $effect(() => {
    if (Object.keys(store.gear).length > 0) void ensureItems(store);
  });

  async function share(): Promise<void> {
    saving = true;
    cardBroken = false;
    copier.reset();
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

{#snippet body()}
  <div class="pl-sec">
    <div class="pl-acts pl-end">
      <div class="pl-field">
        <Field
          id={`share-title-${uid}`}
          label={plannerCopy.shareTitleLabel}
          placeholder={plannerCopy.shareTitlePlaceholder}
          maxlength={MAX_TITLE_LENGTH}
          bind:value={title}
          testid="share-title"
        />
      </div>
      <Button disabled={saving || store.spent === 0} onclick={() => void requestShare()} testid="share-open"
        >{plannerCopy.share}</Button
      >
    </div>
    <div class="pl-acts">
      <Button
        disabled={textSummary === ''}
        onclick={() => copyToClipboard(textSummary, 'text')}
        testid="copy-as-text"
        >{copier.copied('text') ? plannerCopy.copiedText : plannerCopy.copyAsText}</Button
      >
    </div>
    <div
      class="pl-switch"
      class:pl-off={!canIncludeSim}
      title={canIncludeSim ? undefined : simCopy.buildSimUnavailable}
    >
      <span id={includeSimLabelId}>{simCopy.includeSimOnCard}</span>
      <Switch
        checked={includeSimChecked}
        disabled={!canIncludeSim}
        onchange={(next) => (includeSim = next)}
        labelledby={includeSimLabelId}
        testid="share-include-sim"
      />
    </div>

    {#if confirmOpen}
      <div
        class="pl-confirm"
        role="group"
        aria-labelledby={shareConfirmHeadingId}
        data-testid="share-confirm"
      >
        <p
          id={shareConfirmHeadingId}
          tabindex="-1"
          bind:this={confirmHeadingEl}
          class="text-strong outline-none"
        >
          {plannerCopy.shareConfirmTitle}
        </p>
        <p class="text-small text-muted">{plannerCopy.shareConfirmIntro}</p>
        <ul class="text-small text-muted list-disc pl-5">
          {#each shareItems as item (item)}
            <li>{item}</li>
          {/each}
        </ul>
        <div class="pl-acts">
          <Button onclick={() => void confirmedShare()} testid="share-confirm-proceed"
            >{plannerCopy.shareConfirmProceed}</Button
          >
          <Button onclick={cancelShare} testid="share-confirm-cancel">{plannerCopy.shareConfirmCancel}</Button
          >
          <Button disabled={unsavedHref === ''} onclick={copyUnsavedLink} testid="share-confirm-copy-unsaved"
            >{copier.copied('unsaved')
              ? plannerCopy.copiedUnsavedLink
              : plannerCopy.shareConfirmCopyUnsaved}</Button
          >
        </div>
      </div>
    {/if}

    {#if failure}
      <div class="pl-confirm" role="alert" data-testid="share-failure">
        <p class="text-wipe">{failure.message}</p>
        {#each Object.entries(failure.fields) as [field, message] (field)}
          <p class="text-small text-muted">{field}: {message}</p>
        {/each}
        <div class="pl-acts"><Button size="sm" onclick={share}>{plannerCopy.retry}</Button></div>
      </div>
    {/if}

    {#if saved}
      <div class="pl-confirm">
        <div class="pl-acts">
          <a href={saved.url} class="pl-url" data-testid="share-link">{saved.url}</a>
          <Button size="sm" onclick={() => copyToClipboard(saved.url, 'link')}
            >{copier.copied('link') ? plannerCopy.copied : plannerCopy.copyLink}</Button
          >
        </div>
        {#if cardState !== 'idle'}
          <p class="text-small text-muted" data-testid="build-sim-status">
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
            class="pl-card"
            onerror={() => (cardBroken = true)}
          />
        {/if}
      </div>
    {/if}
  </div>
{/snippet}

{#if part}
  <PlannerPart title={plannerCopy.shareTitle} testid="share-panel">{@render body()}</PlannerPart>
{:else}
  <Section title={plannerCopy.shareTitle} panel testid="share-panel">{@render body()}</Section>
{/if}
