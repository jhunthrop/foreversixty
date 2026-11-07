<!-- web/src/components/planner/InGame.svelte -->
<!-- `In the game` (spec 2026-11-08 sections 4.7 and 5.3, decision P6, ruling Q6): the page's one
     primary. Signed in with the header character on this class, `Send this build to Felsnap` is
     the gold primary with the send note under it and `Copy for the game` beneath; another class's
     build disables the send with its sentence and makes Copy the primary; signed out (or no
     character) Copy for the game is the primary, then the hint and the way to the addon. The send
     state machine is the shipped one (`sendActionFor`, `send-build.ts`), moved here from the
     retired character card. -->
<script lang="ts">
  import type { MeCharacter } from '../../lib/account/api';
  import { addonCopy } from '../../lib/addon/copy';
  import { ctaCopy } from '../../lib/addon/cta-copy';
  import { plannerCharacterCardCopy as cardCopy, plannerCopy } from '../../lib/planner/copy';
  import { plannerAddonCode } from '../../lib/planner/current-character-planner';
  import { ensureItems } from '../../lib/planner/items-ready';
  import { characterPointsSpent, sendActionFor, type SendState } from '../../lib/planner/send-state';
  import type { PlannerStore } from '../../lib/planner/store.svelte';
  import { classDisplayName } from '../../lib/sim/spec-label';
  import { LINK } from '../../lib/ui/primitives';
  import CopyForGame from '../CopyForGame.svelte';
  import Button from '../ui/Button.svelte';
  import Section from '../ui/Section.svelte';

  let {
    store,
    character,
    origin = null,
  }: {
    store: PlannerStore;
    /** The header selector's character (null signed out or with none): the send is aimed at it. */
    character: MeCharacter | null;
    /** Set when the build on screen is another player's ('' for a link that names nobody). */
    origin?: { owner: string } | null;
  } = $props();

  let sendState = $state<SendState>({ kind: 'idle' });
  /** The build + character the last successful send was for, and the id that build saved under. */
  let sent = $state<{ key: string; buildId: string } | null>(null);
  let confirmKey: string | null = null;

  /** This build aimed at this character, or null while there is nothing sendable. */
  function sendKey(): string | null {
    if (character === null) return null;
    try {
      return `${character.key}|${JSON.stringify(store.toDraft())}`;
    } catch {
      return null;
    }
  }
  const currentKey = $derived(sendKey());

  // A build changed after it was sent is not sent; a confirm left open over another build closes.
  $effect(() => {
    if (sendState.kind === 'sent' && sent?.key !== currentKey) sendState = { kind: 'idle' };
  });
  $effect(() => {
    if (sendState.kind === 'confirming' && (origin === null || confirmKey !== currentKey)) {
      sendState = { kind: 'idle' };
    }
  });
  $effect(() => {
    if (character !== null && Object.keys(store.gear).length > 0) void ensureItems(store);
  });

  const action = $derived(
    sendActionFor(character, sendState, store.spent, characterPointsSpent(character), {
      classSlug: store.classSlug,
      className: classDisplayName(store.classSlug),
      owner: origin === null ? undefined : origin.owner,
    }),
  );
  /** The build on screen is another class's than the character: Copy becomes the primary. */
  const otherClass = $derived(
    character !== null && character.class !== undefined && character.class.toLowerCase() !== store.classSlug,
  );
  const addonCode = $derived(plannerAddonCode(store));

  async function send(): Promise<void> {
    await ensureItems(store);
    const key = currentKey;
    if (character === null || key === null || sendState.kind === 'sending' || store.spent === 0) return;
    if (otherClass) return;
    if (origin !== null && sendState.kind !== 'confirming') {
      confirmKey = key;
      sendState = { kind: 'confirming' };
      return;
    }
    const target = character.key;
    sendState = { kind: 'sending' };
    const known = sent?.key === key ? sent.buildId : null;
    const { sendBuildToAddon } = await import('../../lib/planner/send-build');
    const outcome = await sendBuildToAddon(store.toDraft(), target, known);
    if (sendKey() !== key) {
      sendState = { kind: 'idle' };
      return;
    }
    if (outcome.kind === 'sent') {
      sent = { key, buildId: outcome.buildId };
      sendState = { kind: 'sent' };
    } else {
      sendState = outcome;
    }
  }

  const note = $derived('note' in action ? action.note : undefined);
  const noteTone = $derived(
    note === undefined
      ? ''
      : note.tone === 'error'
        ? 'text-wipe'
        : note.tone === 'ok'
          ? 'text-kill'
          : 'text-muted',
  );
</script>

<Section title={plannerCopy.inGameTitle} testid="planner-in-game">
  <div class="pl-sec">
    {#if character !== null}
      <div class="pl-acts">
        <Button
          tone={otherClass ? 'secondary' : 'primary'}
          disabled={'disabled' in action && action.disabled === true}
          onclick={() => void send()}
          testid="planner-send">{action.label}</Button
        >
        {#if 'secondary' in action && action.secondary?.testid === 'planner-send-cancel'}
          <Button onclick={() => (sendState = { kind: 'idle' })} testid="planner-send-cancel"
            >{action.secondary.label}</Button
          >
        {/if}
      </div>
      {#if note !== undefined}
        <p
          class="pl-note {noteTone}"
          role={note.tone === 'error' ? 'alert' : 'status'}
          data-testid="planner-send-note"
        >
          {note.text}
          {#if note.link}<a class={LINK} href={note.link.href}>{note.link.label} ›</a>{/if}
        </p>
        {#if note.detail}<p class="text-small text-muted">{note.detail}</p>{/if}
      {/if}
      {#if store.leftOutSlots.length > 0}
        <p class="text-small text-muted" data-testid="planner-left-out">
          {cardCopy.leftOut(store.leftOutSlots)}
        </p>
      {/if}
    {/if}
    <CopyForGame code={store.spent === 0 ? '' : addonCode} primary={character === null || otherClass} />
    <p class="text-small text-muted">{addonCopy.addonCodeHint}</p>
    {#if character === null}
      <p>
        <a class={LINK} href={cardCopy.getTheAddonHref} data-testid="planner-get-addon"
          >{ctaCopy.getAddonAndCompanion} ›</a
        >
      </p>
    {/if}
  </div>
</Section>
