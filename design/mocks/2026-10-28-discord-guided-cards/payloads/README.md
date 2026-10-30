Real handler answers for every guided-card and raid-post board, one `<frame>.json` each (id, title, clock, message; `above` for a reply; `modal` for a dialog).
Regenerate from the repo root, on a private test database:

    cd api && TEST_DATABASE_URL=<private db> GUIDED_PAYLOADS_DIR=$PWD/../design/mocks/2026-10-28-discord-guided-cards/payloads go test -tags payloaddump -p 1 -count=1 -run TestDumpGuidedPayloads ./internal/discordbot/

`gen_guided_cards.mjs` draws each file as `<frame>-real-{desktop,phone}.png` beside the board of the same name.
Test-world values differ from the boards on purpose (base URLs `site.test`/`api.test`, item names, raid ids, nonces); the report lists only differences that are not that.
