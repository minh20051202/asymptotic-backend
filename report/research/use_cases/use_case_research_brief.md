# Research brief: Use Case diagram, analysis, specification, and writing style

Date: 2026-06-16

## Executive summary

Use Case work for this report should follow a conservative rule: the diagram shows actor goals and system scope, while the specification explains the actor-system conversation in readable, testable prose. UML standardizes Use Case semantics and relationships, but it does not mandate one textual template. Therefore, the project should keep the current `Summary -> Flow -> Additional Information` structure and improve the writing style: active voice, clear subject, one main action per step, observable behavior, and explicit success or failure outcome for each branch.

For the AI Agent Financial Gateway, UC01 remains one Use Case. Authentication, idempotency, budget checking, routing, streaming, settlement, and reconciliation are steps or rules inside UC01, not separate Use Cases unless they become independent actor goals or reusable Use Cases across multiple actor goals.

## Source hierarchy

1. OMG UML 2.5.1 is authoritative for Use Case semantics, Actor semantics, system boundary, `include`, and `extend`.
2. OOAD teaching material is useful for methodology: use cases describe `What, not How`; actor is a role outside the system; avoid very small or overly concrete use cases.
3. Cockburn-style Use Case writing and modern BA practice provide textual specification conventions: main success scenario, extensions, preconditions, postconditions, and readable prose.
4. Empirical work on Use Case quality supports the same writing concerns: inconsistent granularity, vague text, solution-oriented wording, long steps, and missing branch outcomes reduce usefulness.

## Use Case diagram rules

- A Use Case must describe useful functionality that yields an observable result of value to an Actor or stakeholder.
- Actor means an external role, not a specific person and not an internal component.
- A diagram should communicate system scope, actor goals, and external systems. It should not show algorithmic steps.
- `include` points from the including Use Case to the included Use Case. It is mainly useful when common behavior is extracted for reuse.
- `extend` points from the extending Use Case to the extended/base Use Case. The base Use Case must remain meaningful without the extension.
- UML has no `<<precondition>>` relationship between Use Cases. Preconditions belong in the textual specification.
- A detailed Use Case diagram is a narrower view of the same model, not an Activity Diagram drawn with ovals.

## Use Case analysis process

1. Define system boundary.
2. Identify external actors and systems.
3. Build an actor-goal list.
4. Keep Use Cases at user-goal level.
5. Remove subfunctions that are only technical steps.
6. Write one short specification for every selected Use Case.
7. For each important Use Case, write the normal successful flow first.
8. Inspect each step for alternative and exception paths.
9. Move formulas, state names, data structures, retry policy, and transaction details into Business Rules or design sections.
10. Verify that every path can become an acceptance test.

## Specification structure recommended for this project

The current structure is acceptable:

- ID
- Name
- Status
- Description
- Scope
- Primary Actor
- Supporting Actor
- Priority
- Trigger
- Pre-Conditions
- Success Post-Conditions
- Failure Post-Conditions
- Basic Flow
- Alternative Flows
- Exception Flows
- Business Rules
- Non-Functional Requirements
- Traceability

Not every Use Case needs a long "fully dressed" form in the final report. UC01 can be fuller because it is central; UC02-UC09 should be concise and readable.

## Writing style rules

- Use active sentences: `Gateway validates the request`, not `The request is validated`.
- Name the subject of each step: `AI Agent`, `Gateway`, `Payment Provider`, `Organization Admin`.
- Keep one main action or response per step.
- Keep Basic Flow as the ordinary success path. Do not embed long `if/else` logic in the Basic Flow.
- Branch names should start with the Basic Flow step number, such as `4a`, `4b`.
- Every branch must say whether the Use Case continues at a specific step, ends successfully, or ends unsuccessfully.
- Avoid UI details unless the interface action is the actual actor goal.
- Avoid implementation details in flow text: SQL locks, table names, hashes, cache synchronization, internal state names, HTTP status codes, class names, and algorithms.
- Use technical terms only when they are part of the external contract, such as API key or idempotency key.
- Put detailed policies into Business Rules with stable IDs.
- Use the same term consistently. Do not alternate between `Agent`, `AI Agent`, `tác nhân`, and `người dùng API` for the same actor.

## Applying the rules to UC01

UC01 should be written as one actor-goal Use Case: `Thực hiện yêu cầu AI`. It can mention streaming because streaming changes externally visible behavior and settlement risk. It should not turn tokenizer, budget formula, row lock, model fallback algorithm, or ledger update details into Basic Flow steps.

The best shape for UC01:

- Basic Flow: 7-10 steps.
- Alternative Flow: cached/replayed request, fallback allowed, budget recalculation, retry before streaming.
- Exception Flow: invalid request, idempotency conflict, no budget, provider failure after streaming, client disconnect, settlement failure.
- Business Rules: idempotency scope, fallback permission, retry limit, minimum output threshold, reservation calculation, reconciliation rule.

## Applying the rules to UC02-UC09

UC02-UC09 should be shorter than UC01. They should focus on the actor goal:

- UC02: top up organization wallet.
- UC03: allocate or recover team budget.
- UC04: configure AI Agent budget limit.
- UC05: manage access keys.
- UC06: view usage and transactions.
- UC07: register organization.
- UC08: manage developers.
- UC09: manage AI models and pricing policy.

The final report should include concise tables. The full markdown specifications can stay under `report_support_documentation/diagrams/chapter_2/use_cases/`.

## Implications for class diagram decomposition

The analysis class diagram should be explained by responsibility groups rather than as one dense figure:

- Boundary classes: actor/system touchpoints.
- Request execution controls: UC01 coordination.
- Finance controls: reservation, settlement, payment, ledger.
- Administration controls: organization, developer, budget, key, model/pricing policy.
- Core entities: organization, user, developer, agent, key, wallet, policy, model, pricing, transaction, ledger, idempotency, trace, payment.

Breaking the class diagram into these sections is consistent with OOAD because each section maps a set of Use Cases to boundary-control-entity responsibilities.

## Sources used

- Object Management Group, Unified Modeling Language 2.5.1, Clause 18 UseCases: `report/report_support_documentation/references/standards/formal-17-12-05.pdf`.
- Trịnh Thành Trung, `Bài 13 - Tổng quan về UML`, OOAD lecture material: `report/report_support_documentation/references/lectures/OOP_Bai13(vi).pdf`.
- Thinhnotes, "Use Case Diagram và 5 sai lầm thường gặp": https://thinhnotes.com/chuyen-nghe-ba/use-case-diagram-va-5-sai-lam-thuong-gap/
- Thinhnotes, "Viết đặc tả Use Case sao đơn giản nhưng hiệu quả?": https://thinhnotes.com/chuyen-nghe-ba/viet-dac-ta-use-case-sao-don-gian-nhung-hieu-qua/
- Alistair Cockburn, `Writing Effective Use Cases`, cited for goal levels and writing conventions.
- Martin Fowler, `Use Case`, cited for lightweight, readable Use Case practice.
- Barros-Justo et al., empirical work on Use Case use in real-world projects, cited for granularity and template consistency concerns.
- Seki et al., empirical work on bad smells in Use Case descriptions, cited for writing quality concerns.
