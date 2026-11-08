package gg.foreversixty.app.activity

import androidx.compose.foundation.Image
import androidx.compose.foundation.background
import androidx.compose.foundation.border
import androidx.compose.foundation.clickable
import androidx.compose.foundation.horizontalScroll
import androidx.compose.foundation.layout.Arrangement
import androidx.compose.foundation.layout.Box
import androidx.compose.foundation.layout.Column
import androidx.compose.foundation.layout.Row
import androidx.compose.foundation.layout.Spacer
import androidx.compose.foundation.layout.fillMaxSize
import androidx.compose.foundation.layout.fillMaxWidth
import androidx.compose.foundation.layout.heightIn
import androidx.compose.foundation.layout.offset
import androidx.compose.foundation.layout.padding
import androidx.compose.foundation.layout.requiredSize
import androidx.compose.foundation.layout.size
import androidx.compose.foundation.layout.width
import androidx.compose.foundation.rememberScrollState
import androidx.compose.foundation.shape.CircleShape
import androidx.compose.foundation.shape.RoundedCornerShape
import androidx.compose.foundation.verticalScroll
import androidx.compose.material3.CircularProgressIndicator
import androidx.compose.material3.HorizontalDivider
import androidx.compose.material3.Text
import androidx.compose.runtime.Composable
import androidx.compose.runtime.LaunchedEffect
import androidx.compose.runtime.rememberCoroutineScope
import androidx.compose.ui.Alignment
import androidx.compose.ui.Modifier
import androidx.compose.ui.draw.clip
import androidx.compose.ui.graphics.ColorFilter
import androidx.compose.ui.platform.LocalResources
import androidx.compose.ui.res.painterResource
import androidx.compose.ui.res.stringResource
import androidx.compose.ui.semantics.Role
import androidx.compose.ui.semantics.contentDescription
import androidx.compose.ui.semantics.heading
import androidx.compose.ui.semantics.selected
import androidx.compose.ui.semantics.semantics
import androidx.compose.ui.unit.dp
import androidx.compose.ui.unit.sp
import androidx.lifecycle.compose.collectAsStateWithLifecycle
import gg.foreversixty.app.R
import gg.foreversixty.app.account.AccountDeps
import gg.foreversixty.app.auth.session.RelativeAge
import gg.foreversixty.app.auth.session.SessionPhase
import gg.foreversixty.app.designsystem.BrandButton
import gg.foreversixty.app.designsystem.BrandButtonKind
import gg.foreversixty.app.designsystem.FsColor
import gg.foreversixty.app.designsystem.FsSpace
import gg.foreversixty.app.designsystem.FsType
import gg.foreversixty.app.designsystem.ParityItemKind
import gg.foreversixty.app.designsystem.ParityRegion
import gg.foreversixty.app.designsystem.ParityScrollRegion
import gg.foreversixty.app.designsystem.parityContent
import gg.foreversixty.app.designsystem.parityItem
import gg.foreversixty.app.designsystem.parityRegion
import gg.foreversixty.app.phrase
import gg.foreversixty.app.items.GameIconImage
import kotlinx.coroutines.launch
import java.time.Instant
import java.time.ZoneId
import java.time.format.DateTimeFormatter
import java.util.Locale

/** The Activity tab bound to the session: signed out it asks to sign in, else the inbox. */
@Composable
fun ActivityTab(feature: ActivityFeature, account: AccountDeps?, onOpen: (ActivityLanding) -> Unit, bar: @Composable () -> Unit) {
    val phase = account?.session?.phase?.collectAsStateWithLifecycle()?.value
    val context = androidx.compose.ui.platform.LocalContext.current
    val scope = rememberCoroutineScope()
    val signedIn = phase is SessionPhase.SignedIn
    LaunchedEffect(signedIn) { if (signedIn) feature.load() }
    Column(Modifier.fillMaxSize().background(FsColor.bg)) {
        bar()
        if (!signedIn) {
            ActivitySignedOut { account?.let { it.session.beginSignIn(it.browser(context)) } }
        } else {
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
            )
        }
    }
}

data class ActivityActions(
    val onFilter: (ActivityFilter) -> Unit = {},
    val onMarkAllRead: () -> Unit = {},
    val onMore: () -> Unit = {},
    val onRetry: () -> Unit = {},
    val onOpen: (gg.foreversixty.app.model.ActivityRow) -> Unit = {},
)

/** Signed out (the activity-out mock): the title, what arrives here, Sign in with Battle.net. */
@Composable
fun ActivitySignedOut(onSignIn: () -> Unit) {
    Column(Modifier.fillMaxWidth().parityContent().padding(horizontal = FsSpace.space4)) {
        Head {
            Title()
            Text(
                stringResource(R.string.activity_out),
                style = FsType.small,
                color = FsColor.muted,
                modifier = Modifier.padding(top = TITLE_GAP).parityRegion(ParityRegion.HEAD_META),
            )
        }
        Spacer(Modifier.heightIn(min = FsSpace.space5))
        BrandButton(stringResource(R.string.signin_button), BrandButtonKind.Primary, onSignIn, height = 44.dp)
    }
}

/** The inbox: the filters, then the rows by day, unread ones dotted, Mark all read on the first day. */
@Composable
fun ActivityScreen(state: ActivityState, filter: ActivityFilter, build: String?, now: Instant, zone: ZoneId, actions: ActivityActions) {
    Column(Modifier.fillMaxSize().verticalScroll(rememberScrollState()).parityContent()) {
        Column(Modifier.padding(horizontal = FsSpace.space4)) { Head { Title() } }
        Filters(filter, actions.onFilter)
        when (state) {
            ActivityState.Loading -> Box(Modifier.fillMaxWidth().padding(FsSpace.space8), Alignment.Center) { CircularProgressIndicator(color = FsColor.gold) }
            ActivityState.Failed -> Column(Modifier.padding(FsSpace.space4)) {
                Text(stringResource(R.string.activity_failed), style = FsType.body, color = FsColor.wipe)
                BrandButton(stringResource(R.string.common_tryAgain), BrandButtonKind.Secondary, actions.onRetry, height = 44.dp, fill = false)
            }
            is ActivityState.Loaded -> Rows(state, build, now, zone, actions)
        }
    }
}

/** The page head: 8 under the bar, the title block 16 below its top, 14 above its foot (the mock's PageHeader). */
@Composable
private fun Head(content: @Composable () -> Unit) {
    Column(
        Modifier
            .padding(top = FsSpace.space2)
            .fillMaxWidth()
            .parityRegion(ParityRegion.HEAD)
            .padding(top = FsSpace.space4, bottom = HEAD_BOTTOM),
    ) {
        Column(Modifier.fillMaxWidth().parityRegion(ParityRegion.HEAD_TITLE)) { content() }
    }
}

@Composable
private fun Title() {
    Text(stringResource(R.string.activity_title), style = FsType.title, color = FsColor.strong)
}

@Composable
private fun Filters(current: ActivityFilter, onFilter: (ActivityFilter) -> Unit) {
    Row(
        Modifier
            .padding(start = FsSpace.space4)
            .parityRegion(ParityRegion.CONTROLS)
            .horizontalScroll(rememberScrollState())
            .padding(end = FsSpace.space4)
            .padding(vertical = CONTROLS_INSET),
        horizontalArrangement = Arrangement.spacedBy(FsSpace.space2),
    ) {
        ParityScrollRegion(ParityRegion.CONTROLS) { ActivityFilter.entries.forEach { option ->
            val on = option == current
            val shape = RoundedCornerShape(6.dp)
            Box(
                Modifier
                    .heightIn(min = FILTER_HEIGHT)
                    .parityItem(ParityItemKind.Button)
                    .background(if (on) FsColor.gold else FsColor.raised, shape)
                    .border(1.dp, if (on) FsColor.gold else FsColor.line, shape)
                    .clickable(role = Role.Tab) { onFilter(option) }
                    .semantics { selected = on }
                    .padding(horizontal = FsSpace.space5),
                contentAlignment = Alignment.Center,
            ) {
                Text(stringResource(option.label), style = FsType.bodyStrong.copy(fontSize = FsType.small.fontSize), color = if (on) FsColor.bg else FsColor.strong)
            }
        } }
    }
}

@Composable
private fun Rows(state: ActivityState.Loaded, build: String?, now: Instant, zone: ZoneId, actions: ActivityActions) {
    val sections = Activity.sections(Activity.items(state.rows), now, zone)
    if (sections.isEmpty()) {
        Text(stringResource(R.string.activity_empty), style = FsType.small, color = FsColor.muted, modifier = Modifier.padding(FsSpace.space4))
        return
    }
    val anyUnread = state.rows.any { !it.read }
    sections.forEachIndexed { index, section ->
        val title = stringResource(section.day.label)
        Column(Modifier.padding(horizontal = FsSpace.space4).padding(top = FsSpace.space7).parityRegion(ParityRegion.section(title))) {
            Row(Modifier.fillMaxWidth().padding(bottom = SECTION_TITLE_GAP), verticalAlignment = Alignment.CenterVertically) {
                Text(title, style = FsType.section, color = FsColor.strong, modifier = Modifier.weight(1f).semantics { heading() })
                if (index == 0 && anyUnread) {
                    Text(
                        stringResource(R.string.activity_markAllRead),
                        style = FsType.bodyStrong.copy(fontSize = FsType.small.fontSize),
                        color = FsColor.strong,
                        modifier = Modifier.clickable(role = Role.Button, onClick = actions.onMarkAllRead),
                    )
                }
            }
            HorizontalDivider(color = FsColor.line)
            section.items.forEach { item -> ActivityRowView(item, section.day, build, now, zone) { actions.onOpen(item.row) } }
        }
    }
    if (state.nextBefore != null) {
        Text(
            stringResource(R.string.activity_more),
            style = FsType.bodyStrong,
            color = FsColor.gold,
            modifier = Modifier.padding(FsSpace.space4).clickable(role = Role.Button, onClick = actions.onMore),
        )
    }
    Spacer(Modifier.heightIn(min = FsSpace.space6))
}

@Composable
private fun ActivityRowView(item: ActivityItem, day: ActivityDay, build: String?, now: Instant, zone: ZoneId, onClick: () -> Unit) {
    val unread = stringResource(R.string.activity_unread)
    Row(
        Modifier
            .fillMaxWidth()
            .parityItem(ParityItemKind.Row)
            .clickable(role = Role.Button, onClick = onClick)
            .padding(vertical = ROW_PAD),
    ) {
        // The unread dot sits in the gutter, left of the icon (the mock).
        Box(Modifier.size(0.dp)) {
            if (!item.row.read) {
                Box(
                    Modifier
                        .offset(x = -DOT_INSET, y = 14.dp)
                        .requiredSize(5.dp)
                        .background(FsColor.gold, CircleShape)
                        .semantics { contentDescription = unread },
                )
            }
        }
        RowIcon(item.row.icon, build)
        Spacer(Modifier.width(FsSpace.space3))
        Column(Modifier.weight(1f)) {
            Text(item.row.title, style = FsType.item, color = FsColor.strong)
            if (item.row.body.isNotBlank()) Text(item.row.body, style = SUB, color = FsColor.muted)
            Text(timeLabel(item.at, day, now, zone), style = WHEN, color = FsColor.muted, modifier = Modifier.padding(top = 2.dp))
        }
    }
    HorizontalDivider(color = FsColor.line)
}

/** The row's game icon (the API's `icon`: the item, the instance, the spec), framed as every game icon is. */
@Composable
private fun RowIcon(icon: String, build: String?) {
    Box(
        Modifier
            .size(ICON)
            .parityItem(ParityItemKind.Image)
            .clip(RoundedCornerShape(4.dp))
            .background(FsColor.raised)
            .border(1.dp, FsColor.line, RoundedCornerShape(4.dp)),
        contentAlignment = Alignment.Center,
    ) {
        GameIconImage(icon, build, ICON)
    }
}

/** Today `8 h ago`, Yesterday `6:02 pm`, This week `Tue`, Earlier `3 Nov`. */
@Composable
private fun timeLabel(at: Instant, day: ActivityDay, now: Instant, zone: ZoneId): String = when (day) {
    ActivityDay.Today -> RelativeAge.of(at, now).phrase(LocalResources.current)
    ActivityDay.Yesterday -> CLOCK.withZone(zone).format(at).lowercase(Locale.UK)
    ActivityDay.ThisWeek -> WEEKDAY.withZone(zone).format(at)
    ActivityDay.Earlier -> DATE.withZone(zone).format(at)
}

private val CLOCK = DateTimeFormatter.ofPattern("h:mm a", Locale.UK)
private val WEEKDAY = DateTimeFormatter.ofPattern("EEE", Locale.UK)
private val DATE = DateTimeFormatter.ofPattern("d MMM", Locale.UK)
private val ICON = 40.dp
private val DOT_INSET = 11.dp
private val FILTER_HEIGHT = 36.dp
private val CONTROLS_INSET = 6.dp
private val TITLE_GAP = 6.dp
private val HEAD_BOTTOM = 14.dp

/** The row's line and its time, as the mock's `.row .sub` (13/18) and `.when` (12/16). */
private val SUB = FsType.small.copy(lineHeight = 18.sp)
private val WHEN = FsType.small.copy(fontSize = 12.sp, lineHeight = 16.sp)
private val ROW_PAD = 10.dp

/** Under a day's title, above its rule (the mock's section head is 34 tall). */
private val SECTION_TITLE_GAP = 10.dp
