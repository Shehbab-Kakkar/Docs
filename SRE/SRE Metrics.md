  # SRE Metrics: SLI, SLO, SLA, and MTTR Explained

This document explains the four key Site Reliability Engineering (SRE) terms: **SLI**, **SLO**, **SLA**, and **MTTR** — what they mean, why they matter, and how they relate, both in technical systems and with a real-world pizza delivery analogy.

---

## 1️⃣ SLI – Service Level Indicator

- **Definition:**  
  A quantitative measure of some aspect of the level of service provided.

- **Examples:**
  - Request success rate (e.g., 99.9%)
  - Latency (e.g., 95% of requests under 300ms)
  - Availability (e.g., service up 99.95% of the time)

- **Purpose:**  
  Used to measure how well a service is performing.

---

## 2️⃣ SLO – Service Level Objective

- **Definition:**  
  A target value or range for a service level that is measured by an SLI.

- **Examples:**
  - “We aim for 99.99% availability over a 30-day period.”
  - “We aim for 95% of requests to complete in under 200ms.”

- **Purpose:**  
  A goal to guide reliability efforts. If the service is outside the SLO, it's considered not meeting expectations.

---

## 3️⃣ SLA – Service Level Agreement

- **Definition:**  
  A formal agreement between a service provider and a customer that outlines the expected level of service, often with penalties for failing to meet it.

- **Examples:**
  - An SLA might guarantee 99.9% uptime and include refunds or service credits if that’s not met.

- **Purpose:**  
  A legal or contractual commitment, often derived from the SLO but with more conservative guarantees.

---

## 4️⃣ MTTR – Mean Time to Recovery/Repair

- **Definition:**  
  The average time it takes to recover from a failure or incident.

- **Formula:**  
  ```
  MTTR = Total downtime / Number of incidents
  ```

- **Purpose:**  
  Indicates how quickly your team can restore service after an outage.

---

## 🍕 Pizza Delivery Analogy

| Term | Pizza Analogy                                            |
|------|---------------------------------------------------------|
| **SLI** | The metric: how many pizzas arrived hot (indicator)       |
| **SLO** | Goal: 99% of pizzas should arrive hot                      |
| **SLA** | Contract: If less than 95% arrive hot, customer gets a refund |
| **MTTR** | If oven breaks, how long does it take to fix it?            |

---

## 📝 Quick Reference Table

| Term  | Stands For                     | What It Is                       | Example                         |
|-------|--------------------------------|----------------------------------|---------------------------------|
| SLI   | Service Level Indicator        | Measurement                      | 99.9% request success rate      |
| SLO   | Service Level Objective        | Target/Goal                      | 99.99% availability in 30 days  |
| SLA   | Service Level Agreement        | Contractual/Legal Commitment     | 99.9% uptime or refund          |
| MTTR  | Mean Time to Recovery/Repair   | Recovery Speed Metric            | 30 minutes average to recover   |

---

## 📚 References

- [Google SRE Book: Service Level Objectives](https://sre.google/sre-book/service-level-objectives/)
- [Wikipedia: Service-level agreement](https://en.wikipedia.org/wiki/Service-level_agreement)
- [Wikipedia: Mean time to repair](https://en.wikipedia.org/wiki/Mean_time_to_repair)

---
