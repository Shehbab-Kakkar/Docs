
# This is not readme.md format it is showing me in the text only. I want o add in the github and command part in the coding format in the github readme.

Absolutely! Below is the README.md with **proper Markdown formatting** for GitHub, where all AWS CLI commands are in code blocks (fenced with triple backticks and labeled with `bash` for syntax highlighting). Headings, sections, and notes are formatted for clear GitHub presentation.

# On-Prem to AWS Migration: CLI Guide

This document details the **step-by-step migration of on-prem applications to AWS using AWS CLI**. All commands are **properly fenced in Markdown code blocks** and ready for GitHub.

## Table of Contents

- [MySQL Migration (Amazon RDS)](#mysql-migration-amazon-rds)
- [Java Spring Boot Deployment (Elastic Beanstalk)](#java-spring-boot-deployment-elastic-beanstalk)
- [RabbitMQ Migration (Amazon MQ)](#rabbitmq-migration-amazon-mq)
- [Application Servers (EC2)](#application-servers-ec2)
- [Load Balancer Setup (ALB)](#load-balancer-setup-alb)
- [Security \& Networking](#security--networking)
- [Monitoring (CloudWatch)](#monitoring-cloudwatch)
- [DNS Cutover (Route 53)](#dns-cutover-route-53)
- [Tips \& Notes](#tips--notes)


## MySQL Migration (Amazon RDS)

**Create RDS MySQL Instance**

```bash
aws rds create-db-instance \
  --db-instance-identifier mymysql-db \
  --db-instance-class db.t3.medium \
  --engine mysql \
  --allocated-storage 50 \
  --master-username admin \
  --master-user-password <password> \
  --vpc-security-group-ids <sg-12345678>
```

**Enable Multi-AZ Deployment**

```bash
aws rds modify-db-instance \
  --db-instance-identifier mymysql-db \
  --multi-az
```

**(Optional) Create AWS DMS Replication Instance**

```bash
aws dms create-replication-instance \
  --replication-instance-identifier my-dms-instance \
  --replication-instance-class dms.t3.medium \
  --allocated-storage 50
```

*Set up source/target endpoints and tasks as required for full data migration.*

## Java Spring Boot Deployment (Elastic Beanstalk)

**Create Beanstalk Application**

```bash
aws elasticbeanstalk create-application --application-name my-springboot-app
```

**Create Beanstalk Environment with Docker Support**

```bash
aws elasticbeanstalk create-environment \
  --application-name my-springboot-app \
  --environment-name my-env \
  --solution-stack-name "64bit Amazon Linux 2 v3.4.17 running Docker" \
  --version-label v1 \
  --option-settings file://eb-options.json
```

*`eb-options.json` should define configuration for VPC, env variables, etc.*

**Update Beanstalk Environment (for New Deployments)**

```bash
aws elasticbeanstalk update-environment \
  --environment-name my-env \
  --version-label v2
```


## RabbitMQ Migration (Amazon MQ)

**Create RabbitMQ Broker**

```bash
aws mq create-broker \
  --broker-name my-rabbitmq-broker \
  --engine-type RABBITMQ \
  --engine-version 3.11.16 \
  --host-instance-type mq.m5.large \
  --deployment-mode SINGLE_INSTANCE \
  --publicly-accessible \
  --users '[{"username":"admin","password":"<password>"}]'
```


## Application Servers (EC2)

**Launch EC2 Application Server**

```bash
aws ec2 run-instances \
  --image-id <ami-id> \
  --count 1 \
  --instance-type t3.medium \
  --key-name <ec2-keypair> \
  --security-group-ids <sg-1234567890> \
  --subnet-id <subnet-123456>
```


## Load Balancer Setup (ALB)

**Create Application Load Balancer**

```bash
aws elbv2 create-load-balancer \
  --name my-alb \
  --subnets <subnet-1> <subnet-2> \
  --security-groups <sg-123456>
```

**Create Target Group**

```bash
aws elbv2 create-target-group \
  --name my-targets \
  --protocol HTTP \
  --port 8080 \
  --vpc-id <vpc-12345>
```

**Register EC2 Instance as Target**

```bash
aws elbv2 register-targets \
  --target-group-arn <target-group-arn> \
  --targets Id=<instance-id>
```

**Create Listener for Load Balancer**

```bash
aws elbv2 create-listener \
  --load-balancer-arn <alb-arn> \
  --protocol HTTP \
  --port 80 \
  --default-actions Type=forward,TargetGroupArn=<target-group-arn>
```


## Security \& Networking

**Create Security Group**

```bash
aws ec2 create-security-group \
  --group-name my-sg \
  --description "My app security group" \
  --vpc-id <vpc-id>
```

**Authorize Ingress to Security Group**

```bash
aws ec2 authorize-security-group-ingress \
  --group-id <sg-id> \
  --protocol tcp \
  --port 8080 \
  --cidr 0.0.0.0/0
```


## Monitoring (CloudWatch)

**Create CloudWatch Alarm**

```bash
aws cloudwatch put-metric-alarm \
  --alarm-name "HighCPUUtilization" \
  --metric-name CPUUtilization \
  --namespace AWS/EC2 \
  --statistic Average \
  --period 300 \
  --threshold 80 \
  --comparison-operator GreaterThanThreshold \
  --dimensions Name=InstanceId,Value=<instance-id> \
  --evaluation-periods 2 \
  --alarm-actions <sns-topic-arn>
```


## DNS Cutover (Route 53)

**Update DNS Record**

```bash
aws route53 change-resource-record-sets \
  --hosted-zone-id <zoneid> \
  --change-batch file://dns-change.json
```

*`dns-change.json` should define your A, CNAME, or other record changes.*

## Tips \& Notes

- Adapt `<...>` placeholders for your environment (IDs, passwords, subnets).
- Use scripting or CI/CD for automation/repeatability.
- Security: Never hard-code secrets in scripts.
- For production, consider Infrastructure-as-Code (CloudFormation/Terraform).
- **Test in sandbox environments before moving to production.**

**Happy migrating! Fork this for your DevOps or Cloud teams and customize as needed!**

> *Markdown fenced code blocks (````bash`) ensure formatting displays as code on GitHub. Use headings, lists, and spacing to keep your README easy to follow.*

