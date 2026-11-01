package gg.foreversixty.app.guides.ui

import androidx.compose.foundation.clickable
import androidx.compose.foundation.layout.Arrangement
import androidx.compose.foundation.layout.Column
import androidx.compose.foundation.layout.ExperimentalLayoutApi
import androidx.compose.foundation.layout.FlowRow
import androidx.compose.foundation.layout.Row
import androidx.compose.foundation.layout.fillMaxWidth
import androidx.compose.foundation.layout.heightIn
import androidx.compose.foundation.layout.padding
import androidx.compose.foundation.layout.widthIn
import androidx.compose.material3.Text
import androidx.compose.runtime.Composable
import androidx.compose.ui.Alignment
import androidx.compose.ui.Modifier
import androidx.compose.ui.platform.LocalResources
import androidx.compose.ui.res.stringResource
import androidx.compose.ui.semantics.Role
import androidx.compose.ui.unit.dp
import gg.foreversixty.app.R
import gg.foreversixty.app.common.FactionMark
import gg.foreversixty.app.designsystem.FsColor
import gg.foreversixty.app.designsystem.FsType
import gg.foreversixty.app.designsystem.MicroLabel
import gg.foreversixty.app.designsystem.Tag
import gg.foreversixty.app.designsystem.quietCard
import gg.foreversixty.app.guides.logic.BuildRow
import gg.foreversixty.app.guides.logic.GuideSection
import gg.foreversixty.app.guides.logic.Guides
import gg.foreversixty.app.guides.logic.GuidesCopy
import gg.foreversixty.app.guides.logic.RotationGroup
import gg.foreversixty.app.guides.logic.BuildKind
import gg.foreversixty.app.items.SpecIcon
import gg.foreversixty.app.model.GuideFile
import gg.foreversixty.app.model.GuideKeyTalent
import gg.foreversixty.app.model.GuideRaceRow
import gg.foreversixty.app.model.GuideRotation
import gg.foreversixty.app.model.GuideSpell
import gg.foreversixty.app.reports.ui.Chevron

/** One section of the reader, under its micro label (spec F9.2, F9.3). */
@Composable
internal fun GuideSectionView(section: GuideSection, guide: GuideFile, actions: GuideReaderActions) {
    val copy = GuidesCopy(LocalResources.current)
    GuideSectionBlock(copy.section(section), Modifier.screenMargins()) {
        when (section) {
            GuideSection.Overview -> Overview(guide)
            GuideSection.Talents -> Talents(guide, copy)
            GuideSection.Rotation -> Rotation(guide.rotation, guide.build, copy)
            GuideSection.Stats -> Stats(guide, copy)
            GuideSection.Gear -> Gear(guide, actions)
            GuideSection.Races -> Races(guide.races)
        }
    }
}

@Composable
private fun Overview(guide: GuideFile) {
    val overview = guide.overview
    Column(Modifier.fillMaxWidth().quietCard().cardInside(), verticalArrangement = Arrangement.spacedBy(GuideSpace.rowGap)) {
        Text(guide.summary, style = FsType.body, color = FsColor.text)
        LabelledLines(stringResource(R.string.guides_strengths), overview.strengths)
        LabelledLines(stringResource(R.string.guides_weaknesses), overview.weaknesses)
        LabelledLines(stringResource(R.string.guides_inRaid), overview.raid)
    }
}

@Composable
private fun LabelledLines(label: String, lines: List<String>) {
    if (lines.isEmpty()) return
    Column(verticalArrangement = Arrangement.spacedBy(6.dp)) {
        MicroLabel(label)
        Lines(lines, bulleted = true)
    }
}

/** Every key talent with its icon, ranks and the game's text; then the raid and leveling builds. */
@Composable
private fun Talents(guide: GuideFile, copy: GuidesCopy) {
    Column(verticalArrangement = Arrangement.spacedBy(GuideSpace.labelGap)) {
        QuietRows(guide.talents.key) { TalentRow(it, guide.build, copy) }
        QuietRows(Guides.builds(guide.talents)) { BuildRowView(it, copy) }
    }
}

@Composable
private fun TalentRow(talent: GuideKeyTalent, build: String, copy: GuidesCopy) {
    IconRow(talent.icon, build) {
        NameAndFigure(talent.name, copy.ranks(talent.maxRank))
        talent.text?.let { Text(it, style = FsType.meta, color = FsColor.muted) }
        talent.why?.let { Text(it, style = FsType.meta, color = FsColor.text) }
    }
}

@Composable
private fun BuildRowView(row: BuildRow, copy: GuidesCopy) {
    Row(
        Modifier.fillMaxWidth().heightIn(min = TARGET).cardInside(),
        horizontalArrangement = Arrangement.spacedBy(GuideSpace.rowGap),
        verticalAlignment = Alignment.CenterVertically,
    ) {
        Text(copy.build(row.kind), style = FsType.bodyStrong, color = FsColor.strong, modifier = Modifier.weight(1f))
        Text(row.totals, style = FsType.mono, color = FsColor.text)
        if (row.kind == BuildKind.Raid) Text(stringResource(R.string.guides_buildLevel), style = FsType.meta, color = FsColor.muted)
    }
}

/** Opener, Priority and Cooldowns: numbered rows of icon, name, `Rank {n}` and the clause; then the footer. */
@Composable
private fun Rotation(rotation: GuideRotation?, build: String, copy: GuidesCopy) {
    Column(verticalArrangement = Arrangement.spacedBy(GuideSpace.labelGap)) {
        Guides.rotation(rotation).forEach { group -> RotationGroupView(group, build, copy) }
        Text(stringResource(R.string.guides_rotationFooter), style = FsType.meta, color = FsColor.muted)
    }
}

@Composable
private fun RotationGroupView(group: RotationGroup, build: String, copy: GuidesCopy) {
    Column(verticalArrangement = Arrangement.spacedBy(6.dp)) {
        MicroLabel(copy.rotation(group.kind))
        QuietRows(group.rows) { row ->
            Row(Modifier.fillMaxWidth().padding(start = GuideSpace.margin), verticalAlignment = Alignment.CenterVertically) {
                Text(row.number.toString(), style = FsType.mono, color = FsColor.muted, modifier = Modifier.widthIn(min = NUMBER_WIDTH))
                SpellRow(row.spell, build, copy)
            }
        }
    }
}

@Composable
private fun SpellRow(spell: GuideSpell, build: String, copy: GuidesCopy) {
    IconRow(spell.icon, build, inset = false, tile = Guides.hasTile(spell)) {
        NameAndFigure(spell.name, copy.rank(spell.rank))
        Text(spell.clause, style = FsType.meta, color = FsColor.muted)
    }
}

/** The priority again, and the hit-cap line when the guide has one; no weights (the sims own numbers). */
@Composable
private fun Stats(guide: GuideFile, copy: GuidesCopy) {
    Column(Modifier.fillMaxWidth().quietCard().cardInside(), verticalArrangement = Arrangement.spacedBy(GuideSpace.rowGap)) {
        NumberedChips(guide.statPriority, copy::statLabel)
        guide.hitCap?.let { Text(it, style = FsType.body, color = FsColor.text) }
    }
}

/** One link row to the BiS tab on this spec's list; no item list here (BiS owns gear). */
@Composable
private fun Gear(guide: GuideFile, actions: GuideReaderActions) {
    val gear = guide.gear ?: return
    Row(
        Modifier.fillMaxWidth().quietCard().heightIn(min = TARGET).clickable(role = Role.Button) { actions.openBis(guide) }.cardInside(),
        verticalAlignment = Alignment.CenterVertically,
    ) {
        Text(stringResource(R.string.guides_gearRow, guide.specName, gear.presetLabel), style = FsType.bodyStrong, color = FsColor.gold, modifier = Modifier.weight(1f))
        Chevron(CHEVRON)
    }
}

@OptIn(ExperimentalLayoutApi::class)
@Composable
private fun Races(races: List<GuideRaceRow>) {
    QuietRows(races) { race ->
        Column(Modifier.fillMaxWidth().cardInside(), verticalArrangement = Arrangement.spacedBy(6.dp)) {
            FlowRow(horizontalArrangement = Arrangement.spacedBy(GuideSpace.labelGap), verticalArrangement = Arrangement.spacedBy(6.dp), itemVerticalAlignment = Alignment.CenterVertically) {
                FactionMark(race.faction, FACTION)
                Text(race.name, style = FsType.bodyStrong, color = FsColor.strong)
                if (race.recommended) Tag(stringResource(R.string.guides_recommended), FsColor.kill)
            }
            race.racials.forEach { racial -> Text("${racial.name}: ${racial.effect}", style = FsType.meta, color = FsColor.muted) }
        }
    }
}

/**
 * A row led by a 40 dp game icon: a missing icon leaves the bordered tile, never a letter. Without
 * [tile] (a line that names no spell) the words start where the tile would.
 */
@Composable
private fun IconRow(icon: String?, build: String, inset: Boolean = true, tile: Boolean = true, words: @Composable () -> Unit) {
    Row(
        Modifier.fillMaxWidth().heightIn(min = ROW_MIN).then(if (inset) Modifier.cardInside() else Modifier.padding(end = GuideSpace.margin, top = GuideSpace.inCard, bottom = GuideSpace.inCard)),
        horizontalArrangement = Arrangement.spacedBy(GuideSpace.rowGap),
        verticalAlignment = Alignment.CenterVertically,
    ) {
        if (tile) SpecIcon(icon, build, ICON)
        Column(Modifier.weight(1f), verticalArrangement = Arrangement.spacedBy(2.dp)) { words() }
    }
}

/** `Lava Burst  Rank 4`: the name, then its figure in mono; wraps rather than cuts. */
@OptIn(ExperimentalLayoutApi::class)
@Composable
private fun NameAndFigure(name: String, figure: String?) {
    FlowRow(horizontalArrangement = Arrangement.spacedBy(6.dp), itemVerticalAlignment = Alignment.Bottom) {
        Text(name, style = FsType.bodyStrong, color = FsColor.strong)
        figure?.let { Text(it, style = FsType.monoMeta, color = FsColor.muted) }
    }
}

private val ICON = 40.dp
private val FACTION = 24.dp
private val ROW_MIN = 68.dp
private val NUMBER_WIDTH = 28.dp
private val CHEVRON = 20.dp
