package gg.foreversixty.app.home.flows

import androidx.compose.runtime.Composable
import androidx.compose.runtime.LaunchedEffect
import androidx.compose.runtime.getValue
import androidx.compose.runtime.mutableStateOf
import androidx.compose.runtime.remember
import androidx.compose.runtime.rememberCoroutineScope
import androidx.compose.runtime.setValue
import androidx.compose.ui.platform.LocalContext
import androidx.compose.ui.platform.LocalResources
import gg.foreversixty.app.activity.Activity
import gg.foreversixty.app.activity.ActivityFeature
import gg.foreversixty.app.activity.ActivityLanding
import gg.foreversixty.app.common.OfflineStamp
import gg.foreversixty.app.common.SavedAnswerStrip
import gg.foreversixty.app.common.rememberHaptics
import gg.foreversixty.app.home.HomeCharacter
import gg.foreversixty.app.home.HomePhase
import gg.foreversixty.app.home.HomeStore
import gg.foreversixty.app.model.CharacterHome
import gg.foreversixty.app.model.RaidBadge
import gg.foreversixty.app.raids.RaidsDestination
import gg.foreversixty.app.raids.raidDestination
import gg.foreversixty.app.raids.RaidsFeature
import gg.foreversixty.app.raids.logic.RaidClock
import gg.foreversixty.app.raids.logic.RaidsTabModel
import gg.foreversixty.app.raids.logic.SignupController
import gg.foreversixty.app.raids.logic.SystemNotificationPermission
import gg.foreversixty.app.raids.flows.NextRaidSignup
import gg.foreversixty.app.flows.FlowPageState
import gg.foreversixty.app.flows.rememberFlowRefresh
import androidx.compose.foundation.rememberScrollState
import java.time.Instant
import java.time.ZoneId
import kotlin.time.Duration.Companion.seconds
import kotlinx.coroutines.launch

/** Where Home's taps go outside Home: the tabs, an Activity row's answer, an item's tooltip. */
data class HomeExits(
    val openGear: () -> Unit = {},
    val openActivity: () -> Unit = {},
    val openRaidTab: () -> Unit = {},
    val openLanding: (ActivityLanding) -> Unit = {},
    val openItem: (UpgradeLine) -> Unit = {},
    val openSims: () -> Unit = {},
    val openWeb: (String) -> Unit = {},
)

/**
 * The signed-in Home for [character] (spec 2026-11-08 4.1): the upgrades answer from [store], the
 * sections, the raid (answered in place through the Raid tab's controller) and New for you, with
 * [topBar] on the hero.
 */
@Composable
fun HomeFlowContent(
    character: HomeCharacter,
    store: HomeStore,
    phase: HomePhase,
    service: HomeFlowService,
    raids: RaidsFeature?,
    activity: ActivityFeature?,
    exits: HomeExits,
    topBar: @Composable () -> Unit = {},
) {
    val res = LocalResources.current
    val context = LocalContext.current
    val scope = rememberCoroutineScope()
    val haptics = rememberHaptics()
    val tabModel = remember(raids) { raids?.let { RaidsTabModel(it.service, it.cache, res, it.now, connectivity = it.connectivity) } }
    val state = remember(character.key) { HomeFlowState(service, tabModel, activity, Instant::now) }
    // One controller per raid, the Raid tab's rule: the answer saves, its haptics and its offline rule.
    val signup = remember(tabModel) { tabModel?.let { NextRaidSignup(it, raids?.now ?: Instant::now) } }
    val deps = remember(raids, haptics) {
        raids?.let { SignupController.Dependencies(it.service, res, haptics, SystemNotificationPermission(context), it.askStore, it.now, scope, HOLD) }
    }
    val home = (phase as? HomePhase.Loaded)?.home ?: (phase as? HomePhase.Failed)?.cached

    suspend fun load() {
        state.load(character)
        deps?.let { signup?.sync(it) }
        raids?.refreshBadge(state.raidContent?.head?.id)
    }

    LaunchedEffect(character.key) { load() }
    LaunchedEffect(home) { state.refreshBestRun(home) }
    val refresh = rememberFlowRefresh {
        store.refresh(character)
        load()
    }

    val controller = signup?.controller
    val guild = state.raidContent?.head
    val now = Instant.now()
    val view = homeView(character, phase, home, state, controller, HomeComposer(res, now, ZoneId.systemDefault()), RaidCardWords(res, RaidClock(res), now),
        unread = activity?.unread ?: state.rows.count { !it.read },
        badge = raids?.badge,
    )
    val actions = HomeActions(
        onGear = exits.openGear,
        onItem = exits.openItem,
        onAnswer = { _, answer ->
            val key = character.key
            scope.launch { controller?.press(key, answer) }
        },
        onLate = { _, minutes -> scope.launch { controller?.pressLate(character.key, minutes) } },
        onRetryAnswer = { scope.launch { controller?.retry(character.key) } },
        onRaid = { id -> openRaid(raids, guild?.id, id, exits) },
        onRaidNight = { id ->
            raids?.navigator?.let { nav ->
                nav.stack.clear()
                nav.push(RaidsDestination.Night(id))
            }
            exits.openRaidTab()
        },
        onReady = { id -> openRaid(raids, guild?.id, id, exits) },
        onPreviewMention = { id -> openRaid(raids, guild?.id, id, exits, unanswered = true) },
        onApplicants = exits.openRaidTab,
        onActivityRow = { row -> Activity.landing(row)?.let(exits.openLanding) },
        onAllActivity = exits.openActivity,
        onSimAgain = exits.openSims,
        onFindGuild = exits.openRaidTab,
        onNotice = exits.openWeb,
        onRetry = {
            scope.launch {
                store.refresh(character)
                load()
            }
        },
    )
    // E5: offline, the last answer shows under the strip with its time; stale, with Try again.
    (phase as? HomePhase.Loaded)?.offlineSince?.let { since ->
        SavedAnswerStrip(OfflineStamp.format(since, now, ZoneId.systemDefault()), phase.stale, actions.onRetry)
    }
    HomeFlowScreen(view, actions, FlowPageState(rememberScrollState(), refresh), topBar)
}

private fun openRaid(raids: RaidsFeature?, guild: Long?, raid: Long, exits: HomeExits, unanswered: Boolean = false) {
    if (raids != null && guild != null) {
        raids.navigator.stack.clear()
        raids.navigator.push(raidDestination(guild, raid, unanswered))
    }
    exits.openRaidTab()
}

/** Home's view from what has loaded: the head at once, then the blocks once the upgrades answer is in. */
private fun homeView(
    character: HomeCharacter,
    phase: HomePhase,
    home: CharacterHome?,
    state: HomeFlowState,
    controller: SignupController?,
    composer: HomeComposer,
    words: RaidCardWords,
    unread: Int,
    badge: RaidBadge?,
): HomeView {
    val head = composer.head(character)
    if (phase is HomePhase.Loading) return HomeView.Loading(head)
    if (phase is HomePhase.Failed && home == null) return HomeView.Failed(head)
    val guild = state.raidContent?.head?.name.orEmpty()
    val card = state.nextRaid?.let { words.card(it, guild, character.key, controller, badge) }
    val live = state.liveRaid?.let { words.live(it, guild, character.key, words.ownAnswer(it, character.key)) }
    val inputs = HomeInputs(state.sections, home, state.bestRun, card, live, state.rows, unread, firstRaids = null)
    return HomeView.Blocks(head, composer.blocks(inputs), home?.dataBuild)
}

private val HOLD = 4.seconds
