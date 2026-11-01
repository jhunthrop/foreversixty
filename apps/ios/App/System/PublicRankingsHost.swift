import Auth
import SwiftUI

/// Lane F15 follow-up (F11): Rankings and character pages are public data, so a rankings or
/// player link opens for a signed-out player and for a player with no character. While Home's
/// sign-in gate holds, this stack carries the gate and pushes what the link asks for; with no
/// character, the board simply has no "your rank".
struct PublicRankingsHost<Gate: View>: View {
    let dependencies: RankingsDependencies?
    let account: RankingsAccount
    let router: RankingsRouter
    @ViewBuilder let gate: () -> Gate
    @State private var watchList: WatchList?

    var body: some View {
        NavigationStack {
            gate()
                .navigationDestination(isPresented: showsBoard) { board }
                .navigationDestination(item: player) { page($0) }
        }
    }

    private var showsBoard: Binding<Bool> {
        Binding(get: { router.showsBoard }, set: { router.showsBoard = $0 })
    }

    private var player: Binding<PlayerRoute?> {
        Binding(get: { router.player }, set: { router.player = $0 })
    }

    @ViewBuilder
    private var board: some View {
        if let dependencies {
            RankingsScreen(dependencies: dependencies, account: account, encounter: router.encounter).id(router.encounter)
        }
    }

    @ViewBuilder
    private func page(_ route: PlayerRoute) -> some View {
        if let dependencies {
            let list = watchList ?? WatchList(service: dependencies.service, signedIn: account.signedIn)
            CharacterScreen(route: route, dependencies: dependencies, account: account, watchList: list)
                .onAppear { if watchList == nil { watchList = list } }
        }
    }
}

extension SessionStore {
    /// Home's gate holds: signed out, signing in, importing, or signed in with no character.
    var homeGateHolds: Bool {
        guard case let .signedIn(me) = phase else { return true }
        return me?.characters.isEmpty == true
    }
}
