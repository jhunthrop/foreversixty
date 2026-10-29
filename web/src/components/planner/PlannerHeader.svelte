<!-- web/src/components/planner/PlannerHeader.svelte -->
<!-- Rebuild spec §4.A, fix round 1 (the approved mock, `day3/shots/boards/Planner.png`):
     the planner's own class-page header -- eyebrow, 64px crest, h1 with the always-present
     level/race suffix (plus a spec clause once one tree leads), a one-line description, the
     class/race selects, and the facts rail (Points left/Spent/Level/DPS) as the header's own
     bottom row. Lives inside the Planner island (not a `.astro` sibling like the BiS/home
     rebuild's `ClassHeader.astro`) because every one of these facts -- the class, the race,
     the level, the leading spec, the live build state -- is the island's own reactive store,
     not build-time frontmatter; `ClassHeader.astro`'s own markup is server-rendered once per
     faction and has no runtime state to read at all (see that component's header comment).

     Fix round 5: this component is now the band's own LEFT COLUMN content only -- no
     background, no border, no radius, no art. `PlannerHeaderBand.svelte` (its only caller)
     owns all of that at the band level, full-bleed across both grid columns (this header's
     own content and the character header's), which a box scoped to this component's own
     column could never reach. The crest still inlines `ClassCrest.astro`'s own recipe (a
     Svelte file cannot import an `.astro` one at all, the same constraint
     `CharacterHeader.svelte` documents). -->
<script lang="ts">
  import { classCrestSrc } from '../../lib/class-crest';
  import { classColorVar } from '../../lib/report/format';
  import { plannerCharacterCardCopy, plannerHeaderCopy } from '../../lib/planner/copy';
  import type { LiveDps } from '../../lib/planner/live-dps.svelte';
  import type { LiveGate } from '../../lib/planner/live-gate';
  import type { PlannerStore } from '../../lib/planner/store.svelte';
  import { specOf } from '../../lib/sim/character';
  import { specLabel } from '../../lib/sim/spec-label';
  import { FACT_FIGURE_CLASS } from '../../lib/planner/styles';
  import PlannerDps from './PlannerDps.svelte';
  import StatTile from '../ui/StatTile.svelte';
  import Select from '../ui/Select.svelte';
  import { scoreUnitForBuild } from '../../lib/planner/score-unit';

  let {
    store,
    live,
    gate,
    onshowdps,
    origin = null,
  }: {
    store: PlannerStore;
    live: LiveDps;
    gate: LiveGate;
    onshowdps: () => void;
    /** Set when the build on screen is another player's: whose it is ('' for a link that names nobody). */
    origin?: { owner: string } | null;
  } = $props();

  const color = $derived(classColorVar(store.classRow?.name));

  /** True once one tree holds strictly more points than each of the other two --
   *  `specKeyFor`'s own tie-break, reused unchanged. A tie, including 0/0/0, is false: the
   *  h1's third clause only ever appears once a build genuinely leans one way. */
  const hasLeadingTree = $derived.by(() => {
    const split = store.split;
    if (split.length === 0) return false;
    const max = Math.max(...split);
    return max > 0 && split.filter((points) => points === max).length === 1;
  });

  /** "Ashenveil's build · Holy Paladin": whose build this is, when it is not the player's own. */
  const originLine = $derived(
    origin === null
      ? ''
      : plannerCharacterCardCopy.originLine(
          origin.owner,
          hasLeadingTree && store.talentIndex !== null
            ? specLabel(specOf(store.talentIndex, [...store.order]))
            : (store.classRow?.name ?? store.classSlug),
        ),
  );

  /** The h1's second half, one entry per segment: race, level, and the spec once a tree leads. */
  const suffixParts = $derived([
    store.raceRow?.name ?? '',
    `Level ${store.level}`,
    ...(hasLeadingTree && store.talentIndex !== null
      ? [specLabel(specOf(store.talentIndex, [...store.order]))]
      : []),
  ]);

  // The bar's one refusal line already carries the planner's own refusals (an illegal move,
  // a read-only build); a failed live estimate is the same kind of fact, so it goes in the
  // same line rather than a second one (SummaryBar.svelte's own reasoning, moved here with
  // the rest of the facts rail).
  const statusLine = $derived(store.refusal ?? (live.state === 'error' ? live.message : null) ?? '');
  const pointsLeft = $derived(Math.max(0, store.pointsAvailable - store.spent));
  const scoreUnit = $derived(scoreUnitForBuild(store.classSlug, store.split));
  // `Select` binds a value of its own. The store is the truth: the pickers mirror it, and only a
  // choice that differs from it (and names something) reaches `selectClass` / `selectRace` -- a
  // bound select that wrote its initial, option-less '' back would reset the build on mount.
  let classValue = $derived(store.classSlug);
  let raceValue = $derived(store.raceSlug);
  $effect(() => {
    if (classValue !== '' && classValue !== store.classSlug) store.selectClass(classValue);
  });
  $effect(() => {
    if (raceValue !== '' && raceValue !== store.raceSlug) store.selectRace(raceValue);
  });
  const classOptions = $derived(store.classes.map((row) => ({ value: row.slug, label: row.name })));
  const raceOptions = $derived(store.legalRaces.map((row) => ({ value: row.slug, label: row.name })));
</script>

<div class="planner-header" data-testid="planner-header">
  <div class="planner-header-content">
    <p class="label planner-header-eyebrow">
      <i class="planner-header-eyebrow-rule" aria-hidden="true"></i>
      {plannerHeaderCopy.eyebrow}
    </p>
    <div class="planner-header-identity">
      <img
        src={classCrestSrc(store.classSlug)}
        alt=""
        width="64"
        height="64"
        loading="eager"
        decoding="async"
        class="planner-header-crest"
        style={`--c:${color}`}
        data-testid="planner-header-crest"
      />
      <div class="planner-header-titles">
        <h1 class="planner-header-h1 heading-page" style={`color:${color}`} data-testid="planner-header-h1">
          <span class="planner-header-seg"
            >{plannerHeaderCopy.h1(store.classRow?.name ?? store.classSlug)}</span
          ><span class="planner-header-suffix">
            <span class="planner-header-sep">·&nbsp;</span>{#each suffixParts as part, index (index)}
              <span class="planner-header-seg">{part}{index < suffixParts.length - 1 ? ' ·' : ''}</span>
              <!-- eslint-disable-next-line svelte/no-useless-mustaches -- the one breakable space between segments -->
              {' '}
            {/each}
          </span>
        </h1>
        {#if originLine !== ''}
          <p class="planner-header-origin" data-testid="planner-build-origin">{originLine}</p>
        {/if}
        <p class="planner-header-description">{plannerHeaderCopy.description}</p>
      </div>
    </div>

    <div class="planner-header-controls">
      <div class="planner-header-select-field">
        <label class="quiet-label" for="planner-class-select">Class</label>

        <Select
          id="planner-class-select"
          compact
          options={classOptions}
          bind:value={classValue}
          disabled={store.readOnly}
        />
      </div>
      <div class="planner-header-select-field">
        <label class="quiet-label" for="planner-race-select">Race</label>
        <Select
          id="planner-race-select"
          compact
          options={raceOptions}
          bind:value={raceValue}
          disabled={store.readOnly}
          testid="planner-race-select"
        />
      </div>
      <span class="planner-header-faction-note">{plannerHeaderCopy.factionFollowsRace}</span>
    </div>

    <!-- The facts rail (spec 4.C, polish spec 5.1): four tiles -- Points left, Spent, Level and
         the role's score -- in a 2 x 2 grid on a phone, one wrapping row from 1100px. Each grid
         child is `min-w-0` (StatTile) and the grid is the column's full width below 640px,
         so nothing runs past the viewport; the DPS note wraps under its figure. The tiles
         stretch to the tallest one, so the rail reads as one row of equal tiles. "Sim this
         build" is its own button in the rail, above the Share panel. -->
    <div class="planner-header-facts" data-testid="planner-facts">
      <StatTile
        label="Points left"
        value={String(pointsLeft)}
        figureTestid="planner-remaining"
        figureClass={FACT_FIGURE_CLASS}
        tone="gold"
      />
      <StatTile
        label="Spent"
        value={`${store.spent}/${store.pointsAvailable}`}
        figureTestid="planner-spent"
        figureClass={FACT_FIGURE_CLASS}
      />
      <StatTile
        label="Level"
        value={String(store.level)}
        figureTestid="planner-level"
        figureClass={FACT_FIGURE_CLASS}
      />
      <PlannerDps
        unit={scoreUnit}
        {live}
        {gate}
        {pointsLeft}
        onshow={onshowdps}
        showSimLink={false}
        unfinishedNote="static"
        tile
      />
    </div>
    <p
      role="status"
      aria-live="polite"
      class="text-muted min-h-[20px] w-full text-[13px] leading-tight"
      data-testid="planner-refusal"
    >
      {statusLine}
    </p>
  </div>
</div>

<style>
  /* Fix round 5: a plain content column, no background/border/radius/overflow of its own --
     `PlannerHeaderBand.svelte` (this component's only caller) owns the band's background,
     full-bleed across both of its grid columns, which a box scoped to this one could never
     reach. */
  .planner-header {
    display: flex;
    flex-direction: column;
    width: 100%;
    min-width: 0;
  }
  .planner-header-content {
    display: flex;
    flex-direction: column;
    gap: 16px;
    padding: 22px 0 18px 0;
  }
  @media (min-width: 1024px) {
    .planner-header-content {
      padding: 28px 0 24px 0;
    }
  }
  .planner-header-eyebrow {
    display: flex;
    align-items: center;
    gap: 10px;
    color: var(--color-gold);
  }
  .planner-header-eyebrow-rule {
    width: 28px;
    height: 1px;
    background: var(--color-gold);
    display: inline-block;
  }
  .planner-header-identity {
    display: flex;
    align-items: center;
    gap: 18px;
    min-width: 0;
  }
  .planner-header-titles {
    display: flex;
    flex-direction: column;
    gap: 6px;
    min-width: 0;
  }
  .planner-header-crest {
    flex-shrink: 0;
    border-radius: 999px;
    box-shadow: 0 0 0 2px var(--c);
    background: var(--color-raised);
    object-fit: cover;
  }
  .planner-header-h1 {
    margin: 0;
  }
  /* A separator dot never starts a line: it is glued to the segment before it, and a segment
     never breaks inside itself. */
  .planner-header-seg {
    white-space: nowrap;
  }
  /* Phone: the class and "talents" on line 1, "Human · Level 9" as a 14px muted line under it,
     so line 1 loses the dot it would end on. From 640 the suffix follows on the same line, its
     leading dot glued into it (no breakable space between "talents" and the dot). */
  .planner-header-sep {
    display: none;
  }
  .planner-header-suffix {
    display: block;
    margin-top: 4px;
    /* A meta line inside the h1: the body face, not the heading face's caps. */
    font-family: var(--font-body);
    color: var(--color-muted);
    font-weight: 600;
    font-size: 14px;
    letter-spacing: 0.1em;
    text-transform: none;
    vertical-align: baseline;
  }
  @media (min-width: 640px) {
    .planner-header-sep {
      display: inline;
    }
    .planner-header-suffix {
      display: inline;
      margin-top: 0;
      margin-inline-start: 0.5em;
    }
  }
  .planner-header-origin {
    margin: 0;
    font-size: 14px;
    font-weight: 700;
    color: var(--color-gold);
  }
  .planner-header-description {
    margin: 0;
    font-size: 14px;
    color: var(--color-nav);
    max-width: 62ch;
  }
  .planner-header-controls {
    display: flex;
    align-items: center;
    gap: 14px;
    flex-wrap: wrap;
  }
  /* The server-rendered selects hold no options until the talent data lands, so a select that
     sizes to its widest option is only as wide as its padding and arrow and then jumps to its
     hydrated width. That widened pair once pushed the faction note onto its own row below 560px
     and shifted everything under the controls by 33px (0.086 of layout shift on /planner at
     360px). A fixed field width, wider than any class or race name at 15px, holds the row the
     same before and after the data arrives. */
  .planner-header-select-field {
    display: flex;
    flex-direction: column;
    gap: 4px;
    width: 152px;
  }
  .planner-header-faction-note {
    font-size: 13px;
    color: var(--color-muted);
  }
  .planner-header-facts {
    /* Polish spec 5.1: the rail is its tiles -- no panel around them. Below 640px a 2 x 2 grid
       of equal columns that is the column's full width (`minmax(0, 1fr)` lets a tile shrink
       below its content, which is what the old `auto` columns could not: the DPS note forced
       the grid past the viewport). From 640px the grid hugs its tiles, two by two; from 1100px
       one wrapping row. */
    display: grid;
    grid-template-columns: repeat(2, minmax(0, 1fr));
    align-items: stretch;
    width: 100%;
    gap: 12px;
  }
  @media (min-width: 640px) {
    .planner-header-facts {
      grid-template-columns: repeat(2, auto);
      width: fit-content;
      max-width: 100%;
      justify-content: start;
    }
  }
  @media (min-width: 1100px) {
    .planner-header-facts {
      display: flex;
      flex-wrap: wrap;
      align-items: stretch;
    }
  }
</style>
