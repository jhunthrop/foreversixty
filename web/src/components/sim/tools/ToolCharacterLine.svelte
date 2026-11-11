<!-- web/src/components/sim/tools/ToolCharacterLine.svelte -->
<!-- Who a tool page is simming (mock design/mocks/ws_simtools.py, region `summary`): one line under
     the run bar, no box and no doll. The crest, the name in the class colour, race, spec and
     level, the item level and where the gear came from, then "Change" (the source switcher) and
     "Open in planner". The worn set is the candidates themselves on Top gear, and the doll on
     /sim; a tool page does not draw it a second time. A level note, when the talents stop short
     of the character's level, sits on the line as "20 talent points". -->
<script lang="ts">
  import CharacterName from '../../ui/CharacterName.svelte';
  import { onMount } from 'svelte';
  import { averageItemLevel } from '../../../lib/gear/paperdoll';
  import { slotsFromWorn } from '../../../lib/gear/paperdoll-slots';
  import { loadTalents } from '../../../lib/planner/load';
  import { indexTalents, type TalentIndex } from '../../../lib/planner/rules';
  import type { Item, TalentFile } from '../../../lib/planner/types';
  import { plannerHrefFor, talentsString, type SimCharacter } from '../../../lib/sim/character';
  import { talentSplit } from '../../../lib/sim/result-band';
  import { simCopy } from '../../../lib/sim/copy';
  import { sourcePill } from '../../../lib/sim/sources';
  import { specLabel } from '../../../lib/sim/spec-label';
  import { LINK } from '../../../lib/ui/primitives';
  import { REGION } from '../../../lib/ui/regions';
  import Crest from '../../ui/Crest.svelte';
  import Button from '../../ui/Button.svelte';

  let {
    character,
    items,
    onchange,
  }: {
    character: SimCharacter;
    items: Map<number, Item>;
    onchange: () => void;
  } = $props();

  let talents = $state<TalentFile | null>(null);
  onMount(() => {
    void loadTalents(character.tree_version, character.class_slug)
      .then((file) => (talents = file))
      .catch(() => (talents = null));
  });
  const talentIndex = $derived<TalentIndex | null>(talents === null ? null : indexTalents(talents));

  const raceWords = (slug: string): string =>
    slug
      .split('-')
      .filter((word) => word !== '')
      .map((word) => word.charAt(0).toUpperCase() + word.slice(1))
      .join(' ');

  const descriptor = $derived(
    `${raceWords(character.race_slug)} ${specLabel(character.spec)} · Level ${character.level}`.trim(),
  );
  const itemLevel = $derived.by(() => {
    if (items.size === 0) return null;
    return averageItemLevel(slotsFromWorn(character.gear, { lookup: (id) => items.get(id) }));
  });
  /** Where the gear came from; a character typed in by hand says nothing (the line is the character). */
  const source = $derived(
    character.source.kind === 'manual' && !character.name
      ? ''
      : sourcePill(character.source, new Date(), character.name),
  );
  /** Talents short of the character's level (a planner build sent at 60): its split, the way every
   *  build reads on these pages ("0/20/0"), never a bare count beside another build's split. */
  const levelNote = $derived(
    character.talent_level < character.level && talentIndex !== null
      ? talentSplit(talentsString(talentIndex, character.point_order))
      : '',
  );
</script>

<div class="flex flex-col gap-1 pt-[14px] pb-1" data-region={REGION.summary} data-testid="sim-character">
  <!-- One line of facts (13 px from the item level on, so a phone holds it on one row), then the
       actions; at 390 the actions take the second row. -->
  <p class="m-0 flex flex-wrap items-center gap-x-2 gap-y-2 text-[15px] leading-6 md:gap-x-3">
    <Crest cls={character.class_slug} size={24} testid="sim-character-class-icon" />
    {#if character.name}<CharacterName
        name={character.name}
        cls={character.class_slug}
        weight="inherit"
        class="font-bold"
        testid="paperdoll-name"
      />{/if}
    <span class="text-text" data-testid="paperdoll-descriptor">{descriptor}</span>
    {#if itemLevel !== null}<span class="text-muted text-[13px] tabular-nums" data-testid="sim-gear-ilvl"
        >{simCopy.toolIlvl(itemLevel)}</span
      >{/if}
    {#if source !== ''}<span class="text-muted" data-testid="paperdoll-source">{source}</span>{/if}
    {#if levelNote !== ''}<span class="text-muted text-[13px] tabular-nums" data-testid="sim-level-note"
        >{levelNote}</span
      >{/if}
    <span class="ml-auto flex items-center gap-4 max-md:w-full max-md:justify-between">
      <a
        class="{LINK} text-[14px] [@media(pointer:coarse)]:-my-3 [@media(pointer:coarse)]:inline-block [@media(pointer:coarse)]:py-3"
        href={plannerHrefFor(character, talentIndex)}
        data-testid="sim-open-planner">{simCopy.openInPlanner} ›</a
      >
      <Button tone="ghost" size="sm" onclick={onchange} testid="sim-change-source"
        >{simCopy.changeSource}</Button
      >
    </span>
  </p>
</div>
