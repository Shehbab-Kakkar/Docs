Docker provides several types of networking drivers to manage how containers communicate with each other, the host, and external networks. Here are the main types:

---

## 1. **Bridge Network (default)**
- **Description:**  
  The default network driver for containers on a single host. Containers get an internal IP and can communicate with each other.
- **Use Case:**  
  Standalone applications on a single host.
- **Command:**  
  Docker creates a bridge network named `bridge` automatically.
- **Example:**  
  `docker run --network bridge ...`

---

## 2. **Host Network**
- **Description:**  
  Removes network isolation between the container and the host. The container shares the host’s network stack.
- **Use Case:**  
  Performance-sensitive applications, or when you need direct access to host networking.
- **Command:**  
  `docker run --network host ...`

---

## 3. **None Network**
- **Description:**  
  Disables all networking for the container. The container has no access to external networks or even other containers.
- **Use Case:**  
  Isolation for security or testing.
- **Command:**  
  `docker run --network none ...`

---

## 4. **Overlay Network**
- **Description:**  
  Enables multi-host networking; containers on different Docker hosts can communicate securely.
- **Use Case:**  
  Docker Swarm or multi-host orchestration.
- **Command:**  
  `docker network create -d overlay my_overlay`

---

## 5. **Macvlan Network**
- **Description:**  
  Assigns a MAC address to a container, making it appear as a physical device on the network.
- **Use Case:**  
  When containers need to be directly accessible on the physical network, e.g., for legacy applications.
- **Command:**  
  `docker network create -d macvlan ...`

---

## 6. **IPvlan Network**
- **Description:**  
  Similar to macvlan but allows more advanced IP address management.
- **Use Case:**  
  When you need finer control over IP assignment.
- **Command:**  
  `docker network create -d ipvlan ...`

---

## 7. **Custom Networks**
- **Description:**  
  User-defined networks (bridge or overlay) for better isolation and network control.
- **Use Case:**  
  Isolating groups of containers, customizing DNS, etc.
- **Command:**  
  `docker network create my_custom_net`

---

## **Summary Table**

| Network Type   | Scope        | Use Case                              | Isolation | Inter-host Communication |
|----------------|--------------|---------------------------------------|-----------|-------------------------|
| Bridge         | Single host  | Default, small setups                 | Yes       | No                      |
| Host           | Single host  | Performance, host-level apps          | No        | No                      |
| None           | Single host  | Security, full isolation              | Yes       | No                      |
| Overlay        | Multi-host   | Swarm, multi-host clusters            | Yes       | Yes                     |
| Macvlan        | Single/multi | Direct access to physical network     | Yes       | Yes (with setup)        |
| IPvlan         | Single/multi | Advanced IP management                | Yes       | Yes (with setup)        |

---

**In summary:**  
Docker networking includes bridge, host, none, overlay, macvlan, ipvlan, and custom networks. Each serves a different use case for container communication, isolation, and connectivity, both within a single host or across multiple hosts.
