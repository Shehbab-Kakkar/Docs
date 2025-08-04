# OpenShift SDN in Latest OpenShift Versions

This document provides information on the Software Defined Networking (SDN) used by the latest versions of the OpenShift Container Platform.

---

## Table of Contents

- [Overview](#overview)
- [Default SDN in Latest OpenShift Versions](#default-sdn-in-latest-openshift-versions)
- [OVN-Kubernetes vs. OpenShift SDN](#ovn-kubernetes-vs-openshift-sdn)
- [Migration and Compatibility](#migration-and-compatibility)
- [References](#references)

---

## Overview

OpenShift relies on a Software Defined Networking (SDN) layer to provide secure, scalable, and high-performance networking for Kubernetes workloads. In earlier versions, the default SDN was **OpenShift SDN**. However, with the evolution of the platform, Red Hat now recommends a new SDN for modern OpenShift clusters.


<p>OVN stands for Open Virtual Network. It is an open source network virtualization project that provides network automation and advanced networking features for virtualized environments, and it is now the foundation of the default SDN (Software Defined Networking) solution in modern OpenShift versions.</p>
---

## Default SDN in Latest OpenShift Versions

As of OpenShift 4.16 and later (2025):

- **OVN-Kubernetes** is the default and recommended SDN for all new OpenShift Container Platform installations.
- **OpenShift SDN** is deprecated and only supported for compatibility with existing clusters that originally used it.

---

## OVN-Kubernetes vs. OpenShift SDN

| Feature                | OVN-Kubernetes            | OpenShift SDN (Deprecated) |
|------------------------|--------------------------|----------------------------|
| Default for new installs | Yes                    | No                         |
| Actively developed     | Yes                      | No                         |
| Advanced networking policies | Yes              | Limited                    |
| IPv6 dual-stack support| Yes                      | No                         |
| Recommended by Red Hat | Yes                      | No                         |

**Note:** OVN-Kubernetes provides advanced network policy support, better scalability, IPv6, and future-proofing for OpenShift environments.

---

## Migration and Compatibility

- New OpenShift clusters (4.16+) use OVN-Kubernetes by default.
- OpenShift SDN remains available only for clusters upgraded from earlier installs.
- Customers are encouraged to migrate to OVN-Kubernetes, as OpenShift SDN will be removed in future major releases.

---

## References

- For official guidance and the latest recommendations, refer to the Red Hat OpenShift release notes and documentation.
- See the OpenShift and OVN-Kubernetes migration guides for step-by-step instructions.

---

**In summary:**  
For any new OpenShift deployment in 2025 and beyond, OVN-Kubernetes is the out-of-the-box SDN, with OpenShift SDN deprecated and not recommended for new clusters.
