package gg.foreversixty.app.model

// Hand-written behaviour on the generated models (ApiModels.kt). Shapes live in
// api/openapi.yaml; only what the spec cannot say belongs here.

/** The account's chosen main, else its first character; null for an account with none. */
val Me.mainCharacter: AccountCharacter?
    get() = characters.firstOrNull { it.key == mainCharacterKey } ?: characters.firstOrNull()
