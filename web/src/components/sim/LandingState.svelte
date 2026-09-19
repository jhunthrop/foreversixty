<!-- web/src/components/sim/LandingState.svelte -->
<!-- What a signed-in member sees when they open /sim: their characters, one button each,
     and no form at all until they ask for one.
     The footnote is not boilerplate. The contract's sim-input has no Armory source yet, so
     the gear behind each of these rows is the member's last addon export or last logged
     fight, and a member who assumed the Armory was being read would be wrong about how
     fresh their gear is. The design's "last-logout gear" is precisely what an addon export
     holds, so this is the feature through the source that exists. -->
<script lang="ts">
  import type { MeCharacter } from '../../lib/account/api';
  import type { CharacterPath } from '../../lib/characters';
  import { rulesetLabel } from '../../lib/characters';
  import { classColorVar } from '../../lib/report/format';
  import { simCopy } from '../../lib/sim/copy';

  let {
    characters,
    busyKey,
    onpick,
    onother,
  }: {
    characters: MeCharacter[];
    busyKey: string | null;
    onpick: (path: CharacterPath) => void;
    onother: () => void;
  } = $props();

  function pathOf(character: MeCharacter): CharacterPath {
    const [region, ruleset, slug] = character.key.split('/');
    return { region, ruleset, slug } as CharacterPath;
  }
</script>

<section class="mx-[18px] flex flex-col gap-3 md:mx-0" data-testid="sim-landing">
  <h2 class="section-title text-[15px]">{simCopy.yourCharacters}</h2>

  <ul class="border-line bg-raised rounded-panel flex flex-col border">
    {#each characters as character (character.key)}
      {@const colour = classColorVar(character.class)}
      <li
        class="border-line-soft flex min-h-11 flex-wrap items-center gap-3 border-b px-3 py-2 last:border-b-0"
        data-testid={`sim-character-${character.key}`}
      >
        <span class="rounded-control h-7 w-7 shrink-0" style={`background: ${colour}`} aria-hidden="true"
        ></span>
        <span class="text-[15px] font-semibold" style={`color: ${colour}`}>{character.name}</span>
        <span class="text-muted text-[13px]">
          {rulesetLabel(character.ruleset)} · {character.region.toUpperCase()}
        </span>
        <button
          type="button"
          class="border-line-warm-strong rounded-control text-strong label ml-auto min-h-11 border px-4 disabled:opacity-50 md:min-h-9"
          disabled={busyKey !== null}
          onclick={() => onpick(pathOf(character))}
          data-testid={`sim-pick-${character.key}`}
        >
          {busyKey === character.key ? simCopy.loading : simCopy.simIt}
        </button>
      </li>
    {/each}
  </ul>

  <p class="text-muted text-[12px]" data-testid="sim-landing-note">{simCopy.landingSourceNote}</p>

  <button
    type="button"
    class="text-nav label min-h-11 self-start underline md:min-h-9"
    onclick={onother}
    data-testid="sim-other-character"
  >
    {simCopy.otherCharacter}
  </button>
</section>
