# Provenance: Database Design, RE, ERD Research

Ngày: 2026-06-18  
Tệp chính: `report/research/database_design/database-design-re-erd-research-2026-06-18.md`

## Skill

- Requested skill: `$deep-research`
- Skill file read: `/home/0xKaBG/.codex/skills/feynman/deep-research/SKILL.md`
- Note: Skill yêu cầu `/deepresearch` workflow, nhưng phiên làm việc hiện không có slash-command runtime riêng. Nghiên cứu được thực hiện trực tiếp bằng web search/open và tổng hợp thủ công theo output format của skill: cited brief + provenance sidecar.

## Scope

Research topic:

> Database design, Requirements Engineering, ERD, focused on design-level modeling rather than implementation details.

In-scope:

- Requirements Engineering as source for data design.
- Conceptual/logical/physical data model distinction.
- ERD components and role in database design.
- Placement of ERD in an OOAD report.
- Application to Asymptotic -- AI Agent Financial Gateway.

Out-of-scope:

- SQL DDL.
- Index tuning.
- Migration strategy.
- Query optimization.
- DBMS-specific physical design.

## Local Sources Consulted

- `report/report_support_documentation/governance/project_source_of_truth.md`
- `report/report_support_documentation/guidelines/diagrams/ooad_diagram_priority_list.md`
- `report/report_support_documentation/governance/OOAD.md`

## Web Sources Consulted

### IEEE/ISO/IEC 29148-2018

- URL: https://standards.ieee.org/ieee/29148/6937/
- Accessed via web open.
- Used for:
  - Requirements Engineering scope.
  - Requirements process/product framing.
  - Good requirement attributes and iterative life-cycle framing.
- Key lines observed:
  - Page identifies the standard as active.
  - Page title: “Systems and software engineering -- Life cycle processes -- Requirements engineering.”
  - Description states that the document covers processes and products related to engineering requirements throughout the life cycle.

### IBM -- Data Modeling

- URL: https://www.ibm.com/think/topics/data-modeling
- Accessed via web open.
- Used for:
  - Definition of data modeling.
  - Conceptual/logical/physical data model distinction.
  - Data modeling process from requirements to entities, attributes, relationships and validation.
- Key lines observed:
  - Data modeling visually represents data points/structures and their relationships.
  - Data models are built around business needs and business rules.
  - Conceptual, logical and physical models differ by abstraction level.

### IBM -- Entity Relationship Diagram

- URL: https://www.ibm.com/think/topics/entity-relationship-diagram
- Accessed via web click/open.
- Used for:
  - ERD definition.
  - ERD components: entity, attribute, relationship, cardinality.
  - Difference between ERD, database schema and data flow diagram.
  - Conceptual/logical/physical ER model distinction.
- Key lines observed:
  - ERD represents how items in a database relate to each other.
  - ERDs convey relationship types between entities.
  - ERDs include entities, attributes and relationships; some include cardinality.

### DBLP -- Chen 1976

- URL: https://dblp.org/rec/journals/tods/Chen76.html
- Accessed via web open.
- Used for:
  - Bibliographic metadata for Peter Chen’s ER model paper.
  - DOI, publication venue, year and pages.
- Key lines observed:
  - Title: “The Entity-Relationship Model - Toward a Unified View of Data.”
  - Author: Peter P. Chen.
  - Venue: ACM Transactions on Database Systems 1(1): 9-36, 1976.
  - DOI: 10.1145/320434.320440.

### ACM DOI

- URL: https://dl.acm.org/doi/10.1145/320434.320440
- Access attempt: returned 403 through browser tool.
- Used only as DOI target metadata, not as a directly read source.

## Synthesis Notes

- IEEE 29148 supports the RE framing: design artifacts should be traceable to requirements and validated as part of requirements/design lifecycle.
- IBM Data Modeling supports separating conceptual/logical/physical models and starting from business requirements.
- IBM ERD supports treating ERD as data structure/relationship modeling rather than process modeling.
- Chen 1976 supports ER model as the theoretical basis for entity-relationship modeling.
- Local OOAD guidance states ERD is not pure OOAD UML, but is useful in software reports and should derive from Entity classes and storage requirements.

## Resulting Recommendation

For this report:

- Keep ERD in Chapter 4, not Chapter 3.
- Use ERD at logical design level.
- Derive ERD from FR/NFR, use case specifications, activity diagrams and Entity classes.
- Split ERD into multiple smaller diagrams if a full system ERD becomes visually dense.
- Avoid implementation details unless they are needed to explain a design constraint.
