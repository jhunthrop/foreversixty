package gg.foreversixty.app.model

import kotlinx.serialization.json.Json

/**
 * The one JSON configuration for every API body, used by the network client and by tests.
 *
 * - Unknown keys are ignored: the API adds fields without a version bump (contracts are floors).
 * - A null for a field that has a default takes the default: Go encodes a nil slice as null.
 * - Absent optional fields decode to null; Go's `omitempty` drops them.
 */
val ApiJson: Json = Json {
    ignoreUnknownKeys = true
    coerceInputValues = true
    explicitNulls = false
}
