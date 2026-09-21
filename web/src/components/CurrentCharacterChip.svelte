<!-- web/src/components/CurrentCharacterChip.svelte -->
<!-- Design 4 (architect review): the visible half of the current-character pointer. Reads
     the pointer it is given (never localStorage itself -- the mounting page owns when to
     read/restore, so this stays a pure render of one prop) and shows up to three links plus
     Forget, in class colour. Fixed min-height so resolving from "no character" to "loaded"
     -- or the reverse, after Forget -- never moves content below it (Global Constraint: no
     layout shift).

     Copy addon code (spec section 1) is a third, optional link: the chip itself only ever
     holds a pointer, never the planner's talent index the addon-code string is built from,
     so the mounting page (the planner) passes the already-built string down and this stays
     a pure render. `copied` has no timer, same idiom as SharePanel.svelte's `copiedFrom`:
     it stands until the character changes underneath it, which a fresh `current` prop
     already does by remounting this component's state from scratch. -->
<script lang="ts">
  import { plannerHrefFor, simHrefFor, type CurrentCharacter } from '../lib/current-character';
  import { currentCharacterCopy } from '../lib/current-character-copy';
  import { classColorVar } from '../lib/report/format';

  let {
    current,
    onforget,
    hasOwnPasteBox = false,
    addonCode = '',
  }: {
    current: CurrentCharacter | null;
    onforget: () => void;
    hasOwnPasteBox?: boolean;
    /** The planner's current build, already encoded as an FS1 string. Empty when the
     *  planner has no build to copy -- the chip then renders no copy-addon link at all,
     *  rather than a link that would fail to copy anything. */
    addonCode?: string;
  } = $props();

  let copied = $state(false);

  async function copyAddonCode(): Promise<void> {
    try {
      await navigator.clipboard.writeText(addonCode);
      copied = true;
    } catch {
      copied = false;
    }
  }

  const HIT_TARGET = 'inline-flex min-h-11 items-center md:min-h-0';
</script>

{#if current !== null}
  <section
    class="border-line bg-raised rounded-panel flex min-h-11 flex-wrap items-center gap-3 border px-3 py-2 text-[13px]"
    data-testid="current-character-chip"
  >
    <span
      class="font-semibold"
      style={`color: ${classColorVar(current.classSlug)}`}
      data-testid="current-character-label"
    >
      {current.label}
    </span>
    <a
      class={`text-nav ${HIT_TARGET}`}
      href={plannerHrefFor(current)}
      data-testid="current-character-planner"
    >
      {currentCharacterCopy.openInPlanner}
    </a>
    <a class={`text-nav ${HIT_TARGET}`} href={simHrefFor(current)} data-testid="current-character-sim">
      {currentCharacterCopy.openInSimulator}
    </a>
    {#if addonCode !== ''}
      <button
        type="button"
        class={`text-nav ${HIT_TARGET}`}
        onclick={copyAddonCode}
        data-testid="current-character-copy-addon"
      >
        {copied ? currentCharacterCopy.copiedAddonCode : currentCharacterCopy.copyAddonCode}
      </button>
    {/if}
    <button
      type="button"
      class="text-muted ml-auto min-h-11 md:min-h-0"
      onclick={onforget}
      data-testid="current-character-forget"
    >
      {currentCharacterCopy.forget}
    </button>
  </section>
{:else if !hasOwnPasteBox}
  <p class="text-muted min-h-11 text-[13px] md:min-h-0" data-testid="current-character-chip">
    {currentCharacterCopy.noCharacterLine}
  </p>
{/if}
