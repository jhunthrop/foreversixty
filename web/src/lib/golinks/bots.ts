// web/src/lib/golinks/bots.ts
// Which /go/ clicks are a crawler or a link preview rather than a person. Every chat app and
// social site fetches a pasted link to draw its card, so without this a link posted in one
// busy Discord would count each unfurl as a visit. The user agent is read here, at the edge,
// and never sent on or stored: the API only hears "bot" or "not".
const BOT_AGENT =
  /bot\b|bot\/|crawl|spider|slurp|preview|externalhit|embedly|whatsapp|skypeuripreview|bitlybot|headless|lighthouse|curl\/|wget|python-|go-http|okhttp|java\/|axios|node-fetch|httpclient|monitor/i;

export function isBotAgent(agent: string): boolean {
  return agent.trim() === '' || BOT_AGENT.test(agent);
}
