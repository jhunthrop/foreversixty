<!-- web/src/components/sim/RotationDisclosure.svelte -->
<!-- Final whole-branch review, C1: main carried two "what it does" controls for the same
     rotation -- this one (born as RotationCard.svelte's own Disclosure, post-run only) and
     a second, plain `<a href="/sim/specs#<spec>">` in SettingsBar.svelte that rendered the
     moment a character loaded and threw the character and the run away on click. Every
     persona repro that mattered (newcomer, tank) never ran a sim first, so every one of
     them hit the still-broken settings bar link. Rather than fix that link in place --
     a second bespoke toggle re-deriving the same drawer -- this is RotationCard's own
     Disclosure block, extracted so SettingsBar.svelte and RotationCard.svelte render the
     same never-navigates control (Disclosure.svelte, Ruling 3's one show/hide primitive)
     instead of one safe copy and one unsafe one.

     Both callers keep the drawer content and the trigger's own accessible name ("what it
     does") identical: the drawer answers "what does this rotation do" the same way whether
     or not a run has finished. `idPrefix` is the one thing that must differ between
     callers: once a run finishes, both copies are on screen at once, and the DOM `id`
     Disclosure derives (for `aria-controls`) would otherwise collide between them -- two
     panels sharing one id, and a trigger's `aria-controls` pointing at whichever happened
     to render first.

     Residual regressions fix, Finding 2: the panel and steps testids are threaded through
     `idPrefix` the same way -- `sim-rotation-drawer-panel`/`-steps` for the settings bar,
     `sim-rotation-card-drawer-panel`/`-steps` for RotationCard -- because both copies are on
     screen at once post-run, and an unscoped `getByTestId` on the old shared literal was a
     strict-mode violation waiting to happen, the exact defect this same wave fixed for
     SpecCard.svelte's per-spec testids. Every existing test that scopes through its own
     container (`sim-rotation-card` or the settings bar's own section) before reaching for
     these testids keeps working unchanged. -->
<script lang="ts">
  import { simCopy } from '../../lib/sim/copy';
  import { rotationDrawerContent } from '../../lib/sim/rotations';
  import Disclosure from './Disclosure.svelte';

  let {
    spec,
    idPrefix,
    disabled = false,
    triggerClass = '',
    triggerTestId,
  }: {
    spec: string;
    /** Unique per caller ("sim-rotation" for the settings bar, "sim-rotation-card" for the
     *  result), so the two copies' own DOM ids never collide once both are on screen. */
    idPrefix: string;
    disabled?: boolean;
    triggerClass?: string;
    triggerTestId?: string;
  } = $props();

  const drawer = $derived(rotationDrawerContent(spec));
</script>

<Disclosure
  label={simCopy.rotationLink}
  id={`${idPrefix}-drawer-${spec}`}
  {disabled}
  {triggerClass}
  panelClass="border-line-soft rounded-panel flex flex-col gap-2 border p-3"
  {triggerTestId}
  panelTestId={`${idPrefix}-drawer-panel`}
>
  <p class="text-[13px]">{drawer.intro}</p>
  {#if drawer.steps.length > 0}
    <ol class="flex list-decimal flex-col gap-1 pl-5 text-[13px]" data-testid={`${idPrefix}-drawer-steps`}>
      {#each drawer.steps as step, index (index)}
        <li>{step}</li>
      {/each}
    </ol>
  {:else}
    <p class="text-muted text-[13px]">{simCopy.rotationDrawerEmpty}</p>
  {/if}
  <a
    href={`/sim/specs#${spec}`}
    target="_blank"
    rel="noopener"
    class="text-nav text-[12px] underline decoration-dotted underline-offset-2"
    data-testid="sim-rotation-drawer-fidelity-link"
  >
    {simCopy.rotationDrawerFidelityLink}
  </a>
</Disclosure>
