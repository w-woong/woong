# Curl

```bash
# Create AppConfig
curl --insecure -H "Content-Type: application/json; charset=utf-8" \
-X POST \
-H 'Authorization: Bearer ab2316584873095f017f6dfa7a9415794f563fcc473eb3fe65b9167e37fd5a4b' \
-d '{"document":{"id":"b69aa108-f12e-4fa0-bf4f-ba002c11a670","name":"Woong"}}' \
https://localhost:49001/v1/woong/appconfig

# Create Home
curl --insecure -H "Content-Type: application/json; charset=utf-8" \
-X POST \
-H 'Authorization: Bearer ab2316584873095f017f6dfa7a9415794f563fcc473eb3fe65b9167e37fd5a4b' \
-d '{"document":{"app_config_id":"b69aa108-f12e-4fa0-bf4f-ba002c11a670","name":"Home"}}' \
https://localhost:49001/v1/woong/home



curl --insecure -H "Content-Type: application/json; charset=utf-8" \
-X GET \
-H 'Authorization: Bearer ab2316584873095f017f6dfa7a9415794f563fcc473eb3fe65b9167e37fd5a4b' \
http://localhost:49001/v1/woong/appconfig/b69aa108-f12e-4fa0-bf4f-ba002c11a670

curl --insecure -H "Content-Type: application/json; charset=utf-8" \
-X GET \
-H 'Authorization: Bearer ab2316584873095f017f6dfa7a9415794f563fcc473eb3fe65b9167e37fd5a4b' \
http://localhost:49001/v1/woong/home/appconfig/b69aa108-f12e-4fa0-bf4f-ba002c11a670
```