<!-- web/src/components/sim/ColdPasteHero.svelte -->
<!-- Persona review 2026-10-06 (retail-raider), §5/§8 item 1: "Let the sim take my real
     gear/talents in one step... no sign-in detour -- the one surface where Raidbots still
     wins". The engine runs as wasm in this browser (lib/sim/engine.ts, /_sim/<version>/),
     so a cold paste can run as fast here as a pasted string does on Raidbots, with no
     queue -- this is the hero that proves it, ahead of everything else SimView.svelte
     renders for a visitor with no tracked character (signed out, or signed in with none).

     Validation is local and synchronous (cold-paste.ts's `validateColdPaste`, the same
     decode AddonPasteBox.svelte already proves on /addon), so Run disables itself the
     instant the box holds something that will not decode, with the same message that
     decoder would give anywhere else on the site -- never a second, hand-written copy of
     it. Loading the character and running it is the parent's job (SimView.svelte's
     `runPastedInput`): this component only ever reads `code` and calls `onrun` with it,
     exactly once, when it already knows that call will succeed. -->
<script lang="ts">
  import { PRIMARY_BUTTON_FIXED } from '../../lib/planner/styles';
  import { landingCopy } from '../../lib/sim/landing-copy';
  import { validateColdPaste } from '../../lib/sim/cold-paste';
  import { BUSY_CLASS } from '../../lib/ui/busy';

  let {
    busy,
    onrun,
  }: {
    busy: boolean;
    onrun: (code: string) => void;
  } = $props();

  let code = $state('');

  const check = $derived(validateColdPaste(code));
  const error = $derived(check.ok ? null : check.message);

  function run(): void {
    if (!check.ok || busy) return;
    onrun(check.code);
  }
</script>

<section
  class="border-line bg-card-top rounded-panel mx-[18px] flex flex-col gap-3 border p-4 md:mx-0"
  data-testid="sim-cold-paste"
>
  <h2 class="section-title text-[15px]">{landingCopy.pasteHeroTitle}</h2>
  <label class="sr-only" for="sim-cold-paste-input">{landingCopy.pasteHeroTitle}</label>
  <textarea
    id="sim-cold-paste-input"
    rows="2"
    class="border-line-warm rounded-control bg-raised text-text min-h-11 w-full border px-3 py-2 font-mono text-[13px]"
    placeholder={landingCopy.pasteHeroPlaceholder}
    aria-invalid={error !== null}
    bind:value={code}
    disabled={busy}
    data-testid="sim-cold-paste-input"></textarea>
  <div class="flex flex-wrap items-center gap-3">
    <button
      type="button"
      class="{PRIMARY_BUTTON_FIXED} px-5 {busy ? BUSY_CLASS : ''}"
      disabled={!check.ok || busy}
      aria-busy={busy}
      onclick={run}
      data-testid="sim-cold-paste-run"
    >
      {landingCopy.pasteHeroRun}
    </button>
    <p class="text-muted text-[12px]">
      {landingCopy.pasteHeroHintLead}
      <a href="/setup" class="text-nav underline" data-testid="sim-cold-paste-setup">
        {landingCopy.pasteHeroHintLink}
      </a>
    </p>
  </div>
  {#if error !== null}
    <p role="alert" class="text-strong text-[13px]" data-testid="sim-cold-paste-error">{error}</p>
  {/if}
  <p class="text-muted text-[12px]">{landingCopy.pasteHeroRunsLocally}</p>
</section>
