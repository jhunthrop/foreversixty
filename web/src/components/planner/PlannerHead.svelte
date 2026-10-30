<!-- web/src/components/planner/PlannerHead.svelte -->
<!-- The planner's page head (site alignment spec 2026-10-29 §5.5, D1): `PageHead` on the page
     ground, no band. The crest is the lead, "{Class} talents" the h1, the character card (the
     shared `CharacterHeader`, its stacked card layout with the outline action and no answer slot,
     design/specs/2026-10-11-character-header.md section 4) the aside, and the Class and Race
     selects with the facts rail the `below` row. Signed out, the card column holds the header's
     one shared signed-out band, so the grid is the same for every visitor. -->
<script lang="ts">
  import { classCrestSrc } from '../../lib/class-crest';
  import { classColorVar } from '../../lib/report/format';
  import { specOf } from '../../lib/sim/character';
  import { specLabel } from '../../lib/sim/spec-label';
  import type { LiveDps } from '../../lib/planner/live-dps.svelte';
  import type { LiveGate } from '../../lib/planner/live-gate';
  import type { PlannerStore } from '../../lib/planner/store.svelte';
  import type { MeCharacter } from '../../lib/account/api';
  import { characterPointsSpent, sendActionFor, type SendState } from '../../lib/planner/send-state';
  import { ensureItems } from '../../lib/planner/items-ready';
  import { plannerCharacterCardCopy, plannerHeaderCopy } from '../../lib/planner/copy';
  import { classDisplayName } from '../../lib/sim/spec-label';
  import { buildHasGear } from '../../lib/addon-export';
  import CharacterHeader from '../character/CharacterHeader.svelte';
  import ExportNotRead from '../character/ExportNotRead.svelte';
  import PageHead from '../ui/PageHead.svelte';
  import PlannerHeadControls from './PlannerHeadControls.svelte';

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

  const color = $derived(classColorVar(store.classRow?.name));

  /** True once one tree holds strictly more points than each of the other two --
   *  `specKeyFor`'s own tie-break, reused unchanged. A tie, including 0/0/0, is false: the
   *  meta line's spec clause only ever appears once a build genuinely leans one way. */
  const hasLeadingTree = $derived.by(() => {
    const split = store.split;
    if (split.length === 0) return false;
    const max = Math.max(...split);
    return max > 0 && split.filter((points) => points === max).length === 1;
  });
  const leadingSpec = $derived(
    hasLeadingTree && store.talentIndex !== null
      ? specLabel(specOf(store.talentIndex, [...store.order]))
      : '',
  );

  /** "Ashenveil's build · Holy Paladin": whose build this is, when it is not the player's own. */
  const originLine = $derived(
    origin === null
      ? ''
      : plannerCharacterCardCopy.originLine(
          origin.owner,
          leadingSpec || (store.classRow?.name ?? store.classSlug),
        ),
  );

  /** The meta line: race, level, and the spec once a tree leads. */
  const metaLine = $derived(
    [store.raceRow?.name ?? '', `Level ${store.level}`, leadingSpec]
      .filter((part) => part !== '')
      .join(' · '),
  );
</script>

{#snippet crest()}
  <img
    src={classCrestSrc(store.classSlug)}
    alt=""
    width="64"
    height="64"
    loading="eager"
    decoding="async"
    class="bg-raised h-16 w-16 shrink-0 rounded-full object-cover"
    style={`box-shadow: 0 0 0 2px ${color}`}
    data-testid="planner-header-crest"
  />
{/snippet}

{#snippet meta()}
  {metaLine}
  {#if originLine !== ''}
    <span class="text-gold block font-bold" data-testid="planner-build-origin">{originLine}</span>
  {/if}
{/snippet}

{#snippet characterCard()}
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
  </div>
{/snippet}

{#snippet controls()}
  <PlannerHeadControls {store} {live} {gate} {onshowdps} />
{/snippet}

<!-- The card is a column of its own beside the head, top-aligned: as the PageHead aside it is
     bottom-aligned, which leaves the title floating under a card-tall gap. -->
<div class="flex flex-col gap-[22px] md:gap-8">
  <div
    class="flex flex-col gap-4 lg:grid lg:grid-cols-[minmax(0,1fr)_440px] lg:items-start lg:gap-6"
    data-testid="planner-header-band"
  >
    <PageHead
      eyebrow={plannerHeaderCopy.eyebrow}
      title={plannerHeaderCopy.h1(store.classRow?.name ?? store.classSlug)}
      titleColor={color}
      job={plannerHeaderCopy.description}
      lead={crest}
      {meta}
      below={controls}
      testid="planner-head"
    />
    {@render characterCard()}
  </div>
  <!-- The unread-export notice runs the full width under the head, not as a third card in the
       character column: beside the title block it left a ~200 px hole under the stat tiles
       (leftovers lane A, planner row). -->
  {#if unread}
    <ExportNotRead capturedAt={character?.build?.captured_at ?? null} testid="planner-export-not-read" />
  {/if}
</div>
