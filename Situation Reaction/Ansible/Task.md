To **restart a service after the 6th task in Ansible**, and **not mix the restart logic into the tasks themselves**, you can **cleanly separate** the service restart as the **7th task** in your playbook. This does not require any conditionals or logic inside the first 6 tasks.

Here’s a simple example of what you might want to do:

```yaml
- name: My Playbook
  hosts: all
  become: yes

  tasks:
    - name: Task 1
      debug:
        msg: "This is task 1"

    - name: Task 2
      debug:
        msg: "This is task 2"

    - name: Task 3
      debug:
        msg: "This is task 3"

    - name: Task 4
      debug:
        msg: "This is task 4"

    - name: Task 5
      debug:
        msg: "This is task 5"

    - name: Task 6
      debug:
        msg: "This is task 6"

    # Restart service here, after task 6
    - name: Restart myservice
      ansible.builtin.service:
        name: myservice
        state: restarted

    - name: Task 8
      debug:
        msg: "This is task 8"

    - name: Task 9
      debug:
        msg: "This is task 9"

    - name: Task 10
      debug:
        msg: "This is task 10"
```

### Key Points:

* The **restart is its own task** — no need to insert conditions into the earlier ones.
* It **executes immediately after task 6**.
* You keep **clear separation of concerns** — tasks do what they need to do, and service management is centralized.

