# Provenance: Cách Viết FR Và NFR Trong OOAD

**Brief:** `report/research/requirements/ooad-fr-nfr-writing-research-2026-06-16.md`  
**Research cutoff:** 2026-06-16  
**Source grades:** A = official standard/government/standards body; B = recognized practitioner framework or official author/publisher page; C = academic/preprint; D = secondary summary.

## Method

The research prioritized authoritative requirements-engineering and software-quality sources. The brief applies them to OOAD reporting and to the local AI Agent Financial Gateway project, but it does not claim to reproduce paid ISO/IEEE standard text beyond public summaries and accessible excerpts. NASA and SEI public reports were used for practical checklists and quality-attribute scenario guidance.

## Source Ledger

| ID | Grade | Source | Claim used | Caveat |
|---|---:|---|---|---|
| S01 | A | ISO/IEC/IEEE 29148:2018, https://www.iso.org/standard/72089.html | Current central standard for requirements engineering; specifies processes, information items, contents, and format guidance for requirements across systems/software life cycles. | Full standard is paywalled; public abstract used. |
| S02 | A | IEEE SA page for IEEE/ISO/IEC 29148-2011, https://standards.ieee.org/ieee/29148/5289/ | 29148 defines the construct of a good requirement, attributes/characteristics, iterative requirements process; replaced IEEE 830/1233/1362 in 2011. | Page is for superseded 2011 standard but useful for public summary; 2018 is current from ISO. |
| S03 | A | NASA Systems Engineering Handbook, Appendix C, https://www.nasa.gov/wp-content/uploads/2018/09/nasa_systems_engineering_handbook_0.pdf | Practical checklist: `shall` as requirement term, clarity, one thought per requirement, completeness, correct level, avoid implementation specifics, consistency, traceability, verifiability/testability. | Systems engineering source, not OOAD-specific; still directly applicable to software requirements. |
| S04 | A | OMG UML 2.5.1, https://www.omg.org/spec/UML/2.5.1/About-UML | UML is a formal graphical language for visualizing/specifying/constructing/documenting artifacts; use in OOAD should be treated as modeling, not a substitute for requirement statements. | Public page does not expose all use-case semantics in-line; full PDF is normative. |
| S05 | A | ISO/IEC 25010:2023, https://www.iso.org/standard/78176.html | Product quality model is appropriate for quality requirements, quality criteria, acceptance criteria, and measures. | Full standard is paywalled; public summary used. |
| S06 | B | ISO 25000 portal summary, https://iso25000.com/index.php/en/iso-25000-standards/iso-25010 | Quality characteristics and subcharacteristics used to classify NFR categories. | Portal summary, not the official ISO standard text. |
| S07 | A | SEI/CMU QAW, https://insights.sei.cmu.edu/library/quality-attribute-workshops-qaws-third-edition/ | Quality Attribute Workshop identifies and refines important quality attributes early, before architecture is created. Supports scenario-based NFR writing. | QAW is a method/report, not a requirements standard. |
| S08 | B | Volere Requirements Specification Template, https://www.volere.org/templates/volere-requirements-specification-template/ | Supports separating goals, stakeholders, constraints, assumptions, business rules, product use cases, FR, NFR, and fit criteria. | Practitioner framework, not formal standard. |
| S09 | B | Software Requirements Essentials, https://softwarereqs.com/ | Practitioner background source from Karl Wiegers/Joy Beatty ecosystem; supports the general recommendation to use templates, reviews, use cases, and quality-attribute prioritization. | Background only; not used as primary evidence for formal claims. |
| S10 | B | Karl Wiegers books page, https://www.karlwiegers.com/books.html | Bibliographic confirmation of Software Requirements and Software Requirements Essentials as recognized practitioner references. | Background only; not used as primary evidence for formal claims. |
| S11 | B | Business Rules Group, https://www.businessrulesgroup.org/first_paper/BRG-whatisBR_3ed.pdf | Business rules identify and articulate rules that define structure and control operation of an enterprise; useful for distinguishing domain policy from system FR. | Practitioner/industry report, not a formal ISO/IEEE standard. |
| S12 | C | Eckhardt et al., arXiv 1611.08868, https://arxiv.org/abs/1611.08868 | NFRs are often documented separately, vague, and lacking quantitative measures; many NFRs still describe system behavior and should be analyzed/tested rigorously. | Preprint; use as cautionary evidence. |
| S13 | C | Behutiye et al., arXiv 1711.08894, https://arxiv.org/abs/1711.08894 | NFR documentation is often neglected; acceptance criteria/backlogs/DoD are used in practice to document NFRs. | Preprint; agile context, not OOAD-specific. |
| S14 | A | OMG SBVR 1.5, https://www.omg.org/spec/SBVR/ | SBVR defines vocabulary and rules for documenting semantics of business vocabularies and business rules for exchange among organizations/tools. | Business-rule standard, not a general SRS template. |

## Verification Notes

- The distinction “FR = what the system does, NFR = quality/constraint under which it operates” is useful pedagogically, but not absolute. Several NFRs can still be expressed as observable behavior and should remain testable.
- Use Cases and FRs overlap but are not identical. Use Cases structure actor-goal scenarios; FRs are atomic, traceable requirement statements derived from those scenarios and other sources.
- NFRs should not be accepted as vague adjectives. They need condition, metric, threshold, and verification method.
- OOAD reports benefit from explicit traceability: Business Goal -> Business Rule -> Use Case -> FR/NFR -> Design -> Test.
