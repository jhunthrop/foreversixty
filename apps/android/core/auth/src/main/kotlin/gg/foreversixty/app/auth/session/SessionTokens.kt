package gg.foreversixty.app.auth.session

import gg.foreversixty.app.auth.TokenStore
import kotlinx.coroutines.channels.BufferOverflow
import kotlinx.coroutines.flow.MutableSharedFlow
import kotlinx.coroutines.flow.SharedFlow
import kotlinx.coroutines.flow.asSharedFlow

/**
 * The token store the HTTP client and sign-in share. It forwards to the real store and reports
 * every 401 discard on [discarded], so [SessionManager] shows signed out the moment any request
 * learns the token is dead (`unauthorized` or `session_expired`).
 */
class SessionTokens(private val base: TokenStore) : TokenStore by base {
    private val events = MutableSharedFlow<Unit>(extraBufferCapacity = 1, onBufferOverflow = BufferOverflow.DROP_OLDEST)

    val discarded: SharedFlow<Unit> = events.asSharedFlow()

    override suspend fun discardBearerToken() {
        base.discardBearerToken()
        events.tryEmit(Unit)
    }
}

/** One small non-secret string kept between launches (the picked character key). */
interface StoredValue {
    fun load(): String?
    fun save(value: String?)
}

class MemoryValue(private var value: String? = null) : StoredValue {
    override fun load(): String? = value
    override fun save(value: String?) {
        this.value = value
    }
}
