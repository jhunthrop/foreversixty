<!-- web/src/components/setup/UnknownRealmQuestion.svelte -->
<!-- The unknown-realm question (lib/account/unknown-realm.ts): one block per Battle.net character
     the import could not type a realm for, "Which realm type is <name> on?" with the site's four
     rulesets as one-tap buttons. A tap saves the answer (POST /v1/me/realm-answers), the pressed
     button busy meanwhile; a refusal shows the API's sentence in the wipe colour. `onAnswered`
     re-reads the account, which drops the answered block. -->
<script lang="ts">
  import { answerRealm } from '../../lib/account/api';
  import type { MeCharacter } from '../../lib/account/me-types';
  import { REALM_CHOICES, unknownRealmCopy } from '../../lib/account/unknown-realm';
  import { busyCopy, withRetry } from '../../lib/ui/copy';
  import Button from '../ui/Button.svelte';
  import CharacterName from '../ui/CharacterName.svelte';

  let {
    characters,
    onAnswered,
    answer = answerRealm,
    testid = 'unknown-realm',
  }: {
    characters: readonly MeCharacter[];
    /** The answer saved: the page reads the account again. */
    onAnswered?: (oldKey: string, newKey: string) => void;
    /** The request; a test or a story passes its own. */
    answer?: (key: string, ruleset: string) => Promise<string>;
    testid?: string;
  } = $props();

  /** The character key whose answer is in flight, and the ruleset pressed. */
  let pending = $state<{ key: string; ruleset: string } | null>(null);
  let errors = $state<Record<string, string>>({});

  /** The API's sentence, capitalised and closed, then the retry word. */
  function refusal(error: unknown): string {
    const raw =
      error instanceof Error && error.message.trim() !== '' ? error.message.trim() : unknownRealmCopy.failed;
    const sentence = `${raw.charAt(0).toUpperCase()}${raw.slice(1)}`;
    return withRetry(sentence.endsWith('.') ? sentence : `${sentence}.`);
  }

  async function choose(key: string, ruleset: string): Promise<void> {
    if (pending !== null) return;
    pending = { key, ruleset };
    errors = { ...errors, [key]: '' };
    try {
      const newKey = await answer(key, ruleset);
      onAnswered?.(key, newKey);
    } catch (error) {
      errors = { ...errors, [key]: refusal(error) };
    } finally {
      pending = null;
    }
  }
</script>

{#if characters.length > 0}
  <div class="unknown-realm border-line-soft border-b" data-testid={testid}>
    {#each characters as character (character.key)}
      {@const headingId = `${testid}-${character.key.replaceAll('/', '-')}`}
      <section
        class="unknown-realm-item flex flex-col gap-3 px-4 py-4"
        aria-labelledby={headingId}
        data-testid={`${testid}-item`}
      >
        <div class="flex flex-col gap-1">
          <h2 id={headingId} class="text-text text-[17px] leading-6 font-semibold">
            {unknownRealmCopy.questionLead}
            <CharacterName name={character.name} cls={character.class} crestSize={20} class="inline-flex" />
            {unknownRealmCopy.questionTail}
          </h2>
          <p class="text-muted text-[14px] leading-5">{unknownRealmCopy.why} {unknownRealmCopy.rpPvp}</p>
        </div>
        <div class="unknown-realm-choices" role="group" aria-labelledby={headingId}>
          {#each REALM_CHOICES as choice (choice.id)}
            <Button
              size="md"
              busy={pending?.key === character.key && pending.ruleset === choice.id}
              busyLabel={busyCopy.saving}
              disabled={pending !== null && pending.key !== character.key}
              onclick={() => void choose(character.key, choice.id)}
              testid={`${testid}-choice-${choice.id}`}
            >
              {choice.label}
            </Button>
          {/each}
        </div>
        {#if errors[character.key]}
          <p class="text-wipe text-[14px] leading-5" role="alert">{errors[character.key]}</p>
        {/if}
      </section>
    {/each}
  </div>
{/if}

<style>
  /* The next thing to do on the plate: the gold rule a step that waits on the player carries. */
  .unknown-realm-item {
    box-shadow: inset 2px 0 0 var(--color-gold);
  }
  .unknown-realm-item + .unknown-realm-item {
    border-top: 1px solid var(--color-line-soft);
  }
  /* Four equal buttons, so "Saving…" never moves its neighbours; two by two on a phone. */
  .unknown-realm-choices {
    display: grid;
    grid-template-columns: repeat(4, minmax(0, 7.5rem));
    gap: 8px;
  }
  .unknown-realm-choices :global(button) {
    width: 100%;
  }
  @media (max-width: 599px) {
    .unknown-realm-choices {
      grid-template-columns: repeat(2, minmax(0, 1fr));
    }
  }
</style>
