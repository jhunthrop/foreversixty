<!-- web/src/components/sim/SimRacePick.svelte -->
<!-- The one question a character without a race owes the sim (a logged fight carries none; an
     addon export reads only "Skyborne"), asked once under the run line. Run stays disabled until
     it is answered: sim/request.ParseRace refuses an empty race, and a failed run is a worse
     answer than a question. -->
<script lang="ts">
  import { simCopy } from '../../lib/sim/copy';
  import type { SimCharacter } from '../../lib/sim/character';
  import type { RaceRow } from '../../lib/planner/types';
  import Select from '../ui/Select.svelte';

  let {
    character,
    races,
    onrace,
  }: { character: SimCharacter; races: readonly RaceRow[]; onrace: (slug: string) => void } = $props();

  const skyborneOnly = $derived(character.source.kind === 'addon');
  const options = $derived(
    (skyborneOnly ? races.filter((race) => race.slug.endsWith('-skyborne')) : races).map((race) => ({
      value: race.slug,
      label: race.name,
    })),
  );
</script>

<div class="flex flex-wrap items-center gap-2 pt-3">
  <label class="text-strong text-[14px]" for="sim-race-field" data-testid="sim-race-prompt"
    >{skyborneOnly ? simCopy.pickSkyborne : simCopy.pickRace}</label
  >
  <Select
    id="sim-race-field"
    size="compact"
    value=""
    placeholder={skyborneOnly ? simCopy.pickSkybornePlaceholder : simCopy.pickRacePlaceholder}
    {options}
    testid="sim-race-picker"
    onchange={(event) => onrace((event.currentTarget as HTMLSelectElement).value)}
  />
</div>
