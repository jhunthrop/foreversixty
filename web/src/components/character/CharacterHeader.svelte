<!-- web/src/components/character/CharacterHeader.svelte -->
<!-- THE shared character header (design/specs/2026-10-11-character-header.md): one panel at the
     top of a page, signed in only -- WHO (crest, label, name, descriptor, sync line), THE PAGE'S
     ANSWER about that character, and THE PAGE'S ONE ACTION. It follows the header selector in
     place and has no switcher of its own. Mounted by the home signed-in panel, the BiS landing,
     the BiS spec page and the planner (the last two as the stacked `card` layout with the outline
     action); Rankings and Logs mount it the same way in later lanes.

     FIXED by this component (a page never restyles or re-words these): the panel, the identity
     (crest 64 px, 56 stacked; label always "Your character"; name 24 px, 22 stacked; the
     descriptor "Gnome Fury Warrior · Level 60 · [flat faction logo] Alliance", the race only when
     the API sent it, split into two fixed lines below 1280 px and in the card layout so a
     separator never starts or ends a line), the sync line ("Addon · 19 days ago · Log out in the
     game to sync", only the words past the age in ember, two-line cap), and every state: loading
     (the ready box hidden under skeletons, exactly the ready height), answer loading, stale, no
     gear on file, under the page's level floor, empty, failed, and the one shared signed-out
     band (a single muted line, no button: the nav selector is the only sign-in control).

     THE API (props):
       action        HeaderAction | ((ctx: { character, belowLevelFloor }) => HeaderAction)
                     The page's one action: { label, href } opens a page or jumps to an anchor;
                     { kind: 'picker', label, href } renders a button that calls `onPick`
                     (Logs: the file picker). The identity link goes to the same href.
       answer        AnswerSpec (lib/character-header/types.ts), optional -- omit it for a page
                     with no answer (the planner). { levelFloor?, needsGear?, strings:
                     { belowLevelFloor?, noGear?: { sentence, link }, failed }, load(character) }.
                     `load` resolves to a HeaderAnswer: { kind: 'figures', label, staleLabel?,
                     figures: [primary, second?], note?: { text, link? }, secondNote?, bar? } (the
                     primary figure and an optional second at the same size; a bar only when the
                     figure is a fraction of a whole), { kind: 'sentence', text } (the empty
                     answer) or { kind: 'none' }. It is re-run, under a skeleton, whenever the
                     selected character changes; a rejection is the failed state with "Try again".
       emphasis      'primary' (a landing page's lead: the filled 48 px gold button, default) or
                     'secondary' (a tool page's context card: the 44 px outline button).
       layout        'band' (full width, three columns from 1024 px, default) or 'card' (always
                     stacked, for a side column).
       whenSignedOut 'band' (the shared signed-out band, default) or 'none' (the home page, whose
                     own signed-out hero is the front door).
       source        a CharacterCardState, to share one session read with the page; default: own.
       onCharacterChange  called with the selected character, or null once the session answered
                     signed out, and again on every selector change.
       onPick        the picker action's handler.

     Server-rendered it is the loading state, with the skeleton shown only when `<html
     data-session>` hints a session (global.css `.ch-hint-only`) and the signed-out band only
     when it does not (`.ch-hint-out-only`), so neither kind of visitor sees the panel resize. -->
<script lang="ts">
  import { untrack } from 'svelte';
  import type { MeCharacter } from '../../lib/account/api';
  import {
    createCharacterCardState,
    type CharacterCardState,
  } from '../../lib/bis/character-card-state.svelte';
  import { classCrestSrc } from '../../lib/class-crest';
  import { answerViewFor, type AnswerPhase, type AnswerView } from '../../lib/character-header/answer-view';
  import { characterHeaderCopy } from '../../lib/character-header/copy';
  import { accountNameOf } from '../../lib/character-selector/rows';
  import { descriptorLines } from '../../lib/character-header/descriptor';
  import {
    PLACEHOLDER_CHARACTER,
    PLACEHOLDER_GEARLESS_CHARACTER,
    placeholderAnswer,
  } from '../../lib/character-header/placeholder';
  import { rememberHeaderGear } from '../../lib/character-header/gear-hint';
  import { planAnswer } from '../../lib/character-header/answer-plan';
  import {
    STALE_SEPARATOR,
    isStaleAddonSync,
    syncText,
    syncView,
    type SyncView,
  } from '../../lib/character-header/sync';
  import type {
    ActionContext,
    AnswerSpec,
    HeaderAction,
    HeaderEmphasis,
    HeaderLayout,
    SignedOutMode,
  } from '../../lib/character-header/types';
  import { factionLogoSmallSrc, factionName } from '../../lib/faction-mark';
  import { classColorVar } from '../../lib/report/format';
  import { classSlugFromName } from '../../lib/report/tree-sizes';
  import { finishedRunsFor } from '../../lib/character-header/finished-sim-line';
  import SupporterChip from '../SupporterChip.svelte';
  import FinishedSimLine from './FinishedSimLine.svelte';
  import CharacterHeaderSignedOut from './CharacterHeaderSignedOut.svelte';

  type ActionProp = HeaderAction | ((context: ActionContext) => HeaderAction);

  let {
    action,
    answer,
    emphasis = 'primary',
    layout = 'band',
    whenSignedOut = 'band',
    nameHeading = false,
    source,
    onCharacterChange,
    onPick,
    onSecondary,
  }: {
    action: ActionProp;
    answer?: AnswerSpec;
    emphasis?: HeaderEmphasis;
    layout?: HeaderLayout;
    whenSignedOut?: SignedOutMode;
    /** Wraps the name in an `<h1>`: the signed-in home page has no other heading. */
    nameHeading?: boolean;
    source?: CharacterCardState;
    onCharacterChange?: (character: MeCharacter | null) => void;
    onPick?: () => void;
    /** The picker action's second button, when it is a press rather than a link. */
    onSecondary?: () => void;
  } = $props();

  const session = untrack(() => source) ?? createCharacterCardState();
  const character = $derived(session.character);
  const isSupporter = $derived(session.me?.entitlements?.supporter_mark === true);
  // The band's "your server sim finished" footer (spec 2026-10-24 3.4): the band layout only,
  // and not once dismissed on this page.
  let finishedDismissed = $state(false);
  const finishedRuns = $derived(layout === 'band' && !finishedDismissed ? finishedRunsFor(session.me) : null);
  let actionEl = $state<HTMLElement | null>(null);

  function onFinishedDismissed(): void {
    finishedDismissed = true;
    actionEl?.focus();
  }

  const view = $derived<'loading' | 'signed-out' | 'no-characters' | 'ready'>(
    character !== null
      ? 'ready'
      : session.status === 'idle' || session.status === 'loading'
        ? 'loading'
        : session.me !== null && session.status === 'ready'
          ? 'no-characters'
          : 'signed-out',
  );

  // The selected character's clock-dependent reads (sync age, stale) are taken when the
  // character changes, not on a timer: the header never re-renders by itself.
  const sync = $derived(character === null ? null : syncView(character, new Date()));
  const stale = $derived(character !== null && isStaleAddonSync(character, new Date()));

  let phase = $state<AnswerPhase>({ phase: 'loading' });
  let attempt = $state(0);

  $effect(() => {
    const current = character;
    const spec = answer;
    void attempt;
    if (current === null || spec === undefined || planAnswer(current, spec) !== 'load') return;
    phase = { phase: 'loading' };
    let cancelled = false;
    spec
      .load(current)
      .then((loaded) => {
        if (!cancelled) phase = { phase: 'ready', answer: loaded };
      })
      .catch(() => {
        if (!cancelled) phase = { phase: 'failed' };
      });
    return () => {
      cancelled = true;
    };
  });

  $effect(() => {
    if (view === 'loading') return;
    const current = character;
    untrack(() => onCharacterChange?.(current));
  });

  const answerView = $derived(character === null ? null : answerViewFor(character, answer, phase, stale));
  const belowLevelFloor = $derived(
    character !== null && answer !== undefined && planAnswer(character, answer) === 'below-floor',
  );
  const resolvedAction = $derived.by((): HeaderAction | null => {
    if (character === null) return null;
    return typeof action === 'function' ? action({ character, belowLevelFloor }) : action;
  });

  const hasAnswerColumn = $derived(answer !== undefined && (view === 'loading' || answerView !== null));

  const PLACEHOLDER_SYNC: SyncView = { source: 'Addon', age: '12 minutes ago', staleTail: null };
  const PLACEHOLDER_ACTION: HeaderAction = { label: ' ', href: '#' };
  type FiguresView = Extract<AnswerView, { kind: 'figures' }>;
  const placeholderView = $derived.by((): AnswerView | null => {
    if (answer === undefined) return null;
    const ghostAnswer = placeholderAnswer(answer.placeholderBar === true);
    return { kind: 'figures', answer: ghostAnswer, label: ghostAnswer.label };
  });
  // A character with no gear on file draws no sync line and, on a page whose answer needs gear,
  // the "add the addon" sentence: the gearless skeleton reserves exactly that box.
  const gearlessPlaceholderView = $derived.by((): AnswerView | null => {
    const view = answerViewFor(PLACEHOLDER_GEARLESS_CHARACTER, answer, { phase: 'loading' }, false);
    return view?.kind === 'loading' ? placeholderView : view;
  });

  // The next page's skeleton takes the shape of the header drawn now (lib/character-header/gear-hint.ts).
  $effect(() => {
    if (view !== 'ready' || character === null) return;
    rememberHeaderGear(character, localStorage);
  });
</script>

<!-- The loading state draws the ready box once, hidden, to reserve its height. Ghost text is
     painted from `data-ghost` by a pseudo-element (.ch-ghost::before), so the reserved box keeps
     the ready line boxes exactly while the served HTML carries no text a crawler or screen
     reader could read. -->
{#snippet txt(value: string, ghost: boolean)}
  {#if ghost}<span class="ch-ghost" data-ghost={value}></span>{:else}{value}{/if}
{/snippet}

{#snippet identity(c: MeCharacter, href: string, syncLine: SyncView | null, ghost: boolean)}
  {@const lines = descriptorLines(c)}
  {@const slug = classSlugFromName(c.class ?? '')}
  {@const color = classColorVar(c.class)}
  <div class="ch-id">
    <a class="ch-link" {href} data-testid={ghost ? undefined : 'character-header-link'}>
      {#if slug !== ''}
        <img
          class="ch-crest"
          src={classCrestSrc(slug)}
          alt=""
          width="64"
          height="64"
          decoding="async"
          style:--c={color}
        />
      {/if}
      <span class="ch-ident">
        <span class="label ch-label">{@render txt(characterHeaderCopy.label, ghost)}</span>
        <span class="ch-name-box">
          <svelte:element
            this={nameHeading && !ghost ? 'h1' : 'span'}
            class="ch-name"
            title={ghost ? undefined : c.name}
            style:color
            data-testid={ghost ? undefined : 'character-header-name'}
            >{@render txt(c.name, ghost)}</svelte:element
          >
          {#if isSupporter && !ghost}<SupporterChip />{/if}
        </span>
      </span>
    </a>
    <div class="ch-meta">
      <span class="ch-desc" data-testid={ghost ? undefined : 'character-header-descriptor'}>
        {#if lines.identity !== ''}<span class="ch-d1">{@render txt(lines.identity, ghost)}</span>{/if}
        {#if lines.identity !== '' && (lines.level !== null || lines.faction !== null)}
          <span class="ch-dot ch-dot-wide" aria-hidden="true">&middot;</span>
        {/if}
        {#if lines.level !== null || lines.faction !== null}
          <span class="ch-d2">
            {#if lines.level !== null}<span>{@render txt(lines.level, ghost)}</span>{/if}
            {#if lines.level !== null && lines.faction !== null}
              <span class="ch-dot" aria-hidden="true">&middot;</span>
            {/if}
            {#if lines.faction !== null}
              <span class="ch-faction">
                <img src={factionLogoSmallSrc(lines.faction)} alt="" width="16" height="16" />
                {@render txt(factionName(lines.faction), ghost)}
              </span>
            {/if}
          </span>
        {/if}
      </span>
      {#if syncLine !== null}
        <span class="ch-sync" data-testid={ghost ? undefined : 'character-header-sync'}
          >{@render txt(syncText(syncLine), ghost)}{#if syncLine.staleTail !== null}{STALE_SEPARATOR}<span
              class="ch-stale"
              data-testid={ghost ? undefined : 'character-header-stale'}>{syncLine.staleTail}</span
            >{/if}</span
        >
      {/if}
    </div>
  </div>
{/snippet}

{#snippet figuresBody(current: FiguresView, ghost: boolean)}
  <span class="label ch-answer-label" data-testid={ghost ? undefined : 'character-header-answer-label'}
    >{@render txt(current.label, ghost)}</span
  >
  <span class="ch-figures">
    {#each current.answer.figures as figure, index (index)}
      <span
        class="ch-figure"
        data-testid={ghost ? undefined : `character-header-figure-${index + 1}`}
        style:color={figure.color}
        >{@render txt(figure.value, ghost)}{#if figure.unit !== undefined}<span class="ch-unit"
            >{` ${figure.unit}`}</span
          >{/if}</span
      >
    {/each}
  </span>
  {#if current.answer.bar !== undefined}
    <span class="ch-track" aria-hidden="true">
      <span class="ch-fill" style:width={`${(current.answer.bar * 100).toFixed(1)}%`}></span>
    </span>
  {/if}
  {#if current.answer.note !== undefined}
    <span class="ch-note" data-testid={ghost ? undefined : 'character-header-note'}
      >{@render txt(current.answer.note.text, ghost)}{#if current.answer.note.link !== undefined}
        &nbsp;&middot;
        <a class="ch-inline-link" href={current.answer.note.link.href}
          >{current.answer.note.link.label}&nbsp;→</a
        >{/if}</span
    >
  {/if}
  {#if current.answer.secondNote !== undefined}
    <a
      class="ch-second-note"
      href={current.answer.secondNote.href}
      data-testid={ghost ? undefined : 'character-header-second-note'}>{current.answer.secondNote.text}</a
    >
  {/if}
{/snippet}

{#snippet answerSlot(current: AnswerView, ghost: boolean)}
  <div
    class="ch-answer"
    data-testid={ghost ? undefined : 'character-header-answer'}
    data-answer={current.kind}
  >
    {#if current.kind === 'loading'}
      <!-- The figures it is waiting for, hidden, so the answer column already has their height. -->
      {#if placeholderView !== null && placeholderView.kind === 'figures'}
        <span class="ch-hidden" aria-hidden="true">{@render figuresBody(placeholderView, true)}</span>
      {/if}
      <span
        class="ch-ask ch-ask-over"
        data-testid={ghost ? undefined : 'character-header-answer-loading'}
        aria-hidden="true"
      >
        <span class="skeleton-block ch-sk-label"></span>
        <span class="skeleton-block ch-sk-figure"></span>
        <span class="skeleton-block ch-sk-note"></span>
      </span>
    {:else if current.kind === 'figures'}
      {@render figuresBody(current, ghost)}
    {:else if current.kind === 'sentence'}
      <span class="ch-sentence" data-testid={ghost ? undefined : 'character-header-sentence'}
        >{current.text}</span
      >
    {:else if current.kind === 'no-gear'}
      <span class="ch-small" data-testid={ghost ? undefined : 'character-header-no-gear'}>
        {@render txt(current.sentence, ghost)}
        <a class="ch-text-link" href={current.link.href}
          >{@render txt(`${current.link.label}\u00a0→`, ghost)}</a
        >
      </span>
    {:else}
      <span class="ch-small" data-testid={ghost ? undefined : 'character-header-failed'}>
        {current.text}
        <button
          type="button"
          class="ch-text-link"
          onclick={() => (attempt += 1)}
          data-testid={ghost ? undefined : 'character-header-retry'}>{characterHeaderCopy.tryAgain}</button
        >
      </span>
    {/if}
  </div>
{/snippet}

{#snippet button(act: HeaderAction, ghost: boolean)}
  {#if act.kind === 'picker'}
    <button
      bind:this={actionEl}
      type="button"
      class="ch-cta"
      disabled={act.disabled === true}
      onclick={() => onPick?.()}
      data-testid={ghost ? undefined : 'character-header-action'}>{@render txt(act.label, ghost)}</button
    >
    {#if act.secondary !== undefined}
      {#if act.secondary.href !== undefined}
        <a
          class="ch-cta ch-cta-second"
          href={act.secondary.href}
          data-testid={act.secondary.testid ?? 'character-header-secondary'}>{act.secondary.label}</a
        >
      {:else}
        <button
          type="button"
          class="ch-cta ch-cta-second"
          onclick={() => onSecondary?.()}
          data-testid={act.secondary.testid ?? 'character-header-secondary'}>{act.secondary.label}</button
        >
      {/if}
    {/if}
    {#if act.note !== undefined}
      <p
        class="ch-note"
        class:ch-note-error={act.note.tone === 'error'}
        role={act.note.tone === 'error' ? 'alert' : 'status'}
        data-testid={ghost ? undefined : 'character-header-note'}
      >
        {act.note.text}
        {#if act.note.link !== undefined}
          <a
            class="ch-note-link"
            href={act.note.link.href}
            data-testid={ghost ? undefined : 'character-header-note-link'}>{act.note.link.label}</a
          >
        {/if}
        {#if act.note.detail !== undefined}
          <span class="ch-note-detail" data-testid={ghost ? undefined : 'character-header-note-detail'}
            >{act.note.detail}</span
          >
        {/if}
      </p>
    {/if}
  {:else}
    <a
      bind:this={actionEl}
      class="ch-cta"
      href={act.href}
      data-testid={ghost ? undefined : 'character-header-action'}>{@render txt(act.label, ghost)}</a
    >
  {/if}
{/snippet}

{#snippet panel(
  c: MeCharacter,
  current: AnswerView | null,
  act: HeaderAction,
  syncLine: SyncView | null,
  ghost: boolean,
)}
  {@render identity(c, act.href, syncLine, ghost)}
  {#if current !== null}{@render answerSlot(current, ghost)}{/if}
  {@render button(act, ghost)}
{/snippet}

{#if view === 'ready' && character !== null && resolvedAction !== null}
  <section
    class="ch card-surface ch-{layout} ch-{emphasis}"
    class:ch-no-answer={!hasAnswerColumn}
    class:ch-has-foot={finishedRuns !== null}
    aria-label={characterHeaderCopy.label}
    data-testid="character-header"
    data-layout={layout}
    data-emphasis={emphasis}
    data-stale={stale}
  >
    {@render panel(character, answerView, resolvedAction, sync, false)}
    {#if finishedRuns !== null}
      <div class="ch-foot"><FinishedSimLine runs={finishedRuns} onDismissed={onFinishedDismissed} /></div>
    {/if}
  </section>
{:else if view === 'no-characters'}
  {#if whenSignedOut === 'band'}<CharacterHeaderSignedOut account={accountNameOf(session.me)} />{/if}
{:else if view === 'signed-out'}
  {#if whenSignedOut === 'band'}<CharacterHeaderSignedOut />{/if}
{:else if view === 'loading'}
  <section
    class="ch card-surface ch-{layout} ch-{emphasis} ch-hint-only ch-loading"
    class:ch-no-answer={!hasAnswerColumn}
    role="status"
    aria-busy="true"
    aria-label={characterHeaderCopy.loading}
    data-testid="character-header-loading"
    data-layout={layout}
  >
    <!-- Two ghosts, one shown: the geared box (sync line, figures) when the pre-paint script
         hinted data-ch-gear, else the gearless box (lib/character-header/gear-hint.ts). -->
    <div class="ch-hidden ch-ghost-geared" aria-hidden="true">
      {@render panel(PLACEHOLDER_CHARACTER, placeholderView, PLACEHOLDER_ACTION, PLACEHOLDER_SYNC, true)}
    </div>
    <div class="ch-hidden ch-ghost-gearless" aria-hidden="true">
      {@render panel(PLACEHOLDER_GEARLESS_CHARACTER, gearlessPlaceholderView, PLACEHOLDER_ACTION, null, true)}
    </div>
    <div class="ch-overlay" aria-hidden="true">
      <span class="ch-sk-identity">
        <span class="skeleton-block ch-sk-crest"></span>
        <span class="ch-sk-text">
          <span class="skeleton-block ch-sk-label"></span>
          <span class="skeleton-block ch-sk-name"></span>
          <span class="skeleton-block ch-sk-meta"></span>
        </span>
      </span>
      {#if answer !== undefined}
        <span class="ch-ask">
          <span class="skeleton-block ch-sk-label"></span>
          <span class="skeleton-block ch-sk-figure"></span>
          <span class="skeleton-block ch-sk-note"></span>
        </span>
      {/if}
      <span class="skeleton-block ch-sk-cta"></span>
    </div>
    {#if layout === 'band'}<div class="ch-foot ch-foot-reserve" aria-hidden="true"></div>{/if}
  </section>
  {#if whenSignedOut === 'band'}
    <div class="ch-hint-out-only"><CharacterHeaderSignedOut /></div>
  {/if}
{/if}

<style>
  .ch {
    --ch-cols: minmax(0, 1fr);
    --ch-gap: 16px;
    --ch-pad-x: 16px;
    --ch-pad-y: 16px;
    --ch-foot-h: 44px;
    --ch-foot-phone-h: 56px;
    --ch-band-min: 136px;
    --ch-crest: 56px;
    --ch-name-size: 22px;
    position: relative;
    display: grid;
    grid-template-columns: var(--ch-cols);
    gap: var(--ch-gap);
    align-items: center;
    box-sizing: border-box;
    padding: var(--ch-pad-y) var(--ch-pad-x);
    /* The surface is the shared card-surface class on the element (site alignment spec
       2026-10-29 S3): the account card's warm gradient, gold-tinted border and top edge. Its
       emphasis comes from the crest, the class-coloured name and its one primary button. */
  }
  @media (min-width: 1024px) {
    .ch-band {
      --ch-cols: minmax(0, 1.3fr) minmax(0, 1fr) auto;
      --ch-gap: 28px;
      --ch-pad-x: 22px;
      --ch-pad-y: 18px;
      --ch-crest: 64px;
      --ch-name-size: 24px;
      /* The tallest state (a two-line stale sync under a two-line identity), so no state
         changes the band's height. */
      min-height: var(--ch-band-min);
    }
    .ch-band.ch-no-answer {
      --ch-cols: minmax(0, 1fr) auto;
    }
  }
  @media (min-width: 1280px) {
    .ch-band {
      --ch-cols: minmax(0, 1.15fr) minmax(0, 1fr) auto;
    }
    .ch-band.ch-no-answer {
      --ch-cols: minmax(0, 1fr) auto;
    }
  }

  /* ---- the done line's footer (FinishedSimLine) ---- */
  /* Full width, flush inside the border: it cancels the panel's padding on three sides and
     the gap above it, so the 136 px of the grid above stays exactly the grid's own. */
  .ch-foot {
    grid-column: 1 / -1;
    margin: calc(var(--ch-pad-y) - var(--ch-gap)) calc(-1 * var(--ch-pad-x)) calc(-1 * var(--ch-pad-y));
    border-radius: 0 0 calc(var(--radius-panel) - 1px) calc(var(--radius-panel) - 1px);
    overflow: hidden;
  }
  /* The reserve (3.7): the skeleton holds the row the line will fill. Hidden unless the
     pre-paint script stamped data-finished (global.css shows it and sizes the skeleton). */
  .ch-foot-reserve {
    display: none;
    min-height: var(--ch-foot-phone-h);
  }
  @media (min-width: 600px) {
    .ch-foot-reserve {
      min-height: var(--ch-foot-h);
    }
  }
  @media (min-width: 1024px) {
    .ch-band.ch-has-foot,
    :global(html[data-finished='1']) .ch-band.ch-loading {
      min-height: calc(var(--ch-band-min) + var(--ch-foot-h));
      grid-template-rows: 1fr auto;
    }
  }

  /* ---- identity ---- */
  .ch-id {
    display: flex;
    flex-direction: column;
    gap: 6px;
    min-width: 0;
  }
  .ch-link {
    display: flex;
    align-items: center;
    gap: 16px;
    min-width: 0;
    color: inherit;
  }
  .ch-link:hover {
    color: inherit;
  }
  .ch-link:hover .ch-name {
    text-decoration: underline;
    text-underline-offset: 4px;
  }
  .ch-link:focus-visible {
    outline: 2px solid var(--color-gold);
    outline-offset: 4px;
    border-radius: var(--radius-panel);
  }
  .ch-crest {
    flex-shrink: 0;
    width: var(--ch-crest);
    height: var(--ch-crest);
    border-radius: 999px;
    object-fit: cover;
    background: var(--color-raised);
    box-shadow: 0 0 0 2px color-mix(in srgb, var(--c) 55%, transparent);
  }
  .ch-ident {
    display: flex;
    flex-direction: column;
    gap: 3px;
    min-width: 0;
  }
  .ch-label {
    color: var(--color-gold);
  }
  .ch-ghost::before {
    content: attr(data-ghost);
  }
  /* A stacked header reserves two name lines so a header switch moves nothing. */
  .ch-name-box {
    display: flex;
    align-items: flex-end;
    /* The chip wraps under the name rather than squeezing it: the name keeps its full width. */
    flex-wrap: wrap;
    gap: 4px 12px;
    align-content: flex-end;
    min-height: calc(var(--ch-name-size) * 1.2 * 2);
  }
  @media (min-width: 1024px) {
    .ch-band .ch-name-box {
      min-height: 0;
    }
  }
  .ch-name {
    max-width: 100%;
    margin: 0;
    font-family: var(--font-heading);
    font-size: var(--ch-name-size);
    font-weight: 800;
    letter-spacing: var(--tracking-heading);
    line-height: 1.2;
    overflow: hidden;
    display: -webkit-box;
    -webkit-line-clamp: 2;
    line-clamp: 2;
    -webkit-box-orient: vertical;
    overflow-wrap: anywhere;
  }
  .ch-meta {
    display: flex;
    flex-direction: column;
    gap: 3px;
    padding-left: calc(var(--ch-crest) + 16px);
  }
  .ch-desc {
    display: flex;
    flex-direction: column;
    gap: 2px;
    font-size: 13px;
    color: #c9c2b2;
  }
  .ch-d2 {
    display: inline-flex;
    align-items: center;
    gap: 6px;
    white-space: nowrap;
  }
  .ch-dot {
    color: var(--color-muted);
  }
  /* The separator between the two lines exists only where the descriptor is one line. */
  .ch-dot-wide {
    display: none;
  }
  .ch-faction {
    display: inline-flex;
    align-items: center;
    gap: 6px;
  }
  .ch-faction img {
    width: 16px;
    height: 16px;
    flex-shrink: 0;
    object-fit: contain;
  }
  @media (min-width: 1280px) {
    .ch-band .ch-desc {
      flex-direction: row;
      align-items: center;
      gap: 6px;
      white-space: nowrap;
    }
    .ch-band .ch-dot-wide {
      display: inline;
    }
  }
  .ch-sync {
    display: -webkit-box;
    -webkit-line-clamp: 2;
    line-clamp: 2;
    -webkit-box-orient: vertical;
    overflow: hidden;
    font-size: 12px;
    color: var(--color-muted);
  }
  .ch-stale {
    color: var(--color-ember);
  }

  /* ---- answer ---- */
  .ch-answer {
    display: flex;
    flex-direction: column;
    align-items: flex-start;
    gap: 6px;
    min-width: 0;
  }
  .ch-answer-label {
    color: var(--color-muted);
  }
  .ch-figures {
    display: flex;
    flex-wrap: wrap;
    align-items: baseline;
    gap: 6px 18px;
    font-family: var(--font-body);
    font-variant-numeric: tabular-nums;
    font-size: 26px;
    font-weight: 500;
    line-height: 1;
    color: var(--color-text);
  }
  .ch-unit {
    font-size: 15px;
    color: var(--color-muted);
  }
  .ch-track {
    align-self: stretch;
    height: 6px;
    border-radius: 3px;
    background: var(--color-line-soft);
    overflow: hidden;
  }
  .ch-fill {
    display: block;
    height: 100%;
    border-radius: 3px;
    background: var(--color-gold);
  }
  .ch-note {
    font-size: 12px;
    color: var(--color-muted);
  }
  .ch-inline-link,
  .ch-text-link {
    font-weight: 700;
    color: var(--color-gold);
    white-space: nowrap;
  }
  .ch-inline-link {
    font-size: 12px;
  }
  .ch-inline-link:hover,
  .ch-text-link:hover {
    color: var(--color-gold-hover);
    text-decoration: underline;
    text-underline-offset: 4px;
  }
  .ch-second-note {
    font-size: 12px;
    color: var(--color-muted);
    text-decoration: underline;
    text-underline-offset: 4px;
  }
  .ch-sentence {
    font-size: 14px;
    color: var(--color-text);
  }
  .ch-small {
    display: flex;
    flex-direction: column;
    align-items: flex-start;
    gap: 2px;
    font-size: 13px;
    color: #c9c2b2;
  }
  .ch-text-link {
    display: inline-flex;
    align-items: center;
    min-height: 44px;
    padding: 0;
    border: 0;
    background: none;
    font-family: inherit;
    font-size: 13px;
  }

  /* ---- action ---- */
  .ch-cta {
    display: inline-flex;
    align-items: center;
    justify-content: center;
    box-sizing: border-box;
    width: 100%;
    padding: 0 20px;
    border-radius: var(--radius-control);
    font-family: inherit;
    font-weight: 700;
    letter-spacing: 0.06em;
    text-transform: uppercase;
    white-space: nowrap;
    transition:
      background 120ms ease-out,
      border-color 120ms ease-out,
      color 120ms ease-out;
  }
  .ch-cta:focus-visible {
    outline: 2px solid var(--color-gold);
    outline-offset: 2px;
  }
  .ch-primary .ch-cta {
    height: 48px;
    border: 1px solid var(--color-gold-deep);
    background: linear-gradient(180deg, #f0cc6c, #c99a3a);
    box-shadow: 0 0 18px rgba(229, 185, 85, 0.18);
    color: #1a1408;
    font-size: 13px;
  }
  .ch-primary .ch-cta:hover {
    background: linear-gradient(180deg, #f5d27a, #d4a444);
    color: #1a1408;
  }
  .ch-primary .ch-cta:active {
    background: linear-gradient(180deg, #d9b55c, #b88a32);
  }
  .ch-secondary .ch-cta {
    height: 44px;
    border: 1px solid var(--color-gold);
    background: transparent;
    color: var(--color-gold);
    font-size: 12px;
  }
  .ch-secondary .ch-cta:hover {
    border-color: var(--color-gold-hover);
    color: var(--color-gold-hover);
    background: color-mix(in srgb, var(--color-gold) 8%, transparent);
  }
  /* The second button sits 8px under the first (the grid gap is 16px), same width, same height,
     quieter: the outline stays, the fill and glow go. */
  .ch-cta-second {
    margin-top: -8px;
    text-decoration: none;
  }
  .ch-primary .ch-cta-second {
    height: 44px;
    border: 1px solid var(--color-gold-deep);
    background: transparent;
    box-shadow: none;
    color: var(--color-gold);
    font-size: 12px;
  }
  .ch-primary .ch-cta-second:hover {
    background: color-mix(in srgb, var(--color-gold) 8%, transparent);
    color: var(--color-gold-hover);
  }
  .ch-cta:disabled {
    cursor: default;
    opacity: 0.6;
  }
  .ch-secondary .ch-cta:disabled:hover {
    border-color: var(--color-gold);
    color: var(--color-gold);
    background: transparent;
  }
  .ch-note {
    margin: 8px 0 0;
    font-size: 12px;
    line-height: 1.5;
    color: var(--color-muted);
  }
  .ch-note-detail {
    display: block;
    margin-top: 2px;
    opacity: 0.85;
  }
  .ch-note-error {
    color: var(--color-ember, #d66e28);
  }
  .ch-note-link {
    color: var(--color-gold);
    text-decoration: underline;
    text-underline-offset: 2px;
  }
  .ch-secondary .ch-cta:active {
    background: color-mix(in srgb, var(--color-gold) 4%, transparent);
  }
  @media (min-width: 1024px) {
    .ch-band .ch-cta {
      width: auto;
      min-width: 210px;
    }
    .ch-band.ch-secondary .ch-cta {
      min-width: 220px;
    }
  }
  @media (min-width: 1280px) {
    .ch-band.ch-primary .ch-cta {
      min-width: 250px;
    }
  }

  /* ---- loading: the ready box hidden, skeletons drawn over it ---- */
  .ch-hidden {
    display: contents;
    visibility: hidden;
  }
  :global(html:not([data-ch-gear='1'])) .ch-ghost-geared,
  :global(html[data-ch-gear='1']) .ch-ghost-gearless {
    display: none;
  }
  /* The answer's own loading state: the hidden figures hold the column, the skeleton sits on top. */
  .ch-answer:has(> .ch-ask-over) {
    position: relative;
  }
  .ch-ask-over {
    position: absolute;
    inset: 0 0 auto;
  }
  .ch-overlay {
    position: absolute;
    inset: 0 0 var(--ch-foot-reserve, 0);
    display: grid;
    grid-template-columns: var(--ch-cols);
    gap: var(--ch-gap);
    align-items: center;
    box-sizing: border-box;
    padding: var(--ch-pad-y) var(--ch-pad-x);
  }
  .ch-sk-identity {
    display: flex;
    align-items: center;
    gap: 16px;
  }
  .ch-sk-text,
  .ch-ask {
    display: flex;
    flex-direction: column;
    gap: 8px;
  }
  .ch-sk-crest {
    flex-shrink: 0;
    width: var(--ch-crest);
    height: var(--ch-crest);
    border-radius: 999px;
  }
  .ch-sk-label {
    width: 150px;
    max-width: 100%;
    height: 11px;
  }
  .ch-sk-name {
    width: 200px;
    max-width: 100%;
    height: 24px;
  }
  .ch-sk-meta {
    width: 240px;
    max-width: 100%;
    height: 13px;
  }
  .ch-sk-figure {
    width: 150px;
    max-width: 100%;
    height: 26px;
  }
  .ch-sk-note {
    width: 100%;
    height: 12px;
  }
  .ch-sk-cta {
    width: 100%;
    height: 48px;
    border-radius: var(--radius-control);
  }
  @media (min-width: 1024px) {
    .ch-band .ch-sk-cta {
      width: 210px;
    }
  }
  @media (min-width: 1280px) {
    .ch-band .ch-sk-cta {
      width: 250px;
    }
  }
  .ch-secondary .ch-sk-cta {
    height: 44px;
  }
</style>
