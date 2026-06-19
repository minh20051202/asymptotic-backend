# Provenance — Chapter 2 deep competitor documentation research

**Artifact:** `report/research/market/chapter-2-deep-competitor-docs-research-2026-06-19.md`  
**Research date:** 2026-06-19  
**Purpose:** Evidence base for Chapter 2, Section 2.1.

## Method

The supplied Gemini deep-research skill could not execute because:

- `/home/0xKaBG/.codex/skills/deep-research/` contained only `SKILL.md`;
- `scripts/research.py` was absent;
- `GEMINI_API_KEY` was unset.

The task continued through an equivalent official-source workflow:

1. Read the Asymptotic source of truth and previous research.
2. Traverse product documentation indexes, `llms.txt`, Enterprise pages and linked subpages.
3. Separate explicit capability, analogous capability and undetermined capability.
4. Inspect architecture and database documentation rather than relying only on feature pages.
5. Add missing direct competitors identified during review.
6. Constrain market-gap claims to public documentation.

## Product branches inspected

### LiteLLM

- Enterprise.
- Multi-tenant architecture.
- Project management.
- Budgets and rate limits.
- A2A Agent Gateway.
- Agent permission management.
- Agent iteration budgets.
- Spend tracking.
- Database schema overview.
- Billing with Lago.
- Pricing calculator.
- Request lifecycle.
- High-availability spend update path.

### Portkey

- Organizations.
- Workspaces.
- Integrations and credential provisioning.
- Five-level budget/rate enforcement.
- Agent Registry.
- Pricing and cost management.

### Cloudflare

- Unified Billing.
- Spend limits.
- Authenticated Gateway.
- BYOK.
- Custom metadata.
- Logging and routing.

### Helicone

- AI Gateway.
- Provider routing.
- Custom cost-based limits.
- Cost tracking.
- Managed credits and error behavior.

### Additional competitors

- OpenRouter organizations, workspaces, key limits, guardrails and BYOK.
- Azure API Management AI Gateway and token-limit policy.
- Kong AI Gateway, AI cost limiting and Metering & Billing.

## Important evidence corrections

1. LiteLLM explicitly supports hierarchical allocation and enforcement.
2. LiteLLM Project Management adds `Organization → Team → Project → Key`.
3. Portkey explicitly supports Workspace allocations below Integration ceilings.
4. OpenRouter supports shared organization credits and key/member guardrails.
5. Cloudflare and Helicone support prepaid credits.
6. Kong supports deduplication in its Metering & Billing event layer.
7. External Agent registration is supported by multiple competitors.

## Report-framing correction

The sample report at
`report/report_support_documentation/references/sample_reports/DOAN1-NGUYENHUYENTRANG.pdf`
was inspected, especially its current-state survey and proposed-system sections.
Its pattern is:

1. inspect named existing products;
2. identify strengths and shortcomings;
3. derive expected functions;
4. reuse suitable existing functions in the proposed system.

Accordingly, the research artifact does not require Asymptotic to demonstrate
absolute feature novelty. Competitor capabilities are classified for:

- direct reuse;
- adaptation to Asymptotic actors and business flow;
- exclusion from the MVP;
- custom design where required.

## Inference boundary

The report infers that LiteLLM does not document strict runtime reservation because:

- budget is checked before provider execution;
- spend update occurs in asynchronous post-request tasks;
- database transactions are not tied to request lifecycle;
- spend updates can be buffered and flushed later.

This does not prove LiteLLM cannot prevent all concurrent overspend through undocumented mechanisms. The report therefore marks strict reservation as undetermined, not absent.

## Scope boundary

The research does not treat:

- token quota as monetary wallet allocation;
- request logs as a financial ledger;
- metadata segmentation as a managed organization hierarchy;
- cost estimation tools as runtime reservation;
- Agent registration as proof of Agent financial ownership;
- vendor documentation as independent proof of production adoption.
