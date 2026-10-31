package gg.foreversixty.app.search.logic

import androidx.compose.runtime.getValue
import androidx.compose.runtime.mutableStateOf
import androidx.compose.runtime.setValue
import gg.foreversixty.app.model.CharacterSearch
import gg.foreversixty.app.model.CharacterSearchRow
import gg.foreversixty.app.network.ApiClient
import gg.foreversixty.app.network.ApiEndpoints
import gg.foreversixty.app.network.ApiError
import gg.foreversixty.app.network.ApiResult
import kotlin.coroutines.CoroutineContext
import kotlinx.coroutines.CoroutineScope
import kotlinx.coroutines.Dispatchers
import kotlinx.coroutines.Job
import kotlinx.coroutines.delay
import kotlinx.coroutines.launch
import kotlinx.coroutines.withContext

/** `GET /v1/characters/search`. Live goes through [ApiClient]; tests pass a fake. */
fun interface CharacterSearchService {
    suspend fun search(query: String): ApiResult<CharacterSearch>
}

class LiveCharacterSearch(private val client: ApiClient) : CharacterSearchService {
    override suspend fun search(query: String) =
        client.send(ApiEndpoints.searchCharacters(q = query, limit = SearchLogic.CHARACTER_LIMIT))
}

/** The search index for the active build, prepared; null when it could not be read. */
fun interface SearchIndexSource {
    suspend fun load(): PreparedIndex?
}

sealed interface IndexPhase {
    data object Loading : IndexPhase
    data class Ready(val index: PreparedIndex) : IndexPhase
    data object Failed : IndexPhase
}

sealed interface CharacterPhase {
    /** Not asked: under two letters, or a scope without characters. */
    data object Idle : CharacterPhase
    data class Loaded(val rows: List<CharacterSearchRow>) : CharacterPhase

    /** No connection: the results say characters need one (spec F12.6). */
    data object Offline : CharacterPhase
    data object Failed : CharacterPhase
}

/** Whose search it is: the current character's class (items it can use first), region and ruleset. */
data class SearchContext(val classSlug: String? = null, val region: String? = null, val ruleset: String? = null)

/**
 * One search session (spec F12.1, F12.2): the field, the scope, the phone-side matches (filtered at
 * once, off the main thread) and the characters (asked 250 ms after typing stops). [scope] owns
 * the jobs; [compute] is where the filter runs.
 */
class SearchStore(
    private val index: SearchIndexSource,
    private val characters: CharacterSearchService,
    private val context: SearchContext,
    private val scope: CoroutineScope,
    private val compute: CoroutineContext = Dispatchers.Default,
) {
    var indexPhase: IndexPhase by mutableStateOf(IndexPhase.Loading)
        private set
    var query: String by mutableStateOf("")
        private set
    var selectedScope: SearchScope by mutableStateOf(SearchScope.All)
        private set
    var local: LocalResults by mutableStateOf(LocalResults())
        private set
    var characterPhase: CharacterPhase by mutableStateOf(CharacterPhase.Idle)
        private set

    private var filterJob: Job? = null
    private var characterJob: Job? = null
    private var askedFor: String? = null

    val normalized: String get() = SearchLogic.normalize(query)

    val sections: List<SearchSection>
        get() = SearchLogic.sections(local, (characterPhase as? CharacterPhase.Loaded)?.rows.orEmpty(), selectedScope)

    suspend fun loadIndex() {
        indexPhase = IndexPhase.Loading
        indexPhase = index.load()?.let { IndexPhase.Ready(it) } ?: IndexPhase.Failed
        refilter()
    }

    fun type(raw: String) {
        query = raw
        refilter()
        askCharacters()
    }

    fun choose(next: SearchScope) {
        selectedScope = next
        askCharacters()
    }

    private fun refilter() {
        filterJob?.cancel()
        val ready = (indexPhase as? IndexPhase.Ready)?.index
        val q = normalized
        if (ready == null || q.isEmpty()) {
            local = LocalResults()
            return
        }
        filterJob = scope.launch { local = withContext(compute) { SearchLogic.filter(ready, q, context.classSlug) } }
    }

    private fun askCharacters() {
        val q = normalized
        if (!SearchLogic.asksCharacters(q, selectedScope)) {
            characterJob?.cancel()
            askedFor = null
            characterPhase = CharacterPhase.Idle
            return
        }
        if (q == askedFor) return
        characterJob?.cancel()
        askedFor = q
        characterJob = scope.launch {
            delay(SearchLogic.CHARACTER_DEBOUNCE_MILLIS)
            characterPhase = when (val answer = characters.search(q)) {
                is ApiResult.Success -> CharacterPhase.Loaded(SearchLogic.orderCharacters(answer.data.rows, context.region, context.ruleset))
                is ApiResult.Failure -> if (answer.error is ApiError.Network) CharacterPhase.Offline else CharacterPhase.Failed
            }
        }
    }
}
