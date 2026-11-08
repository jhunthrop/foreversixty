package gg.foreversixty.app.home.flows

import androidx.compose.foundation.layout.Arrangement
import androidx.compose.foundation.layout.Box
import androidx.compose.foundation.layout.Column
import androidx.compose.foundation.layout.Row
import androidx.compose.foundation.layout.fillMaxSize
import androidx.compose.foundation.layout.fillMaxWidth
import androidx.compose.foundation.layout.padding
import androidx.compose.foundation.layout.width
import androidx.compose.material3.Text
import androidx.compose.runtime.Composable
import androidx.compose.runtime.LaunchedEffect
import androidx.compose.runtime.getValue
import androidx.compose.runtime.mutableIntStateOf
import androidx.compose.runtime.mutableStateOf
import androidx.compose.runtime.remember
import androidx.compose.runtime.setValue
import androidx.compose.ui.Alignment
import androidx.compose.ui.Modifier
import androidx.compose.ui.platform.LocalResources
import androidx.compose.ui.res.stringResource
import androidx.compose.ui.text.font.FontWeight
import androidx.compose.ui.text.style.TextAlign
import androidx.compose.ui.unit.dp
import androidx.compose.ui.unit.sp
import gg.foreversixty.app.R
import gg.foreversixty.app.account.NoticeToast
import gg.foreversixty.app.designsystem.FsColor
import gg.foreversixty.app.designsystem.FsSpace
import gg.foreversixty.app.designsystem.FsType
import gg.foreversixty.app.designsystem.ParityRegion
import gg.foreversixty.app.designsystem.classColor
import gg.foreversixty.app.flows.FlowBlock
import gg.foreversixty.app.flows.FlowButton
import gg.foreversixty.app.flows.FlowCountdown
import gg.foreversixty.app.flows.FlowFeedRow
import gg.foreversixty.app.flows.FlowGap
import gg.foreversixty.app.flows.FlowHead
import gg.foreversixty.app.flows.FlowLead
import gg.foreversixty.app.flows.FlowLeadImage
import gg.foreversixty.app.flows.FlowLink
import gg.foreversixty.app.flows.FlowNote
import gg.foreversixty.app.flows.FlowPage
import gg.foreversixty.app.flows.FlowRow
import gg.foreversixty.app.flows.FlowSection
import gg.foreversixty.app.flows.FlowSkeleton
import gg.foreversixty.app.model.FeedRow
import gg.foreversixty.app.model.KeyDate
import java.time.Instant

/** Signed-out Home's taps. */
data class HomeOutActions(
    val onSignIn: () -> Unit = {},
    val onClass: (String) -> Unit = {},
    val onFeedRow: (String) -> Unit = {},
    val onRetryFeed: () -> Unit = {},
    val onHowToStart: () -> Unit = {},
)

/** Where signed-out Home's sign-in stands: the browser open, and F1.4's notice. */
data class HomeOutSignIn(val signingIn: Boolean = false, val notice: String? = null, val noticeIsFailure: Boolean = true)

/**
 * Signed-out Home (R3, flows F8 step 1): the head with Sign in, Pick your class, the first raids'
 * countdown and `Forever right now` from [service]. [topBar] is the profile button's bar.
 */
@Composable
fun HomeOutRoute(service: HomeOutService, signIn: HomeOutSignIn, actions: HomeOutActions, topBar: @Composable () -> Unit) {
    val res = LocalResources.current
    var dates by remember { mutableStateOf<List<KeyDate>>(emptyList()) }
    var build by remember { mutableStateOf<String?>(null) }
    var rows by remember { mutableStateOf<List<FeedRow>?>(null) }
    var loading by remember { mutableStateOf(true) }
    var attempt by remember { mutableIntStateOf(0) }
    LaunchedEffect(Unit) {
        dates = service.keyDates()
        build = service.iconBuild()
    }
    LaunchedEffect(attempt) {
        loading = true
        rows = service.feed()
        loading = false
    }
    val view = HomeOutComposer(res, Instant.now()).view(service.specs, dates, rows, loading, build)
    Box(Modifier.fillMaxSize()) {
        Column(Modifier.fillMaxSize()) {
            topBar()
            HomeOutScreen(view, signIn.signingIn, actions.copy(onRetryFeed = { attempt++ }))
        }
        signIn.notice?.let { NoticeToast(it, Modifier.padding(horizontal = FsSpace.space4, vertical = FsSpace.space3), isFailure = signIn.noticeIsFailure) }
    }
}

@Composable
fun HomeOutScreen(view: HomeOutView, signingIn: Boolean, actions: HomeOutActions) {
    FlowPage(top = 0.dp) {
        FlowHead(stringResource(R.string.shared_home_signedOutTitle), view.facts, actions = {
            FlowButton(stringResource(R.string.shared_home_signIn), actions.onSignIn, primary = true, small = false, busy = signingIn)
        })
        FlowSection(stringResource(R.string.shared_home_pickClass), top = FsSpace.space5) {
            FlowGap(FsSpace.space3)
            view.classes.forEach { pick ->
                FlowRow(pick.name, pick.specs, FlowLead.Crest(pick.slug, CLASS_CREST), titleColor = classColor(pick.slug), inset = FsSpace.space4, onClick = { actions.onClass(pick.slug) })
            }
        }
        view.raids?.let { FirstRaids(it, view.build) }
        FlowSection(stringResource(R.string.shared_feed_title)) {
            FlowGap(FsSpace.space3)
            Feed(view.feed, actions)
        }
        FlowLink(stringResource(R.string.shared_home_howToStart) + LINK_MARK, Modifier.padding(top = FsSpace.space6), actions.onHowToStart)
    }
}

@Composable
private fun FirstRaids(card: FirstRaidsCard, build: String?) {
    FlowBlock(ParityRegion.COUNTDOWNS, top = 0.dp) {
        FlowGap(FsSpace.space7)
        FlowCountdown(card.eyebrow, card.clock, card.clockLabel) {
            Row(Modifier.fillMaxWidth(), horizontalArrangement = Arrangement.spacedBy(RAID_GAP, Alignment.End)) {
                card.raids.forEach { raid ->
                    Column(Modifier.width(RAID_WIDTH), horizontalAlignment = Alignment.CenterHorizontally) {
                        FlowLeadImage(FlowLead.Icon(raid.icon, build))
                        Text(raid.name, style = RAID_NAME, color = FsColor.nav, textAlign = TextAlign.Center, modifier = Modifier.padding(top = RAID_NAME_GAP))
                    }
                }
            }
        }
    }
}

@Composable
private fun Feed(feed: HomeFeed, actions: HomeOutActions) {
    when (feed) {
        HomeFeed.Loading -> repeat(FEED_SKELETONS) { index ->
            if (index > 0) FlowGap(FsSpace.space3)
            FlowSkeleton(SKELETON_WIDTH, SKELETON_HEIGHT)
        }
        HomeFeed.Failed -> FlowRow(stringResource(R.string.shared_feed_failed), last = true, trail = {
            FlowButton(stringResource(R.string.common_tryAgain), actions.onRetryFeed)
        })
        HomeFeed.Empty -> FlowNote(stringResource(R.string.shared_feed_logsEmpty))
        is HomeFeed.Rows -> feed.lines.forEach { line ->
            FlowFeedRow(line.lead, line.sentence, line.time, line.kind) { actions.onFeedRow(line.href) }
        }
    }
}

private const val LINK_MARK = " ›"
private const val FEED_SKELETONS = 5
private val SKELETON_WIDTH = 361.dp
private val SKELETON_HEIGHT = 56.dp
private val CLASS_CREST = 44.dp
private val RAID_GAP = 18.dp
private val RAID_WIDTH = 88.dp
private val RAID_NAME_GAP = 6.dp
private val RAID_NAME = FsType.small.copy(fontSize = 13.sp, lineHeight = 16.sp, fontWeight = FontWeight.SemiBold)
