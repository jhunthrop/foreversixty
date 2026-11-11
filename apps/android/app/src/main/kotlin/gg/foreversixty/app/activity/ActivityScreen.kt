package gg.foreversixty.app.activity

import androidx.compose.foundation.horizontalScroll
import androidx.compose.foundation.layout.heightIn
import androidx.compose.foundation.layout.wrapContentHeight
import androidx.compose.ui.draw.drawWithContent
import androidx.compose.ui.geometry.Offset
import androidx.compose.ui.graphics.BlendMode
import androidx.compose.ui.graphics.Brush
import androidx.compose.ui.graphics.Color
import androidx.compose.ui.graphics.CompositingStrategy
import androidx.compose.ui.graphics.graphicsLayer
import androidx.compose.foundation.layout.Arrangement
import androidx.compose.foundation.layout.Box
import androidx.compose.foundation.layout.Row
import androidx.compose.foundation.layout.fillMaxWidth
import androidx.compose.foundation.layout.padding
import androidx.compose.foundation.rememberScrollState
import androidx.compose.material3.CircularProgressIndicator
import androidx.compose.material3.minimumInteractiveComponentSize
import androidx.compose.runtime.Composable
import androidx.compose.runtime.LaunchedEffect
import androidx.compose.runtime.rememberCoroutineScope
import androidx.compose.ui.Alignment
import androidx.compose.ui.Modifier
import androidx.compose.ui.platform.LocalContext
import androidx.compose.ui.platform.LocalResources
import androidx.compose.ui.res.stringResource
import androidx.compose.ui.semantics.Role
import androidx.compose.ui.unit.dp
import androidx.lifecycle.compose.collectAsStateWithLifecycle
import gg.foreversixty.app.R
import gg.foreversixty.app.account.AccountDeps
import gg.foreversixty.app.auth.session.SessionPhase
import gg.foreversixty.app.designsystem.BrandButton
import gg.foreversixty.app.designsystem.BrandButtonKind
import gg.foreversixty.app.designsystem.FsColor
import gg.foreversixty.app.designsystem.FsSpace
import gg.foreversixty.app.designsystem.FsTarget
import gg.foreversixty.app.designsystem.PanelFrame
import gg.foreversixty.app.designsystem.ParityItemKind
import gg.foreversixty.app.designsystem.ParityRegion
import gg.foreversixty.app.designsystem.ParityScrollRegion
import gg.foreversixty.app.designsystem.parityItem
import gg.foreversixty.app.designsystem.parityRegion
import gg.foreversixty.app.flows.FLOW_CHIP_SMALL
import gg.foreversixty.app.flows.FlowBandPage
import gg.foreversixty.app.flows.FlowPageState
import gg.foreversixty.app.flows.rememberFlowRefresh
import gg.foreversixty.app.flows.FlowBlock
import gg.foreversixty.app.flows.FlowChip
import gg.foreversixty.app.flows.FlowLink
import gg.foreversixty.app.flows.FlowNote
import gg.foreversixty.app.flows.FlowPanel
import gg.foreversixty.app.flows.FlowSection
import gg.foreversixty.app.model.ActivityRow
import gg.foreversixty.app.phrase
import kotlinx.coroutines.launch
import java.time.Instant
import java.time.ZoneId

/** The Activity tab bound to the session: signed out it asks to sign in, else the inbox. */
@Composable
fun ActivityTab(feature: ActivityFeature, account: AccountDeps?, onOpen: (ActivityLanding) -> Unit, bar: @Composable () -> Unit) {
    val phase = account?.session?.phase?.collectAsStateWithLifecycle()?.value
    val context = LocalContext.current
    val scope = rememberCoroutineScope()
    val signedIn = phase is SessionPhase.SignedIn
    LaunchedEffect(signedIn) { if (signedIn) feature.load() }
    if (!signedIn) {
        ActivitySignedOut(bar) { account?.let { it.session.beginSignIn(it.browser(context)) } }
        return
    }
    val refresh = rememberFlowRefresh { feature.load() }
    ActivityScreen(
        state = feature.state,
        filter = feature.filter,
        build = feature.iconBuild(),
        now = Instant.now(),
        zone = ZoneId.systemDefault(),
        actions = ActivityActions(
            onFilter = { scope.launch { feature.load(it) } },
            onMarkAllRead = { scope.launch { feature.markAllRead() } },
            onMore = { scope.launch { feature.more() } },
            onRetry = { scope.launch { feature.load() } },
            onOpen = { row -> Activity.landing(row)?.let(onOpen) },
        ),
        bar = bar,
        page = FlowPageState(rememberScrollState(), refresh),
    )
}

data class ActivityActions(
    val onFilter: (ActivityFilter) -> Unit = {},
    val onMarkAllRead: () -> Unit = {},
    val onMore: () -> Unit = {},
    val onRetry: () -> Unit = {},
    val onOpen: (ActivityRow) -> Unit = {},
)

/** Signed out (round 5): the band with the bar, then one panel: what arrives here and Sign in with Battle.net. */
@Composable
fun ActivitySignedOut(bar: @Composable () -> Unit, onSignIn: () -> Unit) {
    FlowBandPage(bar) {
        FlowBlock(ParityRegion.section(stringResource(R.string.activity_title)), top = 0.dp) {
            FlowPanel {
                FlowNote(stringResource(R.string.activity_out), Modifier.padding(vertical = FsSpace.space2), color = FsColor.text)
                BrandButton(
                    stringResource(R.string.signin_button),
                    BrandButtonKind.Primary,
                    onSignIn,
                    Modifier.padding(vertical = FsSpace.space3).parityItem(ParityItemKind.Button),
                    height = FsTarget.min,
                )
            }
        }
    }
}

/**
 * The inbox in the round-5 look (mock m-android-activity): the band behind the bar and the group filters,
 * then a panel per day (Today, Yesterday, This week, Earlier), unread rows dotted, Mark all read on the
 * first day.
 */
@Composable
fun ActivityScreen(
    state: ActivityState,
    filter: ActivityFilter,
    build: String?,
    now: Instant,
    zone: ZoneId,
    actions: ActivityActions,
    page: FlowPageState = FlowPageState(rememberScrollState()),
    bar: @Composable () -> Unit,
) {
    // An empty inbox under All has nothing to filter; a filter's own empty page keeps the chips to go back.
    val nothingAtAll = state is ActivityState.Loaded && state.rows.isEmpty() && filter == ActivityFilter.All
    FlowBandPage(bar, head = { if (!nothingAtAll) Filters(filter, actions.onFilter) }, page = page) {
        when (state) {
            ActivityState.Loading -> Box(Modifier.fillMaxWidth().padding(FsSpace.space8), Alignment.Center) { CircularProgressIndicator(color = FsColor.gold) }
            ActivityState.Failed -> FlowPanel {
                FlowNote(stringResource(R.string.activity_failed), Modifier.padding(vertical = FsSpace.space2), color = FsColor.text)
                BrandButton(stringResource(R.string.common_tryAgain), BrandButtonKind.Secondary, actions.onRetry, Modifier.padding(vertical = FsSpace.space3), height = FsTarget.min)
            }
            is ActivityState.Loaded -> Days(state, build, now, zone, actions)
        }
    }
}

/**
 * The groups as chips under the bar, All preselected; they scroll sideways past the edge (the mock's
 * `.chips.scroll`) under a trailing fade while more chips lie past it (round 4 row 9, as iOS).
 */
@Composable
private fun Filters(current: ActivityFilter, onFilter: (ActivityFilter) -> Unit) {
    val scroll = rememberScrollState()
    Row(
        Modifier
            .padding(start = FsSpace.space4)
            .parityRegion(ParityRegion.CONTROLS)
            .trailingFade { scroll.canScrollForward }
            .horizontalScroll(scroll)
            .padding(top = FILTERS_TOP, bottom = FILTERS_BOTTOM, end = FsSpace.space4),
        horizontalArrangement = Arrangement.spacedBy(FsSpace.space2),
    ) {
        ParityScrollRegion(ParityRegion.CONTROLS) {
            ActivityFilter.entries.forEach { option ->
                FlowChip(
                    stringResource(option.label),
                    option == current,
                    { onFilter(option) },
                    Modifier.minimumInteractiveComponentSize(),
                    minHeight = FLOW_CHIP_SMALL,
                    role = Role.Tab,
                )
            }
        }
    }
}

@Composable
private fun Days(state: ActivityState.Loaded, build: String?, now: Instant, zone: ZoneId, actions: ActivityActions) {
    val sections = Activity.sections(Activity.items(state.rows), now, zone)
    if (sections.isEmpty()) {
        FlowBlock(EMPTY_REGION, top = 0.dp) {
            FlowPanel { FlowNote(stringResource(R.string.activity_empty), Modifier.padding(vertical = FsSpace.space2), color = FsColor.text) }
        }
        return
    }
    val anyUnread = state.rows.any { !it.read }
    sections.forEachIndexed { index, section ->
        val markAll = if (index == 0 && anyUnread) stringResource(R.string.activity_markAllRead) else null
        FlowSection(stringResource(section.day.label), action = markAll, onAction = actions.onMarkAllRead, top = 0.dp, panel = PanelFrame.None) {
            section.items.forEachIndexed { at, item ->
                ActivityRowView(
                    row = item.row,
                    line = item.row.body,
                    build = build,
                    last = at == section.items.lastIndex,
                    meta = ActivityTime.label(LocalResources.current, item.at, section.day, now, zone),
                    onClick = { actions.onOpen(item.row) },
                )
            }
        }
    }
    // Round 4 row 7: left under its list, as the web's pager link, at the 48 dp target.
    if (state.nextBefore != null) {
        FlowLink(
            stringResource(R.string.activity_more),
            Modifier.padding(horizontal = FsSpace.space4).heightIn(min = FsTarget.min).wrapContentHeight(Alignment.CenterVertically),
            actions.onMore,
        )
    }
}

/** Fades the trailing [FILTER_FADE] to clear while [active], so a cut chip reads as more to scroll. */
private fun Modifier.trailingFade(active: () -> Boolean): Modifier = this
    .graphicsLayer(compositingStrategy = CompositingStrategy.Offscreen)
    .drawWithContent {
        drawContent()
        if (active()) {
            val start = size.width - FILTER_FADE.toPx()
            drawRect(
                Brush.horizontalGradient(listOf(Color.Black, Color.Transparent), startX = start, endX = size.width),
                topLeft = Offset(start, 0f),
                blendMode = BlendMode.DstIn,
            )
        }
    }

/** How far the chips fade out at the trailing edge (iOS `ActivityFilterChips.fade`). */
private val FILTER_FADE = FsSpace.space6

/** A filter chip's face (the mock's `.chips a`, 7 above and below a 20 line), 48 to the touch. */

/** The chips' faces 8 under the bar and the band's foot 14 under them, less the touch target's 6. */
private val FILTERS_TOP = 2.dp
private val FILTERS_BOTTOM = 8.dp

/** The empty inbox's one line (the mock's `section-empty`). */
private const val EMPTY_REGION = "section-empty"
