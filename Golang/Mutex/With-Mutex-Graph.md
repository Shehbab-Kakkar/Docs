Great question!

Let’s now explore and **compare how the value of `views` changes over time across goroutine executions**:

* ✅ **With Mutex** – ensures correct updates: **no missed increments**
* ❌ **Without Mutex** – leads to race conditions: **some increments are lost**

---

## 🧮 **What is a Value Graph?**

A **value graph** in this context tracks the **value of `views` over time**, as each goroutine runs.

We’ll compare:

1. ✅ **With Mutex** — every increment is isolated
2. ❌ **Without Mutex** — concurrent increments collide

---

## ✅ **With Mutex – Safe, Predictable Growth**

### Assumptions:

* 100 goroutines
* Each increments by 1
* Mutex locks and unlocks ensure **no overlapping**

### Value Graph (With Mutex)

| Goroutine | `views` Value After Execution |
| --------- | ----------------------------- |
| G1        | 1                             |
| G2        | 2                             |
| G3        | 3                             |
| ...       | ...                           |
| G100      | 100                           |

### Graph (With Mutex)

```
views
  ^
100|                                                           ●
 90|                                                    ●
 80|                                             ●
 70|                                      ●
 60|                               ●
 50|                        ●
 40|                 ●
 30|          ●
 20|    ●
 10| ●
 0 +---------------------------------------------------------> goroutines
   0     10    20    30    40    50    60    70    80    90   100
```

✅ **Perfect linear growth** — every increment is preserved.

---

## ❌ **Without Mutex – Unsafe, Inconsistent Growth**

### Assumptions:

* 100 goroutines
* Some race conditions cause overwrites
* Let's say \~10 updates are lost (real number varies per run)

### Value Graph (Without Mutex)

| Goroutine | `views` Value After Execution (Approximate) |
| --------- | ------------------------------------------- |
| G1        | 1                                           |
| G2        | 2                                           |
| G3        | 2  ← collision, lost increment              |
| G4        | 3                                           |
| G5        | 4                                           |
| G6        | 4  ← collision again                        |
| ...       | ...                                         |
| G100      | \~90                                        |

> ⚠️ NOTE: This is **non-deterministic** – results vary on each run!

### Graph (Without Mutex – Example Run)

```
views
  ^
100|                                                          
 90|                                                     ●
 80|                                                ●
 70|                                           ●
 60|                                     ●
 50|                                ●
 40|                          ●
 30|                    ●
 20|              ●
 10|        ●
 0 +---------------------------------------------------------> goroutines
   0     10    20    30    40    50    60    70    80    90   100
```

❌ **Jagged, unpredictable growth** – some updates are lost.

---

## ⚖️ Final Comparison

| Feature             | With Mutex 🟢          | Without Mutex 🔴      |
| ------------------- | ---------------------- | --------------------- |
| Views Growth        | Linear and predictable | Random, non-linear    |
| Final `views` Value | Always 100             | Usually < 100         |
| Race Conditions     | Prevented              | Likely                |
| Thread Safety       | Guaranteed             | None                  |
| Bug Reproduction    | None                   | Hard and inconsistent |

---

## 🧠 Summary

* **With Mutex**: Each goroutine locks before incrementing `views`, ensuring that **no two goroutines interfere** with each other. Result: `views == 100`
* **Without Mutex**: Multiple goroutines may read and write to `views` **simultaneously**, leading to **lost updates**. Result: `views < 100`, varies per run.

---

Would you like me to **generate a real-time visual plot/image** for both graphs to help illustrate this?
