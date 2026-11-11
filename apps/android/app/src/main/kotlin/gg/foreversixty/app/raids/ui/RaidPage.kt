package gg.foreversixty.app.raids.ui

import androidx.compose.ui.graphics.Color
import gg.foreversixty.app.designsystem.surfaceBand
import gg.foreversixty.app.designsystem.FsTopBar
import gg.foreversixty.app.designsystem.screenGround
import android.content.Intent
import gg.foreversixty.app.designsystem.GameColors
import gg.foreversixty.app.designsystem.ScaledType
import android.provider.CalendarContract
import androidx.compose.foundation.background
import androidx.compose.foundation.border
import androidx.compose.foundation.clickable
import androidx.compose.foundation.layout.Arrangement
import androidx.compose.foundation.layout.WindowInsets
import androidx.compose.foundation.layout.Box
import androidx.compose.foundation.layout.Column
import androidx.compose.foundation.layout.FlowRow
import androidx.compose.foundation.layout.Row
import androidx.compose.foundation.layout.fillMaxSize
import androidx.compose.foundation.layout.fillMaxWidth
import androidx.compose.foundation.layout.heightIn
import androidx.compose.foundation.layout.padding
import androidx.compose.foundation.layout.width
import androidx.compose.foundation.rememberScrollState
import androidx.compose.foundation.shape.RoundedCornerShape
import androidx.compose.foundation.verticalScroll
import androidx.compose.material.icons.Icons
import androidx.compose.material.icons.automirrored.filled.ArrowBack
import androidx.compose.material.icons.automirrored.filled.KeyboardArrowRight
import androidx.compose.material.icons.filled.DateRange
import androidx.compose.material.icons.filled.Lock
import androidx.compose.material.icons.filled.MoreVert
import androidx.compose.material3.CircularProgressIndicator
import androidx.compose.material3.DropdownMenu
import androidx.compose.material3.DropdownMenuItem
import androidx.compose.material3.ExperimentalMaterial3Api
import androidx.compose.material3.Icon
import androidx.compose.material3.IconButton
import androidx.compose.material3.Text
import androidx.compose.material3.TopAppBar
import androidx.compose.material3.TopAppBarDefaults
import androidx.compose.runtime.Composable
import androidx.compose.runtime.LaunchedEffect
import androidx.compose.runtime.getValue
import androidx.compose.runtime.mutableStateOf
import androidx.compose.runtime.remember
import androidx.compose.runtime.setValue
import androidx.compose.ui.Alignment
import androidx.compose.ui.Modifier
import androidx.compose.ui.draw.alpha
import androidx.compose.ui.platform.LocalContext
import androidx.compose.ui.platform.LocalUriHandler
import androidx.compose.ui.res.stringResource
import androidx.compose.ui.semantics.clearAndSetSemantics
import androidx.compose.ui.semantics.contentDescription
import androidx.compose.ui.semantics.heading
import androidx.compose.ui.semantics.liveRegion
import androidx.compose.ui.semantics.semantics
import androidx.compose.ui.text.buildAnnotatedString
import androidx.compose.ui.text.withStyle
import androidx.compose.ui.unit.dp
import androidx.core.net.toUri
import gg.foreversixty.app.R
import gg.foreversixty.app.common.SavedAnswerStrip
import gg.foreversixty.app.designsystem.FsColor
import gg.foreversixty.app.designsystem.FsType
import gg.foreversixty.app.designsystem.FsRadius
import gg.foreversixty.app.model.RaidView
import gg.foreversixty.app.raids.logic.NotesText
import gg.foreversixty.app.raids.logic.RaidChip
import gg.foreversixty.app.raids.logic.RaidClock
import gg.foreversixty.app.raids.logic.RaidFacts
import gg.foreversixty.app.raids.logic.RaidLinks
import gg.foreversixty.app.raids.logic.RaidPageModel
import gg.foreversixty.app.raids.logic.SignupController
import gg.foreversixty.app.raids.logic.SignupState
import java.time.Duration
import java.time.Instant
import java.time.ZoneId
import kotlinx.coroutines.delay
import kotlinx.coroutines.launch

/** F4.2 the raid page: loads, then draws [RaidPageContent]. */
@Composable
fun RaidPageScreen(
    model: RaidPageModel,
    clock: RaidClock,
    now: () -> Instant,
    currentKey: String?,
    guildName: String,
    guildUrl: String?,
    openNight: (Long) -> Unit,
    back: () -> Unit,
    requestPermission: () -> Unit,
    toCover: Boolean = false,
) {
    val scope = androidx.compose.runtime.rememberCoroutineScope()
    LaunchedEffect(model) { model.load() }
    LaunchedEffect(model) { model.stayLive() }
    when (val state = model.state) {
        RaidPageModel.State.Loading -> Box(Modifier.fillMaxSize(), contentAlignment = Alignment.Center) {
            CircularProgressIndicator(color = FsColor.gold)
        }
        is RaidPageModel.State.Loaded -> RaidPageContent(
            state.controller, clock, now, currentKey, model.offlineSince, guildName, guildUrl, openNight, back, requestPermission,
            stale = model.stale,
            retry = { scope.launch { model.load() } },
            toCover = toCover,
        )
        RaidPageModel.State.Failed -> PageMessage(back) {
            ErrorRegion(stringResource(R.string.raid_failed)) { scope.launch { model.load() } }
        }
        RaidPageModel.State.NotAvailable -> PageMessage(back) {
            Column(verticalArrangement = Arrangement.spacedBy(12.dp), modifier = Modifier.raidQuietCard().padding(16.dp)) {
                Text(stringResource(R.string.raid_notAvailable), style = ScaledType.body, color = FsColor.text)
                RaidSecondaryButton(stringResource(R.string.raid_back), back)
            }
        }
    }
}

@OptIn(ExperimentalMaterial3Api::class)
@Composable
private fun PageMessage(back: () -> Unit, content: @Composable () -> Unit) {
    Column(Modifier.fillMaxSize().screenGround()) {
        FsTopBar(null, stringResource(R.string.tab_raids), back)
        Box(Modifier.padding(16.dp)) { content() }
    }
}

@Composable
private fun BackButton(back: () -> Unit) {
    IconButton(onClick = back) {
        Icon(Icons.AutoMirrored.Filled.ArrowBack, contentDescription = stringResource(R.string.tab_raids), tint = FsColor.strong)
    }
}

/** The raid page's content, pure over the controller so screenshot tests draw every state. */
@OptIn(ExperimentalMaterial3Api::class)
@Composable
fun RaidPageContent(
    controller: SignupController,
    clock: RaidClock,
    now: () -> Instant,
    currentKey: String?,
    offlineSince: Instant?,
    guildName: String,
    guildUrl: String?,
    openNight: (Long) -> Unit,
    back: () -> Unit,
    requestPermission: () -> Unit,
    stale: Boolean = false,
    retry: () -> Unit = {},
    /** Opened from the Raid tab's `and N more`: the composition's cover scrolls into view. */
    toCover: Boolean = false,
) {
    val raid = controller.raid
    val facts = controller.facts
    val context = LocalContext.current
    val uri = LocalUriHandler.current
    val share = RaidLinks.share(controller.guild, raid.id)
    val web = guildUrl?.let { "$it/raids/${raid.id}" } ?: share
    var menu by remember { mutableStateOf(false) }
    val location = if (guildName.isEmpty()) raid.zoneName else stringResource(R.string.raids_calendarLocation, guildName, raid.zoneName)
    val addToCalendar = { context.startActivity(calendarIntent(raid, facts, location, share)) }
    Box(Modifier.fillMaxSize().screenGround()) {
        Column(Modifier.fillMaxSize()) {
            FsTopBar(raid.title, stringResource(R.string.tab_raids), back, bottomLine = false) {
                    IconButton(onClick = addToCalendar) {
                        Icon(Icons.Filled.DateRange, contentDescription = stringResource(R.string.raid_menu_calendar), tint = FsColor.strong)
                    }
                    Box {
                        IconButton(onClick = { menu = true }) { Icon(Icons.Filled.MoreVert, contentDescription = stringResource(R.string.raids_a11y_more), tint = FsColor.strong) }
                        DropdownMenu(expanded = menu, onDismissRequest = { menu = false }) {
                            DropdownMenuItem(text = { Text(stringResource(R.string.raid_menu_calendar)) }, onClick = { menu = false; addToCalendar() })
                            DropdownMenuItem(text = { Text(stringResource(R.string.raid_menu_share)) }, onClick = {
                                menu = false
                                context.startActivity(Intent.createChooser(Intent(Intent.ACTION_SEND).setType("text/plain").putExtra(Intent.EXTRA_TEXT, share), null))
                            })
                            DropdownMenuItem(text = { Text(stringResource(R.string.raid_menu_web)) }, onClick = { menu = false; uri.openUri(web) })
                        }
                    }
            }
            Column(Modifier.fillMaxSize().verticalScroll(rememberScrollState())) {
                Box(Modifier.fillMaxWidth().surfaceBand().padding(start = 16.dp, end = 16.dp, bottom = 16.dp)) {
                    RaidHead(raid, facts, clock, now, guildName)
                }
                offlineSince?.let { SavedAnswerStrip(clock.stripTime(it, now()), stale, retry) }
                Column(verticalArrangement = Arrangement.spacedBy(22.dp), modifier = Modifier.padding(16.dp)) {
                    CountBlock(facts, muted = facts.chip == RaidChip.Cancelled, ink = FsColor.mutedGround)
                    if (facts.showsOfficerTools && facts.isSignable) OfficerRow(controller, clock)
                    SignupCard(controller, facts, clock, currentKey, requestPermission)
                    if (raid.viewer.raidNight == true) NightRow { openNight(raid.id) }
                    raid.match?.reportId?.takeIf { facts.chip == RaidChip.Done }?.let { report ->
                        GoldLink(stringResource(R.string.raid_report), { uri.openUri(RaidLinks.report(report)) })
                    }
                    Notes(raid)
                    Roster(controller, facts, toCover)
                }
            }
        }
        controller.toast?.let { Toast(it, Modifier.align(Alignment.BottomCenter)) }
    }
}

private fun calendarIntent(raid: RaidView, facts: RaidFacts, location: String, link: String): Intent =
    Intent(Intent.ACTION_INSERT, CalendarContract.Events.CONTENT_URI)
        .putExtra(CalendarContract.Events.TITLE, raid.title)
        .putExtra(CalendarContract.EXTRA_EVENT_BEGIN_TIME, facts.startsAt.toEpochMilli())
        .putExtra(CalendarContract.EXTRA_EVENT_END_TIME, facts.endsAt.toEpochMilli())
        .putExtra(CalendarContract.Events.EVENT_LOCATION, location)
        .putExtra(CalendarContract.Events.DESCRIPTION, link)

/** The page head (spec F4.2 items 1 to 3), on the band: secondary ink is `mutedGround`, AA on the painted art. */
@Composable
fun RaidHead(raid: RaidView, facts: RaidFacts, clock: RaidClock, now: () -> Instant, guildName: String?) {
    Column(verticalArrangement = Arrangement.spacedBy(6.dp)) {
        val large = isLargeText()
        Row(horizontalArrangement = Arrangement.spacedBy(8.dp), verticalAlignment = Alignment.CenterVertically) {
            Text(stringResource(R.string.raid_eyebrow), style = FsType.label, color = FsColor.gold, maxLines = 1, softWrap = false)
            ChipTag(facts.chip)
            if (!large) guildName?.takeIf { it.isNotEmpty() }?.let { Text(it, style = ScaledType.meta, color = FsColor.mutedGround, maxLines = 1) }
        }
        // Large text: the guild name drops to its own line.
        if (large) guildName?.takeIf { it.isNotEmpty() }?.let { Text(it, style = ScaledType.meta, color = FsColor.mutedGround) }
        if (raid.title != raid.zoneName) {
            Text(stringResource(R.string.raids_zoneSuffix, raid.zoneName).trim(), style = ScaledType.body, color = FsColor.mutedGround)
        }
        Text(clock.longDate(facts.startsAt), style = ScaledType.body, color = FsColor.strong)
        TimesLine(facts, clock, showsEnd = true, ink = FsColor.mutedGround)
        CountdownLine(facts, clock, now)
        val guildZone = runCatching { ZoneId.of(raid.guildTimezone) }.getOrNull()
        if (guildZone != null && clock.offsetDiffers(facts.startsAt, guildZone)) {
            Text(stringResource(R.string.raid_guildTime, clock.guildClock(facts.startsAt, guildZone)), style = ScaledType.meta, color = FsColor.mutedGround)
        }
    }
}

/** `Invites 7:45 pm · Pull 8:00 pm · Ends 11:00 pm`: labels in [ink] (muted on a panel, mutedGround on the band), times tabular. */
@Composable
fun TimesLine(facts: RaidFacts, clock: RaidClock, showsEnd: Boolean, ink: Color = FsColor.muted) {
    val parts = buildList {
        facts.inviteAt?.let { add(R.string.raid_invites to clock.clock(it)) }
        add(R.string.raid_pull to clock.clock(facts.startsAt))
        if (showsEnd || facts.inviteAt == null) add(R.string.raid_ends to clock.clock(facts.endsAt))
    }
    val templates = parts.map { (label, time) -> stringResource(label, MARK) to time }
    Text(
        buildAnnotatedString {
            templates.forEachIndexed { index, (template, time) ->
                if (index > 0) withStyle(androidx.compose.ui.text.SpanStyle(color = ink)) { append(RaidFacts.SEPARATOR) }
                val (head, tail) = template.split(MARK, limit = 2).let { it[0] to it.getOrElse(1) { "" } }
                withStyle(androidx.compose.ui.text.SpanStyle(color = ink)) { append(head) }
                // Only the digits are mono; `pm` stays in the body face (UX should 2).
                val digits = time.substringBefore(' ')
                withStyle(ScaledType.tabular.toSpanStyle().copy(color = FsColor.strong)) { append(digits) }
                withStyle(androidx.compose.ui.text.SpanStyle(color = FsColor.strong)) { append(time.removePrefix(digits)) }
                withStyle(androidx.compose.ui.text.SpanStyle(color = ink)) { append(tail) }
            }
        },
        style = ScaledType.body,
    )
}

/** `Invites in 2 days 2 h` (or `Starts …`), the countdown in mono gold, ticking each minute. */
@Composable
fun CountdownLine(facts: RaidFacts, clock: RaidClock, now: () -> Instant) {
    var current by remember { mutableStateOf(now()) }
    LaunchedEffect(Unit) {
        while (true) {
            delay(MINUTE_MILLIS)
            current = now()
        }
    }
    val target = facts.countdownTarget
    val countdown = clock.countdown(target, facts.endsAt, current)
    val template = stringResource(if (facts.inviteAt != null) R.string.raid_countdown_invites else R.string.raid_countdown_starts, MARK)
    val (head, tail) = template.split(MARK, limit = 2).let { it[0] to it.getOrElse(1) { "" } }
    val running = facts.isSignable && target > current
    val a11y = stringResource(R.string.raid_a11y_countdown, clock.spanWords(Duration.between(current, target)))
    Text(
        buildAnnotatedString {
            append(head)
            withStyle(ScaledType.bodyStrong.toSpanStyle().copy(color = if (running) FsColor.gold else FsColor.muted)) { append(countdown) }
            append(tail)
        },
        style = ScaledType.body,
        color = FsColor.text,
        modifier = Modifier.clearAndSetSemantics { contentDescription = a11y },
    )
}

/** The sign-up card (spec F4.2 item 5, F4.9 cancelled). */
@Composable
fun SignupCard(controller: SignupController, facts: RaidFacts, clock: RaidClock, currentKey: String?, requestPermission: () -> Unit) {
    Column(verticalArrangement = Arrangement.spacedBy(14.dp), modifier = Modifier.raidPrimaryCard()) {
        Text(stringResource(R.string.raid_yourSignup), style = ScaledType.sec, color = FsColor.strong, modifier = Modifier.semantics { heading() })
        when (facts.chip) {
            RaidChip.Cancelled -> {
                Text(stringResource(R.string.raid_cancelled), style = ScaledType.body, color = FsColor.text)
                facts.viewerCharacters(currentKey).firstNotNullOfOrNull { c -> SignupState.of(c.state)?.let { it to c.role } }?.let { (state, role) ->
                    Text(
                        stringResource(R.string.raid_cancelledYou, stringResource(state.label), stringResource(gg.foreversixty.app.raids.logic.roleLabel(role))),
                        style = ScaledType.meta,
                        color = FsColor.muted,
                    )
                }
            }
            RaidChip.Done -> Unit
            else -> {
                facts.viewerCharacters(currentKey).forEach { SignupBlockView(controller, it) }
                if (controller.showsReminderAsk) {
                    ReminderAsk(
                        turnOn = { controller.dismissReminderAsk(); requestPermission() },
                        notNow = { controller.dismissReminderAsk() },
                    )
                }
            }
        }
    }
    SpecAskSheet(controller, controller.raid.title, clock.shortDate(facts.startsAt))
}

@Composable
private fun NightRow(open: () -> Unit) {
    Row(
        verticalAlignment = Alignment.CenterVertically,
        modifier = Modifier.raidQuietCard().clickable(onClick = open).heightIn(min = 56.dp).padding(horizontal = 16.dp),
    ) {
        Text(stringResource(R.string.night_title), style = ScaledType.bodyStrong, color = FsColor.strong, modifier = Modifier.weight(1f))
        Icon(Icons.AutoMirrored.Filled.KeyboardArrowRight, contentDescription = null, tint = FsColor.muted)
    }
}

@Composable
private fun Notes(raid: RaidView) {
    Column(verticalArrangement = Arrangement.spacedBy(10.dp)) {
        Text(stringResource(R.string.raid_notes), style = ScaledType.h2, color = FsColor.strong, modifier = Modifier.semantics { heading() })
        if (raid.notesHtml.isBlank()) {
            Text(stringResource(R.string.raid_notesEmpty), style = ScaledType.body, color = FsColor.muted)
        } else {
            Text(NotesText.annotated(raid.notesHtml, FsColor.gold), style = ScaledType.body, color = FsColor.text)
        }
    }
}

@Composable
private fun Toast(text: String, modifier: Modifier) {
    Row(
        horizontalArrangement = Arrangement.spacedBy(10.dp),
        verticalAlignment = Alignment.CenterVertically,
        modifier = modifier
            .padding(16.dp)
            .fillMaxWidth()
            .background(GameColors.sheet, RoundedCornerShape(FsRadius.legacyPanel))
            .border(1.dp, FsColor.line, RoundedCornerShape(FsRadius.legacyPanel))
            .padding(16.dp)
            .semantics { liveRegion = androidx.compose.ui.semantics.LiveRegionMode.Polite },
    ) {
        Icon(Icons.Filled.Lock, contentDescription = null, tint = FsColor.kill)
        Text(text, style = ScaledType.body, color = FsColor.strong)
    }
}

private const val MARK = "\u0000"
private const val MINUTE_MILLIS = 60_000L
