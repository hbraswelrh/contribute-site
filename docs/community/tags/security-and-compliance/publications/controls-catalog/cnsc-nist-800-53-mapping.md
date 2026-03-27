---
title: Cloud Native Security Controls (CNSC) to NIST SP 800-53 Rev. 5 mapping
sidebar_position: 7
toc_max_heading_level: 3
---

<!--
Auto-generated from cnsc-nist-800-53-mapping.yaml.
Do not edit manually.

  go run cmd/catalog/main.go -md index.md
-->

# Cloud Native Security Controls (CNSC) to NIST SP 800-53 Rev. 5 mapping

This document maps the [Cloud Native Security Controls Catalog](./index.md) to NIST 800-53 SP Rev.5 Guidelines.

## Complete Mapping Index

Mapping entry links on the [Cloud Native Security Controls Catalog page](./cnsc-nist-800-53-by-family) jump to these atomic mapping anchors.

| Mapping ID | CNSC guideline | NIST control(s) | Relationship | Remarks |
|------------|----------------|-----------------|--------------|---------|
| <a id="mapping-row-cnsc-1-ia-5-7"></a>`CNSC-1-IA-5-7` | `CNSC-1` | `IA-5(7)` | relates-to | Authenticator Management |
| <a id="mapping-row-cnsc-10-ac-3-13"></a>`CNSC-10-AC-3-13` | `CNSC-10` | `AC-3(13)` | relates-to | Access Enforcement |
| <a id="mapping-row-cnsc-100-ia-3-1"></a>`CNSC-100-IA-3-1` | `CNSC-100` | `IA-3(1)` | relates-to | Cryptographic Bidirectional Authentication |
| <a id="mapping-row-cnsc-101-si-7"></a>`CNSC-101-SI-7` | `CNSC-101` | `SI-7` | relates-to | Software, Firmware, and Information Integrity |
| <a id="mapping-row-cnsc-102-si-7"></a>`CNSC-102-SI-7` | `CNSC-102` | `SI-7` | relates-to | Software, Firmware, and Information Integrity |
| <a id="mapping-row-cnsc-103-si-7"></a>`CNSC-103-SI-7` | `CNSC-103` | `SI-7` | relates-to | Software, Firmware, and Information Integrity |
| <a id="mapping-row-cnsc-104-si-7"></a>`CNSC-104-SI-7` | `CNSC-104` | `SI-7` | relates-to | System and Information Integrity |
| <a id="mapping-row-cnsc-105-ra-5"></a>`CNSC-105-RA-5` | `CNSC-105` | `RA-5` | relates-to | Vulnerability Monitoring and Scanning |
| <a id="mapping-row-cnsc-105-sa-3"></a>`CNSC-105-SA-3` | `CNSC-105` | `SA-3` | relates-to | System Development Life Cycle |
| <a id="mapping-row-cnsc-106-si-7"></a>`CNSC-106-SI-7` | `CNSC-106` | `SI-7` | relates-to | System and Information Integrity |
| <a id="mapping-row-cnsc-107-si-2-3"></a>`CNSC-107-SI-2-3` | `CNSC-107` | `SI-2(3)` | relates-to | System and Information Integrity |
| <a id="mapping-row-cnsc-108-sc-12-3"></a>`CNSC-108-SC-12-3` | `CNSC-108` | `SC-12(3)` | relates-to | Systems and Communication Protection |
| <a id="mapping-row-cnsc-109-si-2-3"></a>`CNSC-109-SI-2-3` | `CNSC-109` | `SI-2(3)` | relates-to | System and Information Integrity |
| <a id="mapping-row-cnsc-11-sc-7-19"></a>`CNSC-11-SC-7-19` | `CNSC-11` | `SC-7(19)` | relates-to | Boundary Protection |
| <a id="mapping-row-cnsc-110-pl-1"></a>`CNSC-110-PL-1` | `CNSC-110` | `PL-1` | relates-to | Policy and Procedures |
| <a id="mapping-row-cnsc-112-ac-6-3"></a>`CNSC-112-AC-6-3` | `CNSC-112` | `AC-6(3)` | relates-to | Least Privilege |
| <a id="mapping-row-cnsc-113-sc-12-2"></a>`CNSC-113-SC-12-2` | `CNSC-113` | `SC-12(2)` | relates-to | Cryptographic Key Establishment and Management \| Symmetric and Asymmetric Keys |
| <a id="mapping-row-cnsc-113-sc-12-3"></a>`CNSC-113-SC-12-3` | `CNSC-113` | `SC-12(3)` | relates-to | Cryptographic Key Establishment and Management \| Symmetric and Asymmetric Keys |
| <a id="mapping-row-cnsc-114-si-2-3"></a>`CNSC-114-SI-2-3` | `CNSC-114` | `SI-2(3)` | relates-to | System and Information Integrity |
| <a id="mapping-row-cnsc-115-ac-6-3"></a>`CNSC-115-AC-6-3` | `CNSC-115` | `AC-6(3)` | relates-to | Least Privilege, Network Access to Privileged Commands |
| <a id="mapping-row-cnsc-116-ca-7"></a>`CNSC-116-CA-7` | `CNSC-116` | `CA-7` | relates-to | Continuous Monitoring |
| <a id="mapping-row-cnsc-116-ir-4"></a>`CNSC-116-IR-4` | `CNSC-116` | `IR-4` | relates-to | Incident Handling, Automated Incident Handling Processes |
| <a id="mapping-row-cnsc-116-ir-4-5"></a>`CNSC-116-IR-4-5` | `CNSC-116` | `IR-4(5)` | relates-to | Incident Handling, Automated Disabling of System |
| <a id="mapping-row-cnsc-117-ir-5-1"></a>`CNSC-117-IR-5-1` | `CNSC-117` | `IR-5(1)` | relates-to | Incident Monitoring, Automated Tracking, Data Collection, and Analysis |
| <a id="mapping-row-cnsc-12-ia-7"></a>`CNSC-12-IA-7` | `CNSC-12` | `IA-7` | relates-to | Cryptographic Module Authentication |
| <a id="mapping-row-cnsc-120-ac-3-3"></a>`CNSC-120-AC-3-3` | `CNSC-120` | `AC-3(3)` | relates-to | Access Enforcement, Mandatory Access Control |
| <a id="mapping-row-cnsc-121-sa-11-2"></a>`CNSC-121-SA-11-2` | `CNSC-121` | `SA-11(2)` | relates-to | Developer Testing and Evaluation, Threat Modeling and Vulnerability Analyses |
| <a id="mapping-row-cnsc-122-ia-9"></a>`CNSC-122-IA-9` | `CNSC-122` | `IA-9` | relates-to | Service Identification and Authentication |
| <a id="mapping-row-cnsc-123-ia-9"></a>`CNSC-123-IA-9` | `CNSC-123` | `IA-9` | relates-to | Service Identification and Authentication |
| <a id="mapping-row-cnsc-124-cm-14"></a>`CNSC-124-CM-14` | `CNSC-124` | `CM-14` | relates-to | Signed Components |
| <a id="mapping-row-cnsc-125-si-6"></a>`CNSC-125-SI-6` | `CNSC-125` | `SI-6` | relates-to | Security and Privacy Function Verification |
| <a id="mapping-row-cnsc-127-sc-8"></a>`CNSC-127-SC-8` | `CNSC-127` | `SC-8` | relates-to | Transmission Confidentiality and Integrity |
| <a id="mapping-row-cnsc-128-si-13"></a>`CNSC-128-SI-13` | `CNSC-128` | `SI-13` | relates-to | Predictable Failure Prevention |
| <a id="mapping-row-cnsc-129-cm-7"></a>`CNSC-129-CM-7` | `CNSC-129` | `CM-7` | relates-to | Least Functionality |
| <a id="mapping-row-cnsc-129-si-7"></a>`CNSC-129-SI-7` | `CNSC-129` | `SI-7` | relates-to | Software, Firmware, and Information Integrity |
| <a id="mapping-row-cnsc-13-ia-7"></a>`CNSC-13-IA-7` | `CNSC-13` | `IA-7` | relates-to | Cryptographic Module Authentication |
| <a id="mapping-row-cnsc-130-sa-9"></a>`CNSC-130-SA-9` | `CNSC-130` | `SA-9` | relates-to | External System Services |
| <a id="mapping-row-cnsc-130-sc-30"></a>`CNSC-130-SC-30` | `CNSC-130` | `SC-30` | relates-to | Concealment and Misdirection |
| <a id="mapping-row-cnsc-131-cp-9"></a>`CNSC-131-CP-9` | `CNSC-131` | `CP-9` | relates-to | System Backup |
| <a id="mapping-row-cnsc-131-mp-6"></a>`CNSC-131-MP-6` | `CNSC-131` | `MP-6` | relates-to | Media Sanitization |
| <a id="mapping-row-cnsc-132-sc-28"></a>`CNSC-132-SC-28` | `CNSC-132` | `SC-28` | relates-to | Protection of Information at Rest |
| <a id="mapping-row-cnsc-135-cm-6"></a>`CNSC-135-CM-6` | `CNSC-135` | `CM-6` | relates-to | Configuration Settings |
| <a id="mapping-row-cnsc-135-sa-8"></a>`CNSC-135-SA-8` | `CNSC-135` | `SA-8` | relates-to | Security and Privacy Engineering Principles |
| <a id="mapping-row-cnsc-135-sc-7"></a>`CNSC-135-SC-7` | `CNSC-135` | `SC-7` | relates-to | Boundary Protection |
| <a id="mapping-row-cnsc-136-cm-6"></a>`CNSC-136-CM-6` | `CNSC-136` | `CM-6` | relates-to | Configuration Settings |
| <a id="mapping-row-cnsc-136-sa-8"></a>`CNSC-136-SA-8` | `CNSC-136` | `SA-8` | relates-to | Security and Privacy Engineering Principles |
| <a id="mapping-row-cnsc-136-sc-7"></a>`CNSC-136-SC-7` | `CNSC-136` | `SC-7` | relates-to | Boundary Protection |
| <a id="mapping-row-cnsc-137-ac-16"></a>`CNSC-137-AC-16` | `CNSC-137` | `AC-16` | relates-to | Security and Privacy Attributes |
| <a id="mapping-row-cnsc-137-ac-4"></a>`CNSC-137-AC-4` | `CNSC-137` | `AC-4` | relates-to | Information Flow Enforcement |
| <a id="mapping-row-cnsc-137-si-7"></a>`CNSC-137-SI-7` | `CNSC-137` | `SI-7` | relates-to | Software, Firmware, and Information Integrity |
| <a id="mapping-row-cnsc-139-cm-14"></a>`CNSC-139-CM-14` | `CNSC-139` | `CM-14` | relates-to | Signed Components |
| <a id="mapping-row-cnsc-14-ia-2-1-2"></a>`CNSC-14-IA-2-1-2` | `CNSC-14` | `IA-2(1)(2)` | relates-to | Identification and Authentication (organizational Users) |
| <a id="mapping-row-cnsc-140-au-10"></a>`CNSC-140-AU-10` | `CNSC-140` | `AU-10` | relates-to | Non-repudiation |
| <a id="mapping-row-cnsc-140-cm-6"></a>`CNSC-140-CM-6` | `CNSC-140` | `CM-6` | relates-to | Configuration Settings |
| <a id="mapping-row-cnsc-141-si-1"></a>`CNSC-141-SI-1` | `CNSC-141` | `SI-1` | relates-to | Policy and Procedures |
| <a id="mapping-row-cnsc-141-si-7"></a>`CNSC-141-SI-7` | `CNSC-141` | `SI-7` | relates-to | Software, Firmware, and Information Integrity |
| <a id="mapping-row-cnsc-142-si-1"></a>`CNSC-142-SI-1` | `CNSC-142` | `SI-1` | relates-to | Policy and Procedures |
| <a id="mapping-row-cnsc-142-si-7"></a>`CNSC-142-SI-7` | `CNSC-142` | `SI-7` | relates-to | Software, Firmware, and Information Integrity |
| <a id="mapping-row-cnsc-143-ia-5"></a>`CNSC-143-IA-5` | `CNSC-143` | `IA-5` | relates-to | Authenticator Management |
| <a id="mapping-row-cnsc-144-ac-4-6"></a>`CNSC-144-AC-4-6` | `CNSC-144` | `AC-4(6)` | relates-to | Information Flow Enforcement, Metadata |
| <a id="mapping-row-cnsc-145-ac-6"></a>`CNSC-145-AC-6` | `CNSC-145` | `AC-6` | relates-to | Least Privilege |
| <a id="mapping-row-cnsc-146-sc-12"></a>`CNSC-146-SC-12` | `CNSC-146` | `SC-12` | relates-to | Cryptographic Key Establishment and Management |
| <a id="mapping-row-cnsc-148-ia-5"></a>`CNSC-148-IA-5` | `CNSC-148` | `IA-5` | relates-to | Authenticator Management |
| <a id="mapping-row-cnsc-148-sc-12"></a>`CNSC-148-SC-12` | `CNSC-148` | `SC-12` | relates-to | Cryptographic Key Establishment and Management |
| <a id="mapping-row-cnsc-148-sc-13"></a>`CNSC-148-SC-13` | `CNSC-148` | `SC-13` | relates-to | Cryptographic Protection |
| <a id="mapping-row-cnsc-148-sc-28-1"></a>`CNSC-148-SC-28-1` | `CNSC-148` | `SC-28(1)` | relates-to | Protection of Information at Rest, Cryptographic Protection |
| <a id="mapping-row-cnsc-148-sc-8"></a>`CNSC-148-SC-8` | `CNSC-148` | `SC-8` | relates-to | Transmission Confidentiality and Integrity |
| <a id="mapping-row-cnsc-149-cm-3-6"></a>`CNSC-149-CM-3-6` | `CNSC-149` | `CM-3(6)` | relates-to | Configuration Change Control |
| <a id="mapping-row-cnsc-15-ac-4-4"></a>`CNSC-15-AC-4-4` | `CNSC-15` | `AC-4(4)` | relates-to | Information Flow Enforcement |
| <a id="mapping-row-cnsc-150-cm-3-2"></a>`CNSC-150-CM-3-2` | `CNSC-150` | `CM-3(2)` | relates-to | Configuration Change Control |
| <a id="mapping-row-cnsc-151-cm-3-4"></a>`CNSC-151-CM-3-4` | `CNSC-151` | `CM-3(4)` | relates-to | NIST SP 800-53 control mapped from CNSC-151 (Build Worker Runtime Security); family: Securing Build Pipelines. |
| <a id="mapping-row-cnsc-152-cm-3-4"></a>`CNSC-152-CM-3-4` | `CNSC-152` | `CM-3(4)` | relates-to | NIST SP 800-53 control mapped from CNSC-152 (Reproducible Builds); family: Securing Build Pipelines. |
| <a id="mapping-row-cnsc-153-cm-3-2"></a>`CNSC-153-CM-3-2` | `CNSC-153` | `CM-3(2)` | relates-to | NIST SP 800-53 control mapped from CNSC-153 (External Requirement Verification); family: Securing Build Pipelines. |
| <a id="mapping-row-cnsc-155-cm-3-1"></a>`CNSC-155-CM-3-1` | `CNSC-155` | `CM-3(1)` | relates-to | NIST SP 800-53 control mapped from CNSC-155 (Build Environment Recording); family: Securing Build Pipelines. |
| <a id="mapping-row-cnsc-156-cm-3-3"></a>`CNSC-156-CM-3-3` | `CNSC-156` | `CM-3(3)` | relates-to | NIST SP 800-53 control mapped from CNSC-156 (Build Environment Automation); family: Securing Build Pipelines. |
| <a id="mapping-row-cnsc-157-cm-3-3"></a>`CNSC-157-CM-3-3` | `CNSC-157` | `CM-3(3)` | relates-to | NIST SP 800-53 control mapped from CNSC-157 (Build Distribution); family: Securing Build Pipelines. |
| <a id="mapping-row-cnsc-158-sa-3"></a>`CNSC-158-SA-3` | `CNSC-158` | `SA-3` | relates-to | System Development Life Cycle |
| <a id="mapping-row-cnsc-16-si-12"></a>`CNSC-16-SI-12` | `CNSC-16` | `SI-12` | relates-to | Information Management and Retention |
| <a id="mapping-row-cnsc-161-ac-2"></a>`CNSC-161-AC-2` | `CNSC-161` | `AC-2` | relates-to | Account Management |
| <a id="mapping-row-cnsc-162-sc-7-3"></a>`CNSC-162-SC-7-3` | `CNSC-162` | `SC-7(3)` | relates-to | Boundary Protection |
| <a id="mapping-row-cnsc-163-ac-5"></a>`CNSC-163-AC-5` | `CNSC-163` | `AC-5` | relates-to | Separation of Duties |
| <a id="mapping-row-cnsc-164-cm-2-2"></a>`CNSC-164-CM-2-2` | `CNSC-164` | `CM-2(2)` | relates-to | Baseline Configuration |
| <a id="mapping-row-cnsc-165-au-9-2"></a>`CNSC-165-AU-9-2` | `CNSC-165` | `AU-9(2)` | relates-to | Protection of Audit Information |
| <a id="mapping-row-cnsc-167-ac-2"></a>`CNSC-167-AC-2` | `CNSC-167` | `AC-2` | relates-to | Account Management |
| <a id="mapping-row-cnsc-168-sc-17"></a>`CNSC-168-SC-17` | `CNSC-168` | `SC-17` | relates-to | Public Key Infrastructure Certificates |
| <a id="mapping-row-cnsc-169-sc-23-5"></a>`CNSC-169-SC-23-5` | `CNSC-169` | `SC-23(5)` | relates-to | Session Authenticity |
| <a id="mapping-row-cnsc-17-ac-16-3"></a>`CNSC-17-AC-16-3` | `CNSC-17` | `AC-16(3)` | relates-to | Security and Privacy Attributes |
| <a id="mapping-row-cnsc-170-si-7"></a>`CNSC-170-SI-7` | `CNSC-170` | `SI-7` | relates-to | Software, Firmware, and Information Integrity |
| <a id="mapping-row-cnsc-171-si-7"></a>`CNSC-171-SI-7` | `CNSC-171` | `SI-7` | relates-to | Software, Firmware, and Information Integrity |
| <a id="mapping-row-cnsc-173-sa-11"></a>`CNSC-173-SA-11` | `CNSC-173` | `SA-11` | relates-to | Developer Testing and Evaluation |
| <a id="mapping-row-cnsc-174-cm-8"></a>`CNSC-174-CM-8` | `CNSC-174` | `CM-8` | relates-to | Information System Component Inventory |
| <a id="mapping-row-cnsc-175-cm-10"></a>`CNSC-175-CM-10` | `CNSC-175` | `CM-10` | relates-to | Software Usage Restrictions |
| <a id="mapping-row-cnsc-179-ra-5"></a>`CNSC-179-RA-5` | `CNSC-179` | `RA-5` | relates-to | Vulnerability Monitoring and Scanning |
| <a id="mapping-row-cnsc-18-sc-12-1"></a>`CNSC-18-SC-12-1` | `CNSC-18` | `SC-12(1)` | relates-to | Cryptographic Key Establishment and Management |
| <a id="mapping-row-cnsc-180-cm-10"></a>`CNSC-180-CM-10` | `CNSC-180` | `CM-10` | relates-to | Software Usage Restrictions |
| <a id="mapping-row-cnsc-181-sa-11-1-8"></a>`CNSC-181-SA-11-1-8` | `CNSC-181` | `SA-11(1)(8)` | relates-to | Developer Testing and Evaluation |
| <a id="mapping-row-cnsc-182-si-7"></a>`CNSC-182-SI-7` | `CNSC-182` | `SI-7` | relates-to | Software, Firmware, and information integrity |
| <a id="mapping-row-cnsc-183-ac-6-3"></a>`CNSC-183-AC-6-3` | `CNSC-183` | `AC-6(3)` | relates-to | Least Privilege |
| <a id="mapping-row-cnsc-184-sc-12-3"></a>`CNSC-184-SC-12-3` | `CNSC-184` | `SC-12(3)` | relates-to | Systems and Communication Protection |
| <a id="mapping-row-cnsc-185-pl-1"></a>`CNSC-185-PL-1` | `CNSC-185` | `PL-1` | relates-to | Policy and Procedures |
| <a id="mapping-row-cnsc-186-ra-5"></a>`CNSC-186-RA-5` | `CNSC-186` | `RA-5` | relates-to | Vulnerability Monitoring and Scanning |
| <a id="mapping-row-cnsc-187-pl-1"></a>`CNSC-187-PL-1` | `CNSC-187` | `PL-1` | relates-to | Policy and Procedures |
| <a id="mapping-row-cnsc-188-pl-1"></a>`CNSC-188-PL-1` | `CNSC-188` | `PL-1` | relates-to | Policy and Procedures |
| <a id="mapping-row-cnsc-189-sa-11"></a>`CNSC-189-SA-11` | `CNSC-189` | `SA-11` | relates-to | Developer Testing and Evaluation |
| <a id="mapping-row-cnsc-19-si-12"></a>`CNSC-19-SI-12` | `CNSC-19` | `SI-12` | relates-to | Information Management and Retention |
| <a id="mapping-row-cnsc-190-sa-8"></a>`CNSC-190-SA-8` | `CNSC-190` | `SA-8` | relates-to | Security Engineering Principles |
| <a id="mapping-row-cnsc-191-ia-2-1"></a>`CNSC-191-IA-2-1` | `CNSC-191` | `IA-2(1)` | relates-to | Identification and Authentication (organizational Users) |
| <a id="mapping-row-cnsc-192-ac-1"></a>`CNSC-192-AC-1` | `CNSC-192` | `AC-1` | relates-to | Remote Access |
| <a id="mapping-row-cnsc-193-ac-2-1"></a>`CNSC-193-AC-2-1` | `CNSC-193` | `AC-2(1)` | relates-to | Prerequisites and criteria for group and role membership are defined. |
| <a id="mapping-row-cnsc-194-ac-2-1"></a>`CNSC-194-AC-2-1` | `CNSC-194` | `AC-2(1)` | relates-to | Usage of automated mechanisms to create, enable, modify, disable, and remove accounts. |
| <a id="mapping-row-cnsc-195-sa-8-23"></a>`CNSC-195-SA-8-23` | `CNSC-195` | `SA-8(23)` | relates-to | Security and Privacy Engineering Principles, Secure Defaults |
| <a id="mapping-row-cnsc-2-ia-9"></a>`CNSC-2-IA-9` | `CNSC-2` | `IA-9` | relates-to | NIST SP 800-53 control mapped from CNSC-2 (Mutual Authentication); family: Access. |
| <a id="mapping-row-cnsc-20-ac-16"></a>`CNSC-20-AC-16` | `CNSC-20` | `AC-16` | relates-to | Security and Privacy Attributes |
| <a id="mapping-row-cnsc-208-ia-2-1-2"></a>`CNSC-208-IA-2-1-2` | `CNSC-208` | `IA-2(1)(2)` | relates-to | Identification and Authentication (organizational Users) |
| <a id="mapping-row-cnsc-21-au-9-3"></a>`CNSC-21-AU-9-3` | `CNSC-21` | `AU-9(3)` | relates-to | Protection of Audit Information |
| <a id="mapping-row-cnsc-22-si-7-9"></a>`CNSC-22-SI-7-9` | `CNSC-22` | `SI-7(9)` | relates-to | Software, Firmware, and Information Integrity |
| <a id="mapping-row-cnsc-23-sc-7"></a>`CNSC-23-SC-7` | `CNSC-23` | `SC-7` | relates-to | Boundary Protection |
| <a id="mapping-row-cnsc-233-sc-7"></a>`CNSC-233-SC-7` | `CNSC-233` | `SC-7` | relates-to | Boundary Protection |
| <a id="mapping-row-cnsc-24-cm-2-2"></a>`CNSC-24-CM-2-2` | `CNSC-24` | `CM-2(2)` | relates-to | Baseline Configuration, Automation Support for Accuracy and Currency |
| <a id="mapping-row-cnsc-24-cm-3-7"></a>`CNSC-24-CM-3-7` | `CNSC-24` | `CM-3(7)` | relates-to | Configuration Change Control, Review System Changes |
| <a id="mapping-row-cnsc-25-au-2"></a>`CNSC-25-AU-2` | `CNSC-25` | `AU-2` | relates-to | Event Logging |
| <a id="mapping-row-cnsc-259-sa-11-1"></a>`CNSC-259-SA-11-1` | `CNSC-259` | `SA-11(1)` | relates-to | Developer Testing and Evaluation \| Static Code Analysis |
| <a id="mapping-row-cnsc-26-cm-2"></a>`CNSC-26-CM-2` | `CNSC-26` | `CM-2` | relates-to | Baseline Configuration |
| <a id="mapping-row-cnsc-26-cm-7"></a>`CNSC-26-CM-7` | `CNSC-26` | `CM-7` | relates-to | Least Functionality |
| <a id="mapping-row-cnsc-265-sa-11-4"></a>`CNSC-265-SA-11-4` | `CNSC-265` | `SA-11(4)` | relates-to | Developer Testing and Evaluation \| Manual Code Reviews |
| <a id="mapping-row-cnsc-27-si-7"></a>`CNSC-27-SI-7` | `CNSC-27` | `SI-7` | relates-to | Software, Firmware, and Information Integrity |
| <a id="mapping-row-cnsc-271-sc-39"></a>`CNSC-271-SC-39` | `CNSC-271` | `SC-39` | relates-to | Process Isolation |
| <a id="mapping-row-cnsc-28-ac-6"></a>`CNSC-28-AC-6` | `CNSC-28` | `AC-6` | relates-to | Least Privilege |
| <a id="mapping-row-cnsc-29-si-7-16"></a>`CNSC-29-SI-7-16` | `CNSC-29` | `SI-7(16)` | relates-to | Software, Firmware, and Information Integrity, Time Limit on Process Execution Without Supervision |
| <a id="mapping-row-cnsc-29-si-7-17"></a>`CNSC-29-SI-7-17` | `CNSC-29` | `SI-7(17)` | relates-to | Software, Firmware, and Information Integrity, Runtime Application Self-protection |
| <a id="mapping-row-cnsc-297-si-7"></a>`CNSC-297-SI-7` | `CNSC-297` | `SI-7` | relates-to | Software, Firmware, and Information Integrity |
| <a id="mapping-row-cnsc-298-si-7"></a>`CNSC-298-SI-7` | `CNSC-298` | `SI-7` | relates-to | Software, Firmware, and Information Integrity |
| <a id="mapping-row-cnsc-3-sc-12"></a>`CNSC-3-SC-12` | `CNSC-3` | `SC-12` | relates-to | Cryptographic Key Establishment and Management |
| <a id="mapping-row-cnsc-30-si-4-13"></a>`CNSC-30-SI-4-13` | `CNSC-30` | `SI-4(13)` | relates-to | System Monitoring, Analyze Traffic and Event Patterns |
| <a id="mapping-row-cnsc-303-sc-12-3"></a>`CNSC-303-SC-12-3` | `CNSC-303` | `SC-12(3)` | relates-to | Systems and Communication Protection |
| <a id="mapping-row-cnsc-31-ac-3"></a>`CNSC-31-AC-3` | `CNSC-31` | `AC-3` | relates-to | Access Enforcement |
| <a id="mapping-row-cnsc-32-cm-2"></a>`CNSC-32-CM-2` | `CNSC-32` | `CM-2` | relates-to | Baseline Configuration |
| <a id="mapping-row-cnsc-32-cm-7"></a>`CNSC-32-CM-7` | `CNSC-32` | `CM-7` | relates-to | Least Functionality |
| <a id="mapping-row-cnsc-33-cm-5"></a>`CNSC-33-CM-5` | `CNSC-33` | `CM-5` | relates-to | Access Restrictions for Change |
| <a id="mapping-row-cnsc-34-cm-5"></a>`CNSC-34-CM-5` | `CNSC-34` | `CM-5` | relates-to | Access Restrictions for Change |
| <a id="mapping-row-cnsc-35-sc-7"></a>`CNSC-35-SC-7` | `CNSC-35` | `SC-7` | relates-to | Boundary Protection |
| <a id="mapping-row-cnsc-36-sc-7"></a>`CNSC-36-SC-7` | `CNSC-36` | `SC-7` | relates-to | Boundary Protection |
| <a id="mapping-row-cnsc-37-cm-5"></a>`CNSC-37-CM-5` | `CNSC-37` | `CM-5` | relates-to | Access Restrictions for Change |
| <a id="mapping-row-cnsc-38-cm-5"></a>`CNSC-38-CM-5` | `CNSC-38` | `CM-5` | relates-to | Access Restrictions for Change |
| <a id="mapping-row-cnsc-39-sc-7"></a>`CNSC-39-SC-7` | `CNSC-39` | `SC-7` | relates-to | Boundary Protection |
| <a id="mapping-row-cnsc-4-sc-12-3"></a>`CNSC-4-SC-12-3` | `CNSC-4` | `SC-12(3)` | relates-to | Cryptographic Key Establishment and Management |
| <a id="mapping-row-cnsc-40-si-4"></a>`CNSC-40-SI-4` | `CNSC-40` | `SI-4` | relates-to | System Monitoring |
| <a id="mapping-row-cnsc-41-si-3"></a>`CNSC-41-SI-3` | `CNSC-41` | `SI-3` | relates-to | Malicious Code Protection |
| <a id="mapping-row-cnsc-42-ra-5"></a>`CNSC-42-RA-5` | `CNSC-42` | `RA-5` | relates-to | Vulnerability Monitoring and Scanning |
| <a id="mapping-row-cnsc-43-au-3"></a>`CNSC-43-AU-3` | `CNSC-43` | `AU-3` | relates-to | Content of Audit Records |
| <a id="mapping-row-cnsc-44-ac-6"></a>`CNSC-44-AC-6` | `CNSC-44` | `AC-6` | relates-to | Least Privilege |
| <a id="mapping-row-cnsc-45-si-7"></a>`CNSC-45-SI-7` | `CNSC-45` | `SI-7` | relates-to | Software, Firmware, and Information Integrity |
| <a id="mapping-row-cnsc-46-sc-12-3"></a>`CNSC-46-SC-12-3` | `CNSC-46` | `SC-12(3)` | relates-to | Systems and Communication Protection |
| <a id="mapping-row-cnsc-47-sc-12-3"></a>`CNSC-47-SC-12-3` | `CNSC-47` | `SC-12(3)` | relates-to | Systems and Communication Protection |
| <a id="mapping-row-cnsc-48-si-4"></a>`CNSC-48-SI-4` | `CNSC-48` | `SI-4` | relates-to | System Monitoring |
| <a id="mapping-row-cnsc-49-sc-28"></a>`CNSC-49-SC-28` | `CNSC-49` | `SC-28` | relates-to | Protection of Information at Rest |
| <a id="mapping-row-cnsc-5-ia-2-12"></a>`CNSC-5-IA-2-12` | `CNSC-5` | `IA-2(12)` | relates-to | Identification and Authentication (Organizational Users) |
| <a id="mapping-row-cnsc-50-cm-8"></a>`CNSC-50-CM-8` | `CNSC-50` | `CM-8` | relates-to | System Component Inventory |
| <a id="mapping-row-cnsc-51-cm-2"></a>`CNSC-51-CM-2` | `CNSC-51` | `CM-2` | relates-to | Baseline Configuration |
| <a id="mapping-row-cnsc-51-cm-7"></a>`CNSC-51-CM-7` | `CNSC-51` | `CM-7` | relates-to | Least Functionality |
| <a id="mapping-row-cnsc-52-cm-5"></a>`CNSC-52-CM-5` | `CNSC-52` | `CM-5` | relates-to | Access Restrictions for Change |
| <a id="mapping-row-cnsc-53-cm-2"></a>`CNSC-53-CM-2` | `CNSC-53` | `CM-2` | relates-to | Baseline Configuration |
| <a id="mapping-row-cnsc-53-cm-7"></a>`CNSC-53-CM-7` | `CNSC-53` | `CM-7` | relates-to | Least Functionality |
| <a id="mapping-row-cnsc-54-si-4"></a>`CNSC-54-SI-4` | `CNSC-54` | `SI-4` | relates-to | System Monitoring |
| <a id="mapping-row-cnsc-55-si-4"></a>`CNSC-55-SI-4` | `CNSC-55` | `SI-4` | relates-to | System Monitoring |
| <a id="mapping-row-cnsc-56-sc-7-21"></a>`CNSC-56-SC-7-21` | `CNSC-56` | `SC-7(21)` | relates-to | Boundary Protection, Isolation of System Components |
| <a id="mapping-row-cnsc-57-sr-4-3"></a>`CNSC-57-SR-4-3` | `CNSC-57` | `SR-4(3)` | relates-to | Provenance, Validate as Genuine and Not Altered |
| <a id="mapping-row-cnsc-57-sr-4-4"></a>`CNSC-57-SR-4-4` | `CNSC-57` | `SR-4(4)` | relates-to | Provenance, Supply Chain Integrity - Pedigree |
| <a id="mapping-row-cnsc-58-si-7-17"></a>`CNSC-58-SI-7-17` | `CNSC-58` | `SI-7(17)` | relates-to | Software, Firmware, and Information Integrity, Runtime Application Self-Protection |
| <a id="mapping-row-cnsc-59-sr-4-3"></a>`CNSC-59-SR-4-3` | `CNSC-59` | `SR-4(3)` | relates-to | Provenance, Validate as Genuine and Not Altered |
| <a id="mapping-row-cnsc-59-sr-4-4"></a>`CNSC-59-SR-4-4` | `CNSC-59` | `SR-4(4)` | relates-to | Provenance, Supply Chain Integrity - Pedigree |
| <a id="mapping-row-cnsc-6-ia-2-6"></a>`CNSC-6-IA-2-6` | `CNSC-6` | `IA-2(6)` | relates-to | Identification and Authentication (Organizational Users) |
| <a id="mapping-row-cnsc-60-cm-3"></a>`CNSC-60-CM-3` | `CNSC-60` | `CM-3` | relates-to | Configuration Change Control |
| <a id="mapping-row-cnsc-62-si-3"></a>`CNSC-62-SI-3` | `CNSC-62` | `SI-3` | relates-to | System and Information Integrity |
| <a id="mapping-row-cnsc-63-sa-3-1"></a>`CNSC-63-SA-3-1` | `CNSC-63` | `SA-3(1)` | relates-to | System Development Life Cycle \| Manage preproduction environment |
| <a id="mapping-row-cnsc-64-sa-8-31"></a>`CNSC-64-SA-8-31` | `CNSC-64` | `SA-8(31)` | relates-to | Security and Privacy Engineering Principles \| Secure System Modification |
| <a id="mapping-row-cnsc-65-sa-11-1"></a>`CNSC-65-SA-11-1` | `CNSC-65` | `SA-11(1)` | relates-to | Developer Testing and Evaluation \| Static Code Analysis |
| <a id="mapping-row-cnsc-66-sa-15"></a>`CNSC-66-SA-15` | `CNSC-66` | `SA-15` | relates-to | Development Process, Standards, and Tools |
| <a id="mapping-row-cnsc-67-sa-11"></a>`CNSC-67-SA-11` | `CNSC-67` | `SA-11` | relates-to | Developer Testing and Evaluation |
| <a id="mapping-row-cnsc-68-sa-11"></a>`CNSC-68-SA-11` | `CNSC-68` | `SA-11` | relates-to | Developer Testing and Evaluation |
| <a id="mapping-row-cnsc-69-sa-11"></a>`CNSC-69-SA-11` | `CNSC-69` | `SA-11` | relates-to | Developer Testing and Evaluation |
| <a id="mapping-row-cnsc-7-ia-2-6"></a>`CNSC-7-IA-2-6` | `CNSC-7` | `IA-2(6)` | relates-to | Identification and Authentication (Organizational Users) |
| <a id="mapping-row-cnsc-70-sa-11"></a>`CNSC-70-SA-11` | `CNSC-70` | `SA-11` | relates-to | Developer Testing and Evaluation |
| <a id="mapping-row-cnsc-71-sa-11-4"></a>`CNSC-71-SA-11-4` | `CNSC-71` | `SA-11(4)` | relates-to | Developer Testing and Evaluation \| Manual Code Reviews |
| <a id="mapping-row-cnsc-73-sa-11"></a>`CNSC-73-SA-11` | `CNSC-73` | `SA-11` | relates-to | Developer Testing and Evaluation |
| <a id="mapping-row-cnsc-74-sa-11"></a>`CNSC-74-SA-11` | `CNSC-74` | `SA-11` | relates-to | Developer Testing and Evaluation |
| <a id="mapping-row-cnsc-75-sa-11"></a>`CNSC-75-SA-11` | `CNSC-75` | `SA-11` | relates-to | Developer Testing and Evaluation |
| <a id="mapping-row-cnsc-76-sa-3-1"></a>`CNSC-76-SA-3-1` | `CNSC-76` | `SA-3(1)` | relates-to | System Development Life Cycle \| Manage preproduction environment |
| <a id="mapping-row-cnsc-77-sc-39"></a>`CNSC-77-SC-39` | `CNSC-77` | `SC-39` | relates-to | Process Isolation |
| <a id="mapping-row-cnsc-78-sa-11-2"></a>`CNSC-78-SA-11-2` | `CNSC-78` | `SA-11(2)` | relates-to | Developer Testing and Evaluation \| Threat Modeling and Vulnerability Analyses |
| <a id="mapping-row-cnsc-8-si-4-2"></a>`CNSC-8-SI-4-2` | `CNSC-8` | `SI-4(2)` | relates-to | System Monitoring |
| <a id="mapping-row-cnsc-85-ra-5"></a>`CNSC-85-RA-5` | `CNSC-85` | `RA-5` | relates-to | Vulnerability Monitoring and Scanning |
| <a id="mapping-row-cnsc-86-sc-39"></a>`CNSC-86-SC-39` | `CNSC-86` | `SC-39` | relates-to | Process Isolation |
| <a id="mapping-row-cnsc-87-sc-39"></a>`CNSC-87-SC-39` | `CNSC-87` | `SC-39` | relates-to | Process Isolation |
| <a id="mapping-row-cnsc-88-sa-1"></a>`CNSC-88-SA-1` | `CNSC-88` | `SA-1` | relates-to | Policy and Procedures |
| <a id="mapping-row-cnsc-89-si-7"></a>`CNSC-89-SI-7` | `CNSC-89` | `SI-7` | relates-to | Software, Firmware, and Information Integrity |
| <a id="mapping-row-cnsc-9-ac-3-13"></a>`CNSC-9-AC-3-13` | `CNSC-9` | `AC-3(13)` | relates-to | Access Enforcement |
| <a id="mapping-row-cnsc-90-si-7"></a>`CNSC-90-SI-7` | `CNSC-90` | `SI-7` | relates-to | Software, Firmware, and Information Integrity |
| <a id="mapping-row-cnsc-91-ra-5"></a>`CNSC-91-RA-5` | `CNSC-91` | `RA-5` | relates-to | Vulnerability Monitoring and Scanning |
| <a id="mapping-row-cnsc-92-sa-1"></a>`CNSC-92-SA-1` | `CNSC-92` | `SA-1` | relates-to | Policy and Procedures |
| <a id="mapping-row-cnsc-93-sa-11-8"></a>`CNSC-93-SA-11-8` | `CNSC-93` | `SA-11(8)` | relates-to | Interactive Application Security Testing |
| <a id="mapping-row-cnsc-94-si-4"></a>`CNSC-94-SI-4` | `CNSC-94` | `SI-4` | relates-to | System Monitoring |
| <a id="mapping-row-cnsc-97-si-4"></a>`CNSC-97-SI-4` | `CNSC-97` | `SI-4` | relates-to | System Monitoring |
| <a id="mapping-row-cnsc-99-ca-8"></a>`CNSC-99-CA-8` | `CNSC-99` | `CA-8` | relates-to | Penetration Testing |
| <a id="mapping-row-cnsc-99-sa-11"></a>`CNSC-99-SA-11` | `CNSC-99` | `SA-11` | relates-to | Developer Testing and Evaluation |



## Document remarks

Auto-generated from YAML family files; 205 atomic guideline-to-control relationships.
