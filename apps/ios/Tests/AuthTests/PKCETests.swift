import Foundation
@testable import Auth
import Testing

struct FixedRandomBytes: RandomBytes {
    let byte: UInt8
    func bytes(_ count: Int) throws(AuthError) -> Data { Data(repeating: byte, count: count) }
}

@Suite("PKCE")
struct PKCETests {
    @Test("S256 challenge matches RFC 7636 Appendix B")
    func rfcVector() {
        let pkce = PKCE(verifier: "dBjftJeZ4CVP-mB92K27uhbUJU1p1r_wW1gFWFOEjXk")
        #expect(pkce.challenge == "E9Melhoa2OwvFrEMTJguCHaoeK1t8URWbuGJSstw-cM")
        #expect(PKCE.method == "S256")
    }

    @Test("a generated verifier is 43 base64url characters")
    func verifierShape() throws {
        let pkce = try PKCE.generate()
        #expect(pkce.verifier.count == 43)
        let allowed = CharacterSet(charactersIn: "ABCDEFGHIJKLMNOPQRSTUVWXYZabcdefghijklmnopqrstuvwxyz0123456789-._~")
        #expect(pkce.verifier.unicodeScalars.allSatisfy(allowed.contains))
        #expect(pkce.challenge == PKCE.challenge(for: pkce.verifier))
    }

    @Test("two generated verifiers differ")
    func unique() throws {
        #expect(try PKCE.generate().verifier != PKCE.generate().verifier)
    }

    @Test("generation uses the injected random source")
    func injectedRandom() throws {
        let a = try PKCE.generate(random: FixedRandomBytes(byte: 0xAB))
        let b = try PKCE.generate(random: FixedRandomBytes(byte: 0xAB))
        #expect(a == b)
    }

    @Test("base64url has no padding or +/ characters")
    func base64URL() {
        #expect(Data([0xFB, 0xFF]).base64URLEncoded() == "-_8")
    }
}
