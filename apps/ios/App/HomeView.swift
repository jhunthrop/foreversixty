import DesignSystem
import Models
import Networking
import SwiftUI

/// What the placeholder home shows. Lane F2 replaces it with the spec's home answer.
enum HomeState: Equatable {
    case loading
    case loaded(HomeCard)
    case empty
    case failed(String)
}

/// The main character's headline, already formatted.
struct HomeCard: Equatable {
    let name: String
    let detail: String
    let guild: String?

    init(name: String, detail: String, guild: String?) {
        self.name = name
        self.detail = detail
        self.guild = guild
    }

    init(_ character: AccountCharacter) {
        let level = character.level.map { "Level \($0)" }
        let kind = [character.spec, character.className].compactMap { $0?.capitalized }.joined(separator: " ")
        self.init(
            name: character.name,
            detail: [level, kind.isEmpty ? nil : kind].compactMap(\.self).joined(separator: " · "),
            guild: character.guild?.name
        )
    }
}

/// Loads `/v1/me` and hands the result to `HomeView`.
struct HomeScreen: View {
    let client: APIClient
    @State private var state: HomeState = .loading

    var body: some View {
        HomeView(state: state)
            .task { state = await load() }
    }

    private func load() async -> HomeState {
        do {
            let me = try await client.send(.getMe)
            return me.mainCharacter.map { .loaded(HomeCard($0)) } ?? .empty
        } catch {
            return .failed(message(for: error))
        }
    }

    private func message(for error: APIError) -> String {
        switch error {
        case let .server(_, body, _): body.message
        case .unauthorized: "Sign in to see your character."
        default: "Could not load your character just now."
        }
    }
}

/// Pure rendering of a `HomeState`; the snapshot test renders this directly.
struct HomeView: View {
    let state: HomeState

    var body: some View {
        NavigationStack {
            ScrollView {
                content
                    .padding(Theme.Spacing.md)
            }
            .background(Theme.Palette.background)
            .navigationTitle("Forever Sixty")
        }
    }

    @ViewBuilder
    private var content: some View {
        switch state {
        case .loading:
            ProgressView().frame(maxWidth: .infinity)
        case let .loaded(card):
            VStack(alignment: .leading, spacing: Theme.Spacing.xs) {
                Text(card.name).font(Theme.Typography.title).foregroundStyle(Theme.Palette.strong)
                Text(card.detail).font(Theme.Typography.body).foregroundStyle(Theme.Palette.text)
                if let guild = card.guild {
                    Text(guild).font(Theme.Typography.meta).foregroundStyle(Theme.Palette.accent)
                }
            }
            .cardSurface()
        case .empty:
            Text("Link a character with the addon or Battle.net to see it here.")
                .font(Theme.Typography.body)
                .foregroundStyle(Theme.Palette.muted)
                .cardSurface()
        case let .failed(message):
            Text(message).font(Theme.Typography.body).foregroundStyle(Theme.Palette.muted).cardSurface()
        }
    }
}

#Preview("Loaded") {
    HomeView(state: .loaded(HomeCard(name: "Obnoxious", detail: "Level 42 · Arms Warrior", guild: "Fixture Guild")))
        .preferredColorScheme(.dark)
}
