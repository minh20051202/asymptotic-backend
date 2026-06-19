# Provenance — Asymptotic competitor documentation deep research

**Research artifact:** `report/research/market/asymptotic-competitor-docs-deep-research-2026-06-18.md`  
**Research date:** 2026-06-18  
**Language:** Vietnamese  
**Source policy:** Official product documentation, API references, official repositories, pricing pages and vendor announcements only.

## 1. Workflow

The named `deep-research` skill was executed using three independent roles:

1. **Researcher:** investigated Cloudflare AI Gateway, Portkey, LiteLLM and Helicone documentation.
2. **Verifier:** investigated Skyfire, Nevermined, x402, Stripe Machine Payments/MPP, Google AP2, Mastercard Agent Pay and Visa Trusted Agent Protocol.
3. **Reviewer:** audited methodology, market positioning and consistency with the Asymptotic source of truth.

The main agent synthesized the outputs and constrained conclusions to:

- `report/report_support_documentation/governance/project_source_of_truth.md`;
- `report/research/market/asymptotic-market-research-synthesis-2026-06-18.md`;
- official sources listed in the final brief.

## 2. Scope controls

### In scope

- AI Gateway capabilities.
- Managed keys and provider credential isolation.
- Organization/team/user/Agent entities.
- Rate, quota and monetary budget enforcement.
- Usage and cost attribution.
- Pre-provider checks.
- Prepaid credits.
- Reservation, reconciliation, idempotency and ledger evidence.
- Adjacent Agent identity, mandate and machine-payment patterns.

### Out of scope

- Independent market-share estimates.
- Claims of production scale not supported by official evidence.
- Legal conclusions.
- Full payment-gateway architecture for Asymptotic.
- Agent creation, hosting or orchestration.

## 3. Evidence-status rules

| Code | Meaning |
|---|---|
| F | Fully matched to the operational definition |
| P | Partial or analogous capability |
| X | Explicitly unsupported |
| U | Undetermined from reviewed public documentation |
| NA | Not applicable |

| Grade | Evidence type |
|---|---|
| E4 | Official technical documentation or API reference |
| E3 | Official beta/preview documentation |
| E2 | Official repository, SDK or sample |
| E1 | Vendor landing page, blog or press release |
| E0 | No adequate public evidence found |

`U/E0` is not evidence of absence.

## 4. Claim provenance

| Claim group | Primary evidence |
|---|---|
| Cloudflare pre-provider spend checks and eventual consistency | CF02 |
| Cloudflare prepaid credits, top-up and auto-top-up | CF03 |
| Cloudflare provider-key storage | CF04 |
| Portkey multi-level budgets and pre-provider enforcement | PK03–PK06 |
| Portkey organization/workspace hierarchy | PK07–PK09 |
| Portkey Agent Registry | PK10 |
| LiteLLM organization/team/user/key hierarchy and budgets | LL02, LL03, LL07 |
| LiteLLM external Agent registration and Agent trace | LL04 |
| LiteLLM Agent/session budget behavior | LL05 |
| LiteLLM per-call spend records | LL06 |
| Helicone current AI Gateway and provider routing | HC01, HC02 |
| Helicone cost-based blocking | HC03 |
| Helicone usage/cost/session observability | HC04, HC07 |
| Skyfire Agent Account, KYA and PAY tokens | SF01–SF04 |
| Nevermined metering and Agent payments | NV01, NV02 |
| x402 payment challenge, settlement and idempotency | X401–X404 |
| Stripe machine-payment lifecycle | ST01, ST02 |
| AP2 mandate model and maturity caveat | AP01, AP02 |
| Mastercard registered Agents and Agentic Tokens | MC01, MC02 |
| Visa intent signatures, anti-replay and deployment status | VS01, VS02 |

Source identifiers resolve to the URLs in Section 14 of the research artifact.

## 5. Corrections to prior synthesis

The deep research supersedes the following earlier assumptions:

1. Cloudflare does support monetary spend limits, pre-provider blocking and prepaid credits.
2. Portkey does provide an Agent Registry and explicit pre-provider budget checks.
3. LiteLLM provides external Agent registration, Agent-linked keys and Agent/session budgets.
4. Helicone is a direct AI Gateway competitor, not only an observability product.
5. Wallet/top-up is not unique because Cloudflare and Helicone provide prepaid credit flows.
6. External Agent registration alone is no longer a defensible differentiator.

## 6. Reviewer findings retained

- Direct AI Gateway products must be compared separately from payment infrastructure.
- Capability criteria must be operationalized.
- “Not found” must not be written as “does not exist.”
- Asymptotic’s hierarchy is a domain model, not automatically a market innovation.
- Differentiation should be expressed as a design focus around consistent per-request financial control.
- Azure API Management AI Gateway, OpenRouter and Kong AI Gateway should be screened in a subsequent research round.

## 7. Known limitations

- The research did not use paid enterprise accounts.
- Product documentation can change after 2026-06-18.
- Vendor documentation does not independently establish adoption, reliability or transaction volume.
- Some maturity classifications remain mixed because individual features can be GA, beta or preview within the same product.
- The listed direct competitors are those named in the input synthesis; the market set is not exhaustive.
