package gg.foreversixty.app.rankings

import gg.foreversixty.app.home.HomeFixtures
import gg.foreversixty.app.model.CharacterProfile
import gg.foreversixty.app.model.ItemsFile
import gg.foreversixty.app.model.Me
import gg.foreversixty.app.model.Phases
import gg.foreversixty.app.model.RankedEncounters
import gg.foreversixty.app.model.RankingsPage
import gg.foreversixty.app.model.WatchedCharacter
import gg.foreversixty.app.model.WatchingList
import gg.foreversixty.app.network.ApiResult
import gg.foreversixty.app.raids.RaidFixture
import gg.foreversixty.app.raids.ok
import gg.foreversixty.app.rankings.logic.RankingsQuery
import gg.foreversixty.app.rankings.logic.RankingsService
import gg.foreversixty.app.rankings.logic.RankingsViewer
import gg.foreversixty.app.rankings.logic.ViewerCharacter
import gg.foreversixty.app.sims.SimsTestData
import java.time.Instant
import java.time.ZoneId

/**
 * F11's test data, the same story as iOS's RankingsTestData: the contract fixtures (Lightningbolltz
 * 14th of 212 Elemental Shamans on Onyxia; Sparkwick pinned 204th; Voltaris's character card).
 */
object RankingsTestData {
    /** The day after the fixtures' newest kill (2026-12-16 20:00Z), inside Phase 1. */
    val now: Instant = Instant.parse("2026-12-17T20:00:00Z")
    val zone: ZoneId = ZoneId.of("America/Chicago")
    const val BUILD = "1.60.1.70707"

    fun page(): RankingsPage = RaidFixture.load(RankingsPage.serializer(), "rankings.list")
    fun offPage(): RankingsPage = RaidFixture.load(RankingsPage.serializer(), "rankings.list.off-page")
    fun empty(): RankingsPage = RaidFixture.load(RankingsPage.serializer(), "rankings.list.empty")
    fun encounters(): RankedEncounters = RaidFixture.load(RankedEncounters.serializer(), "rankings.encounters")
    fun phases(): Phases = RaidFixture.load(Phases.serializer(), "phases")
    fun watching(): WatchingList = RaidFixture.load(WatchingList.serializer(), "watching.list")
    fun watched(): WatchedCharacter = RaidFixture.load(WatchedCharacter.serializer(), "watching.put")
    fun voltaris(): CharacterProfile = RaidFixture.load(CharacterProfile.serializer(), "character.card.ranked")
    fun yours(): CharacterProfile = RaidFixture.load(CharacterProfile.serializer(), "character.card.yours")
    fun noLogs(): CharacterProfile = RaidFixture.load(CharacterProfile.serializer(), "character.card.empty")
    fun items(): ItemsFile = HomeFixtures.staticFile("cdn.items.shaman", ItemsFile.serializer())
    fun me(): Me = RaidFixture.load(Me.serializer(), "me.premium")
    fun specs() = SimsTestData.specs()

    /** Lightningbolltz, an Elemental Shaman on US Normal: the board's own row on page one. */
    val lightningbolltz = ViewerCharacter(
        key = "us/normal/lightningbolltz", name = "Lightningbolltz", region = "us", ruleset = "normal",
        classSlug = "shaman", specName = "Elemental",
    )

    /** Sparkwick, the off-page character (204th): another character of the same account. */
    val sparkwick = lightningbolltz.copy(key = "us/normal/sparkwick", name = "Sparkwick")

    fun viewer(current: ViewerCharacter = lightningbolltz, history: Boolean = true): RankingsViewer = RankingsViewer(
        signedIn = true, current = current, accountKeys = setOf(lightningbolltz.key, sparkwick.key), history = history,
    )
}

/** A scripted rankings service: each call answers from its lambda; every call is recorded. */
class FakeRankingsService : RankingsService {
    var rankingsAnswer: (RankingsQuery) -> ApiResult<RankingsPage> = { ok(RankingsTestData.page()) }
    var encountersAnswer: () -> ApiResult<RankedEncounters> = { ok(RankingsTestData.encounters()) }
    var phasesAnswer: () -> ApiResult<Phases> = { ok(RankingsTestData.phases()) }
    var characterAnswer: (String) -> ApiResult<CharacterProfile> = { ok(RankingsTestData.voltaris()) }
    var watchingAnswer: () -> ApiResult<WatchingList> = { ok(RankingsTestData.watching()) }
    var watchAnswer: () -> ApiResult<WatchedCharacter> = { ok(RankingsTestData.watched()) }
    var unwatchAnswer: () -> ApiResult<Unit> = { ok(Unit) }
    val queries = mutableListOf<RankingsQuery>()
    val calls = mutableListOf<String>()

    override suspend fun rankings(query: RankingsQuery): ApiResult<RankingsPage> {
        queries += query
        calls += "rankings ${query.page}"
        return rankingsAnswer(query)
    }
    override suspend fun encounters(): ApiResult<RankedEncounters> {
        calls += "encounters"
        return encountersAnswer()
    }
    override suspend fun phases(): ApiResult<Phases> {
        calls += "phases"
        return phasesAnswer()
    }
    override suspend fun character(region: String, ruleset: String, name: String): ApiResult<CharacterProfile> {
        calls += "character $region/$ruleset/$name"
        return characterAnswer(name)
    }
    override suspend fun watching(): ApiResult<WatchingList> {
        calls += "watching"
        return watchingAnswer()
    }
    override suspend fun watch(region: String, ruleset: String, name: String): ApiResult<WatchedCharacter> {
        calls += "watch $name"
        return watchAnswer()
    }
    override suspend fun unwatch(region: String, ruleset: String, name: String): ApiResult<Unit> {
        calls += "unwatch $name"
        return unwatchAnswer()
    }
}
