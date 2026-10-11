<!-- web/src/components/bis/BisLandingHeader.svelte -->
<!-- The /bis landing's mount of the shared character header (design/specs/2026-10-11-character-
     header.md section 4): the answer is how many of the character's band's best-in-slot slots it
     wears, the action opens the character's own spec at its own band. The page is complete HTML
     without this: index.astro renders it server-side as the header's loading state (a skeleton
     for a hinted session, the shared signed-out band otherwise), and lib/bis/landing-client.ts
     mounts it over that markup for a hinted session once the page has settled. The band strip
     never moves it. `onCharacter` tells the page which character the header follows. -->
<script lang="ts">
  import type { MeCharacter } from '../../lib/account/api';
  import { landingAction, landingAnswerSpec } from '../../lib/bis/landing-answer';
  import type { LandingConfig } from '../../lib/bis/landing-config';
  import CharacterHeader from '../character/CharacterHeader.svelte';

  let {
    config,
    onCharacter,
  }: { config: LandingConfig; onCharacter?: (character: MeCharacter | null) => void } = $props();
</script>

<CharacterHeader
  action={({ character }) => landingAction(character, config)}
  answer={landingAnswerSpec(config)}
  onCharacterChange={onCharacter}
/>
