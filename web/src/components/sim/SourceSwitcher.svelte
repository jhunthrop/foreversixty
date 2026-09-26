<!-- web/src/components/sim/SourceSwitcher.svelte -->
<!-- The four ways in. Order is deliberate: the addon export is first because it needs no
     account and produces the most accurate character, and "your characters" is last because
     it is the one that needs a sign-in. The site now has a real Battle.net-backed source
     (spec 2026-09-22 §3.3), so signing in gets a member's gear and talents from Blizzard
     directly rather than only from a prior addon export.
     A signed-in member's own character list lives in LandingState.svelte, not here (Task
     18): this switcher only reaches a signed-in member when they explicitly reopen it (from
     the strip's "Change source", or the landing state's "Sim something else"), so the card's
     signed-in body is a way back to that list rather than a second copy of it.

     2026-09-26 layout pass, review round 1: `heroSignIn` promotes the sign-in card to a
     full-width hero above the three paste cards, for a signed-out visitor on /sim only
     (Finding 2 -- the audience already uses Battle.net, so "sign in and go" should not sit
     equal-weight beside three cards that need something to paste). The bulk tool pages
     (ToolsView.svelte) never pass it, so their switcher keeps today's four-card grid. -->
<script lang="ts">
  import { currentCharacterCopy } from '../../lib/current-character-copy';
  import { landingCopy } from '../../lib/sim/landing-copy';
  import { PRIMARY_BUTTON_FIXED } from '../../lib/planner/styles';
  import { rowLink } from '../../lib/report/format';
  import { parseBuildInput } from '../../lib/sim/build-input';
  import { simCopy } from '../../lib/sim/copy';
  import { BUSY_CLASS } from '../../lib/ui/busy';
  import ExampleResultCard from './ExampleResultCard.svelte';

  let {
    busy,
    message,
    signedIn,
    onaddon,
    onbuild,
    onfight,
    onsignin,
    onback = () => {},
    hasCharacters = true,
    heroSignIn = false,
  }: {
    busy: boolean;
    message: string | null;
    signedIn: boolean;
    onaddon: (code: string) => void;
    onbuild: (id: string) => void;
    onfight: (ref: string) => void;
    onsignin: () => void;
    /** Signed-in only: returns to the landing state's character list. */
    onback?: () => void;
    /** False for a signed-in account with no characters: there is nothing to go back to,
     *  so the account card carries only its note. */
    hasCharacters?: boolean;
    /** Finding 2, 2026-09-26 layout pass: renders the sign-in card as a full-width hero
     *  above the three paste cards instead of a fourth, equal-weight one. Has no effect
     *  while `signedIn` is true -- that state is always the "back to your characters" grid. */
    heroSignIn?: boolean;
  } = $props();

  let addonCode = $state('');
  let buildRef = $state('');
  let fightRef = $state('');

  const card = 'border-line bg-card-top rounded-panel flex flex-col gap-3 border p-4';
  const field =
    'border-line-warm rounded-control bg-raised text-text min-h-11 w-full border px-3 py-2 text-[14px]';
  const action = 'border-line-warm-strong rounded-control text-strong label min-h-11 self-start border px-4';
  // The three Load buttons all start a request, so each takes the shared busy look and
  // aria-busy while `busy` holds (design 2026-09-22 spec section 3.2). Their label is
  // "Load" whether or not a load is running -- only the look changes.
  const actionBusy = $derived(busy ? `${action} ${BUSY_CLASS}` : action);

  const showHero = $derived(heroSignIn && !signedIn);

  /** A saved link or id loads the saved build; an unsaved planner link loads its code. */
  function loadBuildInput(): void {
    const input = parseBuildInput(buildRef);
    if (input === null) return;
    if (input.kind === 'code') onaddon(input.code);
    else onbuild(input.id);
  }

  /** A report link carries the fight in ?fight=; a pasted ref already has it after a colon. */
  export function fightRefOf(value: string): string {
    const trimmed = value.trim();
    if (trimmed.includes(':')) return trimmed;
    const [path, query = ''] = trimmed.split('?');
    const id = path.replace(/\/+$/, '').split('/').at(-1) ?? '';
    const fight = new URLSearchParams(query).get('fight') ?? '1';
    return `${id}:${fight}`;
  }
</script>

{#snippet addonCard()}
  <div class={card}>
    <h2 class="section-title text-[15px]">{simCopy.sourceAddonTitle}</h2>
    <p class="text-muted text-[13px]">{simCopy.sourceAddonBody}</p>
    <label class="sr-only" for="sim-addon">{simCopy.sourceAddonTitle}</label>
    <textarea
      id="sim-addon"
      rows="2"
      class={field}
      placeholder="FS1:…"
      bind:value={addonCode}
      disabled={busy}
      data-testid="sim-addon-input"></textarea>
    <button
      type="button"
      class={actionBusy}
      disabled={busy}
      aria-busy={busy}
      onclick={() => onaddon(addonCode)}
      data-testid="sim-addon-load"
    >
      Load
    </button>
    <!-- The addon pointer belongs to this card, not to the row's end where it read as an
         orphaned line under three cards. -->
    <p class="text-muted text-[12px]">
      <a href="/setup" class="{rowLink} text-nav" data-testid="sim-get-addon"
        >{currentCharacterCopy.getTheAddon}</a
      >
    </p>
  </div>
{/snippet}

{#snippet buildCard()}
  <div class={card}>
    <h2 class="section-title text-[15px]">{simCopy.sourceBuildTitle}</h2>
    <p class="text-muted text-[13px]">{simCopy.sourceBuildBody}</p>
    <label class="sr-only" for="sim-build">{simCopy.sourceBuildTitle}</label>
    <input
      id="sim-build"
      class={field}
      placeholder="foreversixty.gg/b/…"
      bind:value={buildRef}
      disabled={busy}
      data-testid="sim-build-input"
    />
    <button
      type="button"
      class={actionBusy}
      disabled={busy}
      aria-busy={busy}
      onclick={loadBuildInput}
      data-testid="sim-build-load">Load</button
    >
  </div>
{/snippet}

{#snippet fightCard()}
  <div class={card}>
    <h2 class="section-title text-[15px]">{simCopy.sourceFightTitle}</h2>
    <p class="text-muted text-[13px]">{simCopy.sourceFightBody}</p>
    <label class="sr-only" for="sim-fight">{simCopy.sourceFightTitle}</label>
    <input
      id="sim-fight"
      class={field}
      placeholder="foreversixty.gg/reports/…?fight=2"
      bind:value={fightRef}
      disabled={busy}
      data-testid="sim-fight-input"
    />
    <button
      type="button"
      class={actionBusy}
      disabled={busy}
      aria-busy={busy}
      onclick={() => onfight(fightRefOf(fightRef))}
      data-testid="sim-fight-load">Load</button
    >
  </div>
{/snippet}

<section class="mx-[18px] flex flex-col gap-3 md:mx-0" data-testid="sim-sources">
  {#if showHero}
    <!-- Finding 5: the intro line and a static example of a finished sim, ahead of every
         load method -- neither paragraph is orphaned: the intro belongs to the page's own
         lead, and the example card is its own bordered panel, never a bare floating line. -->
    <p class="text-muted text-[14px]" data-testid="sim-intro-line">{landingCopy.introLine}</p>
    <ExampleResultCard />

    <!-- Finding 2: the hero. Full width, first, with the one PRIMARY_BUTTON on this view
         and the email-link alternative as text -- the three paste cards below are
         secondary. Finding 3: the DPS-only restriction lives here too, as the caption
         under this heading, its only copy on the page for a signed-out visitor. -->
    <div class={card} data-testid="sim-account-card">
      <h2 class="section-title text-[15px]">{simCopy.sourceAccountTitle}</h2>
      <p class="text-muted text-[13px]" data-testid="sim-scope-note">{landingCopy.scopeCaveat}</p>
      <p class="text-muted text-[13px]">{simCopy.signInToFindCharacters}</p>
      <button
        type="button"
        class="{PRIMARY_BUTTON_FIXED} self-start px-5"
        onclick={onsignin}
        data-testid="sim-signin"
      >
        {landingCopy.signInWithBattlenet}
      </button>
      <a class="text-nav text-[13px] underline" href="/login">{landingCopy.emailLinkInstead}</a>
    </div>

    <div class="grid grid-cols-1 gap-3 md:grid-cols-3">
      {@render addonCard()}
      {@render buildCard()}
      {@render fightCard()}
    </div>
  {:else}
    <div class="grid grid-cols-1 gap-3 md:grid-cols-2">
      {@render addonCard()}
      {@render buildCard()}
      {@render fightCard()}

      <div class={card} data-testid="sim-account-card">
        <h2 class="section-title text-[15px]">{simCopy.sourceAccountTitle}</h2>
        {#if signedIn && hasCharacters}
          <button type="button" class={action} onclick={onback} data-testid="sim-back-to-characters">
            {simCopy.backToCharacters}
          </button>
        {:else if !signedIn}
          <p class="text-muted text-[13px]">{simCopy.signInToFindCharacters}</p>
          <button type="button" class={action} onclick={onsignin} data-testid="sim-signin">
            {landingCopy.signInWithBattlenet}
          </button>
        {/if}
      </div>
    </div>

    <!-- task-2-brief.md: this switcher is what a signed-out visitor reads first on the bulk
         tool pages -- the same scope sentence their own Astro shell already carries above
         the fold, repeated here since it sits below the shell's own copy of it once the
         island mounts. /sim's own hero branch above carries the one copy of this sentence
         it needs instead (Finding 3). -->
    <p class="text-muted text-[12px]" data-testid="sim-sources-scope-note">{simCopy.scopeNote}</p>
  {/if}

  {#if message}
    <p role="alert" class="text-strong text-[13px]" data-testid="sim-source-message">{message}</p>
  {/if}
</section>
