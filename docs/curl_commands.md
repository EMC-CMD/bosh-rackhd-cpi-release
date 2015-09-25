## Tasks and workflows

### Uploading task
```
curl -X PUT -H "Content-Type: application/json" --data @src/github.com/onrack/onrack-cpi/templates/wipe_machine_task.json  "http://$ON_RACK_API_URI:8080/api/common/workflows/tasks"
```

### Uploading workflow
```
curl -X PUT -H "Content-Type: application/json" --data @../templates/create_vm_workflow.json  "http://$ON_RACK_API_URI:8080/api/common/workflows"
```
### Finding task by name
```
curl http://$ON_RACK_API_URI:8080/api/common/workflows/tasks/library | jq '.[] | select(.injectableName == "Task.Os.Delete.CF.VM")'
```

### Finding workflow by name
```
curl http://$ON_RACK_API_URI:8080/api/common/workflows/library | jq '.[] | select(.injectableName == "Graph.CF.DeleteVM")'
```

### Checking status of workflows
```
curl http://$ON_RACK_API_URI:8080/api/common/nodes/55e79ea54e66816f6152fff9/workflows/active | jq .
```

### Watch workflow status
```
watch -c -n 10 'curl http://$ON_RACK_API_URI:8080/api/common/nodes/55e79eb14e66816f6152fffb/workflows/active | jq ". | ._status"'
```

### Canceling active workflow
```
curl -X DELETE http://$ON_RACK_API_URI:8080/api/common/nodes/55e79ea54e66816f6152fff9/workflows/active
```
### Submitting CreateVM workflow
```
curl -X POST -H "Content-Type: application/json" http://$ON_RACK_API_URI:8080/api/common/nodes/55e79eb14e66816f6152fffb/workflows -d '{"name":"Graph.CF.CreateReserveVM","options":{"defaults": {"agentSettingsFile": "env-234567", "agentSettingsPath": "/var/vcap/bosh/baremetal-cpi-agent-env.json", "cid": "vm-1234","downloadDir": "/opt/downloads","registrySettingsFile": "agent-123456", "registrySettingsPath": "/var/vcap/bosh/agent.json", "stemcellFile": "raw-image"}}}'
```

## Node Handling

### Running workflow against node
```
curl -X POST -H "Content-Type: application/json" http://$ON_RACK_API_URI:8080/api/common/nodes/55e79ea54e66816f6152fff9/workflows -d '{"name":"Graph.CF.DeleteVM","options":{}}'
```

### Getting all nodes
```
curl "http://$ON_RACK_API_URI:8080/api/common/nodes" | jq .
```

## Files Handling

### Getting all Files
```
curl "http://$ON_RACK_API_URI:8080/api/common/files/list/all" | jq .
```

### Uploading Files
```
curl -X PUT "http://$ON_RACK_API_URI:8080/api/common/files/env-234567" --upload-file {a_file}
```

### Deleting Files
```
curl -X DELETE "http://$ON_RACK_API_URI:8080/api/common/files/{a_file_uuid}"
```