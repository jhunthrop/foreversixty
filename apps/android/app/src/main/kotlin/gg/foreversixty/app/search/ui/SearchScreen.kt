package gg.foreversixty.app.search.ui

import androidx.annotation.DrawableRes
import androidx.compose.foundation.shape.RoundedCornerShape
import androidx.compose.ui.unit.sp
import gg.foreversixty.app.designsystem.FsRadius
import gg.foreversixty.app.search.logic.WebHandoff
import androidx.compose.foundation.background
import androidx.compose.foundation.border
import androidx.compose.foundation.horizontalScroll
import androidx.compose.foundation.layout.Arrangement
import androidx.compose.foundation.layout.Box
import androidx.compose.foundation.layout.Column
import androidx.compose.foundation.layout.Row
import androidx.compose.foundation.layout.fillMaxSize
import androidx.compose.foundation.layout.fillMaxWidth
import androidx.compose.foundation.layout.height
import androidx.compose.foundation.layout.heightIn
import androidx.compose.foundation.layout.padding
import androidx.compose.foundation.lazy.LazyColumn
import androidx.compose.foundation.lazy.LazyListScope
import androidx.compose.foundation.rememberScrollState
import androidx.compose.foundation.selection.selectable
import androidx.compose.foundation.shape.CircleShape
import androidx.compose.material.icons.Icons
import androidx.compose.material.icons.automirrored.filled.ArrowBack
import androidx.compose.material.icons.filled.Close
import androidx.compose.material.icons.filled.Refresh
import androidx.compose.material3.ExperimentalMaterial3Api
import androidx.compose.material3.Icon
import androidx.compose.material3.IconButton
import androidx.compose.material3.SearchBarDefaults
import androidx.compose.material3.Text
import androidx.compose.material3.TextButton
import androidx.compose.material3.TextFieldDefaults
import androidx.compose.runtime.Composable
import androidx.compose.runtime.LaunchedEffect
import androidx.compose.runtime.remember
import androidx.compose.ui.Alignment
import androidx.compose.ui.Modifier
import androidx.compose.ui.focus.FocusRequester
import androidx.compose.ui.focus.focusRequester
import androidx.compose.ui.graphics.Color
import androidx.compose.ui.platform.LocalResources
import androidx.compose.ui.res.stringResource
import androidx.compose.ui.semantics.Role
import androidx.compose.ui.text.font.FontWeight
import androidx.compose.ui.unit.dp
import gg.foreversixty.app.R
import gg.foreversixty.app.designsystem.BrandButton
import gg.foreversixty.app.designsystem.BrandButtonKind
import gg.foreversixty.app.designsystem.FsColor
import gg.foreversixty.app.designsystem.FsType
import gg.foreversixty.app.designsystem.OfflineStrip
import gg.foreversixty.app.designsystem.quietCard
import gg.foreversixty.app.search.logic.CharacterPhase
import gg.foreversixty.app.search.logic.IndexPhase
import gg.foreversixty.app.search.logic.SearchCopy
import gg.foreversixty.app.search.logic.SearchScope
import gg.foreversixty.app.search.logic.SearchSection
import gg.foreversixty.app.search.logic.SearchTarget

/**
 * The For card's three links (spec F12.1): the character's list, their spec's guide (null when
 * the spec is unknown) and Rankings.
 */
data class ForCharacter(val name: String, val guideLabel: String?)

/** Everything F12.1 to F12.3 and F12.6 draw, pure so the screenshot tests set it directly. */
data class SearchScreenState(
    val query: String,
    val scope: SearchScope,
    val indexPhase: IndexPhase,
    val sections: List<SearchSection>,
    val characterPhase: CharacterPhase,
    val recents: List<SearchTarget>,
    val forCharacter: ForCharacter?,
    val build: String,
    /** The signed-in account's character keys: their rows carry `You`. */
    val ownKeys: Set<String> = emptySet(),
)

/** What the screen's controls do; the route wires them. */
data class SearchActions(
    val onQuery: (String) -> Unit = {},
    val onScope: (SearchScope) -> Unit = {},
    val onBack: () -> Unit = {},
    val onOpen: (SearchTarget) -> Unit = {},
    val onClearRecent: () -> Unit = {},
    val onRetryIndex: () -> Unit = {},
    val onOpenBis: () -> Unit = {},
    val onOpenGuide: () -> Unit = {},
    val onOpenRankings: () -> Unit = {},
)

/** F12: the field with the keyboard up, the scope chips, then the empty state or the results. */
@Composable
fun SearchScreen(state: SearchScreenState, actions: SearchActions, focusField: Boolean = true) {
    val copy = SearchCopy(LocalResources.current)
    Column(Modifier.fillMaxSize().background(FsColor.bg)) {
        SearchField(state.query, actions, focusField)
        ScopeChips(state.scope, copy, actions.onScope)
        LazyColumn(Modifier.fillMaxSize().weight(1f)) {
            if (state.query.isBlank()) emptyState(state, copy, actions) else results(state, copy, actions)
            item { Box(Modifier.height(24.dp)) }
        }
    }
}

/** The Material 3 search field: a 56 dp pill with the back arrow, the placeholder and clear. */
@OptIn(ExperimentalMaterial3Api::class)
@Composable
private fun SearchField(query: String, actions: SearchActions, focus: Boolean) {
    val focusRequester = remember { FocusRequester() }
    LaunchedEffect(focus) { if (focus) focusRequester.requestFocus() }
    val transparent = Color.Transparent
    Box(
        // 56 dp at the default size; it grows with the text so a large query is never cut.
        Modifier.padding(horizontal = SearchLayout.margin, vertical = 8.dp).fillMaxWidth().heightIn(min = 56.dp)
            .background(FsColor.lineSoft, CircleShape),
        contentAlignment = Alignment.Center,
    ) {
        SearchBarDefaults.InputField(
            query = query,
            onQueryChange = actions.onQuery,
            onSearch = {},
            expanded = true,
            onExpandedChange = {},
            modifier = Modifier.fillMaxWidth().focusRequester(focusRequester),
            placeholder = { Text(stringResource(R.string.search_placeholder), style = FsType.lead.copy(fontWeight = FontWeight.Normal), color = FsColor.muted, maxLines = 1) },
            leadingIcon = {
                IconButton(onClick = actions.onBack) {
                    Icon(Icons.AutoMirrored.Filled.ArrowBack, contentDescription = stringResource(R.string.search_cancel), tint = FsColor.strong)
                }
            },
            trailingIcon = if (query.isEmpty()) {
                null
            } else {
                {
                    IconButton(onClick = { actions.onQuery("") }) {
                        Icon(Icons.Filled.Close, contentDescription = stringResource(R.string.search_clear), tint = FsColor.muted)
                    }
                }
            },
            colors = TextFieldDefaults.colors(
                focusedContainerColor = transparent,
                unfocusedContainerColor = transparent,
                focusedIndicatorColor = transparent,
                unfocusedIndicatorColor = transparent,
                focusedTextColor = FsColor.strong,
                unfocusedTextColor = FsColor.strong,
                cursorColor = FsColor.gold,
            ),
        )
    }
}

/** The five scope chips: 36 dp pills in 48 dp targets, one selected (spec §6: `Items, selected`). */
@Composable
private fun ScopeChips(selected: SearchScope, copy: SearchCopy, onScope: (SearchScope) -> Unit) {
    Row(
        Modifier.fillMaxWidth().horizontalScroll(rememberScrollState()).padding(horizontal = SearchLayout.margin),
        horizontalArrangement = Arrangement.spacedBy(CHIP_GAP),
    ) {
        SearchScope.entries.forEach { scope ->
            ScopeChip(copy.scopeLabel(scope), scope == selected) { onScope(scope) }
        }
    }
}

/** A 36 dp chip inside a 48 dp target; the selected one is gold. Shared with the item page's strip. */
@Composable
fun ScopeChip(label: String, selected: Boolean, onClick: () -> Unit) {
    Box(
        Modifier.heightIn(min = 48.dp).selectable(selected = selected, role = Role.Tab, onClick = onClick),
        contentAlignment = Alignment.Center,
    ) {
        val line = if (selected) FsColor.gold else FsColor.line
        Box(
            Modifier.heightIn(min = 36.dp)
                .background(if (selected) FsColor.gold.copy(alpha = SELECTED_WASH) else Color.Transparent, CircleShape)
                .border(1.dp, line, CircleShape)
                .padding(horizontal = CHIP_PADDING, vertical = 6.dp),
            contentAlignment = Alignment.Center,
        ) {
            Text(
                label,
                style = FsType.body.copy(fontSize = CHIP_TEXT, fontWeight = if (selected) FontWeight.Bold else FontWeight.Normal),
                color = if (selected) FsColor.gold else FsColor.text,
            )
        }
    }
}

private const val SELECTED_WASH = 0.10f

/** Sized so the five scope chips fit whole at 412 dp (S8); narrower phones and large text scroll. */
private val CHIP_PADDING = 12.dp
private val CHIP_GAP = 6.dp
private val CHIP_TEXT = 15.sp

/** F12.1: For {name}, then Recent with its Clear. A group with nothing in it is not drawn. */
private fun LazyListScope.emptyState(state: SearchScreenState, copy: SearchCopy, actions: SearchActions) {
    state.forCharacter?.let { mine ->
        item(key = "for") {
            SectionLabel(stringResource(R.string.search_group_forCharacter, mine.name))
            val links = buildList {
                add(QuickLink(R.drawable.ic_shield, stringResource(R.string.search_forCharacter_bis, mine.name), RowTrailing.Chevron, actions.onOpenBis))
                mine.guideLabel?.let { add(QuickLink(R.drawable.ic_guide, it, RowTrailing.of(WebHandoff.GUIDES, RowTrailing.Chevron), actions.onOpenGuide)) }
                add(
                    QuickLink(
                        R.drawable.ic_trophy, stringResource(R.string.search_forCharacter_rankings),
                        RowTrailing.of(WebHandoff.RANKINGS, RowTrailing.Chevron), actions.onOpenRankings,
                    ),
                )
            }
            RowCard(links.size) { i -> LinkRow(links[i].icon, links[i].label, links[i].trailing, links[i].onClick) }
        }
    }
    // Watching (F11's /v1/me/watching) slots in here when that lane lands.
    if (state.recents.isNotEmpty()) {
        item(key = "recent") {
            SectionLabel(stringResource(R.string.search_group_recent)) {
                TextButton(onClick = actions.onClearRecent, modifier = Modifier.heightIn(min = 48.dp)) {
                    Text(stringResource(R.string.search_recent_clear), style = FsType.bodyStrong, color = FsColor.gold)
                }
            }
            RowCard(state.recents.size) { i ->
                val target = state.recents[i]
                TargetRow(
                    target, copy.title(target), copy.kind(target), state.build, titleColor(target),
                    trailing = RowTrailing.of(WebHandoff.opensWebsite(target)),
                ) { actions.onOpen(target) }
            }
        }
    }
    indexState(state.indexPhase, actions)
}

/** F12.2, F12.3, F12.6: the groups, the offline strip, no results. */
private fun LazyListScope.results(state: SearchScreenState, copy: SearchCopy, actions: SearchActions) {
    val offline = state.characterPhase == CharacterPhase.Offline
    val ready = state.indexPhase is IndexPhase.Ready
    if (state.sections.isEmpty() && ready) {
        item(key = "none") { NoResults(state.query.trim()) }
        if (offline) item(key = "offline") { InsetStrip() }
        return
    }
    if (offline) item(key = "offline") { InsetStrip() }
    indexState(state.indexPhase, actions)
    state.sections.forEach { section ->
        item(key = section.group.name) {
            SectionLabel(copy.groupLabel(section.group))
            val rows = section.rows.size + if (section.showsAll) 1 else 0
            RowCard(rows) { i ->
                if (i < section.rows.size) {
                    ResultRow(section.rows[i], copy, state, actions)
                } else {
                    ShowAllRow(section.total) { actions.onScope(section.group.scope) }
                }
            }
        }
    }
    if (state.characterPhase == CharacterPhase.Failed && state.scope in setOf(SearchScope.All, SearchScope.Characters)) {
        item(key = "charactersFailed") {
            Text(
                stringResource(R.string.search_charactersFailed), style = FsType.body, color = FsColor.wipe,
                modifier = Modifier.padding(horizontal = SearchLayout.margin, vertical = 14.dp),
            )
        }
    }
}

@Composable
private fun ResultRow(target: SearchTarget, copy: SearchCopy, state: SearchScreenState, actions: SearchActions) {
    val you = (target as? SearchTarget.Character)?.character?.key?.takeIf { it in state.ownKeys }
        ?.let { stringResource(R.string.search_character_you) }
    // The watch star (F11) joins character rows when /v1/me/watching exists.
    TargetRow(
        target, copy.title(target), copy.meta(target), state.build, titleColor(target),
        youTag = you, trailing = RowTrailing.of(WebHandoff.opensWebsite(target)),
    ) { actions.onOpen(target) }
}

@Composable
private fun ShowAllRow(total: Int, onClick: () -> Unit) {
    TappableRow(onClick, minHeight = 52.dp) {
        Text(stringResource(R.string.search_showAll, total), style = FsType.bodyStrong, color = FsColor.gold, modifier = Modifier.weight(1f))
        Chevron()
    }
}

/** The index's own states: loading, or failed with Try again (spec F12, §5.0). */
private fun LazyListScope.indexState(phase: IndexPhase, actions: SearchActions) {
    when (phase) {
        IndexPhase.Loading -> item(key = "indexLoading") {
            Text(
                stringResource(R.string.search_indexLoading), style = FsType.body, color = FsColor.muted,
                modifier = Modifier.padding(horizontal = SearchLayout.margin, vertical = 14.dp),
            )
        }
        IndexPhase.Failed -> item(key = "indexFailed") {
            Column(Modifier.padding(horizontal = SearchLayout.margin, vertical = 14.dp), verticalArrangement = Arrangement.spacedBy(10.dp)) {
                Text(stringResource(R.string.search_indexFailed), style = FsType.body, color = FsColor.wipe)
                BrandButton(stringResource(R.string.search_tryAgain), BrandButtonKind.Secondary, actions.onRetryIndex, height = 48.dp, fill = false, leadingIcon = Icons.Filled.Refresh)
            }
        }
        is IndexPhase.Ready -> Unit
    }
}

@Composable
private fun NoResults(query: String) {
    Column(
        Modifier.padding(horizontal = SearchLayout.margin).padding(top = 16.dp).fillMaxWidth().quietCard().padding(SearchLayout.margin),
        verticalArrangement = Arrangement.spacedBy(6.dp),
    ) {
        Text(stringResource(R.string.search_noResults_title, query), style = FsType.lead.copy(fontWeight = FontWeight.Normal), color = FsColor.strong)
        Text(stringResource(R.string.search_noResults_body), style = FsType.body, color = FsColor.muted)
    }
}

/** The offline strip inside the results (spec F12.6), a rounded card like the board's. */
@Composable
private fun InsetStrip() {
    // All four corners round, as on iOS.
    Box(Modifier.padding(horizontal = SearchLayout.margin).padding(top = 16.dp)) {
        OfflineStrip(stringResource(R.string.search_offline), RoundedCornerShape(FsRadius.panel))
    }
}

/** One For card link: its glyph, words, and where it goes (in the app or the website). */
private data class QuickLink(@param:DrawableRes val icon: Int, val label: String, val trailing: RowTrailing, val onClick: () -> Unit)
