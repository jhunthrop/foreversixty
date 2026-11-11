<!-- web/src/components/guild/dense/NeedsAWord.svelte -->
<!-- Who the officer talks to before the raid (dense spec 5.1): one Row per raider, the reasons in
     figures, one secondary action: Ping (not answered), Profile (no Discord to ping), Nudge (short
     for the raid, the readiness board's nudge) or Report (rating fell). -->
<script lang="ts">
  import type { Word, WordAction } from '../../../lib/guild/dense/week-board';
  import { CopyState } from '../../../lib/ui/copy-state.svelte';
  import Button from '../../ui/Button.svelte';
  import Row from '../../ui/Row.svelte';
  import Section from '../../ui/Section.svelte';
  import ClassCrestRing from '../../character/ClassCrestRing.svelte';
  import type { ClassSlug } from '../../../lib/ui/primitives';

  let {
    words,
    hrefOf,
    onPing,
    nudgeTextOf,
  }: {
    words: readonly Word[];
    /** Where Profile and Report go for a raider. */
    hrefOf: (word: Word, action: WordAction) => string;
    onPing: () => void;
    /** The raid's nudge for the raider, which Nudge copies ('' when there is none to send). */
    nudgeTextOf: (word: Word) => string;
  } = $props();

  const nudges = new CopyState();
  $effect(() => () => nudges.dispose());

  function nudge(word: Word): void {
    const text = nudgeTextOf(word);
    if (text !== '') void nudges.copy(text, word.key);
  }

  const LABEL: Readonly<Record<WordAction, string>> = {
    profile: 'Profile',
    ping: 'Ping',
    nudge: 'Nudge',
    report: 'Report',
  };
</script>

{#if words.length > 0}
  <Section title="Needs a word" meta={`${words.length} raiders`} panel testid="guild-week-words">
    {#each words as word (word.key)}
      <Row
        title={word.name}
        titleTone={{ classSlug: word.cls as ClassSlug }}
        sub={word.reasons.join(' · ')}
        open
      >
        {#snippet lead()}<ClassCrestRing characterClass={word.cls} size={32} />{/snippet}
        {#snippet trail()}
          {#if word.action === 'ping'}
            <Button tone="ghost" size="sm" onclick={onPing}>{LABEL.ping}</Button>
          {:else if word.action === 'nudge'}
            <Button tone="ghost" size="sm" onclick={() => nudge(word)}
              >{nudges.label(LABEL.nudge, word.key)}</Button
            >
          {:else}
            <Button tone="ghost" size="sm" href={hrefOf(word, word.action)}>{LABEL[word.action]}</Button>
          {/if}
        {/snippet}
      </Row>
    {/each}
  </Section>
{/if}
