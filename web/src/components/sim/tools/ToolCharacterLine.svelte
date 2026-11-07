<!-- web/src/components/sim/tools/ToolCharacterLine.svelte -->
<!-- Who a tool page is simming (mock design/mocks/ws_simtools.py, region `summary`): one line under
     the run bar, no box and no doll. The crest, the name in the class colour, race, spec and
     level, the item level and where the gear came from, then "Change" (the source switcher) and
     "Open in planner". The worn set is the candidates themselves on Top gear, and the doll on
     /sim; a tool page does not draw it a second time. A level note, when the talents stop short
     of the character's level, sits under the line. -->
<script lang="ts">
  import { onMount } from 'svelte';
  import { classColorVar } from '../../../lib/report/format';
  import { gearCopy } from '../../../lib/gear/copy';
  import { averageItemLevel } from '../../../lib/gear/paperdoll';
  import { slotsFromWorn } from '../../../lib/gear/paperdoll-slots';
  import { loadTalents } from '../../../lib/planner/load';
  import { indexTalents, type TalentIndex } from '../../../lib/planner/rules';
  import type { Item, TalentFile } from '../../../lib/planner/types';
  import { plannerHrefFor, type SimCharacter } from '../../../lib/sim/character';
  import { simCopy } from '../../../lib/sim/copy';
  import { handoffCopy } from '../../../lib/sim/handoff-copy';
  import { sourcePill } from '../../../lib/sim/sources';
  import { specLabel } from '../../../lib/sim/spec-label';
  import { LINK } from '../../../lib/ui/primitives';
  import { REGION } from '../../../lib/ui/regions';
  import ClassCrestRing from '../../character/ClassCrestRing.svelte';
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
  const source = $derived(sourcePill(character.source, new Date(), character.name));
  const levelNote = $derived(
    character.talent_level < character.level
      ? `${character.point_order.length} talent points. ${handoffCopy.simmedAtLevel(character.level)}`
      : '',
  );
</script>

<div class="flex flex-col gap-1 pt-[14px] pb-1" data-region={REGION.summary} data-testid="sim-character">
  <p class="m-0 flex flex-wrap items-center gap-x-3 gap-y-2 text-[15px] leading-6">
    <ClassCrestRing characterClass={character.class_slug} size={24} testid="sim-character-class-icon" />
    {#if character.name}<span
        class="font-bold"
        style:color={classColorVar(character.class_slug)}
        data-testid="paperdoll-name">{character.name}</span
      >{/if}
    <span class="text-text" data-testid="paperdoll-descriptor">{descriptor}</span>
    {#if itemLevel !== null}<span class="text-muted tabular-nums" data-testid="sim-gear-ilvl"
        >{gearCopy.paperdoll.averageItemLevel(itemLevel)}</span
      >{/if}
    <span class="text-muted" data-testid="paperdoll-source">{source}</span>
    <span class="ml-auto flex items-center gap-4">
      <a
        class="{LINK} text-[14px]"
        href={plannerHrefFor(character, talentIndex)}
        data-testid="sim-open-planner">{simCopy.openInPlanner} ›</a
      >
      <Button size="sm" onclick={onchange} testid="sim-change-source">{simCopy.changeSource}</Button>
    </span>
  </p>
  {#if levelNote !== ''}<p class="text-small text-muted m-0" data-testid="sim-level-note">{levelNote}</p>{/if}
</div>
