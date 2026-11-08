package gg.foreversixty.app.rankings

import android.content.Context
import android.content.res.Resources
import androidx.activity.compose.BackHandler
import androidx.compose.foundation.background
import androidx.compose.foundation.layout.Box
import androidx.compose.foundation.layout.Column
import androidx.compose.foundation.layout.fillMaxSize
import androidx.compose.foundation.layout.fillMaxWidth
import androidx.compose.foundation.layout.padding
import androidx.compose.material.icons.Icons
import androidx.compose.material.icons.filled.Star
import androidx.compose.material3.ExperimentalMaterial3Api
import androidx.compose.material3.Icon
import androidx.compose.material3.IconButton
import androidx.compose.material3.pulltorefresh.PullToRefreshBox
import androidx.compose.runtime.Composable
import androidx.compose.runtime.LaunchedEffect
import androidx.compose.runtime.getValue
import androidx.compose.runtime.mutableStateOf
import androidx.compose.runtime.produceState
import androidx.compose.runtime.remember
import androidx.compose.runtime.rememberCoroutineScope
import androidx.compose.runtime.setValue
import androidx.compose.ui.Alignment
import androidx.compose.ui.Modifier
import androidx.compose.ui.draw.alpha
import androidx.compose.ui.platform.LocalContext
import androidx.compose.ui.platform.LocalResources
import androidx.compose.ui.res.stringResource
import androidx.compose.ui.unit.dp
import gg.foreversixty.app.R
import gg.foreversixty.app.account.AccountLinks
import gg.foreversixty.app.bis.BisLabels
import gg.foreversixty.app.bis.BisRange
import gg.foreversixty.app.common.rememberHaptics
import gg.foreversixty.app.designsystem.FsColor
import gg.foreversixty.app.flows.FlowBackBar
import gg.foreversixty.app.items.ItemTooltipRequest
import gg.foreversixty.app.items.ItemTooltipSheet
import gg.foreversixty.app.model.CharacterProfile
import gg.foreversixty.app.paperdoll.PaperdollSlot
import gg.foreversixty.app.raids.logic.GameClass
import gg.foreversixty.app.rankings.logic.CharacterCard
import gg.foreversixty.app.rankings.logic.CharacterLink
import gg.foreversixty.app.rankings.logic.CharacterPhase
import gg.foreversixty.app.rankings.logic.CharacterRoute
import gg.foreversixty.app.rankings.logic.CharacterStore
import gg.foreversixty.app.rankings.logic.RankingsViewer
import gg.foreversixty.app.rankings.logic.WatchState
import gg.foreversixty.app.rankings.ui.CharacterActions
import gg.foreversixty.app.rankings.ui.CharacterContent
import gg.foreversixty.app.rankings.ui.CharacterScreenState
import gg.foreversixty.app.rankings.ui.StarOutline
import gg.foreversixty.app.system.logic.ShareText
import kotlinx.coroutines.launch

/** F11.4: one character under its back bar (Watch and Share), its gear's tooltips and its best fights' reports. */
@OptIn(ExperimentalMaterial3Api::class)
@Composable
fun CharacterRouteScreen(
    feature: RankingsFeature,
    route: CharacterRoute,
    viewer: RankingsViewer,
    reports: RankingsReports?,
    openReport: (String) -> Unit,
    back: () -> Unit,
) {
    val context = LocalContext.current
    val res = LocalResources.current
    val scope = rememberCoroutineScope()
    val store = remember(route, viewer) { CharacterStore(route, feature.service, feature.cache, viewer, feature.connectivity, feature.extras, feature.now) }
    val haptics = rememberHaptics()
    var refreshing by remember { mutableStateOf(false) }
    var tooltip by remember { mutableStateOf<ItemTooltipRequest?>(null) }
    LaunchedEffect(store) { store.load() }
    BackHandler(onBack = back)
    val profile = (store.phase as? CharacterPhase.Loaded)?.profile
    val build by produceState<String?>(null, feature) { value = feature.iconBuild() }
    val classSlug = profile?.let(CharacterCard::classSlug)
    val items by produceState(emptyMap(), build, classSlug) {
        val b = build
        if (b != null && classSlug != null) value = feature.catalog.items(b, classSlug).orEmpty()
    }
    val actions = CharacterActions(
        back = back,
        retry = { scope.launch { store.load() } },
        toggleWatch = { scope.launch { store.toggleWatch(haptics) } },
        share = { share(context, res, feature, route, profile) },
        open = { link -> open(context, link, route, profile, feature) },
        openSocket = { slot -> profile?.let { p -> build?.let { b -> tooltip = tooltipRequest(slot, p, b, feature, res) } } },
        openFight = { fight -> openReport(fight.reportId) },
        openActivity = { row -> AccountLinks.open(context, CharacterCard.webUrl(row.href)) },
        lookUp = { name -> feature.searchAgain?.invoke(name) ?: back() },
        signIn = { AccountLinks.open(context, CharacterCard.webUrl(CharacterCard.LOGIN_PATH)) },
        getAddon = { AccountLinks.open(context, AccountLinks.SETUP) },
    )
    val state = CharacterScreenState(route, store.phase, store.watch, viewer, items, build, reports != null, store.extras)
    Column(Modifier.fillMaxSize().background(FsColor.bg)) {
        Box(Modifier.fillMaxWidth()) {
            FlowBackBar(stringResource(R.string.rank_title), back, if (profile != null) actions.share else null)
            val watch = store.watch
            if (profile != null && watch != null) {
                val loaded = store.phase as? CharacterPhase.Loaded
                WatchStar(watch, profile.character.name, enabled = loaded?.offlineSince == null || loaded.stale, actions.toggleWatch, Modifier.align(Alignment.CenterEnd))
            }
        }
        PullToRefreshBox(
            isRefreshing = refreshing,
            onRefresh = { scope.launch { refreshing = true; store.load(); refreshing = false } },
            modifier = Modifier.fillMaxSize(),
        ) {
            CharacterContent(state, feature.specs, feature.now(), actions, feature.zone)
        }
    }
    tooltip?.let { request -> ItemTooltipSheet(request, feature.catalog, onDismiss = { tooltip = null }) }
}

/** The Watch star (spec §7), left of Share: outline muted, filled gold; 50% and disabled offline. */
@Composable
private fun WatchStar(watch: WatchState, name: String, enabled: Boolean, toggle: () -> Unit, modifier: Modifier) {
    val label = stringResource(if (watch.watching) R.string.char_unwatchA11y else R.string.char_watchA11y, name)
    IconButton(onClick = toggle, enabled = enabled, modifier = modifier.padding(end = SHARE_WIDTH).alpha(if (enabled) 1f else DISABLED_ALPHA)) {
        Icon(if (watch.watching) Icons.Filled.Star else StarOutline, contentDescription = label, tint = if (watch.watching) FsColor.gold else FsColor.muted)
    }
}

private fun share(context: Context, res: Resources, feature: RankingsFeature, route: CharacterRoute, profile: CharacterProfile?) {
    val message = ShareText.character(
        // The card's own key names the page, so the link and the sentence are one character.
        res, profile?.character?.key?.let(CharacterCard::shareUrl) ?: CharacterCard.shareUrl(route), profile?.character?.name ?: route.name,
        profile?.let { CharacterCard.specName(it, feature.specs) },
        profile?.character?.className?.let { slug -> GameClass.of(slug)?.let { res.getString(it.label) } },
        profile?.character?.level,
    )
    ShareText.send(context, message)
}

/** A page link: the planner in the app when it is on, everything else on the web. */
private fun open(context: Context, link: CharacterLink, route: CharacterRoute, profile: CharacterProfile?, feature: RankingsFeature) {
    val planner = feature.openPlanner
    val plannerLink = profile?.let(CharacterCard::plannerLink)
    if (link == CharacterLink.Planner && planner != null && plannerLink != null) {
        planner(plannerLink)
        return
    }
    AccountLinks.open(context, CharacterCard.linkUrl(link, route))
}

/** The F2.6 sheet for a socket: the worn item in the character's class items, spec and level band. */
private fun tooltipRequest(slot: PaperdollSlot, profile: CharacterProfile, build: String, feature: RankingsFeature, res: Resources): ItemTooltipRequest? {
    val item = slot.item ?: return null
    val range = profile.character.level?.let(BisRange::forLevel)
    return ItemTooltipRequest(
        itemId = item.id, fallback = item, build = build, classSlug = CharacterCard.classSlug(profile).orEmpty(),
        spec = profile.gearSpec, specName = CharacterCard.specName(profile, feature.specs), band = range?.band,
        bandLabel = range?.let { BisLabels.head(res, it) }, preset = range?.preset, faction = null, wornId = item.id,
    )
}

private const val DISABLED_ALPHA = 0.5f

/** Share's 48 target at the bar's end, which the star sits beside. */
private val SHARE_WIDTH = 56.dp
