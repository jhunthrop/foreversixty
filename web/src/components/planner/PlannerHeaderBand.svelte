<!-- web/src/components/planner/PlannerHeaderBand.svelte -->
<!-- The page's own header BAND -- full-bleed, edge to edge, the same shape the BiS page's
     `ArtPanel` + its own `.header-row` already ship. It breaks out of the site column with the
     shared `.full-bleed` utility and puts its content back in the column with `.page-column`.

     It owns the band's background (flat `--color-raised`, the night-to-`--bg` gradient over it --
     spec §13 keeps the dominant-tree art itself off, `HEADER_ART_ENABLED` below, until it can
     paint from first byte without regressing Lighthouse's TBT again), the inner content wrapper
     (max-width 1344px, centred, the page's own 48px gutters from 1024px, 18px below it), and the
     `1fr 440px` grid (24px gap, bottom-aligned) that puts the shared character header inside the
     band as its own right column -- its stacked card layout with the outline action and no
     answer slot, because the planner's job is the tree, not a figure about the character
     (design/specs/2026-10-11-character-header.md section 4). The same shape
     `ClassHeader.astro`'s own `.header-row` uses on the BiS page. Signed out, the card column
     holds the header's one shared signed-out band, so the grid is the same for every visitor. -->
<script lang="ts">
  import { dataUrl } from '../../lib/planner/load';
  import type { LiveDps } from '../../lib/planner/live-dps.svelte';
  import type { LiveGate } from '../../lib/planner/live-gate';
  import type { PlannerStore } from '../../lib/planner/store.svelte';
  import type { MeCharacter } from '../../lib/account/api';
  import { characterPointsSpent, sendActionFor, type SendState } from '../../lib/planner/send-state';
  import { ensureItems } from '../../lib/planner/items-ready';
  import { plannerCharacterCardCopy } from '../../lib/planner/copy';
  import { classDisplayName } from '../../lib/sim/spec-label';
  import { buildHasGear } from '../../lib/addon-export';
  import CharacterHeader from '../character/CharacterHeader.svelte';
  import ExportNotRead from '../character/ExportNotRead.svelte';
  import PlannerHeader from './PlannerHeader.svelte';

  let {
    store,
    live,
    gate,
    onshowdps,
    oncharacter,
    origin = null,
    simHref,
  }: {
    store: PlannerStore;
    live: LiveDps;
    gate: LiveGate;
    onshowdps: () => void;
    /** Fires with the card's current character (or null), for the band panel's level cut. */
    oncharacter?: (character: MeCharacter | null) => void;
    /** Set when the build on screen is another player's: whose it is ('' for a link that names
     *  nobody). The send asks before it replaces anything, and the header says whose it is. */
    origin?: { owner: string } | null;
    /** The planner's "Sim this build" link, drawn in the card under the send button. */
    simHref?: string;
  } = $props();

  /** The card's current character: the send is aimed at it. Null signed out or with none synced. */
  let character = $state<MeCharacter | null>(null);
  let sendState = $state<SendState>({ kind: 'idle' });
  /** The build + character the last successful send was for, and the id that build saved under. */
  let sent = $state<{ key: string; buildId: string } | null>(null);

  /** This build aimed at this character, or null while there is nothing sendable. */
  function sendKey(): string | null {
    if (character === null) return null;
    try {
      return `${character.key}|${JSON.stringify(store.toDraft())}`;
    } catch {
      return null;
    }
  }
  const currentKey = $derived(sendKey());

  // "Sent" holds only until the build (or the character) changes.
  $effect(() => {
    if (sendState.kind === 'sent' && sent?.key !== currentKey) sendState = { kind: 'idle' };
  });

  // Asking "Send X's build to Y?" holds only while the build and the character stay as they were.
  let confirmKey: string | null = null;
  $effect(() => {
    if (sendState.kind === 'confirming' && (origin === null || confirmKey !== currentKey)) {
      sendState = { kind: 'idle' };
    }
  });

  const action = $derived(
    sendActionFor(character, sendState, store.spent, characterPointsSpent(character), {
      classSlug: store.classSlug,
      className: classDisplayName(store.classSlug),
      owner: origin === null ? undefined : origin.owner,
      simHref,
    }),
  );
  /** The card's character has an addon export that carried no gear and no talents (W.3): the
   *  card says so instead of leaving the empty tree to read as the character's build. */
  const unread = $derived(
    character?.build !== undefined && character.build.source === 'addon' && !buildHasGear(character.build),
  );

  // The card names the worn pieces the planner has no item for, so a character wearing anything
  // asks for the class's catalogue; a visitor with an empty gear map never fetches it.
  $effect(() => {
    if (character !== null && Object.keys(store.gear).length > 0) void ensureItems(store);
  });

  async function send(): Promise<void> {
    // The draft is filtered by the catalogue: have it before the build is read, so the key the
    // result is checked against is the one that was sent.
    await ensureItems(store);
    const key = currentKey;
    if (character === null || key === null || sendState.kind === 'sending' || store.spent === 0) return;
    // Another class's build is never sent to a character (the button is off, this is the guard).
    if (character.class !== undefined && character.class.toLowerCase() !== store.classSlug) return;
    // Another player's build asks once; the yes comes back through here.
    if (origin !== null && sendState.kind !== 'confirming') {
      confirmKey = key;
      sendState = { kind: 'confirming' };
      return;
    }
    const target = character.key;
    sendState = { kind: 'sending' };
    const known = sent?.key === key ? sent.buildId : null;
    // The account API, the save and the queue load with the first press, so the planner
    // island stays under its boot budget.
    const { sendBuildToAddon } = await import('../../lib/planner/send-build');
    const outcome = await sendBuildToAddon(store.toDraft(), target, known);
    // The grid stays editable while this runs: a result for a build that has since changed
    // is not a result for the build on screen.
    if (sendKey() !== key) {
      sendState = { kind: 'idle' };
      return;
    }
    if (outcome.kind === 'sent') {
      sent = { key, buildId: outcome.buildId };
      sendState = { kind: 'sent' };
    } else {
      sendState = outcome;
    }
  }

  /** The dominant tree's own index -- the same "most points, ties to the first tree"
   *  inference the h1 suffix (`PlannerHeader.svelte`) already uses, read here for the
   *  background art instead of the spec label. Moved here from `PlannerHeader.svelte`
   *  (fix round 5): the art has to cover the whole band, both grid columns, not just the
   *  header's own left one, so it can only ever live at this level. */
  const dominantTreeIndex = $derived.by(() => {
    const split = store.split;
    let best = 0;
    for (let i = 1; i < split.length; i += 1) {
      if (split[i] > split[best]) best = i;
    }
    return best;
  });

  /** Spec §13 (2026-10-02): the band ships flat. Art painted after talent data resolves is
   *  always the page's largest paint and lands at ~3.3 s under the audit's throttling,
   *  which put the merged build at 0.89 against the 0.90 performance floor. The art comes
   *  back when the planner can paint it from the first byte (per-class static routes or a
   *  server-rendered class); flipping this constant is the whole switch. */
  const HEADER_ART_ENABLED = false;
  const headerArtSrc = $derived.by(() => {
    if (!HEADER_ART_ENABLED) return null;
    const tree = store.talentIndex?.trees[dominantTreeIndex];
    return tree ? dataUrl(store.treeVersion, `trees/${tree.background}.webp`) : null;
  });
</script>

<div class="planner-header-band full-bleed" data-testid="planner-header-band">
  <!-- A CSS background-image, not an <img>: a nicer fit for a decorative, `aria-hidden`
       backdrop over the whole band -- see `PlannerHeader.svelte`'s own, now-historical
       comment on why this is not a way to dodge LCP candidacy (it isn't; Chrome counts a
       CSS background the moment it is visibly painted) and why the swap is instant, no
       opacity transition, once the art is back in use. -->
  <div
    class="planner-header-band-art"
    class:planner-header-band-art-visible={headerArtSrc !== null}
    style={headerArtSrc !== null ? `background-image:url(${headerArtSrc})` : undefined}
    aria-hidden="true"
  ></div>
  <div class="planner-header-band-gradient" aria-hidden="true"></div>
  <div class="planner-header-band-inner page-column">
    <div class="planner-header-band-grid">
      <PlannerHeader {store} {live} {gate} {onshowdps} {origin} />
      <div class="flex flex-col gap-3">
        <CharacterHeader
          layout="card"
          emphasis="secondary"
          {action}
          onCharacterChange={(next) => {
            character = next;
            oncharacter?.(next);
          }}
          onPick={() => void send()}
          onSecondary={() => (sendState = { kind: 'idle' })}
        />
        {#if character !== null && store.leftOutSlots.length > 0}
          <p class="text-muted text-[13px]" data-testid="planner-left-out">
            {plannerCharacterCardCopy.leftOut(store.leftOutSlots)}
          </p>
        {/if}
        {#if unread}
          <ExportNotRead
            capturedAt={character?.build?.captured_at ?? null}
            testid="planner-export-not-read"
          />
        {/if}
      </div>
    </div>
  </div>
</div>

<style>
  .planner-header-band {
    position: relative;
    overflow: hidden;
    /* Flat `--color-raised` floor (spec §13): the dominant-tree art stays off until it can
       paint from first byte without becoming the page's own LCP/TBT regression again --
       see this file's own `HEADER_ART_ENABLED` comment above. */
    background: var(--color-raised);
  }
  .planner-header-band-art {
    position: absolute;
    inset: 0;
    background-size: cover;
    background-position: center;
    /* Flat until the data resolves and (if the art is ever re-enabled) names the dominant
       tree's own background: opacity 0 so the element paints nothing, the flat
       `--color-raised` band behind it the only "floor" at first paint. No `transition` --
       measured, an animated fade only moved the LCP timestamp later for no candidacy
       benefit (see git history on `PlannerHeader.svelte`'s own prior copy of this note). */
    opacity: 0;
  }
  .planner-header-band-art-visible {
    opacity: 1;
  }
  .planner-header-band-gradient {
    position: absolute;
    inset: 0;
    /* Night-to-bg gradient (unchanged from the header's own prior recipe): reads as the
       page's header band, not a card -- no border, no radius, no inset box. */
    background: linear-gradient(180deg, var(--color-night) 0%, var(--color-bg) 100%);
    opacity: 0.5;
  }
  .planner-header-band-inner {
    position: relative;
    /* The phone inset below 768px; from there the site column's own gutter, so the header's
       left edge is the body's. A scoped `padding-inline` outranks `.page-column`'s gutter, so
       it has to name the gutter itself. */
    padding-inline: var(--page-inset);
  }
  @media (min-width: 768px) {
    .planner-header-band-inner {
      padding-inline: var(--page-gutter);
    }
  }
  .planner-header-band-grid {
    display: grid;
    grid-template-columns: 1fr;
    gap: 16px;
  }
  @media (min-width: 1024px) {
    .planner-header-band-grid {
      grid-template-columns: 1fr 440px;
      gap: 24px;
      align-items: end;
    }
  }
</style>
