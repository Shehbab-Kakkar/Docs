Create an ExternalName service YAML example

```yaml
apiVersion: v1
kind: Service
metadata:
  name: my-external-db
  namespace: default
spec:
  type: ExternalName
  externalName: db.example.com
```
This YAML creates an ExternalName service named `my-external-db` that maps to the external DNS `db.example.com` via CNAME record.[11]
Apply with `kubectl apply -f externalname.yaml`.[1][5]

## Adding Ports
```yaml
spec:
  type: ExternalName
  externalName: db.example.com
  ports:
  - port: 3306
    targetPort: 3306
    protocol: TCP
```

I
