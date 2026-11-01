package gg.foreversixty.app.rankings.ui

import androidx.compose.foundation.border
import androidx.compose.foundation.clickable
import androidx.compose.foundation.layout.Arrangement
import androidx.compose.foundation.layout.Box
import androidx.compose.foundation.layout.Column
import androidx.compose.foundation.layout.Row
import androidx.compose.foundation.layout.fillMaxHeight
import androidx.compose.foundation.layout.fillMaxWidth
import androidx.compose.foundation.layout.heightIn
import androidx.compose.foundation.layout.navigationBarsPadding
import androidx.compose.foundation.layout.padding
import androidx.compose.foundation.layout.size
import androidx.compose.foundation.rememberScrollState
import androidx.compose.foundation.selection.selectable
import androidx.compose.foundation.selection.selectableGroup
import androidx.compose.foundation.selection.toggleable
import androidx.compose.foundation.shape.RoundedCornerShape
import androidx.compose.foundation.verticalScroll
import androidx.compose.material.icons.Icons
import androidx.compose.material.icons.automirrored.filled.ArrowBack
import androidx.compose.material.icons.filled.Check
import androidx.compose.material.icons.filled.KeyboardArrowDown
import androidx.compose.material.icons.materialIcon
import androidx.compose.material3.ExperimentalMaterial3Api
import androidx.compose.material3.HorizontalDivider
import androidx.compose.material3.Icon
import androidx.compose.material3.IconButton
import androidx.compose.material3.ModalBottomSheet
import androidx.compose.material3.Switch
import androidx.compose.material3.SwitchDefaults
import androidx.compose.material3.Text
import androidx.compose.material3.rememberModalBottomSheetState
import androidx.compose.runtime.Composable
import androidx.compose.ui.Alignment
import androidx.compose.ui.Modifier
import androidx.compose.ui.graphics.vector.ImageVector
import androidx.compose.ui.graphics.vector.addPathNodes
import androidx.compose.ui.graphics.SolidColor
import androidx.compose.ui.platform.LocalResources
import androidx.compose.ui.res.stringResource
import androidx.compose.ui.semantics.Role
import androidx.compose.ui.semantics.heading
import androidx.compose.ui.semantics.semantics
import androidx.compose.ui.unit.dp
import gg.foreversixty.app.R
import gg.foreversixty.app.bis.ChoicePills
import gg.foreversixty.app.bis.PickerSection
import gg.foreversixty.app.designsystem.BrandButton
import gg.foreversixty.app.designsystem.BrandButtonKind
import gg.foreversixty.app.designsystem.FsColor
import gg.foreversixty.app.designsystem.FsRadius
import gg.foreversixty.app.designsystem.FsType
import gg.foreversixty.app.designsystem.GameColors
import gg.foreversixty.app.model.RankedEncounter
import gg.foreversixty.app.rankings.logic.Measure
import gg.foreversixty.app.rankings.logic.Rankings
import gg.foreversixty.app.rankings.logic.RankingsCopy
import gg.foreversixty.app.rankings.logic.RankingsFilters
import gg.foreversixty.app.reports.ui.MetersDragHandle
import java.util.Locale

/** What the filters sheet draws: the board's filters, the encounters and the phases on offer. */
data class FiltersSheetState(
    val filters: RankingsFilters,
    val encounters: List<RankedEncounter>,
    /** The phase ids begun, newest first. */
    val phases: List<String>,
    /** The sheet shows the encounter list instead of the filters. */
    val choosingEncounter: Boolean = false,
)

data class FiltersActions(
    val apply: (RankingsFilters) -> Unit = {},
    val reset: () -> Unit = {},
    val chooseEncounter: (Boolean) -> Unit = {},
)

/** F11.2 at its large detent: every choice applies at once, as the BiS picker does. */
@OptIn(ExperimentalMaterial3Api::class)
@Composable
fun FiltersSheet(state: FiltersSheetState, actions: FiltersActions, onDismiss: () -> Unit) {
    ModalBottomSheet(
        onDismissRequest = onDismiss,
        sheetState = rememberModalBottomSheetState(skipPartiallyExpanded = true),
        containerColor = GameColors.sheet,
        dragHandle = { MetersDragHandle() },
    ) {
        Box(Modifier.fillMaxHeight()) { FiltersSheetContent(state, actions) }
    }
}

/** The sheet's body, apart from the sheet so screenshot tests draw it. */
@Composable
fun FiltersSheetContent(state: FiltersSheetState, actions: FiltersActions) {
    if (state.choosingEncounter) {
        EncounterList(state, actions)
        return
    }
    val copy = RankingsCopy(LocalResources.current)
    val f = state.filters
    val apply = actions.apply
    Column(
        Modifier.verticalScroll(rememberScrollState()).navigationBarsPadding().padding(horizontal = 16.dp).padding(bottom = 24.dp),
        verticalArrangement = Arrangement.spacedBy(22.dp),
    ) {
        Title(stringResource(R.string.rank_filters))
        PickerSection(stringResource(R.string.rank_filter_encounter)) { EncounterField(f, state.encounters, actions) }
        PickerSection(stringResource(R.string.rank_filter_measure)) {
            ChoicePills(Measure.entries, copy::measureLabel, f.measure) { apply(f.copy(measure = it)) }
        }
        PickerSection(stringResource(R.string.rank_filter_ruleset)) {
            ChoicePills(Rankings.rulesets, copy::rulesetLabel, f.ruleset) { apply(f.copy(ruleset = it)) }
        }
        PickerSection(stringResource(R.string.rank_filter_region)) {
            ChoicePills(Rankings.regions, copy::regionChip, f.region) { apply(f.copy(region = it)) }
        }
        if (state.phases.isNotEmpty()) {
            PickerSection(stringResource(R.string.rank_filter_phase)) {
                ChoicePills(state.phases, copy::phaseLabel, f.phase) { apply(f.copy(phase = it)) }
            }
        }
        PickerSection(stringResource(R.string.rank_filter_faction)) {
            ChoicePills(Rankings.factions, copy::factionLabel, f.faction) { apply(f.copy(faction = it)) }
        }
        HorizontalDivider(color = FsColor.lineSoft)
        TodayOnly(f.todayOnly) { apply(f.copy(todayOnly = it)) }
        BrandButton(stringResource(R.string.rank_reset), BrandButtonKind.Secondary, actions.reset, height = 48.dp, fill = false)
    }
}

@Composable
private fun Title(text: String) {
    Text(text.uppercase(Locale.ROOT), style = FsType.sec, color = FsColor.strong, modifier = Modifier.semantics { heading() })
}

/** The Encounter row: the current encounter and a chevron, opening the list. */
@Composable
private fun EncounterField(filters: RankingsFilters, encounters: List<RankedEncounter>, actions: FiltersActions) {
    val name = encounters.firstOrNull { it.slug == filters.encounter }?.name ?: filters.encounter
    val shape = RoundedCornerShape(FsRadius.control)
    Row(
        Modifier.fillMaxWidth().heightIn(min = 56.dp).border(1.dp, FsColor.line, shape)
            .clickable(role = Role.Button) { actions.chooseEncounter(true) }.padding(horizontal = 16.dp),
        verticalAlignment = Alignment.CenterVertically,
    ) {
        Text(name, style = FsType.bodyStrong, color = FsColor.strong, modifier = Modifier.weight(1f))
        Icon(Icons.Filled.KeyboardArrowDown, contentDescription = null, tint = FsColor.muted)
    }
}

@Composable
private fun TodayOnly(on: Boolean, set: (Boolean) -> Unit) {
    Row(
        Modifier.fillMaxWidth().heightIn(min = 48.dp).toggleable(value = on, role = Role.Switch, onValueChange = set),
        verticalAlignment = Alignment.CenterVertically,
    ) {
        Text(stringResource(R.string.rank_filter_today), style = FsType.body, color = FsColor.text, modifier = Modifier.weight(1f))
        Switch(
            checked = on, onCheckedChange = null,
            colors = SwitchDefaults.colors(checkedTrackColor = FsColor.gold, checkedThumbColor = FsColor.bg, uncheckedTrackColor = FsColor.raised, uncheckedBorderColor = FsColor.line),
        )
    }
}

/** The encounters with ranked kills: name and `{zone} · {size} players · {kills}`; a pick applies and returns. */
@Composable
private fun EncounterList(state: FiltersSheetState, actions: FiltersActions) {
    val copy = RankingsCopy(LocalResources.current)
    Column(Modifier.verticalScroll(rememberScrollState()).navigationBarsPadding().padding(bottom = 24.dp)) {
        Row(Modifier.padding(horizontal = 4.dp), verticalAlignment = Alignment.CenterVertically) {
            IconButton(onClick = { actions.chooseEncounter(false) }) {
                Icon(Icons.AutoMirrored.Filled.ArrowBack, contentDescription = stringResource(R.string.rank_filters), tint = FsColor.strong)
            }
            Title(stringResource(R.string.rank_filter_encounter))
        }
        Column(Modifier.selectableGroup().padding(horizontal = 16.dp).padding(top = 8.dp)) {
            state.encounters.forEachIndexed { index, encounter ->
                if (index > 0) HorizontalDivider(color = FsColor.lineSoft)
                val current = encounter.slug == state.filters.encounter
                Row(
                    Modifier.fillMaxWidth().heightIn(min = 64.dp)
                        .selectable(selected = current, role = Role.RadioButton) {
                            actions.apply(state.filters.copy(encounter = encounter.slug))
                            actions.chooseEncounter(false)
                        }
                        .padding(vertical = 10.dp),
                    verticalAlignment = Alignment.CenterVertically,
                    horizontalArrangement = Arrangement.spacedBy(12.dp),
                ) {
                    Column(Modifier.weight(1f), verticalArrangement = Arrangement.spacedBy(2.dp)) {
                        Text(encounter.name, style = FsType.bodyStrong, color = if (current) FsColor.goldHover else FsColor.strong)
                        Text(copy.encounterMeta(encounter), style = FsType.meta, color = FsColor.muted)
                    }
                    if (current) Icon(Icons.Filled.Check, contentDescription = null, tint = FsColor.gold, modifier = Modifier.size(20.dp))
                }
            }
        }
    }
}

/** The Filters button's glyph (Material `tune`), drawn here: the core icon set has no filter glyph. */
val FiltersIcon: ImageVector by lazy {
    materialIcon(name = "Filled.Tune") {
        addPath(addPathNodes(TUNE_PATH), fill = SolidColor(androidx.compose.ui.graphics.Color.Black))
    }
}

private const val TUNE_PATH =
    "M3,17v2h6v-2L3,17zM3,5v2h10L13,5L3,5zM13,21v-2h8v-2h-8v-2h-2v6h2zM7,9v2L3,11v2h4v2h2L9,9L7,9zM21,13v-2L11,11v2h10zM15,9h2L17,7h4L21,5h-4L17,3h-2v6z"
