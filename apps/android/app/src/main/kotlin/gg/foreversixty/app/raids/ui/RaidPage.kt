package gg.foreversixty.app.raids.ui

import android.content.Intent
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
import gg.foreversixty.app.designsystem.FsColor
import gg.foreversixty.app.designsystem.FsRadius
import gg.foreversixty.app.model.RaidView
import gg.foreversixty.app.raids.logic.ClassStrip
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
        )
        RaidPageModel.State.Failed -> PageMessage(back) {
            ErrorRegion(stringResource(R.string.raid_failed)) { scope.launch { model.load() } }
        }
        RaidPageModel.State.NotAvailable -> PageMessage(back) {
            Column(verticalArrangement = Arrangement.spacedBy(12.dp), modifier = Modifier.quietCard().padding(16.dp)) {
                Text(stringResource(R.string.raid_notAvailable), style = RaidType.body, color = FsColor.text)
                SecondaryButton(stringResource(R.string.raid_back), back)
            }
        }
    }
}

@OptIn(ExperimentalMaterial3Api::class)
@Composable
private fun PageMessage(back: () -> Unit, content: @Composable () -> Unit) {
    Column(Modifier.fillMaxSize().background(FsColor.bg)) {
        TopAppBar(
            title = {},
            navigationIcon = { BackButton(back) },
            colors = TopAppBarDefaults.topAppBarColors(containerColor = FsColor.bg),
            // The app scaffold already pads for the status bar.
            windowInsets = WindowInsets(0, 0, 0, 0),
        )
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
    Box(Modifier.fillMaxSize().background(FsColor.bg)) {
        Column(Modifier.fillMaxSize()) {
            TopAppBar(
                title = { Text(raid.title.uppercase(), style = RaidType.sec, color = FsColor.strong) },
                navigationIcon = { BackButton(back) },
                actions = {
                    IconButton(onClick = addToCalendar) {
                        Icon(Icons.Filled.DateRange, contentDescription = stringResource(R.string.raid_menu_calendar), tint = FsColor.strong)
                    }
                    Box {
                        IconButton(onClick = { menu = true }) { Icon(Icons.Filled.MoreVert, contentDescription = null, tint = FsColor.strong) }
                        DropdownMenu(expanded = menu, onDismissRequest = { menu = false }) {
                            DropdownMenuItem(text = { Text(stringResource(R.string.raid_menu_calendar)) }, onClick = { menu = false; addToCalendar() })
                            DropdownMenuItem(text = { Text(stringResource(R.string.raid_menu_share)) }, onClick = {
                                menu = false
                                context.startActivity(Intent.createChooser(Intent(Intent.ACTION_SEND).setType("text/plain").putExtra(Intent.EXTRA_TEXT, share), null))
                            })
                            DropdownMenuItem(text = { Text(stringResource(R.string.raid_menu_web)) }, onClick = { menu = false; uri.openUri(web) })
                        }
                    }
                },
                colors = TopAppBarDefaults.topAppBarColors(containerColor = FsColor.bg),
                // The app scaffold already pads for the status bar.
                windowInsets = WindowInsets(0, 0, 0, 0),
            )
            Column(Modifier.fillMaxSize().verticalScroll(rememberScrollState())) {
                offlineSince?.let { OfflineStrip(stringResource(R.string.offline_strip, clock.stripTime(it, now()))) }
                Column(verticalArrangement = Arrangement.spacedBy(22.dp), modifier = Modifier.padding(16.dp)) {
                    RaidHead(raid, facts, clock, now, guildName)
                    CountBlock(facts, muted = facts.chip == RaidChip.Cancelled)
                    if (facts.showsOfficerTools && facts.isSignable) OfficerRow(controller, clock)
                    SignupCard(controller, facts, clock, currentKey, requestPermission)
                    if (raid.viewer.raidNight == true) NightRow { openNight(raid.id) }
                    raid.match?.reportId?.takeIf { facts.chip == RaidChip.Done }?.let { report ->
                        GoldLink(stringResource(R.string.raid_report), { uri.openUri(RaidLinks.report(report)) })
                    }
                    ClassStripView(ClassStrip(raid))
                    Notes(raid)
                    Roster(controller, facts)
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

/** The page head (spec F4.2 items 1 to 3). */
@Composable
fun RaidHead(raid: RaidView, facts: RaidFacts, clock: RaidClock, now: () -> Instant, guildName: String?) {
    Column(verticalArrangement = Arrangement.spacedBy(6.dp)) {
        Row(horizontalArrangement = Arrangement.spacedBy(8.dp), verticalAlignment = Alignment.CenterVertically) {
            Text(stringResource(R.string.raid_eyebrow).uppercase(), style = RaidType.micro, color = FsColor.gold)
            ChipTag(facts.chip)
            guildName?.takeIf { it.isNotEmpty() }?.let { Text(it, style = RaidType.meta, color = FsColor.muted) }
        }
        if (raid.title != raid.zoneName) {
            Text(stringResource(R.string.raids_zoneSuffix, raid.zoneName).trim(), style = RaidType.body, color = FsColor.muted)
        }
        Text(clock.longDate(facts.startsAt), style = RaidType.body, color = FsColor.text)
        TimesLine(facts, clock, showsEnd = true)
        CountdownLine(facts, clock, now)
        val guildZone = runCatching { ZoneId.of(raid.guildTimezone) }.getOrNull()
        if (guildZone != null && clock.offsetDiffers(facts.startsAt, guildZone)) {
            Text(stringResource(R.string.raid_guildTime, clock.guildClock(facts.startsAt, guildZone)), style = RaidType.meta, color = FsColor.muted)
        }
    }
}

/** `Invites 7:45 pm · Pull 8:00 pm · Ends 11:00 pm`, labels muted and times in mono. */
@Composable
fun TimesLine(facts: RaidFacts, clock: RaidClock, showsEnd: Boolean) {
    val parts = buildList {
        facts.inviteAt?.let { add(R.string.raid_invites to clock.clock(it)) }
        add(R.string.raid_pull to clock.clock(facts.startsAt))
        if (showsEnd || facts.inviteAt == null) add(R.string.raid_ends to clock.clock(facts.endsAt))
    }
    val templates = parts.map { (label, time) -> stringResource(label, MARK) to time }
    Text(
        buildAnnotatedString {
            templates.forEachIndexed { index, (template, time) ->
                if (index > 0) withStyle(androidx.compose.ui.text.SpanStyle(color = FsColor.muted)) { append(RaidFacts.SEPARATOR) }
                val (head, tail) = template.split(MARK, limit = 2).let { it[0] to it.getOrElse(1) { "" } }
                withStyle(androidx.compose.ui.text.SpanStyle(color = FsColor.muted)) { append(head) }
                withStyle(RaidType.mono.toSpanStyle().copy(color = FsColor.strong)) { append(time) }
                withStyle(androidx.compose.ui.text.SpanStyle(color = FsColor.muted)) { append(tail) }
            }
        },
        style = RaidType.body,
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
            withStyle(RaidType.mono.toSpanStyle().copy(color = if (running) FsColor.gold else FsColor.muted)) { append(countdown) }
            append(tail)
        },
        style = RaidType.body,
        color = FsColor.text,
        modifier = Modifier.clearAndSetSemantics { contentDescription = a11y },
    )
}

/** The sign-up card (spec F4.2 item 5, F4.9 cancelled). */
@Composable
fun SignupCard(controller: SignupController, facts: RaidFacts, clock: RaidClock, currentKey: String?, requestPermission: () -> Unit) {
    Column(verticalArrangement = Arrangement.spacedBy(14.dp), modifier = Modifier.primaryCard()) {
        Text(stringResource(R.string.raid_yourSignup).uppercase(), style = RaidType.sec, color = FsColor.strong, modifier = Modifier.semantics { heading() })
        when (facts.chip) {
            RaidChip.Cancelled -> {
                Text(stringResource(R.string.raid_cancelled), style = RaidType.body, color = FsColor.text)
                facts.viewerCharacters(currentKey).firstNotNullOfOrNull { c -> SignupState.of(c.state)?.let { it to c.role } }?.let { (state, role) ->
                    Text(
                        stringResource(R.string.raid_cancelledYou, stringResource(state.label), stringResource(gg.foreversixty.app.raids.logic.roleLabel(role))),
                        style = RaidType.meta,
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
    controller.specAsk?.let { ask ->
        controller.character(ask.characterKey)?.let { character ->
            val scope = androidx.compose.runtime.rememberCoroutineScope()
            SpecSheet(
                ask, character, controller.raid.title, clock.shortDate(facts.startsAt),
                pick = { option -> scope.launch { controller.pickSpec(option) } },
                close = { controller.closeSpecSheet() },
            )
        }
    }
}

@Composable
private fun NightRow(open: () -> Unit) {
    Row(
        verticalAlignment = Alignment.CenterVertically,
        modifier = Modifier.quietCard().clickable(onClick = open).heightIn(min = 56.dp).padding(horizontal = 16.dp),
    ) {
        Text(stringResource(R.string.night_title), style = RaidType.bodyStrong, color = FsColor.strong, modifier = Modifier.weight(1f))
        Icon(Icons.AutoMirrored.Filled.KeyboardArrowRight, contentDescription = null, tint = FsColor.muted)
    }
}

/** The classes strip (spec F4.2 item 6): nine crests with In plus Late under each. */
@Composable
fun ClassStripView(strip: ClassStrip) {
    val res = androidx.compose.ui.platform.LocalResources.current
    val label = strip.accessibilityLabel(res)
    val large = androidx.compose.ui.platform.LocalDensity.current.fontScale >= LARGE_FONT_SCALE
    Column(verticalArrangement = Arrangement.spacedBy(10.dp)) {
        SectionLabel(stringResource(R.string.raid_classes))
        FlowRow(
            horizontalArrangement = Arrangement.SpaceBetween,
            verticalArrangement = Arrangement.spacedBy(10.dp),
            maxItemsInEachRow = if (large) LARGE_ROW else GameClassCount,
            modifier = Modifier.fillMaxWidth().clearAndSetSemantics { contentDescription = label },
        ) {
            strip.entries.forEach { (cls, count) ->
                Column(horizontalAlignment = Alignment.CenterHorizontally, modifier = Modifier.alpha(if (count == 0) 0.4f else 1f).width(if (large) 64.dp else 36.dp)) {
                    Crest(cls.slug, if (large) 36.dp else 30.dp)
                    Text("$count", style = RaidType.monoMeta, color = FsColor.strong)
                }
            }
        }
    }
}

@Composable
private fun Notes(raid: RaidView) {
    Column(verticalArrangement = Arrangement.spacedBy(10.dp)) {
        Text(stringResource(R.string.raid_notes).uppercase(), style = RaidType.h2, color = FsColor.strong, modifier = Modifier.semantics { heading() })
        if (raid.notesHtml.isBlank()) {
            Text(stringResource(R.string.raid_notesEmpty), style = RaidType.body, color = FsColor.muted)
        } else {
            Text(NotesText.annotated(raid.notesHtml, FsColor.gold), style = RaidType.body, color = FsColor.text)
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
            .background(RaidColors.sheet, RoundedCornerShape(FsRadius.panel))
            .border(1.dp, FsColor.line, RoundedCornerShape(FsRadius.panel))
            .padding(16.dp)
            .semantics { liveRegion = androidx.compose.ui.semantics.LiveRegionMode.Polite },
    ) {
        Icon(Icons.Filled.Lock, contentDescription = null, tint = FsColor.kill)
        Text(text, style = RaidType.body, color = FsColor.strong)
    }
}

private const val MARK = "\u0000"
private const val MINUTE_MILLIS = 60_000L
private const val LARGE_FONT_SCALE = 1.5f
private const val LARGE_ROW = 5
private const val GameClassCount = 9
