
---

## ✅ Ansible & AWS Questions with Answers

---

### **Q1: In Ansible, how can you restart a service after the 6th task without adding conditions to those tasks?**

**Answer:**
You can simply place the restart task as the 7th task in the playbook. There’s no need to add conditionals. Example:

```yaml
- name: Restart service after 6th task
  ansible.builtin.service:
    name: myservice
    state: restarted
```

---

### **Q2: You have 50 tasks in a playbook. After successfully running the first 25, the playbook failed. How can you run only the remaining tasks without re-running the first 25?**

**Answer:**
Use the `--start-at-task` flag with the name of the 26th task:

```bash
ansible-playbook playbook.yml --start-at-task="Task 26 Name"
```

You can also use `--tags` if tasks are tagged, or `--step` to manually approve tasks.

---

### **Q3: What is the use of `--step` in Ansible?**

**Answer:**
The `--step` flag runs the playbook in interactive mode, allowing you to execute or skip each task manually.

---

### **Q4: How do you notify a handler in Ansible?**

**Answer:**
You add a `notify` directive to a task. Handlers run **only if notified**.

```yaml
- name: Install package
  apt:
    name: nginx
    state: present
  notify: restart nginx

handlers:
  - name: restart nginx
    service:
      name: nginx
      state: restarted
```

---

### **Q5: In Ansible, how can you ensure a task only runs if the previous task made a change?**

**Answer:**
Use **handlers**, or `register` with `when`:

```yaml
- name: Update config file
  template:
    src: config.j2
    dest: /etc/app/config
  register: config_status

- name: Restart app if config changed
  service:
    name: app
    state: restarted
  when: config_status.changed
```

---

### **Q6: What is the difference between `state: restarted` and `state: reloaded` in Ansible service module?**

**Answer:**

* `restarted`: Stops and starts the service. Used for major changes.
* `reloaded`: Applies configuration without full restart (if supported by the service).

---

### **Q7: How do you manage AWS resources using Ansible?**

**Answer:**
Use the **`amazon.aws` or `community.aws`** collections and modules like:

* `ec2_instance`
* `s3_bucket`
* `rds_instance`

Install with:

```bash
ansible-galaxy collection install amazon.aws
```

---

### **Q8: How can you create an EC2 instance in Ansible?**

**Answer:**
Use `amazon.aws.ec2_instance`:

```yaml
- name: Launch EC2
  amazon.aws.ec2_instance:
    name: my-instance
    key_name: my-key
    instance_type: t2.micro
    image_id: ami-12345678
    region: us-east-1
    wait: yes
```

---

### **Q9: What is the purpose of Ansible `register`?**

**Answer:**
`register` captures the output/result of a task so it can be used in `when` conditions or later tasks.

---

### **Q10: What’s the difference between `notify` and `include_tasks` in Ansible?**

**Answer:**

* `notify` triggers **handlers** conditionally (only on change).
* `include_tasks` includes and runs another task file **unconditionally**.

---

If you want **more AWS-specific** or **Ansible troubleshooting** questions, let me know your focus (e.g., EC2, networking, playbooks, IAM), and I can tailor it.
