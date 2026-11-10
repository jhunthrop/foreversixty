package gg.foreversixty.app.flows

import androidx.compose.foundation.background
import androidx.compose.foundation.layout.Box
import androidx.compose.foundation.layout.fillMaxSize
import androidx.compose.runtime.Composable
import androidx.compose.runtime.CompositionLocalProvider
import androidx.compose.ui.Modifier
import androidx.compose.ui.semantics.SemanticsProperties
import androidx.compose.ui.test.SemanticsMatcher
import androidx.compose.ui.test.SemanticsNodeInteraction
import androidx.compose.ui.test.hasClickAction
import androidx.compose.ui.test.junit4.ComposeContentTestRule
import androidx.compose.ui.test.onRoot
import androidx.compose.ui.test.performTouchInput
import com.github.takahirom.roborazzi.ExperimentalRoborazziApi
import com.github.takahirom.roborazzi.captureRoboImage
import com.github.takahirom.roborazzi.captureScreenRoboImage
import gg.foreversixty.app.activity.ActivityFixtures
import gg.foreversixty.app.designsystem.ForeverTheme
import gg.foreversixty.app.designsystem.FsColor
import gg.foreversixty.app.items.LocalItemIconPainter
import gg.foreversixty.app.menu.LocalOpenMenu
import gg.foreversixty.app.menu.LocalPortraitPainter
import gg.foreversixty.app.menu.MenuFixtures
import gg.foreversixty.app.model.GuildRaids
import gg.foreversixty.app.model.RaidView
import gg.foreversixty.app.model.SignupAnswer
import gg.foreversixty.app.model.SignupRequest
import gg.foreversixty.app.navigation.ForeverTab
import gg.foreversixty.app.navigation.ForeverTabScaffold
import gg.foreversixty.app.network.ApiClient
import gg.foreversixty.app.network.ApiConfig
import gg.foreversixty.app.network.ApiError
import gg.foreversixty.app.network.ApiMode
import gg.foreversixty.app.network.ApiResult
import gg.foreversixty.app.network.ClasspathFixtureSource
import gg.foreversixty.app.network.FixtureEngine
import gg.foreversixty.app.network.ForeverHttpClient
import gg.foreversixty.app.network.TokenProvider
import gg.foreversixty.app.raids.logic.LiveRaidService
import gg.foreversixty.app.raids.logic.RaidService
import gg.foreversixty.app.raids.ui.LocalRaidIconBuild
import gg.foreversixty.app.testing.STRICT_SNAPSHOT
import java.io.IOException
import java.time.Instant
import kotlinx.coroutines.CompletableDeferred

/** The flows' moment: raid day, 2 h 48 min before the fixture raid's pull (31, Onyxia's Lair, 2026-12-20 01:00Z). */
object FlowTime {
    val now: Instant = Instant.parse("2026-12-19T22:12:00Z")
}

/** The contract fixtures' player: `me`'s main, signed In for raid 31 of guild 2. */
object FlowPlayer {
    const val KEY = "us/pvp/gorvash"
    const val GUILD = 2L
    const val RAID = 31L
}

/** How long a flow waits for a write through the fixture engine to come back. */
const val SETTLE_MILLIS = 5_000L

/**
 * The raids service in fixture mode: every read and write goes through the real [LiveRaidService]
 * over the contract fixtures, and a sign-up answers with the last raid view echoing the request's
 * state and late_minutes into the player's own sign-up (the fixture's canned answer never changes).
 * [failWith] fails the next writes, [gate] holds a write in flight, [offline] drops every call.
 */
class EchoRaidService(private val live: RaidService = fixtureRaidService()) : RaidService by live {
    val requests = mutableListOf<Pair<String, SignupRequest>>()
    var failWith: ApiError? = null
    var gate: CompletableDeferred<Unit>? = null
    var offline: Boolean = false
    var raidOverride: RaidView? = null

    /** Writes sent and not yet answered, so a flow waits for the fixture engine's thread. */
    @Volatile var inFlight: Int = 0
        private set

    private val views = mutableMapOf<Long, RaidView>()

    /** The guild raid list's reads: a load, a retry or a pull to refresh. */
    var raidsReads: Int = 0
        private set

    override suspend fun me() = if (offline) dropped() else live.me()

    override suspend fun raids(guild: Long): ApiResult<GuildRaids> {
        raidsReads += 1
        if (offline) return dropped()
        return live.raids(guild).also { result ->
            if (result is ApiResult.Success) result.data.planned.orEmpty().forEach { views[it.id] = it }
        }
    }

    override suspend fun raid(guild: Long, raid: Long): ApiResult<RaidView> {
        if (offline) return dropped()
        raidOverride?.let { return ApiResult.Success(it.also { view -> views[view.id] = view }, "flow") }
        return live.raid(guild, raid).also { if (it is ApiResult.Success) views[raid] = it.data }
    }

    override suspend fun signup(guild: Long, raid: Long, characterKey: String, request: SignupRequest): ApiResult<SignupAnswer> {
        requests += characterKey to request
        inFlight += 1
        try {
            gate?.await()
            if (offline) return dropped()
            failWith?.let { return ApiResult.Failure(it) }
            // The real write over the fixture engine: the route, the body's encoding and the answer's decoding.
            val canned = live.signup(guild, raid, characterKey, request)
            if (canned !is ApiResult.Success) return canned
            val echoed = echo(views[raid] ?: canned.data.raid, characterKey, request)
            views[raid] = echoed
            return ApiResult.Success(canned.data.copy(raid = echoed, moved = null), canned.requestId)
        } finally {
            inFlight -= 1
        }
    }

    private fun dropped(): ApiResult<Nothing> = ApiResult.Failure(ApiError.Network(IOException("offline")))

    companion object {
        /** [view] with [key]'s answer set to the request's state and minutes, as the server would store it. */
        fun echo(view: RaidView, key: String, request: SignupRequest): RaidView = view.copy(
            viewer = view.viewer.copy(
                characters = view.viewer.characters.map { if (it.characterKey == key) it.copy(state = request.state) else it },
            ),
            signups = view.signups.map {
                if (it.characterKey == key) it.copy(state = request.state, lateMinutes = request.lateMinutes) else it
            },
        )
    }
}

/** The fixture-mode API client the debug fixture builds use: the contract fixtures, no network. */
fun fixtureRaidService(): RaidService {
    val source = ClasspathFixtureSource(requireNotNull(EchoRaidService::class.java.classLoader))
    val config = ApiConfig(baseUrl = "https://api.example.test", mode = ApiMode.Fixture)
    return LiveRaidService(ApiClient(ForeverHttpClient.create(config, NoToken, FixtureEngine.create(source))))
}

private object NoToken : TokenProvider {
    override suspend fun bearerToken(): String? = null
    override suspend fun discardBearerToken() = Unit
}

/** The app as a flow capture draws it: the theme on the ground, the tab shell and the repo's icons and portraits. */
@Composable
fun FlowFrame(tab: ForeverTab, content: @Composable () -> Unit) {
    ForeverTheme {
        Box(Modifier.fillMaxSize().background(FsColor.bg)) {
            ForeverTabScaffold(selected = tab, onSelect = {}) {
                CompositionLocalProvider(
                    // A ripple fades on the wall clock under Robolectric and lands on a different frame each
                    // run; the steps hold the state, the emulator captures hold the press.
                    androidx.compose.material3.LocalRippleConfiguration provides null,
                    LocalOpenMenu provides {},
                    LocalRaidIconBuild provides ActivityFixtures.BUILD,
                    LocalItemIconPainter provides gg.foreversixty.app.home.HomeFixtures.repoIcons,
                    LocalPortraitPainter provides MenuFixtures.repoPortraits,
                ) { content() }
            }
        }
    }
}

/** Longer than a Material ripple's press and fade (225 ms and 150 ms). */
private const val RIPPLE_SETTLE_MS = 1_000L

/** One step of a flow: src/test/screenshots/flows/<flow>-<NN>-<step>.png, compared exactly. */
class FlowCapture(private val rule: ComposeContentTestRule, private val flow: String) {
    fun step(number: Int, step: String) {
        settle()
        rule.onRoot().captureRoboImage(path(number, step), roborazziOptions = STRICT_SNAPSHOT)
    }

    /** Every window on screen: a bottom sheet or dialog draws in a window of its own. */
    @OptIn(ExperimentalRoborazziApi::class)
    fun screen(number: Int, step: String) {
        settle()
        captureScreenRoboImage(path(number, step), roborazziOptions = STRICT_SNAPSHOT)
    }

    /** Idle, then a fixed stretch of the test clock: a press's ripple lands on the same frame every run. */
    private fun settle() {
        rule.waitForIdle()
        rule.mainClock.advanceTimeBy(RIPPLE_SETTLE_MS)
        rule.waitForIdle()
    }

    private fun path(number: Int, step: String) =
        "src/test/screenshots/flows/$flow-${number.toString().padStart(2, '0')}-$step.png"

    /** The finger down on [node] for the pressed capture, then lifted: the lift is the tap. */
    fun pressed(node: SemanticsNodeInteraction, number: Int, step: String) {
        node.performTouchInput { down(center) }
        step(number, step)
        node.performTouchInput { up() }
    }
}

/** A segment or chip: clickable and carrying its selected state (a head row's answer chip is neither). */
val SELECTABLE: SemanticsMatcher = hasClickAction() and SemanticsMatcher.keyIsDefined(SemanticsProperties.Selected)

/** Waits until every write the flow sent has its answer drawn. */
fun ComposeContentTestRule.settle(service: EchoRaidService) {
    waitUntil(SETTLE_MILLIS) { service.inFlight == 0 }
    waitForIdle()
}
