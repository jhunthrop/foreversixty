<!-- web/src/components/account/AccountMenu.svelte -->
<!-- The header's account chip and menu (spec 2026-09-25, section 2 and 3.3), replacing
     SessionNav.svelte's plain name link. Signed out: the plain "Sign in" link, as before
     (spec section 2, "Signed out it is the Battle.net button, as today") -- SessionNav.svelte
     never linked straight to battlenetStartUrl either, it linked to /login, which does; this
     keeps that same indirection. Signed in: an avatar chip (the main character's portrait,
     else the initial letter of the battletag) that opens a native <details> menu: the
     character list (switch the current-character pointer), Your account, Devices, Plan, Sign
     out. The invisible placeholder below is load-bearing -- see
     web/tests/e2e/sim-tools-phone.spec.ts's `load()` helper, which waits for it to swap to
     real content before measuring hit targets, and SessionNav.svelte's own copy of this same
     trick (kept here verbatim) for the CLS story. -->
<script lang="ts">
  import { fetchMeOnce, signOut, type Me, type MeCharacter } from '../../lib/account/api';
  import { createQueryState } from '../../lib/data/query.svelte';
  import { API_BASE_URL } from '../../lib/planner/config';
  import { mainCharacter, pointerForCharacter } from '../../lib/account/main-character';
  import { writeCurrent, CURRENT_CHARACTER_CHANGED } from '../../lib/current-character';
  import CharacterPortrait from '../character/CharacterPortrait.svelte';

  // One `/v1/me` read, shared with every other island through the client cache
  // (web/src/lib/data/query.ts) -- see Account.svelte and HomeAccountPanel.svelte's own
  // copies of this same call.
  const session = createQueryState<Me | null>(`${API_BASE_URL}/v1/me`, () => fetchMeOnce(), {
    scope: 'private',
    ttlMs: 10 * 60 * 1000,
  });

  const me = $derived(session.data);
  const signedIn = $derived(me !== null);
  const battletag = $derived(me?.user.battletag ?? me?.user.email ?? 'Your account');
  const initial = $derived(battletag.charAt(0).toUpperCase());
  const main = $derived(me !== null ? mainCharacter(me.characters, me.main_character_key) : null);

  let open = $state(false);
  let root: HTMLElement;

  function close(): void {
    open = false;
  }

  function onSwitch(character: MeCharacter): void {
    writeCurrent(pointerForCharacter(character));
    window.dispatchEvent(new Event(CURRENT_CHARACTER_CHANGED));
    close();
  }

  async function onSignOut(): Promise<void> {
    await signOut();
    window.location.assign('/');
  }

  function onDocumentPointerDown(event: PointerEvent): void {
    const target = event.target;
    if (!open || !(target instanceof Node) || root.contains(target)) return;
    close();
  }

  function onDocumentKeydown(event: KeyboardEvent): void {
    if (open && event.key === 'Escape') close();
  }

  $effect(() => {
    document.addEventListener('pointerdown', onDocumentPointerDown);
    document.addEventListener('keydown', onDocumentKeydown);
    return () => {
      document.removeEventListener('pointerdown', onDocumentPointerDown);
      document.removeEventListener('keydown', onDocumentKeydown);
    };
  });
</script>

<div class="flex items-center gap-3 text-[13px]" data-testid="session-nav" bind:this={root}>
  {#if session.status === 'loading'}
    <span class="invisible inline-flex min-h-11 items-center px-2 md:min-h-0 md:px-0" aria-hidden="true">
      Sign in
    </span>
  {:else if signedIn}
    <details bind:open data-testid="account-menu" class="relative">
      <summary class="flex min-h-11 list-none items-center gap-2 px-1 marker:content-none md:min-h-0">
        {#if main !== null}
          <CharacterPortrait character={main} size="sm" testid="account-menu-portrait" />
        {:else}
          <span
            class="bg-raised border-line text-text flex h-7 w-7 shrink-0 items-center justify-center rounded-full border text-[12px] font-bold"
            >{initial}</span
          >
        {/if}
        <span class="text-nav hover:text-strong">{battletag}</span>
      </summary>
      <div
        class="bg-raised border-line rounded-panel absolute top-full right-0 z-30 mt-2 flex w-[220px] flex-col border py-2"
      >
        {#if me !== null && me.characters.length > 0}
          <div class="border-line-soft flex flex-col border-b pb-2">
            {#each me.characters as character (character.key)}
              <button
                type="button"
                class="text-nav hover:text-strong flex min-h-11 items-center gap-2 px-4 text-left"
                data-testid={`account-menu-character-${character.key}`}
                onclick={() => onSwitch(character)}
              >
                <CharacterPortrait {character} size="sm" testid={`account-menu-portrait-${character.key}`} />
                <span>{character.name}</span>
              </button>
            {/each}
          </div>
        {/if}
        <a class="text-nav hover:text-strong flex min-h-11 items-center px-4" href="/account">Your account</a>
        <a class="text-nav hover:text-strong flex min-h-11 items-center px-4" href="/account#devices"
          >Devices</a
        >
        <a class="text-nav hover:text-strong flex min-h-11 items-center px-4" href="/account#plan">Plan</a>
        <button
          type="button"
          class="text-nav hover:text-strong flex min-h-11 items-center px-4 text-left"
          data-testid="account-menu-signout"
          onclick={() => void onSignOut()}
        >
          Sign out
        </button>
      </div>
    </details>
  {:else}
    <a
      href="/login"
      class="text-nav hover:text-strong inline-flex min-h-11 items-center px-2 md:min-h-0 md:px-0"
    >
      Sign in
    </a>
  {/if}
</div>
