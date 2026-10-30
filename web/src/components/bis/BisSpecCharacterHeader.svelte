<!-- web/src/components/bis/BisSpecCharacterHeader.svelte -->
<!-- The BiS spec page's "Your character" card: the shared character header in its stacked card
     layout with the outline action (design/specs/2026-10-11-character-header.md section 4,
     replacing the old CharacterCard.svelte). Mounted `client:idle-after-load` (never
     `client:load`): this region sits well below the page's own h1, its LCP element, so a
     slightly deferred hydration costs nothing there, and an island fetched before the first
     paint would count toward the simulated LCP (src/directives/idle-after-load.ts).

     Also owns the page-level side effects the old card carried: the header race select defaults
     to the visitor's own character race once signed in (the two panels must agree, never the
     faction default), and is revealed once the session has settled either way. -->
<script lang="ts">
  import { createCharacterCardState } from '../../lib/bis/character-card-state.svelte';
  import { specPageAction, specPageAnswerSpec, type SpecPage } from '../../lib/bis/spec-header';
  import CharacterHeader from '../character/CharacterHeader.svelte';

  let {
    specKey,
    classSlug,
    specName,
    build,
    plannerHref = '/planner',
  }: SpecPage & { plannerHref?: string } = $props();

  const session = createCharacterCardState();

  /**
   * Reaches outside this island's own root to both `ClassHeader` copies (`[data-testid=
   * "bis-race-select"]`, one per faction, only one ever visible at a time), the same
   * cross-island DOM-reach pattern `HomeAccountPanel.svelte`'s own `inert` effect uses, since the
   * select lives in a plain Astro component this island cannot otherwise reach. Matched by the
   * OPTION's own text against `character.race` rather than a slug, so this needs no second
   * race-name-to-slug mapping: `ClassHeader` already renders each option's real display name as
   * its text content. A race the page's class cannot be (no matching option) leaves the select on
   * its own per-faction default rather than force a selection that does not exist here.
   */
  $effect(() => {
    const race = session.character?.race;
    if (race === undefined) return;
    document.querySelectorAll<HTMLSelectElement>('[data-testid="bis-race-select"]').forEach((select) => {
      const match = Array.from(select.options).find(
        (option) => option.textContent?.trim().toLowerCase() === race.toLowerCase(),
      );
      if (match === undefined || select.value === match.value) return;
      select.value = match.value;
      // The shared Select's face follows its select through `change` (lib/ui/select-boot.ts).
      select.dispatchEvent(new Event('change', { bubbles: true }));
    });
  });

  /**
   * Reveals the race select once the session has settled, one way or the other
   * (`ClassHeader`'s `:global(html[data-session='1']:not([data-race-select-synced]))` rule):
   * fires on `ready` AND `failed`, not only when a race was applied above, because a visitor
   * with no characters, a race this class cannot be, or a stale session hint all still need the
   * select uncovered, showing its own honest per-faction default rather than hide forever.
   */
  $effect(() => {
    if (session.status === 'loading' || session.status === 'idle') return;
    document.documentElement.setAttribute('data-race-select-synced', '1');
  });
</script>

<CharacterHeader
  source={session}
  layout="card"
  emphasis="secondary"
  action={specPageAction(plannerHref)}
  answer={specPageAnswerSpec({ specKey, classSlug, specName, build })}
/>
