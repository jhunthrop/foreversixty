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
 * sections, the raid (answered in place through the Raid tab's controller) and New for you.
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
) {
    val res = LocalResources.current
    val context = LocalContext.current
    val scope = rememberCoroutineScope()
    val haptics = rememberHaptics()
    val tabModel = remember(raids) { raids?.let { RaidsTabModel(it.service, it.cache, res, it.now, connectivity = it.connectivity) } }
    val state = remember(character.key) { HomeFlowState(service, tabModel, activity, Instant::now) }
    var controller by remember { mutableStateOf<SignupController?>(null) }
    val home = (phase as? HomePhase.Loaded)?.home ?: (phase as? HomePhase.Failed)?.cached

    LaunchedEffect(character.key) {
        state.load(character)
        raids?.refreshBadge(state.raidContent?.head?.id)
    }
    LaunchedEffect(home) { state.refreshBestRun(home) }

    // One controller per raid, the Raid tab's: the answer saves, its haptics and its offline rule.
    val raid = state.nextRaid
    val guild = state.raidContent?.head
    if (raids != null && raid != null && guild != null && controller?.raid?.id != raid.id) {
        controller = SignupController(
            guild.id,
            raid,
            SignupController.Dependencies(raids.service, res, haptics, SystemNotificationPermission(context), raids.askStore, raids.now, scope, HOLD),
        )
    }
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
                state.load(character)
            }
        },
    )
    // E5: offline, the last answer shows under the strip with its time; stale, with Try again.
    (phase as? HomePhase.Loaded)?.offlineSince?.let { since ->
        SavedAnswerStrip(OfflineStamp.format(since, now, ZoneId.systemDefault()), phase.stale, actions.onRetry)
    }
    HomeFlowScreen(view, actions)
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
    val next = state.nextRaid
    val card = next?.let { raid ->
        val own = raid.signups.firstOrNull { it.own && it.characterKey == character.key }
        val answer = controller?.takeIf { it.raid.id == raid.id }?.block(character.key)?.shown ?: words.ownAnswer(raid, character.key)
        words.card(raid, guild, character.key, answer, lateOpen = own?.lateMinutes == null, badge = badge)
    }
    val live = state.liveRaid?.let { words.live(it, guild, character.key, words.ownAnswer(it, character.key)) }
    val inputs = HomeInputs(state.sections, home, state.bestRun, card, live, state.rows, unread, firstRaids = null)
    return HomeView.Blocks(head, composer.blocks(inputs), home?.dataBuild)
}

private val HOLD = 4.seconds
