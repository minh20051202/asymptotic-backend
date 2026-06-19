# Provenance: AI Agent Financial Gateway Research

**Brief:** `report/research/market/ai-agent-financial-gateway-research-2026-06-16.md`  
**Research cutoff:** June 16, 2026  
**Source grades:** A = law/regulator/standard/official technical spec; B = official vendor documentation or announcement; C = reputable secondary reporting; D = preprint or unaudited research.

## Method

The research focused on systems that let AI agents spend, receive, authorize, settle, or reconcile money. Generic AI finance products were excluded unless they exposed payment, wallet, mandate, credential, or settlement infrastructure. The survey emphasized public English-language sources, US/EU legal and regulatory anchors, and globally visible payment/commerce protocols. It is not an exhaustive jurisdiction-by-jurisdiction legal survey or a guarantee that no private deployment exists.

Claims were separated into:

- Deployed: public APIs, live checkout, live documentation, or explicit live availability.
- Limited: preview, pilot, constrained rollout, sandbox, or unclear access.
- Announced: public plan without evidence of broad availability.
- Standard: specification or reference implementation, not itself a regulated financial service.

Vendor pages were treated as self-reported. Legal and regulatory conclusions were anchored to primary law, regulator, or standards sources where available.

## Source Ledger

| ID | Grade | Source | Claim used | Caveat |
|---|---:|---|---|---|
| S01 | B | Mastercard Agent Pay press release, Apr. 29, 2025, https://www.mastercard.com/news/press/2025/april/mastercard-unveils-agent-pay-pioneering-agentic-payments-technology-to-power-commerce-in-the-age-of-ai/ | Mastercard announced registered agent payments, tokenization, permissions, and ecosystem visibility. | Announcement/pilot evidence; no broad volume disclosed. |
| S02 | A/B | Google AP2 repository, https://github.com/google-agentic-commerce/AP2 | AP2 provides signed mandate artifacts for agent payment authorization. | Standard/reference implementation; not a regulated rail. |
| S03 | B | OpenAI Instant Checkout, Sep. 29, 2025, https://openai.com/index/buy-it-in-chatgpt/ | US ChatGPT users could purchase from Etsy sellers with user confirmation. | Human-present checkout, not full autonomous spending. |
| S04 | A/B | Agentic Commerce Protocol repository, https://github.com/agentic-commerce-protocol/agentic-commerce-protocol | ACP defines commerce/checkout/order/payment-handler flows and is beta with production-linked implementations. | Beta standard; ecosystem adoption still emerging. |
| S05 | B | Stripe Agentic Commerce Suite, Dec. 11, 2025, https://stripe.com/newsroom/news/agentic-commerce-suite | Stripe supports merchants selling through agents and Shared Payment Tokens. | Vendor announcement; customer rollout breadth not independently audited. |
| S06 | A/B | UCP repository and Google developer post, https://github.com/Universal-Commerce-Protocol/ucp and https://developers.googleblog.com/under-the-hood-universal-commerce-protocol-ucp/ | UCP spans discovery, checkout, order lifecycle, identity linking, payment-token exchange, and multiple transports. | Protocol, not licensing/settlement/compliance. |
| S07 | A | A2A specification, https://a2a-protocol.org/latest/specification/ | A2A handles agent discovery and interoperability; auth/security schemes are externalized. | Not a financial authorization standard. |
| S08 | B | Stripe Machine Payments docs, https://docs.stripe.com/payments/machine.md | Stripe documents programmatic agent payments over stablecoin and payment-token rails. | Live/preview status varies by rail and region. |
| S09 | B | Stripe MPP docs, https://docs.stripe.com/payments/machine/mpp.md | MPP uses HTTP 402 challenge/authorization/retry/receipt flow; supports crypto and SPT/card patterns. | API preview/version details can change. |
| S10 | B | Coinbase AgentKit, https://docs.cdp.coinbase.com/agent-kit/welcome | AgentKit exposes wallets and on-chain actions for agents. | Crypto/on-chain enabler, not a complete financial gateway. |
| S11 | B | Crossmint wallet infrastructure, https://www.crossmint.com/products/wallet-infrastructure | Crossmint claims programmable wallets for users, agents, and backend workflows with policy/compliance features. | Vendor assertion; exact licensing/custody facts require diligence. |
| S12 | B | Privy server wallets, https://docs.privy.io/guide/server-wallets/ | Managed wallet fleets can support policies, approvals, and idempotent transactions. | Enabler, not agent-specific regulated gateway. |
| S13 | B | Payman AI, https://paymanai.com/ | Payman claims controlled banking/payment execution under policies and audit trails. | Limited public deployment detail; volume undisclosed. |
| S14 | B | Link agents and skill file, https://link.com/agents and https://link.com/skill.md | Link gives agents one-time cards or SPT-like credentials under user approval and spend limits. | Current model emphasizes explicit approval; granular autonomous controls are still developing. |
| S15 | A | W3C VC Data Model 2.0, May 15, 2025, https://www.w3.org/TR/vc-data-model-2.0/ | Verifiable credentials carry issuer claims; verifiers decide fitness, proof, and status. | Credentials do not equal legal authority. |
| S16 | A | ERC-8004 draft, https://eips.ethereum.org/EIPS/eip-8004 | Draft defines agent identity, reputation, validation registries, and wallet pointer. | Payments explicitly out of scope. |
| S17 | A | RFC 9700, Jan. 2025, https://www.rfc-editor.org/rfc/rfc9700.html | OAuth BCP favors sender-constrained tokens, least privilege, audience restriction, and replay defenses. | API security standard, not legal mandate. |
| S18 | A | OpenID FAPI 2.0 and RFC 9396 RAR, https://openid.net/specs/fapi-security-profile-2_0-final.html and https://www.rfc-editor.org/rfc/rfc9396.html | High-security delegated API access and rich authorization request structures. | Does not itself perform KYC, suitability, settlement, or licensing. |
| S19 | A | NIST SP 800-162 ABAC, https://csrc.nist.gov/pubs/sp/800/162/upd2/final | Attribute-based access control supports deterministic policy enforcement. | Must be implemented with domain-specific financial attributes. |
| S20 | A | NIST SP 800-207 Zero Trust, https://csrc.nist.gov/pubs/sp/800/207/final | Continuous, contextual policy enforcement is appropriate for high-value systems. | Not finance-specific. |
| S21 | A | NIST FIPS 186-5, https://csrc.nist.gov/pubs/fips/186-5/final | Digital signatures prove key-based signing and integrity, not full legal consent. | Legal attribution remains separate. |
| S22 | A | NIST FIPS 140-3, https://csrc.nist.gov/pubs/fips/140-3/final | Cryptographic modules protect keys but do not validate transaction semantics. | HSM/MPC is necessary but insufficient. |
| S23 | A | OFAC Framework, https://ofac.treasury.gov/media/16331/download?inline | Sanctions programs require risk-based, current screening and controls. | Implementation details depend on business and jurisdiction. |
| S24 | A | E-SIGN Act, https://uscode.house.gov/view.xhtml?path=/prelim@title15/chapter96&edition=prelim | Electronic agents can participate in electronic contracts, but actions are attributable to persons/entities. | Does not create AI legal personhood. |
| S25 | A | FinCEN CVC guidance FIN-2019-G001, https://www.fincen.gov/sites/default/files/2019-05/FinCEN%20Guidance%20CVC%20FINAL%20508.pdf | Money-transmitter/custody analysis is functional; hosted wallets generally trigger obligations. | State and other federal regimes may add obligations. |
| S26 | A | GENIUS Act, Pub. L. 119-27, https://www.govinfo.gov/content/pkg/PLAW-119publ27/html/PLAW-119publ27.htm | Stablecoin issuer framework and effective-date trigger; default statutory date is 18 months after enactment absent earlier final-reg trigger. | Must check final implementing regulations before launch. |
| S27 | A | PCI DSS v4.0.1 library, https://www.pcisecuritystandards.org/document_library/?category=pcidss&document=pci_dss | PCI scope follows storage, processing, or transmission of cardholder data. | Tokenization reduces scope but may not eliminate service-provider duties. |
| S28 | B | Visa Trusted Agent Protocol, https://developer.visa.com/capabilities/trusted-agent-protocol | Visa describes time-bound, purpose/merchant-specific agent signatures and intent metadata. | Development/deployment language; no universal GA proof. |
| S29 | B | KYA/KYAPay, https://kyapay.org/ | KYA binds agent/developer/platform/user identity context; KYAPay adds intent and payment token ideas. | Associated with vendor ecosystem, not neutral legal KYC replacement. |
| S30 | B | x402, https://www.x402.org/ | HTTP 402 payment flow for stablecoin/machine payments with public live metrics. | Metrics are not all provably autonomous AI-agent transactions. |
| S31 | B | Circle Agent Stack, https://developers.circle.com/agent-stack | Circle provides agent wallet, USDC, policy, and nanopayment tooling. | Vendor docs; production usage distribution not independently audited. |
| S32 | B | Circle Gateway, https://developers.circle.com/gateway | Circle describes unified USDC balances and fast cross-chain routing. | Stablecoin settlement still has issuer/liquidity/compliance risk. |
| S33 | B | Skyfire, https://skyfire.xyz/ | Skyfire claims KYA, agent wallets, cards, stablecoins, mandates, and checkout. | No audited public volume found. |
| S34 | B | Nevermined, https://nevermined.ai/ | Nevermined claims live agent payments, virtual cards, metering, and Stripe settlement. | Vendor page; independent volume not found. |
| S35 | B | Catena, https://catena.com/ | Catena claims agent accounts, deterministic policies, audit trails, cards, ACH, wires, and stablecoins. | Early market signal; volume undisclosed. |
| S36 | D | "Five Attacks on x402," arXiv, Jan. 2026, https://arxiv.org/abs/2601.19850 | Emerging attack patterns include replay, race, and quote-substitution risks around payment flows. | Preprint; use as risk signal, not settled fact. |
| S37 | D | AP2 red-team preprint, arXiv, Jun. 2026, https://arxiv.org/abs/2606.02538 | Mandate protocols still need adversarial testing around prompt/tool/payment attacks. | Preprint; not peer-reviewed. |
| S38 | A | CFPB Regulation E, https://www.consumerfinance.gov/rules-policy/regulations/1005/2/ | Consumer unauthorized EFT rules depend on authorization/access-device facts. | Agent delegation complicates but does not erase Reg E analysis. |
| S39 | A | UCC 4A-202, https://www.law.cornell.edu/ucc/4A/4A-202 | Commercial payment orders depend on authorization and commercially reasonable security procedures. | State enactments and contracts matter. |
| S40 | A | EU AI Act, https://eur-lex.europa.eu/eli/reg/2024/1689/oj/eng | AI Act regulates AI systems by risk/role but does not create a financial-agent license. | Sector-specific rules remain. |
| S41 | A | MiCA, https://eur-lex.europa.eu/eli/reg/2023/1114/oj/eng | EU crypto-asset services such as custody/exchange/transfer generally require authorization. | Financial instruments may fall under MiFID instead. |

## Verification Notes

- The strongest evidence for actual autonomous payment execution is stablecoin/API payment infrastructure, especially x402-style flows. However, public x402 metrics do not classify traffic by autonomous AI agent versus other software clients.
- The strongest evidence for mainstream commerce adoption is human-confirmed checkout, especially OpenAI/Stripe ACP and ChatGPT Instant Checkout. This is commercially important but does not prove broad autonomous spending.
- Card networks are moving toward agent visibility, tokenization, and constrained credentials, but public sources still describe much of this as pilot, rollout, development, or partnership activity.
- Identity protocols should be treated as operational identity. They supplement, but do not replace, principal KYC/KYB, beneficial-owner diligence, or legal authority.
- The report intentionally avoids market-size estimates because current public evidence does not support a reliable AI-agent-specific payments TAM or transaction-volume figure.
- Receiving, settlement, and accounting were evaluated mostly through what the reviewed vendor/protocol sources expose publicly. The absence finding is limited to those sources: ACP, Stripe, x402/MPP, Circle, Skyfire, Nevermined, Catena, Payman, and wallet-infrastructure providers show checkout, payment proof, receipt, log, dashboard, or settlement components, but not a complete public autonomous seller onboarding, merchant-of-record, chargeback, tax, GL-close, and audited reconciliation stack.
