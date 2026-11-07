// web/src/lib/not-found/answer-mount.ts
// Paints the 404's answer (ui/CharacterAnswer's markup, `#nf-answer`) after load, only when a
// session is hinted, and again on every change of the header selector's character; only the
// newest read paints. The guides' painter draws it; the third link is the smaller one.
import { fetchMeOnce } from '../account/api';
import { onCurrentCharacterChange, readCurrent } from '../current-character';
import { paintGuideAnswer } from '../guides/guide-answer-mount';
import { afterLoad, hasSessionHint } from '../page-load';
import { CHARACTER_LINK_INDEX, notFoundAnswerFrom } from './answer';

export function mountNotFoundAnswer(): void {
  const root = document.getElementById('nf-answer');
  if (root === null) return;
  if (!hasSessionHint()) {
    root.dataset.state = 'hidden';
    return;
  }
  let latest = 0;
  const answer = async (): Promise<void> => {
    const asked = ++latest;
    let found;
    try {
      found = notFoundAnswerFrom(await fetchMeOnce(), readCurrent());
    } catch {
      found = notFoundAnswerFrom(null, null, true);
    }
    if (asked !== latest) return;
    paintGuideAnswer(root, found, () => void answer());
    const small = root.querySelectorAll<HTMLElement>('[data-field="links"] a')[CHARACTER_LINK_INDEX];
    if (small !== undefined) small.dataset.size = 'small';
  };
  afterLoad(() => {
    void answer();
    onCurrentCharacterChange(() => void answer());
  });
}
