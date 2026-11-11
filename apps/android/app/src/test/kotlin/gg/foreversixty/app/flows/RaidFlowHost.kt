package gg.foreversixty.app.flows

import androidx.activity.compose.BackHandler
import androidx.compose.foundation.background
import androidx.compose.foundation.layout.Column
import androidx.compose.foundation.layout.fillMaxSize
import androidx.compose.runtime.Composable
import androidx.compose.runtime.LaunchedEffect
import androidx.compose.runtime.getValue
import androidx.compose.runtime.mutableIntStateOf
import androidx.compose.runtime.mutableStateOf
import androidx.compose.runtime.remember
import androidx.compose.runtime.rememberCoroutineScope
import androidx.compose.runtime.setValue
import androidx.compose.ui.Modifier
import gg.foreversixty.app.common.SavedAnswerStrip
import gg.foreversixty.app.designsystem.FsColor
import gg.foreversixty.app.home.flows.HomeActions
import gg.foreversixty.app.menu.CharacterBar
import gg.foreversixty.app.raids.MemoryAskStore
import gg.foreversixty.app.raids.MemoryRaidCache
import gg.foreversixty.app.raids.RaidsDestination
import gg.foreversixty.app.raids.RaidsNavigator
import gg.foreversixty.app.raids.RecordingHaptics
import gg.foreversixty.app.raids.clock
import gg.foreversixty.app.raids.flows.RaidNightRoute
import gg.foreversixty.app.raids.flows.RaidTabActions
import gg.foreversixty.app.raids.flows.RaidTabExtras
import gg.foreversixty.app.raids.flows.NextRaidSignup
import androidx.compose.foundation.ScrollState
import androidx.compose.runtime.CompositionLocalProvider
import gg.foreversixty.app.common.LocalHaptics
import androidx.compose.foundation.rememberScrollState
import gg.foreversixty.app.raids.flows.RaidTabPresenter
import gg.foreversixty.app.raids.flows.RaidTabScreen
import gg.foreversixty.app.raids.logic.NotificationPermission
import gg.foreversixty.app.raids.logic.RaidNightModel
import gg.foreversixty.app.raids.logic.RaidPageModel
import gg.foreversixty.app.raids.logic.RaidsTabModel
import gg.foreversixty.app.raids.logic.RaidsTabState
import gg.foreversixty.app.raids.logic.SignupController
import gg.foreversixty.app.raids.resources
import gg.foreversixty.app.raids.ui.RaidPageScreen
import java.time.Instant
import kotlin.time.Duration.Companion.seconds
import kotlinx.coroutines.CoroutineScope
import kotlinx.coroutines.launch

/**
 * The Raids tab as [gg.foreversixty.app.raids.RaidsRoute] wires it, over a fake clock: the tab model,
 * the next raid's [SignupController] (the route's [NextRaidSignup]), the root's answer taps (the route's
 * private `raidTabActions`) and the stack's Raid and Night pages. RaidsRoute itself cannot be hosted
 * deterministically: RaidsFeature fixes `now` to Instant::now and builds its own LiveRaidService.
 */
class RaidFlow(
    val service: EchoRaidService = EchoRaidService(),
    val haptics: RecordingHaptics = RecordingHaptics(),
    val now: () -> Instant = { FlowTime.now },
) {
    val navigator = RaidsNavigator()

    /** The session's current character, RaidsRoute's `characterKey`; null follows the tab's. */
    var currentKey: String? = null

    /** The root's scroll as the screen hoists it, for the push-and-pop flow. */
    var rootScroll: ScrollState? = null
        private set

    /** RaidFlowScreen hands over the root's scroll. */
    fun bindRoot(scroll: ScrollState) {
        rootScroll = scroll
    }
    private val cache = MemoryRaidCache()
    val tab = RaidsTabModel(service, cache, resources(), now)

    /** RaidsRoute's NextRaidSignup: the next raid's controller, and the page's writes shared with it. */
    val signup = NextRaidSignup(tab, now)
    val next: SignupController? get() = signup.controller

    /** The screen's scope and dependencies, once composed. */
    private var bound: Pair<CoroutineScope, SignupController.Dependencies>? = null

    /** Reloads the tab in the screen's scope, as the route's retry does. */
    fun reloadNow() {
        val (scope, deps) = checkNotNull(bound) { "RaidFlowScreen is not composed" }
        scope.launch { reload(deps) }
    }

    fun bind(scope: CoroutineScope): SignupController.Dependencies =
        deps(scope).also { bound = scope to it }

    private fun deps(scope: CoroutineScope) = SignupController.Dependencies(
        service = service,
        res = resources(),
        haptics = haptics,
        permission = NotificationPermission { false },
        askStore = MemoryAskStore(),
        now = now,
        scope = scope,
        hold = HOLD,
    )

    suspend fun reload(deps: SignupController.Dependencies) {
        tab.currentKey = FlowPlayer.KEY
        tab.load()
        signup.sync(deps)
    }

    /** RaidsRoute's root taps: the answers through the next raid's controller, the raid pushed. */
    fun actions(scope: CoroutineScope): RaidTabActions {
        val content = (tab.state as? RaidsTabState.Loaded)?.content
        val key = content?.character?.key
        val guild = content?.head?.id
        return RaidTabActions(
            card = HomeActions(
                onAnswer = { _, state -> if (key != null) scope.launch { next?.press(key, state) } },
                onLate = { _, minutes -> if (key != null) scope.launch { next?.pressLate(key, minutes) } },
                onRetryAnswer = { if (key != null) scope.launch { next?.retry(key) } },
                onRaid = { raid -> guild?.let { navigator.push(RaidsDestination.Raid(it, raid)) } },
            ),
        )
    }

    companion object {
        /** RaidsRoute's HOLD: `Saved.` stays 4 s. */
        val HOLD = 4.seconds
    }
}

/** The Raids tab and its stack, inside the tab shell. */
@Composable
fun RaidFlowScreen(flow: RaidFlow) {
    // The light tick at the pull's trigger point (R14) lands on the flow's recorder.
    CompositionLocalProvider(LocalHaptics provides flow.haptics) { RaidFlowBody(flow) }
}

@Composable
private fun RaidFlowBody(flow: RaidFlow) {
    val scope = rememberCoroutineScope()
    val deps = remember(flow) { flow.bind(scope) }
    val clock = remember { clock() }
    LaunchedEffect(flow) { flow.reload(deps) }
    // RaidsRoute hoists the root's scroll above the stack, so a push and pop keep the root's place.
    val rootScroll = rememberScrollState().also(flow::bindRoot)
    val refresh = rememberFlowRefresh { flow.reload(deps) }
    BackHandler(enabled = flow.navigator.stack.isNotEmpty()) { flow.navigator.pop() }
    RaidFlowStack(flow, deps, clock, rootScroll, refresh, scope)
}

@Composable
private fun RaidFlowStack(
    flow: RaidFlow,
    deps: SignupController.Dependencies,
    clock: gg.foreversixty.app.raids.logic.RaidClock,
    rootScroll: ScrollState,
    refresh: FlowRefresh,
    scope: CoroutineScope,
) {
    val content = (flow.tab.state as? RaidsTabState.Loaded)?.content
    when (val top = flow.navigator.stack.lastOrNull()) {
        null -> Column(Modifier.fillMaxSize().background(FsColor.bg)) {
            val view = RaidTabPresenter(resources(), clock, flow.now()).view(flow.tab.state, flow.next, null, false, RaidTabExtras())
            if (view != null) {
                RaidTabScreen(view, flow.actions(scope), topBar = { CharacterBar(content?.character, "Raid") }, page = FlowPageState(rootScroll, refresh)) {
                    content?.offlineSince?.let { SavedAnswerStrip(clock.stripTime(it, flow.now()), content.stale, flow::reloadNow) }
                }
            }
        }
        is RaidsDestination.Raid -> {
            val model = remember(top) {
                RaidPageModel(top.guild, top.raid, deps, MemoryRaidCache(), top.open, flow.signup::replaced, currentKey = flow.currentKey ?: content?.character?.key)
            }
            RaidPageScreen(
                model = model,
                clock = clock,
                now = flow.now,
                currentKey = content?.character?.key,
                guildName = content?.head?.name.orEmpty(),
                guildUrl = content?.head?.webUrl,
                openNight = { flow.navigator.push(RaidsDestination.Night(it)) },
                back = { flow.navigator.pop() },
                requestPermission = {},
            )
        }
        is RaidsDestination.Night -> {
            val model = remember(top) { RaidNightModel(top.raid, flow.service, MemoryRaidCache(), flow.now) }
            RaidNightRoute(
                model, clock, flow.now, content?.planned?.firstOrNull { it.id == top.raid }?.compTarget, null, null,
                openRoster = { content?.head?.id?.let { flow.navigator.push(RaidsDestination.Raid(it, top.raid)) } },
                openWeb = {},
            ) { flow.navigator.pop() }
        }
        else -> error("the raid flows open no $top")
    }
}
