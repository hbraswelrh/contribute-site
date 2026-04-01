---
title: Cloud Native Security Controls Catalog — NIST mappings by family
sidebar_position: 8
toc_max_heading_level: 3
---

<!--
Auto-generated from cnsc-nist-800-53-mapping.yaml plus catalog guidelines (for family grouping).
Do not edit manually.

  go run cmd/catalog/main.go -md index.md
-->

# Cloud Native Security Controls Catalog

## Guideline families (NIST mappings)

**Jump to a section:**

[Access Control](#access) · [Compute](#compute) · [Deploy](#deploy) · [Develop](#develop) · [Distribute](#distribute) · [Securing Artefacts](#securing-artefacts) · [Securing Build Pipelines](#securing-build-pipelines) · [Securing Materials](#securing-materials) · [Securing the Source Code](#securing-the-source-code) · [Security Assurance](#security-assurance) · [Storage](#storage).

---


<a id="access"></a>

## Access Control

Guidelines for access control models and identity forwarding.



<a id="cnsc-1"></a>

### Runtime Secret Injection | `CNSC-1`
**Originating Document**: `Cloud Native Security Whitepaper v1.0`


**Objective**

Secrets are injected at runtime

**Guideline Mappings**

**`NIST-800-53`**

| NIST control | Remarks |
|--------------|---------|
| `IA-5(7)` | Authenticator Management |


---



<a id="cnsc-10"></a>

### ABAC and RBAC | `CNSC-10`
**Originating Document**: `Cloud Native Security Whitepaper v1.0`


**Objective**

ABAC and RBAC are used

**Guideline Mappings**

**`NIST-800-53`**

| NIST control | Remarks |
|--------------|---------|
| `AC-3(13)` | Access Enforcement |


---



<a id="cnsc-11"></a>

### Authorization and Identity Management | `CNSC-11`
**Originating Document**: `Cloud Native Security Whitepaper v1.0`


**Objective**

End user identity is capable of being accepted, consumed, and forwarded on for contextual or dynamic authorization

**Guideline Mappings**

**`NIST-800-53`**

| NIST control | Remarks |
|--------------|---------|
| `SC-7(19)` | Boundary Protection |


---



<a id="cnsc-12"></a>

### Cluster Authentication Management | `CNSC-12`
**Originating Document**:
- `Cloud Native Security Whitepaper v1.0`
- `Cloud Native Security Whitepaper v2.0`



**Objective**

All cluster and workloads operators are authenticated

**Guideline Mappings**

**`NIST-800-53`**

| NIST control | Remarks |
|--------------|---------|
| `IA-7` | Cryptographic Module Authentication |


---



<a id="cnsc-13"></a>

### Authentication Policy Management | `CNSC-13`
**Originating Document**:
- `Cloud Native Security Whitepaper v1.0`
- `Cloud Native Security Whitepaper v2.0`



**Objective**

Cluster and workloads operate actions are evaluated against access control policies governing context, purpose, and output

**Guideline Mappings**

**`NIST-800-53`**

| NIST control | Remarks |
|--------------|---------|
| `IA-7` | Cryptographic Module Authentication |


---



<a id="cnsc-14"></a>

### Multi-factor Authentication | `CNSC-14`
**Originating Document**: `Cloud Native Security Whitepaper v1.0`


**Objective**

Identity federation uses multi-factor authentication

**Guideline Mappings**

**`NIST-800-53`**

| NIST control | Remarks |
|--------------|---------|
| `IA-2(1)(2)` | Identification and Authentication (organizational Users) |


---



<a id="cnsc-15"></a>

### HSMs Protection of Cryptographic Secrets | `CNSC-15`
**Originating Document**:
- `Cloud Native Security Whitepaper v1.0`
- `Cloud Native Security Whitepaper v2.0`



**Objective**

HSMs are used to physically protect cryptographic secrets with an encryption key residing in the HSM

**Guideline Mappings**

**`NIST-800-53`**

| NIST control | Remarks |
|--------------|---------|
| `AC-4(4)` | Information Flow Enforcement |


---



<a id="cnsc-16"></a>

### Secrets Management | `CNSC-16`
**Originating Document**:
- `Cloud Native Security Whitepaper v1.0`
- `Cloud Native Security Whitepaper v2.0`



**Objective**

Secrets should have a short expiration period or time to live

**Guideline Mappings**

**`NIST-800-53`**

| NIST control | Remarks |
|--------------|---------|
| `SI-12` | Information Management and Retention |


---



<a id="cnsc-17"></a>

### Secrets Lifecycle Management | `CNSC-17`
**Originating Document**:
- `Cloud Native Security Whitepaper v1.0`
- `Cloud Native Security Whitepaper v2.0`



**Objective**

Time to live and expiration period on secrets is verified to prevent reuse

**Guideline Mappings**

**`NIST-800-53`**

| NIST control | Remarks |
|--------------|---------|
| `AC-16(3)` | Security and Privacy Attributes |


---



<a id="cnsc-18"></a>

### Secrets Management System | `CNSC-18`
**Originating Document**:
- `Cloud Native Security Whitepaper v1.0`
- `Cloud Native Security Whitepaper v2.0`



**Objective**

Secrets management systems are highly available

**Guideline Mappings**

**`NIST-800-53`**

| NIST control | Remarks |
|--------------|---------|
| `SC-12(1)` | Cryptographic Key Establishment and Management |


---



<a id="cnsc-19"></a>

### Secrets Rotation Management | `CNSC-19`
**Originating Document**:
- `Cloud Native Security Whitepaper v1.0`
- `Cloud Native Security Whitepaper v2.0`



**Objective**

Long-lived secrets adhere to periodic rotation and revocation

**Guideline Mappings**

**`NIST-800-53`**

| NIST control | Remarks |
|--------------|---------|
| `SI-12` | Information Management and Retention |


---



<a id="cnsc-2"></a>

### Mutual Authentication | `CNSC-2`
**Originating Document**: `Cloud Native Security Whitepaper v1.0`


**Objective**

Applications and workloads are explicitly authorized to communicate with each other using mutual authentication

**Guideline Mappings**

**`NIST-800-53`**

| NIST control | Remarks |
|--------------|---------|
| `IA-9` | NIST SP 800-53 control mapped from CNSC-2 (Mutual Authentication); family: Access. |


---



<a id="cnsc-20"></a>

### Secrets Protection | `CNSC-20`
**Originating Document**:
- `Cloud Native Security Whitepaper v1.0`
- `Cloud Native Security Whitepaper v2.0`



**Objective**

Secrets are distributed through secured communication channels protected commensurate with the level of access or data they are protecting

**Guideline Mappings**

**`NIST-800-53`**

| NIST control | Remarks |
|--------------|---------|
| `AC-16` | Security and Privacy Attributes |


---



<a id="cnsc-208"></a>

### Multi-factor Authentication | `CNSC-208`
**Originating Document**: `Cloud Native Security Whitepaper v2.0`


**Objective**

Identity federation uses multi-factor authentication for human users

**Guideline Mappings**

**`NIST-800-53`**

| NIST control | Remarks |
|--------------|---------|
| `IA-2(1)(2)` | Identification and Authentication (organizational Users) |


---



<a id="cnsc-21"></a>

### Secret Injection Lifecycle | `CNSC-21`
**Originating Document**:
- `Cloud Native Security Whitepaper v1.0`
- `Cloud Native Security Whitepaper v2.0`



**Objective**

Secrets injected at runtime are masked or dropped from logs, audit, or system dumps

**Guideline Mappings**

**`NIST-800-53`**

| NIST control | Remarks |
|--------------|---------|
| `AU-9(3)` | Protection of Audit Information |


---



<a id="cnsc-3"></a>

### Cryptographic Key Management | `CNSC-3`
**Originating Document**: `Cloud Native Security Whitepaper v1.0`


**Objective**

Keys are rotated frequently

**Guideline Mappings**

**`NIST-800-53`**

| NIST control | Remarks |
|--------------|---------|
| `SC-12` | Cryptographic Key Establishment and Management |


---



<a id="cnsc-4"></a>

### Cryptographic Key Lifecycle Management | `CNSC-4`
**Originating Document**: `Cloud Native Security Whitepaper v1.0`


**Objective**

Key lifespan is short

**Guideline Mappings**

**`NIST-800-53`**

| NIST control | Remarks |
|--------------|---------|
| `SC-12(3)` | Cryptographic Key Establishment and Management |


---



<a id="cnsc-5"></a>

### Sensitive Credential Management | `CNSC-5`
**Originating Document**: `Cloud Native Security Whitepaper v1.0`


**Objective**

Credentials and keys protecting sensitive workloads (health/finance/etc) are generated and managed independent of a cloud service provider

**Guideline Mappings**

**`NIST-800-53`**

| NIST control | Remarks |
|--------------|---------|
| `IA-2(12)` | Identification and Authentication (Organizational Users) |


---



<a id="cnsc-6"></a>

### Identification and Authentication | `CNSC-6`
**Originating Document**: `Cloud Native Security Whitepaper v1.0`


**Objective**

Authentication and authorization are determined independently

**Guideline Mappings**

**`NIST-800-53`**

| NIST control | Remarks |
|--------------|---------|
| `IA-2(6)` | Identification and Authentication (Organizational Users) |


---



<a id="cnsc-7"></a>

### Authentication and Authorization Enforcement | `CNSC-7`
**Originating Document**: `Cloud Native Security Whitepaper v1.0`


**Objective**

Authentication and authorization are enforced independently

**Guideline Mappings**

**`NIST-800-53`**

| NIST control | Remarks |
|--------------|---------|
| `IA-2(6)` | Identification and Authentication (Organizational Users) |


---



<a id="cnsc-8"></a>

### Continuous System Monitoring | `CNSC-8`
**Originating Document**: `Cloud Native Security Whitepaper v1.0`


**Objective**

Access control and file permissions are updated in real-time

**Guideline Mappings**

**`NIST-800-53`**

| NIST control | Remarks |
|--------------|---------|
| `SI-4(2)` | System Monitoring |


---



<a id="cnsc-9"></a>

### Privileged-based Authorization | `CNSC-9`
**Originating Document**: `Cloud Native Security Whitepaper v1.0`


**Objective**

Authorization for workloads is granted based on attributes and roles/permissions previously assigned

**Guideline Mappings**

**`NIST-800-53`**

| NIST control | Remarks |
|--------------|---------|
| `AC-3(13)` | Access Enforcement |


---




<a id="compute"></a>

## Compute

Guidelines for securing compute infrastructure including bootstrapping, isolation, monitoring, and runtime security.



<a id="cnsc-22"></a>

### Compute Bootstrapping Verification | `CNSC-22`
**Originating Document**:
- `Cloud Native Security Whitepaper v1.0`
- `Cloud Native Security Whitepaper v2.0`



**Objective**

Bootstrapping is employed to verify correct physical and logical location of compute

**Guideline Mappings**

**`NIST-800-53`**

| NIST control | Remarks |
|--------------|---------|
| `SI-7(9)` | Software, Firmware, and Information Integrity |


---



<a id="cnsc-23"></a>

### Boundary Management | `CNSC-23`
**Originating Document**:
- `Cloud Native Security Whitepaper v1.0`
- `Cloud Native Security Whitepaper v2.0`



**Objective**

Disparate data sensitive workloads are not run on the same host OS kernel

**Guideline Mappings**

**`NIST-800-53`**

| NIST control | Remarks |
|--------------|---------|
| `SC-7` | Boundary Protection |


---



<a id="cnsc-233"></a>

### Data Trust Management | `CNSC-233`
**Originating Document**: `Cloud Native Security Whitepaper v2.0`


**Objective**

Use a service mesh that eliminates implicit trust through data-in-motion protection (i.e. confidentiality, integrity, authentication, authorization)

**Guideline Mappings**

**`NIST-800-53`**

| NIST control | Remarks |
|--------------|---------|
| `SC-7` | Boundary Protection |


---



<a id="cnsc-24"></a>

### Runtime Configuration Monitoring | `CNSC-24`
**Originating Document**:
- `Cloud Native Security Whitepaper v1.0`
- `Cloud Native Security Whitepaper v2.0`



**Objective**

Monitor and detect any changes to the initial configurations made in runtime

**Guideline Mappings**

**`NIST-800-53`**

| NIST control | Remarks |
|--------------|---------|
| `CM-2(2)`<br />`CM-3(7)` | <ul><li>Baseline Configuration, Automation Support for Accuracy and Currency</li><li>Configuration Change Control, Review System Changes</li></ul> |


---



<a id="cnsc-25"></a>

### API Auditing Implementation | `CNSC-25`
**Originating Document**:
- `Cloud Native Security Whitepaper v1.0`
- `Cloud Native Security Whitepaper v2.0`



**Objective**

API auditing is enabled with a filter for a specific set of API Groups or verbs

**Guideline Mappings**

**`NIST-800-53`**

| NIST control | Remarks |
|--------------|---------|
| `AU-2` | Event Logging |


---



<a id="cnsc-26"></a>

### Operating System Configuration | `CNSC-26`
**Originating Document**:
- `Cloud Native Security Whitepaper v1.0`
- `Cloud Native Security Whitepaper v2.0`



**Objective**

Container specific operating systems are in use

**Guideline Mappings**

**`NIST-800-53`**

| NIST control | Remarks |
|--------------|---------|
| `CM-2`<br />`CM-7` | <ul><li>Baseline Configuration</li><li>Least Functionality</li></ul> |


---



<a id="cnsc-27"></a>

### Trust Implementation | `CNSC-27`
**Originating Document**:
- `Cloud Native Security Whitepaper v1.0`
- `Cloud Native Security Whitepaper v2.0`



**Objective**

The hardware root of trust is based in a Trusted Platform Module (TPM) or virtual TPM (vTPM)

**Guideline Mappings**

**`NIST-800-53`**

| NIST control | Remarks |
|--------------|---------|
| `SI-7` | Software, Firmware, and Information Integrity |


---



<a id="cnsc-28"></a>

### Least Privilege | `CNSC-28`
**Originating Document**:
- `Cloud Native Security Whitepaper v1.0`
- `Cloud Native Security Whitepaper v2.0`



**Objective**

Minimize administrative access to the control plane

**Guideline Mappings**

**`NIST-800-53`**

| NIST control | Remarks |
|--------------|---------|
| `AC-6` | Least Privilege |


---



<a id="cnsc-29"></a>

### Resource Control Management | `CNSC-29`
**Originating Document**:
- `Cloud Native Security Whitepaper v1.0`
- `Cloud Native Security Whitepaper v2.0`



**Objective**

Object level and resource requests and limits are controlled through cgroups

**Guideline Mappings**

**`NIST-800-53`**

| NIST control | Remarks |
|--------------|---------|
| `SI-7(16)`<br />`SI-7(17)` | <ul><li>Software, Firmware, and Information Integrity, Time Limit on Process Execution Without Supervision</li><li>Software, Firmware, and Information Integrity, Runtime Application Self-protection</li></ul> |


---



<a id="cnsc-30"></a>

### System Alert Monitoring | `CNSC-30`
**Originating Document**:
- `Cloud Native Security Whitepaper v1.0`
- `Cloud Native Security Whitepaper v2.0`



**Objective**

Systems processing alerts are periodically tuned for false positives

**Guideline Mappings**

**`NIST-800-53`**

| NIST control | Remarks |
|--------------|---------|
| `SI-4(13)` | System Monitoring, Analyze Traffic and Event Patterns |


---



<a id="cnsc-31"></a>

### Control Plane Configuration | `CNSC-31`
**Originating Document**:
- `Cloud Native Security Whitepaper v1.0`
- `Cloud Native Security Whitepaper v2.0`



**Objective**

All orchestrator control plane components are configured to communicate via mutual authentication and certificate validation with a periodically rotated certificate

**Guideline Mappings**

**`NIST-800-53`**

| NIST control | Remarks |
|--------------|---------|
| `AC-3` | Access Enforcement |


---



<a id="cnsc-32"></a>

### Baseline Configured Functionality | `CNSC-32`
**Originating Document**:
- `Cloud Native Security Whitepaper v1.0`
- `Cloud Native Security Whitepaper v2.0`



**Objective**

Only sanctioned capabilities and system calls (e.g. seccomp filters), are allowed to execute or be invoked in a container by the host operating system

**Guideline Mappings**

**`NIST-800-53`**

| NIST control | Remarks |
|--------------|---------|
| `CM-2`<br />`CM-7` | <ul><li>Baseline Configuration</li><li>Least Functionality</li></ul> |


---



<a id="cnsc-33"></a>

### Critical Change Management | `CNSC-33`
**Originating Document**:
- `Cloud Native Security Whitepaper v1.0`
- `Cloud Native Security Whitepaper v2.0`



**Objective**

Changes to critical mount points and files are prevented, monitored, and alerted

**Guideline Mappings**

**`NIST-800-53`**

| NIST control | Remarks |
|--------------|---------|
| `CM-5` | Access Restrictions for Change |


---



<a id="cnsc-34"></a>

### Runtime Configuration for Change Management | `CNSC-34`
**Originating Document**:
- `Cloud Native Security Whitepaper v1.0`
- `Cloud Native Security Whitepaper v2.0`



**Objective**

Runtime configuration control prevents changes to binaries, certificates, and remote access configurations

**Guideline Mappings**

**`NIST-800-53`**

| NIST control | Remarks |
|--------------|---------|
| `CM-5` | Access Restrictions for Change |


---



<a id="cnsc-35"></a>

### Runtime Boundary Protection Management | `CNSC-35`
**Originating Document**:
- `Cloud Native Security Whitepaper v1.0`
- `Cloud Native Security Whitepaper v2.0`



**Objective**

Runtime configuration prevents ingress and egress network access for containers to only what is required to operate

**Guideline Mappings**

**`NIST-800-53`**

| NIST control | Remarks |
|--------------|---------|
| `SC-7` | Boundary Protection |


---



<a id="cnsc-36"></a>

### Boundary Protection Management | `CNSC-36`
**Originating Document**:
- `Cloud Native Security Whitepaper v1.0`
- `Cloud Native Security Whitepaper v2.0`



**Objective**

Policies are defined that restrict communications to only occur between sanctioned microservice pairs

**Guideline Mappings**

**`NIST-800-53`**

| NIST control | Remarks |
|--------------|---------|
| `SC-7` | Boundary Protection |


---



<a id="cnsc-37"></a>

### Policy Enforcement Management | `CNSC-37`
**Originating Document**:
- `Cloud Native Security Whitepaper v1.0`
- `Cloud Native Security Whitepaper v2.0`



**Objective**

Use a policy agent to control and enforce authorized, signed container images

**Guideline Mappings**

**`NIST-800-53`**

| NIST control | Remarks |
|--------------|---------|
| `CM-5` | Access Restrictions for Change |


---



<a id="cnsc-38"></a>

### Policy Enforcement Management | `CNSC-38`
**Originating Document**:
- `Cloud Native Security Whitepaper v1.0`
- `Cloud Native Security Whitepaper v2.0`



**Objective**

Use a policy agent to control provenance assurance for operational workloads

**Guideline Mappings**

**`NIST-800-53`**

| NIST control | Remarks |
|--------------|---------|
| `CM-5` | Access Restrictions for Change |


---



<a id="cnsc-39"></a>

### Data Trust Management | `CNSC-39`
**Originating Document**: `Cloud Native Security Whitepaper v1.0`


**Objective**

Use a service mesh that eliminates implicit trust through data-in-motion encryption (data in transit)

**Guideline Mappings**

**`NIST-800-53`**

| NIST control | Remarks |
|--------------|---------|
| `SC-7` | Boundary Protection |


---



<a id="cnsc-40"></a>

### System Monitoring Components | `CNSC-40`
**Originating Document**:
- `Cloud Native Security Whitepaper v1.0`
- `Cloud Native Security Whitepaper v2.0`



**Objective**

Use components that detect, track, aggregate and report system calls and network traffic from a container

**Guideline Mappings**

**`NIST-800-53`**

| NIST control | Remarks |
|--------------|---------|
| `SI-4` | System Monitoring |


---



<a id="cnsc-41"></a>

### Dynamic Workload Scanning | `CNSC-41`
**Originating Document**:
- `Cloud Native Security Whitepaper v1.0`
- `Cloud Native Security Whitepaper v2.0`



**Objective**

Workloads should be dynamically scanned to detect malicious or insidious behavior for which no known occurrence yet exists

**Guideline Mappings**

**`NIST-800-53`**

| NIST control | Remarks |
|--------------|---------|
| `SI-3` | Malicious Code Protection |


---



<a id="cnsc-42"></a>

### Continuous Monitoring and Scanning | `CNSC-42`
**Originating Document**:
- `Cloud Native Security Whitepaper v1.0`
- `Cloud Native Security Whitepaper v2.0`



**Objective**

Environments are continuously scanned to detect new vulnerabilities in workloads

**Guideline Mappings**

**`NIST-800-53`**

| NIST control | Remarks |
|--------------|---------|
| `RA-5` | Vulnerability Monitoring and Scanning |


---



<a id="cnsc-43"></a>

### Audit Event Logging | `CNSC-43`
**Originating Document**:
- `Cloud Native Security Whitepaper v1.0`
- `Cloud Native Security Whitepaper v2.0`



**Objective**

Actionable audit events are generated that correlate/contextualize data from logs into "information" that can drive decision trees/incident response

**Guideline Mappings**

**`NIST-800-53`**

| NIST control | Remarks |
|--------------|---------|
| `AU-3` | Content of Audit Records |


---



<a id="cnsc-44"></a>

### Privilege Management | `CNSC-44`
**Originating Document**:
- `Cloud Native Security Whitepaper v1.0`
- `Cloud Native Security Whitepaper v2.0`



**Objective**

Segregation of duties and the principle of least privilege is enforced

**Guideline Mappings**

**`NIST-800-53`**

| NIST control | Remarks |
|--------------|---------|
| `AC-6` | Least Privilege |


---



<a id="cnsc-45"></a>

### Information Integrity | `CNSC-45`
**Originating Document**:
- `Cloud Native Security Whitepaper v1.0`
- `Cloud Native Security Whitepaper v2.0`



**Objective**

Non-compliant violations are detected based on a pre-configured set of rules defined by the organization's policies

**Guideline Mappings**

**`NIST-800-53`**

| NIST control | Remarks |
|--------------|---------|
| `SI-7` | Software, Firmware, and Information Integrity |


---



<a id="cnsc-46"></a>

### Key Management Storage | `CNSC-46`
**Originating Document**:
- `Cloud Native Security Whitepaper v1.0`
- `Cloud Native Security Whitepaper v2.0`



**Objective**

Native secret stores encrypt with keys from an external Key Management Store (KMS)

**Guideline Mappings**

**`NIST-800-53`**

| NIST control | Remarks |
|--------------|---------|
| `SC-12(3)` | Systems and Communication Protection |


---



<a id="cnsc-47"></a>

### Secret Storage Configuration | `CNSC-47`
**Originating Document**:
- `Cloud Native Security Whitepaper v1.0`
- `Cloud Native Security Whitepaper v2.0`



**Objective**

Native secret stores are not configured for base64 encoding or stored in clear-text in the key-value store by default

**Guideline Mappings**

**`NIST-800-53`**

| NIST control | Remarks |
|--------------|---------|
| `SC-12(3)` | Systems and Communication Protection |


---



<a id="cnsc-48"></a>

### System Monitoring | `CNSC-48`
**Originating Document**:
- `Cloud Native Security Whitepaper v1.0`
- `Cloud Native Security Whitepaper v2.0`



**Objective**

Network traffic to malicious domains is detected and denied

**Guideline Mappings**

**`NIST-800-53`**

| NIST control | Remarks |
|--------------|---------|
| `SI-4` | System Monitoring |


---



<a id="cnsc-49"></a>

### Sensitive Data Encryption | `CNSC-49`
**Originating Document**:
- `Cloud Native Security Whitepaper v1.0`
- `Cloud Native Security Whitepaper v2.0`



**Objective**

Use encrypted containers for sensitive sources, methods, and data

**Guideline Mappings**

**`NIST-800-53`**

| NIST control | Remarks |
|--------------|---------|
| `SC-28` | Protection of Information at Rest |


---



<a id="cnsc-50"></a>

### SBOM Management | `CNSC-50`
**Originating Document**:
- `Cloud Native Security Whitepaper v1.0`
- `Cloud Native Security Whitepaper v2.0`



**Objective**

Use SBOMs to identify current deployments of vulnerable libraries, dependencies, and packages

**Guideline Mappings**

**`NIST-800-53`**

| NIST control | Remarks |
|--------------|---------|
| `CM-8` | System Component Inventory |


---



<a id="cnsc-51"></a>

### Functionality Management | `CNSC-51`
**Originating Document**:
- `Cloud Native Security Whitepaper v1.0`
- `Cloud Native Security Whitepaper v2.0`



**Objective**

Processes must execute only functions explicitly defined in an allow list

**Guideline Mappings**

**`NIST-800-53`**

| NIST control | Remarks |
|--------------|---------|
| `CM-2`<br />`CM-7` | <ul><li>Baseline Configuration</li><li>Least Functionality</li></ul> |


---



<a id="cnsc-52"></a>

### Access and Change Restrictions | `CNSC-52`
**Originating Document**:
- `Cloud Native Security Whitepaper v1.0`
- `Cloud Native Security Whitepaper v2.0`



**Objective**

Functions are not be allowed to make changes to critical file system mount points

**Guideline Mappings**

**`NIST-800-53`**

| NIST control | Remarks |
|--------------|---------|
| `CM-5` | Access Restrictions for Change |


---



<a id="cnsc-53"></a>

### Access Configuration | `CNSC-53`
**Originating Document**:
- `Cloud Native Security Whitepaper v1.0`
- `Cloud Native Security Whitepaper v2.0`



**Objective**

Function access is only permitted to sanctioned services

**Guideline Mappings**

**`NIST-800-53`**

| NIST control | Remarks |
|--------------|---------|
| `CM-2`<br />`CM-7` | <ul><li>Baseline Configuration</li><li>Least Functionality</li></ul> |


---



<a id="cnsc-54"></a>

### System Monitoring | `CNSC-54`
**Originating Document**:
- `Cloud Native Security Whitepaper v1.0`
- `Cloud Native Security Whitepaper v2.0`



**Objective**

Egress network connection is monitored to detect and prevent access to C&C (command and control) and other malicious network domains

**Guideline Mappings**

**`NIST-800-53`**

| NIST control | Remarks |
|--------------|---------|
| `SI-4` | System Monitoring |


---



<a id="cnsc-55"></a>

### System Monitoring Management | `CNSC-55`
**Originating Document**:
- `Cloud Native Security Whitepaper v1.0`
- `Cloud Native Security Whitepaper v2.0`



**Objective**

Ingress network inspection is employed detect and remove malicious payloads and commands

**Guideline Mappings**

**`NIST-800-53`**

| NIST control | Remarks |
|--------------|---------|
| `SI-4` | System Monitoring |


---



<a id="cnsc-56"></a>

### System Component Isolation | `CNSC-56`
**Originating Document**:
- `Cloud Native Security Whitepaper v1.0`
- `Cloud Native Security Whitepaper v2.0`



**Objective**

Serverless functions are run in tenant-based resource or performance isolation for similar data classifications

**Guideline Mappings**

**`NIST-800-53`**

| NIST control | Remarks |
|--------------|---------|
| `SC-7(21)` | Boundary Protection, Isolation of System Components |


---




<a id="deploy"></a>

## Deploy

Guidelines for securing software deployments, including artifact verification, runtime policy enforcement, freshness validation, update management, logging, and incident response.



<a id="cnsc-170"></a>

### Artifact Verification | `CNSC-170`
**Originating Document**: `Software Supply Chain Best Practices v1.0`


**Objective**

Clients can perform verification of artefacts and associated metadata

**Guideline Mappings**

**`NIST-800-53`**

| NIST control | Remarks |
|--------------|---------|
| `SI-7` | Software, Firmware, and Information Integrity |


---



<a id="cnsc-171"></a>

### Freshness Verification | `CNSC-171`
**Originating Document**: `Software Supply Chain Best Practices v1.0`


**Objective**

Clients can verify the freshness of files

**Guideline Mappings**

**`NIST-800-53`**

| NIST control | Remarks |
|--------------|---------|
| `SI-7` | Software, Firmware, and Information Integrity |


---



<a id="cnsc-172"></a>

### Update Framework | `CNSC-172`
**Originating Document**: `Software Supply Chain Best Practices v1.0`


**Objective**

A framework is used for managing software updates

*No NIST SP 800-53 mapping row in this document for this guideline.*

---



<a id="cnsc-57"></a>

### Trust Confirmation | `CNSC-57`
**Originating Document**:
- `Cloud Native Security Whitepaper v1.0`
- `Cloud Native Security Whitepaper v2.0`



**Objective**

Trust confirmation verifies the image has a valid signature from an authorized source

**Guideline Mappings**

**`NIST-800-53`**

| NIST control | Remarks |
|--------------|---------|
| `SR-4(3)`<br />`SR-4(4)` | <ul><li>Provenance, Validate as Genuine and Not Altered</li><li>Provenance, Supply Chain Integrity - Pedigree</li></ul> |


---



<a id="cnsc-58"></a>

### Runtime Policy Enforcement | `CNSC-58`
**Originating Document**:
- `Cloud Native Security Whitepaper v1.0`
- `Cloud Native Security Whitepaper v2.0`



**Objective**

Image runtime policies are enforced prior to deployment

**Guideline Mappings**

**`NIST-800-53`**

| NIST control | Remarks |
|--------------|---------|
| `SI-7(17)` | Software, Firmware, and Information Integrity, Runtime Application Self-Protection |


---



<a id="cnsc-59"></a>

### Image Integrity Verification | `CNSC-59`
**Originating Document**:
- `Cloud Native Security Whitepaper v1.0`
- `Cloud Native Security Whitepaper v2.0`



**Objective**

Image integrity and signature are verified prior to deployment

**Guideline Mappings**

**`NIST-800-53`**

| NIST control | Remarks |
|--------------|---------|
| `SR-4(3)`<br />`SR-4(4)` | <ul><li>Provenance, Validate as Genuine and Not Altered</li><li>Provenance, Supply Chain Integrity - Pedigree</li></ul> |


---



<a id="cnsc-60"></a>

### Application Logging | `CNSC-60`
**Originating Document**:
- `Cloud Native Security Whitepaper v1.0`
- `Cloud Native Security Whitepaper v2.0`



**Objective**

Applications provide logs regarding authentication, authorization, actions, and failures

**Guideline Mappings**

**`NIST-800-53`**

| NIST control | Remarks |
|--------------|---------|
| `CM-3` | Configuration Change Control |


---



<a id="cnsc-61"></a>

### Forensics Integration | `CNSC-61`
**Originating Document**:
- `Cloud Native Security Whitepaper v1.0`
- `Cloud Native Security Whitepaper v2.0`



**Objective**

Forensics capabilities are integrated into an incident response plan and procedures

*No NIST SP 800-53 mapping row in this document for this guideline.*

---



<a id="cnsc-62"></a>

### Behavioral Analysis | `CNSC-62`
**Originating Document**:
- `Cloud Native Security Whitepaper v1.0`
- `Cloud Native Security Whitepaper v2.0`



**Objective**

AI, ML, or statistical modeling are used for behavioural and heuristic environment analysis to detect unwanted activities

**Guideline Mappings**

**`NIST-800-53`**

| NIST control | Remarks |
|--------------|---------|
| `SI-3` | System and Information Integrity |


---




<a id="develop"></a>

## Develop

Guidelines for secure software development practices including environment segregation, testing, code review, and CI server hardening.



<a id="cnsc-195"></a>

### Secure Configuration Defaults | `CNSC-195`
**Originating Document**: `Cloud Native Security Whitepaper v2.0`


**Objective**

Secure configuration is implemented as the default state of the system

**Guideline Mappings**

**`NIST-800-53`**

| NIST control | Remarks |
|--------------|---------|
| `SA-8(23)` | Security and Privacy Engineering Principles, Secure Defaults |


---



<a id="cnsc-259"></a>

### Early Vulnerability Scanning | `CNSC-259`
**Originating Document**: `Cloud Native Security Whitepaper v2.0`


**Objective**

Integrate vulnerability and configuration scanning in both the IDE and at the CI system during pull request

**Guideline Mappings**

**`NIST-800-53`**

| NIST control | Remarks |
|--------------|---------|
| `SA-11(1)` | Developer Testing and Evaluation \| Static Code Analysis |


---



<a id="cnsc-265"></a>

### Code Review Requirements | `CNSC-265`
**Originating Document**: `Cloud Native Security Whitepaper v2.0`


**Objective**

Implement at least one other non-author reviewer/approver prior to merging

**Guideline Mappings**

**`NIST-800-53`**

| NIST control | Remarks |
|--------------|---------|
| `SA-11(4)` | Developer Testing and Evaluation \| Manual Code Reviews |


---



<a id="cnsc-271"></a>

### CI Server Isolation | `CNSC-271`
**Originating Document**: `Cloud Native Security Whitepaper v2.0`


**Objective**

Continuous integration server is isolated and hardened

**Guideline Mappings**

**`NIST-800-53`**

| NIST control | Remarks |
|--------------|---------|
| `SC-39` | Process Isolation |


---



<a id="cnsc-63"></a>

### Production Environment | `CNSC-63`
**Originating Document**:
- `Cloud Native Security Whitepaper v1.0`
- `Cloud Native Security Whitepaper v2.0`



**Objective**

A dedicated production environment is established

**Guideline Mappings**

**`NIST-800-53`**

| NIST control | Remarks |
|--------------|---------|
| `SA-3(1)` | System Development Life Cycle \| Manage preproduction environment |


---



<a id="cnsc-64"></a>

### Dynamic Deployments | `CNSC-64`
**Originating Document**:
- `Cloud Native Security Whitepaper v1.0`
- `Cloud Native Security Whitepaper v2.0`



**Objective**

Dynamic deployments are leveraged for safer releases

**Guideline Mappings**

**`NIST-800-53`**

| NIST control | Remarks |
|--------------|---------|
| `SA-8(31)` | Security and Privacy Engineering Principles \| Secure System Modification |


---



<a id="cnsc-65"></a>

### Early Vulnerability Scanning | `CNSC-65`
**Originating Document**: `Cloud Native Security Whitepaper v1.0`


**Objective**

Vulnerability and configuration scanning is integrated in the IDE or at the pull request

**Guideline Mappings**

**`NIST-800-53`**

| NIST control | Remarks |
|--------------|---------|
| `SA-11(1)` | Developer Testing and Evaluation \| Static Code Analysis |


---



<a id="cnsc-66"></a>

### Environment Segregation | `CNSC-66`
**Originating Document**:
- `Cloud Native Security Whitepaper v1.0`
- `Cloud Native Security Whitepaper v2.0`



**Objective**

Dedicated development, testing, and production environments are established

**Guideline Mappings**

**`NIST-800-53`**

| NIST control | Remarks |
|--------------|---------|
| `SA-15` | Development Process, Standards, and Tools |


---



<a id="cnsc-67"></a>

### Business-Critical Code Testing | `CNSC-67`
**Originating Document**:
- `Cloud Native Security Whitepaper v1.0`
- `Cloud Native Security Whitepaper v2.0`



**Objective**

Tests are built for business-critical code

**Guideline Mappings**

**`NIST-800-53`**

| NIST control | Remarks |
|--------------|---------|
| `SA-11` | Developer Testing and Evaluation |


---



<a id="cnsc-68"></a>

### Infrastructure Testing | `CNSC-68`
**Originating Document**:
- `Cloud Native Security Whitepaper v1.0`
- `Cloud Native Security Whitepaper v2.0`



**Objective**

Tests are built for business-critical infrastructure

**Guideline Mappings**

**`NIST-800-53`**

| NIST control | Remarks |
|--------------|---------|
| `SA-11` | Developer Testing and Evaluation |


---



<a id="cnsc-69"></a>

### Local Test Execution | `CNSC-69`
**Originating Document**:
- `Cloud Native Security Whitepaper v1.0`
- `Cloud Native Security Whitepaper v2.0`



**Objective**

Test suites are able to be run locally

**Guideline Mappings**

**`NIST-800-53`**

| NIST control | Remarks |
|--------------|---------|
| `SA-11` | Developer Testing and Evaluation |


---



<a id="cnsc-70"></a>

### Shared Test Execution | `CNSC-70`
**Originating Document**:
- `Cloud Native Security Whitepaper v1.0`
- `Cloud Native Security Whitepaper v2.0`



**Objective**

Test suites are available to run in a shared environment

**Guideline Mappings**

**`NIST-800-53`**

| NIST control | Remarks |
|--------------|---------|
| `SA-11` | Developer Testing and Evaluation |


---



<a id="cnsc-71"></a>

### Code Review Requirements | `CNSC-71`
**Originating Document**: `Cloud Native Security Whitepaper v1.0`


**Objective**

Two non-author reviewers or approvers are required prior to merging

**Guideline Mappings**

**`NIST-800-53`**

| NIST control | Remarks |
|--------------|---------|
| `SA-11(4)` | Developer Testing and Evaluation \| Manual Code Reviews |


---



<a id="cnsc-72"></a>

### Code Quality | `CNSC-72`
**Originating Document**:
- `Cloud Native Security Whitepaper v1.0`
- `Cloud Native Security Whitepaper v2.0`



**Objective**

Code is clean and well commented

*No NIST SP 800-53 mapping row in this document for this guideline.*

---



<a id="cnsc-73"></a>

### Full Infrastructure Testing | `CNSC-73`
**Originating Document**:
- `Cloud Native Security Whitepaper v1.0`
- `Cloud Native Security Whitepaper v2.0`



**Objective**

Full infrastructure tests are used

**Guideline Mappings**

**`NIST-800-53`**

| NIST control | Remarks |
|--------------|---------|
| `SA-11` | Developer Testing and Evaluation |


---



<a id="cnsc-74"></a>

### Regression Testing | `CNSC-74`
**Originating Document**:
- `Cloud Native Security Whitepaper v1.0`
- `Cloud Native Security Whitepaper v2.0`



**Objective**

Regression tests are used

**Guideline Mappings**

**`NIST-800-53`**

| NIST control | Remarks |
|--------------|---------|
| `SA-11` | Developer Testing and Evaluation |


---



<a id="cnsc-75"></a>

### Security Regression Testing | `CNSC-75`
**Originating Document**:
- `Cloud Native Security Whitepaper v1.0`
- `Cloud Native Security Whitepaper v2.0`



**Objective**

Test suites are updated against new and emerging threats and developed into security regression tests

**Guideline Mappings**

**`NIST-800-53`**

| NIST control | Remarks |
|--------------|---------|
| `SA-11` | Developer Testing and Evaluation |


---



<a id="cnsc-76"></a>

### Testing Environment | `CNSC-76`
**Originating Document**:
- `Cloud Native Security Whitepaper v1.0`
- `Cloud Native Security Whitepaper v2.0`



**Objective**

A dedicated testing environment is established

**Guideline Mappings**

**`NIST-800-53`**

| NIST control | Remarks |
|--------------|---------|
| `SA-3(1)` | System Development Life Cycle \| Manage preproduction environment |


---



<a id="cnsc-77"></a>

### CI Server Isolation | `CNSC-77`
**Originating Document**: `Cloud Native Security Whitepaper v1.0`


**Objective**

Continuous integration server is isolated

**Guideline Mappings**

**`NIST-800-53`**

| NIST control | Remarks |
|--------------|---------|
| `SC-39` | Process Isolation |


---



<a id="cnsc-78"></a>

### Threat-Informed Test Development | `CNSC-78`
**Originating Document**:
- `Cloud Native Security Whitepaper v1.0`
- `Cloud Native Security Whitepaper v2.0`



**Objective**

Threat model results are used to determine ROI for test development

**Guideline Mappings**

**`NIST-800-53`**

| NIST control | Remarks |
|--------------|---------|
| `SA-11(2)` | Developer Testing and Evaluation \| Threat Modeling and Vulnerability Analyses |


---




<a id="distribute"></a>

## Distribute

Guidelines for secure distribution of container images, packages, and artifacts including signing, scanning, and registry security.



<a id="cnsc-100"></a>

### Registry Authentication | `CNSC-100`
**Originating Document**:
- `Cloud Native Security Whitepaper v1.0`
- `Cloud Native Security Whitepaper v2.0`



**Objective**

Registries require mutually authenticated TLS for all registry connections

**Guideline Mappings**

**`NIST-800-53`**

| NIST control | Remarks |
|--------------|---------|
| `IA-3(1)` | Cryptographic Bidirectional Authentication |


---



<a id="cnsc-101"></a>

### Image Signing | `CNSC-101`
**Originating Document**:
- `Cloud Native Security Whitepaper v1.0`
- `Cloud Native Security Whitepaper v2.0`



**Objective**

Image and metadata are signed

**Guideline Mappings**

**`NIST-800-53`**

| NIST control | Remarks |
|--------------|---------|
| `SI-7` | Software, Firmware, and Information Integrity |


---



<a id="cnsc-102"></a>

### Configuration Signing | `CNSC-102`
**Originating Document**: `Cloud Native Security Whitepaper v1.0`


**Objective**

Configuration is signed

**Guideline Mappings**

**`NIST-800-53`**

| NIST control | Remarks |
|--------------|---------|
| `SI-7` | Software, Firmware, and Information Integrity |


---



<a id="cnsc-103"></a>

### Package Signing | `CNSC-103`
**Originating Document**: `Cloud Native Security Whitepaper v1.0`


**Objective**

Package is signed

**Guideline Mappings**

**`NIST-800-53`**

| NIST control | Remarks |
|--------------|---------|
| `SI-7` | Software, Firmware, and Information Integrity |


---



<a id="cnsc-104"></a>

### Image Integrity Validation | `CNSC-104`
**Originating Document**:
- `Cloud Native Security Whitepaper v1.0`
- `Cloud Native Security Whitepaper v2.0`



**Objective**

Integrity of images is validated

**Guideline Mappings**

**`NIST-800-53`**

| NIST control | Remarks |
|--------------|---------|
| `SI-7` | System and Information Integrity |


---



<a id="cnsc-105"></a>

### Image Vulnerability Scanning | `CNSC-105`
**Originating Document**:
- `Cloud Native Security Whitepaper v1.0`
- `Cloud Native Security Whitepaper v2.0`



**Objective**

Images are scanned for vulnerabilities and malware

**Guideline Mappings**

**`NIST-800-53`**

| NIST control | Remarks |
|--------------|---------|
| `RA-5`<br />`SA-3` | <ul><li>Vulnerability Monitoring and Scanning</li><li>System Development Life Cycle</li></ul> |


---



<a id="cnsc-106"></a>

### Key Revocation | `CNSC-106`
**Originating Document**:
- `Cloud Native Security Whitepaper v1.0`
- `Cloud Native Security Whitepaper v2.0`



**Objective**

Image signing key revocation is enabled in the event of compromise

**Guideline Mappings**

**`NIST-800-53`**

| NIST control | Remarks |
|--------------|---------|
| `SI-7` | System and Information Integrity |


---



<a id="cnsc-107"></a>

### Security Update Prioritization | `CNSC-107`
**Originating Document**:
- `Cloud Native Security Whitepaper v1.0`
- `Cloud Native Security Whitepaper v2.0`



**Objective**

Security updates are prioritized

**Guideline Mappings**

**`NIST-800-53`**

| NIST control | Remarks |
|--------------|---------|
| `SI-2(3)` | System and Information Integrity |


---



<a id="cnsc-108"></a>

### Credential Protection | `CNSC-108`
**Originating Document**: `Cloud Native Security Whitepaper v1.0`


**Objective**

HSMs or credential managers are used for protecting credentials

**Guideline Mappings**

**`NIST-800-53`**

| NIST control | Remarks |
|--------------|---------|
| `SC-12(3)` | Systems and Communication Protection |


---



<a id="cnsc-109"></a>

### Scanning Remediation | `CNSC-109`
**Originating Document**:
- `Cloud Native Security Whitepaper v1.0`
- `Cloud Native Security Whitepaper v2.0`



**Objective**

Container image scanning findings are acted upon

**Guideline Mappings**

**`NIST-800-53`**

| NIST control | Remarks |
|--------------|---------|
| `SI-2(3)` | System and Information Integrity |


---



<a id="cnsc-110"></a>

### Compliance Enforcement | `CNSC-110`
**Originating Document**:
- `Cloud Native Security Whitepaper v1.0`
- `Cloud Native Security Whitepaper v2.0`



**Objective**

Organizational compliance rules are enforced

**Guideline Mappings**

**`NIST-800-53`**

| NIST control | Remarks |
|--------------|---------|
| `PL-1` | Policy and Procedures |


---



<a id="cnsc-111"></a>

### Infrastructure Hardening | `CNSC-111`
**Originating Document**:
- `Cloud Native Security Whitepaper v1.0`
- `Cloud Native Security Whitepaper v2.0`



**Objective**

Incremental hardening of the infrastructure is employed

*No NIST SP 800-53 mapping row in this document for this guideline.*

---



<a id="cnsc-112"></a>

### Public Registry Access Control | `CNSC-112`
**Originating Document**:
- `Cloud Native Security Whitepaper v1.0`
- `Cloud Native Security Whitepaper v2.0`



**Objective**

Pulls from public registries are controlled and only from authorized engineers or internal registries

**Guideline Mappings**

**`NIST-800-53`**

| NIST control | Remarks |
|--------------|---------|
| `AC-6(3)` | Least Privilege |


---



<a id="cnsc-113"></a>

### Image Encryption Management | `CNSC-113`
**Originating Document**:
- `Cloud Native Security Whitepaper v1.0`
- `Cloud Native Security Whitepaper v2.0`



**Objective**

Image encryption is coupled with key management attestation and/or authorization and credential distribution

**Guideline Mappings**

**`NIST-800-53`**

| NIST control | Remarks |
|--------------|---------|
| `SC-12(2)`<br />`SC-12(3)` | <ul><li>Cryptographic Key Establishment and Management \| Symmetric and Asymmetric Keys</li><li>Cryptographic Key Establishment and Management \| Symmetric and Asymmetric Keys</li></ul> |


---



<a id="cnsc-114"></a>

### Risk-Based Remediation Prioritization | `CNSC-114`
**Originating Document**:
- `Cloud Native Security Whitepaper v1.0`
- `Cloud Native Security Whitepaper v2.0`



**Objective**

At-risk applications are prioritized for remediation by exploit maturity and vulnerable path presence in addition to the CVSS score

**Guideline Mappings**

**`NIST-800-53`**

| NIST control | Remarks |
|--------------|---------|
| `SI-2(3)` | System and Information Integrity |


---



<a id="cnsc-274"></a>

### Signing Key Revocation | `CNSC-274`
**Originating Document**: `Cloud Native Security Whitepaper v2.0`


**Objective**

Should software artifacts become untrusted due to compromise or other incident, teams should revoke signing keys to ensure repudiation

*No NIST SP 800-53 mapping row in this document for this guideline.*

---



<a id="cnsc-297"></a>

### Configuration Signing | `CNSC-297`
**Originating Document**: `Cloud Native Security Whitepaper v2.0`


**Objective**

Workload-related configuration is signed

**Guideline Mappings**

**`NIST-800-53`**

| NIST control | Remarks |
|--------------|---------|
| `SI-7` | Software, Firmware, and Information Integrity |


---



<a id="cnsc-298"></a>

### Package Signing | `CNSC-298`
**Originating Document**: `Cloud Native Security Whitepaper v2.0`


**Objective**

Workload-related package is signed

**Guideline Mappings**

**`NIST-800-53`**

| NIST control | Remarks |
|--------------|---------|
| `SI-7` | Software, Firmware, and Information Integrity |


---



<a id="cnsc-303"></a>

### Credential Protection | `CNSC-303`
**Originating Document**: `Cloud Native Security Whitepaper v2.0`


**Objective**

HSMs or credential managers should be used for protecting credentials. If this is not possible, software-based credential managers should be used.

**Guideline Mappings**

**`NIST-800-53`**

| NIST control | Remarks |
|--------------|---------|
| `SC-12(3)` | Systems and Communication Protection |


---



<a id="cnsc-79"></a>

### Trust Verification | `CNSC-79`
**Originating Document**: `Cloud Native Security Whitepaper v1.0`


**Objective**

Trust is verified

*No NIST SP 800-53 mapping row in this document for this guideline.*

---



<a id="cnsc-80"></a>

### Staging Registry Management | `CNSC-80`
**Originating Document**:
- `Cloud Native Security Whitepaper v1.0`
- `Cloud Native Security Whitepaper v2.0`



**Objective**

Artifacts ready for deployment are managed in a staging or pre-prod registry

*No NIST SP 800-53 mapping row in this document for this guideline.*

---



<a id="cnsc-81"></a>

### Container Image Hardening | `CNSC-81`
**Originating Document**:
- `Cloud Native Security Whitepaper v1.0`
- `Cloud Native Security Whitepaper v2.0`



**Objective**

Container images are hardened following best practices

*No NIST SP 800-53 mapping row in this document for this guideline.*

---



<a id="cnsc-82"></a>

### Static Application Security Testing | `CNSC-82`
**Originating Document**:
- `Cloud Native Security Whitepaper v1.0`
- `Cloud Native Security Whitepaper v2.0`



**Objective**

Static application security testing (SAST) is performed

*No NIST SP 800-53 mapping row in this document for this guideline.*

---



<a id="cnsc-83"></a>

### Test Pyramid Adherence | `CNSC-83`
**Originating Document**:
- `Cloud Native Security Whitepaper v1.0`
- `Cloud Native Security Whitepaper v2.0`



**Objective**

Test suites follow the test pyramid

*No NIST SP 800-53 mapping row in this document for this guideline.*

---



<a id="cnsc-84"></a>

### Private Development Registry | `CNSC-84`
**Originating Document**:
- `Cloud Native Security Whitepaper v1.0`
- `Cloud Native Security Whitepaper v2.0`



**Objective**

Artifacts undergoing active development are held in a private registry

*No NIST SP 800-53 mapping row in this document for this guideline.*

---



<a id="cnsc-85"></a>

### Manifest Scanning | `CNSC-85`
**Originating Document**:
- `Cloud Native Security Whitepaper v1.0`
- `Cloud Native Security Whitepaper v2.0`



**Objective**

Application manifests are scanned in CI pipeline

**Guideline Mappings**

**`NIST-800-53`**

| NIST control | Remarks |
|--------------|---------|
| `RA-5` | Vulnerability Monitoring and Scanning |


---



<a id="cnsc-86"></a>

### CI Server Isolation | `CNSC-86`
**Originating Document**:
- `Cloud Native Security Whitepaper v1.0`
- `Cloud Native Security Whitepaper v2.0`



**Objective**

CI servers for sensitive workloads are isolated from other workloads

**Guideline Mappings**

**`NIST-800-53`**

| NIST control | Remarks |
|--------------|---------|
| `SC-39` | Process Isolation |


---



<a id="cnsc-87"></a>

### Privileged Build Isolation | `CNSC-87`
**Originating Document**:
- `Cloud Native Security Whitepaper v1.0`
- `Cloud Native Security Whitepaper v2.0`



**Objective**

Builds requiring elevated privileges run on dedicated servers

**Guideline Mappings**

**`NIST-800-53`**

| NIST control | Remarks |
|--------------|---------|
| `SC-39` | Process Isolation |


---



<a id="cnsc-88"></a>

### Build Policy Enforcement | `CNSC-88`
**Originating Document**:
- `Cloud Native Security Whitepaper v1.0`
- `Cloud Native Security Whitepaper v2.0`



**Objective**

Build policies are enforced on the CI pipeline

**Guideline Mappings**

**`NIST-800-53`**

| NIST control | Remarks |
|--------------|---------|
| `SA-1` | Policy and Procedures |


---



<a id="cnsc-89"></a>

### Pipeline Metadata Signing | `CNSC-89`
**Originating Document**:
- `Cloud Native Security Whitepaper v1.0`
- `Cloud Native Security Whitepaper v2.0`



**Objective**

Pipeline metadata is signed

**Guideline Mappings**

**`NIST-800-53`**

| NIST control | Remarks |
|--------------|---------|
| `SI-7` | Software, Firmware, and Information Integrity |


---



<a id="cnsc-90"></a>

### Build Stage Verification | `CNSC-90`
**Originating Document**:
- `Cloud Native Security Whitepaper v1.0`
- `Cloud Native Security Whitepaper v2.0`



**Objective**

Build stages are verified prior to the next stage executing

**Guideline Mappings**

**`NIST-800-53`**

| NIST control | Remarks |
|--------------|---------|
| `SI-7` | Software, Firmware, and Information Integrity |


---



<a id="cnsc-91"></a>

### CI Pipeline Scanning | `CNSC-91`
**Originating Document**:
- `Cloud Native Security Whitepaper v1.0`
- `Cloud Native Security Whitepaper v2.0`



**Objective**

Images are scanned within the CI pipeline

**Guideline Mappings**

**`NIST-800-53`**

| NIST control | Remarks |
|--------------|---------|
| `RA-5` | Vulnerability Monitoring and Scanning |


---



<a id="cnsc-92"></a>

### Pipeline Compliance Integration | `CNSC-92`
**Originating Document**:
- `Cloud Native Security Whitepaper v1.0`
- `Cloud Native Security Whitepaper v2.0`



**Objective**

Vulnerability scans are coupled with pipeline compliance rules

**Guideline Mappings**

**`NIST-800-53`**

| NIST control | Remarks |
|--------------|---------|
| `SA-1` | Policy and Procedures |


---



<a id="cnsc-93"></a>

### Dynamic Application Security Testing | `CNSC-93`
**Originating Document**:
- `Cloud Native Security Whitepaper v1.0`
- `Cloud Native Security Whitepaper v2.0`



**Objective**

Dynamic application security testing (DAST) is performed

**Guideline Mappings**

**`NIST-800-53`**

| NIST control | Remarks |
|--------------|---------|
| `SA-11(8)` | Interactive Application Security Testing |


---



<a id="cnsc-94"></a>

### Application Instrumentation | `CNSC-94`
**Originating Document**:
- `Cloud Native Security Whitepaper v1.0`
- `Cloud Native Security Whitepaper v2.0`



**Objective**

Application instrumentation is employed

**Guideline Mappings**

**`NIST-800-53`**

| NIST control | Remarks |
|--------------|---------|
| `SI-4` | System Monitoring |


---



<a id="cnsc-95"></a>

### Test Traceability | `CNSC-95`
**Originating Document**:
- `Cloud Native Security Whitepaper v1.0`
- `Cloud Native Security Whitepaper v2.0`



**Objective**

Automated test results map back to requirements

*No NIST SP 800-53 mapping row in this document for this guideline.*

---



<a id="cnsc-96"></a>

### Infrastructure Security Testing | `CNSC-96`
**Originating Document**:
- `Cloud Native Security Whitepaper v1.0`
- `Cloud Native Security Whitepaper v2.0`



**Objective**

Infrastructure security tests are employed

*No NIST SP 800-53 mapping row in this document for this guideline.*

---



<a id="cnsc-97"></a>

### Security Health Verification | `CNSC-97`
**Originating Document**:
- `Cloud Native Security Whitepaper v1.0`
- `Cloud Native Security Whitepaper v2.0`



**Objective**

Tests to verify security health are executed at build and deploy time

**Guideline Mappings**

**`NIST-800-53`**

| NIST control | Remarks |
|--------------|---------|
| `SI-4` | System Monitoring |


---



<a id="cnsc-98"></a>

### IaC Policy Controls | `CNSC-98`
**Originating Document**:
- `Cloud Native Security Whitepaper v1.0`
- `Cloud Native Security Whitepaper v2.0`



**Objective**

Infrastructure as Code is subject to the same pipeline policy controls as application code

*No NIST SP 800-53 mapping row in this document for this guideline.*

---



<a id="cnsc-99"></a>

### Automated Security Testing | `CNSC-99`
**Originating Document**:
- `Cloud Native Security Whitepaper v1.0`
- `Cloud Native Security Whitepaper v2.0`



**Objective**

Security testing is automated

**Guideline Mappings**

**`NIST-800-53`**

| NIST control | Remarks |
|--------------|---------|
| `CA-8`<br />`SA-11` | <ul><li>Penetration Testing</li><li>Developer Testing and Evaluation</li></ul> |


---




<a id="securing-artefacts"></a>

## Securing Artefacts

Guidelines for securing artefacts, including build process attestation, signing frameworks, key rotation, OCI registry support, and encryption.



<a id="cnsc-141"></a>

### Build Process Attestation | `CNSC-141`
**Originating Document**: `Software Supply Chain Best Practices v1.0`


**Objective**

Every step in the build process is signed and attested for process integrity

**Guideline Mappings**

**`NIST-800-53`**

| NIST control | Remarks |
|--------------|---------|
| `SI-1`<br />`SI-7` | <ul><li>Policy and Procedures</li><li>Software, Firmware, and Information Integrity</li></ul> |


---



<a id="cnsc-142"></a>

### Build Signature Verification | `CNSC-142`
**Originating Document**: `Software Supply Chain Best Practices v1.0`


**Objective**

Every step in the build process verifies previously generated signatures

**Guideline Mappings**

**`NIST-800-53`**

| NIST control | Remarks |
|--------------|---------|
| `SI-1`<br />`SI-7` | <ul><li>Policy and Procedures</li><li>Software, Firmware, and Information Integrity</li></ul> |


---



<a id="cnsc-143"></a>

### Artifact Signing Framework | `CNSC-143`
**Originating Document**: `Software Supply Chain Best Practices v1.0`


**Objective**

A framework is used to manage signing of artefacts

**Guideline Mappings**

**`NIST-800-53`**

| NIST control | Remarks |
|--------------|---------|
| `IA-5` | Authenticator Management |


---



<a id="cnsc-144"></a>

### Attestation Store | `CNSC-144`
**Originating Document**: `Software Supply Chain Best Practices v1.0`


**Objective**

A store is used to manage attestations

**Guideline Mappings**

**`NIST-800-53`**

| NIST control | Remarks |
|--------------|---------|
| `AC-4(6)` | Information Flow Enforcement, Metadata |


---



<a id="cnsc-145"></a>

### Certification Authorization | `CNSC-145`
**Originating Document**: `Software Supply Chain Best Practices v1.0`


**Objective**

Artefacts any given party is authorized to certify are limited

**Guideline Mappings**

**`NIST-800-53`**

| NIST control | Remarks |
|--------------|---------|
| `AC-6` | Least Privilege |


---



<a id="cnsc-146"></a>

### Key Rotation and Revocation | `CNSC-146`
**Originating Document**: `Software Supply Chain Best Practices v1.0`


**Objective**

Rotation and revocation of private keys is supported

**Guideline Mappings**

**`NIST-800-53`**

| NIST control | Remarks |
|--------------|---------|
| `SC-12` | Cryptographic Key Establishment and Management |


---



<a id="cnsc-147"></a>

### OCI Registry Support | `CNSC-147`
**Originating Document**: `Software Supply Chain Best Practices v1.0`


**Objective**

A container registry that supports OCI image-spec images is used

*No NIST SP 800-53 mapping row in this document for this guideline.*

---



<a id="cnsc-148"></a>

### Artifact Encryption | `CNSC-148`
**Originating Document**: `Software Supply Chain Best Practices v1.0`


**Objective**

Artefacts are encrypted before distribution and only authorized platforms have decryption capabilities

**Guideline Mappings**

**`NIST-800-53`**

| NIST control | Remarks |
|--------------|---------|
| `IA-5`<br />`SC-12`<br />`SC-13`<br />`SC-28(1)`<br />`SC-8` | <ul><li>Authenticator Management</li><li>Cryptographic Key Establishment and Management</li><li>Cryptographic Protection</li><li>Protection of Information at Rest, Cryptographic Protection</li><li>Transmission Confidentiality and Integrity</li></ul> |


---




<a id="securing-build-pipelines"></a>

## Securing Build Pipelines

Guidelines for securing build pipelines, ensuring cryptographic guarantees, validation, and secure build environments.



<a id="cnsc-149"></a>

### Cryptographic Policy Guarantee | `CNSC-149`
**Originating Document**: `Software Supply Chain Best Practices v1.0`


**Objective**

Policy adherence is cryptographically guaranteed

**Guideline Mappings**

**`NIST-800-53`**

| NIST control | Remarks |
|--------------|---------|
| `CM-3(6)` | Configuration Change Control |


---



<a id="cnsc-150"></a>

### Environment and Dependency Validation | `CNSC-150`
**Originating Document**: `Software Supply Chain Best Practices v1.0`


**Objective**

Environments and dependencies are validated before usage

**Guideline Mappings**

**`NIST-800-53`**

| NIST control | Remarks |
|--------------|---------|
| `CM-3(2)` | Configuration Change Control |


---



<a id="cnsc-151"></a>

### Build Worker Runtime Security | `CNSC-151`
**Originating Document**: `Software Supply Chain Best Practices v1.0`


**Objective**

Runtime security of build workers is validated

**Guideline Mappings**

**`NIST-800-53`**

| NIST control | Remarks |
|--------------|---------|
| `CM-3(4)` | NIST SP 800-53 control mapped from CNSC-151 (Build Worker Runtime Security); family: Securing Build Pipelines. |


---



<a id="cnsc-152"></a>

### Reproducible Builds | `CNSC-152`
**Originating Document**: `Software Supply Chain Best Practices v1.0`


**Objective**

Build artefacts are validated through verifiably reproducible builds

**Guideline Mappings**

**`NIST-800-53`**

| NIST control | Remarks |
|--------------|---------|
| `CM-3(4)` | NIST SP 800-53 control mapped from CNSC-152 (Reproducible Builds); family: Securing Build Pipelines. |


---



<a id="cnsc-153"></a>

### External Requirement Verification | `CNSC-153`
**Originating Document**: `Software Supply Chain Best Practices v1.0`


**Objective**

External requirements from the build process are locked and verified

**Guideline Mappings**

**`NIST-800-53`**

| NIST control | Remarks |
|--------------|---------|
| `CM-3(2)` | NIST SP 800-53 control mapped from CNSC-153 (External Requirement Verification); family: Securing Build Pipelines. |


---



<a id="cnsc-154"></a>

### Build Determinism | `CNSC-154`
**Originating Document**: `Software Supply Chain Best Practices v1.0`


**Objective**

Sources of non-determinism are found and eliminated

*No NIST SP 800-53 mapping row in this document for this guideline.*

---



<a id="cnsc-155"></a>

### Build Environment Recording | `CNSC-155`
**Originating Document**: `Software Supply Chain Best Practices v1.0`


**Objective**

The build environment is recorded

**Guideline Mappings**

**`NIST-800-53`**

| NIST control | Remarks |
|--------------|---------|
| `CM-3(1)` | NIST SP 800-53 control mapped from CNSC-155 (Build Environment Recording); family: Securing Build Pipelines. |


---



<a id="cnsc-156"></a>

### Build Environment Automation | `CNSC-156`
**Originating Document**: `Software Supply Chain Best Practices v1.0`


**Objective**

Creation of the build environment is automated

**Guideline Mappings**

**`NIST-800-53`**

| NIST control | Remarks |
|--------------|---------|
| `CM-3(3)` | NIST SP 800-53 control mapped from CNSC-156 (Build Environment Automation); family: Securing Build Pipelines. |


---



<a id="cnsc-157"></a>

### Build Distribution | `CNSC-157`
**Originating Document**: `Software Supply Chain Best Practices v1.0`


**Objective**

Builds are distributed across different infrastructure

**Guideline Mappings**

**`NIST-800-53`**

| NIST control | Remarks |
|--------------|---------|
| `CM-3(3)` | NIST SP 800-53 control mapped from CNSC-157 (Build Distribution); family: Securing Build Pipelines. |


---



<a id="cnsc-158"></a>

### Pipeline as Code | `CNSC-158`
**Originating Document**: `Software Supply Chain Best Practices v1.0`


**Objective**

Build and related CI/CD steps are automated through a pipeline delivered as code

**Guideline Mappings**

**`NIST-800-53`**

| NIST control | Remarks |
|--------------|---------|
| `SA-3` | System Development Life Cycle |


---



<a id="cnsc-159"></a>

### Pipeline Standardization | `CNSC-159`
**Originating Document**: `Software Supply Chain Best Practices v1.0`


**Objective**

Pipelines are standardized across projects

*No NIST SP 800-53 mapping row in this document for this guideline.*

---



<a id="cnsc-160"></a>

### Software Factory Platform | `CNSC-160`
**Originating Document**: `Software Supply Chain Best Practices v1.0`


**Objective**

A secured orchestration platform is provisioned to host the software factory

*No NIST SP 800-53 mapping row in this document for this guideline.*

---



<a id="cnsc-161"></a>

### Single-Use Build Workers | `CNSC-161`
**Originating Document**: `Software Supply Chain Best Practices v1.0`


**Objective**

Build workers are single use

**Guideline Mappings**

**`NIST-800-53`**

| NIST control | Remarks |
|--------------|---------|
| `AC-2` | Account Management |


---



<a id="cnsc-162"></a>

### Software Factory Network Isolation | `CNSC-162`
**Originating Document**: `Software Supply Chain Best Practices v1.0`


**Objective**

Software factory has minimal network connectivity

**Guideline Mappings**

**`NIST-800-53`**

| NIST control | Remarks |
|--------------|---------|
| `SC-7(3)` | Boundary Protection |


---



<a id="cnsc-163"></a>

### Build Worker Duty Segregation | `CNSC-163`
**Originating Document**: `Software Supply Chain Best Practices v1.0`


**Objective**

Duties of each build worker are segregated

**Guideline Mappings**

**`NIST-800-53`**

| NIST control | Remarks |
|--------------|---------|
| `AC-5` | Separation of Duties |


---



<a id="cnsc-164"></a>

### Build Worker Environment Control | `CNSC-164`
**Originating Document**: `Software Supply Chain Best Practices v1.0`


**Objective**

Build worker environment and commands are passed in

**Guideline Mappings**

**`NIST-800-53`**

| NIST control | Remarks |
|--------------|---------|
| `CM-2(2)` | Baseline Configuration |


---



<a id="cnsc-165"></a>

### Secured Output Storage | `CNSC-165`
**Originating Document**: `Software Supply Chain Best Practices v1.0`


**Objective**

Output is written to a separate secured storage repo

**Guideline Mappings**

**`NIST-800-53`**

| NIST control | Remarks |
|--------------|---------|
| `AU-9(2)` | Protection of Audit Information |


---



<a id="cnsc-166"></a>

### Pipeline Modification Control | `CNSC-166`
**Originating Document**: `Software Supply Chain Best Practices v1.0`


**Objective**

Pipeline modification is only allowed through pipeline as code

*No NIST SP 800-53 mapping row in this document for this guideline.*

---



<a id="cnsc-167"></a>

### User Role Definition | `CNSC-167`
**Originating Document**: `Software Supply Chain Best Practices v1.0`


**Objective**

User roles are defined

**Guideline Mappings**

**`NIST-800-53`**

| NIST control | Remarks |
|--------------|---------|
| `AC-2` | Account Management |


---



<a id="cnsc-168"></a>

### Root of Trust Establishment | `CNSC-168`
**Originating Document**: `Software Supply Chain Best Practices v1.0`


**Objective**

Established practices are followed for establishing a root of trust from an offline source

**Guideline Mappings**

**`NIST-800-53`**

| NIST control | Remarks |
|--------------|---------|
| `SC-17` | Public Key Infrastructure Certificates |


---



<a id="cnsc-169"></a>

### Short-Lived Certificates | `CNSC-169`
**Originating Document**: `Software Supply Chain Best Practices v1.0`


**Objective**

Short-lived workload certificates are used

**Guideline Mappings**

**`NIST-800-53`**

| NIST control | Remarks |
|--------------|---------|
| `SC-23(5)` | Session Authenticity |


---




<a id="securing-materials"></a>

## Securing Materials

Guidelines for securing materials, including third-party artifact verification, SBOM generation, dependency tracking, vulnerability scanning, and software composition analysis.



<a id="cnsc-173"></a>

### Third-Party Artifact Verification | `CNSC-173`
**Originating Document**: `Software Supply Chain Best Practices v1.0`


**Objective**

Third party artefacts and open source libraries are verified

**Guideline Mappings**

**`NIST-800-53`**

| NIST control | Remarks |
|--------------|---------|
| `SA-11` | Developer Testing and Evaluation |


---



<a id="cnsc-174"></a>

### Third-Party SBOM Requirements | `CNSC-174`
**Originating Document**: `Software Supply Chain Best Practices v1.0`


**Objective**

SBOM is required from third party suppliers

**Guideline Mappings**

**`NIST-800-53`**

| NIST control | Remarks |
|--------------|---------|
| `CM-8` | Information System Component Inventory |


---



<a id="cnsc-175"></a>

### Dependency Tracking | `CNSC-175`
**Originating Document**: `Software Supply Chain Best Practices v1.0`


**Objective**

Dependencies between open source components are tracked

**Guideline Mappings**

**`NIST-800-53`**

| NIST control | Remarks |
|--------------|---------|
| `CM-10` | Software Usage Restrictions |


---



<a id="cnsc-176"></a>

### Source-Based Library Builds | `CNSC-176`
**Originating Document**: `Software Supply Chain Best Practices v1.0`


**Objective**

Libraries are built based upon source code

*No NIST SP 800-53 mapping row in this document for this guideline.*

---



<a id="cnsc-177"></a>

### Trusted Package Sources | `CNSC-177`
**Originating Document**: `Software Supply Chain Best Practices v1.0`


**Objective**

Trusted package managers and repositories are defined and prioritized

*No NIST SP 800-53 mapping row in this document for this guideline.*

---



<a id="cnsc-178"></a>

### Immutable SBOM Generation | `CNSC-178`
**Originating Document**: `Software Supply Chain Best Practices v1.0`


**Objective**

An immutable SBOM of the code is generated

*No NIST SP 800-53 mapping row in this document for this guideline.*

---



<a id="cnsc-179"></a>

### Software Vulnerability Scanning | `CNSC-179`
**Originating Document**: `Software Supply Chain Best Practices v1.0`


**Objective**

Software is scanned for vulnerabilities

**Guideline Mappings**

**`NIST-800-53`**

| NIST control | Remarks |
|--------------|---------|
| `RA-5` | Vulnerability Monitoring and Scanning |


---



<a id="cnsc-180"></a>

### License Compliance Scanning | `CNSC-180`
**Originating Document**: `Software Supply Chain Best Practices v1.0`


**Objective**

Software is scanned for license implications

**Guideline Mappings**

**`NIST-800-53`**

| NIST control | Remarks |
|--------------|---------|
| `CM-10` | Software Usage Restrictions |


---



<a id="cnsc-181"></a>

### Software Composition Analysis | `CNSC-181`
**Originating Document**: `Software Supply Chain Best Practices v1.0`


**Objective**

Software composition analysis is run on ingested software

**Guideline Mappings**

**`NIST-800-53`**

| NIST control | Remarks |
|--------------|---------|
| `SA-11(1)(8)` | Developer Testing and Evaluation |


---




<a id="securing-the-source-code"></a>

## Securing the Source Code

Guidelines for securing the source code, including commit signing, branch protection, secret prevention, access control, MFA enforcement, and automated security scanning.



<a id="cnsc-182"></a>

### Commit and Tag Signing | `CNSC-182`
**Originating Document**: `Software Supply Chain Best Practices v1.0`


**Objective**

Commits and tags are signed

**Guideline Mappings**

**`NIST-800-53`**

| NIST control | Remarks |
|--------------|---------|
| `SI-7` | Software, Firmware, and information integrity |


---



<a id="cnsc-183"></a>

### Branch Protection Attestation | `CNSC-183`
**Originating Document**: `Software Supply Chain Best Practices v1.0`


**Objective**

Full attestation and verification is enforced for protected branches

**Guideline Mappings**

**`NIST-800-53`**

| NIST control | Remarks |
|--------------|---------|
| `AC-6(3)` | Least Privilege |


---



<a id="cnsc-184"></a>

### Secret Commit Prevention | `CNSC-184`
**Originating Document**: `Software Supply Chain Best Practices v1.0`


**Objective**

Secrets are not committed to the source code repository unless encrypted

**Guideline Mappings**

**`NIST-800-53`**

| NIST control | Remarks |
|--------------|---------|
| `SC-12(3)` | Systems and Communication Protection |


---



<a id="cnsc-185"></a>

### Repository Access Definition | `CNSC-185`
**Originating Document**: `Software Supply Chain Best Practices v1.0`


**Objective**

Individuals or teams with write access to a repository are defined

**Guideline Mappings**

**`NIST-800-53`**

| NIST control | Remarks |
|--------------|---------|
| `PL-1` | Policy and Procedures |


---



<a id="cnsc-186"></a>

### Automated Security Scanning | `CNSC-186`
**Originating Document**: `Software Supply Chain Best Practices v1.0`


**Objective**

Software security scanning and testing is automated

**Guideline Mappings**

**`NIST-800-53`**

| NIST control | Remarks |
|--------------|---------|
| `RA-5` | Vulnerability Monitoring and Scanning |


---



<a id="cnsc-187"></a>

### Contribution Policy Enforcement | `CNSC-187`
**Originating Document**: `Software Supply Chain Best Practices v1.0`


**Objective**

Contribution policies are established and adhered to

**Guideline Mappings**

**`NIST-800-53`**

| NIST control | Remarks |
|--------------|---------|
| `PL-1` | Policy and Procedures |


---



<a id="cnsc-188"></a>

### Functional Role Definition | `CNSC-188`
**Originating Document**: `Software Supply Chain Best Practices v1.0`


**Objective**

Roles are defined and aligned to functional responsibilities

**Guideline Mappings**

**`NIST-800-53`**

| NIST control | Remarks |
|--------------|---------|
| `PL-1` | Policy and Procedures |


---



<a id="cnsc-189"></a>

### Four-Eyes Principle | `CNSC-189`
**Originating Document**: `Software Supply Chain Best Practices v1.0`


**Objective**

An independent four-eyes principle is enforced

**Guideline Mappings**

**`NIST-800-53`**

| NIST control | Remarks |
|--------------|---------|
| `SA-11` | Developer Testing and Evaluation |


---



<a id="cnsc-190"></a>

### Branch Protection Rules | `CNSC-190`
**Originating Document**: `Software Supply Chain Best Practices v1.0`


**Objective**

Branch protection rules are used

**Guideline Mappings**

**`NIST-800-53`**

| NIST control | Remarks |
|--------------|---------|
| `SA-8` | Security Engineering Principles |


---



<a id="cnsc-191"></a>

### Repository MFA Enforcement | `CNSC-191`
**Originating Document**: `Software Supply Chain Best Practices v1.0`


**Objective**

MFA is enforced for accessing source code repositories

**Guideline Mappings**

**`NIST-800-53`**

| NIST control | Remarks |
|--------------|---------|
| `IA-2(1)` | Identification and Authentication (organizational Users) |


---



<a id="cnsc-192"></a>

### SSH Key Access | `CNSC-192`
**Originating Document**: `Software Supply Chain Best Practices v1.0`


**Objective**

SSH keys are used to provide developers access to source code repositories

**Guideline Mappings**

**`NIST-800-53`**

| NIST control | Remarks |
|--------------|---------|
| `AC-1` | Remote Access |


---



<a id="cnsc-193"></a>

### Key Rotation Policy | `CNSC-193`
**Originating Document**: `Software Supply Chain Best Practices v1.0`


**Objective**

A key rotation policy is maintained

**Guideline Mappings**

**`NIST-800-53`**

| NIST control | Remarks |
|--------------|---------|
| `AC-2(1)` | Prerequisites and criteria for group and role membership are defined. |


---



<a id="cnsc-194"></a>

### Ephemeral Credentials | `CNSC-194`
**Originating Document**: `Software Supply Chain Best Practices v1.0`


**Objective**

Short-lived or ephemeral credentials are used for machine and service access

**Guideline Mappings**

**`NIST-800-53`**

| NIST control | Remarks |
|--------------|---------|
| `AC-2(1)` | Usage of automated mechanisms to create, enable, modify, disable, and remove accounts. |


---




<a id="security-assurance"></a>

## Security Assurance

Guidelines for security assurance, including network policies, incident response, platform hardening, threat modeling, identity management, and GitOps security practices.



<a id="cnsc-115"></a>

### East-West Network Policy | `CNSC-115`
**Originating Document**:
- `Cloud Native Security Whitepaper v1.0`
- `Cloud Native Security Whitepaper v2.0`



**Objective**

Network policies enforce east-west network communication within the container deployment is limited to only that which is authorized for access

**Guideline Mappings**

**`NIST-800-53`**

| NIST control | Remarks |
|--------------|---------|
| `AC-6(3)` | Least Privilege, Network Access to Privileged Commands |


---



<a id="cnsc-116"></a>

### Incident Response | `CNSC-116`
**Originating Document**:
- `Cloud Native Security Whitepaper v1.0`
- `Cloud Native Security Whitepaper v2.0`



**Objective**

Incident response considers cloud native workloads

**Guideline Mappings**

**`NIST-800-53`**

| NIST control | Remarks |
|--------------|---------|
| `CA-7`<br />`IR-4`<br />`IR-4(5)` | <ul><li>Continuous Monitoring</li><li>Incident Handling, Automated Incident Handling Processes</li><li>Incident Handling, Automated Disabling of System</li></ul> |


---



<a id="cnsc-117"></a>

### Incident Monitoring | `CNSC-117`
**Originating Document**:
- `Cloud Native Security Whitepaper v1.0`
- `Cloud Native Security Whitepaper v2.0`



**Objective**

Incident response accounts for appropriate evidence handling and collection of cloud native workloads

**Guideline Mappings**

**`NIST-800-53`**

| NIST control | Remarks |
|--------------|---------|
| `IR-5(1)` | Incident Monitoring, Automated Tracking, Data Collection, and Analysis |


---



<a id="cnsc-118"></a>

### Security Assurance Management | `CNSC-118`
**Originating Document**:
- `Cloud Native Security Whitepaper v1.0`
- `Cloud Native Security Whitepaper v2.0`



**Objective**

Rootless builds are employed

*No NIST SP 800-53 mapping row in this document for this guideline.*

---



<a id="cnsc-119"></a>

### Workload and Deployment Isolation | `CNSC-119`
**Originating Document**:
- `Cloud Native Security Whitepaper v1.0`
- `Cloud Native Security Whitepaper v2.0`



**Objective**

Cgroups and system groups are used to isolate workloads and deployments

*No NIST SP 800-53 mapping row in this document for this guideline.*

---



<a id="cnsc-120"></a>

### Mandatory Access Controls | `CNSC-120`
**Originating Document**:
- `Cloud Native Security Whitepaper v1.0`
- `Cloud Native Security Whitepaper v2.0`



**Objective**

MAC implementations are employed

**Guideline Mappings**

**`NIST-800-53`**

| NIST control | Remarks |
|--------------|---------|
| `AC-3(3)` | Access Enforcement, Mandatory Access Control |


---



<a id="cnsc-121"></a>

### Threat Modeling and Vulnerability Analysis | `CNSC-121`
**Originating Document**:
- `Cloud Native Security Whitepaper v1.0`
- `Cloud Native Security Whitepaper v2.0`



**Objective**

Threat model code and infrastructure

**Guideline Mappings**

**`NIST-800-53`**

| NIST control | Remarks |
|--------------|---------|
| `SA-11(2)` | Developer Testing and Evaluation, Threat Modeling and Vulnerability Analyses |


---



<a id="cnsc-122"></a>

### Authentication Management | `CNSC-122`
**Originating Document**:
- `Cloud Native Security Whitepaper v1.0`
- `Cloud Native Security Whitepaper v2.0`



**Objective**

Entities are able to independently authenticate other identities

**Guideline Mappings**

**`NIST-800-53`**

| NIST control | Remarks |
|--------------|---------|
| `IA-9` | Service Identification and Authentication |


---



<a id="cnsc-123"></a>

### Identity Management | `CNSC-123`
**Originating Document**:
- `Cloud Native Security Whitepaper v1.0`
- `Cloud Native Security Whitepaper v2.0`



**Objective**

Each entity can create proof of who the identity is

**Guideline Mappings**

**`NIST-800-53`**

| NIST control | Remarks |
|--------------|---------|
| `IA-9` | Service Identification and Authentication |


---



<a id="cnsc-124"></a>

### Trusted Components | `CNSC-124`
**Originating Document**:
- `Cloud Native Security Whitepaper v1.0`
- `Cloud Native Security Whitepaper v2.0`



**Objective**

Orchestrator is running on a trusted OS, BIOS, etc

**Guideline Mappings**

**`NIST-800-53`**

| NIST control | Remarks |
|--------------|---------|
| `CM-14` | Signed Components |


---



<a id="cnsc-125"></a>

### Security Verification | `CNSC-125`
**Originating Document**:
- `Cloud Native Security Whitepaper v1.0`
- `Cloud Native Security Whitepaper v2.0`



**Objective**

Orchestrator verifies the claims of a container

**Guideline Mappings**

**`NIST-800-53`**

| NIST control | Remarks |
|--------------|---------|
| `SI-6` | Security and Privacy Function Verification |


---



<a id="cnsc-126"></a>

### Network Policy Management | `CNSC-126`
**Originating Document**:
- `Cloud Native Security Whitepaper v1.0`
- `Cloud Native Security Whitepaper v2.0`



**Objective**

Orchestrator network policies are used in conjunction with a service mesh

*No NIST SP 800-53 mapping row in this document for this guideline.*

---



<a id="cnsc-196"></a>

### Supply Chain Security Best Practices | `CNSC-196`
**Originating Document**: `Cloud Native Security Whitepaper v2.0`


**Objective**

Supply chain security best practices are adhered to

*No NIST SP 800-53 mapping row in this document for this guideline.*

---



<a id="cnsc-197"></a>

### Repository and Branch Access Control | `CNSC-197`
**Originating Document**: `Cloud Native Security Whitepaper v2.0`


**Objective**

Access to repository and branches is restricted

*No NIST SP 800-53 mapping row in this document for this guideline.*

---



<a id="cnsc-198"></a>

### Git Secret Prevention | `CNSC-198`
**Originating Document**: `Cloud Native Security Whitepaper v2.0`


**Objective**

Unencrypted credentials or secrets are never stored in the Git repository and sensitive data is blocked from being pushed

*No NIST SP 800-53 mapping row in this document for this guideline.*

---



<a id="cnsc-199"></a>

### Commit Identity Enforcement | `CNSC-199`
**Originating Document**: `Cloud Native Security Whitepaper v2.0`


**Objective**

Strong identity is enforced with GPG Signed Commits for accountability and traceability

*No NIST SP 800-53 mapping row in this document for this guideline.*

---



<a id="cnsc-200"></a>

### Linear History Enforcement | `CNSC-200`
**Originating Document**: `Cloud Native Security Whitepaper v2.0`


**Objective**

Linear history is required and commit history is maintained by disallowing force pushes

*No NIST SP 800-53 mapping row in this document for this guideline.*

---



<a id="cnsc-201"></a>

### Branching Policy Enforcement | `CNSC-201`
**Originating Document**: `Cloud Native Security Whitepaper v2.0`


**Objective**

Branching policy is enforced with main branch protection and code review required before merging

*No NIST SP 800-53 mapping row in this document for this guideline.*

---



<a id="cnsc-202"></a>

### Git Tool Security Monitoring | `CNSC-202`
**Originating Document**: `Cloud Native Security Whitepaper v2.0`


**Objective**

Vulnerabilities are monitored and Git and GitOps tools are kept up to date

*No NIST SP 800-53 mapping row in this document for this guideline.*

---



<a id="cnsc-203"></a>

### Repository Credential Management | `CNSC-203`
**Originating Document**: `Cloud Native Security Whitepaper v2.0`


**Objective**

SSH keys and Personal Access Tokens are rotated and unauthorized access to Git repositories is blocked

*No NIST SP 800-53 mapping row in this document for this guideline.*

---



<a id="cnsc-204"></a>

### Technical Account Management | `CNSC-204`
**Originating Document**: `Cloud Native Security Whitepaper v2.0`


**Objective**

Dedicated non-user technical accounts are used for access with frequently rotated short-lived credentials

*No NIST SP 800-53 mapping row in this document for this guideline.*

---



<a id="cnsc-205"></a>

### Privilege Escalation Prevention | `CNSC-205`
**Originating Document**: `Cloud Native Security Whitepaper v2.0`


**Objective**

Users who can elevate permissions to remove security features are limited to prevent deletion of audit trails and silencing of alerts

*No NIST SP 800-53 mapping row in this document for this guideline.*

---




<a id="storage"></a>

## Storage

Guidelines for securing storage, including encryption at rest, data availability, integrity validation, namespace boundaries, access policies, and artifact registry management.



<a id="cnsc-127"></a>

### Control Plane Authentication | `CNSC-127`
**Originating Document**:
- `Cloud Native Security Whitepaper v1.0`
- `Cloud Native Security Whitepaper v2.0`



**Objective**

Storage control plane management interface requires mutual authentication and TLS for connections

**Guideline Mappings**

**`NIST-800-53`**

| NIST control | Remarks |
|--------------|---------|
| `SC-8` | Transmission Confidentiality and Integrity |


---



<a id="cnsc-128"></a>

### Data Availability Mechanism | `CNSC-128`
**Originating Document**:
- `Cloud Native Security Whitepaper v1.0`
- `Cloud Native Security Whitepaper v2.0`



**Objective**

Data availability is achieved through parity or mirroring, erasure coding or replicas

**Guideline Mappings**

**`NIST-800-53`**

| NIST control | Remarks |
|--------------|---------|
| `SI-13` | Predictable Failure Prevention |


---



<a id="cnsc-129"></a>

### Integrity Validation | `CNSC-129`
**Originating Document**:
- `Cloud Native Security Whitepaper v1.0`
- `Cloud Native Security Whitepaper v2.0`



**Objective**

Hashing and checksums are added to blocks, objects or files

**Guideline Mappings**

**`NIST-800-53`**

| NIST control | Remarks |
|--------------|---------|
| `CM-7`<br />`SI-7` | <ul><li>Least Functionality</li><li>Software, Firmware, and Information Integrity</li></ul> |


---



<a id="cnsc-130"></a>

### Data Source Storage Management | `CNSC-130`
**Originating Document**:
- `Cloud Native Security Whitepaper v1.0`
- `Cloud Native Security Whitepaper v2.0`



**Objective**

Data backup storage locations employ like access controls and security policies to that of the data storage source

**Guideline Mappings**

**`NIST-800-53`**

| NIST control | Remarks |
|--------------|---------|
| `SA-9`<br />`SC-30` | <ul><li>External System Services</li><li>Concealment and Misdirection</li></ul> |


---



<a id="cnsc-131"></a>

### System Backup | `CNSC-131`
**Originating Document**:
- `Cloud Native Security Whitepaper v1.0`
- `Cloud Native Security Whitepaper v2.0`



**Objective**

Secure erasure adhering to OPAL standards is employed for returned or non-functional devices

**Guideline Mappings**

**`NIST-800-53`**

| NIST control | Remarks |
|--------------|---------|
| `CP-9`<br />`MP-6` | <ul><li>System Backup</li><li>Media Sanitization</li></ul> |


---



<a id="cnsc-132"></a>

### Encryption of Data at Rest | `CNSC-132`
**Originating Document**:
- `Cloud Native Security Whitepaper v1.0`
- `Cloud Native Security Whitepaper v2.0`



**Objective**

Encryption at rest considers data path, size, and frequency of access when determining additional security protections and cryptographic algorithms to employ

**Guideline Mappings**

**`NIST-800-53`**

| NIST control | Remarks |
|--------------|---------|
| `SC-28` | Protection of Information at Rest |


---



<a id="cnsc-133"></a>

### Encryption Requirements | `CNSC-133`
**Originating Document**:
- `Cloud Native Security Whitepaper v1.0`
- `Cloud Native Security Whitepaper v2.0`



**Objective**

Caching is considered for determining encryption requirements in architectures

*No NIST SP 800-53 mapping row in this document for this guideline.*

---



<a id="cnsc-134"></a>

### Boundary Protection | `CNSC-134`
**Originating Document**:
- `Cloud Native Security Whitepaper v1.0`
- `Cloud Native Security Whitepaper v2.0`



**Objective**

Namespaces have defined trust boundaries to cordon access to volumes

*No NIST SP 800-53 mapping row in this document for this guideline.*

---



<a id="cnsc-135"></a>

### Security Policy Management | `CNSC-135`
**Originating Document**:
- `Cloud Native Security Whitepaper v1.0`
- `Cloud Native Security Whitepaper v2.0`



**Objective**

Security policies are used to prevent containers from accessing volume mounts on worker nodes

**Guideline Mappings**

**`NIST-800-53`**

| NIST control | Remarks |
|--------------|---------|
| `CM-6`<br />`SA-8`<br />`SC-7` | <ul><li>Configuration Settings</li><li>Security and Privacy Engineering Principles</li><li>Boundary Protection</li></ul> |


---



<a id="cnsc-136"></a>

### Security Policy Enforcement | `CNSC-136`
**Originating Document**:
- `Cloud Native Security Whitepaper v1.0`
- `Cloud Native Security Whitepaper v2.0`



**Objective**

Security policies are used enforce authorized worker node access to volumes

**Guideline Mappings**

**`NIST-800-53`**

| NIST control | Remarks |
|--------------|---------|
| `CM-6`<br />`SA-8`<br />`SC-7` | <ul><li>Configuration Settings</li><li>Security and Privacy Engineering Principles</li><li>Boundary Protection</li></ul> |


---



<a id="cnsc-137"></a>

### Information Flow Management | `CNSC-137`
**Originating Document**:
- `Cloud Native Security Whitepaper v1.0`
- `Cloud Native Security Whitepaper v2.0`



**Objective**

Volume UID and GID are inaccessible to containers

**Guideline Mappings**

**`NIST-800-53`**

| NIST control | Remarks |
|--------------|---------|
| `AC-16`<br />`AC-4`<br />`SI-7` | <ul><li>Security and Privacy Attributes</li><li>Information Flow Enforcement</li><li>Software, Firmware, and Information Integrity</li></ul> |


---



<a id="cnsc-138"></a>

### Artifact Registry Management | `CNSC-138`
**Originating Document**:
- `Cloud Native Security Whitepaper v1.0`
- `Cloud Native Security Whitepaper v2.0`



**Objective**

Artifact registry supports OCI artifacts

*No NIST SP 800-53 mapping row in this document for this guideline.*

---



<a id="cnsc-139"></a>

### Signed Artifact Support | `CNSC-139`
**Originating Document**:
- `Cloud Native Security Whitepaper v1.0`
- `Cloud Native Security Whitepaper v2.0`



**Objective**

Artifact registry supports signed artifacts

**Guideline Mappings**

**`NIST-800-53`**

| NIST control | Remarks |
|--------------|---------|
| `CM-14` | Signed Components |


---



<a id="cnsc-140"></a>

### Artifact Registry Policy Verification | `CNSC-140`
**Originating Document**:
- `Cloud Native Security Whitepaper v1.0`
- `Cloud Native Security Whitepaper v2.0`



**Objective**

Artifact registry verifies artifacts against organizational policies

**Guideline Mappings**

**`NIST-800-53`**

| NIST control | Remarks |
|--------------|---------|
| `AU-10`<br />`CM-6` | <ul><li>Non-repudiation</li><li>Configuration Settings</li></ul> |


---



