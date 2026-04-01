---
title: Cloud Native Security Controls (CNSC) to NIST SP 800-53 Rev. 5 Mapping
sidebar_position: 7
toc_max_heading_level: 3
---

<!--
Auto-generated from cnsc-nist-800-53-mapping.yaml.
Do not edit manually.

  go run cmd/catalog/main.go -md index.md
-->

# Cloud Native Security Controls (CNSC) to NIST SP 800-53 Rev. 5 Mapping

This document maps the [Cloud Native Security Controls Catalog](./index.md) to NIST 800-53 SP Rev.5 Guidelines.

## Cross-walked Controls

For NIST alignments by catalog family, see the [Cloud Native Security Controls Catalog](./cnsc-nist-800-53-by-family) page.

| CNSC guideline | Relationship | NIST control(s) | Remarks |
|----------------|--------------|-----------------|---------|
| <a id="mapping-row-cnsc-1-ia-5-7"></a>`CNSC-1` | relates-to | `IA-5(7)` | Authenticator Management |
| <a id="mapping-row-cnsc-10-ac-3-13"></a>`CNSC-10` | relates-to | `AC-3(13)` | Access Enforcement |
| <a id="mapping-row-cnsc-100-ia-3-1"></a>`CNSC-100` | relates-to | `IA-3(1)` | Cryptographic Bidirectional Authentication |
| <a id="mapping-row-cnsc-101-si-7"></a>`CNSC-101` | relates-to | `SI-7` | Software, Firmware, and Information Integrity |
| <a id="mapping-row-cnsc-102-si-7"></a>`CNSC-102` | relates-to | `SI-7` | Software, Firmware, and Information Integrity |
| <a id="mapping-row-cnsc-103-si-7"></a>`CNSC-103` | relates-to | `SI-7` | Software, Firmware, and Information Integrity |
| <a id="mapping-row-cnsc-104-si-7"></a>`CNSC-104` | relates-to | `SI-7` | System and Information Integrity |
| <a id="mapping-row-cnsc-105-ra-5"></a><a id="mapping-row-cnsc-105-sa-3"></a>`CNSC-105` | relates-to | `RA-5`<br />`SA-3` | <ul><li>Vulnerability Monitoring and Scanning</li><li>System Development Life Cycle</li></ul> |
| <a id="mapping-row-cnsc-106-si-7"></a>`CNSC-106` | relates-to | `SI-7` | System and Information Integrity |
| <a id="mapping-row-cnsc-107-si-2-3"></a>`CNSC-107` | relates-to | `SI-2(3)` | System and Information Integrity |
| <a id="mapping-row-cnsc-108-sc-12-3"></a>`CNSC-108` | relates-to | `SC-12(3)` | Systems and Communication Protection |
| <a id="mapping-row-cnsc-109-si-2-3"></a>`CNSC-109` | relates-to | `SI-2(3)` | System and Information Integrity |
| <a id="mapping-row-cnsc-11-sc-7-19"></a>`CNSC-11` | relates-to | `SC-7(19)` | Boundary Protection |
| <a id="mapping-row-cnsc-110-pl-1"></a>`CNSC-110` | relates-to | `PL-1` | Policy and Procedures |
| <a id="mapping-row-cnsc-112-ac-6-3"></a>`CNSC-112` | relates-to | `AC-6(3)` | Least Privilege |
| <a id="mapping-row-cnsc-113-sc-12-2"></a><a id="mapping-row-cnsc-113-sc-12-3"></a>`CNSC-113` | relates-to | `SC-12(2)`<br />`SC-12(3)` | <ul><li>Cryptographic Key Establishment and Management \| Symmetric and Asymmetric Keys</li><li>Cryptographic Key Establishment and Management \| Symmetric and Asymmetric Keys</li></ul> |
| <a id="mapping-row-cnsc-114-si-2-3"></a>`CNSC-114` | relates-to | `SI-2(3)` | System and Information Integrity |
| <a id="mapping-row-cnsc-115-ac-6-3"></a>`CNSC-115` | relates-to | `AC-6(3)` | Least Privilege, Network Access to Privileged Commands |
| <a id="mapping-row-cnsc-116-ca-7"></a><a id="mapping-row-cnsc-116-ir-4"></a><a id="mapping-row-cnsc-116-ir-4-5"></a>`CNSC-116` | relates-to | `CA-7`<br />`IR-4`<br />`IR-4(5)` | <ul><li>Continuous Monitoring</li><li>Incident Handling, Automated Incident Handling Processes</li><li>Incident Handling, Automated Disabling of System</li></ul> |
| <a id="mapping-row-cnsc-117-ir-5-1"></a>`CNSC-117` | relates-to | `IR-5(1)` | Incident Monitoring, Automated Tracking, Data Collection, and Analysis |
| <a id="mapping-row-cnsc-12-ia-7"></a>`CNSC-12` | relates-to | `IA-7` | Cryptographic Module Authentication |
| <a id="mapping-row-cnsc-120-ac-3-3"></a>`CNSC-120` | relates-to | `AC-3(3)` | Access Enforcement, Mandatory Access Control |
| <a id="mapping-row-cnsc-121-sa-11-2"></a>`CNSC-121` | relates-to | `SA-11(2)` | Developer Testing and Evaluation, Threat Modeling and Vulnerability Analyses |
| <a id="mapping-row-cnsc-122-ia-9"></a>`CNSC-122` | relates-to | `IA-9` | Service Identification and Authentication |
| <a id="mapping-row-cnsc-123-ia-9"></a>`CNSC-123` | relates-to | `IA-9` | Service Identification and Authentication |
| <a id="mapping-row-cnsc-124-cm-14"></a>`CNSC-124` | relates-to | `CM-14` | Signed Components |
| <a id="mapping-row-cnsc-125-si-6"></a>`CNSC-125` | relates-to | `SI-6` | Security and Privacy Function Verification |
| <a id="mapping-row-cnsc-127-sc-8"></a>`CNSC-127` | relates-to | `SC-8` | Transmission Confidentiality and Integrity |
| <a id="mapping-row-cnsc-128-si-13"></a>`CNSC-128` | relates-to | `SI-13` | Predictable Failure Prevention |
| <a id="mapping-row-cnsc-129-cm-7"></a><a id="mapping-row-cnsc-129-si-7"></a>`CNSC-129` | relates-to | `CM-7`<br />`SI-7` | <ul><li>Least Functionality</li><li>Software, Firmware, and Information Integrity</li></ul> |
| <a id="mapping-row-cnsc-13-ia-7"></a>`CNSC-13` | relates-to | `IA-7` | Cryptographic Module Authentication |
| <a id="mapping-row-cnsc-130-sa-9"></a><a id="mapping-row-cnsc-130-sc-30"></a>`CNSC-130` | relates-to | `SA-9`<br />`SC-30` | <ul><li>External System Services</li><li>Concealment and Misdirection</li></ul> |
| <a id="mapping-row-cnsc-131-cp-9"></a><a id="mapping-row-cnsc-131-mp-6"></a>`CNSC-131` | relates-to | `CP-9`<br />`MP-6` | <ul><li>System Backup</li><li>Media Sanitization</li></ul> |
| <a id="mapping-row-cnsc-132-sc-28"></a>`CNSC-132` | relates-to | `SC-28` | Protection of Information at Rest |
| <a id="mapping-row-cnsc-135-cm-6"></a><a id="mapping-row-cnsc-135-sa-8"></a><a id="mapping-row-cnsc-135-sc-7"></a>`CNSC-135` | relates-to | `CM-6`<br />`SA-8`<br />`SC-7` | <ul><li>Configuration Settings</li><li>Security and Privacy Engineering Principles</li><li>Boundary Protection</li></ul> |
| <a id="mapping-row-cnsc-136-cm-6"></a><a id="mapping-row-cnsc-136-sa-8"></a><a id="mapping-row-cnsc-136-sc-7"></a>`CNSC-136` | relates-to | `CM-6`<br />`SA-8`<br />`SC-7` | <ul><li>Configuration Settings</li><li>Security and Privacy Engineering Principles</li><li>Boundary Protection</li></ul> |
| <a id="mapping-row-cnsc-137-ac-16"></a><a id="mapping-row-cnsc-137-ac-4"></a><a id="mapping-row-cnsc-137-si-7"></a>`CNSC-137` | relates-to | `AC-16`<br />`AC-4`<br />`SI-7` | <ul><li>Security and Privacy Attributes</li><li>Information Flow Enforcement</li><li>Software, Firmware, and Information Integrity</li></ul> |
| <a id="mapping-row-cnsc-139-cm-14"></a>`CNSC-139` | relates-to | `CM-14` | Signed Components |
| <a id="mapping-row-cnsc-14-ia-2-1-2"></a>`CNSC-14` | relates-to | `IA-2(1)(2)` | Identification and Authentication (organizational Users) |
| <a id="mapping-row-cnsc-140-au-10"></a><a id="mapping-row-cnsc-140-cm-6"></a>`CNSC-140` | relates-to | `AU-10`<br />`CM-6` | <ul><li>Non-repudiation</li><li>Configuration Settings</li></ul> |
| <a id="mapping-row-cnsc-141-si-1"></a><a id="mapping-row-cnsc-141-si-7"></a>`CNSC-141` | relates-to | `SI-1`<br />`SI-7` | <ul><li>Policy and Procedures</li><li>Software, Firmware, and Information Integrity</li></ul> |
| <a id="mapping-row-cnsc-142-si-1"></a><a id="mapping-row-cnsc-142-si-7"></a>`CNSC-142` | relates-to | `SI-1`<br />`SI-7` | <ul><li>Policy and Procedures</li><li>Software, Firmware, and Information Integrity</li></ul> |
| <a id="mapping-row-cnsc-143-ia-5"></a>`CNSC-143` | relates-to | `IA-5` | Authenticator Management |
| <a id="mapping-row-cnsc-144-ac-4-6"></a>`CNSC-144` | relates-to | `AC-4(6)` | Information Flow Enforcement, Metadata |
| <a id="mapping-row-cnsc-145-ac-6"></a>`CNSC-145` | relates-to | `AC-6` | Least Privilege |
| <a id="mapping-row-cnsc-146-sc-12"></a>`CNSC-146` | relates-to | `SC-12` | Cryptographic Key Establishment and Management |
| <a id="mapping-row-cnsc-148-ia-5"></a><a id="mapping-row-cnsc-148-sc-12"></a><a id="mapping-row-cnsc-148-sc-13"></a><a id="mapping-row-cnsc-148-sc-28-1"></a><a id="mapping-row-cnsc-148-sc-8"></a>`CNSC-148` | relates-to | `IA-5`<br />`SC-12`<br />`SC-13`<br />`SC-28(1)`<br />`SC-8` | <ul><li>Authenticator Management</li><li>Cryptographic Key Establishment and Management</li><li>Cryptographic Protection</li><li>Protection of Information at Rest, Cryptographic Protection</li><li>Transmission Confidentiality and Integrity</li></ul> |
| <a id="mapping-row-cnsc-149-cm-3-6"></a>`CNSC-149` | relates-to | `CM-3(6)` | Configuration Change Control |
| <a id="mapping-row-cnsc-15-ac-4-4"></a>`CNSC-15` | relates-to | `AC-4(4)` | Information Flow Enforcement |
| <a id="mapping-row-cnsc-150-cm-3-2"></a>`CNSC-150` | relates-to | `CM-3(2)` | Configuration Change Control |
| <a id="mapping-row-cnsc-151-cm-3-4"></a>`CNSC-151` | relates-to | `CM-3(4)` | NIST SP 800-53 control mapped from CNSC-151 (Build Worker Runtime Security); family: Securing Build Pipelines. |
| <a id="mapping-row-cnsc-152-cm-3-4"></a>`CNSC-152` | relates-to | `CM-3(4)` | NIST SP 800-53 control mapped from CNSC-152 (Reproducible Builds); family: Securing Build Pipelines. |
| <a id="mapping-row-cnsc-153-cm-3-2"></a>`CNSC-153` | relates-to | `CM-3(2)` | NIST SP 800-53 control mapped from CNSC-153 (External Requirement Verification); family: Securing Build Pipelines. |
| <a id="mapping-row-cnsc-155-cm-3-1"></a>`CNSC-155` | relates-to | `CM-3(1)` | NIST SP 800-53 control mapped from CNSC-155 (Build Environment Recording); family: Securing Build Pipelines. |
| <a id="mapping-row-cnsc-156-cm-3-3"></a>`CNSC-156` | relates-to | `CM-3(3)` | NIST SP 800-53 control mapped from CNSC-156 (Build Environment Automation); family: Securing Build Pipelines. |
| <a id="mapping-row-cnsc-157-cm-3-3"></a>`CNSC-157` | relates-to | `CM-3(3)` | NIST SP 800-53 control mapped from CNSC-157 (Build Distribution); family: Securing Build Pipelines. |
| <a id="mapping-row-cnsc-158-sa-3"></a>`CNSC-158` | relates-to | `SA-3` | System Development Life Cycle |
| <a id="mapping-row-cnsc-16-si-12"></a>`CNSC-16` | relates-to | `SI-12` | Information Management and Retention |
| <a id="mapping-row-cnsc-161-ac-2"></a>`CNSC-161` | relates-to | `AC-2` | Account Management |
| <a id="mapping-row-cnsc-162-sc-7-3"></a>`CNSC-162` | relates-to | `SC-7(3)` | Boundary Protection |
| <a id="mapping-row-cnsc-163-ac-5"></a>`CNSC-163` | relates-to | `AC-5` | Separation of Duties |
| <a id="mapping-row-cnsc-164-cm-2-2"></a>`CNSC-164` | relates-to | `CM-2(2)` | Baseline Configuration |
| <a id="mapping-row-cnsc-165-au-9-2"></a>`CNSC-165` | relates-to | `AU-9(2)` | Protection of Audit Information |
| <a id="mapping-row-cnsc-167-ac-2"></a>`CNSC-167` | relates-to | `AC-2` | Account Management |
| <a id="mapping-row-cnsc-168-sc-17"></a>`CNSC-168` | relates-to | `SC-17` | Public Key Infrastructure Certificates |
| <a id="mapping-row-cnsc-169-sc-23-5"></a>`CNSC-169` | relates-to | `SC-23(5)` | Session Authenticity |
| <a id="mapping-row-cnsc-17-ac-16-3"></a>`CNSC-17` | relates-to | `AC-16(3)` | Security and Privacy Attributes |
| <a id="mapping-row-cnsc-170-si-7"></a>`CNSC-170` | relates-to | `SI-7` | Software, Firmware, and Information Integrity |
| <a id="mapping-row-cnsc-171-si-7"></a>`CNSC-171` | relates-to | `SI-7` | Software, Firmware, and Information Integrity |
| <a id="mapping-row-cnsc-173-sa-11"></a>`CNSC-173` | relates-to | `SA-11` | Developer Testing and Evaluation |
| <a id="mapping-row-cnsc-174-cm-8"></a>`CNSC-174` | relates-to | `CM-8` | Information System Component Inventory |
| <a id="mapping-row-cnsc-175-cm-10"></a>`CNSC-175` | relates-to | `CM-10` | Software Usage Restrictions |
| <a id="mapping-row-cnsc-179-ra-5"></a>`CNSC-179` | relates-to | `RA-5` | Vulnerability Monitoring and Scanning |
| <a id="mapping-row-cnsc-18-sc-12-1"></a>`CNSC-18` | relates-to | `SC-12(1)` | Cryptographic Key Establishment and Management |
| <a id="mapping-row-cnsc-180-cm-10"></a>`CNSC-180` | relates-to | `CM-10` | Software Usage Restrictions |
| <a id="mapping-row-cnsc-181-sa-11-1-8"></a>`CNSC-181` | relates-to | `SA-11(1)(8)` | Developer Testing and Evaluation |
| <a id="mapping-row-cnsc-182-si-7"></a>`CNSC-182` | relates-to | `SI-7` | Software, Firmware, and information integrity |
| <a id="mapping-row-cnsc-183-ac-6-3"></a>`CNSC-183` | relates-to | `AC-6(3)` | Least Privilege |
| <a id="mapping-row-cnsc-184-sc-12-3"></a>`CNSC-184` | relates-to | `SC-12(3)` | Systems and Communication Protection |
| <a id="mapping-row-cnsc-185-pl-1"></a>`CNSC-185` | relates-to | `PL-1` | Policy and Procedures |
| <a id="mapping-row-cnsc-186-ra-5"></a>`CNSC-186` | relates-to | `RA-5` | Vulnerability Monitoring and Scanning |
| <a id="mapping-row-cnsc-187-pl-1"></a>`CNSC-187` | relates-to | `PL-1` | Policy and Procedures |
| <a id="mapping-row-cnsc-188-pl-1"></a>`CNSC-188` | relates-to | `PL-1` | Policy and Procedures |
| <a id="mapping-row-cnsc-189-sa-11"></a>`CNSC-189` | relates-to | `SA-11` | Developer Testing and Evaluation |
| <a id="mapping-row-cnsc-19-si-12"></a>`CNSC-19` | relates-to | `SI-12` | Information Management and Retention |
| <a id="mapping-row-cnsc-190-sa-8"></a>`CNSC-190` | relates-to | `SA-8` | Security Engineering Principles |
| <a id="mapping-row-cnsc-191-ia-2-1"></a>`CNSC-191` | relates-to | `IA-2(1)` | Identification and Authentication (organizational Users) |
| <a id="mapping-row-cnsc-192-ac-1"></a>`CNSC-192` | relates-to | `AC-1` | Remote Access |
| <a id="mapping-row-cnsc-193-ac-2-1"></a>`CNSC-193` | relates-to | `AC-2(1)` | Prerequisites and criteria for group and role membership are defined. |
| <a id="mapping-row-cnsc-194-ac-2-1"></a>`CNSC-194` | relates-to | `AC-2(1)` | Usage of automated mechanisms to create, enable, modify, disable, and remove accounts. |
| <a id="mapping-row-cnsc-195-sa-8-23"></a>`CNSC-195` | relates-to | `SA-8(23)` | Security and Privacy Engineering Principles, Secure Defaults |
| <a id="mapping-row-cnsc-2-ia-9"></a>`CNSC-2` | relates-to | `IA-9` | NIST SP 800-53 control mapped from CNSC-2 (Mutual Authentication); family: Access. |
| <a id="mapping-row-cnsc-20-ac-16"></a>`CNSC-20` | relates-to | `AC-16` | Security and Privacy Attributes |
| <a id="mapping-row-cnsc-208-ia-2-1-2"></a>`CNSC-208` | relates-to | `IA-2(1)(2)` | Identification and Authentication (organizational Users) |
| <a id="mapping-row-cnsc-21-au-9-3"></a>`CNSC-21` | relates-to | `AU-9(3)` | Protection of Audit Information |
| <a id="mapping-row-cnsc-22-si-7-9"></a>`CNSC-22` | relates-to | `SI-7(9)` | Software, Firmware, and Information Integrity |
| <a id="mapping-row-cnsc-23-sc-7"></a>`CNSC-23` | relates-to | `SC-7` | Boundary Protection |
| <a id="mapping-row-cnsc-233-sc-7"></a>`CNSC-233` | relates-to | `SC-7` | Boundary Protection |
| <a id="mapping-row-cnsc-24-cm-2-2"></a><a id="mapping-row-cnsc-24-cm-3-7"></a>`CNSC-24` | relates-to | `CM-2(2)`<br />`CM-3(7)` | <ul><li>Baseline Configuration, Automation Support for Accuracy and Currency</li><li>Configuration Change Control, Review System Changes</li></ul> |
| <a id="mapping-row-cnsc-25-au-2"></a>`CNSC-25` | relates-to | `AU-2` | Event Logging |
| <a id="mapping-row-cnsc-259-sa-11-1"></a>`CNSC-259` | relates-to | `SA-11(1)` | Developer Testing and Evaluation \| Static Code Analysis |
| <a id="mapping-row-cnsc-26-cm-2"></a><a id="mapping-row-cnsc-26-cm-7"></a>`CNSC-26` | relates-to | `CM-2`<br />`CM-7` | <ul><li>Baseline Configuration</li><li>Least Functionality</li></ul> |
| <a id="mapping-row-cnsc-265-sa-11-4"></a>`CNSC-265` | relates-to | `SA-11(4)` | Developer Testing and Evaluation \| Manual Code Reviews |
| <a id="mapping-row-cnsc-27-si-7"></a>`CNSC-27` | relates-to | `SI-7` | Software, Firmware, and Information Integrity |
| <a id="mapping-row-cnsc-271-sc-39"></a>`CNSC-271` | relates-to | `SC-39` | Process Isolation |
| <a id="mapping-row-cnsc-28-ac-6"></a>`CNSC-28` | relates-to | `AC-6` | Least Privilege |
| <a id="mapping-row-cnsc-29-si-7-16"></a><a id="mapping-row-cnsc-29-si-7-17"></a>`CNSC-29` | relates-to | `SI-7(16)`<br />`SI-7(17)` | <ul><li>Software, Firmware, and Information Integrity, Time Limit on Process Execution Without Supervision</li><li>Software, Firmware, and Information Integrity, Runtime Application Self-protection</li></ul> |
| <a id="mapping-row-cnsc-297-si-7"></a>`CNSC-297` | relates-to | `SI-7` | Software, Firmware, and Information Integrity |
| <a id="mapping-row-cnsc-298-si-7"></a>`CNSC-298` | relates-to | `SI-7` | Software, Firmware, and Information Integrity |
| <a id="mapping-row-cnsc-3-sc-12"></a>`CNSC-3` | relates-to | `SC-12` | Cryptographic Key Establishment and Management |
| <a id="mapping-row-cnsc-30-si-4-13"></a>`CNSC-30` | relates-to | `SI-4(13)` | System Monitoring, Analyze Traffic and Event Patterns |
| <a id="mapping-row-cnsc-303-sc-12-3"></a>`CNSC-303` | relates-to | `SC-12(3)` | Systems and Communication Protection |
| <a id="mapping-row-cnsc-31-ac-3"></a>`CNSC-31` | relates-to | `AC-3` | Access Enforcement |
| <a id="mapping-row-cnsc-32-cm-2"></a><a id="mapping-row-cnsc-32-cm-7"></a>`CNSC-32` | relates-to | `CM-2`<br />`CM-7` | <ul><li>Baseline Configuration</li><li>Least Functionality</li></ul> |
| <a id="mapping-row-cnsc-33-cm-5"></a>`CNSC-33` | relates-to | `CM-5` | Access Restrictions for Change |
| <a id="mapping-row-cnsc-34-cm-5"></a>`CNSC-34` | relates-to | `CM-5` | Access Restrictions for Change |
| <a id="mapping-row-cnsc-35-sc-7"></a>`CNSC-35` | relates-to | `SC-7` | Boundary Protection |
| <a id="mapping-row-cnsc-36-sc-7"></a>`CNSC-36` | relates-to | `SC-7` | Boundary Protection |
| <a id="mapping-row-cnsc-37-cm-5"></a>`CNSC-37` | relates-to | `CM-5` | Access Restrictions for Change |
| <a id="mapping-row-cnsc-38-cm-5"></a>`CNSC-38` | relates-to | `CM-5` | Access Restrictions for Change |
| <a id="mapping-row-cnsc-39-sc-7"></a>`CNSC-39` | relates-to | `SC-7` | Boundary Protection |
| <a id="mapping-row-cnsc-4-sc-12-3"></a>`CNSC-4` | relates-to | `SC-12(3)` | Cryptographic Key Establishment and Management |
| <a id="mapping-row-cnsc-40-si-4"></a>`CNSC-40` | relates-to | `SI-4` | System Monitoring |
| <a id="mapping-row-cnsc-41-si-3"></a>`CNSC-41` | relates-to | `SI-3` | Malicious Code Protection |
| <a id="mapping-row-cnsc-42-ra-5"></a>`CNSC-42` | relates-to | `RA-5` | Vulnerability Monitoring and Scanning |
| <a id="mapping-row-cnsc-43-au-3"></a>`CNSC-43` | relates-to | `AU-3` | Content of Audit Records |
| <a id="mapping-row-cnsc-44-ac-6"></a>`CNSC-44` | relates-to | `AC-6` | Least Privilege |
| <a id="mapping-row-cnsc-45-si-7"></a>`CNSC-45` | relates-to | `SI-7` | Software, Firmware, and Information Integrity |
| <a id="mapping-row-cnsc-46-sc-12-3"></a>`CNSC-46` | relates-to | `SC-12(3)` | Systems and Communication Protection |
| <a id="mapping-row-cnsc-47-sc-12-3"></a>`CNSC-47` | relates-to | `SC-12(3)` | Systems and Communication Protection |
| <a id="mapping-row-cnsc-48-si-4"></a>`CNSC-48` | relates-to | `SI-4` | System Monitoring |
| <a id="mapping-row-cnsc-49-sc-28"></a>`CNSC-49` | relates-to | `SC-28` | Protection of Information at Rest |
| <a id="mapping-row-cnsc-5-ia-2-12"></a>`CNSC-5` | relates-to | `IA-2(12)` | Identification and Authentication (Organizational Users) |
| <a id="mapping-row-cnsc-50-cm-8"></a>`CNSC-50` | relates-to | `CM-8` | System Component Inventory |
| <a id="mapping-row-cnsc-51-cm-2"></a><a id="mapping-row-cnsc-51-cm-7"></a>`CNSC-51` | relates-to | `CM-2`<br />`CM-7` | <ul><li>Baseline Configuration</li><li>Least Functionality</li></ul> |
| <a id="mapping-row-cnsc-52-cm-5"></a>`CNSC-52` | relates-to | `CM-5` | Access Restrictions for Change |
| <a id="mapping-row-cnsc-53-cm-2"></a><a id="mapping-row-cnsc-53-cm-7"></a>`CNSC-53` | relates-to | `CM-2`<br />`CM-7` | <ul><li>Baseline Configuration</li><li>Least Functionality</li></ul> |
| <a id="mapping-row-cnsc-54-si-4"></a>`CNSC-54` | relates-to | `SI-4` | System Monitoring |
| <a id="mapping-row-cnsc-55-si-4"></a>`CNSC-55` | relates-to | `SI-4` | System Monitoring |
| <a id="mapping-row-cnsc-56-sc-7-21"></a>`CNSC-56` | relates-to | `SC-7(21)` | Boundary Protection, Isolation of System Components |
| <a id="mapping-row-cnsc-57-sr-4-3"></a><a id="mapping-row-cnsc-57-sr-4-4"></a>`CNSC-57` | relates-to | `SR-4(3)`<br />`SR-4(4)` | <ul><li>Provenance, Validate as Genuine and Not Altered</li><li>Provenance, Supply Chain Integrity - Pedigree</li></ul> |
| <a id="mapping-row-cnsc-58-si-7-17"></a>`CNSC-58` | relates-to | `SI-7(17)` | Software, Firmware, and Information Integrity, Runtime Application Self-Protection |
| <a id="mapping-row-cnsc-59-sr-4-3"></a><a id="mapping-row-cnsc-59-sr-4-4"></a>`CNSC-59` | relates-to | `SR-4(3)`<br />`SR-4(4)` | <ul><li>Provenance, Validate as Genuine and Not Altered</li><li>Provenance, Supply Chain Integrity - Pedigree</li></ul> |
| <a id="mapping-row-cnsc-6-ia-2-6"></a>`CNSC-6` | relates-to | `IA-2(6)` | Identification and Authentication (Organizational Users) |
| <a id="mapping-row-cnsc-60-cm-3"></a>`CNSC-60` | relates-to | `CM-3` | Configuration Change Control |
| <a id="mapping-row-cnsc-62-si-3"></a>`CNSC-62` | relates-to | `SI-3` | System and Information Integrity |
| <a id="mapping-row-cnsc-63-sa-3-1"></a>`CNSC-63` | relates-to | `SA-3(1)` | System Development Life Cycle \| Manage preproduction environment |
| <a id="mapping-row-cnsc-64-sa-8-31"></a>`CNSC-64` | relates-to | `SA-8(31)` | Security and Privacy Engineering Principles \| Secure System Modification |
| <a id="mapping-row-cnsc-65-sa-11-1"></a>`CNSC-65` | relates-to | `SA-11(1)` | Developer Testing and Evaluation \| Static Code Analysis |
| <a id="mapping-row-cnsc-66-sa-15"></a>`CNSC-66` | relates-to | `SA-15` | Development Process, Standards, and Tools |
| <a id="mapping-row-cnsc-67-sa-11"></a>`CNSC-67` | relates-to | `SA-11` | Developer Testing and Evaluation |
| <a id="mapping-row-cnsc-68-sa-11"></a>`CNSC-68` | relates-to | `SA-11` | Developer Testing and Evaluation |
| <a id="mapping-row-cnsc-69-sa-11"></a>`CNSC-69` | relates-to | `SA-11` | Developer Testing and Evaluation |
| <a id="mapping-row-cnsc-7-ia-2-6"></a>`CNSC-7` | relates-to | `IA-2(6)` | Identification and Authentication (Organizational Users) |
| <a id="mapping-row-cnsc-70-sa-11"></a>`CNSC-70` | relates-to | `SA-11` | Developer Testing and Evaluation |
| <a id="mapping-row-cnsc-71-sa-11-4"></a>`CNSC-71` | relates-to | `SA-11(4)` | Developer Testing and Evaluation \| Manual Code Reviews |
| <a id="mapping-row-cnsc-73-sa-11"></a>`CNSC-73` | relates-to | `SA-11` | Developer Testing and Evaluation |
| <a id="mapping-row-cnsc-74-sa-11"></a>`CNSC-74` | relates-to | `SA-11` | Developer Testing and Evaluation |
| <a id="mapping-row-cnsc-75-sa-11"></a>`CNSC-75` | relates-to | `SA-11` | Developer Testing and Evaluation |
| <a id="mapping-row-cnsc-76-sa-3-1"></a>`CNSC-76` | relates-to | `SA-3(1)` | System Development Life Cycle \| Manage preproduction environment |
| <a id="mapping-row-cnsc-77-sc-39"></a>`CNSC-77` | relates-to | `SC-39` | Process Isolation |
| <a id="mapping-row-cnsc-78-sa-11-2"></a>`CNSC-78` | relates-to | `SA-11(2)` | Developer Testing and Evaluation \| Threat Modeling and Vulnerability Analyses |
| <a id="mapping-row-cnsc-8-si-4-2"></a>`CNSC-8` | relates-to | `SI-4(2)` | System Monitoring |
| <a id="mapping-row-cnsc-85-ra-5"></a>`CNSC-85` | relates-to | `RA-5` | Vulnerability Monitoring and Scanning |
| <a id="mapping-row-cnsc-86-sc-39"></a>`CNSC-86` | relates-to | `SC-39` | Process Isolation |
| <a id="mapping-row-cnsc-87-sc-39"></a>`CNSC-87` | relates-to | `SC-39` | Process Isolation |
| <a id="mapping-row-cnsc-88-sa-1"></a>`CNSC-88` | relates-to | `SA-1` | Policy and Procedures |
| <a id="mapping-row-cnsc-89-si-7"></a>`CNSC-89` | relates-to | `SI-7` | Software, Firmware, and Information Integrity |
| <a id="mapping-row-cnsc-9-ac-3-13"></a>`CNSC-9` | relates-to | `AC-3(13)` | Access Enforcement |
| <a id="mapping-row-cnsc-90-si-7"></a>`CNSC-90` | relates-to | `SI-7` | Software, Firmware, and Information Integrity |
| <a id="mapping-row-cnsc-91-ra-5"></a>`CNSC-91` | relates-to | `RA-5` | Vulnerability Monitoring and Scanning |
| <a id="mapping-row-cnsc-92-sa-1"></a>`CNSC-92` | relates-to | `SA-1` | Policy and Procedures |
| <a id="mapping-row-cnsc-93-sa-11-8"></a>`CNSC-93` | relates-to | `SA-11(8)` | Interactive Application Security Testing |
| <a id="mapping-row-cnsc-94-si-4"></a>`CNSC-94` | relates-to | `SI-4` | System Monitoring |
| <a id="mapping-row-cnsc-97-si-4"></a>`CNSC-97` | relates-to | `SI-4` | System Monitoring |
| <a id="mapping-row-cnsc-99-ca-8"></a><a id="mapping-row-cnsc-99-sa-11"></a>`CNSC-99` | relates-to | `CA-8`<br />`SA-11` | <ul><li>Penetration Testing</li><li>Developer Testing and Evaluation</li></ul> |



## Document remarks

Auto-generated from YAML family files; 205 atomic guideline-to-control relationships.
