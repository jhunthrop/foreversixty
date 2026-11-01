package gg.foreversixty.app.raids.ui

import androidx.compose.foundation.background
import gg.foreversixty.app.designsystem.Tag
import gg.foreversixty.app.designsystem.FsType
import androidx.compose.foundation.border
import androidx.compose.foundation.clickable
import androidx.compose.foundation.layout.Arrangement
import androidx.compose.foundation.layout.Box
import androidx.compose.foundation.layout.Column
import androidx.compose.foundation.layout.Row
import androidx.compose.foundation.layout.Spacer
import androidx.compose.foundation.layout.fillMaxSize
import androidx.compose.foundation.layout.fillMaxWidth
import androidx.compose.foundation.layout.height
import androidx.compose.foundation.layout.heightIn
import androidx.compose.foundation.layout.padding
import androidx.compose.foundation.layout.size
import androidx.compose.foundation.layout.width
import androidx.compose.foundation.rememberScrollState
import androidx.compose.foundation.shape.CircleShape
import androidx.compose.foundation.shape.RoundedCornerShape
import androidx.compose.foundation.verticalScroll
import androidx.compose.material.icons.Icons
import androidx.compose.material.icons.automirrored.filled.KeyboardArrowRight
import androidx.compose.material.icons.filled.DateRange
import androidx.compose.material.icons.filled.Person
import androidx.compose.material3.HorizontalDivider
import androidx.compose.material3.Icon
import androidx.compose.material3.Text
import androidx.compose.runtime.Composable
import androidx.compose.ui.Alignment
import androidx.compose.ui.Modifier
import androidx.compose.ui.draw.clip
import androidx.compose.ui.layout.ContentScale
import androidx.compose.ui.res.stringResource
import androidx.compose.ui.semantics.clearAndSetSemantics
import androidx.compose.ui.semantics.contentDescription
import androidx.compose.ui.semantics.heading
import androidx.compose.ui.semantics.semantics
import androidx.compose.ui.unit.dp
import coil3.compose.AsyncImage
import gg.foreversixty.app.R
import gg.foreversixty.app.common.FactionMark
import gg.foreversixty.app.designsystem.OfflineStrip
import gg.foreversixty.app.designsystem.FsColor
import gg.foreversixty.app.designsystem.FsRadius
import gg.foreversixty.app.model.RaidNightRow
import gg.foreversixty.app.model.RaidView
import gg.foreversixty.app.raids.logic.GuildHead
import gg.foreversixty.app.raids.logic.PlannedRaids
import gg.foreversixty.app.raids.logic.RaidClock
import gg.foreversixty.app.raids.logic.RaidFacts
import gg.foreversixty.app.raids.logic.RaidsTabContent
import gg.foreversixty.app.raids.logic.RaidsTabState
import gg.foreversixty.app.raids.logic.RowChip
import gg.foreversixty.app.raids.logic.SignupController
import gg.foreversixty.app.raids.logic.roleLabel
import java.time.Instant
import kotlinx.coroutines.launch

/** F4.1 the Raids tab's content, pure over its state so screenshot tests draw every state. */
@Composable
fun RaidsTabContentView(
    state: RaidsTabState,
    clock: RaidClock,
    now: () -> Instant,
    nextController: SignupController?,
    openRaid: (RaidView, Boolean) -> Unit,
    openNight: (Long) -> Unit,
    openWeb: (String) -> Unit,
    retry: () -> Unit,
    requestPermission: () -> Unit,
    guildRows: @Composable (RaidsTabContent) -> Unit = {},
) {
    Column(Modifier.fillMaxSize().background(FsColor.bg).verticalScroll(rememberScrollState())) {
        (state as? RaidsTabState.Loaded)?.content?.offlineSince?.let {
            OfflineStrip(stringResource(R.string.offline_strip, clock.stripTime(it, now())))
        }
        Column(verticalArrangement = Arrangement.spacedBy(22.dp), modifier = Modifier.padding(16.dp)) {
            when (state) {
                RaidsTabState.Loading -> Skeleton()
                RaidsTabState.SignedOut -> Message(stringResource(R.string.raids_signedOut_title), stringResource(R.string.raids_signedOut_body), null)
                is RaidsTabState.NoGuild -> Message(
                    stringResource(R.string.raids_noGuild_title),
                    stringResource(R.string.raids_noGuild_body, state.name),
                    stringResource(R.string.raids_noGuild_note),
                )
                is RaidsTabState.Waiting -> {
                    GuildHeadView(state.head)
                    Column(verticalArrangement = Arrangement.spacedBy(12.dp), modifier = Modifier.raidPrimaryCard()) {
                        Row(verticalAlignment = Alignment.CenterVertically) {
                            Box(Modifier.weight(1f)) {
                                WhoLine(state.character.name, state.character.className.orEmpty(), state.character.spec, null, "")
                            }
                            Tag(stringResource(R.string.raids_waitingTag), FsColor.gold)
                        }
                        Text(stringResource(R.string.raids_waiting), style = FsType.body, color = FsColor.text)
                    }
                }
                RaidsTabState.Failed -> ErrorRegion(stringResource(R.string.raids_failed), retry)
                is RaidsTabState.Loaded -> {
                    GuildHeadView(state.content.head)
                    // Lane F16: Roster and Raid history, for a verified member only.
                    guildRows(state.content)
                    Loaded(state.content, clock, now, nextController, openRaid, openNight, openWeb, requestPermission)
                }
            }
        }
    }
}

@Composable
private fun Message(title: String, body: String, note: String?) {
    Column(verticalArrangement = Arrangement.spacedBy(12.dp), modifier = Modifier.padding(top = 24.dp)) {
        Box(
            contentAlignment = Alignment.Center,
            modifier = Modifier.size(44.dp).background(FsColor.gold.copy(alpha = 0.12f), RoundedCornerShape(FsRadius.control)),
        ) { Icon(Icons.Filled.DateRange, contentDescription = null, tint = FsColor.gold) }
        Text(title.uppercase(), style = FsType.h2, color = FsColor.strong, modifier = Modifier.semantics { heading() })
        Text(body, style = FsType.body, color = FsColor.text)
        note?.let { Text(it, style = FsType.meta, color = FsColor.muted) }
    }
}

@Composable
private fun Loaded(
    content: RaidsTabContent,
    clock: RaidClock,
    now: () -> Instant,
    nextController: SignupController?,
    openRaid: (RaidView, Boolean) -> Unit,
    openNight: (Long) -> Unit,
    openWeb: (String) -> Unit,
    requestPermission: () -> Unit,
) {
    val planned = PlannedRaids(content.planned, now())
    val next = planned.next
    if (next != null && nextController != null && nextController.raid.id == next.id) {
        NextRaidCard(nextController, clock, now, content.character.key, requestPermission) { roster -> openRaid(nextController.raid, roster) }
    } else if (next == null) {
        Column(verticalArrangement = Arrangement.spacedBy(6.dp), modifier = Modifier.raidPrimaryCard()) {
            Text(stringResource(R.string.raids_empty), style = FsType.bodyStrong, color = FsColor.strong)
            Text(stringResource(R.string.raids_emptyNote), style = FsType.meta, color = FsColor.muted)
        }
    }
    content.waiting?.let { waiting ->
        Row(
            horizontalArrangement = Arrangement.spacedBy(12.dp),
            verticalAlignment = Alignment.CenterVertically,
            modifier = Modifier.raidQuietCard().clickable { content.head.webUrl?.let(openWeb) }.padding(14.dp),
        ) {
            Box(
                contentAlignment = Alignment.Center,
                modifier = Modifier.size(36.dp).background(FsColor.gold.copy(alpha = 0.12f), RoundedCornerShape(FsRadius.control)),
            ) { Icon(Icons.Filled.Person, contentDescription = null, tint = FsColor.gold) }
            Column {
                Text(androidx.compose.ui.res.pluralStringResource(R.plurals.raids_pending, waiting, waiting), style = FsType.bodyStrong, color = FsColor.strong)
                Text(stringResource(R.string.raids_pendingLine), style = FsType.meta, color = FsColor.muted)
            }
        }
    }
    if (planned.also.isNotEmpty()) {
        Column(verticalArrangement = Arrangement.spacedBy(8.dp)) {
            SectionLabel(stringResource(R.string.raids_planned), trailing = "${planned.also.size}")
            Column(Modifier.raidQuietCard()) {
                planned.also.forEachIndexed { index, raid ->
                    if (index > 0) HorizontalDivider(color = FsColor.lineSoft)
                    PlannedRow(raid, clock, now(), content.character.key) { openRaid(raid, false) }
                }
            }
        }
    }
    if (content.nights.isNotEmpty()) {
        Column(verticalArrangement = Arrangement.spacedBy(8.dp)) {
            SectionLabel(stringResource(R.string.raids_nights))
            Column(Modifier.raidQuietCard()) {
                content.nights.forEachIndexed { index, night ->
                    if (index > 0) HorizontalDivider(color = FsColor.lineSoft)
                    NightRow(night, clock) { night.raidId?.let(openNight) }
                }
            }
        }
    }
}

/** The guild head: crest 40, the name, `{Faction} · {Region} {Ruleset} · {n} on the roster`. */
@Composable
fun GuildHeadView(head: GuildHead) {
    val crest = @Composable {
        // The guild crest; with none, the faction mark (v1 spec 4.3, as the site's GuildCrest falls back).
        if (head.crestUrl != null) {
            Box(Modifier.size(40.dp).clip(CircleShape).background(FsColor.raised).border(1.dp, FsColor.line, CircleShape)) {
                AsyncImage(model = head.crestUrl, contentDescription = null, contentScale = ContentScale.Crop, modifier = Modifier.size(40.dp))
            }
        } else {
            FactionMark(head.faction.orEmpty(), 40.dp)
        }
    }
    val text = @Composable {
        Column {
            Text(head.name.uppercase(), style = FsType.cardHead, color = FsColor.strong, modifier = Modifier.semantics { heading() })
            head.line?.let { Text(it, style = FsType.meta, color = FsColor.muted) }
        }
    }
    if (isLargeText()) {
        Column(verticalArrangement = Arrangement.spacedBy(8.dp), modifier = Modifier.semantics(mergeDescendants = true) {}) { crest(); text() }
    } else {
        Row(horizontalArrangement = Arrangement.spacedBy(12.dp), verticalAlignment = Alignment.CenterVertically, modifier = Modifier.semantics(mergeDescendants = true) {}) {
            crest()
            text()
        }
    }
}

/** The next raid card (spec F4.1 item 2), the tab's one primary surface. */
@Composable
fun NextRaidCard(
    controller: SignupController,
    clock: RaidClock,
    now: () -> Instant,
    currentKey: String?,
    requestPermission: () -> Unit,
    open: (Boolean) -> Unit,
) {
    val raid = controller.raid
    val facts = controller.facts
    val a11y = listOf(clock.longDate(facts.startsAt), raid.title).joinToString(", ")
    Column(verticalArrangement = Arrangement.spacedBy(14.dp), modifier = Modifier.raidPrimaryCard()) {
        Column(verticalArrangement = Arrangement.spacedBy(14.dp), modifier = Modifier.clickable { open(false) }) {
            val heading = @Composable {
                Column(verticalArrangement = Arrangement.spacedBy(4.dp), modifier = Modifier.semantics(mergeDescendants = true) { contentDescription = a11y }) {
                    SectionLabel(stringResource(R.string.raids_next), color = FsColor.gold)
                    Text(raid.title.uppercase(), style = FsType.cardHead, color = FsColor.strong)
                    TimesLine(facts, clock, showsEnd = false)
                    CountdownLine(facts, clock, now)
                }
            }
            if (isLargeText()) {
                // Large text: the date block above the title instead of beside it.
                Column(verticalArrangement = Arrangement.spacedBy(10.dp)) {
                    DateBlock(facts.startsAt, clock)
                    heading()
                }
            } else {
                Row(horizontalArrangement = Arrangement.spacedBy(12.dp)) {
                    DateBlock(facts.startsAt, clock)
                    heading()
                }
            }
            CountBlock(facts)
        }
        HorizontalDivider(color = FsColor.lineSoft)
        SectionLabel(stringResource(R.string.raid_yourSignup))
        facts.viewerCharacters(currentKey).firstOrNull()?.let { SignupBlockView(controller, it, showsNote = false, reservesStatus = false) }
        if (controller.showsReminderAsk) {
            ReminderAsk(turnOn = { controller.dismissReminderAsk(); requestPermission() }, notNow = { controller.dismissReminderAsk() })
        }
        GoldLink(stringResource(R.string.raids_whoIsComing), { open(true) }, trailing = true)
    }
    controller.specAsk?.let { ask ->
        controller.character(ask.characterKey)?.let { character ->
            val scope = androidx.compose.runtime.rememberCoroutineScope()
            SpecSheet(ask, character, raid.title, clock.shortDate(facts.startsAt), pick = { option ->
                scope.launch { controller.pickSpec(option) }
            }, close = { controller.closeSpecSheet() })
        }
    }
}

/** An Also planned row (min 76): date block, title, when, count with a 64 dp bar, one chip. */
@Composable
fun PlannedRow(raid: RaidView, clock: RaidClock, now: Instant, currentKey: String?, open: () -> Unit) {
    val facts = RaidFacts(raid, now)
    val res = androidx.compose.ui.platform.LocalResources.current
    val large = isLargeText()
    Row(
        horizontalArrangement = Arrangement.spacedBy(12.dp),
        modifier = Modifier.fillMaxWidth().clickable(onClick = open).heightIn(min = 76.dp).padding(14.dp).semantics(mergeDescendants = true) {},
    ) {
        // The date block keeps its width; the text column takes the rest.
        DateBlock(facts.startsAt, clock, if (large) 56.dp else 48.dp, 52.dp)
        Column(verticalArrangement = Arrangement.spacedBy(4.dp), modifier = Modifier.weight(1f)) {
            Text(raid.title, style = FsType.bodyStrong, color = FsColor.strong, maxLines = 2)
            Text(clock.whenPhrase(facts.startsAt, now), style = FsType.meta, color = FsColor.muted)
            Row(horizontalArrangement = Arrangement.spacedBy(8.dp), verticalAlignment = Alignment.CenterVertically) {
                Text(if (raid.size == null) facts.countLine(res) else facts.ratio, style = FsType.monoMeta, color = FsColor.strong)
                if (!large) facts.fill?.let { FillBar(it, 64.dp) }
            }
            // Large text: the chip moves under the title.
            if (large) PlannedChip(facts, currentKey)
        }
        if (!large) PlannedChip(facts, currentKey)
    }
}

@Composable
private fun PlannedChip(facts: RaidFacts, currentKey: String?) {
    run {
        when (val chip = facts.rowChip(currentKey)) {
            is RowChip.Status -> ChipTag(chip.chip)
            is RowChip.Signed -> Tag(
                if (chip.state.isComing) {
                    stringResource(R.string.raids_chip_signed, stringResource(chip.state.label), stringResource(roleLabel(chip.role)))
                } else {
                    stringResource(chip.state.label)
                },
                FsColor.kill,
            )
            RowChip.NotSigned -> Tag(stringResource(R.string.raid_chip_notSigned), FsColor.gold)
        }
    }
}

@Composable
private fun NightRow(night: RaidNightRow, clock: RaidClock, open: () -> Unit) {
    val meta = stringResource(
        R.string.raids_nightMeta,
        clock.shortDate(Instant.parse(night.createdAt)),
        androidx.compose.ui.res.pluralStringResource(R.plurals.night_kill, night.killCount, night.killCount),
        androidx.compose.ui.res.pluralStringResource(R.plurals.night_wipes, night.wipeCount, night.wipeCount),
    )
    Row(
        verticalAlignment = Alignment.CenterVertically,
        modifier = Modifier.fillMaxWidth().clickable(onClick = open).heightIn(min = 56.dp).padding(14.dp).semantics(mergeDescendants = true) {},
    ) {
        Column(Modifier.weight(1f)) {
            Text(night.title, style = FsType.bodyStrong, color = FsColor.strong)
            Text(meta, style = FsType.meta, color = FsColor.muted)
        }
        Icon(Icons.AutoMirrored.Filled.KeyboardArrowRight, contentDescription = null, tint = FsColor.muted)
    }
}

/** First load with nothing cached: the head, the card and a row at their real heights. */
@Composable
private fun Skeleton() {
    val bar = @Composable { w: Int, h: Int ->
        Box(Modifier.width(w.dp).height(h.dp).background(FsColor.lineSoft, RoundedCornerShape(4.dp)))
    }
    Column(verticalArrangement = Arrangement.spacedBy(22.dp), modifier = Modifier.clearAndSetSemantics {}) {
        Row(horizontalArrangement = Arrangement.spacedBy(12.dp)) {
            Box(Modifier.size(40.dp).background(FsColor.lineSoft, CircleShape))
            Column(verticalArrangement = Arrangement.spacedBy(6.dp)) { bar(160, 22); bar(220, 13) }
        }
        Column(verticalArrangement = Arrangement.spacedBy(14.dp), modifier = Modifier.raidPrimaryCard()) {
            Row(horizontalArrangement = Arrangement.spacedBy(12.dp)) {
                Box(Modifier.size(52.dp, 56.dp).background(FsColor.lineSoft, RoundedCornerShape(FsRadius.control)))
                Column(verticalArrangement = Arrangement.spacedBy(6.dp)) { bar(70, 11); bar(180, 22); bar(200, 15); bar(150, 15) }
            }
            bar(240, 15)
            Box(Modifier.fillMaxWidth().height(48.dp).background(FsColor.lineSoft, RoundedCornerShape(FsRadius.control)))
            Spacer(Modifier.height(4.dp))
            bar(90, 20)
        }
    }
}
